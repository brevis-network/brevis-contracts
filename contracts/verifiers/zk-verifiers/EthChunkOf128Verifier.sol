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

import "../interfaces/IZkpVerifier.sol";

import "./common/Pairing.sol";
import "./common/Constants.sol";
import "./common/Common.sol";

pragma solidity ^0.8.0;

contract EthChunkOf128Verifier is IZkpVerifier {
    using Pairing for *;

    function verifyingKey() internal pure returns (Common.VerifyingKey memory vk) {
        vk.alfa1 = Pairing.G1Point(
            uint256(3731296059100206882180451755503172983109296248442025944723039592906698751900),
            uint256(259282599373131156598819103264659325586675381205265829375869139666859279077)
        );
        vk.beta2 = Pairing.G2Point(
            [
                uint256(1542547783169253521197735286554664372551314003473574629092788070944269927009),
                uint256(3442843350567772397992411822441194496033000804528083490373991767861898380921)
            ],
            [
                uint256(548716800731408953921339589333269049204399133705913270541257072650667905637),
                uint256(19884421027672845815718445738066364430735990147602958067065404264457524686219)
            ]
        );
        vk.gamma2 = Pairing.G2Point(
            [
                uint256(11891988479373600508373950254994848776611254114498681881881793265865565715386),
                uint256(19931296744367177659683198721889325174160042107805115279855692658605438538701)
            ],
            [
                uint256(16677912734950263118736987260922842186425341971840174170121099525719169899146),
                uint256(7417603967557513391474517980599665601132457400115579956964885296221617800673)
            ]
        );
        vk.delta2 = Pairing.G2Point(
            [
                uint256(16517284527623427684849545816824631314997223078114802689115370057079771769284),
                uint256(9556334973004218826551566300456739568039857994289235505930178517047984960309)
            ],
            [
                uint256(11378844714980444222011796238800318515164232281952868936120207698945684705011),
                uint256(16717114042048232586709438632101306668327363244916399631576013967002311482733)
            ]
        );
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
        uint256[8] memory input
    ) public view returns (bool r) {
        Common.Proof memory proof;
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

        Common.VerifyingKey memory vk = verifyingKey();

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

        vk_x.X = uint256(20771483654708554343899445707962704346619674363636303975277986884195939178109); // vk.K[0].X
        vk_x.Y = uint256(7136981524249642767758675836778771393360023884114860800337043409235713151709); // vk.K[0].Y
        mul_input[0] = uint256(19233685918932192923050625893183018999852678211997471298974889334412747883339); // vk.K[1].X
        mul_input[1] = uint256(10573418048246313728599286941747928760901755928929380194465994346023631757706); // vk.K[1].Y
        mul_input[2] = input[0];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[1] * input[0]
        mul_input[0] = uint256(14573802110847817293393990384189386188819275610796892844002198465549922624165); // vk.K[2].X
        mul_input[1] = uint256(6366177769334632983948241800448594297640371227299129420889272762628401203322); // vk.K[2].Y
        mul_input[2] = input[1];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[2] * input[1]
        mul_input[0] = uint256(4423470128720469023828937120320047480914076991874402659152085174133214111513); // vk.K[3].X
        mul_input[1] = uint256(16661537357790150960310016685835020561095798528309478888052879927100385369883); // vk.K[3].Y
        mul_input[2] = input[2];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[3] * input[2]
        mul_input[0] = uint256(20954264198739863123465273905250064029734563486205900387553306667464142762653); // vk.K[4].X
        mul_input[1] = uint256(15400479026677042446536213348710994852526762940505119021682575524837659837800); // vk.K[4].Y
        mul_input[2] = input[3];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[4] * input[3]
        mul_input[0] = uint256(0); // vk.K[5].X
        mul_input[1] = uint256(0); // vk.K[5].Y
        mul_input[2] = input[4];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[5] * input[4]
        mul_input[0] = uint256(0); // vk.K[6].X
        mul_input[1] = uint256(0); // vk.K[6].Y
        mul_input[2] = input[5];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[6] * input[5]
        mul_input[0] = uint256(15581619746676589215919086055186955357437498065118852273796328444039455983836); // vk.K[7].X
        mul_input[1] = uint256(2079225735174412294634686160304546497644529522569168218992038814057683036684); // vk.K[7].Y
        mul_input[2] = input[6];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[7] * input[6]
        mul_input[0] = uint256(5271944653916881751571075049622758937543583004433508630909066776841583056730); // vk.K[8].X
        mul_input[1] = uint256(5018511801199739199792530205278441732520613352432208789698314328778478334533); // vk.K[8].Y
        mul_input[2] = input[7];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[8] * input[7]
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
        uint256[8] memory input;
        input[0] = uint256(bytes32(proofData[320:352]));
        input[1] = uint256(bytes32(proofData[352:384]));
        input[2] = uint256(bytes32(proofData[384:416]));
        input[3] = uint256(bytes32(proofData[416:448]));
        input[4] = uint256(bytes32(proofData[448:480]));
        input[5] = uint256(bytes32(proofData[480:512]));
        input[6] = uint256(bytes32(proofData[512:544]));
        input[7] = uint256(bytes32(proofData[544:576]));

        return verifyProof(a, b, c, commit, input);
    }
}
