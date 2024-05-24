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

contract BatchZkProofVerifier {
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
            uint256(15238308597703998611024237714335999796365016940312182319280395757706967719293),
            uint256(10219911876785786802002801370470398321650839519404409433848585724096395912813)
        );
        vk.beta2 = Pairing.G2Point(
            [
                uint256(16724817789808363935019716420097303583783816372788140679808387107301975309262),
                uint256(5778445566676027117037680092491029447507126512417385078353864383782020250397)
            ],
            [
                uint256(14058479183418968750184727536762806238931313078241995735760477278048306135851),
                uint256(14109384065032405558523914752328776367729991421617607120062845893114961770297)
            ]
        );
        vk.gamma2 = Pairing.G2Point(
            [
                uint256(4849374573467867131743034891183148299921951523707019746444468744028093260842),
                uint256(19198999624409799465959639686144554617014642257205908251362417016643142369977)
            ],
            [
                uint256(1572433243919907845314059575345991674289709827340315857346388493812775131387),
                uint256(19144661012025957028820388582997525098651337600859903355485005371794448709326)
            ]
        );
        vk.delta2 = Pairing.G2Point(
            [
                uint256(8675268962751885909765400209018563103958042547099968862734739916509027970783),
                uint256(21334095358979846169682840614747683942013654095180122458302288133636118962193)
            ],
            [
                uint256(20806544570626068324185407569920089157125215733452486800579671830019604406420),
                uint256(19320796094476058926763619514277544257927035899524445390040846751509191005713)
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
        uint256[14] memory input
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

        vk_x.X = uint256(4037306610926595749762075796746609702552307636987816780131058632335019249158); // vk.K[0].X
        vk_x.Y = uint256(17289477422575608593697732739310292230294333841423503415113301787380105217187); // vk.K[0].Y
        mul_input[0] = uint256(10037096222996786580190418719833957206453843602932270170664055188562568771027); // vk.K[1].X
        mul_input[1] = uint256(15320486340840873689164565592037333715383138535348014841809641122171818810036); // vk.K[1].Y
        mul_input[2] = input[0];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[1] * input[0]
        mul_input[0] = uint256(9896557402339891076396748574063140253246858155466527062441237909938610823754); // vk.K[2].X
        mul_input[1] = uint256(5164310492746652270672196862256673842804683282368846128824605431888437028315); // vk.K[2].Y
        mul_input[2] = input[1];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[2] * input[1]
        mul_input[0] = uint256(15705570664843676305244923084737310423934337534095385957113352057709739538579); // vk.K[3].X
        mul_input[1] = uint256(5657315344176937602330667514195982637595830588415888903950150885407794100439); // vk.K[3].Y
        mul_input[2] = input[2];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[3] * input[2]
        mul_input[0] = uint256(8491601659316077373107675776449969162172043482262510422712755851858933893670); // vk.K[4].X
        mul_input[1] = uint256(21665343136300961104345653535947413453573579433298425965534286592198041321210); // vk.K[4].Y
        mul_input[2] = input[3];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[4] * input[3]
        mul_input[0] = uint256(16487339789800816774433659496727708804625896411052663921373489568514791282316); // vk.K[5].X
        mul_input[1] = uint256(5392557275131571448254591082275082405708268619325456715079559988921856614947); // vk.K[5].Y
        mul_input[2] = input[4];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[5] * input[4]
        mul_input[0] = uint256(5760744510920437359711530772130596454523074712251763024976643549629349341029); // vk.K[6].X
        mul_input[1] = uint256(20561869527787897331749432992557176728073011740569310617207323846113958072098); // vk.K[6].Y
        mul_input[2] = input[5];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[6] * input[5]
        mul_input[0] = uint256(8938844239960091665696359855399717603558250728238750045539124092231985209597); // vk.K[7].X
        mul_input[1] = uint256(21085636969309297587411330198492565336581397650211483814114532934744000075984); // vk.K[7].Y
        mul_input[2] = input[6];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[7] * input[6]
        mul_input[0] = uint256(12117014598692072663950543644225023407901520491661942865593284582952909022782); // vk.K[8].X
        mul_input[1] = uint256(6563607654709757135971478433160655085349594669842762414756198183188100749916); // vk.K[8].Y
        mul_input[2] = input[7];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[8] * input[7]
        mul_input[0] = uint256(1876069974761176132969801959712873511146873644882078591170959518600066434088); // vk.K[9].X
        mul_input[1] = uint256(17254199576946899577090130294451675870206545287522281806930431198181899225222); // vk.K[9].Y
        mul_input[2] = input[8];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[9] * input[8]
        mul_input[0] = uint256(1701790252141422939101054766324470273552474316692868216771515791226546222304); // vk.K[10].X
        mul_input[1] = uint256(9908164510280999756965157678715899350039013940657119416333308690576851757152); // vk.K[10].Y
        mul_input[2] = input[9];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[10] * input[9]
        mul_input[0] = uint256(16591342763437189049323540618884654501108469055224263209450502142315414405548); // vk.K[11].X
        mul_input[1] = uint256(8178344437670034159374480625490785281447132078868100938476050070035463834108); // vk.K[11].Y
        mul_input[2] = input[10];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[11] * input[10]
        mul_input[0] = uint256(14703140412931501992135632763258522183956323585947888724160204311927087339174); // vk.K[12].X
        mul_input[1] = uint256(4725913873629201216398987571240145458971217726264472912550476744148652845198); // vk.K[12].Y
        mul_input[2] = input[11];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[12] * input[11]
        mul_input[0] = uint256(19018247690552355099153068992785533501519072800079329665523330167239555113095); // vk.K[13].X
        mul_input[1] = uint256(21134992072574410505124430108199227538266930095351542616114475670234988309129); // vk.K[13].Y
        mul_input[2] = input[12];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[13] * input[12]
        mul_input[0] = uint256(13640822404210647006784067638649610643922969553672280134406469238818114008997); // vk.K[14].X
        mul_input[1] = uint256(14210714846154522557123511382005640854325868943946774859420797463161407589257); // vk.K[14].Y
        mul_input[2] = input[13];
        accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[14] * input[13]
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
        uint256[14] memory input;
        input[13] = uint256(bytes32(proofData[320:352])); //input last one is cpub

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

        input[12] = uint256(bytes32(proofData[416:448]));

        return verifyProof(a, b, c, commit, input);
    }
}
