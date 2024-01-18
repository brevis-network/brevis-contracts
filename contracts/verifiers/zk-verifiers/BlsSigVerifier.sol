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

contract BlsSigVerifier {
    using Pairing for *;

    function verifyingKey() internal pure returns (Common.VerifyingKey memory vk) {
        vk.alfa1 = Pairing.G1Point(
            uint256(21869404648590355938070204738007299921184879677994422527706836467860465229555),
            uint256(13498271808119839955057715147407595718888788089303053071109523938531313129416)
        );
        vk.beta2 = Pairing.G2Point(
            [
                uint256(7054962807852101821777521907353900534536574912206623342548135225636684065633),
                uint256(4719416372386397569789716378331929165562736304329438825528404248445356317544)
            ],
            [
                uint256(13169505134780753056527210184700054053183554009975495323937739848223108944491),
                uint256(13592098286878802627104334887812977484641971308142750321766707760557234071693)
            ]
        );
        vk.gamma2 = Pairing.G2Point(
            [
                uint256(11095415866555931179835746321035950972580128700079993342944687987088771893970),
                uint256(7608157602633458693059833022944239778069565713137388978977277542137468568611)
            ],
            [
                uint256(7401180895810745229430756020474788835440387402515164454484661092797156083108),
                uint256(5065358031358114712449190279086624673751222971320486961839316362446988673960)
            ]
        );
        vk.delta2 = Pairing.G2Point(
            [
                uint256(2222475616788908183739316851660307727267648260775064003405822118215303226516),
                uint256(6757963293650080478631547193181808365039329301693745170213066300772412893432)
            ],
            [
                uint256(16109832433313432721291101899523165130162303404627529133543410010343809968099),
                uint256(10657128623625091271138067059727783952095522932162725865057316893805238179881)
            ]
        );
    }

    /*
     * @returns Whether the proof is valid given the hardcoded verifying key
     *          above and the public inputs
     */
    function verifyBlsSigProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[35] memory input
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

        vk_x.X = uint256(6164202379403353093337803285728957014471789019875115975131401894724388184318); // vk.K[0].X
        vk_x.Y = uint256(14653739865386807698202111307501669964169649892515646451656712652300441924746); // vk.K[0].Y
        mul_input[0] = uint256(21624815114078889503955414395662096302081445496963466787344578976263660902728); // vk.K[1].X
        mul_input[1] = uint256(11975795059976038387412140269866198306105499844184502711829412068023274602635); // vk.K[1].Y
        mul_input[2] = input[0];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[1] * input[0]
        mul_input[0] = uint256(9775291361201569299414467699407168277620212957980866255045879806029778744670); // vk.K[2].X
        mul_input[1] = uint256(17233924503171010558175232794027883068335710383496735149565234989213088796768); // vk.K[2].Y
        mul_input[2] = input[1];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[2] * input[1]
        mul_input[0] = uint256(8261354145794168978236113350349424817333204461878829563010649753110739562635); // vk.K[3].X
        mul_input[1] = uint256(228648003771409636961945287326030341918541270047335109757141629333782275554); // vk.K[3].Y
        mul_input[2] = input[2];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[3] * input[2]
        mul_input[0] = uint256(17152002075022784984752660010686481854578129701150390242442680875098241592838); // vk.K[4].X
        mul_input[1] = uint256(11918344010421497075133630718996356084858247140809583870914170975825762239902); // vk.K[4].Y
        mul_input[2] = input[3];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[4] * input[3]
        mul_input[0] = uint256(6724640989773315527339379483030265695332390420189583344806142015699008773038); // vk.K[5].X
        mul_input[1] = uint256(861669679975036023917296038423011939661116903611653803520085082890905199789); // vk.K[5].Y
        mul_input[2] = input[4];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[5] * input[4]
        mul_input[0] = uint256(14251387140848206404864695272174413529185588726241264808479413090259577951538); // vk.K[6].X
        mul_input[1] = uint256(4877123074037160235642133175015361159040197060194606461929469075209090452909); // vk.K[6].Y
        mul_input[2] = input[5];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[6] * input[5]
        mul_input[0] = uint256(20478957346453652200579299273452777262775946284794283970558676720582194409443); // vk.K[7].X
        mul_input[1] = uint256(13820772652701693231224118191632029591356297457861304887694791784009996804692); // vk.K[7].Y
        mul_input[2] = input[6];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[7] * input[6]
        mul_input[0] = uint256(9234465648938974703786224967061256507261169398685469391550057760202443899331); // vk.K[8].X
        mul_input[1] = uint256(13214614408686260510380993851179383149022140264065385584700637692530775154803); // vk.K[8].Y
        mul_input[2] = input[7];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[8] * input[7]
        mul_input[0] = uint256(8322398570000964255587165898643030029017010777972305270331988207302834192671); // vk.K[9].X
        mul_input[1] = uint256(2792163176420440857016104913456882943905817335378546111765666996712066935576); // vk.K[9].Y
        mul_input[2] = input[8];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[9] * input[8]
        mul_input[0] = uint256(15797084070376204854261757656819325712271250278301568410976291761898112633721); // vk.K[10].X
        mul_input[1] = uint256(16921845299282876706614570130523339313885521170603933721890067977426560293948); // vk.K[10].Y
        mul_input[2] = input[9];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[10] * input[9]
        mul_input[0] = uint256(406216972217762516055114890594086577323982284148453070046117814489609787072); // vk.K[11].X
        mul_input[1] = uint256(19680277307775257363705979553150354779381952871529724380294836363819544498123); // vk.K[11].Y
        mul_input[2] = input[10];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[11] * input[10]
        mul_input[0] = uint256(12525306314981123491352345009709438907206215415831226491632499998999204127459); // vk.K[12].X
        mul_input[1] = uint256(1658709794415567495634480941878933577381576154054415708255836274811180964933); // vk.K[12].Y
        mul_input[2] = input[11];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[12] * input[11]
        mul_input[0] = uint256(243194178389731910341565843545971823670153019065524768068260919793967722188); // vk.K[13].X
        mul_input[1] = uint256(1539016065323199244386584217201350236710807235938447628544504644138676080000); // vk.K[13].Y
        mul_input[2] = input[12];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[13] * input[12]
        mul_input[0] = uint256(9275972589095639516678595044389709393960973430061520882565539743422692685960); // vk.K[14].X
        mul_input[1] = uint256(4472610206252234523359317365636135395350815197903460649733786399406367486575); // vk.K[14].Y
        mul_input[2] = input[13];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[14] * input[13]
        mul_input[0] = uint256(1604039017774374075569025682010500744317749653925017699242881946260024023542); // vk.K[15].X
        mul_input[1] = uint256(9410905005395438689346852685696727671986553838639094357993770247718588779274); // vk.K[15].Y
        mul_input[2] = input[14];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[15] * input[14]
        mul_input[0] = uint256(5467079506302237253715880360217400787071012184171130827989740479628417662627); // vk.K[16].X
        mul_input[1] = uint256(5224193576447153852009901997143890474156655585587640006673196987947931700778); // vk.K[16].Y
        mul_input[2] = input[15];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[16] * input[15]
        mul_input[0] = uint256(9387444906146485207075936664808214655500107570142417551032840004938425984101); // vk.K[17].X
        mul_input[1] = uint256(590813342019945519768071606508477036322091281941727992935444194199808517685); // vk.K[17].Y
        mul_input[2] = input[16];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[17] * input[16]
        mul_input[0] = uint256(19041927819469860627785730717901058973852083133163312846046958804599105964228); // vk.K[18].X
        mul_input[1] = uint256(8145703607669957376931460933689362105951308370224326599319773750727102492008); // vk.K[18].Y
        mul_input[2] = input[17];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[18] * input[17]
        mul_input[0] = uint256(14204777538249598371293791017226113964614885045747036692434310643177244084169); // vk.K[19].X
        mul_input[1] = uint256(19991551163786857983848828955926174440512514907619539468704022616567132258153); // vk.K[19].Y
        mul_input[2] = input[18];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[19] * input[18]
        mul_input[0] = uint256(12417075001679265443105210537412885798344552810848770848740273623901433226365); // vk.K[20].X
        mul_input[1] = uint256(15254102101044498161161923790765048216173887701468938517549705046497697348129); // vk.K[20].Y
        mul_input[2] = input[19];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[20] * input[19]
        mul_input[0] = uint256(10269550220001463320881104580554222531390353399943986353550597742767825203938); // vk.K[21].X
        mul_input[1] = uint256(3029115017209685411907558073483483867359830128318261715424531870119465329684); // vk.K[21].Y
        mul_input[2] = input[20];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[21] * input[20]
        mul_input[0] = uint256(14428083165341078925702200873211485650187256256509974692495884984971265925464); // vk.K[22].X
        mul_input[1] = uint256(15475503186436372842471391388091832734881597469954031845230603258731548786526); // vk.K[22].Y
        mul_input[2] = input[21];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[22] * input[21]
        mul_input[0] = uint256(21628050343072376849090277477176470142622908553816436115710816402067062636840); // vk.K[23].X
        mul_input[1] = uint256(7029190031936503335713201259562887395453755842232447142827705567066954030015); // vk.K[23].Y
        mul_input[2] = input[22];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[23] * input[22]
        mul_input[0] = uint256(18110637118732954586253866494394595281955623057396063638731715904832753062595); // vk.K[24].X
        mul_input[1] = uint256(10041565796468001571466611556015291077335046776933681659744217988421358102715); // vk.K[24].Y
        mul_input[2] = input[23];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[24] * input[23]
        mul_input[0] = uint256(4394594477670212487624118222291685582484342594449655582983945008137107825); // vk.K[25].X
        mul_input[1] = uint256(12907667141942256119052629834176444916095324468551648766101314800519980891911); // vk.K[25].Y
        mul_input[2] = input[24];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[25] * input[24]
        mul_input[0] = uint256(12264813020217849970868777479799238229912660701320711442659475763187093184540); // vk.K[26].X
        mul_input[1] = uint256(5418405552541349472839540236004659645077530645792067120765938377537697612686); // vk.K[26].Y
        mul_input[2] = input[25];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[26] * input[25]
        mul_input[0] = uint256(511205244113194326604079314682967281483973984567702129900795395158408587823); // vk.K[27].X
        mul_input[1] = uint256(13182904816261728693468832675109011780365501025377828890958755203977021114190); // vk.K[27].Y
        mul_input[2] = input[26];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[27] * input[26]
        mul_input[0] = uint256(7745322370569289256054483431857334835907602968150216319025615839493340926445); // vk.K[28].X
        mul_input[1] = uint256(802527185108559510753348953543835956246583756974633024055374364959752960622); // vk.K[28].Y
        mul_input[2] = input[27];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[28] * input[27]
        mul_input[0] = uint256(6443616155328954964370112634631799535966126116548305083954877582628003863726); // vk.K[29].X
        mul_input[1] = uint256(6748022417080637760377192763249012207074862674787085243054196391894862853769); // vk.K[29].Y
        mul_input[2] = input[28];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[29] * input[28]
        mul_input[0] = uint256(4746053807226732815609107613694422909828283203346089197433464043542316944158); // vk.K[30].X
        mul_input[1] = uint256(568459565911058268471985303141965010407745186442811181897604717750346101795); // vk.K[30].Y
        mul_input[2] = input[29];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[30] * input[29]
        mul_input[0] = uint256(6867388841298402377650181552003032287359815924707163591172560189355289362811); // vk.K[31].X
        mul_input[1] = uint256(12681153196280456072370427543675941721359193477428355329469521843056250842522); // vk.K[31].Y
        mul_input[2] = input[30];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[31] * input[30]
        mul_input[0] = uint256(11834855045548628237299252472287217659541095251983817826749302275066363050032); // vk.K[32].X
        mul_input[1] = uint256(10549551457031005862882545350575007326729695512101543170018470906439946899475); // vk.K[32].Y
        mul_input[2] = input[31];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[32] * input[31]
        mul_input[0] = uint256(21080571058062105169480371434944542317741872583578405624882037076932528614624); // vk.K[33].X
        mul_input[1] = uint256(11386667982944779644843347272450639356560600194842785817671488724791644458315); // vk.K[33].Y
        mul_input[2] = input[32];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[33] * input[32]
        mul_input[0] = uint256(8323118186735247530367920845955364944904052300793452286806441510422146880773); // vk.K[34].X
        mul_input[1] = uint256(7046818066858629220640333055452104424201007141089175803564675939513429614352); // vk.K[34].Y
        mul_input[2] = input[33];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[34] * input[33]
        mul_input[0] = uint256(12620149755261046671755404694759801128385638190500539207079026239458336064522); // vk.K[35].X
        mul_input[1] = uint256(10994846953135033560648325534438412658436366579529892185303274351838980685981); // vk.K[35].Y
        mul_input[2] = input[34];
        Common.accumulate(mul_input, q, add_input, vk_x); // vk_x += vk.K[35] * input[34]
        if (commit[0] != 0 || commit[1] != 0) {
            vk_x = Pairing.plus(vk_x, proof.Commit);
        }

        return
            Pairing.pairing(Pairing.negate(proof.A), proof.B, vk.alfa1, vk.beta2, vk_x, vk.gamma2, proof.C, vk.delta2);
    }
}
