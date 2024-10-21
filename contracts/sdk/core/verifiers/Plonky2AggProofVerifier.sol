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
    uint256 constant ALPHA_X = 12454639809842667797394630752713125293942486255440070292261011069638209746471;
    uint256 constant ALPHA_Y = 15070132500661191985113312776840364149227735719016533076417811736233853166815;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 14058705832211816682066839172722736629452548018725194211391425038677442608388;
    uint256 constant BETA_NEG_X_1 = 17804226448727052455902945802974516765081608503149671319997178058597811465108;
    uint256 constant BETA_NEG_Y_0 = 16968671663864704468603744894609679097412821293448412652759893538983278093511;
    uint256 constant BETA_NEG_Y_1 = 9770908036313116798967463002067633841001727429059466786511012992267890254408;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 172798160947796533735987019290706361968120975932024050566772674902613552896;
    uint256 constant GAMMA_NEG_X_1 = 16267280011774932988250022148628777512720137986207287816170425666902611826220;
    uint256 constant GAMMA_NEG_Y_0 = 16782902920448377215582660374577945264584434951738912973034376157859442903241;
    uint256 constant GAMMA_NEG_Y_1 = 1200352801813531414623118777798778780918574624986484083254433270177776695085;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 16715091264394372825910973277706090845285413839346696909850920441501713575861;
    uint256 constant DELTA_NEG_X_1 = 628999461778456374476051498969811336278955444386627300191498531330744870721;
    uint256 constant DELTA_NEG_Y_0 = 11814424011112156410183227391276866194862181757089289421210876934526850739408;
    uint256 constant DELTA_NEG_Y_1 = 21421525857278434164019175757933421249465148957266128717700205937703158185884;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 19764112198930581842141071815733380966672010013541591689576199773814835453336;
    uint256 constant PEDERSEN_G_X_1 = 4410785297896959438995012775397418225794446167574159270543377733417857884647;
    uint256 constant PEDERSEN_G_Y_0 = 11884569594740335759883765413676209330794564713158149929681152709699082972916;
    uint256 constant PEDERSEN_G_Y_1 = 4524915528697232139860100934965979558142590460356430894197559431787418227478;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 =
        12607697000684170912122929684112745292284916963921872699462113197565979831494;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 =
        12414118764379179668081870000439014135916616054051755764925269833325737276078;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 =
        19607314541962267319414455640857464632527469506859149332444863330202799597995;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 =
        14827596690164094730431743569600948747230464750185328123920781268041329660488;

    // Constant and public input points
    uint256 constant CONSTANT_X = 17248754784954784304119625497289746762186351665638754436839291470992613401487;
    uint256 constant CONSTANT_Y = 5695082879045810079875299544755485340692823731671202523371257132854657014133;
    uint256 constant PUB_0_X = 5301254920570855357657704570207461271207809093897436827308701206870458083407;
    uint256 constant PUB_0_Y = 8585028929703964460928729350114800282113835330818362400015805091773135548383;
    uint256 constant PUB_1_X = 10513977477828921391501624276550584594137184964382482461882297623074384679991;
    uint256 constant PUB_1_Y = 16781822363202394627973185866884529697032141346633640915706381972247697459717;
    uint256 constant PUB_2_X = 1794654636925950890529949302831578808051728380961238285897279587755122082421;
    uint256 constant PUB_2_Y = 8021717449672890912066837954473856071733800277886171835531545342414240033343;
    uint256 constant PUB_3_X = 17860765407182479893598378055099467270344305495284012251939801199029522192562;
    uint256 constant PUB_3_Y = 21774970136118375471568176971816900819924288719868932079216836033111544334201;
    uint256 constant PUB_4_X = 9137979302822181778190865131728770525013979066904037930253829001928561220402;
    uint256 constant PUB_4_Y = 8272658343841756424872816522167139717295933050368542222566897157025870276134;

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
