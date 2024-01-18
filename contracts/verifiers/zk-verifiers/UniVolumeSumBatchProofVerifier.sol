// SPDX-License-Identifier: AML
//
// Copyright 2017 Christian Reitwiessner
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

// 2019 OKIMS

pragma solidity ^0.8.0;

library Pairing {
    uint256 constant PRIME_Q = 21888242871839275222246405745257275088696311157297823662689037894645226208583;

    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    // Encoding of field elements is: X[0] * z + X[1]
    struct G2Point {
        uint256[2] X;
        uint256[2] Y;
    }

    /*
     * @return The negation of p, i.e. p.plus(p.negate()) should be zero.
     */
    function negate(G1Point memory p) internal pure returns (G1Point memory) {
        // The prime q in the base field F_q for G1
        if (p.X == 0 && p.Y == 0) {
            return G1Point(0, 0);
        } else {
            return G1Point(p.X, PRIME_Q - (p.Y % PRIME_Q));
        }
    }

    /*
     * @return The sum of two points of G1
     */
    function plus(G1Point memory p1, G1Point memory p2) internal view returns (G1Point memory r) {
        uint256[4] memory input;
        input[0] = p1.X;
        input[1] = p1.Y;
        input[2] = p2.X;
        input[3] = p2.Y;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 6, input, 0xc0, r, 0x60)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }

        require(success, "pairing-add-failed");
    }

    /*
     * Same as plus but accepts raw input instead of struct
     * @return The sum of two points of G1, one is represented as array
     */
    function plus_raw(uint256[4] memory input, G1Point memory r) internal view {
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 6, input, 0xc0, r, 0x60)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }

        require(success, "pairing-add-failed");
    }

    /*
     * @return The product of a point on G1 and a scalar, i.e.
     *         p == p.scalar_mul(1) and p.plus(p) == p.scalar_mul(2) for all
     *         points p.
     */
    function scalar_mul(G1Point memory p, uint256 s) internal view returns (G1Point memory r) {
        uint256[3] memory input;
        input[0] = p.X;
        input[1] = p.Y;
        input[2] = s;
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 7, input, 0x80, r, 0x60)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }
        require(success, "pairing-mul-failed");
    }

    /*
     * Same as scalar_mul but accepts raw input instead of struct,
     * Which avoid extra allocation. provided input can be allocated outside and re-used multiple times
     */
    function scalar_mul_raw(uint256[3] memory input, G1Point memory r) internal view {
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 7, input, 0x80, r, 0x60)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }
        require(success, "pairing-mul-failed");
    }

    /* @return The result of computing the pairing check
     *         e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1
     *         For example,
     *         pairing([P1(), P1().negate()], [P2(), P2()]) should return true.
     */
    function pairing(
        G1Point memory a1,
        G2Point memory a2,
        G1Point memory b1,
        G2Point memory b2,
        G1Point memory c1,
        G2Point memory c2,
        G1Point memory d1,
        G2Point memory d2
    ) internal view returns (bool) {
        G1Point[4] memory p1 = [a1, b1, c1, d1];
        G2Point[4] memory p2 = [a2, b2, c2, d2];
        uint256 inputSize = 24;
        uint256[] memory input = new uint256[](inputSize);

        for (uint256 i = 0; i < 4; i++) {
            uint256 j = i * 6;
            input[j + 0] = p1[i].X;
            input[j + 1] = p1[i].Y;
            input[j + 2] = p2[i].X[0];
            input[j + 3] = p2[i].X[1];
            input[j + 4] = p2[i].Y[0];
            input[j + 5] = p2[i].Y[1];
        }

        uint256[1] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 8, add(input, 0x20), mul(inputSize, 0x20), out, 0x20)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }

        require(success, "pairing-opcode-failed");

        return out[0] != 0;
    }
}

