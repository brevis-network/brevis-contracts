contractroot=..
out=$contractroot/out
contractdir=$contractroot/contracts

contract_path_args="$contractdir/interfaces/*.sol $contractdir/light-client-eth/*.sol $contractdir/light-client-eth/common/*.sol $contractdir/light-client-eth/interfaces/*.sol $contractdir/chunk-sync/*.sol $contractdir/verifiers/interfaces/*.sol $contractdir/verifiers/*.sol $contractdir/verifiers/interfaces/*.sol $contractdir/smt/*.sol "

echo "run solc"
solc --overwrite --optimize --via-ir --pretty-json \
  --base-path $contractroot \
  --allow-paths $contractdir \
  --optimize-runs 800 \
  --combined-json abi,bin \
  -o $contractroot/out \
  '@openzeppelin/'=./node_modules/@openzeppelin/ \
  'solidity-rlp'=./node_modules/solidity-rlp/ \
  'solidity-bytes-utils'=./node_modules/solidity-bytes-utils/ \
  $contract_path_args

echo "run abigen"
abigen -combined-json $out/combined.json -pkg eth -out bindings.go
echo "clean up"
rm -rf $contractroot/out
echo "done"