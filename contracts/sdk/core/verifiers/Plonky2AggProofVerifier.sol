// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @title Groth16 verifier template.
/// @author Remco Bloemen
/// @notice Supports verifying Groth16 proofs. Proofs can be in uncompressed
/// (256 bytes) and compressed (128 bytes) format. A view function is provided
/// to compress proofs.
/// @notice See <https://2π.com/23/bn254-compression> for further explanation.
contract Plonky2AggProofVerifier {

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
    uint256 constant ALPHA_X = 18528005100652142185651448115945208981606486485125287834673377562318730932645;
    uint256 constant ALPHA_Y = 6934256914272587171417385458321750823327038732992209218016284866907350901631;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 17398279421220782971683362974499061297663399136559904597515951854670640930533;
    uint256 constant BETA_NEG_X_1 = 9597415710432537555295181310584051494349852196208956942308755798030456348445;
    uint256 constant BETA_NEG_Y_0 = 12552087533025170243003254464256647117426884615888424718421204540429912422226;
    uint256 constant BETA_NEG_Y_1 = 12774856770953585283643391443089399529151236951968415802063106324832738028778;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 4232509125452353759379442192265774784844292158972696372979377377877606203349;
    uint256 constant GAMMA_NEG_X_1 = 2037251090331568979206382698886837545525022530674062358847267977540304162956;
    uint256 constant GAMMA_NEG_Y_0 = 9332066037278655901898321262897557396686529893482538030585465938123378655515;
    uint256 constant GAMMA_NEG_Y_1 = 1954581634125944447937474766089334458464007610893048392418663692146522540180;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 5743366975238211595434447868739957537755833151337869116146911624544932342346;
    uint256 constant DELTA_NEG_X_1 = 15564046426623817987465990162984286364709419878129905556385063694195706402982;
    uint256 constant DELTA_NEG_Y_0 = 4821448421572386419305425143664207506856750281740418661966099453023874371243;
    uint256 constant DELTA_NEG_Y_1 = 3762685771259665578269773653491283128074805605662357060386388798905911322270;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 8200151898097693893924693494218112323111137741943661618866020258226737484690;
    uint256 constant PEDERSEN_G_X_1 = 339684239955717580477005784623686646758996833667904987940912658755622992119;
    uint256 constant PEDERSEN_G_Y_0 = 5429001831343467033543737962261080360341837417649349553857113244564755897137;
    uint256 constant PEDERSEN_G_Y_1 = 2963329897730790074584294797238813837913764062653341142265475407074977857789;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 = 17124541897286272868249120319356675068156673196281740074444019772026212496790;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 = 15430431486819796964627011574699634520623612927515510006746128400367697027029;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 = 5314565612731024468674617684006532513271627793070819078226219999220871877710;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 = 12981393382811674498426158985106834223701443417853821268580626933313409297616;

    // Constant and public input points
    uint256 constant CONSTANT_X = 11632587137175963500759160382363676059042501910716224795476401182750971570050;
    uint256 constant CONSTANT_Y = 21297278341355033768361254531373605759527941692703852884659193259868597257772;
    uint256 constant PUB_0_X = 7231784866304085789215947880056913368681300388350025903518750579786766765607;
    uint256 constant PUB_0_Y = 19161947333775088067437853288800226218251231802857126851866282697883127560889;
    uint256 constant PUB_1_X = 15871706201371427865014450360246144987973204270285304395354452052394893365085;
    uint256 constant PUB_1_Y = 21318153598021825164572065302700674623728812859698920164384998475991659891039;
    uint256 constant PUB_2_X = 10646523027605994273940192129173578667711108069086154859127277288296427019024;
    uint256 constant PUB_2_Y = 17735601254226496133408474789690691727351685564232292195453104432422603581886;
    uint256 constant PUB_3_X = 15641174821056639396268207222814733182462930401937724785057952326396461616935;
    uint256 constant PUB_3_Y = 6436573545471145945105511070708323505989949995762215349761794131457669128349;
    uint256 constant PUB_4_X = 10324046772086138590249367031696010510557359734540490651965428638506783126895;
    uint256 constant PUB_4_Y = 687519780340405754354983526872303225946761493446264138754544934069054061225;

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
        input[1] = uint256(uint128(bytes16(proofData[384:400]))); // merkleRoot 0
        input[0] = uint256(uint128(bytes16(proofData[400:416]))); // merkleRoot 1
        input[3] = uint256(uint128(bytes16(proofData[416:432]))); // proofIdsCommit 0
        input[2] = uint256(uint128(bytes16(proofData[432:448]))); // proofIdsCommit 1

        return this.verifyProof(proof, commitment, commitmentPOK, input);
    }
}
