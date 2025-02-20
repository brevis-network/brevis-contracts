// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @title Groth16 verifier template.
/// @author Remco Bloemen
/// @notice Supports verifying Groth16 proofs. Proofs can be in uncompressed
/// (256 bytes) and compressed (128 bytes) format. A view function is provided
/// to compress proofs.
/// @notice See <https://2π.com/23/bn254-compression> for further explanation.
contract BrevisPlonky2SmtVerifier {
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
    uint256 constant ALPHA_X = 9711627955724311834784488649640453336865142251193944219245711722439808200371;
    uint256 constant ALPHA_Y = 2505492180822423263041788355812321952759547577899451782511577455005757262923;

    // Groth16 beta point in G2 in powers of i
    uint256 constant BETA_NEG_X_0 = 13499871872605582160891168838829521262451463065113808703943528539835480766637;
    uint256 constant BETA_NEG_X_1 = 15820632028657616015057294708134858252841059230313550478583050800247359581668;
    uint256 constant BETA_NEG_Y_0 = 12573886416676420363793835962707721718375062605004734481621777844586734840058;
    uint256 constant BETA_NEG_Y_1 = 17241979172256983835296876038192904084559798366531371721860514963600379797728;

    // Groth16 gamma point in G2 in powers of i
    uint256 constant GAMMA_NEG_X_0 = 533823341112532332181875592415981129573597537830924362231409714066441860387;
    uint256 constant GAMMA_NEG_X_1 = 12348343873624267645852071153801626064988242899882169083465143781800401707028;
    uint256 constant GAMMA_NEG_Y_0 = 17632037757301536792061392519062620953482381884525709241195276603115804280991;
    uint256 constant GAMMA_NEG_Y_1 = 8847137175606651042410049397316706456512545587952246616766240538562941278989;

    // Groth16 delta point in G2 in powers of i
    uint256 constant DELTA_NEG_X_0 = 20781536877202454612313904077056964109565138894038139857017864417168276560552;
    uint256 constant DELTA_NEG_X_1 = 1855809939155669843883400186204044222539672494056438732094818963231743115356;
    uint256 constant DELTA_NEG_Y_0 = 18893290922867552704630406592968945566656800474066642978211441389325229765791;
    uint256 constant DELTA_NEG_Y_1 = 3156392774641462836878961591011179191215975872658186634315222489938736390860;
    // Pedersen G point in G2 in powers of i
    uint256 constant PEDERSEN_G_X_0 = 7590769814777716702368274323822866203310371035775963421714290244640267968810;
    uint256 constant PEDERSEN_G_X_1 = 13938521311676760312862721950667387012021892932709293153309340203335041800131;
    uint256 constant PEDERSEN_G_Y_0 = 16526346463830461447055374657535131780765081106386257238916332479510606353171;
    uint256 constant PEDERSEN_G_Y_1 = 17209259504677688674058987403089010076299754933268187705508266825052530638006;

    // Pedersen GRootSigmaNeg point in G2 in powers of i
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_0 =
        9414394291441915470653952017248010394073630170817386807456454754534113486211;
    uint256 constant PEDERSEN_GROOTSIGMANEG_X_1 =
        993594027156440135435878430487261545559346686988903913189119717580184039914;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_0 =
        6753985584860452045608779244881915107652171843322826399403112115513279076356;
    uint256 constant PEDERSEN_GROOTSIGMANEG_Y_1 =
        20326437300128747457179208676007395546113513461444267108581568995017323902270;

    // Constant and public input points
    uint256 constant CONSTANT_X = 16394830929394585354971696367942597959752012849364332032406395500147703696364;
    uint256 constant CONSTANT_Y = 18990408379527419892373164294228897292913339848455176086560734160755260265556;
    uint256 constant PUB_0_X = 1541846356150730107649357576989558778593922841044037975616107812810761390288;
    uint256 constant PUB_0_Y = 12165766330834405566640925553296798813533309205026909616247549209542384721769;
    uint256 constant PUB_1_X = 21044130314309058741590430830727323323167235138300104168139318366407917460634;
    uint256 constant PUB_1_Y = 6076021234838724413024445698409144369514907023593371385124049879264360116914;
    uint256 constant PUB_2_X = 16624608716437642121497813738496178101080484698345802682659089838844161464017;
    uint256 constant PUB_2_Y = 10079905137910246263777471805555085968086808156311073864304068017881822427903;
    uint256 constant PUB_3_X = 12195366618644555674077950137013938458590421350318166883167722941029843893946;
    uint256 constant PUB_3_Y = 4305785304004972390928951397679187128288606373956332808827657642751230470232;
    uint256 constant PUB_4_X = 18281912648112712652823898977108230597943776269455822175403060820026923549081;
    uint256 constant PUB_4_Y = 12455636581371328050369695832676869257569189466133302902328476078575183979595;
    uint256 constant PUB_5_X = 7958475172690672829748492988249572878991386204317802601773893461459343248265;
    uint256 constant PUB_5_Y = 9389532256402089527316365862774204613814889098999755171094402267653410979428;
    uint256 constant PUB_6_X = 419128560521272670682877864475153060065413905152921277659659034298547869000;
    uint256 constant PUB_6_Y = 6372879698814001335093396030028684780920775852316901267801270385648097815751;
    uint256 constant PUB_7_X = 14500643290997515859841688368473103213584819763054667829457080972623677945522;
    uint256 constant PUB_7_Y = 2510796519049536186265902727442848210238425627035721385815340217938059673026;
    uint256 constant PUB_8_X = 13714803245493495615164497351979481060151182221430936280030201759204693767852;
    uint256 constant PUB_8_Y = 8075942977904763478797072463106634173664284978886652995431784609076641694121;
    uint256 constant PUB_9_X = 7162979031372881348274523157895919204553728918944491306315207279319570802655;
    uint256 constant PUB_9_Y = 19440880952494561475061748206555677803194228459619058424983217515493100789065;

    /// Compute the public input linear combination.
    /// @notice Reverts with PublicInputNotInField if the input is not in the field.
    /// @notice Computes the multi-scalar-multiplication of the public input
    /// elements and the verification key including the constant term.
    /// @param input The public inputs. These are elements of the scalar field Fr.
    /// @return x The X coordinate of the resulting G1 point.
    /// @return y The Y coordinate of the resulting G1 point.
    function publicInputMSM(
        uint256[9] memory input,
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
            mstore(add(g, 0x40), s)
            success := and(success, lt(s, R))
            success := and(success, staticcall(gas(), PRECOMPILE_MUL, g, 0x60, g, 0x40))
            success := and(success, staticcall(gas(), PRECOMPILE_ADD, f, 0x80, f, 0x40))
            mstore(g, PUB_9_X)
            mstore(add(g, 0x20), PUB_9_Y)

            s := mload(add(input, 288))
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
        uint256[9] memory input
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
}
