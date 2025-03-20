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
    uint256 constant ALPHA_X = 20130115150190108019054123324150266355963634520665803044119931121920433877502;
    uint256 constant ALPHA_Y = 14285561117546692228188264117658849531292193248102265302826521527658774130730;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 12173582291697991928038706774854257923669411098639177950151278294306876399185;
    uint256 constant BETA_NEG_X_1 = 15132618022772140126948816923300076218627540588695058850129828561330319272394;
    uint256 constant BETA_NEG_Y_0 = 21443297657689754130694783041616163165528212213189433353793846678725256110968;
    uint256 constant BETA_NEG_Y_1 = 10501575552957665887945655290567158183091839264534194878186069451520528855363;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 21190170695766197642735844330659825381395925533125398227031524291838460125060;
    uint256 constant GAMMA_NEG_X_1 = 19203026818133896529218674625910724378196155667663875895123556718339151509868;
    uint256 constant GAMMA_NEG_Y_0 = 11205639707683230252739576378165133363686151523964908586275163372934586463856;
    uint256 constant GAMMA_NEG_Y_1 = 10214449308851856821972963282423085145860104790900192284349190463458888620188;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 10399319314398376417820802724200232724959666061017285602953069221524861793753;
    uint256 constant DELTA_NEG_X_1 = 452886078276742855748568059052645802297245110497791390730521524044942878705;
    uint256 constant DELTA_NEG_Y_0 = 18843179484068923140237924738630828752463118041111414293311624672247179311708;
    uint256 constant DELTA_NEG_Y_1 = 6578419443316232797720362393723449118629184777106657540493243117798965076474;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 18840956241968029731100325856496885824192733037252226446060524528156530111308;
    uint256 constant PEDERSEN_G_X_1 = 21274503460330121600530385632507567737303199317766957577320329332713158025884;
    uint256 constant PEDERSEN_G_Y_0 = 21703706571156547469091964352250046820952425344667079956389688490794984176885;
    uint256 constant PEDERSEN_G_Y_1 = 17581065412830892652496072086768587586258265586086612253128718326201955073618;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 =
        9477073819196801833214615315952587768947385509800268046198048251598908341847;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 =
        8733489005457477856766509430546022763271987272883810365436552332880971951842;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 =
        855147140980254462115987301424675648082659807931923152867586423105627459442;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 =
        21617656332261003847761610550493473011912376763721964847518217346669211124634;

    // Constant and public input points
    uint256 constant CONSTANT_X = 12096385008008835625928732081395595134182198718451766067043585113744474120373;
    uint256 constant CONSTANT_Y = 7743770677267848242696443432401875246365945607818899290902759876388111845102;
    uint256 constant PUB_0_X = 11517708531354116769553564301121006653788372143772228218883516720808567029236;
    uint256 constant PUB_0_Y = 2152127335932304681966390681193568647248500526879538976519507779251881398051;
    uint256 constant PUB_1_X = 2411623673232227440182427003605066260447834996624159200713714450678525673813;
    uint256 constant PUB_1_Y = 20302495699790144734302725144246911109225787328411619703039824705266308249975;
    uint256 constant PUB_2_X = 18806945557217583070771738455262435430543222198660506094198443254708010681201;
    uint256 constant PUB_2_Y = 488608118252318333241074170519058509672247917809847727458802505257044789554;
    uint256 constant PUB_3_X = 3271889172046179794175799562498488678635115520079974329494182174355427750959;
    uint256 constant PUB_3_Y = 13888544571628847530483236925932420001521195941868097389115207259520893393129;
    uint256 constant PUB_4_X = 15735991015174072795703790087522471561699190385808970605894593460553789112495;
    uint256 constant PUB_4_Y = 12221542038876261891140618611503962465999310138418282525833171556590552238586;
    uint256 constant PUB_5_X = 8008225254116302360001416404876069698232064530453411337872616393927786537230;
    uint256 constant PUB_5_Y = 10268217310132774904598799980256756069898428929972715439619293670935450693965;

    /// Compute the public input linear combination.
    /// @notice Reverts with PublicInputNotInField if the input is not in the field.
    /// @notice Computes the multi-scalar-multiplication of the public input
    /// elements and the verification key including the constant term.
    /// @param input The public inputs. These are elements of the scalar field Fr.
    /// @return x The X coordinate of the resulting G1 point.
    /// @return y The Y coordinate of the resulting G1 point.
    function publicInputMSM(
        uint256[5] calldata input,
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
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_5_X)
            mstore(add(g, 0x20), PUB_5_Y)

            s := calldataload(add(input, 160))
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
        uint256[5] calldata input
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

        uint256[5] memory input;
        input[1] = uint256(uint128(bytes16(proofData[384:400]))); // merkleRoot 0
        input[0] = uint256(uint128(bytes16(proofData[400:416]))); // merkleRoot 1
        input[3] = uint256(uint128(bytes16(proofData[416:432]))); // proofIdsCommit 0
        input[2] = uint256(uint128(bytes16(proofData[432:448]))); // proofIdsCommit 1
        input[4] = uint256(bytes32(proofData[448:480])); // sub proof vk hash, mimc bn254

        return this.verifyProof(proof, commitment, commitmentPOK, input);
    }
}
