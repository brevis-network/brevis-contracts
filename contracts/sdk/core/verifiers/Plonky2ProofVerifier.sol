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
    uint256 constant ALPHA_X = 15662887438510689113245403064437449301918771302994027747332214452842251083773;
    uint256 constant ALPHA_Y = 18810300503474274362307932430574365648616556487532802808782878613157079508257;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 1077993296495892462123890573528146580970055695742816509342512580347427004106;
    uint256 constant BETA_NEG_X_1 = 4344471847245836152409578130598243501451474367983009336320695180707670321623;
    uint256 constant BETA_NEG_Y_0 = 17116308193210263820311934916334629486385792794672366677420917882990527887794;
    uint256 constant BETA_NEG_Y_1 = 10690646525382930297040163652996602033677571308448605096343007845753736817129;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 11525950012120967595324962393357382558808781082618958195163038978519861073815;
    uint256 constant GAMMA_NEG_X_1 = 5415201663983032552924115535229767007087665787910829160640321816509820946238;
    uint256 constant GAMMA_NEG_Y_0 = 8251365719042446374632169442407961719001754484602546891328843770306595746080;
    uint256 constant GAMMA_NEG_Y_1 = 427917952069507608206295397114532500963944794210439183969234376946262008142;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 12098457871850348740765743461253079768086445531438241855356220773595764120360;
    uint256 constant DELTA_NEG_X_1 = 21335598216768396020225752984239559871722776521835944117236449579727409828783;
    uint256 constant DELTA_NEG_Y_0 = 10888755558616063241203233357012428315979327354609754347768626845621276174535;
    uint256 constant DELTA_NEG_Y_1 = 8261893971866071791650192113407752063286189472181116230523168756517396090439;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 9351415229473439228125436398296772821293004972291979786073913940004492958508;
    uint256 constant PEDERSEN_G_X_1 = 14719190946487056136461798877282772469708116895990459717545812373183330375358;
    uint256 constant PEDERSEN_G_Y_0 = 12953845885897627588836045041572617798178175604232216295410751958127023747326;
    uint256 constant PEDERSEN_G_Y_1 = 18666732610208156221473739148692793838408719187618260589200353480581221843461;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 =
        21505379607885122009953351770408508416300000766254241868101458101697608187837;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 =
        2030938172933112431869863941948069252617984473137149834426985863056323485998;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 =
        16421963408192857261465288505251823743005906191763480780698561612688354193564;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 =
        6106969836540272191339684324668517442136584166977361754486796138964061701414;

    // Constant and public input points
    uint256 constant CONSTANT_X = 4066879786622539654360163533447495143816198716738061491754616469283063625572;
    uint256 constant CONSTANT_Y = 16799107499947660767993676898199808608457227948087367139003956661676343664958;
    uint256 constant PUB_0_X = 6826919586776771174046135192389928579546269329760933023971949947409663549985;
    uint256 constant PUB_0_Y = 1813883115212151780253186580736779327099709600622357661834089578305617663483;
    uint256 constant PUB_1_X = 6109956587526078870320882566812893080905733994489570352193766748754037242199;
    uint256 constant PUB_1_Y = 16509152429262282547110357212112996820225121169378745427183438537130840914550;
    uint256 constant PUB_2_X = 5487754820386608172271494347324111595058628346601312539519950441508704201253;
    uint256 constant PUB_2_Y = 8311775301199115149947804710535022566106829036461077401113632907429740869542;
    uint256 constant PUB_3_X = 3530145260290829983784155725719148969916552677941286361838155207491369173809;
    uint256 constant PUB_3_Y = 8172348758451956606475842772858386451775280930758162499758836814184724224403;
    uint256 constant PUB_4_X = 2211761900442233383723332938122942719031302197211731388082615895687902355197;
    uint256 constant PUB_4_Y = 17003602372273517324569255535037298093466829138081794416417844077108946162769;
    uint256 constant PUB_5_X = 21666654307061514712733853315343775746371521290985330557992585818558393200511;
    uint256 constant PUB_5_Y = 8200350480001575514838494296462233486909735824174878340151914113753134023860;
    uint256 constant PUB_6_X = 20107157134272471496844182341815345628794737378089538544290745009847585664605;
    uint256 constant PUB_6_Y = 3762390217765803306161839768826061187577750540546463251022688223693718495355;
    uint256 constant PUB_7_X = 8730441759520402300673040564463242851287640656314889147135106176295744061696;
    uint256 constant PUB_7_Y = 6938879651126100567255381924403425815420512898847539067063784929495235791095;

    /// Compute the public input linear combination.
    /// @notice Reverts with PublicInputNotInField if the input is not in the field.
    /// @notice Computes the multi-scalar-multiplication of the public input
    /// elements and the verification key including the constant term.
    /// @param input The public inputs. These are elements of the scalar field Fr.
    /// @return x The X coordinate of the resulting G1 point.
    /// @return y The Y coordinate of the resulting G1 point.
    function publicInputMSM(
        uint256[7] memory input,
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
        uint256[7] memory input
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

        uint256[7] memory input;
        input[0] = uint256(bytes32(proofData[384:416])); // commit hash
        input[1] = uint256(uint128(bytes16(proofData[416:432]))); // smt root 0
        input[2] = uint256(uint128(bytes16(proofData[432:448]))); // smt root 1
        input[3] = uint256(uint128(bytes16(proofData[448:464]))); // output commitment 0
        input[4] = uint256(uint128(bytes16(proofData[464:480]))); // output commitment 1
        input[5] = uint256(bytes32(proofData[480:512])); // app vk hash
        input[6] = uint256(bytes32(proofData[512:544])); // dummy commitment

        return this.verifyProof(proof, commitment, commitmentPOK, input);
    }
}
