
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @title Groth16 verifier template.
/// @author Remco Bloemen
/// @notice Supports verifying Groth16 proofs. Proofs can be in uncompressed
/// (256 bytes) and compressed (128 bytes) format. A view function is provided
/// to compress proofs.
/// @notice See <https://2π.com/23/bn254-compression> for further explanation.
contract Plonky2ProofVerifierForxLayer {
    
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
    uint256 constant ALPHA_X = 16941860205729802054752993470384300197293358649359706940281569884876143849840;
    uint256 constant ALPHA_Y = 13394057771814856230214750307737187189899932134484893800747581458086081533658;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 19376995576545274164422704811955514682839953295923042914314742995899959615660;
    uint256 constant BETA_NEG_X_1 = 21868334384024131284204756717012453304696569873580356066965011731235952676966;
    uint256 constant BETA_NEG_Y_0 = 12278658128973420787340717003814710928948551775583460093253659416356567177704;
    uint256 constant BETA_NEG_Y_1 = 150209684864089997978366835595789142737663547234126199731939172890473484960;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 882478628596069499100798018711415565283133540502433408649537021990639793304;
    uint256 constant GAMMA_NEG_X_1 = 18030021345095537783305818872351071945058996265297007503943091336207688588416;
    uint256 constant GAMMA_NEG_Y_0 = 888999880467229598325292163972511327508583842485402544002372464663257907423;
    uint256 constant GAMMA_NEG_Y_1 = 12503987446671243481884282202916062362672859286321391905449565490380552988497;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 16401501412610436179791024763541132569094930990627606654548550911991228069918;
    uint256 constant DELTA_NEG_X_1 = 8003820435230069581871795633564894185525694133552045759036075857698353844767;
    uint256 constant DELTA_NEG_Y_0 = 8660519134829653668367971591385318994212415066613978836646859848940767271643;
    uint256 constant DELTA_NEG_Y_1 = 17249935355182213687303842453353875808023380144831690770688061811080433948731;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 12405836372395435985425235576174826650648275285430633133802852707908493414903;
    uint256 constant PEDERSEN_G_X_1 = 13658066220314453704088643681180975609003098260827998189786184791889900039169;
    uint256 constant PEDERSEN_G_Y_0 = 16661322741961211487976464338642933066238174688140160452752666486585397216138;
    uint256 constant PEDERSEN_G_Y_1 = 2471556477446606044567382216323556520646965011532911625585891661254114202820;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 = 13461884536574381732989038503740343510680190513459109204811206277714232110762;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 = 15041495481447641666818224585557646886172049299972296944601850456002793424276;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 = 14202656463787153687167922146708345991735495796836944024972696339384200458152;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 = 13848576030828357870643246877278888546442101521708840488194133893740815607300;

    // Constant and public input points
    uint256 constant CONSTANT_X = 3175410241884927963331166726277081995922583857612514257948619677109974892991;
    uint256 constant CONSTANT_Y = 15416712579334141859182728839546756150888155712541552439056649183356970958085;
    uint256 constant PUB_0_X = 161719209698983373501751073216935812521546179758422097386974487847826391550;
    uint256 constant PUB_0_Y = 20136500238789400434749681585584717277672278312228007708365391346350208233007;
    uint256 constant PUB_1_X = 18446040354197002276092439222739821176489228084520157944201624012682477282803;
    uint256 constant PUB_1_Y = 2926777123788159659621796261339158339661038926974214870290963794751983236350;
    uint256 constant PUB_2_X = 3305231687971983810581305335629621944798342718031239707802060096532300715270;
    uint256 constant PUB_2_Y = 20515687856100209724658310700094199946435391297088473114435977373856477840176;
    uint256 constant PUB_3_X = 17417585999763368155091029807614685026885404167728904930345820375804414487500;
    uint256 constant PUB_3_Y = 9122968218472234942697671053138403404531966538878340626323817418148630293239;
    uint256 constant PUB_4_X = 5593309657274025082011402991226807190637196629910659461208288875873949036606;
    uint256 constant PUB_4_Y = 20867664909170623026518437561955941474236087867326313875133485396110390093972;
    uint256 constant PUB_5_X = 8596593948962788370427025607943291710277857580264156739544065875068607952593;
    uint256 constant PUB_5_Y = 684964755690599682850861200537642840498241249341870256085917017144178736992;
    uint256 constant PUB_6_X = 14196555572321690709849080385423978719680877670086588461205591958321521201164;
    uint256 constant PUB_6_Y = 7191782616084889130633211579428022211680932196705632512370251019929091596906;
    uint256 constant PUB_7_X = 9241340272403770484779177914637955704724749136635592544371289160335576064213;
    uint256 constant PUB_7_Y = 3572773935089660196407904573250146473163590220054285343100550120064527652226;

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
