// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @title Groth16 verifier template.
/// @author Remco Bloemen
/// @notice Supports verifying Groth16 proofs. Proofs can be in uncompressed
/// (256 bytes) and compressed (128 bytes) format. A view function is provided
/// to compress proofs.
/// @notice See <https://2π.com/23/bn254-compression> for further explanation.
contract Bn254Agg16Bn254Verifier {
    /// Some of the provided public input values are larger than the field modulus.
    /// @dev Public input elements are not automatically reduced, as this is can be
    /// a dangerous source of bugs.
    error PublicInputNotInField();

    /// The proof is invalid.
    /// @dev This can mean that provided Groth16 proof points are not on their
    /// curves, that pairing equation fails, or that the proof is not for the
    /// provided public input.
    error ProofInvalid();
    /// The commitment is invalid
    /// @dev This can mean that provided commitment points and/or proof of knowledge are not on their
    /// curves, that pairing equation fails, or that the commitment and/or proof of knowledge is not for the
    /// commitment key.
    error CommitmentInvalid();

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
    uint256 constant ALPHA_X = 1702561344448085281094789774600128263574687193187158205383665062185976344077;
    uint256 constant ALPHA_Y = 7301803822542645859693868319673688816340240735245295748912341029232073884893;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 18447598694048246903669311741796039889581197963734674587204001415120693662269;
    uint256 constant BETA_NEG_X_1 = 9906792182456187874093099216765147414966609121010960408702046751943428047884;
    uint256 constant BETA_NEG_Y_0 = 8212080886742499081890418886403412084273567012234973032882039331422457759102;
    uint256 constant BETA_NEG_Y_1 = 11498531656360166433778404625960782488873912593839517306158651415443690583476;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 2491710546609360521903143794380373439301931295077953562212882998240993333168;
    uint256 constant GAMMA_NEG_X_1 = 7747370486270831540667976309020769645992343189153115852884869854194470189192;
    uint256 constant GAMMA_NEG_Y_0 = 3223246835726549345027047391658349688925290193702683330987447146534706846416;
    uint256 constant GAMMA_NEG_Y_1 = 13767384139254646155439085491171425872623257436338424656611276147733312508813;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 21780434321833569700455680174765178407605639447990776771414320305650209407328;
    uint256 constant DELTA_NEG_X_1 = 16666987695231315272653987812257126345790246676011502930648946730270386439075;
    uint256 constant DELTA_NEG_Y_0 = 5057694115975673707013606317563384282790453820773866054449478572268869819639;
    uint256 constant DELTA_NEG_Y_1 = 11811922725085387303004152384361723091514764781991635525808959876317918058910;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 4138036332849951528840855374237403236724199383419025770430014562403649391142;
    uint256 constant PEDERSEN_G_X_1 = 2293045096486387846661300474038431228563424529677972807151991868057625612775;
    uint256 constant PEDERSEN_G_Y_0 = 1087705685416253402974179660337882124517465852693998675374591321185478844126;
    uint256 constant PEDERSEN_G_Y_1 = 19181010027527305347086822924084846014678367928458620110851926246811149924856;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 =
        1364659900616619261366837825352415018281838459958758994298972335695928663992;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 =
        8539606595213343683994667486011158634175978350433009945885353496301776477429;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 =
        4189300617291115316791544645370959035880022218495965504661993378452301869987;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 =
        11401306171979238369177558573213960126147089129804204285535711539750966120541;

    // Constant and public input points
    uint256 constant CONSTANT_X = 14218196459025072344164335469527566142437235019634603139488016000363196487251;
    uint256 constant CONSTANT_Y = 10862497750089339891552565252625097482282524312765781030524460466023606654217;
    uint256 constant PUB_0_X = 18467530036186918258053217278651749755767463729818206783069178679765552786737;
    uint256 constant PUB_0_Y = 3566954893092823524798162335477480636113872220392730697109719553185661702975;
    uint256 constant PUB_1_X = 9328996871313912269678054381737497891833697057833703195939667911089258579558;
    uint256 constant PUB_1_Y = 7089244843500989485069059571976392390446814363347059437315621079814604054132;
    uint256 constant PUB_2_X = 20066608599991512310741760098179945900067357185262509493484535865887356750249;
    uint256 constant PUB_2_Y = 19229941365329726717492564627814137547281282810440695031053635679829463739330;
    uint256 constant PUB_3_X = 10659026592484994710191837691080173295581106172882564842322840613513395930103;
    uint256 constant PUB_3_Y = 3233351632225662099164535535249156731403323281035211072218995698649726250163;
    uint256 constant PUB_4_X = 18062474645004074368737176496637516696228042501784432534460721095322061661000;
    uint256 constant PUB_4_Y = 16302132265137321535158122968164860777126480886662859913133410317899246408539;

    uint256 constant MOD_R = 21888242871839275222246405745257275088548364400416034343698204186575808495617;

    /// Compute the public input linear combination.
    /// @notice Reverts with PublicInputNotInField if the input is not in the field.
    /// @notice Computes the multi-scalar-multiplication of the public input
    /// elements and the verification key including the constant term.
    /// @param input The public inputs. These are elements of the scalar field Fr.
    /// @return x The X coordinate of the resulting G1 point.
    /// @return y The Y coordinate of the resulting G1 point.
    function publicInputMSM(
        uint256[4] calldata input,
        uint256 publicCommit,
        uint256[2] calldata commit
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
            s := calldataload(input)
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_1_X)
            mstore(add(g, 0x20), PUB_1_Y)
            s := calldataload(add(input, 32))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_2_X)
            mstore(add(g, 0x20), PUB_2_Y)
            s := calldataload(add(input, 64))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_3_X)
            mstore(add(g, 0x20), PUB_3_Y)
            s := calldataload(add(input, 96))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_4_X)
            mstore(add(g, 0x20), PUB_4_Y)

            s := calldataload(add(input, 128))
            mstore(add(g, 0x40), publicCommit)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))

            s := calldataload(commit)
            mstore(g, s) // save commit[0]
            s := calldataload(add(commit, 32))
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
        uint256[8] calldata proof,
        uint256[2] calldata commit,
        uint256[2] calldata knowledgeProof,
        uint256[4] calldata input
    ) public view returns (bool) {
        uint256 inputFr = uint256(keccak256(abi.encodePacked(commit[0], commit[1]))) % MOD_R;
        (uint256 x, uint256 y) = publicInputMSM(input, inputFr, commit);

        // Note: The precompile expects the F2 coefficients in big-endian order.
        // Note: The pairing precompile rejects unreduced values, so we won't check that here.

        bool success;
        assembly ("memory-safe") {
            let f := mload(0x40) // Free memory pointer.

            // Copy points (A, B, C) to memory. They are already in correct encoding.
            // This is pairing e(A, B) and G1 of e(C, -δ).
            calldatacopy(f, proof, 0x100)

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
            c := calldataload(commit)
            mstore(add(f, 0x300), c) // save commitment[0]
            c := calldataload(add(commit, 32))
            mstore(add(f, 0x320), c) // save commitment[1]

            mstore(add(f, 0x340), PEDERSEN_G_X_1)
            mstore(add(f, 0x360), PEDERSEN_G_X_0)
            mstore(add(f, 0x380), PEDERSEN_G_Y_1)
            mstore(add(f, 0x3a0), PEDERSEN_G_Y_0)

            c := calldataload(knowledgeProof)
            mstore(add(f, 0x3c0), c) // save knowledgeProof[0]
            c := calldataload(add(knowledgeProof, 32))
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

        uint256[4] memory input;
        input[0] = uint256(uint128(bytes16(proofData[384:400])));
        input[1] = uint256(uint128(bytes16(proofData[400:416])));
        input[2] = uint256(uint128(bytes16(proofData[416:432])));
        input[3] = uint256(uint128(bytes16(proofData[432:448])));

        return this.verifyProof(proof, commitment, commitmentPOK, input);
    }
}
