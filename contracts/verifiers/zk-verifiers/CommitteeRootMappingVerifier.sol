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

import "./common/Pairing.sol";
import "./common/Constants.sol";
import "./common/Common.sol";

contract CommitteeRootMappingVerifier {
    using Pairing for *;

    function verifyingKey1() private pure returns (Common.VerifyingKey memory vk) {
        vk.alfa1 = Pairing.G1Point(
            uint256(4625995678875839184227102343980957941553435037863367632170514069470978075482),
            uint256(7745472346822620166365670179252096531675980956628675937691452644416704349631)
        );
        vk.beta2 = Pairing.G2Point(
            [
                uint256(16133906051290029359415836500687237322258320219528941728637152470582101797559),
                uint256(9982592290591904397750372202184781412509742437847499064025507928193374812763)
            ],
            [
                uint256(20447084996628162496147084243623314997274147610235538549283479856317752366847),
                uint256(10652060452474388359080900509291122865897396777233890537481945528644944582649)
            ]
        );
        vk.gamma2 = Pairing.G2Point(
            [
                uint256(14205774305928561884273671177098614973303096843515928049981466843882075090453),
                uint256(6194647019556442694746623566240152360142526955447025858054760757353994166695)
            ],
            [
                uint256(720177741655577944140882804072173464461234581005085937938128202222496044348),
                uint256(15180859461535417805311870856102250988010112023636345871703449475067945282517)
            ]
        );
        vk.delta2 = Pairing.G2Point(
            [
                uint256(2075341858515413383107490988194322113274273165071779395977011288835607214232),
                uint256(21779842329350845285414688998042134519611654255235365675696046856282966715158)
            ],
            [
                uint256(4310903133868833376693610009744123646701594778591654462646551313203044329349),
                uint256(8934039419334185533732134671857943150009456594043165319933471646801466475060)
            ]
        );
    }

    /*
     * @returns Whether the proof is valid given the hardcoded verifying key
     *          above and the public inputs
     */
    function verifyCommitteeRootMappingProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[33] memory input
    ) public view returns (bool r) {
        Common.Proof memory proof;
        proof.A = Pairing.G1Point(a[0], a[1]);
        proof.B = Pairing.G2Point([b[0][0], b[0][1]], [b[1][0], b[1][1]]);
        proof.C = Pairing.G1Point(c[0], c[1]);

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

        Common.VerifyingKey memory vk = verifyingKey1();

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

        vk_x.X = uint256(20552480178503420105472757749758256930777503163697981232418248899738739436302); // vk.K[0].X
        vk_x.Y = uint256(21874644052683447189335205444383300629386926406593895540736254865290692175330); // vk.K[0].Y
        mul_input[0] = uint256(2419465434811246925970456918943785845329721675292263546063218305166868830301); // vk.K[1].X
        mul_input[1] = uint256(224414837900933448241244127409926533084118787014653569685139207760162770563); // vk.K[1].Y
        mul_input[2] = input[0];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[1] * input[0]
        mul_input[0] = uint256(20237582094031100903111658800543003981446659818658320070287593450545147260932); // vk.K[2].X
        mul_input[1] = uint256(9498936270692258262448475366106441134297508170417707117017418182506243810929); // vk.K[2].Y
        mul_input[2] = input[1];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[2] * input[1]
        mul_input[0] = uint256(21686431407509598771022896245105442713057757617842882639916055310118549735455); // vk.K[3].X
        mul_input[1] = uint256(18587475580363988870337779644366478839186363821430368900189877147428300473925); // vk.K[3].Y
        mul_input[2] = input[2];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[3] * input[2]
        mul_input[0] = uint256(4190323520659374373641761976155873288531237902311450285189695279890286046705); // vk.K[4].X
        mul_input[1] = uint256(8044837422277408304807431419004307582225876792722238390231063677200212676904); // vk.K[4].Y
        mul_input[2] = input[3];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[4] * input[3]
        mul_input[0] = uint256(2652622379392044318082038991710242104342228971779836360052332572087628421201); // vk.K[5].X
        mul_input[1] = uint256(406860223885500452975843681654102213552218004006375181643914225581644355831); // vk.K[5].Y
        mul_input[2] = input[4];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[5] * input[4]
        mul_input[0] = uint256(6057918943482398019697118579402810827270820344972408585195554580949838772589); // vk.K[6].X
        mul_input[1] = uint256(5060377211716517826689871487122513539243478809827924728351043431363438746264); // vk.K[6].Y
        mul_input[2] = input[5];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[6] * input[5]
        mul_input[0] = uint256(3687702938753468537462497928786246235243684882237823906440956320376037461563); // vk.K[7].X
        mul_input[1] = uint256(1208686206265801496727901652555022795816232879429721718984614404615694111083); // vk.K[7].Y
        mul_input[2] = input[6];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[7] * input[6]
        mul_input[0] = uint256(11710614008104008246282861623202747769385618500144669344475214097509828684593); // vk.K[8].X
        mul_input[1] = uint256(5065836875015911503963590142184023993405575153173968399414211124081308802733); // vk.K[8].Y
        mul_input[2] = input[7];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[8] * input[7]
        mul_input[0] = uint256(544404787870686540959136485911507545335221912755631162384362056307403363961); // vk.K[9].X
        mul_input[1] = uint256(2345869893991024974950769006226913293849021455623995373213361343160988457751); // vk.K[9].Y
        mul_input[2] = input[8];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[9] * input[8]
        mul_input[0] = uint256(2209389364146280288951908471817129375141759543141552284740145921306411049406); // vk.K[10].X
        mul_input[1] = uint256(9042259349973012497614444570261244747029883119587798835387806797437998198439); // vk.K[10].Y
        mul_input[2] = input[9];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[10] * input[9]
        mul_input[0] = uint256(5329749415213215279150815169017002879660981652478899879932293459107956198272); // vk.K[11].X
        mul_input[1] = uint256(1269241490245981774317800992176787362067828005821041854984670483140659381972); // vk.K[11].Y
        mul_input[2] = input[10];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[11] * input[10]
        mul_input[0] = uint256(4943793813361186613838184379271444100858893499387902057809188182513783485846); // vk.K[12].X
        mul_input[1] = uint256(9275690329715777324103642003412034648418070562981699307031172873365106078545); // vk.K[12].Y
        mul_input[2] = input[11];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[12] * input[11]
        mul_input[0] = uint256(12729498268013982038852548044563174517696421517428254680176367740849220266709); // vk.K[13].X
        mul_input[1] = uint256(7546589572574852665535613703939452808321148398493753492131740521875420626909); // vk.K[13].Y
        mul_input[2] = input[12];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[13] * input[12]
        mul_input[0] = uint256(9333085734209829031122997463964247926338222396225058317742956090059153031592); // vk.K[14].X
        mul_input[1] = uint256(4043123151744068929699760825751364162242644369436915556155534564396462636465); // vk.K[14].Y
        mul_input[2] = input[13];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[14] * input[13]
        mul_input[0] = uint256(3698686717106590496650986585007797659650605418055308742433506982460764492730); // vk.K[15].X
        mul_input[1] = uint256(9179617523334761636265229485895993306228474412981061346064728177636515751968); // vk.K[15].Y
        mul_input[2] = input[14];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[15] * input[14]
        mul_input[0] = uint256(15521850592660810728436432508964964041834382081916421935161893482249902884387); // vk.K[16].X
        mul_input[1] = uint256(5449901017503560405242500659614777785834634841695450826672263537767974100219); // vk.K[16].Y
        mul_input[2] = input[15];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[16] * input[15]
        mul_input[0] = uint256(20102906107256118088436001377164222872704427733042089123636772674622559816716); // vk.K[17].X
        mul_input[1] = uint256(12498854682789208487185327670228889940757953195079617884138082484806034246784); // vk.K[17].Y
        mul_input[2] = input[16];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[17] * input[16]
        mul_input[0] = uint256(9455841695606475800176819517076441035373288808813491909032241063291148788930); // vk.K[18].X
        mul_input[1] = uint256(5760837211388967374979882368837632355372021503182733102840122488409476353553); // vk.K[18].Y
        mul_input[2] = input[17];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[18] * input[17]
        mul_input[0] = uint256(1446991383552871512734012954692326283314249519870143612600792757960520781278); // vk.K[19].X
        mul_input[1] = uint256(9834470268591454131741863361237282178002203711883219940241340793939995038767); // vk.K[19].Y
        mul_input[2] = input[18];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[19] * input[18]
        mul_input[0] = uint256(1059357485615144832413353841149751938707953460935522780194084907196702253731); // vk.K[20].X
        mul_input[1] = uint256(10815569476482003993766770423385630209543201328293985898718647153832884016017); // vk.K[20].Y
        mul_input[2] = input[19];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[20] * input[19]
        mul_input[0] = uint256(7433245970798099608332042376067563625513377267096206052430761000239299269566); // vk.K[21].X
        mul_input[1] = uint256(12741834193487831964894419250386047831198155854304448707022734193570700410821); // vk.K[21].Y
        mul_input[2] = input[20];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[21] * input[20]
        mul_input[0] = uint256(8648224634225961431530490440075030243542463588893169022877288417966438069777); // vk.K[22].X
        mul_input[1] = uint256(16540610842070555034877322476339116325277917786072762919274678110762172365508); // vk.K[22].Y
        mul_input[2] = input[21];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[22] * input[21]
        mul_input[0] = uint256(16908648218709781420138074614673957046034248547088691701260866141074824824919); // vk.K[23].X
        mul_input[1] = uint256(20980273428957053574278769661356962533672481733183512384951407225298181139010); // vk.K[23].Y
        mul_input[2] = input[22];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[23] * input[22]
        mul_input[0] = uint256(20934252423600973663175987808002009495824217352345209099319606411155218995932); // vk.K[24].X
        mul_input[1] = uint256(9987927206019920292163635872827487165514620975045002130414615160938718715749); // vk.K[24].Y
        mul_input[2] = input[23];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[24] * input[23]
        mul_input[0] = uint256(9602737041922572073213386264444643405537681976425696147506639312256088109115); // vk.K[25].X
        mul_input[1] = uint256(5030838233095700558123674330813925820525997306253984515590208165812087573689); // vk.K[25].Y
        mul_input[2] = input[24];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[25] * input[24]
        mul_input[0] = uint256(20088832978375886523413495106079569725269630343909328763686584839952109161933); // vk.K[26].X
        mul_input[1] = uint256(8311397503596416021728705867174781915782892850820869993294450806608979293432); // vk.K[26].Y
        mul_input[2] = input[25];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[26] * input[25]
        mul_input[0] = uint256(15729968276421379987872047780863974781795109674620595131198333451598870913212); // vk.K[27].X
        mul_input[1] = uint256(11755585053459843437112320638816029546922021127794137048950074210155862560131); // vk.K[27].Y
        mul_input[2] = input[26];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[27] * input[26]
        mul_input[0] = uint256(5783930197610380391486193680213891260111080319012345925622032738683845648623); // vk.K[28].X
        mul_input[1] = uint256(15914052883335873414184612431500787588848752068877353731383121390711998005745); // vk.K[28].Y
        mul_input[2] = input[27];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[28] * input[27]
        mul_input[0] = uint256(13576027419855184371737615151659181815220661446877879847199764825219880625500); // vk.K[29].X
        mul_input[1] = uint256(2191728030944522062213775267825510142676636904535936426097088151735038661017); // vk.K[29].Y
        mul_input[2] = input[28];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[29] * input[28]
        mul_input[0] = uint256(17443744306907421274656073114832682866914815795994710278637727590770342132904); // vk.K[30].X
        mul_input[1] = uint256(6204265850197846880732314988280474321915051365218910504902500465319260176648); // vk.K[30].Y
        mul_input[2] = input[29];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[30] * input[29]
        mul_input[0] = uint256(7667236600173703281656707827902729453577123223272717952708859478183847798002); // vk.K[31].X
        mul_input[1] = uint256(3073364345901477288521870238026227645583520851820532416933060479253244595356); // vk.K[31].Y
        mul_input[2] = input[30];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[31] * input[30]
        mul_input[0] = uint256(9980877541970177898146397507672456369445448128646497326829193893755401659297); // vk.K[32].X
        mul_input[1] = uint256(11845859001496825643147981605740249183632753870257747701403057774143489519069); // vk.K[32].Y
        mul_input[2] = input[31];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[32] * input[31]
        mul_input[0] = uint256(12453897189547283279636360437482740153245209912090247350145743599538029507132); // vk.K[33].X
        mul_input[1] = uint256(6469937287375115226432040539121250021511388797917475330256634615436829876816); // vk.K[33].Y
        mul_input[2] = input[32];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[33] * input[32]

        return
            Pairing.pairing(Pairing.negate(proof.A), proof.B, vk.alfa1, vk.beta2, vk_x, vk.gamma2, proof.C, vk.delta2);
    }
}
