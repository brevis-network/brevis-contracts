// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @title Groth16 verifier template.
/// @author Remco Bloemen
/// @notice Supports verifying Groth16 proofs. Proofs can be in uncompressed
/// (256 bytes) and compressed (128 bytes) format. A view function is provided
/// to compress proofs.
/// @notice See <https://2π.com/23/bn254-compression> for further explanation.
contract SMTUpdateCircuitProofOnOpVerifier {
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
    uint256 constant ALPHA_X = 15419906342546775518056107071206082363336491762605521910208413869208927134694;
    uint256 constant ALPHA_Y = 8925206659625855244561366728843285323744854262045895593028094624373620731922;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 11215850788231168341298441616563847207476572509165490513874803298212456302963;
    uint256 constant BETA_NEG_X_1 = 8238279398116533936998213301569180178754815422316673139943232388076301667239;
    uint256 constant BETA_NEG_Y_0 = 21832070218819376232664150960538222039779319910810640357249292304324844753485;
    uint256 constant BETA_NEG_Y_1 = 9645998624345071248952629980493114216707576571248096877536647337508115727673;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 10601238188791374550285170434453601938335804995266934091716139317125293638278;
    uint256 constant GAMMA_NEG_X_1 = 13677398231709499629906376603684156136667816516228700490473454800982690413108;
    uint256 constant GAMMA_NEG_Y_0 = 17041542147468964349964039722920041227280608734181340612910512772859614705745;
    uint256 constant GAMMA_NEG_Y_1 = 12907952842297082648757487550998246103074200984829304277298928430029071100637;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 8232190889596979213738055480496452392103977043528541478688126078314412227487;
    uint256 constant DELTA_NEG_X_1 = 1065931679919289878696551123315192093061984499916404628172416994815927094095;
    uint256 constant DELTA_NEG_Y_0 = 12552350698181263435179086441908958505534872197386922161059078318191731336963;
    uint256 constant DELTA_NEG_Y_1 = 6511382691711036021804729349675340256940372423678254876986120440897243279309;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 18805755783857632130494169760142971459211085766908956007665772744607805349331;
    uint256 constant PEDERSEN_G_X_1 = 6577130346294027956106129016242575755376501376806793281099727040898065194630;
    uint256 constant PEDERSEN_G_Y_0 = 3063828328789064990096030216023403162709221064383189238160145907411059445654;
    uint256 constant PEDERSEN_G_Y_1 = 3518510766686242646084167009169793143410416394108769285506541954245359469430;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 =
        20573844106836851216683495680715692316448816169379645474552662784955401807782;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 =
        10173198368954426359560222122775027390195584468790431924474000180170482567391;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 =
        8933462681644689232793574626123804034150276190704931603444382781462281429595;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 =
        17690363985422853860548334460761950515160248144161805124506500282167531442429;

    // Constant and public input points
    uint256 constant CONSTANT_X = 17781150640532246683478288430048881803672435485525463173154420895146090513125;
    uint256 constant CONSTANT_Y = 16641756872356175845605867643980429319093093385697558171090848764055531414861;
    uint256 constant PUB_0_X = 21762954875859522910283083606892400572088600507337073996229569604242842249270;
    uint256 constant PUB_0_Y = 19663976138561437876960571951888052852491438808124896768715500012952888092467;
    uint256 constant PUB_1_X = 1282725618717198352780293692830891443025482340068889157440025831574202963944;
    uint256 constant PUB_1_Y = 5141998626009341098281690417433558725725706726366284990672385112933672017322;
    uint256 constant PUB_2_X = 1240446175031511932337110215869561493021210843231095144396369783697007421678;
    uint256 constant PUB_2_Y = 6097377749748449085475680060974341379881798127128503720246824284984211576709;
    uint256 constant PUB_3_X = 7697944521653177457169858557057282285863319527054233090334250110271583519921;
    uint256 constant PUB_3_Y = 7468913458484340056196103603113691168882715555713943918188547289763943605351;
    uint256 constant PUB_4_X = 14709897161989503194601316482911961165234009336256324577023792817414900623418;
    uint256 constant PUB_4_Y = 4749531004971325732649341235206855488889769848132077726737632639652670873328;
    uint256 constant PUB_5_X = 4382497355660004011670078123988761997560027829055718786348255676075333012297;
    uint256 constant PUB_5_Y = 12909374075361927801637650590989430487973660002571842025096047421198361103419;
    uint256 constant PUB_6_X = 3045310679983037167284855434152608108004309361015643912970003163614423182850;
    uint256 constant PUB_6_Y = 3855115339596168614477475672569798808024932255699838809951164243482027405378;
    uint256 constant PUB_7_X = 9759470637889957138632189878058376019817094047617385414100630665386579435922;
    uint256 constant PUB_7_Y = 18723926700666823905805244351111906954741145157310895366231249412894474043073;
    uint256 constant PUB_8_X = 7834429028994559907616029020492869687127296893521648940777211614458493363574;
    uint256 constant PUB_8_Y = 8246737296674016630208054535734405751622077371125335638920262402621189665145;
    uint256 constant PUB_9_X = 16292799440761393277617698507073614168294446014741578359675905806367574298656;
    uint256 constant PUB_9_Y = 11592802447601385600742313198658048780149326795010307963294046197593644867387;

    /// Compute the public input linear combination.
    /// @notice Reverts with PublicInputNotInField if the input is not in the field.
    /// @notice Computes the multi-scalar-multiplication of the public input
    /// elements and the verification key including the constant term.
    /// @param input The public inputs. These are elements of the scalar field Fr.
    /// @return x The X coordinate of the resulting G1 point.
    /// @return y The Y coordinate of the resulting G1 point.
    function publicInputMSM(
        uint256[9] calldata input,
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
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_6_X)
            mstore(add(g, 0x20), PUB_6_Y)
            s := calldataload(add(input, 192))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_7_X)
            mstore(add(g, 0x20), PUB_7_Y)
            s := calldataload(add(input, 224))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_8_X)
            mstore(add(g, 0x20), PUB_8_Y)
            s := calldataload(add(input, 256))
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_9_X)
            mstore(add(g, 0x20), PUB_9_Y)

            s := calldataload(add(input, 288))
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
        uint256[9] calldata input
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
}
