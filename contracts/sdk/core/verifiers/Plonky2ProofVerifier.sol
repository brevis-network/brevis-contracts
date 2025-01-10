// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @title Groth16 verifier template.
/// @author Remco Bloemen
/// @notice Supports verifying Groth16 proofs. Proofs can be in uncompressed
/// (256 bytes) and compressed (128 bytes) format. A view function is provided
/// to compress proofs.
/// @notice See <https://2π.com/23/bn254-compression> for further explanation.
contract Plonky2ProofVerifier {

    /// Some of the provided public input values are larger than the field modulus.
    /// @dev Public input elements are not automatically reduced, as this is can be
    /// a dangerous source of bugs.
    error PublicInputNotInField();

    /// The proof is invalid.
    /// @dev This can mean that provided Groth16 proof points are not on their
    /// curves, that pairing equation fails, or that the proof is not for the
    /// provided public input.
    error ProofInvalid();

    // Addresses of precompiles
    uint256 constant PRECOMPILE_MODEXP = 0x05;
    uint256 constant PRECOMPILE_ADD = 0x06;
    uint256 constant PRECOMPILE_MUL = 0x07;
    uint256 constant PRECOMPILE_VERIFY = 0x08;

    // Base field Fp order P and scalar field Fr order R.
    // For BN254 these are computed as follows:
    //     t = 4965661367192848881
    //     P = 36⋅t⁴ + 36⋅t³ + 24⋅t² + 6⋅t + 1
    //     R = 36⋅t⁴ + 36⋅t³ + 18⋅t² + 6⋅t + 1
    uint256 constant P = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47;
    uint256 constant R = 0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001;

    uint256 constant MOD_R = 21888242871839275222246405745257275088548364400416034343698204186575808495617;

    // Extension field Fp2 = Fp[i] / (i² + 1)
    // Note: This is the complex extension field of Fp with i² = -1.
    //       Values in Fp2 are represented as a pair of Fp elements (a₀, a₁) as a₀ + a₁⋅i.
    // Note: The order of Fp2 elements is *opposite* that of the pairing contract, which
    //       expects Fp2 elements in order (a₁, a₀). This is also the order in which
    //       Fp2 elements are encoded in the public interface as this became convention.

    // Constants in Fp
    uint256 constant FRACTION_1_2_FP = 0x183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4;
    uint256 constant FRACTION_27_82_FP = 0x2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5;
    uint256 constant FRACTION_3_82_FP = 0x2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775;

    // Exponents for inversions and square roots mod P
    uint256 constant EXP_INVERSE_FP = 0x30644E72E131A029B85045B68181585D97816A916871CA8D3C208C16D87CFD45; // P - 2
    uint256 constant EXP_SQRT_FP = 0xC19139CB84C680A6E14116DA060561765E05AA45A1C72A34F082305B61F3F52; // (P + 1) / 4;

    // Groth16 alpha point in G1
    uint256 constant ALPHA_X = 18465672538572564110654870691316325222021992541514285645578648286480786538043;
    uint256 constant ALPHA_Y = 21318369084500624145230170437471414962110739155584103631872407006931000719081;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 1040934873188637874721541777395400134344790444720280161296646593532173730879;
    uint256 constant BETA_NEG_X_1 = 21203149842354738305675940308731461414235117610508844826685244696613144395114;
    uint256 constant BETA_NEG_Y_0 = 20076936249467790883877454784404644422501896758885419229097224928738792835265;
    uint256 constant BETA_NEG_Y_1 = 2459199669942949699053072822108300156641275279832963140959067874708576299125;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 6059673474152865292208372428057277555067191546992828187163075795034387982684;
    uint256 constant GAMMA_NEG_X_1 = 2200677557695067462945195597720996649703788039350802315018332453212473743869;
    uint256 constant GAMMA_NEG_Y_0 = 10092241198505883499555652859929762413882464091119020398047530381285943984786;
    uint256 constant GAMMA_NEG_Y_1 = 21286073834610239846841985325799339034585228067529718636516495411173134569066;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 5344071998990808677332281293642118919337725540233985062631787460694689702395;
    uint256 constant DELTA_NEG_X_1 = 11901471681027137046078343917837675565067130452319830767756899273182938364194;
    uint256 constant DELTA_NEG_Y_0 = 7834055415058692361827881253540091609941628412960888105718454960822411612662;
    uint256 constant DELTA_NEG_Y_1 = 17720506151939299472218543306401068259753543292713438528751656406999736440972;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 6393803814431136249711904496147113806772225217875498274348077568022216966194;
    uint256 constant PEDERSEN_G_X_1 = 18223629497189435754374964942736898365918452931244862664751655427973826729520;
    uint256 constant PEDERSEN_G_Y_0 = 4235628718310107962325375263455966483914752227291678083549466575299923769953;
    uint256 constant PEDERSEN_G_Y_1 = 12551152224778183746967252297709784257390971681317025530621998697583948386128;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 = 6089982018257586125691546818153485831256710272353810737912324424487892997515;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 = 12836689740842574789368138826839755411670026694279368281504468294735556334751;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 = 3444632640261882176700784395590915819299342807262587376710639596072980388912;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 = 12266165675992077611702917298065019238899185575482686701882521777742237843663;

    // Constant and public input points
    uint256 constant CONSTANT_X = 17842415870269424190244400246372136209419975945593128523349246584337781730921;
    uint256 constant CONSTANT_Y = 7042238044267267264637487016795321217401540174665380877531022979957498020275;
    uint256 constant PUB_0_X = 16186693208175688567296675061429980671337849774679408148730224084838727627600;
    uint256 constant PUB_0_Y = 5564967360710461738939137212844995628073613866865372165506267691893772378435;
    uint256 constant PUB_1_X = 19185691551367161844270253358356127172363397946230066102832345592912265185969;
    uint256 constant PUB_1_Y = 21397499293375596041265394305640320941421652533837157916717221810741203924643;
    uint256 constant PUB_2_X = 13680609380625884427338478352628435922230060588270038714460114492457055869969;
    uint256 constant PUB_2_Y = 48725041707931937228479085859534563226932914277863536072220894623917783412;
    uint256 constant PUB_3_X = 19243250727887833276140842951424478634647586685361821072623494949057503753941;
    uint256 constant PUB_3_Y = 3036130038910477094393204337949754078465272361484598360854466760835948458284;
    uint256 constant PUB_4_X = 8793272448675004869888301961411226199340935223921367630652069693924897208752;
    uint256 constant PUB_4_Y = 3947661938373102982350422975747175921136376017871457172785914879174365403307;
    uint256 constant PUB_5_X = 7198341914176272591251869195246140711855892562238811323079671329342239072457;
    uint256 constant PUB_5_Y = 14449488327536499440225762752769824529169218981356146941799306752635876045236;
    uint256 constant PUB_6_X = 12320246081351483278216357874148389582869288963342115537645502890238607390060;
    uint256 constant PUB_6_Y = 5253944111384833810814770597929372170644460146257218871826721291539451400428;
    uint256 constant PUB_7_X = 17121930814812207462294183921742339351989330337653860264078407653837592749835;
    uint256 constant PUB_7_Y = 13347616006694463130092922482964418127787237454760678331769101530384626181776;
    uint256 constant PUB_8_X = 5833718279367228660795785901405569878653262736056782039772801635723483987970;
    uint256 constant PUB_8_Y = 13564180361743660497577365074783208480423846294458126860915870662201881063244;

    /// Compute the public input linear combination.
    /// @notice Reverts with PublicInputNotInField if the input is not in the field.
    /// @notice Computes the multi-scalar-multiplication of the public input
    /// elements and the verification key including the constant term.
    /// @param input The public inputs. These are elements of the scalar field Fr.
    /// @return x The X coordinate of the resulting G1 point.
    /// @return y The Y coordinate of the resulting G1 point.
    function publicInputMSM(
        uint256[8] memory input,
        uint256 publicCommit,
        uint256[2] memory commit
    ) internal view returns (uint256 x, uint256 y) {
        // Note: The ECMUL precompile does not reject unreduced values, so we check this.
        // Note: Unrolling this loop does not cost much extra in code-size, the bulk of the
        //       code-size is in the PUB_ constants.
        // ECMUL has input (x, y, scalar) and output (x', y').
        // ECADD has input (x1, y1, x2, y2) and output (x', y').
        // We call them such that ecmul output is already in the second point
        // argument to ECADD so we can have a tight loop.
        bool success = true;
        assembly ("memory-safe") {
            let f := mload(0x40)
            let g := add(f, 0x40)
            let s
            mstore(f, CONSTANT_X)
            mstore(add(f, 0x20), CONSTANT_Y)
            mstore(g, PUB_0_X)
            mstore(add(g, 0x20), PUB_0_Y)
            s := mload(input)
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_1_X)
            mstore(add(g, 0x20), PUB_1_Y)
            s := mload(add(input, 32))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_2_X)
            mstore(add(g, 0x20), PUB_2_Y)
            s := mload(add(input, 64))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_3_X)
            mstore(add(g, 0x20), PUB_3_Y)
            s := mload(add(input, 96))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_4_X)
            mstore(add(g, 0x20), PUB_4_Y)
            s := mload(add(input, 128))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_5_X)
            mstore(add(g, 0x20), PUB_5_Y)
            s := mload(add(input, 160))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_6_X)
            mstore(add(g, 0x20), PUB_6_Y)
            s := mload(add(input, 192))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_7_X)
            mstore(add(g, 0x20), PUB_7_Y)
            s := mload(add(input, 224))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_8_X)
            mstore(add(g, 0x20), PUB_8_Y)

            s := mload(add(input, 256))
            mstore(add(g, 0x40), publicCommit)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))

            s := mload(commit)
            mstore(g, s) // save commit[0]
            s := mload(add(commit, 32))
            mstore(add(g, 0x20), s) // save commit[1]

            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))

            x := mload(f)
            y := mload(add(f, 0x20))
        }
        if (!success) {
            // Either Public input not in field, or verification key invalid.
            // We assume the contract is correctly generated, so the verification key is valid.
            revert PublicInputNotInField();
        }
    }

    /// Verify an uncompressed Groth16 proof.
    /// @notice Reverts with InvalidProof if the proof is invalid or
    /// with PublicInputNotInField the public input is not reduced.
    /// @notice There is no return value. If the function does not revert, the
    /// proof was successfully verified.
    /// @param proof the points (A, B, C) in EIP-197 format matching the output
    /// of compressProof.
    /// @param input the public input field elements in the scalar field Fr.
    /// Elements must be reduced.
    function verifyProof(
        uint256[8] memory proof,
        uint256[2] memory commitment,
        uint256[2] memory commitmentPOK,
        uint256[8] memory input
    ) public view returns (bool) {
        uint256 inputFr = uint256(keccak256(abi.encodePacked(commitment[0], commitment[1]))) % MOD_R;
        (uint256 x, uint256 y) = publicInputMSM(input, inputFr, commitment);

        // Note: The precompile expects the F2 coefficients in big-endian order.
        // Note: The pairing precompile rejects unreduced values, so we won't check that here.

        bool success;

        uint256 a0 = proof[0];
        uint256 a1 = proof[1];
        uint256 b00 = proof[2];
        uint256 b01 = proof[3];
        uint256 b10 = proof[4];
        uint256 b11 = proof[5];
        uint256 c0 = proof[6];
        uint256 c1 = proof[7];

        assembly ("memory-safe") {
            let f := mload(0x40) // Free memory pointer.

            // Copy points (A, B, C) to memory. They are already in correct encoding.
            // This is pairing e(A, B) and G1 of e(C, -δ).
            mstore(f, a0)
            mstore(add(f, 0x20), a1)
            mstore(add(f, 0x40), b00)
            mstore(add(f, 0x60), b01)
            mstore(add(f, 0x80), b10)
            mstore(add(f, 0xa0), b11)
            mstore(add(f, 0xc0), c0)
            mstore(add(f, 0xe0), c1)

            // Complete e(C, -δ) and write e(α, -β), e(L_pub, -γ) to memory.
            // OPT: This could be better done using a single codecopy, but
            //      Solidity (unlike standalone Yul) doesn't provide a way to
            //      to do this.
            mstore(add(f, 0x100), DELTA_NEG_X_1)
            mstore(add(f, 0x120), DELTA_NEG_X_0)
            mstore(add(f, 0x140), DELTA_NEG_Y_1)
            mstore(add(f, 0x160), DELTA_NEG_Y_0)
            mstore(add(f, 0x180), ALPHA_X)
            mstore(add(f, 0x1a0), ALPHA_Y)
            mstore(add(f, 0x1c0), BETA_NEG_X_1)
            mstore(add(f, 0x1e0), BETA_NEG_X_0)
            mstore(add(f, 0x200), BETA_NEG_Y_1)
            mstore(add(f, 0x220), BETA_NEG_Y_0)
            mstore(add(f, 0x240), x)
            mstore(add(f, 0x260), y)
            mstore(add(f, 0x280), GAMMA_NEG_X_1)
            mstore(add(f, 0x2a0), GAMMA_NEG_X_0)
            mstore(add(f, 0x2c0), GAMMA_NEG_Y_1)
            mstore(add(f, 0x2e0), GAMMA_NEG_Y_0)

            let c
            c := mload(commitment)
            mstore(add(f, 0x300), c) // save commitment[0]
            c := mload(add(commitment, 32))
            mstore(add(f, 0x320), c) // save commitment[1]

            mstore(add(f, 0x340), PEDERSEN_G_X_1)
            mstore(add(f, 0x360), PEDERSEN_G_X_0)
            mstore(add(f, 0x380), PEDERSEN_G_Y_1)
            mstore(add(f, 0x3a0), PEDERSEN_G_Y_0)

            c := mload(commitmentPOK)
            mstore(add(f, 0x3c0), c) // save knowledgeProof[0]
            c := mload(add(commitmentPOK, 32))
            mstore(add(f, 0x3e0), c) // save knowledgeProof[1]

            mstore(add(f, 0x400), PEDERSEN_GROOTSIGMANEG_X_1)
            mstore(add(f, 0x420), PEDERSEN_GROOTSIGMANEG_X_0)
            mstore(add(f, 0x440), PEDERSEN_GROOTSIGMANEG_Y_1)
            mstore(add(f, 0x460), PEDERSEN_GROOTSIGMANEG_Y_0)

            // Check pairing equation.
            success := staticcall(gas(), PRECOMPILE_VERIFY, f, 0x480, f, 0x20)
            // Also check returned value (both are either 1 or 0).
            success := and(success, mload(f))
        }
        if (!success) {
            // Either proof or verification key invalid.
            // We assume the contract is correctly generated, so the verification key is valid.
            revert ProofInvalid();
        }
        return success;
    }

    function verifyRaw(bytes calldata proofData) external view returns (bool) {
        uint256[8] memory proof;
        proof[0] = uint256(bytes32(proofData[:32]));
        proof[1] = uint256(bytes32(proofData[32:64]));
        proof[2] = uint256(bytes32(proofData[64:96]));
        proof[3] = uint256(bytes32(proofData[96:128]));
        proof[4] = uint256(bytes32(proofData[128:160]));
        proof[5] = uint256(bytes32(proofData[160:192]));
        proof[6] = uint256(bytes32(proofData[192:224]));
        proof[7] = uint256(bytes32(proofData[224:256]));

        uint256[2] memory commitment;
        commitment[0] = uint256(bytes32(proofData[256:288]));
        commitment[1] = uint256(bytes32(proofData[288:320]));

        uint256[2] memory commitmentPOK;
        commitmentPOK[0] = uint256(bytes32(proofData[320:352]));
        commitmentPOK[1] = uint256(bytes32(proofData[352:384]));

        uint256[8] memory input;
        input[0] = uint256(bytes32(proofData[384:416])); // commit hash
        input[1] = uint256(uint128(bytes16(proofData[416:432]))); // smt root 0
        input[2] = uint256(uint128(bytes16(proofData[432:448]))); // smt root 1
        input[3] = uint256(uint128(bytes16(proofData[448:464]))); // output commitment 0
        input[4] = uint256(uint128(bytes16(proofData[464:480]))); // output commitment 1
        input[5] = uint256(bytes32(proofData[480:512])); // circuit digest
        input[6] = uint256(bytes32(proofData[512:544])); // dummy commitment
        input[7] = uint256(bytes32(proofData[544:576])); // app vk hash

        return this.verifyProof(proof, commitment, commitmentPOK, input);
    }
}