contract UniswapSumVolumeVerifier {
    using Pairing for *;

    uint256 constant SNARK_SCALAR_FIELD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 constant PRIME_Q = 21888242871839275222246405745257275088696311157297823662689037894645226208583;

    struct VerifyingKey {
        Pairing.G1Point alfa1;
        Pairing.G2Point beta2;
        Pairing.G2Point gamma2;
        Pairing.G2Point delta2;
        // []G1Point IC (K in gnark) appears directly in verifyProof
    }

    struct Proof {
        Pairing.G1Point A;
        Pairing.G2Point B;
        Pairing.G1Point C;
        Pairing.G1Point Commit;
    }

    function verifyingKey() internal pure returns (VerifyingKey memory vk) {
        vk.alfa1 = Pairing.G1Point(
            uint256(3495230055929497346524681914586082958637336070006425695650037337616459459713),
            uint256(13753882227354172537719220698401028787447604517083211931696250241972767344257)
        );
        vk.beta2 = Pairing.G2Point(
            [
                uint256(9529526955822643166713138918618344102520995486763026962735335599892251106838),
                uint256(2464465499724313558415179073776930898864050287713306974091404815724545909227)
            ],
            [
                uint256(15074798363319050879741497669624606757963157190142504005523462594822261102450),
                uint256(464914921358567082785856687638465550489736092792418088341895894511280945913)
            ]
        );
        vk.gamma2 = Pairing.G2Point(
            [
                uint256(18323292397109129422382196971115802720688348570662216906371630979463998663255),
                uint256(18384364197079121392172955779335315940276562614669669157553064126292143355696)
            ],
            [
                uint256(16531624336128789285484541487434779060606819552202803713762390100841748140653),
                uint256(11518011523246896632165372896881302004128909393851548096877484216038904664262)
            ]
        );
        vk.delta2 = Pairing.G2Point(
            [
                uint256(1178643840928344118367511210012490217317617715502861347657476471907416848261),
                uint256(13818638828224359334160545290781714007429988703928190519308381000249416073647)
            ],
            [
                uint256(15645916288686682818855820599634260251448564762122602777476548616329880425235),
                uint256(19719121708103678947618130808572978408986723003707401372848514085169315456458)
            ]
        );
    }

    // accumulate scalarMul(mul_input) into q
    // that is computes sets q = (mul_input[0:2] * mul_input[3]) + q
    function accumulate(
        uint256[3] memory mul_input,
        Pairing.G1Point memory p,
        uint256[4] memory buffer,
        Pairing.G1Point memory q
    ) internal view {
        // computes p = mul_input[0:2] * mul_input[3]
        Pairing.scalar_mul_raw(mul_input, p);

        // point addition inputs
        buffer[0] = q.X;
        buffer[1] = q.Y;
        buffer[2] = p.X;
        buffer[3] = p.Y;

        // q = p + q
        Pairing.plus_raw(buffer, q);
    }

    /*
     * @returns Whether the proof is valid given the hardcoded verifying key
     *          above and the public inputs
     */
    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[32] memory input
    ) public view returns (bool r) {
        Proof memory proof;
        proof.A = Pairing.G1Point(a[0], a[1]);
        proof.B = Pairing.G2Point([b[0][0], b[0][1]], [b[1][0], b[1][1]]);
        proof.C = Pairing.G1Point(c[0], c[1]);
        proof.Commit = Pairing.G1Point(commit[0], commit[1]);

        // Make sure that proof.A, B, and C are each less than the prime q
        require(proof.A.X < PRIME_Q, "verifier-aX-gte-prime-q");
        require(proof.A.Y < PRIME_Q, "verifier-aY-gte-prime-q");

        require(proof.B.X[0] < PRIME_Q, "verifier-bX0-gte-prime-q");
        require(proof.B.Y[0] < PRIME_Q, "verifier-bY0-gte-prime-q");

        require(proof.B.X[1] < PRIME_Q, "verifier-bX1-gte-prime-q");
        require(proof.B.Y[1] < PRIME_Q, "verifier-bY1-gte-prime-q");

        require(proof.C.X < PRIME_Q, "verifier-cX-gte-prime-q");
        require(proof.C.Y < PRIME_Q, "verifier-cY-gte-prime-q");

        // Make sure that every input is less than the snark scalar field
        for (uint256 i = 0; i < input.length; i++) {
            require(input[i] < SNARK_SCALAR_FIELD, "verifier-gte-snark-scalar-field");
        }

        VerifyingKey memory vk = verifyingKey();

        // Compute the linear combination vk_x
        Pairing.G1Point memory vk_x = Pairing.G1Point(0, 0);

        // Buffer reused for addition p1 + p2 to avoid memory allocations
        // [0:2] -> p1.X, p1.Y ; [2:4] -> p2.X, p2.Y
        uint256[4] memory add_input;

        // Buffer reused for multiplication p1 * s
        // [0:2] -> p1.X, p1.Y ; [3] -> s
        uint256[3] memory mul_input;

        // temporary point to avoid extra allocations in accumulate
        Pairing.G1Point memory q = Pairing.G1Point(0, 0);

        vk_x.X = uint256(3620881018601535948764113707540002800766616898056894894793980318744329588476); // vk.K[0].X
        vk_x.Y = uint256(300646094362362905849323155903199192485967625441437051640838123294159659928); // vk.K[0].Y
        mul_input[0] = uint256(16511302356454745997404439488652451285736881078509641937140907170323794923078); // vk.K[1].X
        mul_input[1] = uint256(13979873842895528414815874777928984249793359493678552666246089433064448940346); // vk.K[1].Y
        mul_input[2] = input[0];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[1] * input[0]
        mul_input[0] = uint256(16162029783796135716466504914415131452437585361381180674902989407359087565901); // vk.K[2].X
        mul_input[1] = uint256(18404073456915309307997395924841781067447142183725612029369873277840707704242); // vk.K[2].Y
        mul_input[2] = input[1];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[2] * input[1]
        mul_input[0] = uint256(17588141939748872355433032015033149489504451304409308678359904313587558461364); // vk.K[3].X
        mul_input[1] = uint256(9877847949926987634696640311782491942243506139179306379315549936549426326347); // vk.K[3].Y
        mul_input[2] = input[2];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[3] * input[2]
        mul_input[0] = uint256(13840149287499207252672532729422031401662411635781184637774485129612931893541); // vk.K[4].X
        mul_input[1] = uint256(16834569386584243740924735600886192412772276914114394862809581776523075778627); // vk.K[4].Y
        mul_input[2] = input[3];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[4] * input[3]
        mul_input[0] = uint256(6151843274529598689884421931187025413603492686362392064154916369852980736521); // vk.K[5].X
        mul_input[1] = uint256(4014316135042207688405379224803416428664048236986628773758367403964373969868); // vk.K[5].Y
        mul_input[2] = input[4];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[5] * input[4]
        mul_input[0] = uint256(7165833303457693223071003934993715758733804129346115451694979217601043209283); // vk.K[6].X
        mul_input[1] = uint256(19168551232118017312772818610728980714401792932292485086395374991956910041868); // vk.K[6].Y
        mul_input[2] = input[5];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[6] * input[5]
        mul_input[0] = uint256(6362227164324135345687793357279968687946789005128442003685288831861831948394); // vk.K[7].X
        mul_input[1] = uint256(12031949098767277205924176130036314514438398292628177002721092282931372678161); // vk.K[7].Y
        mul_input[2] = input[6];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[7] * input[6]
        mul_input[0] = uint256(15209330211935576444815377960564220344951489270571527938412482719465340313902); // vk.K[8].X
        mul_input[1] = uint256(20807791260801661491234475993168497166008740219234654180930999999351992930171); // vk.K[8].Y
        mul_input[2] = input[7];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[8] * input[7]
        mul_input[0] = uint256(14314021892362833109087291405109828555589796733615058590306751151870184262115); // vk.K[9].X
        mul_input[1] = uint256(7192653808322235714448055382401917433034818334138505171700358186537609790557); // vk.K[9].Y
        mul_input[2] = input[8];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[9] * input[8]
        mul_input[0] = uint256(7809147924148156244723641472623721051727302559592529158554729907599176474080); // vk.K[10].X
        mul_input[1] = uint256(12712832040014178377122172427734380347809301986206537193310265116667640011695); // vk.K[10].Y
        mul_input[2] = input[9];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[10] * input[9]
        mul_input[0] = uint256(17133538140823213822017317959075668391802621430713819540959998671194416150096); // vk.K[11].X
        mul_input[1] = uint256(14251353948801439234795711765324997782831858613453830197625628218309938050276); // vk.K[11].Y
        mul_input[2] = input[10];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[11] * input[10]
        mul_input[0] = uint256(19321244880305504341586592299223834682372170730216799854464099916364976031460); // vk.K[12].X
        mul_input[1] = uint256(141384221847407634837094408671734061765677717785191809154136040723129995606); // vk.K[12].Y
        mul_input[2] = input[11];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[12] * input[11]
        mul_input[0] = uint256(336041225158662771537518670541160893062564048460575382926533612570414170097); // vk.K[13].X
        mul_input[1] = uint256(10918982459799224835368294012459604256559989043377595847134969583125197754454); // vk.K[13].Y
        mul_input[2] = input[12];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[13] * input[12]
        mul_input[0] = uint256(21559709156099493975346282197006554290950742786035406077663071438344705391972); // vk.K[14].X
        mul_input[1] = uint256(19072116008698743271864462955420245819042747669622547950758672803329132619632); // vk.K[14].Y
        mul_input[2] = input[13];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[14] * input[13]
        mul_input[0] = uint256(1340235371816563427441312786195698324864478195647701208840959896844508970166); // vk.K[15].X
        mul_input[1] = uint256(1766272641737323144117128246930196764749125329839272104918786931835236094743); // vk.K[15].Y
        mul_input[2] = input[14];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[15] * input[14]
        mul_input[0] = uint256(1550560755218617380525601885133107571746559150968028373774626947585083549330); // vk.K[16].X
        mul_input[1] = uint256(315691894405950573915609253467540777407485489995150855193641188260766513259); // vk.K[16].Y
        mul_input[2] = input[15];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[16] * input[15]
        mul_input[0] = uint256(3325392570782813063019753948809962527981215536614554945310077189924843829938); // vk.K[17].X
        mul_input[1] = uint256(3327508066506625644839072551534339167929243325291770723378112767737174477812); // vk.K[17].Y
        mul_input[2] = input[16];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[17] * input[16]
        mul_input[0] = uint256(18617173141205640870225871488992131079850846831461448020396530714052859160556); // vk.K[18].X
        mul_input[1] = uint256(11753901861308170637034017733962497698129672994844671545969154180216281571381); // vk.K[18].Y
        mul_input[2] = input[17];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[18] * input[17]
        mul_input[0] = uint256(2985829483032027990071860677283289958023154954973569614325287799675897390064); // vk.K[19].X
        mul_input[1] = uint256(15547351916694485388834924403454110597971696947160846187371664775904233218143); // vk.K[19].Y
        mul_input[2] = input[18];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[19] * input[18]
        mul_input[0] = uint256(17635087365283006812368883212694639641651576013614671691719739651672447538335); // vk.K[20].X
        mul_input[1] = uint256(159513181283689860665148961521746235355754376928613106765249341450537136077); // vk.K[20].Y
        mul_input[2] = input[19];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[20] * input[19]
        mul_input[0] = uint256(13128592090707926521398592682494506623484338764919803356525822853041568077640); // vk.K[21].X
        mul_input[1] = uint256(19751765373496258733563020832667218693470610630042194679484661406994968400378); // vk.K[21].Y
        mul_input[2] = input[20];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[21] * input[20]
        mul_input[0] = uint256(5795729047155020399123458662564619918364702991024951706553490717912946628101); // vk.K[22].X
        mul_input[1] = uint256(18802172385998807309968221656568966673946600901576864587595601741558818122443); // vk.K[22].Y
        mul_input[2] = input[21];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[22] * input[21]
        mul_input[0] = uint256(1676216849920825711058724666933328990355447541597809238785336864343570717468); // vk.K[23].X
        mul_input[1] = uint256(18656946087511866723377584017711547850258056256544067851500297975868096576811); // vk.K[23].Y
        mul_input[2] = input[22];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[23] * input[22]
        mul_input[0] = uint256(19660198267422159372436389965731885882657272213099326871859624121077065368735); // vk.K[24].X
        mul_input[1] = uint256(14803765777107938189350319449638295803862678259382785661220716872771581361556); // vk.K[24].Y
        mul_input[2] = input[23];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[24] * input[23]
        mul_input[0] = uint256(4280746727217651510425610381921096212467127748186587615401622671385714141734); // vk.K[25].X
        mul_input[1] = uint256(11528548281267871998808903656347750089935207126757086660587422708609484012036); // vk.K[25].Y
        mul_input[2] = input[24];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[25] * input[24]
        mul_input[0] = uint256(20835651915617114059097998094511865344876734911114886492858968028937054347634); // vk.K[26].X
        mul_input[1] = uint256(21005185482816714532869566191931759396843276758174469690851139973501863758799); // vk.K[26].Y
        mul_input[2] = input[25];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[26] * input[25]
        mul_input[0] = uint256(11948950872831022709006416557742459104185833843389517177886596971420893680229); // vk.K[27].X
        mul_input[1] = uint256(7335830878886154292349766117972418767256057510393424815198592316149139043369); // vk.K[27].Y
        mul_input[2] = input[26];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[27] * input[26]
        mul_input[0] = uint256(6103065927741102049934457912997815151724470876271676287037925321105934851540); // vk.K[28].X
        mul_input[1] = uint256(11207711290502750990419645547301772745865249660420004695612149539847127734879); // vk.K[28].Y
        mul_input[2] = input[27];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[28] * input[27]
        mul_input[0] = uint256(12699939612392256041625759284144088819738222452837163635780627107945402415093); // vk.K[29].X
        mul_input[1] = uint256(19449441644528026683652941000116088958281940845304758060607805991413978736727); // vk.K[29].Y
        mul_input[2] = input[28];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[29] * input[28]
        mul_input[0] = uint256(8242362142410956355466962576597726871365630778446659381373761481130205949351); // vk.K[30].X
        mul_input[1] = uint256(10399171122301269786492415753678662204349768547233633180958874533751659907005); // vk.K[30].Y
        mul_input[2] = input[29];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[30] * input[29]
        mul_input[0] = uint256(6336371553116220172813403619960638546995537910090005969827728669751067599096); // vk.K[31].X
        mul_input[1] = uint256(7069996817011795529448190724021782250820806678981299879138939334802128742262); // vk.K[31].Y
        mul_input[2] = input[30];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[31] * input[30]
        mul_input[0] = uint256(11882112752134222842030567033273257549236676785328489485893694350900354933480); // vk.K[32].X
        mul_input[1] = uint256(15657966439730622678882171755431979942667106625900348982775656465677049366076); // vk.K[32].Y
        mul_input[2] = input[31];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[32] * input[31]
        if (commit[0] != 0 || commit[1] != 0) {
            vk_x = Pairing.plus(vk_x, proof.Commit);
        }

        return
            Pairing.pairing(Pairing.negate(proof.A), proof.B, vk.alfa1, vk.beta2, vk_x, vk.gamma2, proof.C, vk.delta2);
    }

    function verifyRaw(bytes calldata proofData) external view returns (bool) {
        uint256[2] memory a;
        a[0] = uint256(bytes32(proofData[:32]));
        a[1] = uint256(bytes32(proofData[32:64]));
        uint256[2][2] memory b;
        b[0][0] = uint256(bytes32(proofData[64:96]));
        b[0][1] = uint256(bytes32(proofData[96:128]));
        b[1][0] = uint256(bytes32(proofData[128:160]));
        b[1][1] = uint256(bytes32(proofData[160:192]));
        uint256[2] memory c;
        c[0] = uint256(bytes32(proofData[192:224]));
        c[1] = uint256(bytes32(proofData[224:256]));
        uint256[2] memory commit;
        commit[0] = uint256(bytes32(proofData[256:288]));
        commit[1] = uint256(bytes32(proofData[288:320]));
        uint256[32] memory input;
        input[31] = uint256(bytes32(proofData[320:352])); //input last one is cpub

        input[0] = uint256(uint64(bytes8(proofData[376:384]))); // emulated field with 6 limbs
        input[1] = uint256(uint64(bytes8(proofData[368:376])));
        input[2] = uint256(uint64(bytes8(proofData[360:368])));
        input[3] = uint256(uint64(bytes8(proofData[352:360])));
        input[4] = 0;
        input[5] = 0;

        input[6] = uint256(uint64(bytes8(proofData[408:416])));
        input[7] = uint256(uint64(bytes8(proofData[400:408])));
        input[8] = uint256(uint64(bytes8(proofData[392:400])));
        input[9] = uint256(uint64(bytes8(proofData[384:392])));
        input[10] = 0;
        input[11] = 0;

        input[12] = uint256(uint64(bytes8(proofData[440:448])));
        input[13] = uint256(uint64(bytes8(proofData[432:440])));
        input[14] = uint256(uint64(bytes8(proofData[424:432])));
        input[15] = uint256(uint64(bytes8(proofData[416:424])));
        input[16] = 0;
        input[17] = 0;

        input[18] = uint256(uint64(bytes8(proofData[472:480])));
        input[19] = uint256(uint64(bytes8(proofData[464:472])));
        input[20] = uint256(uint64(bytes8(proofData[456:464])));
        input[21] = uint256(uint64(bytes8(proofData[448:456])));
        input[22] = 0;
        input[23] = 0;

        input[24] = uint256(uint64(bytes8(proofData[504:512])));
        input[25] = uint256(uint64(bytes8(proofData[496:504])));
        input[26] = uint256(uint64(bytes8(proofData[488:496])));
        input[27] = uint256(uint64(bytes8(proofData[480:488])));
        input[28] = 0;
        input[29] = 0;

        input[30] = uint256(bytes32(proofData[512:544]));

        return verifyProof(a, b, c, commit, input);
    }
}
