// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AnchorBlocksBlockHashWitness is an auto generated low-level Go binding around an user-defined struct.
type AnchorBlocksBlockHashWitness struct {
	Left  []byte
	Right []byte
}

// BeaconBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type BeaconBlockHeader struct {
	Slot          uint64
	ProposerIndex uint64
	ParentRoot    [32]byte
	StateRoot     [32]byte
	BodyRoot      [32]byte
}

// ExecutionPayload is an auto generated low-level Go binding around an user-defined struct.
type ExecutionPayload struct {
	StateRoot   LeafWithBranch
	BlockHash   LeafWithBranch
	BlockNumber LeafWithBranch
}

// HeaderWithExecution is an auto generated low-level Go binding around an user-defined struct.
type HeaderWithExecution struct {
	Beacon        BeaconBlockHeader
	Execution     ExecutionPayload
	ExecutionRoot LeafWithBranch
}

// IBeaconVerifierProof is an auto generated low-level Go binding around an user-defined struct.
type IBeaconVerifierProof struct {
	A          [2]*big.Int
	B          [2][2]*big.Int
	C          [2]*big.Int
	Commitment [2]*big.Int
}

// IBlockChunksBlockHashWitness is an auto generated low-level Go binding around an user-defined struct.
type IBlockChunksBlockHashWitness struct {
	ChainId        uint64
	BlkNum         uint32
	ClaimedBlkHash [32]byte
	PrevHash       [32]byte
	NumFinal       uint32
	MerkleProof    [7][32]byte
}

// IReceiptVerifierLogInfo is an auto generated low-level Go binding around an user-defined struct.
type IReceiptVerifierLogInfo struct {
	Addr   common.Address
	Topics [][32]byte
	Data   []byte
}

// IReceiptVerifierReceiptInfo is an auto generated low-level Go binding around an user-defined struct.
type IReceiptVerifierReceiptInfo struct {
	Success bool
	ChainId uint64
	BlkHash [32]byte
	BlkNum  uint32
	BlkTime uint64
	Logs    []IReceiptVerifierLogInfo
}

// ISMTSmtUpdate is an auto generated low-level Go binding around an user-defined struct.
type ISMTSmtUpdate struct {
	NewSmtRoot          [32]byte
	EndBlockNum         uint64
	EndBlockHash        [32]byte
	NextChunkMerkleRoot [32]byte
	Proof               [8]*big.Int
	Commit              [2]*big.Int
	KnowledgeProof      [2]*big.Int
}

// ISlotValueVerifierSlotInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlotValueVerifierSlotInfo struct {
	ChainId     uint64
	AddrHash    [32]byte
	BlkHash     [32]byte
	SlotKeyHash [32]byte
	SlotValue   [32]byte
	BlkNum      uint32
}

// ITxVerifierTxInfo is an auto generated low-level Go binding around an user-defined struct.
type ITxVerifierTxInfo struct {
	ChainId   uint64
	Nonce     uint64
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Gas       *big.Int
	To        common.Address
	Value     *big.Int
	Data      []byte
	From      common.Address
	BlkNum    uint32
	BlkHash   [32]byte
	BlkTime   uint64
}

// LeafWithBranch is an auto generated low-level Go binding around an user-defined struct.
type LeafWithBranch struct {
	Leaf   [32]byte
	Branch [][32]byte
}

// LightClientOptimisticUpdate is an auto generated low-level Go binding around an user-defined struct.
type LightClientOptimisticUpdate struct {
	AttestedHeader HeaderWithExecution
	SyncAggregate  SyncAggregate
	SignatureSlot  uint64
}

// LightClientUpdate is an auto generated low-level Go binding around an user-defined struct.
type LightClientUpdate struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	FinalityBranch                    [][32]byte
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteeBranch           [][32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}

// SyncAggregate is an auto generated low-level Go binding around an user-defined struct.
type SyncAggregate struct {
	Participation uint64
	PoseidonRoot  [32]byte
	Commitment    *big.Int
	Proof         IBeaconVerifierProof
}

// AnchorBlocksMetaData contains all meta data concerning the AnchorBlocks contract.
var AnchorBlocksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lightClient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"AnchorBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lightClient\",\"outputs\":[{\"internalType\":\"contractIEthereumLightClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"attestedHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"}],\"internalType\":\"structLightClientOptimisticUpdate\",\"name\":\"hb\",\"type\":\"tuple\"}],\"name\":\"processUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"attestedHeader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"}],\"internalType\":\"structLightClientOptimisticUpdate\",\"name\":\"hb\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"left\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"right\",\"type\":\"bytes\"}],\"internalType\":\"structAnchorBlocks.BlockHashWitness[]\",\"name\":\"chainProof\",\"type\":\"tuple[]\"}],\"name\":\"processUpdateWithChainProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lightClient\",\"type\":\"address\"}],\"name\":\"setLightClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346100a457601f61119d38819003918201601f19168301916001600160401b038311848410176100a8578084926020946040528339810103126100a457516001600160a01b0390818116908190036100a4575f5460018060a01b03199033828216175f55604051933391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a360015416176001556110e090816100bd8239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60406080815260049081361015610014575f80fd5b5f91823560e01c9081632e3466f6146105725781633e553bab1461051d578163715018a6146104b05781638da5cb5b1461048a578163b5700e6814610462578163f25b3f991461043a578163f2fde38b14610344578163fb8cfb51146100a4575063fbbea34c14610083575f80fd5b346100a057816003193601126100a0576020906003549051908152f35b5080fd5b8383346100a05760603660031901126100a05767ffffffffffffffff8335818111610340576100d69036908601610796565b9160249485359560449081359585871161033c573660238801121561033c578684013596610103886106a4565b976101108751998a610669565b80895260209784898b019260051b8401019236841161033857858101925b8484106102c35750505050508651156102835761014a90610a37565b97909887518a03998a11610272578099829a5b89518c101561021a57866101718d8c610fef565b515151036101d95789896101ca6101d3938f8361019d826101956101bc9489610fef565b515197610fef565b5101518d519485936101b28286018099611017565b9081520190611017565b03601f198101835282610669565b5190209b610fcd565b9a61015d565b875162461bcd60e51b81528088018a90526010818701527f696e76616c6964206c656674206c656e0000000000000000000000000000000081880152606490fd5b85858a8a8e8b950361023457876102318888610f24565b80f35b5162461bcd60e51b8152938401526012908301527f696e76616c696420636861696e50726f6f66000000000000000000000000000090820152606490fd5b50634e487b7160e01b815260118452fd5b506014906064957f696e76616c69642070726f6f66206c656e6774680000000000000000000000009495519562461bcd60e51b8752860152840152820152fd5b83358381116103345782018a6023198236030112610334578a51906102e782610601565b8881013585811161033057610301908a36918401016109ba565b82528981013585811161033057916103218e94928b8695369201016109ba565b8382015281520193019261012e565b8f80fd5b8d80fd5b8b80fd5b8780fd5b8380fd5b91905034610436576020366003190112610436578135916001600160a01b03908184168094036104325784549182169261037f338514611040565b84156103c957505073ffffffffffffffffffffffffffffffffffffffff1916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b8480fd5b8280fd5b9050346104365760203660031901126104365760209282913581526002845220549051908152f35b5050346100a057816003193601126100a0576020906001600160a01b03600154169051908152f35b5050346100a057816003193601126100a0576001600160a01b0360209254169051908152f35b833461051a578060031936011261051a5780805473ffffffffffffffffffffffffffffffffffffffff196001600160a01b038216916104f0338414611040565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b8390346100a05760203660031901126100a057356001600160a01b0380821680920361043657610551908354163314611040565b73ffffffffffffffffffffffffffffffffffffffff19600154161760015580f35b839150346100a05760203660031901126100a057803567ffffffffffffffff8111610436576105a76105ac9136908401610796565b610a37565b9182156105be57509061023191610f24565b606490602086519162461bcd60e51b8352820152600f60248201527f656d70747920626c6f636b4861736800000000000000000000000000000000006044820152fd5b6040810190811067ffffffffffffffff82111761061d57604052565b634e487b7160e01b5f52604160045260245ffd5b6060810190811067ffffffffffffffff82111761061d57604052565b6080810190811067ffffffffffffffff82111761061d57604052565b90601f8019910116810190811067ffffffffffffffff82111761061d57604052565b359067ffffffffffffffff821682036106a057565b5f80fd5b67ffffffffffffffff811161061d5760051b60200190565b91906040838203126106a057604051926106d584610601565b8381358152602091828101359067ffffffffffffffff82116106a057019280601f850112156106a0578335610709816106a4565b946107176040519687610669565b818652848087019260051b8201019283116106a05784809101915b83831061074157505050500152565b8235815291810191859101610732565b9080601f830112156106a0576040519161076a83610601565b8290604081019283116106a057905b8282106107865750505090565b8135815260209182019101610779565b9190828103926101e084126106a0576040908151926107b484610631565b839582359067ffffffffffffffff918281116106a05784019182840360e081126106a05760a08751916107e683610631565b126106a05786519160a083018381108282111761061d5788526108088561068b565b835260209461081886820161068b565b868501528881013589850152606093848201358582015260808201356080820152835260a08101358281116106a057810184818903126106a05789519061085e82610631565b80358481116106a057896108739183016106bc565b8252878101358481116106a0578961088c9183016106bc565b888301528a810135908482116106a0576108a8918a91016106bc565b8a8201528684015260c08101359182116106a0576108c8918791016106bc565b8188015287526101a0601f198301126106a0578551936108e78561064d565b6108f284870161068b565b8552868601358486015281860135878601526101408093607f1901126106a05786519261091e8461064d565b61092b8260808901610751565b84528160df880112156106a057875161094381610601565b80918801918383116106a0578960c08a01915b8483106109a1575050508261097e869488969461098c9461099c9c9a6101c09c9a0152610751565b8a8501526101808701610751565b818301528301528601520161068b565b910152565b88906109ad8785610751565b8152019101908a90610956565b81601f820112156106a05780359067ffffffffffffffff821161061d57604051926109ef601f8401601f191660200185610669565b828452602083830101116106a057815f926020809301838601378301015290565b5f915b60028310610a2057505050565b600190825181526020809101920192019190610a13565b602091828201928351925f94600367ffffffffffffffff8096511602858116908103610f105761040011610ecb57815191604094858401908151936080865101519887865196015195968b975b8751891015610b365760ff8911610b22576001808a1b15610b0e578a918e9160198c1c811603610af657610ac5610ad091610abf8d8d610fef565b5161108b565b8d5191828092611017565b039060025afa15610aec57610ae68c5198610fcd565b97610a84565b89513d8d823e3d90fd5b610ac5610ad091610b078d8d610fef565b519061108b565b634e487b7160e01b8e52601260045260248efd5b634e487b7160e01b8d52601160045260248dfd5b92959a9194975092955003610e87578501938685510151938151519887865196015195968b975b8751891015610bc55760ff8911610b22576001808a1b15610b0e578a918e9160168c1c811603610bb457610ac5610b9891610abf8d8d610fef565b039060025afa15610aec57610bae8c5198610fcd565b97610b5d565b610ac5610b9891610b078d8d610fef565b92959a9194975092955003610e5757859051015192515196858451940151939489955b8551871015610c8d5760ff8711610c7957600180881b15610c655788918c91601c8a1c811603610c5457610c23610c2e91610abf8b8b610fef565b8b5191828092611017565b039060025afa15610c4a57610c448a5196610fcd565b95610be8565b87513d8b823e3d90fd5b610c23610c2e91610b078b8b610fef565b634e487b7160e01b8c52601260045260248cfd5b634e487b7160e01b8b52601160045260248bfd5b929550929790935095949503610e275760016001600160a01b0381541690828686015116938551519051833b15610e23579160608b9492969360808a51988996631ee9fa5560e11b88526004880152888151166024880152888d8201511660448801528b8101516064880152838101516084880152015160a48601528681511660c48601528a81015160e486015289810151610104860152015190610d3761012485018351610a10565b818a0151858b8b61016488015b60028410610df85750610264969450610d7f935060609250859150610d7290889601516101e4870190610a10565b0151610224840190610a10565b5afa8015610dee57610dc7575b50508290510190815101515184905b838210610dab5750505101515190565b600895861b60ff821617951c90610dc190610fcd565b90610d9b565b8196929611610dda57825293825f610d8c565b634e487b7160e01b82526041600452602482fd5b84513d89823e3d90fd5b9091929496985083959750610e0f81889551610a10565b0193019101908c9593918b8b8a9795610d44565b8a80fd5b835162461bcd60e51b81526004810186905260096024820152683130b210383937b7b360b91b6044820152606490fd5b865162461bcd60e51b81526004810187905260096024820152683130b210383937b7b360b91b6044820152606490fd5b865162461bcd60e51b815260048101879052601360248201527f626164206578656320726f6f742070726f6f66000000000000000000000000006044820152606490fd5b60405162461bcd60e51b815260048101849052601260248201527f71756f72756d206e6f74207265616368656400000000000000000000000000006044820152606490fd5b634e487b7160e01b87526011600452602487fd5b90815f52600260205260405f2054610f8857816040917fa9aaf84657c346a7eafe57cf0403ab0be7867b79a5fdd5e5ab3527fbfc739d85935f52600260205280835f20556003548211610f7f575b82519182526020820152a1565b81600355610f72565b60405162461bcd60e51b815260206004820152601960248201527f626c6f636b206861736820616c726561647920657869737473000000000000006044820152606490fd5b5f198114610fdb5760010190565b634e487b7160e01b5f52601160045260245ffd5b80518210156110035760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b908151915f5b83811061102d575050015f815290565b806020809284010151818501520161101d565b1561104757565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b91906040519260208401526040830152604082526110a882610631565b56fea26469706673582212202c2f7b3a1d0378238ca18fbd5d8ccd8bd2bd19b4186952e9562f7510551eb33864736f6c63430008140033",
}

// AnchorBlocksABI is the input ABI used to generate the binding from.
// Deprecated: Use AnchorBlocksMetaData.ABI instead.
var AnchorBlocksABI = AnchorBlocksMetaData.ABI

// AnchorBlocksBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AnchorBlocksMetaData.Bin instead.
var AnchorBlocksBin = AnchorBlocksMetaData.Bin

// DeployAnchorBlocks deploys a new Ethereum contract, binding an instance of AnchorBlocks to it.
func DeployAnchorBlocks(auth *bind.TransactOpts, backend bind.ContractBackend, _lightClient common.Address) (common.Address, *types.Transaction, *AnchorBlocks, error) {
	parsed, err := AnchorBlocksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AnchorBlocksBin), backend, _lightClient)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AnchorBlocks{AnchorBlocksCaller: AnchorBlocksCaller{contract: contract}, AnchorBlocksTransactor: AnchorBlocksTransactor{contract: contract}, AnchorBlocksFilterer: AnchorBlocksFilterer{contract: contract}}, nil
}

// AnchorBlocks is an auto generated Go binding around an Ethereum contract.
type AnchorBlocks struct {
	AnchorBlocksCaller     // Read-only binding to the contract
	AnchorBlocksTransactor // Write-only binding to the contract
	AnchorBlocksFilterer   // Log filterer for contract events
}

// AnchorBlocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnchorBlocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorBlocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnchorBlocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorBlocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnchorBlocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorBlocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnchorBlocksSession struct {
	Contract     *AnchorBlocks     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnchorBlocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnchorBlocksCallerSession struct {
	Contract *AnchorBlocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AnchorBlocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnchorBlocksTransactorSession struct {
	Contract     *AnchorBlocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AnchorBlocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnchorBlocksRaw struct {
	Contract *AnchorBlocks // Generic contract binding to access the raw methods on
}

// AnchorBlocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnchorBlocksCallerRaw struct {
	Contract *AnchorBlocksCaller // Generic read-only contract binding to access the raw methods on
}

// AnchorBlocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnchorBlocksTransactorRaw struct {
	Contract *AnchorBlocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnchorBlocks creates a new instance of AnchorBlocks, bound to a specific deployed contract.
func NewAnchorBlocks(address common.Address, backend bind.ContractBackend) (*AnchorBlocks, error) {
	contract, err := bindAnchorBlocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnchorBlocks{AnchorBlocksCaller: AnchorBlocksCaller{contract: contract}, AnchorBlocksTransactor: AnchorBlocksTransactor{contract: contract}, AnchorBlocksFilterer: AnchorBlocksFilterer{contract: contract}}, nil
}

// NewAnchorBlocksCaller creates a new read-only instance of AnchorBlocks, bound to a specific deployed contract.
func NewAnchorBlocksCaller(address common.Address, caller bind.ContractCaller) (*AnchorBlocksCaller, error) {
	contract, err := bindAnchorBlocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorBlocksCaller{contract: contract}, nil
}

// NewAnchorBlocksTransactor creates a new write-only instance of AnchorBlocks, bound to a specific deployed contract.
func NewAnchorBlocksTransactor(address common.Address, transactor bind.ContractTransactor) (*AnchorBlocksTransactor, error) {
	contract, err := bindAnchorBlocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorBlocksTransactor{contract: contract}, nil
}

// NewAnchorBlocksFilterer creates a new log filterer instance of AnchorBlocks, bound to a specific deployed contract.
func NewAnchorBlocksFilterer(address common.Address, filterer bind.ContractFilterer) (*AnchorBlocksFilterer, error) {
	contract, err := bindAnchorBlocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnchorBlocksFilterer{contract: contract}, nil
}

// bindAnchorBlocks binds a generic wrapper to an already deployed contract.
func bindAnchorBlocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnchorBlocksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnchorBlocks *AnchorBlocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnchorBlocks.Contract.AnchorBlocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnchorBlocks *AnchorBlocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.AnchorBlocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnchorBlocks *AnchorBlocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.AnchorBlocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnchorBlocks *AnchorBlocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnchorBlocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnchorBlocks *AnchorBlocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnchorBlocks *AnchorBlocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(bytes32)
func (_AnchorBlocks *AnchorBlocksCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AnchorBlocks.contract.Call(opts, &out, "blocks", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(bytes32)
func (_AnchorBlocks *AnchorBlocksSession) Blocks(arg0 *big.Int) ([32]byte, error) {
	return _AnchorBlocks.Contract.Blocks(&_AnchorBlocks.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(bytes32)
func (_AnchorBlocks *AnchorBlocksCallerSession) Blocks(arg0 *big.Int) ([32]byte, error) {
	return _AnchorBlocks.Contract.Blocks(&_AnchorBlocks.CallOpts, arg0)
}

// LatestBlockNum is a free data retrieval call binding the contract method 0xfbbea34c.
//
// Solidity: function latestBlockNum() view returns(uint256)
func (_AnchorBlocks *AnchorBlocksCaller) LatestBlockNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnchorBlocks.contract.Call(opts, &out, "latestBlockNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNum is a free data retrieval call binding the contract method 0xfbbea34c.
//
// Solidity: function latestBlockNum() view returns(uint256)
func (_AnchorBlocks *AnchorBlocksSession) LatestBlockNum() (*big.Int, error) {
	return _AnchorBlocks.Contract.LatestBlockNum(&_AnchorBlocks.CallOpts)
}

// LatestBlockNum is a free data retrieval call binding the contract method 0xfbbea34c.
//
// Solidity: function latestBlockNum() view returns(uint256)
func (_AnchorBlocks *AnchorBlocksCallerSession) LatestBlockNum() (*big.Int, error) {
	return _AnchorBlocks.Contract.LatestBlockNum(&_AnchorBlocks.CallOpts)
}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_AnchorBlocks *AnchorBlocksCaller) LightClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnchorBlocks.contract.Call(opts, &out, "lightClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_AnchorBlocks *AnchorBlocksSession) LightClient() (common.Address, error) {
	return _AnchorBlocks.Contract.LightClient(&_AnchorBlocks.CallOpts)
}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_AnchorBlocks *AnchorBlocksCallerSession) LightClient() (common.Address, error) {
	return _AnchorBlocks.Contract.LightClient(&_AnchorBlocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnchorBlocks *AnchorBlocksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnchorBlocks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnchorBlocks *AnchorBlocksSession) Owner() (common.Address, error) {
	return _AnchorBlocks.Contract.Owner(&_AnchorBlocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnchorBlocks *AnchorBlocksCallerSession) Owner() (common.Address, error) {
	return _AnchorBlocks.Contract.Owner(&_AnchorBlocks.CallOpts)
}

// ProcessUpdate is a paid mutator transaction binding the contract method 0x2e3466f6.
//
// Solidity: function processUpdate((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) hb) returns()
func (_AnchorBlocks *AnchorBlocksTransactor) ProcessUpdate(opts *bind.TransactOpts, hb LightClientOptimisticUpdate) (*types.Transaction, error) {
	return _AnchorBlocks.contract.Transact(opts, "processUpdate", hb)
}

// ProcessUpdate is a paid mutator transaction binding the contract method 0x2e3466f6.
//
// Solidity: function processUpdate((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) hb) returns()
func (_AnchorBlocks *AnchorBlocksSession) ProcessUpdate(hb LightClientOptimisticUpdate) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.ProcessUpdate(&_AnchorBlocks.TransactOpts, hb)
}

// ProcessUpdate is a paid mutator transaction binding the contract method 0x2e3466f6.
//
// Solidity: function processUpdate((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) hb) returns()
func (_AnchorBlocks *AnchorBlocksTransactorSession) ProcessUpdate(hb LightClientOptimisticUpdate) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.ProcessUpdate(&_AnchorBlocks.TransactOpts, hb)
}

// ProcessUpdateWithChainProof is a paid mutator transaction binding the contract method 0xfb8cfb51.
//
// Solidity: function processUpdateWithChainProof((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) hb, bytes32 blockHash, (bytes,bytes)[] chainProof) returns()
func (_AnchorBlocks *AnchorBlocksTransactor) ProcessUpdateWithChainProof(opts *bind.TransactOpts, hb LightClientOptimisticUpdate, blockHash [32]byte, chainProof []AnchorBlocksBlockHashWitness) (*types.Transaction, error) {
	return _AnchorBlocks.contract.Transact(opts, "processUpdateWithChainProof", hb, blockHash, chainProof)
}

// ProcessUpdateWithChainProof is a paid mutator transaction binding the contract method 0xfb8cfb51.
//
// Solidity: function processUpdateWithChainProof((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) hb, bytes32 blockHash, (bytes,bytes)[] chainProof) returns()
func (_AnchorBlocks *AnchorBlocksSession) ProcessUpdateWithChainProof(hb LightClientOptimisticUpdate, blockHash [32]byte, chainProof []AnchorBlocksBlockHashWitness) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.ProcessUpdateWithChainProof(&_AnchorBlocks.TransactOpts, hb, blockHash, chainProof)
}

// ProcessUpdateWithChainProof is a paid mutator transaction binding the contract method 0xfb8cfb51.
//
// Solidity: function processUpdateWithChainProof((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) hb, bytes32 blockHash, (bytes,bytes)[] chainProof) returns()
func (_AnchorBlocks *AnchorBlocksTransactorSession) ProcessUpdateWithChainProof(hb LightClientOptimisticUpdate, blockHash [32]byte, chainProof []AnchorBlocksBlockHashWitness) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.ProcessUpdateWithChainProof(&_AnchorBlocks.TransactOpts, hb, blockHash, chainProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnchorBlocks *AnchorBlocksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorBlocks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnchorBlocks *AnchorBlocksSession) RenounceOwnership() (*types.Transaction, error) {
	return _AnchorBlocks.Contract.RenounceOwnership(&_AnchorBlocks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnchorBlocks *AnchorBlocksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AnchorBlocks.Contract.RenounceOwnership(&_AnchorBlocks.TransactOpts)
}

// SetLightClient is a paid mutator transaction binding the contract method 0x3e553bab.
//
// Solidity: function setLightClient(address _lightClient) returns()
func (_AnchorBlocks *AnchorBlocksTransactor) SetLightClient(opts *bind.TransactOpts, _lightClient common.Address) (*types.Transaction, error) {
	return _AnchorBlocks.contract.Transact(opts, "setLightClient", _lightClient)
}

// SetLightClient is a paid mutator transaction binding the contract method 0x3e553bab.
//
// Solidity: function setLightClient(address _lightClient) returns()
func (_AnchorBlocks *AnchorBlocksSession) SetLightClient(_lightClient common.Address) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.SetLightClient(&_AnchorBlocks.TransactOpts, _lightClient)
}

// SetLightClient is a paid mutator transaction binding the contract method 0x3e553bab.
//
// Solidity: function setLightClient(address _lightClient) returns()
func (_AnchorBlocks *AnchorBlocksTransactorSession) SetLightClient(_lightClient common.Address) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.SetLightClient(&_AnchorBlocks.TransactOpts, _lightClient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnchorBlocks *AnchorBlocksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AnchorBlocks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnchorBlocks *AnchorBlocksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.TransferOwnership(&_AnchorBlocks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnchorBlocks *AnchorBlocksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AnchorBlocks.Contract.TransferOwnership(&_AnchorBlocks.TransactOpts, newOwner)
}

// AnchorBlocksAnchorBlockUpdatedIterator is returned from FilterAnchorBlockUpdated and is used to iterate over the raw logs and unpacked data for AnchorBlockUpdated events raised by the AnchorBlocks contract.
type AnchorBlocksAnchorBlockUpdatedIterator struct {
	Event *AnchorBlocksAnchorBlockUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnchorBlocksAnchorBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnchorBlocksAnchorBlockUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnchorBlocksAnchorBlockUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnchorBlocksAnchorBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnchorBlocksAnchorBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnchorBlocksAnchorBlockUpdated represents a AnchorBlockUpdated event raised by the AnchorBlocks contract.
type AnchorBlocksAnchorBlockUpdated struct {
	BlockNum  *big.Int
	BlockHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnchorBlockUpdated is a free log retrieval operation binding the contract event 0xa9aaf84657c346a7eafe57cf0403ab0be7867b79a5fdd5e5ab3527fbfc739d85.
//
// Solidity: event AnchorBlockUpdated(uint256 blockNum, bytes32 blockHash)
func (_AnchorBlocks *AnchorBlocksFilterer) FilterAnchorBlockUpdated(opts *bind.FilterOpts) (*AnchorBlocksAnchorBlockUpdatedIterator, error) {

	logs, sub, err := _AnchorBlocks.contract.FilterLogs(opts, "AnchorBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &AnchorBlocksAnchorBlockUpdatedIterator{contract: _AnchorBlocks.contract, event: "AnchorBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchAnchorBlockUpdated is a free log subscription operation binding the contract event 0xa9aaf84657c346a7eafe57cf0403ab0be7867b79a5fdd5e5ab3527fbfc739d85.
//
// Solidity: event AnchorBlockUpdated(uint256 blockNum, bytes32 blockHash)
func (_AnchorBlocks *AnchorBlocksFilterer) WatchAnchorBlockUpdated(opts *bind.WatchOpts, sink chan<- *AnchorBlocksAnchorBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _AnchorBlocks.contract.WatchLogs(opts, "AnchorBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnchorBlocksAnchorBlockUpdated)
				if err := _AnchorBlocks.contract.UnpackLog(event, "AnchorBlockUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnchorBlockUpdated is a log parse operation binding the contract event 0xa9aaf84657c346a7eafe57cf0403ab0be7867b79a5fdd5e5ab3527fbfc739d85.
//
// Solidity: event AnchorBlockUpdated(uint256 blockNum, bytes32 blockHash)
func (_AnchorBlocks *AnchorBlocksFilterer) ParseAnchorBlockUpdated(log types.Log) (*AnchorBlocksAnchorBlockUpdated, error) {
	event := new(AnchorBlocksAnchorBlockUpdated)
	if err := _AnchorBlocks.contract.UnpackLog(event, "AnchorBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnchorBlocksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AnchorBlocks contract.
type AnchorBlocksOwnershipTransferredIterator struct {
	Event *AnchorBlocksOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnchorBlocksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnchorBlocksOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnchorBlocksOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnchorBlocksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnchorBlocksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnchorBlocksOwnershipTransferred represents a OwnershipTransferred event raised by the AnchorBlocks contract.
type AnchorBlocksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnchorBlocks *AnchorBlocksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AnchorBlocksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AnchorBlocks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AnchorBlocksOwnershipTransferredIterator{contract: _AnchorBlocks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnchorBlocks *AnchorBlocksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AnchorBlocksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AnchorBlocks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnchorBlocksOwnershipTransferred)
				if err := _AnchorBlocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnchorBlocks *AnchorBlocksFilterer) ParseOwnershipTransferred(log types.Log) (*AnchorBlocksOwnershipTransferred, error) {
	event := new(AnchorBlocksOwnershipTransferred)
	if err := _AnchorBlocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconVerifierMetaData contains all meta data concerning the BeaconVerifier contract.
var BeaconVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[35]\",\"name\":\"input\",\"type\":\"uint256[35]\"}],\"name\":\"verifyBlsSigProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[33]\",\"name\":\"input\",\"type\":\"uint256[33]\"}],\"name\":\"verifyCommitteeRootMappingProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signingRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteePoseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"participation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"verifySignatureProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sszRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"verifySyncCommitteeRootMappingProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657613067908161001b8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c9081630cc26769146100575750806352356da0146100525780637a5387811461004d5763ab00dde614610048575f80fd5b610662565b6105dd565b61052b565b346100bc576105a03660031901126100bc5761007236610236565b9061007c36610331565b61008536610273565b9061008f366102b1565b923661016312156100bc5760206100b2868686866100ac3661015a565b93610ad1565b6040519015158152f35b80fd5b634e487b7160e01b5f52604160045260245ffd5b604051906080820182811067ffffffffffffffff8211176100f357604052565b6100bf565b604051906040820182811067ffffffffffffffff8211176100f357604052565b60405190610460820182811067ffffffffffffffff8211176100f357604052565b60405190610420820182811067ffffffffffffffff8211176100f357604052565b90610163610118565b91826105a491821161019057610144905b82821061018057505050565b8135815260209182019101610174565b5f80fd5b9061019d6100f8565b918261018491821161019057610144905b8282106101ba57505050565b81358152602091820191016101ae565b906101d36100f8565b91826101c491821161019057610184905b8282106101f057505050565b81358152602091820191016101e4565b906102096100f8565b918261014491821161019057610104905b82821061022657505050565b813581526020918201910161021a565b8060231215610190576102476100f8565b90816044918211610190576004905b8282106102635750505090565b8135815260209182019101610256565b8060e31215610190576102846100f8565b90816101049182116101905760c4905b8282106102a15750505090565b8135815260209182019101610294565b806101231215610190576102c36100f8565b908161014491821161019057610104905b8282106102e15750505090565b81358152602091820191016102d4565b9080601f83011215610190576103056100f8565b80926040810192831161019057905b8282106103215750505090565b8135815260209182019101610314565b8060631215610190576103426100f8565b908160c491808311610190576044915b838310610360575050505090565b602060409161036f84866102f1565b815201920191610352565b8060e312156101905761038b6100f8565b9081610144918083116101905760c4915b8383106103aa575050505090565b60206040916103b984866102f1565b81520192019161039c565b8060a31215610190576103d56100f8565b908161010491808311610190576084915b8383106103f4575050505090565b602060409161040384866102f1565b8152019201916103e6565b90610140608319830112610190576104246100d3565b918060a31215610190576104366100f8565b60c481838211610190576084905b8282106104905750505083526104598161037a565b60208401528061016312156101905761047181610194565b6040840152806101a3121561019057610489906101ca565b6060830152565b8135815260209182019101610444565b90610140604319830112610190576104b66100d3565b918060631215610190576104c86100f8565b608481838211610190576044905b82821061051b5750505083526104eb816103c4565b60208401528061012312156101905761050381610200565b60408401528061016312156101905761048990610194565b81358152602091820191016104d6565b34610190576101c0366003190112610190576105463661040e565b61054e610118565b906104603683375f6004355b602082106105b3576105af61059d858560443561040083015260243561042083015260643561044083015280519060208101516060604083015192015192610ad1565b60405190151581529081906020820190565b0390f35b8060ff6105d792166105cd6105c785610757565b87610779565b5260081c91610749565b9061055a565b3461019057610180366003190112610190576105f8366104a0565b610600610139565b906104203683375f6004355b60208210610638576105af61059d85856024356104008301528051906040602082015191015191611c65565b601f9082820391821161065d57610657916105cd60ff8316918761078f565b9061060c565b610735565b34610190576105203660031901126101905761067d36610236565b61068636610331565b9061069036610273565b91366101231215610190576106a3610139565b9283916105249336851161019057602095610104905b8682106106cc5750506100b29450611c65565b813581529087019087016106b9565b6106e36100d3565b906080368337565b604051906060820182811067ffffffffffffffff8211176100f3576040526060368337565b604051906020820182811067ffffffffffffffff8211176100f3576040526020368337565b634e487b7160e01b5f52601160045260245ffd5b5f19811461065d5760010190565b601f0390601f821161065d57565b634e487b7160e01b5f52603260045260245ffd5b90602381101561078a5760051b0190565b610765565b90602181101561078a5760051b0190565b6107a86100f8565b906107b16100f8565b604036823782526107c06100f8565b60403682376020830152565b6107d46100d3565b906107dd6100f8565b5f9081815281602082015283526107f26107a0565b60208401526107ff6100f8565b81815281602082015260408401526108156100f8565b9080825260208201526060830152565b1561082c57565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d710000000000000000006044820152606490fd5b1561087857565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d710000000000000000006044820152606490fd5b156108c457565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561091057565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561095c57565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d7100000000000000006044820152606490fd5b156109a857565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d7100000000000000006044820152606490fd5b156109f457565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d710000000000000000006044820152606490fd5b15610a4057565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d710000000000000000006044820152606490fd5b15610a8c57565b60405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c64006044820152606490fd5b94939492919092610ae06107cc565b8151602080930151610af06100f8565b918252838201528152610b016100f8565b948051518652610b12815160200190565b5183870152610b3183610b236100f8565b920180515183525160200190565b5183820152610b3e6100f8565b95865282860152818101948552818351930151610b596100f8565b9384528284015260408101928352835193828101948551610b786100f8565b9182528482015260608301908152610c2c7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610bb78186515110610825565b610bc681878751015110610871565b610bd4818a515151106108bd565b610be481878b5101515110610909565b610bfb81610bf48b515160200190565b5110610955565b610c1481610c0d888c51015160200190565b51106109a1565b610c2181885151106109ed565b858751015110610a39565b5f5b602381106119685750906119419798939291610c486119e9565b96610c516100f8565b925f845261190884878101985f8a528199610c6a6106db565b918a610440610c776106eb565b927f2065b90c648581703a4ef82833653ae713aaf62c2dc4ef26b0a9bbbdf254b48a610ca16100f8565b955f87525f858801527f0da0d09dcc32c2d20c9905307190ffb91538db928804c70e7ed77639f2cee0fe8952527f2fcf362c494439bcae24ab0ab7dd0bd40825ed18725c1d11d25eeb863f24194884527f184edce371c121d112278a4d1239f9d65421fb00e688d7612320bb5f66e7409d8385017f1a7a0f4ef55687795fa98c4585fc66e26ddea1f6e161a837ef4a4f1ae9c8808b815282519460408701958652610d4e8a8a8a8a612c69565b7f159c9e6c6ad47c3114dd4bdc88dc34359cae49de8306c8f56c9ba9d56429755e87527f261a0e3bad2b8c7b4057a7708b68ddaa9684eaa9d458088e7a2fe7813e08d06082528301518552610da589898989612c69565b7f1243c2c01b1a238bd0937eed8a5eb5f962eae681000356540ae1cfb36e1e308b86527e816901d1be816971e5f7e84b32a92c58a9f8971ff921adc37884d47db225e2815260408301518552610dfd89898989612c69565b7f25ebb06beeca9f8b08c511a86423c8f8705f115fe942220b9f8e30d54b016e0686527f1a598b2d183a54a959959d562879ae4b48df2203151d223973543b7d9584c59e815260608301518552610e5689898989612c69565b7f0ede030d344e453627bd0d2e849cd89947ed04b1825b0d4f7d6a8bf8b6ca5bae86527f01e7b010c4ab8cfc4791d0886bb39e1e8785b51a2a2165514d1fe026b8de24ad815260808301518552610eaf89898989612c69565b7f1f81ffd062f9644e01e392d39b8de8e030afd731c770e58f96050ed1d36d553286527f0ac85a9509ed9a6e21c1e933b10794303a4b77d9fd1fb036e966fef320cb4dad815260a08301518552610f0889898989612c69565b7f2d46add97dc6a65ee2f3dd3ea61060bf9fd510929a701e4cec0913f8ab1cafe386527f1e8e47b54e79299b14dfe604a3c8ce10ee4cb2d09be71e23628c13888c29a254815260c08301518552610f6189898989612c69565b7f146a85b7d5644a318ee8d2a44d995f580695fb92ade1dc0bbfc84cb4010ac1c386527f1d3734ea6fc621a0710ba60b024e87e0442525b2c77aff46634f6c5c0035e073815260e08301518552610fba89898989612c69565b7f12664f87a4a89be5df17d8f4cb797e102a083e505835ddfa68bfff968ded011f86527f062c4ef046ea89c77def686012d175292e6cca3adb8dc9e0234bbbe4384e6b188152610100830151855261101489898989612c69565b7f22ecd52df7b85f6392fa550619d401fedebde3332b7c8857f3b26caac17f4b7986527f25696ccea69b88905a95af209f6daa0d638426ae494b1b5800d1bb32aa271c3c8152610120830151855261106e89898989612c69565b7ee5e920eb2bd31bc7480c75e93c11a2a8f421af3e287f37a87453a96b4dd6c086527f2b82a4685f51b881675e3e4958242a11585b1aa2211bef1ba101ade59d484fcb815261014083015185526110c789898989612c69565b7f1bb112783a4bd8e95decec6012a6b0c06d076f07806241e64d5bb279055ce2e386527f03aacbe5b76da6ffe5a38b5f74bb1defdf60afde8f7bed77c2103c7d6d285c458152610160830151855261112189898989612c69565b7e89a49d6c462af5737f4f74e89ee3fd5fcda9b129a6885f4402b7191ce06ecc86527f03670d5e8e16d0f9e9351a53b707cddd87f0df01f7e71f8ef942b35a63c54d808152610180830151855261117a89898989612c69565b7f148203b1cf918d850c2e4eb482623111d69a0d149273d88bd472cc2dc667788886527f09e3685cddd844c117894cf1560ee45625a29890eb111e539598ac6c8510b26f81526101a083015185526111d489898989612c69565b7f038bda99e81e5aee528c18e38ab4a8806508a531d1f22f6618919bffb81f59f686527f14ce622b4ba47284ea8c421b00498ea220fb6f88026998098cbb21d38d8e0b0a81526101c0830151855261122e89898989612c69565b7f0c1641e14f8c4509f0c675448654f877363c981ca8c18363b549cfb115737ea386527f0b8cc9ed761dfbad8b821e125c571ba83a857c405c10df3bdeb64fc9b3248e2a81526101e0830151855261128889898989612c69565b7f14c11b03d9e6d4e5c71174b991b0b1b63f5a8539d4449e10a08275be0454646586527f014e6374d23fc81a10b61fd108e7b0e59003f8d3ed87edd40a722aafbc1c2635815261020083015185526112e289898989612c69565b7f2a195965e3a4ebe550289af22bb4c1118e21dc2c74be94ad6455e8f6eb70c8c486527f12024e0b3b82946c93024f8e1159da64dc1d3c72c49929836a5f2577d5a0ed688152610220830151855261133c89898989612c69565b7f1f679e8223e56ae364c7fe5b1eb44e3b3c66cbf45645c364d308c505539047c986527f2c32d125628fae7d840d3f28e83a7fa88112f60ee52f5f86fba53c08e474ff698152610240830151855261139689898989612c69565b7f1b73d0bdb2a03d112f31e25b60799a767fb82a9d6418db824c632ffba80be47d86527f21b9840d5d347552e43ea2e7bd19560353b633a278b602c6074025ddb9f63a21815261026083015185526113f089898989612c69565b7f16b45c0468819f85893ab1c4877c4fe4f49c146b976419b48aa07a0f6cbfb2e286527f06b26b18879ce8c03298c0302e22e9be3407e4b50aa15e3153b0eea99ce024148152610280830151855261144a89898989612c69565b7f1fe6017431f3e3861894f3e1871a3627f3fc61c832f3e951b3e55c86c4b9615886527f2236d30c0a8738dcac59bc76c975d2ac9c9f1347adfc85c977fc196f4a963f5e81526102a083015185526114a489898989612c69565b7f2fd10aed9958e2d8257e4a70a742fa4774402f36a25babac21c8e0b5f661c12886527f0f8a616d021292af83c9631904c885899f5fe78a5489c14462c97f8472ce5bbf81526102c083015185526114fe89898989612c69565b7f280a41e1efcd026c87f851e0180aad3ac57df1f93286a57f53c400268d8752c386527f1633532825ddc9c0da04246a44b706dcfd57ae2f4c1bb69738ff8433d5b2a8bb81526102e0830151855261155889898989612c69565b7e027cbc9193ee97f7eec57e57cb0ff7347cc0b2586a1637f4ce954bea3ff97186527f1c897b0f8a07cbc7a4b69597227129b4a12ba68e1926ecc7a45a4ec4bdc5bf07815261030083015185526115b189898989612c69565b7f1b1da35eea8e3139d38e9db84386f59853a18040d1a2216b74679e08c191a01c86527f0bfab57806284de52685f6dda04330043efee9399c75b78e48b5b1d7cb80038e8152610320830151855261160b89898989612c69565b7f012154f85b76ea46de9dc3f61d7c053aa9a583e3e2e57d6a076db599b1326a2f86527f1d25427c48b7647c1efe27b5e7da3240ee333d288ebabeb8b45e30c113c6474e8152610340830151855261166589898989612c69565b7f111fb275c27d543c507e0c685728727d2344f736a345419dd396d083296211ed86527f01c636dbb603223ca61aee282e75c2ee554f6639e813f990cc7e045128e9ba6e815261036083015185526116bf89898989612c69565b7f0e3ef51ec2992fd4fd4e08d2fc6c02cb6586ec574edec92b74583bd38cc15cae86527f0eeb3ece8b4b83ce8946832b6dd7f35204669e47ccd8b1fc31cba6c71808f6898152610380830151855261171989898989612c69565b7f0a7e2bd7bb8aeb9e84739db84898a9115aa023c279d2df4536366e445e618b1e86527f0141bc992ed56ad3af847ed62afe254e174e6df8efbc36cf3314adad1244b42381526103a0830151855261177389898989612c69565b7f0f2ecde94b061c256edb823ac557ce52f907f612791c85a2e66fb888ce8a417b86527f1c09474255a3b4c33f9452b362e6352c65acb1406a454a1d4b212538d529ed9a81526103c083015185526117cd89898989612c69565b7f1a2a4a634641112a1b940ba3b089193dfde76611dc7a7c29538a0f93bbdaf83086527f1752d5a3839dedbad8cd819b7b86a82982d1c7663453d236c559353725c2901381526103e0830151855261182789898989612c69565b7f2e9b2e2e4921cd57e24d3215d41cfa43545b0fa8907380f2b4eca856b242ace086527f192ca02d2e86b7636626a919c871396ba8108cc7f2358ded277a32ed4ba10d4b8152610400830151855261188189898989612c69565b7f1266b7cbb61c28d580a6aa8e4a6b3455bf5c925fe9321c0b0afd01596c4a950586527f0f945b9129e6749912477338802d286af2f6aca2e71cdf199f387d570acf5f10815261042083015185526118db89898989612c69565b7f1be6c06af2b7182fc509c0c72a3874173c0e9b8ef208c89a6ffcbe343f74440a86525201519052612c69565b51159081159161195d575b50611944575b50506119259051612cba565b9451908451908501519160606040870151955196015196612ed3565b90565b519193506119259161195591612d8d565b92905f611919565b90505115155f611913565b806119a17f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000161199a6119a6948e610779565b5110610a85565b610749565b610c2e565b6119b36100d3565b906119bc6100f8565b5f81525f602082015282526119cf6107a0565b60208301526119dc6107a0565b60408301526104896107a0565b6119f16119ab565b906119fa6100f8565b7f3059a4f6581fbdcc0d5e847306a4862568fbf62d22d49958fc3902e4bc379ef381526020907f1dd7c04824ae7a26cbb4bb2e8e647030df4de01e51764344c9da30b1ac0317c8828201528352611a4f6100f8565b7f0f98f7aa65c680ca7cd4b7f95afc5f6827e95398c599befdd9a9eca741e46b6181527f0a6f18525c7167edf6945aa57ecf8e0dde824c50912f1fe71574e86908579b6882820152611a9f6100f8565b7f1d1dacfe7971320b875aa2dbcdafb33d4141ca0b0435904e1e2ead83b600d26b81527f1e0cdae1aa15580307c121c8518d1b513fb8bfc62718065a60e88eee79d0288d83820152611aef6100f8565b9182528282015281840152611b026100f8565b7f1887c867c4428fd8c7157ac7f5e81a19271ea37ac336aa87203e65bd77cdcad281527f10d21089c03935120870563d17d271a9165f3d541cf4b41a450b8c3741e8442382820152611b526100f8565b7f105ceb8102cb4bd76c903c3f045988d7407aab02a26e2b0ac08c58af1edec5a481527f0b32e41ba74a1a65c885129ca2c3c87475d584c75ebc553cb79d4468de6483a883820152611ba26100f8565b918252828201526040840152611bb66100f8565b907f04e9e06a4684b3f9cfea22a0b5d19239c957ba0b12a17f2d9dc9d1e63f9ae49482527f0ef0df2626365d3222024595b1cd400614d2db4a442bb59f5ab585b4717528f881830152611c076100f8565b907f239dd78f7b5dba6a6d81f994b3060a73e4d7602aeb8909fd9785a0f1e04367e382527f178fb89664e86e6758f974887c6d9d19cd52b518c16d799e75556a7ae9a2582981830152611c586100f8565b9283528201526060830152565b9091939293611c726107cc565b918051602080920151611c836100f8565b918252828201528352611c946100f8565b938051518552611ca5815160200190565b5182860152611cb682610b236100f8565b5182820152611cc36100f8565b94855281850152808301938452808251920151611cde6100f8565b9283528183015260408301918252611d847f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47611d1d8186515110610825565b611d2c81848751015110610871565b611d3a8187515151106108bd565b611d4a8184885101515110610909565b611d5a81610bf488515160200190565b611d6c81610c0d858951015160200190565b611d7981855151106109ed565b828451015110610a39565b5f5b602181106129b25750946119419495611925611da06129e9565b946129ac611dac6100f8565b80955f82528582015f8152611dbf6106db565b91611dc86106eb565b907f305c9c1aa4a3294d7d6f331d65dc097fd9b9011350a6065eed3d56ad4d48a5e2611df26100f8565b935f85525f8b8601527f2d704aa2e65d5ad168e2ebdd814a37bf7f58961077a78e2ce2eea371cf886b0e8752527f05595e70b8c63dfb8fe2f8adb49c225ee5e6f783b5736f3221d423194db5585d8252610400898301917e7f03a8f871280b33c0292e973247afd9cddfa419e978678bde69323baece8383527f0e4ddacfcf167969963d2bb01ca92fa86b4bfd92c37917bad4ab760f0279625081519360408601948552611ea289898989612c69565b7f2cbe10c7c83f6dbe1bc89736f5dc9a9b91e55be6941a4b99b058acd8001fb00486527f1500351867036612a9db15b6f7d4198993f31565af0610ed3fda8d92ffb5e67181528c8301518552611efa89898989612c69565b7f2ff215cfa1c7c99fc0b382d9d7225a0636ccd9a01be83959da430a3a25d4481f86527f2918236b7c008f70fc98cc3bcc41eb60fb8b85d02b4e83100de7a9a7eed34645815260408301518552611f5389898989612c69565b7f0943a399c312616b46deee38e49b364b3449d7bab638b580df78aece26ecedf186527f11c937747adbd7b45585d385c0174241e937dce9fa818ca66d00f550b3f3d128815260608301518552611fac89898989612c69565b7f05dd54c0736f8a8d838a097e3151776cc6f602439ecc7833d223dc6cc69d585186527ee6465472f5b3647daa18848088f62f4fe1e57c401172d39f45b827bcc898f781526080830151855261200489898989612c69565b7f0d64a91e0e28a2e96b12a3211ebd9f66d854efbc1bd17fccc28e33993722b96d86527f0b30126ec8f40991e90ef34a78e10acb9848fe65e84547e64346fbaebafc9a98815260a0830151855261205d89898989612c69565b7f08272a78392bca7c2597f09f39daf6f7129808e6a6b34a49239dbf2a264e4e3b86527f02ac17971af65a980f6ab150b8ebaf573008d90c0af4eeed28d50374e8eed16b815260c083015185526120b689898989612c69565b7f19e3f8ff1265325376056bd2155edf79762433ec7b24c2196701da40ff2e6b3186527f0b33297cf3ae84083dab64e559bccd29e271c3e7d9ba307b7d623d223d091ead815260e0830151855261210f89898989612c69565b7f01341f47e05793a19ea7b92ff3b84c73f7223d56104d070f246e00eb1db7967986527f052fb71e652150885399cfb863b33f0dc0dfe9b279d0bf29df0fc403810fdb178152610100830151855261216989898989612c69565b7f04e2785630031d901e87c72d18c1e526142d7b79ee5b4360f5f9373d385aa1be86527f13fdbcc7f866114bdd363bc99c54b3ea0921b9b5526e46885291a00f8f9feea7815261012083015185526121c389898989612c69565b7f0bc888012014ac70c1692250c46020392de91756724c6c890edbc8d860450b8086527f02ce5d7f6c5c1ec90a48c08531305543ef94f25d0d03f67124280afb560406d48152610140830151855261221d89898989612c69565b7f0aee169033d6ddce5c5dde351a6ce45eb60dc3018752274e14f68e3f19a5359686527f1481dacc3a815550ede26690b83395cb031c17fb4aca75baa0a74acc404233518152610160830151855261227789898989612c69565b7f1c24a404b575734133f0c03bbf71c74c5b03a34f1d0817662c7d9293eaecaad586527f10af37e497d439239dfdab04bce89b1eaeaa12bc5f327129816a7a3841863fdd815261018083015185526122d189898989612c69565b7f14a256de1a2c7bb25672a1acc5406b90543c8a3b8c7c6e0a1895f019171141a886527f08f0539c9fd5ef363053e6560e4769e20e56dd0a510c30dce8adab0230c5fdb181526101a0830151855261232b89898989612c69565b7f082d61eb34a0a6544527e7db6b9bd16a8f6488733c83bee559ec5378191c93ba86527f144b7ab9c8fd2fb71b51c102820f2b08303be60f9fdd313e68f412d3f027a82081526101c0830151855261238589898989612c69565b7f22510e5ce22c30374993b2a360ceed3bdc20bd64b8d14cb3baedf76ddbd8062386527f0c0c88dfbf63ebc976d642a63c3d22288c546570b101f0219b3e2f3af5bedcfb81526101e083015185526123df89898989612c69565b7f2c71d775cc194f6e13408a12a33cb48babccfee137654d1443371de1d0f30c0c86527f1ba219dea8d4ffd8339c1c10cda690451c10fb5058f36ce7e1407118d871cc808152610200830151855261243989898989612c69565b7f14e7d115c5cfbe3c075697f305b8660abf41c5725a40557d3e14c9703aef64c286527f0cbc84b02d09b3f498b122ab5819248195aea678e0a41e744967da6bd8d0ce118152610220830151855261249389898989612c69565b7f0332f7d5660e970f229a174367929acabfa2f9fdab763460fd7acebddd944dde86527f15be1ce817121a7c25340b8d9c50a584a179d3dffd489f54311a12f922a6942f815261024083015185526124ed89898989612c69565b7f0257933903a2e91846df829f8084008ddf5fc35dd8d4acdebd426bff0d97e2a386527f17e9653840e81e1a68076e0c5f8c89f61463791e918991df86c86f63dccd93918152610260830151855261254789898989612c69565b7f106f1170be9c02c979b3d6e1d43737530d6bfc444c16df873400384d39e393be86527f1c2b9f619d809bb543e712ac0ee22cc3c6aa99ae73c97246c9769c9c98713fc5815261028083015185526125a189898989612c69565b7f131eb8c00ed76432c870a74c71365748a51807c021e678beb083b0e7e8b5b61186527f2491a76ab72146d0aeb330815df908ad5dd6cd86201e97d220b63d0d8d0f3ac481526102a083015185526125fb89898989612c69565b7f2561f4abb9fabeeb813dcc6d4d487d8f6e36fdd18805e785cccdf2b0a2ff085786527f2e6269f87539d6b464a25b6bd4522d1e2b78c3918f48e79de55dd640261faa4281526102c0830151855261265589898989612c69565b7f2e485df27f23a93b97e296061758e7dd3d34c4722c6fd7ae249433c55d259adc86527f1614f76a407ac31a5acb91266c2c7f54166ee17112e69f0b5e5fad32dafe5f6581526102e083015185526126af89898989612c69565b7f153af4e0fe4af748819ca675d8781da95111510763de4bec8abf25bca637703b86527f0b1f5a812c51999ebc7ac97de6ab2409b894ad5fb02d45f0798acc2548344cb98152610300830151855261270989898989612c69565b7f2c69e0646f6bfc70dfd02cf64d0b781ab481cafbd190dba9ee603f8160c44dcd86527f126015936956b109beba47938f9808cd9eed7ba5fc4531e6f6f267cadb13f4f88152610320830151855261276389898989612c69565b7f22c6d8b6cf6965d431abf72b985b44f1d0831026a43d6dc8cbcb9ca85ab4a0bc86527f19fd6ce3da2b55331cdd361a63f29e95c52f07bcf9cbe355077e5141bf020d83815261034083015185526127bd89898989612c69565b7f0cc996a6a427bcf59dad5cb6a2da92164e142372347a8ca5b1b3b32d8b20a0ef86527f232f08e45f51e15d57617fc960278b4dd236ed78e9014d43950751ea862841f18152610360830151855261281789898989612c69565b7f1e03c25c13870ae3e127e009be017aa0f47c7b53fc8636bd519ad68f035aa55c86527f04d8795eee1d4bcb8a4042ed861d91024ec5ce76c4fdac892d0c0047987451998152610380830151855261287189898989612c69565b7f2690cf34bdf3837f3036c3c61e73f94f8026f6a6e9be13695cf81ccebbac7ca886527f0db77d728541f6ab723d2dc8389a97864f9e61349120d35d76c004579c3fc10881526103a083015185526128cb89898989612c69565b7f10f3808b8fb7eb5be9d22fa5e7b4599e94116d5803607ae38bd3c637c10224f286527f06cb766b59e904c5b47098c0f98af9c4444f614cf303f1c939fd24b6bd60cc9c81526103c0830151855261292589898989612c69565b7f1610f9fbade90d90feada79ae229d67175dee93355fbbf229133ce0d6e75e3a186527f1a3084c2af6e7f823d2044866e00b80af74b5586f26c9685943786f355712ddd81526103e0830151855261297f89898989612c69565b7f1b88a7f08e12e3e28b9b10c0c0cbfd8d7df8b8a8fae1840b33c1cd4d24c8b23c86525201519052612c69565b51612cba565b806119a17f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000161199a6129e4948b61078f565b611d86565b6129f16119ab565b906129fa6100f8565b7f0a3a3884405b8d1fe46693685f02ba163634fd63d14bf91b6a433825b9ed6b5a81526020907f111fc830b029cfb2c94e450b570cf3be3eac81076213c2dcb1e1059330c605bf828201528352612a4f6100f8565b7f23ab779b99bf99c421500d8fe70c4e84fc1ff6eefdf3e92d8f581d046cb3eab781527f1611f26f3d9d6b19c4a418d02f19f6796be688f08507bc59ee5f9862dd46fa5b82820152612a9f6100f8565b7f2d34a3d654ca9ea36195f8167d653fa7240b0af8acad4b224aed268f9d8756ff81527f178cda417a663a79267fba64b28caf8fc8484866bfe0f423cb8d3b7da164d7f983820152612aef6100f8565b9182528282015281840152612b026100f8565b7f1f682eee4eeb25b38c3bff07fad9aaeb8c1ae87a95472a7819a57fd8b37a6e1581527f0db20bc4434468f4ce7f5888da80c6013c5392645400eee1ddbb77b0696ea1a782820152612b526100f8565b7f01979b2d16e0fb974244f72e399fd4d24be132523f4aeb010c75f26b6452d53c81527f21900fdcdfde4102dbbcd9525e925c0f4ea5317aefc7a1c350753b5c9741ebd583820152612ba26100f8565b918252828201526040840152612bb66100f8565b907f04969a13dd24e7586c1e7e668f9be1cfab2bfb7baf9e48cd94428a55b4cfb89882527f3026f4334a515ea181839681e5a601e08615013a7355b0a0ad1c6ffce279eb1681830152612c076100f8565b907f0987e27c310f4a785adc7dfc5324848dc4b1b4957907733a04c889777c88a78582527f13c07cb3a59387f85f315e9b41060f8a993a3c3d22113439d63f9be212afc23481830152611c586100f8565b90600481101561078a5760051b0190565b90929160608460806107cf19946007865a01fa15612cb857600660c0926020606096865185528187015182860152805160408601520151868401525a01fa8015612cb857612cb690612d41565b565bfe5b5f6020612cc56100f8565b8281520152805190811580612d35575b15612cf0575050612ce46100f8565b5f81525f602082015290565b602001517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790819006810390811161065d57612d2a6100f8565b918252602082015290565b50602081015115612cd5565b15612d4857565b60405162461bcd60e51b815260206004820152601260248201527f70616972696e672d6164642d6661696c656400000000000000000000000000006044820152606490fd5b60609092919260c0612d9d6100f8565b915f83525f60208401526020839681612db46100d3565b936080368637805185520151828401528051604084015201518482015260066107cf195a01fa8015612cb857612cb690612d41565b60405190610320820182811067ffffffffffffffff8211176100f35760405260188252610300366020840137565b9060068202918083046006149015171561065d57565b906001820180921161065d57565b906002820180921161065d57565b906003820180921161065d57565b906004820180921161065d57565b906005820180921161065d57565b805182101561078a5760209160051b010190565b15612e8e57565b60405162461bcd60e51b815260206004820152601560248201527f70616972696e672d6f70636f64652d6661696c656400000000000000000000006044820152606490fd5b9491959692909396612ee36100d3565b95865260209788978888015260408701526060860152612f016100d3565b9384528584015260408301526060820152612f1a612de9565b915f5b60048110612f5557505050610300612f33610710565b9384920160086107cf195a01fa8015612cb857612f4f90612e87565b51151590565b6130279192939450612f6681612e17565b612f708285612c58565b5151612f7c8288612e73565b5286612f888386612c58565b510151612f9d612f9783612e2d565b88612e73565b52612fa88286612c58565b515151612fb7612f9783612e3b565b52612fcd612fc58387612c58565b515160200190565b51612fda612f9783612e49565b5286612fe68387612c58565b51015151612ff6612f9783612e57565b5261302161301b6130148961300b868a612c58565b51015160200190565b5192612e65565b87612e73565b52610749565b9084939291612f1d56fea2646970667358221220e0f5fe942fa41ef5d7525d668cb9b820694b990e39df2c375409c733dfb1813064736f6c63430008140033",
}

// BeaconVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BeaconVerifierMetaData.ABI instead.
var BeaconVerifierABI = BeaconVerifierMetaData.ABI

// BeaconVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeaconVerifierMetaData.Bin instead.
var BeaconVerifierBin = BeaconVerifierMetaData.Bin

// DeployBeaconVerifier deploys a new Ethereum contract, binding an instance of BeaconVerifier to it.
func DeployBeaconVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BeaconVerifier, error) {
	parsed, err := BeaconVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeaconVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeaconVerifier{BeaconVerifierCaller: BeaconVerifierCaller{contract: contract}, BeaconVerifierTransactor: BeaconVerifierTransactor{contract: contract}, BeaconVerifierFilterer: BeaconVerifierFilterer{contract: contract}}, nil
}

// BeaconVerifier is an auto generated Go binding around an Ethereum contract.
type BeaconVerifier struct {
	BeaconVerifierCaller     // Read-only binding to the contract
	BeaconVerifierTransactor // Write-only binding to the contract
	BeaconVerifierFilterer   // Log filterer for contract events
}

// BeaconVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeaconVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeaconVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeaconVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeaconVerifierSession struct {
	Contract     *BeaconVerifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeaconVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeaconVerifierCallerSession struct {
	Contract *BeaconVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BeaconVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeaconVerifierTransactorSession struct {
	Contract     *BeaconVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BeaconVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeaconVerifierRaw struct {
	Contract *BeaconVerifier // Generic contract binding to access the raw methods on
}

// BeaconVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeaconVerifierCallerRaw struct {
	Contract *BeaconVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeaconVerifierTransactorRaw struct {
	Contract *BeaconVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeaconVerifier creates a new instance of BeaconVerifier, bound to a specific deployed contract.
func NewBeaconVerifier(address common.Address, backend bind.ContractBackend) (*BeaconVerifier, error) {
	contract, err := bindBeaconVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconVerifier{BeaconVerifierCaller: BeaconVerifierCaller{contract: contract}, BeaconVerifierTransactor: BeaconVerifierTransactor{contract: contract}, BeaconVerifierFilterer: BeaconVerifierFilterer{contract: contract}}, nil
}

// NewBeaconVerifierCaller creates a new read-only instance of BeaconVerifier, bound to a specific deployed contract.
func NewBeaconVerifierCaller(address common.Address, caller bind.ContractCaller) (*BeaconVerifierCaller, error) {
	contract, err := bindBeaconVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconVerifierCaller{contract: contract}, nil
}

// NewBeaconVerifierTransactor creates a new write-only instance of BeaconVerifier, bound to a specific deployed contract.
func NewBeaconVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconVerifierTransactor, error) {
	contract, err := bindBeaconVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconVerifierTransactor{contract: contract}, nil
}

// NewBeaconVerifierFilterer creates a new log filterer instance of BeaconVerifier, bound to a specific deployed contract.
func NewBeaconVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconVerifierFilterer, error) {
	contract, err := bindBeaconVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconVerifierFilterer{contract: contract}, nil
}

// bindBeaconVerifier binds a generic wrapper to an already deployed contract.
func bindBeaconVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeaconVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconVerifier *BeaconVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconVerifier.Contract.BeaconVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconVerifier *BeaconVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconVerifier.Contract.BeaconVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconVerifier *BeaconVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconVerifier.Contract.BeaconVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconVerifier *BeaconVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconVerifier *BeaconVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconVerifier *BeaconVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyBlsSigProof is a free data retrieval call binding the contract method 0x0cc26769.
//
// Solidity: function verifyBlsSigProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[35] input) view returns(bool r)
func (_BeaconVerifier *BeaconVerifierCaller) VerifyBlsSigProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [35]*big.Int) (bool, error) {
	var out []interface{}
	err := _BeaconVerifier.contract.Call(opts, &out, "verifyBlsSigProof", a, b, c, commit, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlsSigProof is a free data retrieval call binding the contract method 0x0cc26769.
//
// Solidity: function verifyBlsSigProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[35] input) view returns(bool r)
func (_BeaconVerifier *BeaconVerifierSession) VerifyBlsSigProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [35]*big.Int) (bool, error) {
	return _BeaconVerifier.Contract.VerifyBlsSigProof(&_BeaconVerifier.CallOpts, a, b, c, commit, input)
}

// VerifyBlsSigProof is a free data retrieval call binding the contract method 0x0cc26769.
//
// Solidity: function verifyBlsSigProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[35] input) view returns(bool r)
func (_BeaconVerifier *BeaconVerifierCallerSession) VerifyBlsSigProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [35]*big.Int) (bool, error) {
	return _BeaconVerifier.Contract.VerifyBlsSigProof(&_BeaconVerifier.CallOpts, a, b, c, commit, input)
}

// VerifyCommitteeRootMappingProof is a free data retrieval call binding the contract method 0xab00dde6.
//
// Solidity: function verifyCommitteeRootMappingProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[33] input) view returns(bool r)
func (_BeaconVerifier *BeaconVerifierCaller) VerifyCommitteeRootMappingProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [33]*big.Int) (bool, error) {
	var out []interface{}
	err := _BeaconVerifier.contract.Call(opts, &out, "verifyCommitteeRootMappingProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCommitteeRootMappingProof is a free data retrieval call binding the contract method 0xab00dde6.
//
// Solidity: function verifyCommitteeRootMappingProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[33] input) view returns(bool r)
func (_BeaconVerifier *BeaconVerifierSession) VerifyCommitteeRootMappingProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [33]*big.Int) (bool, error) {
	return _BeaconVerifier.Contract.VerifyCommitteeRootMappingProof(&_BeaconVerifier.CallOpts, a, b, c, input)
}

// VerifyCommitteeRootMappingProof is a free data retrieval call binding the contract method 0xab00dde6.
//
// Solidity: function verifyCommitteeRootMappingProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[33] input) view returns(bool r)
func (_BeaconVerifier *BeaconVerifierCallerSession) VerifyCommitteeRootMappingProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [33]*big.Int) (bool, error) {
	return _BeaconVerifier.Contract.VerifyCommitteeRootMappingProof(&_BeaconVerifier.CallOpts, a, b, c, input)
}

// VerifySignatureProof is a free data retrieval call binding the contract method 0x52356da0.
//
// Solidity: function verifySignatureProof(bytes32 signingRoot, bytes32 syncCommitteePoseidonRoot, uint256 participation, uint256 commitment, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_BeaconVerifier *BeaconVerifierCaller) VerifySignatureProof(opts *bind.CallOpts, signingRoot [32]byte, syncCommitteePoseidonRoot [32]byte, participation *big.Int, commitment *big.Int, p IBeaconVerifierProof) (bool, error) {
	var out []interface{}
	err := _BeaconVerifier.contract.Call(opts, &out, "verifySignatureProof", signingRoot, syncCommitteePoseidonRoot, participation, commitment, p)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignatureProof is a free data retrieval call binding the contract method 0x52356da0.
//
// Solidity: function verifySignatureProof(bytes32 signingRoot, bytes32 syncCommitteePoseidonRoot, uint256 participation, uint256 commitment, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_BeaconVerifier *BeaconVerifierSession) VerifySignatureProof(signingRoot [32]byte, syncCommitteePoseidonRoot [32]byte, participation *big.Int, commitment *big.Int, p IBeaconVerifierProof) (bool, error) {
	return _BeaconVerifier.Contract.VerifySignatureProof(&_BeaconVerifier.CallOpts, signingRoot, syncCommitteePoseidonRoot, participation, commitment, p)
}

// VerifySignatureProof is a free data retrieval call binding the contract method 0x52356da0.
//
// Solidity: function verifySignatureProof(bytes32 signingRoot, bytes32 syncCommitteePoseidonRoot, uint256 participation, uint256 commitment, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_BeaconVerifier *BeaconVerifierCallerSession) VerifySignatureProof(signingRoot [32]byte, syncCommitteePoseidonRoot [32]byte, participation *big.Int, commitment *big.Int, p IBeaconVerifierProof) (bool, error) {
	return _BeaconVerifier.Contract.VerifySignatureProof(&_BeaconVerifier.CallOpts, signingRoot, syncCommitteePoseidonRoot, participation, commitment, p)
}

// VerifySyncCommitteeRootMappingProof is a free data retrieval call binding the contract method 0x7a538781.
//
// Solidity: function verifySyncCommitteeRootMappingProof(bytes32 sszRoot, bytes32 poseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_BeaconVerifier *BeaconVerifierCaller) VerifySyncCommitteeRootMappingProof(opts *bind.CallOpts, sszRoot [32]byte, poseidonRoot [32]byte, p IBeaconVerifierProof) (bool, error) {
	var out []interface{}
	err := _BeaconVerifier.contract.Call(opts, &out, "verifySyncCommitteeRootMappingProof", sszRoot, poseidonRoot, p)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySyncCommitteeRootMappingProof is a free data retrieval call binding the contract method 0x7a538781.
//
// Solidity: function verifySyncCommitteeRootMappingProof(bytes32 sszRoot, bytes32 poseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_BeaconVerifier *BeaconVerifierSession) VerifySyncCommitteeRootMappingProof(sszRoot [32]byte, poseidonRoot [32]byte, p IBeaconVerifierProof) (bool, error) {
	return _BeaconVerifier.Contract.VerifySyncCommitteeRootMappingProof(&_BeaconVerifier.CallOpts, sszRoot, poseidonRoot, p)
}

// VerifySyncCommitteeRootMappingProof is a free data retrieval call binding the contract method 0x7a538781.
//
// Solidity: function verifySyncCommitteeRootMappingProof(bytes32 sszRoot, bytes32 poseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_BeaconVerifier *BeaconVerifierCallerSession) VerifySyncCommitteeRootMappingProof(sszRoot [32]byte, poseidonRoot [32]byte, p IBeaconVerifierProof) (bool, error) {
	return _BeaconVerifier.Contract.VerifySyncCommitteeRootMappingProof(&_BeaconVerifier.CallOpts, sszRoot, poseidonRoot, p)
}

// BlockChunksMetaData contains all meta data concerning the BlockChunks contract.
var BlockChunksMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateAnchorBlockProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numFinal\",\"type\":\"uint32\"}],\"name\":\"UpdateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateVerifierAddress\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"anchorBlockProviders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"startBlockNumber\",\"type\":\"uint32\"}],\"name\":\"historicalRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"claimedBlkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"numFinal\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[7]\",\"name\":\"merkleProof\",\"type\":\"bytes32[7]\"}],\"internalType\":\"structIBlockChunks.BlockHashWitness\",\"name\":\"witness\",\"type\":\"tuple\"}],\"name\":\"isBlockHashValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"updateAnchorBlockProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nextRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nextNumFinal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"}],\"name\":\"updateOld\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"}],\"name\":\"updateRecent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"}],\"name\":\"updateVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"verifierAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005a575f8054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3610f19908161005f8239f35b5f80fdfe60406080815260049081361015610014575f80fd5b5f91823560e01c9081631d8441a1146109f557816350a26d48146107ac57816352cfc560146105a5578163544f73a4146103985781635d727e9714610357578163715018a6146102f75781638da5cb5b146102d1578163b551a1871461027a578163ec4ffc52146101c5578163f2fde38b146100dc575063f5cec6af14610099575f80fd5b346100d85760203660031901126100d8576001600160a01b038160209367ffffffffffffffff6100c7610aa4565b168152600185522054169051908152f35b5080fd5b919050346101c15760203660031901126101c1578135916001600160a01b03908184168094036101bd57845491821692610117338514610b03565b84156101545750506001600160a01b031916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b8480fd5b8280fd5b5050346100d857806003193601126100d8577ffd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f90610201610aa4565b61027461020c610abf565b926001600160a01b03610223818854163314610b03565b67ffffffffffffffff8416875260016020528187209085166001600160a01b031982541617905551928392839092916001600160a01b0360209167ffffffffffffffff604085019616845216910152565b0390a180f35b5050346100d857806003193601126100d857610294610aa4565b6024359063ffffffff82168092036102cd5792829167ffffffffffffffff60209516825260038552828220908252845220549051908152f35b8380fd5b5050346100d857816003193601126100d8576001600160a01b0360209254169051908152f35b83346103545780600319360112610354578080546001600160a01b03196001600160a01b0382169161032a338414610b03565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b5050346100d85760203660031901126100d8576001600160a01b038160209367ffffffffffffffff610387610aa4565b168152600285522054169051908152f35b839150346100d8576101803660031901126100d857604435801561056257607f6103c0610eaa565b16936103d3856103ce610eaa565b610b4e565b833567ffffffffffffffff8116810361055e57906104149167ffffffffffffffff165f52600360205263ffffffff60405f2091165f5260205260405f205490565b91821561051b579184925b60ff841660078110156104b55787811c6001166104875761043f90610ebd565b35835190602082019283528482015283815261045a81610c10565b519020925b60ff809116908114610474576001019261041f565b634e487b7160e01b865260118552602486fd5b61049090610ebd565b359083519060208201928352848201528381526104ac81610c10565b5190209261045f565b5060843593508563ffffffff8516850361035457506105026105106020958551928391888301956064358791604493918352602083015263ffffffff60e01b9060e01b1660408201520190565b03601f198101835282610c40565b519020149051908152f35b815162461bcd60e51b8152602081860152601a60248201527f626c6b20686973746f7279206e6f742073746f726564207965740000000000006044820152606490fd5b8580fd5b835162461bcd60e51b8152602081840152601a60248201527f636c61696d6564426c6b48617368206e6f742070726573656e740000000000006044820152606490fd5b8383346100d85760803660031901126100d8576105c0610aa4565b60443563ffffffff9283821682036101bd5767ffffffffffffffff6064358181116107a857859291906105f69036908a01610ad5565b61060281839993610e12565b8682607f819f979d969c94161561061890610bc4565b6106229083610b4e565b16607f1461062f90610b78565b16600101858111610794576102749896948b9a98969461070a6106ed957fa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e9f9e956106a4610710966107159867ffffffffffffffff165f52600360205263ffffffff60405f2091165f5260205260405f205490565b8b5160208101928352602435604082015260e09390931b7fffffffff00000000000000000000000000000000000000000000000000000000166060840152919788906064820190565b0397610701601f19998a8101835282610c40565b51902014610c62565b8b610cfa565b610cae565b83516020810186815260408201889052600160e71b606083015260649283018252916107419082610c40565b5190209187168a526003602052828a209088168a5260205281892055519485948590949360809363ffffffff859467ffffffffffffffff60a0860199168552166020840152604083015260608201520152565b50634e487b7160e01b8c5260118d5260248cfd5b8680fd5b919050346101c157806003193601126101c1576107c7610aa4565b67ffffffffffffffff9260243584811161055e576107e89036908301610ad5565b9190936107f58386610e12565b929593969163ffffffff9991999a8b8b1694608086018d81116109e2578d165f1901908d82116109e2578d169c8e959493929161083491168e14610b78565b610841607f8d1615610bc4565b8a169586855287602095600287526001600160a01b039182912054161561099f579085918f808a8c92526002855220541689519e8f9263f25b3f9960e01b8452830152815a91602492fa9b8c15610995578d9c61094a575b50610274989796959493926108d79261070a610710937fa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e9f14610c62565b835181810186815260208101889052600160e71b6040820152906108fe8160448401610502565b519020928b5260038152838b20918b525281892055519485948590949360809363ffffffff859467ffffffffffffffff60a0860199168552166020840152604083015260608201520152565b9198979695949380939c5082813d831161098e575b6109698183610c40565b8101031261098a5790519a9697959694959394929391929190610274610899565b8c80fd5b503d61095f565b87513d8f823e3d90fd5b885162461bcd60e51b8152808301879052601d60248201527f636861696e20616e63686f722070726f7669646572206e6f74207365740000006044820152606490fd5b634e487b7160e01b8f526011885260248ffd5b5050346100d857806003193601126100d8577f4b8c49e37c813f3cbe140bc7b6a5662bd14e48311b5864c106ca3f9c4a2bc02b90610a31610aa4565b610274610a3c610abf565b926001600160a01b03610a53818854163314610b03565b67ffffffffffffffff8416875260026020528187209085166001600160a01b031982541617905551928392839092916001600160a01b0360209167ffffffffffffffff604085019616845216910152565b6004359067ffffffffffffffff82168203610abb57565b5f80fd5b602435906001600160a01b0382168203610abb57565b9181601f84011215610abb5782359167ffffffffffffffff8311610abb5760208381860195010111610abb57565b15610b0a57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b63ffffffff9182169082160391908211610b6457565b634e487b7160e01b5f52601160045260245ffd5b15610b7f57565b60405162461bcd60e51b815260206004820152600d60248201527f6e6565642031323820626c6b73000000000000000000000000000000000000006044820152606490fd5b15610bcb57565b60405162461bcd60e51b815260206004820152601460248201527f6e6565642073746172742066726f6d20313238780000000000000000000000006044820152606490fd5b6060810190811067ffffffffffffffff821117610c2c57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff821117610c2c57604052565b15610c6957565b60405162461bcd60e51b815260206004820152601360248201527f656e6448617368206e6f7420636f7272656374000000000000000000000000006044820152606490fd5b15610cb557565b60405162461bcd60e51b815260206004820152600f60248201527f70726f6f66206e6f742076616c696400000000000000000000000000000000006044820152606490fd5b67ffffffffffffffff16915f91838352602093600185526001600160a01b03908160408620541615610dcd5791839186938652600184526040862054169060446040518096819582946322bb937360e11b8452886004850152816024850152848401378181018301899052601f01601f191681010301915afa928315610dc2578293610d87575b50505090565b9080929350813d8311610dbb575b610d9f8183610c40565b810103126100d8575190811515820361035457505f8080610d81565b503d610d95565b6040513d84823e3d90fd5b60405162461bcd60e51b815260048101879052601660248201527f636861696e207665726966696572206e6f7420736574000000000000000000006044820152606490fd5b9190918261016011610abb5761018092808411610abb57608092610170830135841c610140840135851b1794826101a011610abb576101c090838211610abb576101b0850135861c90850135861b1794836101e011610abb578361020011610abb576101f0850135811c91850135901b17928261022011610abb5761021c81013560e01c9261024011610abb5761023c013560e01c90565b60243563ffffffff81168103610abb5790565b6007811015610ecf5760051b60a40190565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212209a76cb696151019727c90001fe248cba5c170828342da8cc1ab4115427c1810464736f6c63430008140033",
}

// BlockChunksABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockChunksMetaData.ABI instead.
var BlockChunksABI = BlockChunksMetaData.ABI

// BlockChunksBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockChunksMetaData.Bin instead.
var BlockChunksBin = BlockChunksMetaData.Bin

// DeployBlockChunks deploys a new Ethereum contract, binding an instance of BlockChunks to it.
func DeployBlockChunks(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlockChunks, error) {
	parsed, err := BlockChunksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockChunksBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockChunks{BlockChunksCaller: BlockChunksCaller{contract: contract}, BlockChunksTransactor: BlockChunksTransactor{contract: contract}, BlockChunksFilterer: BlockChunksFilterer{contract: contract}}, nil
}

// BlockChunks is an auto generated Go binding around an Ethereum contract.
type BlockChunks struct {
	BlockChunksCaller     // Read-only binding to the contract
	BlockChunksTransactor // Write-only binding to the contract
	BlockChunksFilterer   // Log filterer for contract events
}

// BlockChunksCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockChunksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockChunksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockChunksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockChunksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockChunksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockChunksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockChunksSession struct {
	Contract     *BlockChunks      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockChunksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockChunksCallerSession struct {
	Contract *BlockChunksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BlockChunksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockChunksTransactorSession struct {
	Contract     *BlockChunksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BlockChunksRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockChunksRaw struct {
	Contract *BlockChunks // Generic contract binding to access the raw methods on
}

// BlockChunksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockChunksCallerRaw struct {
	Contract *BlockChunksCaller // Generic read-only contract binding to access the raw methods on
}

// BlockChunksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockChunksTransactorRaw struct {
	Contract *BlockChunksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockChunks creates a new instance of BlockChunks, bound to a specific deployed contract.
func NewBlockChunks(address common.Address, backend bind.ContractBackend) (*BlockChunks, error) {
	contract, err := bindBlockChunks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockChunks{BlockChunksCaller: BlockChunksCaller{contract: contract}, BlockChunksTransactor: BlockChunksTransactor{contract: contract}, BlockChunksFilterer: BlockChunksFilterer{contract: contract}}, nil
}

// NewBlockChunksCaller creates a new read-only instance of BlockChunks, bound to a specific deployed contract.
func NewBlockChunksCaller(address common.Address, caller bind.ContractCaller) (*BlockChunksCaller, error) {
	contract, err := bindBlockChunks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockChunksCaller{contract: contract}, nil
}

// NewBlockChunksTransactor creates a new write-only instance of BlockChunks, bound to a specific deployed contract.
func NewBlockChunksTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockChunksTransactor, error) {
	contract, err := bindBlockChunks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockChunksTransactor{contract: contract}, nil
}

// NewBlockChunksFilterer creates a new log filterer instance of BlockChunks, bound to a specific deployed contract.
func NewBlockChunksFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockChunksFilterer, error) {
	contract, err := bindBlockChunks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockChunksFilterer{contract: contract}, nil
}

// bindBlockChunks binds a generic wrapper to an already deployed contract.
func bindBlockChunks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockChunksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockChunks *BlockChunksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockChunks.Contract.BlockChunksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockChunks *BlockChunksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockChunks.Contract.BlockChunksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockChunks *BlockChunksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockChunks.Contract.BlockChunksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockChunks *BlockChunksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockChunks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockChunks *BlockChunksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockChunks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockChunks *BlockChunksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockChunks.Contract.contract.Transact(opts, method, params...)
}

// AnchorBlockProviders is a free data retrieval call binding the contract method 0x5d727e97.
//
// Solidity: function anchorBlockProviders(uint64 ) view returns(address)
func (_BlockChunks *BlockChunksCaller) AnchorBlockProviders(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _BlockChunks.contract.Call(opts, &out, "anchorBlockProviders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnchorBlockProviders is a free data retrieval call binding the contract method 0x5d727e97.
//
// Solidity: function anchorBlockProviders(uint64 ) view returns(address)
func (_BlockChunks *BlockChunksSession) AnchorBlockProviders(arg0 uint64) (common.Address, error) {
	return _BlockChunks.Contract.AnchorBlockProviders(&_BlockChunks.CallOpts, arg0)
}

// AnchorBlockProviders is a free data retrieval call binding the contract method 0x5d727e97.
//
// Solidity: function anchorBlockProviders(uint64 ) view returns(address)
func (_BlockChunks *BlockChunksCallerSession) AnchorBlockProviders(arg0 uint64) (common.Address, error) {
	return _BlockChunks.Contract.AnchorBlockProviders(&_BlockChunks.CallOpts, arg0)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0xb551a187.
//
// Solidity: function historicalRoots(uint64 chainId, uint32 startBlockNumber) view returns(bytes32)
func (_BlockChunks *BlockChunksCaller) HistoricalRoots(opts *bind.CallOpts, chainId uint64, startBlockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _BlockChunks.contract.Call(opts, &out, "historicalRoots", chainId, startBlockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HistoricalRoots is a free data retrieval call binding the contract method 0xb551a187.
//
// Solidity: function historicalRoots(uint64 chainId, uint32 startBlockNumber) view returns(bytes32)
func (_BlockChunks *BlockChunksSession) HistoricalRoots(chainId uint64, startBlockNumber uint32) ([32]byte, error) {
	return _BlockChunks.Contract.HistoricalRoots(&_BlockChunks.CallOpts, chainId, startBlockNumber)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0xb551a187.
//
// Solidity: function historicalRoots(uint64 chainId, uint32 startBlockNumber) view returns(bytes32)
func (_BlockChunks *BlockChunksCallerSession) HistoricalRoots(chainId uint64, startBlockNumber uint32) ([32]byte, error) {
	return _BlockChunks.Contract.HistoricalRoots(&_BlockChunks.CallOpts, chainId, startBlockNumber)
}

// IsBlockHashValid is a free data retrieval call binding the contract method 0x544f73a4.
//
// Solidity: function isBlockHashValid((uint64,uint32,bytes32,bytes32,uint32,bytes32[7]) witness) view returns(bool)
func (_BlockChunks *BlockChunksCaller) IsBlockHashValid(opts *bind.CallOpts, witness IBlockChunksBlockHashWitness) (bool, error) {
	var out []interface{}
	err := _BlockChunks.contract.Call(opts, &out, "isBlockHashValid", witness)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockHashValid is a free data retrieval call binding the contract method 0x544f73a4.
//
// Solidity: function isBlockHashValid((uint64,uint32,bytes32,bytes32,uint32,bytes32[7]) witness) view returns(bool)
func (_BlockChunks *BlockChunksSession) IsBlockHashValid(witness IBlockChunksBlockHashWitness) (bool, error) {
	return _BlockChunks.Contract.IsBlockHashValid(&_BlockChunks.CallOpts, witness)
}

// IsBlockHashValid is a free data retrieval call binding the contract method 0x544f73a4.
//
// Solidity: function isBlockHashValid((uint64,uint32,bytes32,bytes32,uint32,bytes32[7]) witness) view returns(bool)
func (_BlockChunks *BlockChunksCallerSession) IsBlockHashValid(witness IBlockChunksBlockHashWitness) (bool, error) {
	return _BlockChunks.Contract.IsBlockHashValid(&_BlockChunks.CallOpts, witness)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockChunks *BlockChunksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockChunks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockChunks *BlockChunksSession) Owner() (common.Address, error) {
	return _BlockChunks.Contract.Owner(&_BlockChunks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockChunks *BlockChunksCallerSession) Owner() (common.Address, error) {
	return _BlockChunks.Contract.Owner(&_BlockChunks.CallOpts)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_BlockChunks *BlockChunksCaller) VerifierAddresses(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _BlockChunks.contract.Call(opts, &out, "verifierAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_BlockChunks *BlockChunksSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _BlockChunks.Contract.VerifierAddresses(&_BlockChunks.CallOpts, arg0)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_BlockChunks *BlockChunksCallerSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _BlockChunks.Contract.VerifierAddresses(&_BlockChunks.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockChunks *BlockChunksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockChunks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockChunks *BlockChunksSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockChunks.Contract.RenounceOwnership(&_BlockChunks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockChunks *BlockChunksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockChunks.Contract.RenounceOwnership(&_BlockChunks.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockChunks *BlockChunksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockChunks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockChunks *BlockChunksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockChunks.Contract.TransferOwnership(&_BlockChunks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockChunks *BlockChunksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockChunks.Contract.TransferOwnership(&_BlockChunks.TransactOpts, newOwner)
}

// UpdateAnchorBlockProvider is a paid mutator transaction binding the contract method 0x1d8441a1.
//
// Solidity: function updateAnchorBlockProvider(uint64 _chainId, address _provider) returns()
func (_BlockChunks *BlockChunksTransactor) UpdateAnchorBlockProvider(opts *bind.TransactOpts, _chainId uint64, _provider common.Address) (*types.Transaction, error) {
	return _BlockChunks.contract.Transact(opts, "updateAnchorBlockProvider", _chainId, _provider)
}

// UpdateAnchorBlockProvider is a paid mutator transaction binding the contract method 0x1d8441a1.
//
// Solidity: function updateAnchorBlockProvider(uint64 _chainId, address _provider) returns()
func (_BlockChunks *BlockChunksSession) UpdateAnchorBlockProvider(_chainId uint64, _provider common.Address) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateAnchorBlockProvider(&_BlockChunks.TransactOpts, _chainId, _provider)
}

// UpdateAnchorBlockProvider is a paid mutator transaction binding the contract method 0x1d8441a1.
//
// Solidity: function updateAnchorBlockProvider(uint64 _chainId, address _provider) returns()
func (_BlockChunks *BlockChunksTransactorSession) UpdateAnchorBlockProvider(_chainId uint64, _provider common.Address) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateAnchorBlockProvider(&_BlockChunks.TransactOpts, _chainId, _provider)
}

// UpdateOld is a paid mutator transaction binding the contract method 0x52cfc560.
//
// Solidity: function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes proofData) returns()
func (_BlockChunks *BlockChunksTransactor) UpdateOld(opts *bind.TransactOpts, chainId uint64, nextRoot [32]byte, nextNumFinal uint32, proofData []byte) (*types.Transaction, error) {
	return _BlockChunks.contract.Transact(opts, "updateOld", chainId, nextRoot, nextNumFinal, proofData)
}

// UpdateOld is a paid mutator transaction binding the contract method 0x52cfc560.
//
// Solidity: function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes proofData) returns()
func (_BlockChunks *BlockChunksSession) UpdateOld(chainId uint64, nextRoot [32]byte, nextNumFinal uint32, proofData []byte) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateOld(&_BlockChunks.TransactOpts, chainId, nextRoot, nextNumFinal, proofData)
}

// UpdateOld is a paid mutator transaction binding the contract method 0x52cfc560.
//
// Solidity: function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes proofData) returns()
func (_BlockChunks *BlockChunksTransactorSession) UpdateOld(chainId uint64, nextRoot [32]byte, nextNumFinal uint32, proofData []byte) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateOld(&_BlockChunks.TransactOpts, chainId, nextRoot, nextNumFinal, proofData)
}

// UpdateRecent is a paid mutator transaction binding the contract method 0x50a26d48.
//
// Solidity: function updateRecent(uint64 chainId, bytes proofData) returns()
func (_BlockChunks *BlockChunksTransactor) UpdateRecent(opts *bind.TransactOpts, chainId uint64, proofData []byte) (*types.Transaction, error) {
	return _BlockChunks.contract.Transact(opts, "updateRecent", chainId, proofData)
}

// UpdateRecent is a paid mutator transaction binding the contract method 0x50a26d48.
//
// Solidity: function updateRecent(uint64 chainId, bytes proofData) returns()
func (_BlockChunks *BlockChunksSession) UpdateRecent(chainId uint64, proofData []byte) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateRecent(&_BlockChunks.TransactOpts, chainId, proofData)
}

// UpdateRecent is a paid mutator transaction binding the contract method 0x50a26d48.
//
// Solidity: function updateRecent(uint64 chainId, bytes proofData) returns()
func (_BlockChunks *BlockChunksTransactorSession) UpdateRecent(chainId uint64, proofData []byte) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateRecent(&_BlockChunks.TransactOpts, chainId, proofData)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_BlockChunks *BlockChunksTransactor) UpdateVerifierAddress(opts *bind.TransactOpts, _chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _BlockChunks.contract.Transact(opts, "updateVerifierAddress", _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_BlockChunks *BlockChunksSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateVerifierAddress(&_BlockChunks.TransactOpts, _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_BlockChunks *BlockChunksTransactorSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _BlockChunks.Contract.UpdateVerifierAddress(&_BlockChunks.TransactOpts, _chainId, _verifierAddress)
}

// BlockChunksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockChunks contract.
type BlockChunksOwnershipTransferredIterator struct {
	Event *BlockChunksOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockChunksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockChunksOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockChunksOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockChunksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockChunksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockChunksOwnershipTransferred represents a OwnershipTransferred event raised by the BlockChunks contract.
type BlockChunksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockChunks *BlockChunksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockChunksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockChunks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockChunksOwnershipTransferredIterator{contract: _BlockChunks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockChunks *BlockChunksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockChunksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockChunks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockChunksOwnershipTransferred)
				if err := _BlockChunks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockChunks *BlockChunksFilterer) ParseOwnershipTransferred(log types.Log) (*BlockChunksOwnershipTransferred, error) {
	event := new(BlockChunksOwnershipTransferred)
	if err := _BlockChunks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockChunksUpdateAnchorBlockProviderIterator is returned from FilterUpdateAnchorBlockProvider and is used to iterate over the raw logs and unpacked data for UpdateAnchorBlockProvider events raised by the BlockChunks contract.
type BlockChunksUpdateAnchorBlockProviderIterator struct {
	Event *BlockChunksUpdateAnchorBlockProvider // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockChunksUpdateAnchorBlockProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockChunksUpdateAnchorBlockProvider)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockChunksUpdateAnchorBlockProvider)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockChunksUpdateAnchorBlockProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockChunksUpdateAnchorBlockProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockChunksUpdateAnchorBlockProvider represents a UpdateAnchorBlockProvider event raised by the BlockChunks contract.
type BlockChunksUpdateAnchorBlockProvider struct {
	ChainId    uint64
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAnchorBlockProvider is a free log retrieval operation binding the contract event 0x4b8c49e37c813f3cbe140bc7b6a5662bd14e48311b5864c106ca3f9c4a2bc02b.
//
// Solidity: event UpdateAnchorBlockProvider(uint64 chainId, address newAddress)
func (_BlockChunks *BlockChunksFilterer) FilterUpdateAnchorBlockProvider(opts *bind.FilterOpts) (*BlockChunksUpdateAnchorBlockProviderIterator, error) {

	logs, sub, err := _BlockChunks.contract.FilterLogs(opts, "UpdateAnchorBlockProvider")
	if err != nil {
		return nil, err
	}
	return &BlockChunksUpdateAnchorBlockProviderIterator{contract: _BlockChunks.contract, event: "UpdateAnchorBlockProvider", logs: logs, sub: sub}, nil
}

// WatchUpdateAnchorBlockProvider is a free log subscription operation binding the contract event 0x4b8c49e37c813f3cbe140bc7b6a5662bd14e48311b5864c106ca3f9c4a2bc02b.
//
// Solidity: event UpdateAnchorBlockProvider(uint64 chainId, address newAddress)
func (_BlockChunks *BlockChunksFilterer) WatchUpdateAnchorBlockProvider(opts *bind.WatchOpts, sink chan<- *BlockChunksUpdateAnchorBlockProvider) (event.Subscription, error) {

	logs, sub, err := _BlockChunks.contract.WatchLogs(opts, "UpdateAnchorBlockProvider")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockChunksUpdateAnchorBlockProvider)
				if err := _BlockChunks.contract.UnpackLog(event, "UpdateAnchorBlockProvider", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAnchorBlockProvider is a log parse operation binding the contract event 0x4b8c49e37c813f3cbe140bc7b6a5662bd14e48311b5864c106ca3f9c4a2bc02b.
//
// Solidity: event UpdateAnchorBlockProvider(uint64 chainId, address newAddress)
func (_BlockChunks *BlockChunksFilterer) ParseUpdateAnchorBlockProvider(log types.Log) (*BlockChunksUpdateAnchorBlockProvider, error) {
	event := new(BlockChunksUpdateAnchorBlockProvider)
	if err := _BlockChunks.contract.UnpackLog(event, "UpdateAnchorBlockProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockChunksUpdateEventIterator is returned from FilterUpdateEvent and is used to iterate over the raw logs and unpacked data for UpdateEvent events raised by the BlockChunks contract.
type BlockChunksUpdateEventIterator struct {
	Event *BlockChunksUpdateEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockChunksUpdateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockChunksUpdateEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockChunksUpdateEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockChunksUpdateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockChunksUpdateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockChunksUpdateEvent represents a UpdateEvent event raised by the BlockChunks contract.
type BlockChunksUpdateEvent struct {
	ChainId          uint64
	StartBlockNumber uint32
	PrevHash         [32]byte
	Root             [32]byte
	NumFinal         uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateEvent is a free log retrieval operation binding the contract event 0xa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e.
//
// Solidity: event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal)
func (_BlockChunks *BlockChunksFilterer) FilterUpdateEvent(opts *bind.FilterOpts) (*BlockChunksUpdateEventIterator, error) {

	logs, sub, err := _BlockChunks.contract.FilterLogs(opts, "UpdateEvent")
	if err != nil {
		return nil, err
	}
	return &BlockChunksUpdateEventIterator{contract: _BlockChunks.contract, event: "UpdateEvent", logs: logs, sub: sub}, nil
}

// WatchUpdateEvent is a free log subscription operation binding the contract event 0xa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e.
//
// Solidity: event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal)
func (_BlockChunks *BlockChunksFilterer) WatchUpdateEvent(opts *bind.WatchOpts, sink chan<- *BlockChunksUpdateEvent) (event.Subscription, error) {

	logs, sub, err := _BlockChunks.contract.WatchLogs(opts, "UpdateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockChunksUpdateEvent)
				if err := _BlockChunks.contract.UnpackLog(event, "UpdateEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateEvent is a log parse operation binding the contract event 0xa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e.
//
// Solidity: event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal)
func (_BlockChunks *BlockChunksFilterer) ParseUpdateEvent(log types.Log) (*BlockChunksUpdateEvent, error) {
	event := new(BlockChunksUpdateEvent)
	if err := _BlockChunks.contract.UnpackLog(event, "UpdateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockChunksUpdateVerifierAddressIterator is returned from FilterUpdateVerifierAddress and is used to iterate over the raw logs and unpacked data for UpdateVerifierAddress events raised by the BlockChunks contract.
type BlockChunksUpdateVerifierAddressIterator struct {
	Event *BlockChunksUpdateVerifierAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockChunksUpdateVerifierAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockChunksUpdateVerifierAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockChunksUpdateVerifierAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockChunksUpdateVerifierAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockChunksUpdateVerifierAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockChunksUpdateVerifierAddress represents a UpdateVerifierAddress event raised by the BlockChunks contract.
type BlockChunksUpdateVerifierAddress struct {
	ChainId    uint64
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifierAddress is a free log retrieval operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_BlockChunks *BlockChunksFilterer) FilterUpdateVerifierAddress(opts *bind.FilterOpts) (*BlockChunksUpdateVerifierAddressIterator, error) {

	logs, sub, err := _BlockChunks.contract.FilterLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return &BlockChunksUpdateVerifierAddressIterator{contract: _BlockChunks.contract, event: "UpdateVerifierAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifierAddress is a free log subscription operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_BlockChunks *BlockChunksFilterer) WatchUpdateVerifierAddress(opts *bind.WatchOpts, sink chan<- *BlockChunksUpdateVerifierAddress) (event.Subscription, error) {

	logs, sub, err := _BlockChunks.contract.WatchLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockChunksUpdateVerifierAddress)
				if err := _BlockChunks.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateVerifierAddress is a log parse operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_BlockChunks *BlockChunksFilterer) ParseUpdateVerifierAddress(log types.Log) (*BlockChunksUpdateVerifierAddress, error) {
	event := new(BlockChunksUpdateVerifierAddress)
	if err := _BlockChunks.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlsSigVerifierMetaData contains all meta data concerning the BlsSigVerifier contract.
var BlsSigVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[35]\",\"name\":\"input\",\"type\":\"uint256[35]\"}],\"name\":\"verifyBlsSigProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657611bfb908161001b8239f35b5f80fdfe6040608081526004361015610012575f80fd5b5f803560e01c630cc2676914610026575f80fd5b346100c4576105a03660031901126100c45761004136610120565b9036606312156100c4576100536100db565b908160c4913683116100c457506044905b8282106100ab576100a786610096878761007d36610161565b6100863661019f565b916100903661021f565b93610651565b905190151581529081906020820190565b0390f35b602086916100b936856101df565b815201910190610064565b80fd5b634e487b7160e01b5f52604160045260245ffd5b604051906040820182811067ffffffffffffffff8211176100fb57604052565b6100c7565b604051906080820182811067ffffffffffffffff8211176100fb57604052565b806023121561015d576101316100db565b9081604491821161015d576004905b82821061014d5750505090565b8135815260209182019101610140565b5f80fd5b8060e3121561015d576101726100db565b908161010491821161015d5760c4905b82821061018f5750505090565b8135815260209182019101610182565b80610123121561015d576101b16100db565b908161014491821161015d57610104905b8282106101cf5750505090565b81358152602091820191016101c2565b9080601f8301121561015d576101f36100db565b80926040810192831161015d57905b82821061020f5750505090565b8135815260209182019101610202565b80610163121561015d5760405190610460820182811067ffffffffffffffff8211176100fb57604052816105a491821161015d57610144905b8282106102655750505090565b8135815260209182019101610258565b61027d610100565b906080368337565b604051906060820182811067ffffffffffffffff8211176100fb576040526060368337565b604051906020820182811067ffffffffffffffff8211176100fb576040526020368337565b6102d76100db565b906102e06100db565b604036823782526102ef6100db565b60403682376020830152565b610303610100565b9061030c6100db565b5f9081815281602082015283526103216102cf565b602084015261032e6100db565b81815281602082015260408401526103446100db565b9080825260208201526060830152565b634e487b7160e01b5f52603260045260245ffd5b1561036f57565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d710000000000000000006044820152606490fd5b156103bb57565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d710000000000000000006044820152606490fd5b1561040757565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561045357565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561049f57565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d7100000000000000006044820152606490fd5b156104eb57565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d7100000000000000006044820152606490fd5b1561053757565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d710000000000000000006044820152606490fd5b1561058357565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d710000000000000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b5f1981146105ea5760010190565b6105c8565b9060238110156106005760051b0190565b610354565b1561060c57565b60405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c64006044820152606490fd5b949394929190926106606102fb565b81516020809301516106706100db565b9182528382015281526106816100db565b948051518652610692815160200190565b51838701526106b1836106a36100db565b920180515183525160200190565b51838201526106be6100db565b958652828601528181019485528183519301516106d96100db565b93845282840152604081019283528351938281019485516106f86100db565b91825284820152606083019081526107ac7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd476107378186515110610368565b610746818787510151106103b4565b610754818a51515110610400565b61076481878b510151511061044c565b61077b816107748b515160200190565b5110610498565b6107948161078d888c51015160200190565b51106104e4565b6107a18188515110610530565b85875101511061057c565b5f5b602381106114e85750906114c197989392916107c8611570565b966107d16100db565b925f845261148884878101985f8a5281996107ea610275565b918a6104406107f7610285565b927f2065b90c648581703a4ef82833653ae713aaf62c2dc4ef26b0a9bbbdf254b48a6108216100db565b955f87525f858801527f0da0d09dcc32c2d20c9905307190ffb91538db928804c70e7ed77639f2cee0fe8952527f2fcf362c494439bcae24ab0ab7dd0bd40825ed18725c1d11d25eeb863f24194884527f184edce371c121d112278a4d1239f9d65421fb00e688d7612320bb5f66e7409d8385017f1a7a0f4ef55687795fa98c4585fc66e26ddea1f6e161a837ef4a4f1ae9c8808b8152825194604087019586526108ce8a8a8a8a6117fd565b7f159c9e6c6ad47c3114dd4bdc88dc34359cae49de8306c8f56c9ba9d56429755e87527f261a0e3bad2b8c7b4057a7708b68ddaa9684eaa9d458088e7a2fe7813e08d06082528301518552610925898989896117fd565b7f1243c2c01b1a238bd0937eed8a5eb5f962eae681000356540ae1cfb36e1e308b86527e816901d1be816971e5f7e84b32a92c58a9f8971ff921adc37884d47db225e281526040830151855261097d898989896117fd565b7f25ebb06beeca9f8b08c511a86423c8f8705f115fe942220b9f8e30d54b016e0686527f1a598b2d183a54a959959d562879ae4b48df2203151d223973543b7d9584c59e8152606083015185526109d6898989896117fd565b7f0ede030d344e453627bd0d2e849cd89947ed04b1825b0d4f7d6a8bf8b6ca5bae86527f01e7b010c4ab8cfc4791d0886bb39e1e8785b51a2a2165514d1fe026b8de24ad815260808301518552610a2f898989896117fd565b7f1f81ffd062f9644e01e392d39b8de8e030afd731c770e58f96050ed1d36d553286527f0ac85a9509ed9a6e21c1e933b10794303a4b77d9fd1fb036e966fef320cb4dad815260a08301518552610a88898989896117fd565b7f2d46add97dc6a65ee2f3dd3ea61060bf9fd510929a701e4cec0913f8ab1cafe386527f1e8e47b54e79299b14dfe604a3c8ce10ee4cb2d09be71e23628c13888c29a254815260c08301518552610ae1898989896117fd565b7f146a85b7d5644a318ee8d2a44d995f580695fb92ade1dc0bbfc84cb4010ac1c386527f1d3734ea6fc621a0710ba60b024e87e0442525b2c77aff46634f6c5c0035e073815260e08301518552610b3a898989896117fd565b7f12664f87a4a89be5df17d8f4cb797e102a083e505835ddfa68bfff968ded011f86527f062c4ef046ea89c77def686012d175292e6cca3adb8dc9e0234bbbe4384e6b1881526101008301518552610b94898989896117fd565b7f22ecd52df7b85f6392fa550619d401fedebde3332b7c8857f3b26caac17f4b7986527f25696ccea69b88905a95af209f6daa0d638426ae494b1b5800d1bb32aa271c3c81526101208301518552610bee898989896117fd565b7ee5e920eb2bd31bc7480c75e93c11a2a8f421af3e287f37a87453a96b4dd6c086527f2b82a4685f51b881675e3e4958242a11585b1aa2211bef1ba101ade59d484fcb81526101408301518552610c47898989896117fd565b7f1bb112783a4bd8e95decec6012a6b0c06d076f07806241e64d5bb279055ce2e386527f03aacbe5b76da6ffe5a38b5f74bb1defdf60afde8f7bed77c2103c7d6d285c4581526101608301518552610ca1898989896117fd565b7e89a49d6c462af5737f4f74e89ee3fd5fcda9b129a6885f4402b7191ce06ecc86527f03670d5e8e16d0f9e9351a53b707cddd87f0df01f7e71f8ef942b35a63c54d8081526101808301518552610cfa898989896117fd565b7f148203b1cf918d850c2e4eb482623111d69a0d149273d88bd472cc2dc667788886527f09e3685cddd844c117894cf1560ee45625a29890eb111e539598ac6c8510b26f81526101a08301518552610d54898989896117fd565b7f038bda99e81e5aee528c18e38ab4a8806508a531d1f22f6618919bffb81f59f686527f14ce622b4ba47284ea8c421b00498ea220fb6f88026998098cbb21d38d8e0b0a81526101c08301518552610dae898989896117fd565b7f0c1641e14f8c4509f0c675448654f877363c981ca8c18363b549cfb115737ea386527f0b8cc9ed761dfbad8b821e125c571ba83a857c405c10df3bdeb64fc9b3248e2a81526101e08301518552610e08898989896117fd565b7f14c11b03d9e6d4e5c71174b991b0b1b63f5a8539d4449e10a08275be0454646586527f014e6374d23fc81a10b61fd108e7b0e59003f8d3ed87edd40a722aafbc1c263581526102008301518552610e62898989896117fd565b7f2a195965e3a4ebe550289af22bb4c1118e21dc2c74be94ad6455e8f6eb70c8c486527f12024e0b3b82946c93024f8e1159da64dc1d3c72c49929836a5f2577d5a0ed6881526102208301518552610ebc898989896117fd565b7f1f679e8223e56ae364c7fe5b1eb44e3b3c66cbf45645c364d308c505539047c986527f2c32d125628fae7d840d3f28e83a7fa88112f60ee52f5f86fba53c08e474ff6981526102408301518552610f16898989896117fd565b7f1b73d0bdb2a03d112f31e25b60799a767fb82a9d6418db824c632ffba80be47d86527f21b9840d5d347552e43ea2e7bd19560353b633a278b602c6074025ddb9f63a2181526102608301518552610f70898989896117fd565b7f16b45c0468819f85893ab1c4877c4fe4f49c146b976419b48aa07a0f6cbfb2e286527f06b26b18879ce8c03298c0302e22e9be3407e4b50aa15e3153b0eea99ce0241481526102808301518552610fca898989896117fd565b7f1fe6017431f3e3861894f3e1871a3627f3fc61c832f3e951b3e55c86c4b9615886527f2236d30c0a8738dcac59bc76c975d2ac9c9f1347adfc85c977fc196f4a963f5e81526102a08301518552611024898989896117fd565b7f2fd10aed9958e2d8257e4a70a742fa4774402f36a25babac21c8e0b5f661c12886527f0f8a616d021292af83c9631904c885899f5fe78a5489c14462c97f8472ce5bbf81526102c0830151855261107e898989896117fd565b7f280a41e1efcd026c87f851e0180aad3ac57df1f93286a57f53c400268d8752c386527f1633532825ddc9c0da04246a44b706dcfd57ae2f4c1bb69738ff8433d5b2a8bb81526102e083015185526110d8898989896117fd565b7e027cbc9193ee97f7eec57e57cb0ff7347cc0b2586a1637f4ce954bea3ff97186527f1c897b0f8a07cbc7a4b69597227129b4a12ba68e1926ecc7a45a4ec4bdc5bf0781526103008301518552611131898989896117fd565b7f1b1da35eea8e3139d38e9db84386f59853a18040d1a2216b74679e08c191a01c86527f0bfab57806284de52685f6dda04330043efee9399c75b78e48b5b1d7cb80038e8152610320830151855261118b898989896117fd565b7f012154f85b76ea46de9dc3f61d7c053aa9a583e3e2e57d6a076db599b1326a2f86527f1d25427c48b7647c1efe27b5e7da3240ee333d288ebabeb8b45e30c113c6474e815261034083015185526111e5898989896117fd565b7f111fb275c27d543c507e0c685728727d2344f736a345419dd396d083296211ed86527f01c636dbb603223ca61aee282e75c2ee554f6639e813f990cc7e045128e9ba6e8152610360830151855261123f898989896117fd565b7f0e3ef51ec2992fd4fd4e08d2fc6c02cb6586ec574edec92b74583bd38cc15cae86527f0eeb3ece8b4b83ce8946832b6dd7f35204669e47ccd8b1fc31cba6c71808f68981526103808301518552611299898989896117fd565b7f0a7e2bd7bb8aeb9e84739db84898a9115aa023c279d2df4536366e445e618b1e86527f0141bc992ed56ad3af847ed62afe254e174e6df8efbc36cf3314adad1244b42381526103a083015185526112f3898989896117fd565b7f0f2ecde94b061c256edb823ac557ce52f907f612791c85a2e66fb888ce8a417b86527f1c09474255a3b4c33f9452b362e6352c65acb1406a454a1d4b212538d529ed9a81526103c0830151855261134d898989896117fd565b7f1a2a4a634641112a1b940ba3b089193dfde76611dc7a7c29538a0f93bbdaf83086527f1752d5a3839dedbad8cd819b7b86a82982d1c7663453d236c559353725c2901381526103e083015185526113a7898989896117fd565b7f2e9b2e2e4921cd57e24d3215d41cfa43545b0fa8907380f2b4eca856b242ace086527f192ca02d2e86b7636626a919c871396ba8108cc7f2358ded277a32ed4ba10d4b81526104008301518552611401898989896117fd565b7f1266b7cbb61c28d580a6aa8e4a6b3455bf5c925fe9321c0b0afd01596c4a950586527f0f945b9129e6749912477338802d286af2f6aca2e71cdf199f387d570acf5f108152610420830151855261145b898989896117fd565b7f1be6c06af2b7182fc509c0c72a3874173c0e9b8ef208c89a6ffcbe343f74440a865252015190526117fd565b5115908115916114dd575b506114c4575b50506114a5905161184e565b9451908451908501519160606040870151955196015196611a67565b90565b519193506114a5916114d591611921565b92905f611499565b90505115155f611493565b806115217f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000161151a611526948e6105ef565b5110610605565b6105dc565b6107ae565b611533610100565b9061153c6100db565b5f81525f6020820152825261154f6102cf565b602083015261155c6102cf565b60408301526115696102cf565b6060830152565b61157861152b565b906115816100db565b7f3059a4f6581fbdcc0d5e847306a4862568fbf62d22d49958fc3902e4bc379ef381526020907f1dd7c04824ae7a26cbb4bb2e8e647030df4de01e51764344c9da30b1ac0317c88282015283526115d66100db565b7f0f98f7aa65c680ca7cd4b7f95afc5f6827e95398c599befdd9a9eca741e46b6181527f0a6f18525c7167edf6945aa57ecf8e0dde824c50912f1fe71574e86908579b68828201526116266100db565b7f1d1dacfe7971320b875aa2dbcdafb33d4141ca0b0435904e1e2ead83b600d26b81527f1e0cdae1aa15580307c121c8518d1b513fb8bfc62718065a60e88eee79d0288d838201526116766100db565b91825282820152818401526116896100db565b7f1887c867c4428fd8c7157ac7f5e81a19271ea37ac336aa87203e65bd77cdcad281527f10d21089c03935120870563d17d271a9165f3d541cf4b41a450b8c3741e84423828201526116d96100db565b7f105ceb8102cb4bd76c903c3f045988d7407aab02a26e2b0ac08c58af1edec5a481527f0b32e41ba74a1a65c885129ca2c3c87475d584c75ebc553cb79d4468de6483a8838201526117296100db565b91825282820152604084015261173d6100db565b907f04e9e06a4684b3f9cfea22a0b5d19239c957ba0b12a17f2d9dc9d1e63f9ae49482527f0ef0df2626365d3222024595b1cd400614d2db4a442bb59f5ab585b4717528f88183015261178e6100db565b907f239dd78f7b5dba6a6d81f994b3060a73e4d7602aeb8909fd9785a0f1e04367e382527f178fb89664e86e6758f974887c6d9d19cd52b518c16d799e75556a7ae9a25829818301526117df6100db565b9283528201526060830152565b9060048110156106005760051b0190565b90929160608460806107cf19946007865a01fa1561184c57600660c0926020606096865185528187015182860152805160408601520151868401525a01fa801561184c5761184a906118d5565b565bfe5b5f60206118596100db565b82815201528051908115806118c9575b156118845750506118786100db565b5f81525f602082015290565b602001517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479081900681039081116105ea576118be6100db565b918252602082015290565b50602081015115611869565b156118dc57565b60405162461bcd60e51b815260206004820152601260248201527f70616972696e672d6164642d6661696c656400000000000000000000000000006044820152606490fd5b60609092919260c06119316100db565b915f83525f60208401526020839681611948610100565b936080368637805185520151828401528051604084015201518482015260066107cf195a01fa801561184c5761184a906118d5565b60405190610320820182811067ffffffffffffffff8211176100fb5760405260188252610300366020840137565b906006820291808304600614901517156105ea57565b90600182018092116105ea57565b90600282018092116105ea57565b90600382018092116105ea57565b90600482018092116105ea57565b90600582018092116105ea57565b80518210156106005760209160051b010190565b15611a2257565b60405162461bcd60e51b815260206004820152601560248201527f70616972696e672d6f70636f64652d6661696c656400000000000000000000006044820152606490fd5b9491959692909396611a77610100565b95865260209788978888015260408701526060860152611a95610100565b9384528584015260408301526060820152611aae61197d565b915f5b60048110611ae957505050610300611ac76102aa565b9384920160086107cf195a01fa801561184c57611ae390611a1b565b51151590565b611bbb9192939450611afa816119ab565b611b0482856117ec565b5151611b108288611a07565b5286611b1c83866117ec565b510151611b31611b2b836119c1565b88611a07565b52611b3c82866117ec565b515151611b4b611b2b836119cf565b52611b61611b5983876117ec565b515160200190565b51611b6e611b2b836119dd565b5286611b7a83876117ec565b51015151611b8a611b2b836119eb565b52611bb5611baf611ba889611b9f868a6117ec565b51015160200190565b51926119f9565b87611a07565b526105dc565b9084939291611ab156fea26469706673582212205b768a5c4be76785e2e101930ea12f80a6d2ae1435ed5b2bf4ddf2a5732ce31064736f6c63430008140033",
}

// BlsSigVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BlsSigVerifierMetaData.ABI instead.
var BlsSigVerifierABI = BlsSigVerifierMetaData.ABI

// BlsSigVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlsSigVerifierMetaData.Bin instead.
var BlsSigVerifierBin = BlsSigVerifierMetaData.Bin

// DeployBlsSigVerifier deploys a new Ethereum contract, binding an instance of BlsSigVerifier to it.
func DeployBlsSigVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlsSigVerifier, error) {
	parsed, err := BlsSigVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlsSigVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlsSigVerifier{BlsSigVerifierCaller: BlsSigVerifierCaller{contract: contract}, BlsSigVerifierTransactor: BlsSigVerifierTransactor{contract: contract}, BlsSigVerifierFilterer: BlsSigVerifierFilterer{contract: contract}}, nil
}

// BlsSigVerifier is an auto generated Go binding around an Ethereum contract.
type BlsSigVerifier struct {
	BlsSigVerifierCaller     // Read-only binding to the contract
	BlsSigVerifierTransactor // Write-only binding to the contract
	BlsSigVerifierFilterer   // Log filterer for contract events
}

// BlsSigVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlsSigVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlsSigVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlsSigVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlsSigVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlsSigVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlsSigVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlsSigVerifierSession struct {
	Contract     *BlsSigVerifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlsSigVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlsSigVerifierCallerSession struct {
	Contract *BlsSigVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BlsSigVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlsSigVerifierTransactorSession struct {
	Contract     *BlsSigVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BlsSigVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlsSigVerifierRaw struct {
	Contract *BlsSigVerifier // Generic contract binding to access the raw methods on
}

// BlsSigVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlsSigVerifierCallerRaw struct {
	Contract *BlsSigVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BlsSigVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlsSigVerifierTransactorRaw struct {
	Contract *BlsSigVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlsSigVerifier creates a new instance of BlsSigVerifier, bound to a specific deployed contract.
func NewBlsSigVerifier(address common.Address, backend bind.ContractBackend) (*BlsSigVerifier, error) {
	contract, err := bindBlsSigVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlsSigVerifier{BlsSigVerifierCaller: BlsSigVerifierCaller{contract: contract}, BlsSigVerifierTransactor: BlsSigVerifierTransactor{contract: contract}, BlsSigVerifierFilterer: BlsSigVerifierFilterer{contract: contract}}, nil
}

// NewBlsSigVerifierCaller creates a new read-only instance of BlsSigVerifier, bound to a specific deployed contract.
func NewBlsSigVerifierCaller(address common.Address, caller bind.ContractCaller) (*BlsSigVerifierCaller, error) {
	contract, err := bindBlsSigVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlsSigVerifierCaller{contract: contract}, nil
}

// NewBlsSigVerifierTransactor creates a new write-only instance of BlsSigVerifier, bound to a specific deployed contract.
func NewBlsSigVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BlsSigVerifierTransactor, error) {
	contract, err := bindBlsSigVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlsSigVerifierTransactor{contract: contract}, nil
}

// NewBlsSigVerifierFilterer creates a new log filterer instance of BlsSigVerifier, bound to a specific deployed contract.
func NewBlsSigVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BlsSigVerifierFilterer, error) {
	contract, err := bindBlsSigVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlsSigVerifierFilterer{contract: contract}, nil
}

// bindBlsSigVerifier binds a generic wrapper to an already deployed contract.
func bindBlsSigVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlsSigVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlsSigVerifier *BlsSigVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlsSigVerifier.Contract.BlsSigVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlsSigVerifier *BlsSigVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlsSigVerifier.Contract.BlsSigVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlsSigVerifier *BlsSigVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlsSigVerifier.Contract.BlsSigVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlsSigVerifier *BlsSigVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlsSigVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlsSigVerifier *BlsSigVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlsSigVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlsSigVerifier *BlsSigVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlsSigVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyBlsSigProof is a free data retrieval call binding the contract method 0x0cc26769.
//
// Solidity: function verifyBlsSigProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[35] input) view returns(bool r)
func (_BlsSigVerifier *BlsSigVerifierCaller) VerifyBlsSigProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [35]*big.Int) (bool, error) {
	var out []interface{}
	err := _BlsSigVerifier.contract.Call(opts, &out, "verifyBlsSigProof", a, b, c, commit, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlsSigProof is a free data retrieval call binding the contract method 0x0cc26769.
//
// Solidity: function verifyBlsSigProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[35] input) view returns(bool r)
func (_BlsSigVerifier *BlsSigVerifierSession) VerifyBlsSigProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [35]*big.Int) (bool, error) {
	return _BlsSigVerifier.Contract.VerifyBlsSigProof(&_BlsSigVerifier.CallOpts, a, b, c, commit, input)
}

// VerifyBlsSigProof is a free data retrieval call binding the contract method 0x0cc26769.
//
// Solidity: function verifyBlsSigProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[35] input) view returns(bool r)
func (_BlsSigVerifier *BlsSigVerifierCallerSession) VerifyBlsSigProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [35]*big.Int) (bool, error) {
	return _BlsSigVerifier.Contract.VerifyBlsSigProof(&_BlsSigVerifier.CallOpts, a, b, c, commit, input)
}

// CommitteeRootMappingVerifierMetaData contains all meta data concerning the CommitteeRootMappingVerifier contract.
var CommitteeRootMappingVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[33]\",\"name\":\"input\",\"type\":\"uint256[33]\"}],\"name\":\"verifyCommitteeRootMappingProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610016576119f1908161001b8239f35b5f80fdfe6040608081526004361015610012575f80fd5b5f803560e01c63ab00dde614610026575f80fd5b346100bb576105203660031901126100bb5761004136610117565b9036606312156100bb576100536100d2565b908160c4913683116100bb57506044905b8282106100a25761009e8661008d878761007d36610158565b90610087366101d6565b92610608565b905190151581529081906020820190565b0390f35b602086916100b03685610196565b815201910190610064565b80fd5b634e487b7160e01b5f52604160045260245ffd5b604051906040820182811067ffffffffffffffff8211176100f257604052565b6100be565b604051906080820182811067ffffffffffffffff8211176100f257604052565b8060231215610154576101286100d2565b90816044918211610154576004905b8282106101445750505090565b8135815260209182019101610137565b5f80fd5b8060e31215610154576101696100d2565b90816101049182116101545760c4905b8282106101865750505090565b8135815260209182019101610179565b9080601f83011215610154576101aa6100d2565b80926040810192831161015457905b8282106101c65750505090565b81358152602091820191016101b9565b8061012312156101545760405190610420820182811067ffffffffffffffff8211176100f2576040528161052491821161015457610104905b82821061021c5750505090565b813581526020918201910161020f565b6102346100f7565b906080368337565b604051906060820182811067ffffffffffffffff8211176100f2576040526060368337565b604051906020820182811067ffffffffffffffff8211176100f2576040526020368337565b61028e6100d2565b906102976100d2565b604036823782526102a66100d2565b60403682376020830152565b6102ba6100f7565b906102c36100d2565b5f9081815281602082015283526102d8610286565b60208401526102e56100d2565b81815281602082015260408401526102fb6100d2565b9080825260208201526060830152565b634e487b7160e01b5f52603260045260245ffd5b1561032657565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d710000000000000000006044820152606490fd5b1561037257565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d710000000000000000006044820152606490fd5b156103be57565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561040a57565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561045657565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d7100000000000000006044820152606490fd5b156104a257565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d7100000000000000006044820152606490fd5b156104ee57565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d710000000000000000006044820152606490fd5b1561053a57565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d710000000000000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b5f1981146105a15760010190565b61057f565b9060218110156105b75760051b0190565b61030b565b156105c357565b60405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c64006044820152606490fd5b90919392936106156102b2565b9180516020809201516106266100d2565b9182528282015283526106376100d2565b938051518552610648815160200190565b5182860152610667826106596100d2565b920180515183525160200190565b51828201526106746100d2565b9485528185015280830193845280825192015161068f6100d2565b92835281830152604083019182526107437f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd476106ce818651511061031f565b6106dd8184875101511061036b565b6106eb8187515151106103b7565b6106fb8184885101515110610403565b6107128161070b88515160200190565b511061044f565b61072b81610724858951015160200190565b511061049b565b61073881855151106104e7565b828451015110610533565b5f5b6021811061139057509461138d949561137161075f611418565b9461136b61076b6100d2565b80955f82528582015f815261077e61022c565b9161078761023c565b907f305c9c1aa4a3294d7d6f331d65dc097fd9b9011350a6065eed3d56ad4d48a5e26107b16100d2565b935f85525f8b8601527f2d704aa2e65d5ad168e2ebdd814a37bf7f58961077a78e2ce2eea371cf886b0e8752527f05595e70b8c63dfb8fe2f8adb49c225ee5e6f783b5736f3221d423194db5585d8252610400898301917e7f03a8f871280b33c0292e973247afd9cddfa419e978678bde69323baece8383527f0e4ddacfcf167969963d2bb01ca92fa86b4bfd92c37917bad4ab760f0279625081519360408601948552610861898989896116a5565b7f2cbe10c7c83f6dbe1bc89736f5dc9a9b91e55be6941a4b99b058acd8001fb00486527f1500351867036612a9db15b6f7d4198993f31565af0610ed3fda8d92ffb5e67181528c83015185526108b9898989896116a5565b7f2ff215cfa1c7c99fc0b382d9d7225a0636ccd9a01be83959da430a3a25d4481f86527f2918236b7c008f70fc98cc3bcc41eb60fb8b85d02b4e83100de7a9a7eed34645815260408301518552610912898989896116a5565b7f0943a399c312616b46deee38e49b364b3449d7bab638b580df78aece26ecedf186527f11c937747adbd7b45585d385c0174241e937dce9fa818ca66d00f550b3f3d12881526060830151855261096b898989896116a5565b7f05dd54c0736f8a8d838a097e3151776cc6f602439ecc7833d223dc6cc69d585186527ee6465472f5b3647daa18848088f62f4fe1e57c401172d39f45b827bcc898f78152608083015185526109c3898989896116a5565b7f0d64a91e0e28a2e96b12a3211ebd9f66d854efbc1bd17fccc28e33993722b96d86527f0b30126ec8f40991e90ef34a78e10acb9848fe65e84547e64346fbaebafc9a98815260a08301518552610a1c898989896116a5565b7f08272a78392bca7c2597f09f39daf6f7129808e6a6b34a49239dbf2a264e4e3b86527f02ac17971af65a980f6ab150b8ebaf573008d90c0af4eeed28d50374e8eed16b815260c08301518552610a75898989896116a5565b7f19e3f8ff1265325376056bd2155edf79762433ec7b24c2196701da40ff2e6b3186527f0b33297cf3ae84083dab64e559bccd29e271c3e7d9ba307b7d623d223d091ead815260e08301518552610ace898989896116a5565b7f01341f47e05793a19ea7b92ff3b84c73f7223d56104d070f246e00eb1db7967986527f052fb71e652150885399cfb863b33f0dc0dfe9b279d0bf29df0fc403810fdb1781526101008301518552610b28898989896116a5565b7f04e2785630031d901e87c72d18c1e526142d7b79ee5b4360f5f9373d385aa1be86527f13fdbcc7f866114bdd363bc99c54b3ea0921b9b5526e46885291a00f8f9feea781526101208301518552610b82898989896116a5565b7f0bc888012014ac70c1692250c46020392de91756724c6c890edbc8d860450b8086527f02ce5d7f6c5c1ec90a48c08531305543ef94f25d0d03f67124280afb560406d481526101408301518552610bdc898989896116a5565b7f0aee169033d6ddce5c5dde351a6ce45eb60dc3018752274e14f68e3f19a5359686527f1481dacc3a815550ede26690b83395cb031c17fb4aca75baa0a74acc4042335181526101608301518552610c36898989896116a5565b7f1c24a404b575734133f0c03bbf71c74c5b03a34f1d0817662c7d9293eaecaad586527f10af37e497d439239dfdab04bce89b1eaeaa12bc5f327129816a7a3841863fdd81526101808301518552610c90898989896116a5565b7f14a256de1a2c7bb25672a1acc5406b90543c8a3b8c7c6e0a1895f019171141a886527f08f0539c9fd5ef363053e6560e4769e20e56dd0a510c30dce8adab0230c5fdb181526101a08301518552610cea898989896116a5565b7f082d61eb34a0a6544527e7db6b9bd16a8f6488733c83bee559ec5378191c93ba86527f144b7ab9c8fd2fb71b51c102820f2b08303be60f9fdd313e68f412d3f027a82081526101c08301518552610d44898989896116a5565b7f22510e5ce22c30374993b2a360ceed3bdc20bd64b8d14cb3baedf76ddbd8062386527f0c0c88dfbf63ebc976d642a63c3d22288c546570b101f0219b3e2f3af5bedcfb81526101e08301518552610d9e898989896116a5565b7f2c71d775cc194f6e13408a12a33cb48babccfee137654d1443371de1d0f30c0c86527f1ba219dea8d4ffd8339c1c10cda690451c10fb5058f36ce7e1407118d871cc8081526102008301518552610df8898989896116a5565b7f14e7d115c5cfbe3c075697f305b8660abf41c5725a40557d3e14c9703aef64c286527f0cbc84b02d09b3f498b122ab5819248195aea678e0a41e744967da6bd8d0ce1181526102208301518552610e52898989896116a5565b7f0332f7d5660e970f229a174367929acabfa2f9fdab763460fd7acebddd944dde86527f15be1ce817121a7c25340b8d9c50a584a179d3dffd489f54311a12f922a6942f81526102408301518552610eac898989896116a5565b7f0257933903a2e91846df829f8084008ddf5fc35dd8d4acdebd426bff0d97e2a386527f17e9653840e81e1a68076e0c5f8c89f61463791e918991df86c86f63dccd939181526102608301518552610f06898989896116a5565b7f106f1170be9c02c979b3d6e1d43737530d6bfc444c16df873400384d39e393be86527f1c2b9f619d809bb543e712ac0ee22cc3c6aa99ae73c97246c9769c9c98713fc581526102808301518552610f60898989896116a5565b7f131eb8c00ed76432c870a74c71365748a51807c021e678beb083b0e7e8b5b61186527f2491a76ab72146d0aeb330815df908ad5dd6cd86201e97d220b63d0d8d0f3ac481526102a08301518552610fba898989896116a5565b7f2561f4abb9fabeeb813dcc6d4d487d8f6e36fdd18805e785cccdf2b0a2ff085786527f2e6269f87539d6b464a25b6bd4522d1e2b78c3918f48e79de55dd640261faa4281526102c08301518552611014898989896116a5565b7f2e485df27f23a93b97e296061758e7dd3d34c4722c6fd7ae249433c55d259adc86527f1614f76a407ac31a5acb91266c2c7f54166ee17112e69f0b5e5fad32dafe5f6581526102e0830151855261106e898989896116a5565b7f153af4e0fe4af748819ca675d8781da95111510763de4bec8abf25bca637703b86527f0b1f5a812c51999ebc7ac97de6ab2409b894ad5fb02d45f0798acc2548344cb9815261030083015185526110c8898989896116a5565b7f2c69e0646f6bfc70dfd02cf64d0b781ab481cafbd190dba9ee603f8160c44dcd86527f126015936956b109beba47938f9808cd9eed7ba5fc4531e6f6f267cadb13f4f881526103208301518552611122898989896116a5565b7f22c6d8b6cf6965d431abf72b985b44f1d0831026a43d6dc8cbcb9ca85ab4a0bc86527f19fd6ce3da2b55331cdd361a63f29e95c52f07bcf9cbe355077e5141bf020d838152610340830151855261117c898989896116a5565b7f0cc996a6a427bcf59dad5cb6a2da92164e142372347a8ca5b1b3b32d8b20a0ef86527f232f08e45f51e15d57617fc960278b4dd236ed78e9014d43950751ea862841f1815261036083015185526111d6898989896116a5565b7f1e03c25c13870ae3e127e009be017aa0f47c7b53fc8636bd519ad68f035aa55c86527f04d8795eee1d4bcb8a4042ed861d91024ec5ce76c4fdac892d0c00479874519981526103808301518552611230898989896116a5565b7f2690cf34bdf3837f3036c3c61e73f94f8026f6a6e9be13695cf81ccebbac7ca886527f0db77d728541f6ab723d2dc8389a97864f9e61349120d35d76c004579c3fc10881526103a0830151855261128a898989896116a5565b7f10f3808b8fb7eb5be9d22fa5e7b4599e94116d5803607ae38bd3c637c10224f286527f06cb766b59e904c5b47098c0f98af9c4444f614cf303f1c939fd24b6bd60cc9c81526103c083015185526112e4898989896116a5565b7f1610f9fbade90d90feada79ae229d67175dee93355fbbf229133ce0d6e75e3a186527f1a3084c2af6e7f823d2044866e00b80af74b5586f26c9685943786f355712ddd81526103e0830151855261133e898989896116a5565b7f1b88a7f08e12e3e28b9b10c0c0cbfd8d7df8b8a8fae1840b33c1cd4d24c8b23c865252015190526116a5565b516116ec565b945190845190850151916060604087015195519601519661185d565b90565b806113c97f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016113c26113ce948b6105a6565b51106105bc565b610593565b610745565b6113db6100f7565b906113e46100d2565b5f81525f602082015282526113f7610286565b6020830152611404610286565b6040830152611411610286565b6060830152565b6114206113d3565b906114296100d2565b7f0a3a3884405b8d1fe46693685f02ba163634fd63d14bf91b6a433825b9ed6b5a81526020907f111fc830b029cfb2c94e450b570cf3be3eac81076213c2dcb1e1059330c605bf82820152835261147e6100d2565b7f23ab779b99bf99c421500d8fe70c4e84fc1ff6eefdf3e92d8f581d046cb3eab781527f1611f26f3d9d6b19c4a418d02f19f6796be688f08507bc59ee5f9862dd46fa5b828201526114ce6100d2565b7f2d34a3d654ca9ea36195f8167d653fa7240b0af8acad4b224aed268f9d8756ff81527f178cda417a663a79267fba64b28caf8fc8484866bfe0f423cb8d3b7da164d7f98382015261151e6100d2565b91825282820152818401526115316100d2565b7f1f682eee4eeb25b38c3bff07fad9aaeb8c1ae87a95472a7819a57fd8b37a6e1581527f0db20bc4434468f4ce7f5888da80c6013c5392645400eee1ddbb77b0696ea1a7828201526115816100d2565b7f01979b2d16e0fb974244f72e399fd4d24be132523f4aeb010c75f26b6452d53c81527f21900fdcdfde4102dbbcd9525e925c0f4ea5317aefc7a1c350753b5c9741ebd5838201526115d16100d2565b9182528282015260408401526115e56100d2565b907f04969a13dd24e7586c1e7e668f9be1cfab2bfb7baf9e48cd94428a55b4cfb89882527f3026f4334a515ea181839681e5a601e08615013a7355b0a0ad1c6ffce279eb16818301526116366100d2565b907f0987e27c310f4a785adc7dfc5324848dc4b1b4957907733a04c889777c88a78582527f13c07cb3a59387f85f315e9b41060f8a993a3c3d22113439d63f9be212afc234818301526116876100d2565b9283528201526060830152565b9060048110156105b75760051b0190565b90929160608460806107cf19946007865a01fa156116ea57600660c0926020606096865185528187015182860152805160408601520151868401525a01fa156116ea57565bfe5b5f60206116f76100d2565b8281520152805190811580611767575b156117225750506117166100d2565b5f81525f602082015290565b602001517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479081900681039081116105a15761175c6100d2565b918252602082015290565b50602081015115611707565b60405190610320820182811067ffffffffffffffff8211176100f25760405260188252610300366020840137565b906006820291808304600614901517156105a157565b90600182018092116105a157565b90600282018092116105a157565b90600382018092116105a157565b90600482018092116105a157565b90600582018092116105a157565b80518210156105b75760209160051b010190565b1561181857565b60405162461bcd60e51b815260206004820152601560248201527f70616972696e672d6f70636f64652d6661696c656400000000000000000000006044820152606490fd5b949195969290939661186d6100f7565b9586526020978897888801526040870152606086015261188b6100f7565b93845285840152604083015260608201526118a4611773565b915f5b600481106118df575050506103006118bd610261565b9384920160086107cf195a01fa80156116ea576118d990611811565b51151590565b6119b191929394506118f0816117a1565b6118fa8285611694565b515161190682886117fd565b52866119128386611694565b510151611927611921836117b7565b886117fd565b526119328286611694565b515151611941611921836117c5565b5261195761194f8387611694565b515160200190565b51611964611921836117d3565b52866119708387611694565b51015151611980611921836117e1565b526119ab6119a561199e89611995868a611694565b51015160200190565b51926117ef565b876117fd565b52610593565b90849392916118a756fea2646970667358221220bf0f6f3901e7ec1963d44db8824ed962ed3cb608cebb41343b26937e468152ff64736f6c63430008140033",
}

// CommitteeRootMappingVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitteeRootMappingVerifierMetaData.ABI instead.
var CommitteeRootMappingVerifierABI = CommitteeRootMappingVerifierMetaData.ABI

// CommitteeRootMappingVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitteeRootMappingVerifierMetaData.Bin instead.
var CommitteeRootMappingVerifierBin = CommitteeRootMappingVerifierMetaData.Bin

// DeployCommitteeRootMappingVerifier deploys a new Ethereum contract, binding an instance of CommitteeRootMappingVerifier to it.
func DeployCommitteeRootMappingVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CommitteeRootMappingVerifier, error) {
	parsed, err := CommitteeRootMappingVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitteeRootMappingVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitteeRootMappingVerifier{CommitteeRootMappingVerifierCaller: CommitteeRootMappingVerifierCaller{contract: contract}, CommitteeRootMappingVerifierTransactor: CommitteeRootMappingVerifierTransactor{contract: contract}, CommitteeRootMappingVerifierFilterer: CommitteeRootMappingVerifierFilterer{contract: contract}}, nil
}

// CommitteeRootMappingVerifier is an auto generated Go binding around an Ethereum contract.
type CommitteeRootMappingVerifier struct {
	CommitteeRootMappingVerifierCaller     // Read-only binding to the contract
	CommitteeRootMappingVerifierTransactor // Write-only binding to the contract
	CommitteeRootMappingVerifierFilterer   // Log filterer for contract events
}

// CommitteeRootMappingVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitteeRootMappingVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeRootMappingVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitteeRootMappingVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeRootMappingVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitteeRootMappingVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeRootMappingVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitteeRootMappingVerifierSession struct {
	Contract     *CommitteeRootMappingVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CommitteeRootMappingVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitteeRootMappingVerifierCallerSession struct {
	Contract *CommitteeRootMappingVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CommitteeRootMappingVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitteeRootMappingVerifierTransactorSession struct {
	Contract     *CommitteeRootMappingVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CommitteeRootMappingVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitteeRootMappingVerifierRaw struct {
	Contract *CommitteeRootMappingVerifier // Generic contract binding to access the raw methods on
}

// CommitteeRootMappingVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitteeRootMappingVerifierCallerRaw struct {
	Contract *CommitteeRootMappingVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// CommitteeRootMappingVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitteeRootMappingVerifierTransactorRaw struct {
	Contract *CommitteeRootMappingVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitteeRootMappingVerifier creates a new instance of CommitteeRootMappingVerifier, bound to a specific deployed contract.
func NewCommitteeRootMappingVerifier(address common.Address, backend bind.ContractBackend) (*CommitteeRootMappingVerifier, error) {
	contract, err := bindCommitteeRootMappingVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitteeRootMappingVerifier{CommitteeRootMappingVerifierCaller: CommitteeRootMappingVerifierCaller{contract: contract}, CommitteeRootMappingVerifierTransactor: CommitteeRootMappingVerifierTransactor{contract: contract}, CommitteeRootMappingVerifierFilterer: CommitteeRootMappingVerifierFilterer{contract: contract}}, nil
}

// NewCommitteeRootMappingVerifierCaller creates a new read-only instance of CommitteeRootMappingVerifier, bound to a specific deployed contract.
func NewCommitteeRootMappingVerifierCaller(address common.Address, caller bind.ContractCaller) (*CommitteeRootMappingVerifierCaller, error) {
	contract, err := bindCommitteeRootMappingVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeRootMappingVerifierCaller{contract: contract}, nil
}

// NewCommitteeRootMappingVerifierTransactor creates a new write-only instance of CommitteeRootMappingVerifier, bound to a specific deployed contract.
func NewCommitteeRootMappingVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitteeRootMappingVerifierTransactor, error) {
	contract, err := bindCommitteeRootMappingVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeRootMappingVerifierTransactor{contract: contract}, nil
}

// NewCommitteeRootMappingVerifierFilterer creates a new log filterer instance of CommitteeRootMappingVerifier, bound to a specific deployed contract.
func NewCommitteeRootMappingVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitteeRootMappingVerifierFilterer, error) {
	contract, err := bindCommitteeRootMappingVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitteeRootMappingVerifierFilterer{contract: contract}, nil
}

// bindCommitteeRootMappingVerifier binds a generic wrapper to an already deployed contract.
func bindCommitteeRootMappingVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitteeRootMappingVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeRootMappingVerifier.Contract.CommitteeRootMappingVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeRootMappingVerifier.Contract.CommitteeRootMappingVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeRootMappingVerifier.Contract.CommitteeRootMappingVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeRootMappingVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeRootMappingVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeRootMappingVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyCommitteeRootMappingProof is a free data retrieval call binding the contract method 0xab00dde6.
//
// Solidity: function verifyCommitteeRootMappingProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[33] input) view returns(bool r)
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierCaller) VerifyCommitteeRootMappingProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [33]*big.Int) (bool, error) {
	var out []interface{}
	err := _CommitteeRootMappingVerifier.contract.Call(opts, &out, "verifyCommitteeRootMappingProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCommitteeRootMappingProof is a free data retrieval call binding the contract method 0xab00dde6.
//
// Solidity: function verifyCommitteeRootMappingProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[33] input) view returns(bool r)
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierSession) VerifyCommitteeRootMappingProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [33]*big.Int) (bool, error) {
	return _CommitteeRootMappingVerifier.Contract.VerifyCommitteeRootMappingProof(&_CommitteeRootMappingVerifier.CallOpts, a, b, c, input)
}

// VerifyCommitteeRootMappingProof is a free data retrieval call binding the contract method 0xab00dde6.
//
// Solidity: function verifyCommitteeRootMappingProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[33] input) view returns(bool r)
func (_CommitteeRootMappingVerifier *CommitteeRootMappingVerifierCallerSession) VerifyCommitteeRootMappingProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [33]*big.Int) (bool, error) {
	return _CommitteeRootMappingVerifier.Contract.VerifyCommitteeRootMappingProof(&_CommitteeRootMappingVerifier.CallOpts, a, b, c, input)
}

// CommonMetaData contains all meta data concerning the Common contract.
var CommonMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea2646970667358221220deebfb740ceca6a43cf1cc2a2a4a4f9160fcbc359b0237b7c14a13f0983fbb4764736f6c63430008140033",
}

// CommonABI is the input ABI used to generate the binding from.
// Deprecated: Use CommonMetaData.ABI instead.
var CommonABI = CommonMetaData.ABI

// CommonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommonMetaData.Bin instead.
var CommonBin = CommonMetaData.Bin

// DeployCommon deploys a new Ethereum contract, binding an instance of Common to it.
func DeployCommon(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Common, error) {
	parsed, err := CommonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommonBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Common{CommonCaller: CommonCaller{contract: contract}, CommonTransactor: CommonTransactor{contract: contract}, CommonFilterer: CommonFilterer{contract: contract}}, nil
}

// Common is an auto generated Go binding around an Ethereum contract.
type Common struct {
	CommonCaller     // Read-only binding to the contract
	CommonTransactor // Write-only binding to the contract
	CommonFilterer   // Log filterer for contract events
}

// CommonCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommonSession struct {
	Contract     *Common           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommonCallerSession struct {
	Contract *CommonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CommonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommonTransactorSession struct {
	Contract     *CommonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommonRaw struct {
	Contract *Common // Generic contract binding to access the raw methods on
}

// CommonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommonCallerRaw struct {
	Contract *CommonCaller // Generic read-only contract binding to access the raw methods on
}

// CommonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommonTransactorRaw struct {
	Contract *CommonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommon creates a new instance of Common, bound to a specific deployed contract.
func NewCommon(address common.Address, backend bind.ContractBackend) (*Common, error) {
	contract, err := bindCommon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Common{CommonCaller: CommonCaller{contract: contract}, CommonTransactor: CommonTransactor{contract: contract}, CommonFilterer: CommonFilterer{contract: contract}}, nil
}

// NewCommonCaller creates a new read-only instance of Common, bound to a specific deployed contract.
func NewCommonCaller(address common.Address, caller bind.ContractCaller) (*CommonCaller, error) {
	contract, err := bindCommon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommonCaller{contract: contract}, nil
}

// NewCommonTransactor creates a new write-only instance of Common, bound to a specific deployed contract.
func NewCommonTransactor(address common.Address, transactor bind.ContractTransactor) (*CommonTransactor, error) {
	contract, err := bindCommon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommonTransactor{contract: contract}, nil
}

// NewCommonFilterer creates a new log filterer instance of Common, bound to a specific deployed contract.
func NewCommonFilterer(address common.Address, filterer bind.ContractFilterer) (*CommonFilterer, error) {
	contract, err := bindCommon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommonFilterer{contract: contract}, nil
}

// bindCommon binds a generic wrapper to an already deployed contract.
func bindCommon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Common *CommonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Common.Contract.CommonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Common *CommonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.Contract.CommonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Common *CommonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Common.Contract.CommonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Common *CommonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Common.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Common *CommonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Common *CommonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Common.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// EthereumLightClientMetaData contains all meta data concerning the EthereumLightClient contract.
var EthereumLightClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"genesisValidatorsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"_forkEpochs\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"_forkVersions\",\"type\":\"bytes4[]\"},{\"internalType\":\"uint64\",\"name\":\"_finalizedSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteePoseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_zkVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionStateRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalityUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"forkVersion\",\"type\":\"bytes4\"}],\"name\":\"ForkVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionStateRoot\",\"type\":\"bytes32\"}],\"name\":\"OptimisticUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sszRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"}],\"name\":\"SyncCommitteeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bestValidUpdate\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"attestedHeader\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"finalizedHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"nextSyncCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextSyncCommitteePoseidonRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"nextSyncCommitteeRootMappingProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"forkVersion\",\"type\":\"bytes4\"}],\"name\":\"computeDomain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"computeSigningRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSyncCommitteePoseidonRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSyncCommitteeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedExecutionStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forkEpochs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forkVersions\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestFinalizedSlotAndCommitteeRoots\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"currentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSyncCommitteePoseidonRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSyncCommitteeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticExecutionStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processLightClientForceUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"attestedHeader\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"finalizedHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalityBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"nextSyncCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nextSyncCommitteeBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"nextSyncCommitteePoseidonRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"nextSyncCommitteeRootMappingProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"}],\"internalType\":\"structLightClientUpdate\",\"name\":\"update\",\"type\":\"tuple\"}],\"name\":\"processLightClientUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes4\",\"name\":\"forkVersion\",\"type\":\"bytes4\"}],\"name\":\"updateForkVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"}],\"name\":\"verifyCommitteeSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkVerifier\",\"outputs\":[{\"internalType\":\"contractIBeaconVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234620004b55762004c95803803806200001d81620004f6565b92833981019061010081830312620004b55780519160208083015192604081015160018060401b0390818111620004b55782019284601f85011215620004b557835192620000756200006f856200051c565b620004f6565b9482868681520183819660051b83010191888311620004b55784809101915b838310620004db57505050506060810151838111620004b557810186601f82011215620004b557805196620000cd6200006f896200051c565b9184838a8152019085829a60051b820101928311620004b5578501905b828210620004b957505050620001036080830162000534565b9260a08301519560c08401519760e0809501519a60018060a01b039c8d8d16809d03620004b55760805260a0525198868a11620004a1576801000000000000000091828b11620004a157603c548b603c558b81811062000445575b505099935f809b603c825285822060029780891c92845b848110620003ea57506003198216909103908162000393575b505050505051918783116200037f5782116200036b57603d5482603d5580831062000310575b5091603d8a52808a20928260031c948b5b868110620002c357506007198416909303928362000262575b505050505050501660018060401b031984541617835560045560055560018060a01b03199182603e541617603e55603f54913390831617603f557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0604051933393169180a3614733908162000562823960805181614397015260a0518161445b0152f35b958b93929193968c915b848310620002895750505050505001555f808080808080620001de565b90919293949784620002b56001928b51851c9087871b60031b9163ffffffff809116831b921b19161790565b99019594930191906200026c565b8c8d5b8560088210620002df57505086820155600101620001c5565b62000306869360019397518d1c9088881b60031b9163ffffffff809116831b921b19161790565b92019401620002c6565b6200034190603d8c52848c20600780860160031c820192601c878a1b168062000348575b500160031c019062000549565b5f620001b4565b62000364905f198601908154905f199060200360031b1c169055565b5f62000334565b634e487b7160e01b8a52604160045260248afd5b634e487b7160e01b8b52604160045260248bfd5b9084898e5b838310620003b257505050505001558a5f8080806200018e565b8497620003dd9160019495965116908560031b60031b9160018060401b03809116831b921b19161790565b9601929101898e62000398565b90919293948d8a82915b600483106200041357505050848201558f949392919060010162000175565b85519190950194166001600160401b03908116600683901b90811b91901b1990921691909117906001018a8f620003f4565b6200047691603c5f526003875f2091601882850160021c840194831b16806200047e575b500160021c019062000549565b5f8b6200015e565b6200049a905f198601908154905f199060200360031b1c169055565b5f62000469565b634e487b7160e01b5f52604160045260245ffd5b5f80fd5b81516001600160e01b031981168103620004b5578152908501908501620000ea565b8190620004e88462000534565b815201910190849062000094565b6040519190601f01601f191682016001600160401b03811183821017620004a157604052565b6001600160401b038111620004a15760051b60200190565b51906001600160401b0382168203620004b557565b81811062000555575050565b5f81556001016200054956fe60806040526004361015610011575f80fd5b5f3560e01c8063031523dd146114d957806312420766146114aa57806339536c8f1461148d5780633cf5ea9e146114675780633dd3f4aa1461127d57806343a6c5a61461124d57806365e700de1461123057806367b49cc714611213578063715018a6146111a9578063751f7f15146111365780638da5cb5b14611110578063a1a9ad5514610c40578063a4059e0714610c23578063aae3913b14610beb578063ab556e9f14610aa0578063ba67ee4814610352578063baa94ea21461030c578063bcbaf770146102cb578063c5190436146102ae578063d180236914610289578063d6df096d14610263578063e153d7991461022d578063e1861b08146102105763f2fde38b14610121575f80fd5b3461020c57602036600319011261020c576004356001600160a01b0380821680920361020c57603f5490811690610159338314612f01565b82156101a15773ffffffffffffffffffffffffffffffffffffffff19168217603f557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b5f80fd5b3461020c575f36600319011261020c576020600754604051908152f35b3461020c575f36600319011261020c5760606001600160401b035f54166004546006549060405192835260208301526040820152f35b3461020c575f36600319011261020c5760206001600160a01b03603e5416604051908152f35b3461020c575f36600319011261020c5760206001600160401b035f5416604051908152f35b3461020c575f36600319011261020c576020600154604051908152f35b3461020c57602036600319011261020c57600435603c5481101561020c576001600160401b036102fc602092612ec2565b9190546040519260031b1c168152f35b3461020c57602036600319011261020c57600435603d5481101561020c57610335602091612e6f565b905460405160039290921b1c60e01b6001600160e01b0319168152f35b3461020c575f36600319011261020c5760405161036e81612976565b60405161037a81612940565b6001600160401b03600854818116835260401c1660208201526009546040820152600a546060820152600b54608082015281526040516103b981612976565b6040516103c58161295b565b600c54815260405180816020600d549283815201600d5f527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5925f5b818110610a87575050610416925003826129ac565b602082015281526040516104298161295b565b600e54815260405180816020600f549283815201600f5f527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802925f5b818110610a6e57505061047a925003826129ac565b602082015260208201526040516104908161295b565b601054815260405180816020601154928381520160115f527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68925f5b818110610a555750506104e1925003826129ac565b6020820152604082015260208201526040516104fc8161295b565b601254815260405180816020601354928381520160135f527f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a090925f5b818110610a3c57505061054d925003826129ac565b6020820152604082015260405161056381612976565b60405161056f81612940565b6001600160401b03601454818116835260401c16602082015260155460408201526016546060820152601754608082015281526040516105ae81612976565b6040516105ba8161295b565b601854815260405180602060195491828152019060195f527f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c9695905f5b818110610a26575050508161060c9103826129ac565b6020820152815260405161061f8161295b565b601a548152604051806020601b54918281520190601b5f527f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc1905f5b818110610a1057505050816106719103826129ac565b602082015260208201526040516106878161295b565b601c548152604051806020601d54918281520190601d5f527f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f905f5b8181106109fa57505050816106d99103826129ac565b6020820152604082015260208201526040516106f48161295b565b601e548152604051806020601f54918281520190601f5f527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d807905f5b8181106109e457505050816107469103826129ac565b6020820152604082015260215490602354906040519361076585612991565b60405160245f825b600282106109cd575050506107818161295b565b8552604051926107908461295b565b83926026965f945b600290818710156107c0579060206001926107b28c612cb9565b815201990195019497610798565b5050928686602086015260405180602a905f905b600282106109b6575050506107e88161295b565b604086015260405180602c905f905b6002821061099f5750505061080b8161295b565b60608601526040519161081d83612991565b6001600160401b03602e54168352602f54602084015260305460408401526040519161084883612991565b60405160315f825b60028210610988575050506108648161295b565b8352604051946108738661295b565b85946033985f965b600290818910156108a3579060206001926108958e612cb9565b8152019b019701969961087b565b505088886020880152604051806037905f905b6002821061097257505050926109386109169893604093866108da6109689861295b565b858701526108e6612c84565b6060870152606084019586526109246001600160401b03603b54169886519c8d9c8d610380908181520190612d33565b8c810360208e015290612d33565b98858b015260608a01526080890190612e03565b6001600160401b038151166101c088015260208101516101e0880152015161020086015251610220850190612e03565b6103608301520390f35b60016020819285548152019301910190916108b6565b825481526001928301929190910190602001610850565b8254815260019283019291909101906020016107f7565b8254815260019283019291909101906020016107d4565b82548152600192830192919091019060200161076d565b8254845260209093019260019283019201610730565b82548452602090930192600192830192016106c3565b825484526020909301926001928301920161065b565b82548452602090930192600192830192016105f6565b8454835260019485019486945060209093019201610538565b84548352600194850194869450602090930192016104cc565b8454835260019485019486945060209093019201610465565b8454835260019485019486945060209093019201610401565b3461020c57604036600319011261020c57610ab96129cd565b602435906001600160e01b031982169081830361020c57610ae66001600160a01b03603f54163314612f01565b8115610ba657603c5491600160401b9283811015610b9257806001610b0e9201603c55612ec2565b926001600160401b038091169382549060031b9185831b921b1916179055603d5492831015610b92577f3d992c45d9456d8ebe181b6a66a3721421393afaa297791373e7569c1abcc8af93610b6b84600160409601603d55612e6f565b63ffffffff829392549160031b9260e01c831b921b191617905582519182526020820152a1005b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152601060248201527f62616420666f726b2076657273696f6e000000000000000000000000000000006044820152606490fd5b3461020c57602036600319011261020c576004356001600160e01b03198116810361020c57610c1b602091614440565b604051908152f35b3461020c575f36600319011261020c576020600454604051908152f35b3461020c575f36600319011261020c57610c666001600160a01b03603f54163314612f01565b610c6e614395565b6001600160401b0390815f541690620151808201908183116110fc57831611156110b75781600854169081156110725782601454161115610cc4575b610cba610cb5612f4c565b613ffc565b610cc26136af565b005b610ce4906001600160401b03166001600160401b03196014541617601455565b610d188160085460401c1667ffffffffffffffff60401b6014549160401b169067ffffffffffffffff60401b191617601455565b600954601555600a54601655600b54601755600c54601855600d54818111610b9257600160401b90818111610b925760195481601955808210611032575b507fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb58054905f915b838310610ffc5750505050600e54601a55600f54828111610b9257818111610b9257601b5481601b55808210610fbc575b507f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac8028054905f915b838310610f865750505050601054601c55601154828111610b9257818111610b9257601d5481601d55808210610f46575b507f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c688054905f915b838310610f105750505050601254601e55601354918211610b92578111610b9257601f5481601f55808210610ed0575b507f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09090601f5f5281545f905b828210610e9a575050610caa565b60018091940191825494817fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d8070155019092610e8c565b817fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d80791820191015b818110610f055750610e60565b5f8155600101610ef8565b60018091920192835492817f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f0155019190610e30565b817f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f91820191015b818110610f7b5750610e08565b5f8155600101610f6e565b60018091920192835492817f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc10155019190610dd7565b817f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc191820191015b818110610ff15750610daf565b5f8155600101610fe4565b60018091920192835492817f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c96950155019190610d7e565b817f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c969591820191015b8181106110675750610d56565b5f815560010161105a565b60405162461bcd60e51b815260206004820152601460248201527f6e6f20626573742076616c6964207570646174650000000000000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601260248201527f74696d656f7574206e6f742070617373656400000000000000000000000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b3461020c575f36600319011261020c5760206001600160a01b03603f5416604051908152f35b3461020c57366003190160c0811261020c5760a01361020c5760405161115b81612940565b6001600160401b03600435818116810361020c578252602435908116810361020c5781610c1b9160208094015260443560408201526064356060820152608435608082015260a435906144d5565b3461020c575f36600319011261020c575f603f5473ffffffffffffffffffffffffffffffffffffffff196001600160a01b038216916111e9338414612f01565b16603f557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461020c575f36600319011261020c576020600654604051908152f35b3461020c575f36600319011261020c576020600554604051908152f35b3461020c575f36600319011261020c57600354600254604080519283526001600160401b03909116602083015290f35b3461020c5761026036600319011261020c576112976129cd565b60a036602319011261020c576040516112af81612940565b6001600160401b0390602435828116810361020c57815260443592828416840361020c57602093848301526064356040830152608435606083015260a43560808301526101a060c31936011261020c576040519261130c84612991565b60c435908116810361020c57835260e435848401526101043560408401526101403661012319011261020c5760405161134481612991565b36610143121561020c5760405161135a8161295b565b806101649136831161020c5787610124915b84831061145857505050825236610183121561020c5760405161138e8161295b565b806101e49236841161020c57905b888483106114405791505083015236610203121561020c576040516113c08161295b565b806102249236841161020c578890915b84831061143157505050604083015236610243121561020c57604051906113f68261295b565b816102649136831161020c57905b82821061142257505050610cc2955060608201526060840152613b71565b81358152908801908801611404565b823581529181019181016113d0565b60409161144d3685612ba5565b81520191019061139c565b8235815291810191810161136c565b3461020c575f36600319011261020c5760206001600160401b0360025416604051908152f35b3461020c575f36600319011261020c576020600354604051908152f35b3461020c575f36600319011261020c576001545f54604080519283526001600160401b03909116602083015290f35b3461020c5760031960203682011261020c57600435906001600160401b03821161020c576103c090823603011261020c576115146080612924565b80600401356001600160401b03811161020c576115379060043691840101612aa2565b60805260248101356001600160401b03811161020c5761155d9060043691840101612aa2565b60a05260448101356001600160401b03811161020c5761158390600436918401016129f7565b60c052606481013560e05260848101356001600160401b03811161020c576115b190600436918401016129f7565b6101005260a4810135610120526115cb3660c48301612bea565b610140526101a03682900361020319011261020c576001600160401b0361164e916116436103a4604051926115ff84612991565b61160c61020482016129e3565b845261022481013560208501526102448101356040850152611632366102648301612bea565b6060850152610160849052016129e3565b610180525116614375565b611660611659612f4c565b60806141d6565b80801561291d575b156128d85760016001600160401b0360e06080015151161115612893576001600160401b03611695614395565b816080515151169182911611908161287c575b5015612837575f5460805151516006549182159190828061282a575b8061280f575b6001600160401b0382166001600160401b03841611908115612807575b50156127c25760c0515161262e5760a05151606001516125e9575b610100515161235b57505050506060608001511580612350575b1561230b575b611764608051604051906117358261295b565b600a82527f6f7074696d6973746963000000000000000000000000000000000000000000006020830152613e00565b610180516080515161016051611785929091906001600160401b0316613b71565b6118cd575b80806118ac575b611832575b806117ac575b6117a257005b610cba6080613ffc565b5060206080015151516001600160401b035f54166001600160401b038216119081156117d9575b5061179c565b61010051511515915081611825575b81611805575b50806117fb575b816117d3565b50600654156117f5565b90506607ffffffffffff806080515151600d1c1691600d1c1614816117ee565b60c05151151591506117e8565b7f27eafee2f5b2c935fa7666d231360699c41fe21db57034136b1af3c43f13df896080516001600160401b03602082015151519182600355515116806001600160401b031960025416176002556118a460405192839283602090939291936001600160401b0360408201951681520152565b0390a1611796565b506001600160401b036080515151166001600160401b036002541610611791565b608051608081516001600160401b038151166008549067ffffffffffffffff60401b602084015160401b16916fffffffffffffffffffffffffffffffff1916171760085560408101516009556060810151600a550151600b556020810151602081518051600c5501518051906001600160401b038211610b9257600160401b8211610b9257600d5482600d558083106122ca575b50602001600d5f525f5b828110612296575050506020808201518051600e5501518051906001600160401b038211610b9257600160401b8211610b9257600f5482600f55808310612255575b50602001600f5f525f5b82811061222157505050604001518051601055602001518051906001600160401b038211610b9257600160401b8211610b9257601154826011558083106121e0575b5060200160115f525f5b8281106121ac57505050604001518051601255602001518051906001600160401b038211610b9257600160401b8211610b92576013548260135580831061216b575b5060200160135f525f5b8281106121375750505060206080015160808151611a8d6001600160401b038251166001600160401b03166001600160401b03196014541617601455565b611ac76001600160401b0360208301511667ffffffffffffffff60401b6014549160401b169067ffffffffffffffff60401b191617601455565b604081015160155560608101516016550151601755602081015160208151805160185501518051906001600160401b038211610b9257600160401b8211610b9257601954826019558083106120f6575b5060200160195f525f5b8281106120c2575050506020808201518051601a5501518051906001600160401b038211610b9257600160401b8211610b9257601b5482601b55808310612081575b50602001601b5f525f5b82811061204d57505050604001518051601c55602001518051906001600160401b038211610b9257600160401b8211610b9257601d5482601d5580831061200c575b50602001601d5f525f5b828110611fd857505050604001518051601e55602001518051906001600160401b038211610b9257600160401b8211610b9257601f5482601f55808310611f97575b50602001601f5f525f5b828110611f6357505060c051805191506001600160401b038211610b9257600160401b8211610b925760205482602055808310611f22575b5060200160205f525f5b828110611eee57505060e05160215550610100518051906001600160401b038211610b9257600160401b8211610b925760225482602255808310611ead575b5060200160225f525f5b828110611e7957505061012051602355506101405180515f5b60028110611e645750506020810151906026915f905b6002821015611cfd5780515f5b60028110611ce95750506002602060019201940191019092611cbf565b600190602083519301928188015501611ccc565b5050905060408101515f5b60028110611e4f575050606001515f5b60028110611e3a575050606060e0608001516001600160401b038151166001600160401b0319602e541617602e556020810151602f556040810151603055015180515f5b60028110611e255750506020810151906033915f905b6002821015611db05780515f5b60028110611d9c5750506002602060019201940191019092611d72565b600190602083519301928188015501611d7f565b5050905060408101515f5b60028110611e10575050606001515f5b60028110611dfb5750506001600160401b0361010060800151166001600160401b0319603b541617603b5561178a565b60019060208351930192816039015501611dcb565b60019060208351930192816037015501611dbb565b60019060208351930192816031015501611d5c565b6001906020835193019281602c015501611d18565b6001906020835193019281602a015501611d08565b60019060208351930192816024015501611ca9565b60019060208351930192817f61035b26e3e9eee00e0d72fd1ee8ddca6894550dca6916ea2ac6baa90d11e510015501611c90565b7f61035b26e3e9eee00e0d72fd1ee8ddca6894550dca6916ea2ac6baa90d11e510908382015b8183018110611ee3575050611c86565b5f8155600101611ed3565b60019060208351930192817fc97bfaf2f8ee708c303a06d134f5ecd8389ae0432af62dc132a24118292866bb015501611c47565b7fc97bfaf2f8ee708c303a06d134f5ecd8389ae0432af62dc132a24118292866bb908382015b8183018110611f58575050611c3d565b5f8155600101611f48565b60019060208351930192817fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d807015501611c05565b7fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d807908382015b8183018110611fcd575050611bfb565b5f8155600101611fbd565b60019060208351930192817f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f015501611bb9565b7f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f908382015b8183018110612042575050611baf565b5f8155600101612032565b60019060208351930192817f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc1015501611b6d565b7f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc1908382015b81830181106120b7575050611b63565b5f81556001016120a7565b60019060208351930192817f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c9695015501611b21565b7f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c9695908382015b818301811061212c575050611b17565b5f815560010161211c565b60019060208351930192817f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a090015501611a4f565b7f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a090908382015b81830181106121a1575050611a45565b5f8155600101612191565b60019060208351930192817f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68015501611a03565b7f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68908382015b81830181106122165750506119f9565b5f8155600101612206565b60019060208351930192817f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac8020155016119b7565b7f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802908382015b818301811061228b5750506119ad565b5f815560010161227b565b60019060208351930192817fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb501550161196b565b7fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5908382015b8183018110612300575050611961565b5f81556001016122f0565b60405162461bcd60e51b815260206004820152601c60248201527f6e6f206e6578742073796e6320636f6d6d69747465652070726f6f66000000006044820152606490fd5b50610120511561171c565b6607ffffffffffff8091600d1c1691600d1c161490816125e0575b5061258d575b5060e051610100516080515160600151935f9391929091905b83518510156124415760ff85116110fc576001851b1561242d575f6020916001806037891c16148214612415576123d96123ec916123d3898961450d565b516144b8565b8360405192828480945193849201613d99565b8101039060025afa1561240a576124045f51946144ff565b93612395565b6040513d5f823e3d90fd5b6123d96123ec91612426898961450d565b51906144b8565b634e487b7160e01b5f52601260045260245ffd5b925092509203612548576001600160a01b03603e5416602060606080015161018460a0608001519160c060800151946124996040519687958694637a53878160e01b8652600486015260248501526044840190612e03565b5afa90811561240a575f91612519575b506117225760405162461bcd60e51b815260206004820152602a60248201527f626164206e6578742073796e6320636f6d6d697474656520726f6f74206d617060448201527f70696e672070726f6f66000000000000000000000000000000000000000000006064820152608490fd5b61253b915060203d602011612541575b61253381836129ac565b810190613af5565b836124a9565b503d612529565b60405162461bcd60e51b815260206004820152601d60248201527f626164206e6578742073796e6320636f6d6d69747465652070726f6f660000006044820152606490fd5b60e0510361259b578261237c565b60405162461bcd60e51b815260206004820152601760248201527f626164206e6578742073796e6320636f6d6d69747465650000000000000000006044820152606490fd5b90501584612376565b60405162461bcd60e51b815260206004820152600c60248201527f6e6f2066696e2070726f6f6600000000000000000000000000000000000000006044820152606490fd5b60a0515180516001600160401b0316612694575060a0515160600151156117025760405162461bcd60e51b815260206004820152601e60248201527f67656e65736973206865616465722073686f756c6420626520656d70747900006044820152606490fd5b6126a5909693969592949195614521565b9560406080015194606060805151015197965f975b87518910156127275760ff89116110fc576001891b1561242d575f60209160018060698d1c16148214612716576123d96126f8916123d38d8d61450d565b8101039060025afa1561240a576127105f51986144ff565b976126ba565b6123d96126f8916124268d8d61450d565b929598919497509295500361277d5760a051604051612778916127498261295b565b600982527f66696e616c697a656400000000000000000000000000000000000000000000006020830152613e00565b611702565b60405162461bcd60e51b815260206004820152600d60248201527f6261642066696e2070726f6f66000000000000000000000000000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f6261642061747420736c6f74206f7220636f6d6d6974746565000000000000006044820152606490fd5b9050876116e7565b506607ffffffffffff8082600d1c169083600d1c16146116ca565b50610100515115156116c4565b60405162461bcd60e51b815260206004820152600860248201527f62616420736c6f740000000000000000000000000000000000000000000000006044820152606490fd5b60a05151516001600160401b0316109050836116a8565b60405162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682070617274696369706174696f6e00000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601260248201527f71756f72756d206e6f74207265616368656400000000000000000000000000006044820152606490fd5b5081611668565b61012081019081106001600160401b03821117610b9257604052565b60a081019081106001600160401b03821117610b9257604052565b604081019081106001600160401b03821117610b9257604052565b606081019081106001600160401b03821117610b9257604052565b608081019081106001600160401b03821117610b9257604052565b90601f801991011681019081106001600160401b03821117610b9257604052565b600435906001600160401b038216820361020c57565b35906001600160401b038216820361020c57565b9080601f8301121561020c578135906001600160401b038211610b92578160051b60405193602093612a2b858401876129ac565b8552838086019282010192831161020c578301905b828210612a4e575050505090565b81358152908301908301612a40565b919060408382031261020c5760405190612a768261295b565b8193803583526020810135916001600160401b03831161020c57602092612a9d92016129f7565b910152565b91908281039260e0841261020c57604090815192612abf84612976565b60a084961261020c578251612ad381612940565b612adc836129e3565b8152612aea602084016129e3565b602082015283830135848201526060830135606082015260808301356080820152845260a08201356001600160401b039081811161020c57830160608184031261020c57845190612b3a82612976565b803583811161020c5784612b4f918301612a5d565b8252602081013583811161020c5784612b69918301612a5d565b6020830152858101359083821161020c57612b8691859101612a5d565b85820152602086015260c083013590811161020c57612a9d9201612a5d565b9080601f8301121561020c5760405191612bbe8361295b565b82906040810192831161020c57905b828210612bda5750505090565b8135815260209182019101612bcd565b9190916101408184031261020c57604092835191612c0783612991565b8294612c138383612ba5565b845282605f8301121561020c578051612c2b8161295b565b8060c084019185831161020c57838501905b838210612c6b575050928492612c6160609661010094612a9d9760208b0152612ba5565b9087015201612ba5565b60208591612c798985612ba5565b815201910190612c3d565b6040519060395f835b60028210612ca357505050612ca18261295b565b565b6001602081928554815201930191019091612c8d565b60405191905f835b60028210612cd557505050612ca18261295b565b6001602081928554815201930191019091612cc1565b6060906040830190805184526020928380920151946040838201528551809452019301915f5b828110612d1f575050505090565b835185529381019392810192600101612d11565b612dd991608082516001600160401b0380825116845260208201511660208401526040810151604084015260608101516060840152015160808201526040612dc8602084015160e060a0850152612d978151606060e0870152610140860190612ceb565b83612db560208401519260df199384898303016101008a0152612ceb565b9201519085830301610120860152612ceb565b9201519060c0818403910152612ceb565b90565b5f915b60028310612dec57505050565b600190825181526020809101920192019190612ddf565b9190612e10818451612ddc565b60208381015193906040905f908483015b60028310612e5357505050612ca193945081612e4960609261010094015160c0860190612ddc565b0151910190612ddc565b818482612e636001948c51612ddc565b01980192019196612e21565b90603d54821015612eae57603d5f52601c8260031c7fece66cfdbd22e3f37d348a3d8e19074452862cd65fd4b9a11f0336d1ac6d1dc3019260021b1690565b634e487b7160e01b5f52603260045260245ffd5b90603c54821015612eae57603c5f5260188260021c7fc6bb06cb7f92603de181bf256cd16846b93b752a170ff24824098b31aa008a7e019260031b1690565b15612f0857565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6040908151612f5a81612924565b80835193612f6785612976565b805193612f7385612940565b600854916001600160401b0394858416875285602094831c168488015260095482880152600a549660609788820152600b5460808201528852815197612fb889612976565b825198612fc48a61295b565b600c548a528351998a87600d549c8d8152015f809d600d82527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5915b81811061368557505050816130169103826129ac565b87820152815283516130278161295b565b600e5481528451808c89600f54928381520191600f82527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802915b81811061367157505050816130779103826129ac565b8782015286820152835161308a8161295b565b60105481528451808c89601154928381520191601182527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68915b81811061365d57505050816130da9103826129ac565b87820152848201528582015282516130f18161295b565b60125481528351808760135491828152019060138d527f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a090908d5b81811061364957505050816131419103826129ac565b86820152838201528152815161315681612976565b825161316181612940565b876014548181168352851c1686820152601554848201526016548982015260175460808201528152825161319481612976565b835161319f8161295b565b60185481528451808c89601954928381520191601982527f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c9695915b81811061363557505050816131ef9103826129ac565b87820152815283516132008161295b565b601a5481528451808c89601b54928381520191601b82527f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc1915b81811061362157505050816132509103826129ac565b878201528682015283516132638161295b565b601c5481528451808c89601d54928381520191601d82527f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f915b81811061360d57505050816132b39103826129ac565b87820152848201528582015282516132ca8161295b565b601e54815283518087601f54918281520190601f8d527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d807908d5b8181106135f9575050508161331a9103826129ac565b81870152818401528185015281518454808252858a52818601907fc97bfaf2f8ee708c303a06d134f5ecd8389ae0432af62dc132a24118292866bb908b5b8181106135e5575050508161336e9103826129ac565b82820152602154878201528151808560225491828152019060228b527f61035b26e3e9eee00e0d72fd1ee8ddca6894550dca6916ea2ac6baa90d11e510908b5b8181106135d157505050816133c49103826129ac565b608082015260235460a08201528151966133dd88612991565b825160248a825b600282106135bb575050506133f88161295b565b88528251966134068861295b565b8996602695895b600290818b10156134345790896001926134268b612cb9565b81520198019901989661340d565b50509398919750939891945085820152875180602a908b905b600282106135a5575050506134618161295b565b818901528751602c8a825b6002821061358f575050506134808161295b565b8482015260c087015286519561349587612991565b85602e54168752602f5485880152603054888801528751956134b687612991565b885160318b825b60028210613579575050506134d18161295b565b87528851976134df8961295b565b8a976033968a5b600290818c101561350d57908a6001926134ff8c612cb9565b81520199019a0199976134e6565b5050949991965094999196929750828501528051918260379b905b600282106135645750505061010096979899506135448261295b565b83015261354f612c84565b8183015282015260e0840152603b5416910152565b60018381928f548152019d019101909b613528565b82548152600192830192919091019089016134bd565b825481526001928301929190910190880161346c565b825481526001928301929190910190880161344d565b82548152600192830192919091019088016133e4565b8254845292880192600192830192016133ae565b825484529288019260019283019201613358565b82548452928a019260019283019201613304565b82548452928b01926001928301920161329d565b82548452928b01926001928301920161323a565b82548452928b0192600192830192016131d9565b82548452928a01926001928301920161312b565b82548452928b0192600192830192016130c4565b82548452928b019260019283019201613061565b82548452928b019260019283019201613000565b8181106136a4575050565b5f8155600101613699565b5f806008558060095580600a5580600b5580600c55600d5481600d5580613ab3575b5080600e55600f5481600f5580613a71575b50806010556011548160115580613a2f575b508060125560135481601355806139ed575b50806014558060155580601655806017558060185560195481601955806139ab575b5080601a55601b5481601b5580613969575b5080601c55601d5481601d5580613927575b5080601e55601f5481601f55806138e5575b5060205481602055806138a3575b50806021556022548160225580613861575b508060235560245b60268110613856575060265b602a81106138415750602a5b602c81106138365750602c5b602e811061382b575080602e5580602f558060305560315b60338110613820575060335b6037811061380b575060375b60398110613800575060395b603b81106137f55750603b55565b8181556001016137e7565b8181556001016137db565b8061381a600280930182613699565b016137cf565b8181556001016137c3565b8181556001016137ab565b81815560010161379f565b80613850600280930182613699565b01613793565b818155600101613787565b602282527f61035b26e3e9eee00e0d72fd1ee8ddca6894550dca6916ea2ac6baa90d11e510908101905b818110613898575061377f565b82815560010161388b565b602082527fc97bfaf2f8ee708c303a06d134f5ecd8389ae0432af62dc132a24118292866bb908101905b8181106138da575061376d565b8281556001016138cd565b601f82527fa03837a25210ee280c2113ff4b77ca23440b19d4866cca721c801278fd08d807908101905b81811061391c575061375f565b82815560010161390f565b601d82527f6d4407e7be21f808e6509aa9fa9143369579dd7d760fe20a2c09680fc146134f908101905b81811061395e575061374d565b828155600101613951565b601b82527f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc1908101905b8181106139a0575061373b565b828155600101613993565b601982527f944998273e477b495144fb8794c914197f3ccb46be2900f4698fd0ef743c9695908101905b8181106139e25750613729565b8281556001016139d5565b601382527f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a090908101905b818110613a245750613707565b828155600101613a17565b601182527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68908101905b818110613a6657506136f5565b828155600101613a59565b600f82527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802908101905b818110613aa857506136e3565b828155600101613a9b565b600d82527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5908101905b818110613aea57506136d1565b828155600101613add565b9081602091031261020c5751801515810361020c5790565b9060016001600160401b03809316019182116110fc57565b15613b2c57565b60405162461bcd60e51b815260206004820152601160248201527f62616420706f736569646f6e20726f6f740000000000000000000000000000006044820152606490fd5b6607ffffffffffff805f54600d1c169082600d1c169160065415155f14613d32578183148015613d18575b15613cd357602093613bc6613bc16707ffffffffffffff613bcc945b60051c166143cf565b614440565b906144d5565b9103613cb9576101c4600554613be6848601518214613b25565b915b6001600160a01b03603e5416613c3e6001600160401b038751169660606040820151910151906040519889978896630291ab6d60e51b885260048801526024870152604486015260648501526084840190612e03565b5afa90811561240a575f91613c9b575b5015613c5657565b60405162461bcd60e51b815260206004820152601160248201527f62616420626c73207369672070726f6f660000000000000000000000000000006044820152606490fd5b613cb3915060203d81116125415761253381836129ac565b5f613c4e565b6101c4600754613ccd848601518214613b25565b91613be8565b60405162461bcd60e51b815260206004820152601060248201527f6261642073696720706572696f642032000000000000000000000000000000006044820152606490fd5b506001600160401b03613d2a83613b0d565b168314613b9c565b818303613d5457602093613bc6613bc16707ffffffffffffff613bcc94613bb8565b60405162461bcd60e51b815260206004820152601060248201527f6261642073696720706572696f642031000000000000000000000000000000006044820152606490fd5b5f5b838110613daa5750505f910152565b8181015183820152602001613d9b565b15613dc25750565b6044604051809262461bcd60e51b825260206004830152613df28151809281602486015260208686019101613d99565b601f01601f19168101030190fd5b9190916020908181015192604094858301608081519451015185855195015194935f945b8651861015613eaa5760ff86116110fc57600180871b1561242d5788915f916019891c811603613e9957613e5f613e71916123d38a8c61450d565b838d5192828480945193849201613d99565b8101039060025afa15613e8f57613e895f51956144ff565b94613e24565b88513d5f823e3d90fd5b613e5f613e71916124268a8c61450d565b909397929550613f0e9194508851907f626164206578656320726f6f742070726f6f662000000000000000000000000088830152885195613f086034848b8d0199613ef8818484018d613d99565b81010360148101865201846129ac565b14613dba565b5192515184845194015193925f935b8551851015613fa35760ff85116110fc57600180861b1561242d5787915f916012881c811603613f9257613f58613f6a916123d3898b61450d565b838c5192828480945193849201613d99565b8101039060025afa15613f8857613f825f51946144ff565b93613f1d565b87513d5f823e3d90fd5b613f58613f6a91612426898b61450d565b9450603a919350947f626164206578656320737461746520726f6f742070726f6f662000000000000095613fec612ca198613f08959a5198899485015251809285850190613d99565b810103601a8101865201846129ac565b602081019081515151906001600160401b0392838316926607ffffffffffff805f54600d1c1691600d1c16906006548681155f14614163575050810361411e57827fdc7ba15c782b181b9d54a996db4ab8a32182bff2bfb09e4aca8ed9ea9e5380c7916140a661407f60a060606020980151948560065501519283600755613b0d565b91604051938493846040919493926001600160401b03606083019616825260208201520152565b0390a15b5101515151905f5492831681116140c057505050565b807f4d75bcddf849ad697dd4b9e37ec69f14240170e980101fcd9e57bb000527c24893836001556001600160401b031916175f5561411960405192839283602090939291936001600160401b0360408201951681520152565b0390a1565b60405162461bcd60e51b815260206004820152600f60248201527f6d69736d6174636820706572696f6400000000000000000000000000000000006044820152606490fd5b61417260209693949294613b0d565b168114614182575b5050506140aa565b7fdc7ba15c782b181b9d54a996db4ab8a32182bff2bfb09e4aca8ed9ea9e5380c7926004556007546005556141cb61407f60a06060850151948560065501519283600755613b0d565b0390a15f808061417a565b60e08201906001600160401b03918281515116156143315760e082016141ff8482515116614375565b61420c8584515116614375565b1590801591801583150361431b57505080614309575b6142f85761422f8361433f565b6142388661433f565b1515811515036142ef57506040830151511515604086015151151581036142ef5784906142aa575b809151511691515116908181036142a25750508181515151168284515151169081810361429a575050816101008092015116920151161190565b109392505050565b119392505050565b50602083015151516607ffffffffffff90818086515151600d1c1691600d1c1614908060208801515151600d1c169087515151600d1c161481036142ef575083614260565b94505050505090565b909350829150515116915151161090565b50838151511684835151161415614222565b95509550505050508161432c575090565b905090565b5060e0015151161515919050565b60808101515115159081614351575090565b905061010081515151916607ffffffffffff9182910151600d1c1691600d1c161490565b60036001600160401b03809216029081169081036110fc57610400111590565b7f000000000000000000000000000000000000000000000000000000000000000042034281116110fc57600c6001600160401b0391041690565b603d545f1992908381019081116110fc575b6143ea81612ec2565b906001600160401b03809154600393841b1c169085161015614428575080156144145783016143e1565b60245f634e487b7160e01b81526011600452fd5b9250614435919350612e6f565b9054911b1c60e01b90565b5f61449d6020926040518481019163ffffffff60e01b1682527f000000000000000000000000000000000000000000000000000000000000000060408201526040815261448c81612976565b604051928392839251928391613d99565b8101039060025afa1561240a575f5160201c600760f81b1790565b9190604051926020840152604083015260408252612ca182612976565b6144ed6123d96020936144e85f94614521565b6144b8565b8101039060025afa1561240a575f5190565b5f1981146110fc5760010190565b8051821015612eae5760209160051b010190565b6001600160401b039061454f614539838351166146cf565b61454960209485850151166146cf565b906144b8565b918060409261456684519586815194859201613d99565b825f86819760029581010390855afa156146b857845183866145a5614593888601516060870151906144b8565b83895192828480945193849201613d99565b81010390865afa156146c557856145d46145c286938351906144b8565b83885192828480945193849201613d99565b81010390855afa156146b8578285614615608082519401518751848101918252838982015288815261460581612976565b8851928392839251928391613d99565b81010390855afa156146b85784518386614652875183810190838252838a82015289815261464281612976565b8951928392839251928391613d99565b81010390865afa156146c5578561466f6145c286938351906144b8565b81010390855afa156146b8579061469e61468c86938451906144b8565b84865192828480945193849201613d99565b810103915afa156146ae57505190565b51903d90823e3d90fd5b50505051903d90823e3d90fd5b84513d87823e3d90fd5b5f9081905b602082106146e157505090565b600892831b60ff821617921c906146f7906144ff565b906146d456fea264697066735822122025eae12edff26e7a9b24b40adcdb80fe0e0897dc8a934421610b30f9549ef0a164736f6c63430008140033",
}

// EthereumLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumLightClientMetaData.ABI instead.
var EthereumLightClientABI = EthereumLightClientMetaData.ABI

// EthereumLightClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthereumLightClientMetaData.Bin instead.
var EthereumLightClientBin = EthereumLightClientMetaData.Bin

// DeployEthereumLightClient deploys a new Ethereum contract, binding an instance of EthereumLightClient to it.
func DeployEthereumLightClient(auth *bind.TransactOpts, backend bind.ContractBackend, genesisTime *big.Int, genesisValidatorsRoot [32]byte, _forkEpochs []uint64, _forkVersions [][4]byte, _finalizedSlot uint64, syncCommitteeRoot [32]byte, syncCommitteePoseidonRoot [32]byte, _zkVerifier common.Address) (common.Address, *types.Transaction, *EthereumLightClient, error) {
	parsed, err := EthereumLightClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthereumLightClientBin), backend, genesisTime, genesisValidatorsRoot, _forkEpochs, _forkVersions, _finalizedSlot, syncCommitteeRoot, syncCommitteePoseidonRoot, _zkVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthereumLightClient{EthereumLightClientCaller: EthereumLightClientCaller{contract: contract}, EthereumLightClientTransactor: EthereumLightClientTransactor{contract: contract}, EthereumLightClientFilterer: EthereumLightClientFilterer{contract: contract}}, nil
}

// EthereumLightClient is an auto generated Go binding around an Ethereum contract.
type EthereumLightClient struct {
	EthereumLightClientCaller     // Read-only binding to the contract
	EthereumLightClientTransactor // Write-only binding to the contract
	EthereumLightClientFilterer   // Log filterer for contract events
}

// EthereumLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumLightClientSession struct {
	Contract     *EthereumLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthereumLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumLightClientCallerSession struct {
	Contract *EthereumLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EthereumLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumLightClientTransactorSession struct {
	Contract     *EthereumLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EthereumLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumLightClientRaw struct {
	Contract *EthereumLightClient // Generic contract binding to access the raw methods on
}

// EthereumLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumLightClientCallerRaw struct {
	Contract *EthereumLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumLightClientTransactorRaw struct {
	Contract *EthereumLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereumLightClient creates a new instance of EthereumLightClient, bound to a specific deployed contract.
func NewEthereumLightClient(address common.Address, backend bind.ContractBackend) (*EthereumLightClient, error) {
	contract, err := bindEthereumLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthereumLightClient{EthereumLightClientCaller: EthereumLightClientCaller{contract: contract}, EthereumLightClientTransactor: EthereumLightClientTransactor{contract: contract}, EthereumLightClientFilterer: EthereumLightClientFilterer{contract: contract}}, nil
}

// NewEthereumLightClientCaller creates a new read-only instance of EthereumLightClient, bound to a specific deployed contract.
func NewEthereumLightClientCaller(address common.Address, caller bind.ContractCaller) (*EthereumLightClientCaller, error) {
	contract, err := bindEthereumLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientCaller{contract: contract}, nil
}

// NewEthereumLightClientTransactor creates a new write-only instance of EthereumLightClient, bound to a specific deployed contract.
func NewEthereumLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumLightClientTransactor, error) {
	contract, err := bindEthereumLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientTransactor{contract: contract}, nil
}

// NewEthereumLightClientFilterer creates a new log filterer instance of EthereumLightClient, bound to a specific deployed contract.
func NewEthereumLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumLightClientFilterer, error) {
	contract, err := bindEthereumLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientFilterer{contract: contract}, nil
}

// bindEthereumLightClient binds a generic wrapper to an already deployed contract.
func bindEthereumLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthereumLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumLightClient *EthereumLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumLightClient.Contract.EthereumLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumLightClient *EthereumLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.EthereumLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumLightClient *EthereumLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.EthereumLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumLightClient *EthereumLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumLightClient *EthereumLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumLightClient *EthereumLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.contract.Transact(opts, method, params...)
}

// BestValidUpdate is a free data retrieval call binding the contract method 0xba67ee48.
//
// Solidity: function bestValidUpdate() view returns(((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) attestedHeader, ((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) finalizedHeader, bytes32 nextSyncCommitteeRoot, bytes32 nextSyncCommitteePoseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) nextSyncCommitteeRootMappingProof, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate, uint64 signatureSlot)
func (_EthereumLightClient *EthereumLightClientCaller) BestValidUpdate(opts *bind.CallOpts) (struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "bestValidUpdate")

	outstruct := new(struct {
		AttestedHeader                    HeaderWithExecution
		FinalizedHeader                   HeaderWithExecution
		NextSyncCommitteeRoot             [32]byte
		NextSyncCommitteePoseidonRoot     [32]byte
		NextSyncCommitteeRootMappingProof IBeaconVerifierProof
		SyncAggregate                     SyncAggregate
		SignatureSlot                     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestedHeader = *abi.ConvertType(out[0], new(HeaderWithExecution)).(*HeaderWithExecution)
	outstruct.FinalizedHeader = *abi.ConvertType(out[1], new(HeaderWithExecution)).(*HeaderWithExecution)
	outstruct.NextSyncCommitteeRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.NextSyncCommitteePoseidonRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.NextSyncCommitteeRootMappingProof = *abi.ConvertType(out[4], new(IBeaconVerifierProof)).(*IBeaconVerifierProof)
	outstruct.SyncAggregate = *abi.ConvertType(out[5], new(SyncAggregate)).(*SyncAggregate)
	outstruct.SignatureSlot = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// BestValidUpdate is a free data retrieval call binding the contract method 0xba67ee48.
//
// Solidity: function bestValidUpdate() view returns(((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) attestedHeader, ((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) finalizedHeader, bytes32 nextSyncCommitteeRoot, bytes32 nextSyncCommitteePoseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) nextSyncCommitteeRootMappingProof, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate, uint64 signatureSlot)
func (_EthereumLightClient *EthereumLightClientSession) BestValidUpdate() (struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}, error) {
	return _EthereumLightClient.Contract.BestValidUpdate(&_EthereumLightClient.CallOpts)
}

// BestValidUpdate is a free data retrieval call binding the contract method 0xba67ee48.
//
// Solidity: function bestValidUpdate() view returns(((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) attestedHeader, ((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) finalizedHeader, bytes32 nextSyncCommitteeRoot, bytes32 nextSyncCommitteePoseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) nextSyncCommitteeRootMappingProof, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate, uint64 signatureSlot)
func (_EthereumLightClient *EthereumLightClientCallerSession) BestValidUpdate() (struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}, error) {
	return _EthereumLightClient.Contract.BestValidUpdate(&_EthereumLightClient.CallOpts)
}

// ComputeDomain is a free data retrieval call binding the contract method 0xaae3913b.
//
// Solidity: function computeDomain(bytes4 forkVersion) view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) ComputeDomain(opts *bind.CallOpts, forkVersion [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "computeDomain", forkVersion)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeDomain is a free data retrieval call binding the contract method 0xaae3913b.
//
// Solidity: function computeDomain(bytes4 forkVersion) view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) ComputeDomain(forkVersion [4]byte) ([32]byte, error) {
	return _EthereumLightClient.Contract.ComputeDomain(&_EthereumLightClient.CallOpts, forkVersion)
}

// ComputeDomain is a free data retrieval call binding the contract method 0xaae3913b.
//
// Solidity: function computeDomain(bytes4 forkVersion) view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) ComputeDomain(forkVersion [4]byte) ([32]byte, error) {
	return _EthereumLightClient.Contract.ComputeDomain(&_EthereumLightClient.CallOpts, forkVersion)
}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0x751f7f15.
//
// Solidity: function computeSigningRoot((uint64,uint64,bytes32,bytes32,bytes32) header, bytes32 domain) pure returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) ComputeSigningRoot(opts *bind.CallOpts, header BeaconBlockHeader, domain [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "computeSigningRoot", header, domain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0x751f7f15.
//
// Solidity: function computeSigningRoot((uint64,uint64,bytes32,bytes32,bytes32) header, bytes32 domain) pure returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) ComputeSigningRoot(header BeaconBlockHeader, domain [32]byte) ([32]byte, error) {
	return _EthereumLightClient.Contract.ComputeSigningRoot(&_EthereumLightClient.CallOpts, header, domain)
}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0x751f7f15.
//
// Solidity: function computeSigningRoot((uint64,uint64,bytes32,bytes32,bytes32) header, bytes32 domain) pure returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) ComputeSigningRoot(header BeaconBlockHeader, domain [32]byte) ([32]byte, error) {
	return _EthereumLightClient.Contract.ComputeSigningRoot(&_EthereumLightClient.CallOpts, header, domain)
}

// CurrentSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0x65e700de.
//
// Solidity: function currentSyncCommitteePoseidonRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) CurrentSyncCommitteePoseidonRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "currentSyncCommitteePoseidonRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0x65e700de.
//
// Solidity: function currentSyncCommitteePoseidonRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) CurrentSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.CurrentSyncCommitteePoseidonRoot(&_EthereumLightClient.CallOpts)
}

// CurrentSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0x65e700de.
//
// Solidity: function currentSyncCommitteePoseidonRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) CurrentSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.CurrentSyncCommitteePoseidonRoot(&_EthereumLightClient.CallOpts)
}

// CurrentSyncCommitteeRoot is a free data retrieval call binding the contract method 0xa4059e07.
//
// Solidity: function currentSyncCommitteeRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) CurrentSyncCommitteeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "currentSyncCommitteeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentSyncCommitteeRoot is a free data retrieval call binding the contract method 0xa4059e07.
//
// Solidity: function currentSyncCommitteeRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) CurrentSyncCommitteeRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.CurrentSyncCommitteeRoot(&_EthereumLightClient.CallOpts)
}

// CurrentSyncCommitteeRoot is a free data retrieval call binding the contract method 0xa4059e07.
//
// Solidity: function currentSyncCommitteeRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) CurrentSyncCommitteeRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.CurrentSyncCommitteeRoot(&_EthereumLightClient.CallOpts)
}

// FinalizedExecutionStateRoot is a free data retrieval call binding the contract method 0xc5190436.
//
// Solidity: function finalizedExecutionStateRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) FinalizedExecutionStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "finalizedExecutionStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedExecutionStateRoot is a free data retrieval call binding the contract method 0xc5190436.
//
// Solidity: function finalizedExecutionStateRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) FinalizedExecutionStateRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.FinalizedExecutionStateRoot(&_EthereumLightClient.CallOpts)
}

// FinalizedExecutionStateRoot is a free data retrieval call binding the contract method 0xc5190436.
//
// Solidity: function finalizedExecutionStateRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) FinalizedExecutionStateRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.FinalizedExecutionStateRoot(&_EthereumLightClient.CallOpts)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_EthereumLightClient *EthereumLightClientCaller) FinalizedExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "finalizedExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_EthereumLightClient *EthereumLightClientSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _EthereumLightClient.Contract.FinalizedExecutionStateRootAndSlot(&_EthereumLightClient.CallOpts)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_EthereumLightClient *EthereumLightClientCallerSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _EthereumLightClient.Contract.FinalizedExecutionStateRootAndSlot(&_EthereumLightClient.CallOpts)
}

// FinalizedSlot is a free data retrieval call binding the contract method 0xd1802369.
//
// Solidity: function finalizedSlot() view returns(uint64)
func (_EthereumLightClient *EthereumLightClientCaller) FinalizedSlot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "finalizedSlot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FinalizedSlot is a free data retrieval call binding the contract method 0xd1802369.
//
// Solidity: function finalizedSlot() view returns(uint64)
func (_EthereumLightClient *EthereumLightClientSession) FinalizedSlot() (uint64, error) {
	return _EthereumLightClient.Contract.FinalizedSlot(&_EthereumLightClient.CallOpts)
}

// FinalizedSlot is a free data retrieval call binding the contract method 0xd1802369.
//
// Solidity: function finalizedSlot() view returns(uint64)
func (_EthereumLightClient *EthereumLightClientCallerSession) FinalizedSlot() (uint64, error) {
	return _EthereumLightClient.Contract.FinalizedSlot(&_EthereumLightClient.CallOpts)
}

// ForkEpochs is a free data retrieval call binding the contract method 0xbcbaf770.
//
// Solidity: function forkEpochs(uint256 ) view returns(uint64)
func (_EthereumLightClient *EthereumLightClientCaller) ForkEpochs(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "forkEpochs", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkEpochs is a free data retrieval call binding the contract method 0xbcbaf770.
//
// Solidity: function forkEpochs(uint256 ) view returns(uint64)
func (_EthereumLightClient *EthereumLightClientSession) ForkEpochs(arg0 *big.Int) (uint64, error) {
	return _EthereumLightClient.Contract.ForkEpochs(&_EthereumLightClient.CallOpts, arg0)
}

// ForkEpochs is a free data retrieval call binding the contract method 0xbcbaf770.
//
// Solidity: function forkEpochs(uint256 ) view returns(uint64)
func (_EthereumLightClient *EthereumLightClientCallerSession) ForkEpochs(arg0 *big.Int) (uint64, error) {
	return _EthereumLightClient.Contract.ForkEpochs(&_EthereumLightClient.CallOpts, arg0)
}

// ForkVersions is a free data retrieval call binding the contract method 0xbaa94ea2.
//
// Solidity: function forkVersions(uint256 ) view returns(bytes4)
func (_EthereumLightClient *EthereumLightClientCaller) ForkVersions(opts *bind.CallOpts, arg0 *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "forkVersions", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ForkVersions is a free data retrieval call binding the contract method 0xbaa94ea2.
//
// Solidity: function forkVersions(uint256 ) view returns(bytes4)
func (_EthereumLightClient *EthereumLightClientSession) ForkVersions(arg0 *big.Int) ([4]byte, error) {
	return _EthereumLightClient.Contract.ForkVersions(&_EthereumLightClient.CallOpts, arg0)
}

// ForkVersions is a free data retrieval call binding the contract method 0xbaa94ea2.
//
// Solidity: function forkVersions(uint256 ) view returns(bytes4)
func (_EthereumLightClient *EthereumLightClientCallerSession) ForkVersions(arg0 *big.Int) ([4]byte, error) {
	return _EthereumLightClient.Contract.ForkVersions(&_EthereumLightClient.CallOpts, arg0)
}

// LatestFinalizedSlotAndCommitteeRoots is a free data retrieval call binding the contract method 0xe153d799.
//
// Solidity: function latestFinalizedSlotAndCommitteeRoots() view returns(uint64 slot, bytes32 currentRoot, bytes32 nextRoot)
func (_EthereumLightClient *EthereumLightClientCaller) LatestFinalizedSlotAndCommitteeRoots(opts *bind.CallOpts) (struct {
	Slot        uint64
	CurrentRoot [32]byte
	NextRoot    [32]byte
}, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "latestFinalizedSlotAndCommitteeRoots")

	outstruct := new(struct {
		Slot        uint64
		CurrentRoot [32]byte
		NextRoot    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Slot = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.CurrentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.NextRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// LatestFinalizedSlotAndCommitteeRoots is a free data retrieval call binding the contract method 0xe153d799.
//
// Solidity: function latestFinalizedSlotAndCommitteeRoots() view returns(uint64 slot, bytes32 currentRoot, bytes32 nextRoot)
func (_EthereumLightClient *EthereumLightClientSession) LatestFinalizedSlotAndCommitteeRoots() (struct {
	Slot        uint64
	CurrentRoot [32]byte
	NextRoot    [32]byte
}, error) {
	return _EthereumLightClient.Contract.LatestFinalizedSlotAndCommitteeRoots(&_EthereumLightClient.CallOpts)
}

// LatestFinalizedSlotAndCommitteeRoots is a free data retrieval call binding the contract method 0xe153d799.
//
// Solidity: function latestFinalizedSlotAndCommitteeRoots() view returns(uint64 slot, bytes32 currentRoot, bytes32 nextRoot)
func (_EthereumLightClient *EthereumLightClientCallerSession) LatestFinalizedSlotAndCommitteeRoots() (struct {
	Slot        uint64
	CurrentRoot [32]byte
	NextRoot    [32]byte
}, error) {
	return _EthereumLightClient.Contract.LatestFinalizedSlotAndCommitteeRoots(&_EthereumLightClient.CallOpts)
}

// NextSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0xe1861b08.
//
// Solidity: function nextSyncCommitteePoseidonRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) NextSyncCommitteePoseidonRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "nextSyncCommitteePoseidonRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0xe1861b08.
//
// Solidity: function nextSyncCommitteePoseidonRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) NextSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.NextSyncCommitteePoseidonRoot(&_EthereumLightClient.CallOpts)
}

// NextSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0xe1861b08.
//
// Solidity: function nextSyncCommitteePoseidonRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) NextSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.NextSyncCommitteePoseidonRoot(&_EthereumLightClient.CallOpts)
}

// NextSyncCommitteeRoot is a free data retrieval call binding the contract method 0x67b49cc7.
//
// Solidity: function nextSyncCommitteeRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) NextSyncCommitteeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "nextSyncCommitteeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextSyncCommitteeRoot is a free data retrieval call binding the contract method 0x67b49cc7.
//
// Solidity: function nextSyncCommitteeRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) NextSyncCommitteeRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.NextSyncCommitteeRoot(&_EthereumLightClient.CallOpts)
}

// NextSyncCommitteeRoot is a free data retrieval call binding the contract method 0x67b49cc7.
//
// Solidity: function nextSyncCommitteeRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) NextSyncCommitteeRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.NextSyncCommitteeRoot(&_EthereumLightClient.CallOpts)
}

// OptimisticExecutionStateRoot is a free data retrieval call binding the contract method 0x39536c8f.
//
// Solidity: function optimisticExecutionStateRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCaller) OptimisticExecutionStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "optimisticExecutionStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OptimisticExecutionStateRoot is a free data retrieval call binding the contract method 0x39536c8f.
//
// Solidity: function optimisticExecutionStateRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientSession) OptimisticExecutionStateRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.OptimisticExecutionStateRoot(&_EthereumLightClient.CallOpts)
}

// OptimisticExecutionStateRoot is a free data retrieval call binding the contract method 0x39536c8f.
//
// Solidity: function optimisticExecutionStateRoot() view returns(bytes32)
func (_EthereumLightClient *EthereumLightClientCallerSession) OptimisticExecutionStateRoot() ([32]byte, error) {
	return _EthereumLightClient.Contract.OptimisticExecutionStateRoot(&_EthereumLightClient.CallOpts)
}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_EthereumLightClient *EthereumLightClientCaller) OptimisticExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "optimisticExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_EthereumLightClient *EthereumLightClientSession) OptimisticExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _EthereumLightClient.Contract.OptimisticExecutionStateRootAndSlot(&_EthereumLightClient.CallOpts)
}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_EthereumLightClient *EthereumLightClientCallerSession) OptimisticExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _EthereumLightClient.Contract.OptimisticExecutionStateRootAndSlot(&_EthereumLightClient.CallOpts)
}

// OptimisticSlot is a free data retrieval call binding the contract method 0x3cf5ea9e.
//
// Solidity: function optimisticSlot() view returns(uint64)
func (_EthereumLightClient *EthereumLightClientCaller) OptimisticSlot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "optimisticSlot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OptimisticSlot is a free data retrieval call binding the contract method 0x3cf5ea9e.
//
// Solidity: function optimisticSlot() view returns(uint64)
func (_EthereumLightClient *EthereumLightClientSession) OptimisticSlot() (uint64, error) {
	return _EthereumLightClient.Contract.OptimisticSlot(&_EthereumLightClient.CallOpts)
}

// OptimisticSlot is a free data retrieval call binding the contract method 0x3cf5ea9e.
//
// Solidity: function optimisticSlot() view returns(uint64)
func (_EthereumLightClient *EthereumLightClientCallerSession) OptimisticSlot() (uint64, error) {
	return _EthereumLightClient.Contract.OptimisticSlot(&_EthereumLightClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthereumLightClient *EthereumLightClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthereumLightClient *EthereumLightClientSession) Owner() (common.Address, error) {
	return _EthereumLightClient.Contract.Owner(&_EthereumLightClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthereumLightClient *EthereumLightClientCallerSession) Owner() (common.Address, error) {
	return _EthereumLightClient.Contract.Owner(&_EthereumLightClient.CallOpts)
}

// VerifyCommitteeSignature is a free data retrieval call binding the contract method 0x3dd3f4aa.
//
// Solidity: function verifyCommitteeSignature(uint64 signatureSlot, (uint64,uint64,bytes32,bytes32,bytes32) header, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate) view returns()
func (_EthereumLightClient *EthereumLightClientCaller) VerifyCommitteeSignature(opts *bind.CallOpts, signatureSlot uint64, header BeaconBlockHeader, syncAggregate SyncAggregate) error {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "verifyCommitteeSignature", signatureSlot, header, syncAggregate)

	if err != nil {
		return err
	}

	return err

}

// VerifyCommitteeSignature is a free data retrieval call binding the contract method 0x3dd3f4aa.
//
// Solidity: function verifyCommitteeSignature(uint64 signatureSlot, (uint64,uint64,bytes32,bytes32,bytes32) header, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate) view returns()
func (_EthereumLightClient *EthereumLightClientSession) VerifyCommitteeSignature(signatureSlot uint64, header BeaconBlockHeader, syncAggregate SyncAggregate) error {
	return _EthereumLightClient.Contract.VerifyCommitteeSignature(&_EthereumLightClient.CallOpts, signatureSlot, header, syncAggregate)
}

// VerifyCommitteeSignature is a free data retrieval call binding the contract method 0x3dd3f4aa.
//
// Solidity: function verifyCommitteeSignature(uint64 signatureSlot, (uint64,uint64,bytes32,bytes32,bytes32) header, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate) view returns()
func (_EthereumLightClient *EthereumLightClientCallerSession) VerifyCommitteeSignature(signatureSlot uint64, header BeaconBlockHeader, syncAggregate SyncAggregate) error {
	return _EthereumLightClient.Contract.VerifyCommitteeSignature(&_EthereumLightClient.CallOpts, signatureSlot, header, syncAggregate)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_EthereumLightClient *EthereumLightClientCaller) ZkVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumLightClient.contract.Call(opts, &out, "zkVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_EthereumLightClient *EthereumLightClientSession) ZkVerifier() (common.Address, error) {
	return _EthereumLightClient.Contract.ZkVerifier(&_EthereumLightClient.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_EthereumLightClient *EthereumLightClientCallerSession) ZkVerifier() (common.Address, error) {
	return _EthereumLightClient.Contract.ZkVerifier(&_EthereumLightClient.CallOpts)
}

// ProcessLightClientForceUpdate is a paid mutator transaction binding the contract method 0xa1a9ad55.
//
// Solidity: function processLightClientForceUpdate() returns()
func (_EthereumLightClient *EthereumLightClientTransactor) ProcessLightClientForceUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumLightClient.contract.Transact(opts, "processLightClientForceUpdate")
}

// ProcessLightClientForceUpdate is a paid mutator transaction binding the contract method 0xa1a9ad55.
//
// Solidity: function processLightClientForceUpdate() returns()
func (_EthereumLightClient *EthereumLightClientSession) ProcessLightClientForceUpdate() (*types.Transaction, error) {
	return _EthereumLightClient.Contract.ProcessLightClientForceUpdate(&_EthereumLightClient.TransactOpts)
}

// ProcessLightClientForceUpdate is a paid mutator transaction binding the contract method 0xa1a9ad55.
//
// Solidity: function processLightClientForceUpdate() returns()
func (_EthereumLightClient *EthereumLightClientTransactorSession) ProcessLightClientForceUpdate() (*types.Transaction, error) {
	return _EthereumLightClient.Contract.ProcessLightClientForceUpdate(&_EthereumLightClient.TransactOpts)
}

// ProcessLightClientUpdate is a paid mutator transaction binding the contract method 0x031523dd.
//
// Solidity: function processLightClientUpdate((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),bytes32[],bytes32,bytes32[],bytes32,(uint256[2],uint256[2][2],uint256[2],uint256[2]),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) update) returns()
func (_EthereumLightClient *EthereumLightClientTransactor) ProcessLightClientUpdate(opts *bind.TransactOpts, update LightClientUpdate) (*types.Transaction, error) {
	return _EthereumLightClient.contract.Transact(opts, "processLightClientUpdate", update)
}

// ProcessLightClientUpdate is a paid mutator transaction binding the contract method 0x031523dd.
//
// Solidity: function processLightClientUpdate((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),bytes32[],bytes32,bytes32[],bytes32,(uint256[2],uint256[2][2],uint256[2],uint256[2]),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) update) returns()
func (_EthereumLightClient *EthereumLightClientSession) ProcessLightClientUpdate(update LightClientUpdate) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.ProcessLightClientUpdate(&_EthereumLightClient.TransactOpts, update)
}

// ProcessLightClientUpdate is a paid mutator transaction binding the contract method 0x031523dd.
//
// Solidity: function processLightClientUpdate((((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])),bytes32[],bytes32,bytes32[],bytes32,(uint256[2],uint256[2][2],uint256[2],uint256[2]),(uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])),uint64) update) returns()
func (_EthereumLightClient *EthereumLightClientTransactorSession) ProcessLightClientUpdate(update LightClientUpdate) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.ProcessLightClientUpdate(&_EthereumLightClient.TransactOpts, update)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthereumLightClient *EthereumLightClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumLightClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthereumLightClient *EthereumLightClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthereumLightClient.Contract.RenounceOwnership(&_EthereumLightClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthereumLightClient *EthereumLightClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthereumLightClient.Contract.RenounceOwnership(&_EthereumLightClient.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthereumLightClient *EthereumLightClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumLightClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthereumLightClient *EthereumLightClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.TransferOwnership(&_EthereumLightClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthereumLightClient *EthereumLightClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.TransferOwnership(&_EthereumLightClient.TransactOpts, newOwner)
}

// UpdateForkVersion is a paid mutator transaction binding the contract method 0xab556e9f.
//
// Solidity: function updateForkVersion(uint64 epoch, bytes4 forkVersion) returns()
func (_EthereumLightClient *EthereumLightClientTransactor) UpdateForkVersion(opts *bind.TransactOpts, epoch uint64, forkVersion [4]byte) (*types.Transaction, error) {
	return _EthereumLightClient.contract.Transact(opts, "updateForkVersion", epoch, forkVersion)
}

// UpdateForkVersion is a paid mutator transaction binding the contract method 0xab556e9f.
//
// Solidity: function updateForkVersion(uint64 epoch, bytes4 forkVersion) returns()
func (_EthereumLightClient *EthereumLightClientSession) UpdateForkVersion(epoch uint64, forkVersion [4]byte) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.UpdateForkVersion(&_EthereumLightClient.TransactOpts, epoch, forkVersion)
}

// UpdateForkVersion is a paid mutator transaction binding the contract method 0xab556e9f.
//
// Solidity: function updateForkVersion(uint64 epoch, bytes4 forkVersion) returns()
func (_EthereumLightClient *EthereumLightClientTransactorSession) UpdateForkVersion(epoch uint64, forkVersion [4]byte) (*types.Transaction, error) {
	return _EthereumLightClient.Contract.UpdateForkVersion(&_EthereumLightClient.TransactOpts, epoch, forkVersion)
}

// EthereumLightClientFinalityUpdateIterator is returned from FilterFinalityUpdate and is used to iterate over the raw logs and unpacked data for FinalityUpdate events raised by the EthereumLightClient contract.
type EthereumLightClientFinalityUpdateIterator struct {
	Event *EthereumLightClientFinalityUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthereumLightClientFinalityUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumLightClientFinalityUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthereumLightClientFinalityUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthereumLightClientFinalityUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumLightClientFinalityUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumLightClientFinalityUpdate represents a FinalityUpdate event raised by the EthereumLightClient contract.
type EthereumLightClientFinalityUpdate struct {
	Slot               *big.Int
	ExecutionStateRoot [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFinalityUpdate is a free log retrieval operation binding the contract event 0x4d75bcddf849ad697dd4b9e37ec69f14240170e980101fcd9e57bb000527c248.
//
// Solidity: event FinalityUpdate(uint256 slot, bytes32 executionStateRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) FilterFinalityUpdate(opts *bind.FilterOpts) (*EthereumLightClientFinalityUpdateIterator, error) {

	logs, sub, err := _EthereumLightClient.contract.FilterLogs(opts, "FinalityUpdate")
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientFinalityUpdateIterator{contract: _EthereumLightClient.contract, event: "FinalityUpdate", logs: logs, sub: sub}, nil
}

// WatchFinalityUpdate is a free log subscription operation binding the contract event 0x4d75bcddf849ad697dd4b9e37ec69f14240170e980101fcd9e57bb000527c248.
//
// Solidity: event FinalityUpdate(uint256 slot, bytes32 executionStateRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) WatchFinalityUpdate(opts *bind.WatchOpts, sink chan<- *EthereumLightClientFinalityUpdate) (event.Subscription, error) {

	logs, sub, err := _EthereumLightClient.contract.WatchLogs(opts, "FinalityUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumLightClientFinalityUpdate)
				if err := _EthereumLightClient.contract.UnpackLog(event, "FinalityUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalityUpdate is a log parse operation binding the contract event 0x4d75bcddf849ad697dd4b9e37ec69f14240170e980101fcd9e57bb000527c248.
//
// Solidity: event FinalityUpdate(uint256 slot, bytes32 executionStateRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) ParseFinalityUpdate(log types.Log) (*EthereumLightClientFinalityUpdate, error) {
	event := new(EthereumLightClientFinalityUpdate)
	if err := _EthereumLightClient.contract.UnpackLog(event, "FinalityUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumLightClientForkVersionUpdatedIterator is returned from FilterForkVersionUpdated and is used to iterate over the raw logs and unpacked data for ForkVersionUpdated events raised by the EthereumLightClient contract.
type EthereumLightClientForkVersionUpdatedIterator struct {
	Event *EthereumLightClientForkVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthereumLightClientForkVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumLightClientForkVersionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthereumLightClientForkVersionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthereumLightClientForkVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumLightClientForkVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumLightClientForkVersionUpdated represents a ForkVersionUpdated event raised by the EthereumLightClient contract.
type EthereumLightClientForkVersionUpdated struct {
	Epoch       uint64
	ForkVersion [4]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterForkVersionUpdated is a free log retrieval operation binding the contract event 0x3d992c45d9456d8ebe181b6a66a3721421393afaa297791373e7569c1abcc8af.
//
// Solidity: event ForkVersionUpdated(uint64 epoch, bytes4 forkVersion)
func (_EthereumLightClient *EthereumLightClientFilterer) FilterForkVersionUpdated(opts *bind.FilterOpts) (*EthereumLightClientForkVersionUpdatedIterator, error) {

	logs, sub, err := _EthereumLightClient.contract.FilterLogs(opts, "ForkVersionUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientForkVersionUpdatedIterator{contract: _EthereumLightClient.contract, event: "ForkVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchForkVersionUpdated is a free log subscription operation binding the contract event 0x3d992c45d9456d8ebe181b6a66a3721421393afaa297791373e7569c1abcc8af.
//
// Solidity: event ForkVersionUpdated(uint64 epoch, bytes4 forkVersion)
func (_EthereumLightClient *EthereumLightClientFilterer) WatchForkVersionUpdated(opts *bind.WatchOpts, sink chan<- *EthereumLightClientForkVersionUpdated) (event.Subscription, error) {

	logs, sub, err := _EthereumLightClient.contract.WatchLogs(opts, "ForkVersionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumLightClientForkVersionUpdated)
				if err := _EthereumLightClient.contract.UnpackLog(event, "ForkVersionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForkVersionUpdated is a log parse operation binding the contract event 0x3d992c45d9456d8ebe181b6a66a3721421393afaa297791373e7569c1abcc8af.
//
// Solidity: event ForkVersionUpdated(uint64 epoch, bytes4 forkVersion)
func (_EthereumLightClient *EthereumLightClientFilterer) ParseForkVersionUpdated(log types.Log) (*EthereumLightClientForkVersionUpdated, error) {
	event := new(EthereumLightClientForkVersionUpdated)
	if err := _EthereumLightClient.contract.UnpackLog(event, "ForkVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumLightClientOptimisticUpdateIterator is returned from FilterOptimisticUpdate and is used to iterate over the raw logs and unpacked data for OptimisticUpdate events raised by the EthereumLightClient contract.
type EthereumLightClientOptimisticUpdateIterator struct {
	Event *EthereumLightClientOptimisticUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthereumLightClientOptimisticUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumLightClientOptimisticUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthereumLightClientOptimisticUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthereumLightClientOptimisticUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumLightClientOptimisticUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumLightClientOptimisticUpdate represents a OptimisticUpdate event raised by the EthereumLightClient contract.
type EthereumLightClientOptimisticUpdate struct {
	Slot               *big.Int
	ExecutionStateRoot [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOptimisticUpdate is a free log retrieval operation binding the contract event 0x27eafee2f5b2c935fa7666d231360699c41fe21db57034136b1af3c43f13df89.
//
// Solidity: event OptimisticUpdate(uint256 slot, bytes32 executionStateRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) FilterOptimisticUpdate(opts *bind.FilterOpts) (*EthereumLightClientOptimisticUpdateIterator, error) {

	logs, sub, err := _EthereumLightClient.contract.FilterLogs(opts, "OptimisticUpdate")
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientOptimisticUpdateIterator{contract: _EthereumLightClient.contract, event: "OptimisticUpdate", logs: logs, sub: sub}, nil
}

// WatchOptimisticUpdate is a free log subscription operation binding the contract event 0x27eafee2f5b2c935fa7666d231360699c41fe21db57034136b1af3c43f13df89.
//
// Solidity: event OptimisticUpdate(uint256 slot, bytes32 executionStateRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) WatchOptimisticUpdate(opts *bind.WatchOpts, sink chan<- *EthereumLightClientOptimisticUpdate) (event.Subscription, error) {

	logs, sub, err := _EthereumLightClient.contract.WatchLogs(opts, "OptimisticUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumLightClientOptimisticUpdate)
				if err := _EthereumLightClient.contract.UnpackLog(event, "OptimisticUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOptimisticUpdate is a log parse operation binding the contract event 0x27eafee2f5b2c935fa7666d231360699c41fe21db57034136b1af3c43f13df89.
//
// Solidity: event OptimisticUpdate(uint256 slot, bytes32 executionStateRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) ParseOptimisticUpdate(log types.Log) (*EthereumLightClientOptimisticUpdate, error) {
	event := new(EthereumLightClientOptimisticUpdate)
	if err := _EthereumLightClient.contract.UnpackLog(event, "OptimisticUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumLightClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthereumLightClient contract.
type EthereumLightClientOwnershipTransferredIterator struct {
	Event *EthereumLightClientOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthereumLightClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumLightClientOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthereumLightClientOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthereumLightClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumLightClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumLightClientOwnershipTransferred represents a OwnershipTransferred event raised by the EthereumLightClient contract.
type EthereumLightClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthereumLightClient *EthereumLightClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthereumLightClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthereumLightClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientOwnershipTransferredIterator{contract: _EthereumLightClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthereumLightClient *EthereumLightClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthereumLightClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthereumLightClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumLightClientOwnershipTransferred)
				if err := _EthereumLightClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthereumLightClient *EthereumLightClientFilterer) ParseOwnershipTransferred(log types.Log) (*EthereumLightClientOwnershipTransferred, error) {
	event := new(EthereumLightClientOwnershipTransferred)
	if err := _EthereumLightClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumLightClientSyncCommitteeUpdatedIterator is returned from FilterSyncCommitteeUpdated and is used to iterate over the raw logs and unpacked data for SyncCommitteeUpdated events raised by the EthereumLightClient contract.
type EthereumLightClientSyncCommitteeUpdatedIterator struct {
	Event *EthereumLightClientSyncCommitteeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthereumLightClientSyncCommitteeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumLightClientSyncCommitteeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthereumLightClientSyncCommitteeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthereumLightClientSyncCommitteeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumLightClientSyncCommitteeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumLightClientSyncCommitteeUpdated represents a SyncCommitteeUpdated event raised by the EthereumLightClient contract.
type EthereumLightClientSyncCommitteeUpdated struct {
	Period       *big.Int
	SszRoot      [32]byte
	PoseidonRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSyncCommitteeUpdated is a free log retrieval operation binding the contract event 0xdc7ba15c782b181b9d54a996db4ab8a32182bff2bfb09e4aca8ed9ea9e5380c7.
//
// Solidity: event SyncCommitteeUpdated(uint256 period, bytes32 sszRoot, bytes32 poseidonRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) FilterSyncCommitteeUpdated(opts *bind.FilterOpts) (*EthereumLightClientSyncCommitteeUpdatedIterator, error) {

	logs, sub, err := _EthereumLightClient.contract.FilterLogs(opts, "SyncCommitteeUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumLightClientSyncCommitteeUpdatedIterator{contract: _EthereumLightClient.contract, event: "SyncCommitteeUpdated", logs: logs, sub: sub}, nil
}

// WatchSyncCommitteeUpdated is a free log subscription operation binding the contract event 0xdc7ba15c782b181b9d54a996db4ab8a32182bff2bfb09e4aca8ed9ea9e5380c7.
//
// Solidity: event SyncCommitteeUpdated(uint256 period, bytes32 sszRoot, bytes32 poseidonRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) WatchSyncCommitteeUpdated(opts *bind.WatchOpts, sink chan<- *EthereumLightClientSyncCommitteeUpdated) (event.Subscription, error) {

	logs, sub, err := _EthereumLightClient.contract.WatchLogs(opts, "SyncCommitteeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumLightClientSyncCommitteeUpdated)
				if err := _EthereumLightClient.contract.UnpackLog(event, "SyncCommitteeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSyncCommitteeUpdated is a log parse operation binding the contract event 0xdc7ba15c782b181b9d54a996db4ab8a32182bff2bfb09e4aca8ed9ea9e5380c7.
//
// Solidity: event SyncCommitteeUpdated(uint256 period, bytes32 sszRoot, bytes32 poseidonRoot)
func (_EthereumLightClient *EthereumLightClientFilterer) ParseSyncCommitteeUpdated(log types.Log) (*EthereumLightClientSyncCommitteeUpdated, error) {
	event := new(EthereumLightClientSyncCommitteeUpdated)
	if err := _EthereumLightClient.contract.UnpackLog(event, "SyncCommitteeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelpersMetaData contains all meta data concerning the Helpers contract.
var HelpersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea2646970667358221220ca02bdd00b3565f52b5f35b0bbba26d5e9e9ea0c39ae706c9c462384cb31a38b64736f6c63430008140033",
}

// HelpersABI is the input ABI used to generate the binding from.
// Deprecated: Use HelpersMetaData.ABI instead.
var HelpersABI = HelpersMetaData.ABI

// HelpersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelpersMetaData.Bin instead.
var HelpersBin = HelpersMetaData.Bin

// DeployHelpers deploys a new Ethereum contract, binding an instance of Helpers to it.
func DeployHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helpers, error) {
	parsed, err := HelpersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helpers{HelpersCaller: HelpersCaller{contract: contract}, HelpersTransactor: HelpersTransactor{contract: contract}, HelpersFilterer: HelpersFilterer{contract: contract}}, nil
}

// Helpers is an auto generated Go binding around an Ethereum contract.
type Helpers struct {
	HelpersCaller     // Read-only binding to the contract
	HelpersTransactor // Write-only binding to the contract
	HelpersFilterer   // Log filterer for contract events
}

// HelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelpersSession struct {
	Contract     *Helpers          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelpersCallerSession struct {
	Contract *HelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelpersTransactorSession struct {
	Contract     *HelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelpersRaw struct {
	Contract *Helpers // Generic contract binding to access the raw methods on
}

// HelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelpersCallerRaw struct {
	Contract *HelpersCaller // Generic read-only contract binding to access the raw methods on
}

// HelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelpersTransactorRaw struct {
	Contract *HelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelpers creates a new instance of Helpers, bound to a specific deployed contract.
func NewHelpers(address common.Address, backend bind.ContractBackend) (*Helpers, error) {
	contract, err := bindHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helpers{HelpersCaller: HelpersCaller{contract: contract}, HelpersTransactor: HelpersTransactor{contract: contract}, HelpersFilterer: HelpersFilterer{contract: contract}}, nil
}

// NewHelpersCaller creates a new read-only instance of Helpers, bound to a specific deployed contract.
func NewHelpersCaller(address common.Address, caller bind.ContractCaller) (*HelpersCaller, error) {
	contract, err := bindHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelpersCaller{contract: contract}, nil
}

// NewHelpersTransactor creates a new write-only instance of Helpers, bound to a specific deployed contract.
func NewHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*HelpersTransactor, error) {
	contract, err := bindHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelpersTransactor{contract: contract}, nil
}

// NewHelpersFilterer creates a new log filterer instance of Helpers, bound to a specific deployed contract.
func NewHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*HelpersFilterer, error) {
	contract, err := bindHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelpersFilterer{contract: contract}, nil
}

// bindHelpers binds a generic wrapper to an already deployed contract.
func bindHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelpersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpers *HelpersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helpers.Contract.HelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpers *HelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpers.Contract.HelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpers *HelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpers.Contract.HelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpers *HelpersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpers *HelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpers *HelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpers.Contract.contract.Transact(opts, method, params...)
}

// IAnchorBlocksMetaData contains all meta data concerning the IAnchorBlocks contract.
var IAnchorBlocksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAnchorBlocksABI is the input ABI used to generate the binding from.
// Deprecated: Use IAnchorBlocksMetaData.ABI instead.
var IAnchorBlocksABI = IAnchorBlocksMetaData.ABI

// IAnchorBlocks is an auto generated Go binding around an Ethereum contract.
type IAnchorBlocks struct {
	IAnchorBlocksCaller     // Read-only binding to the contract
	IAnchorBlocksTransactor // Write-only binding to the contract
	IAnchorBlocksFilterer   // Log filterer for contract events
}

// IAnchorBlocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAnchorBlocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAnchorBlocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAnchorBlocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAnchorBlocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAnchorBlocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAnchorBlocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAnchorBlocksSession struct {
	Contract     *IAnchorBlocks    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAnchorBlocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAnchorBlocksCallerSession struct {
	Contract *IAnchorBlocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAnchorBlocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAnchorBlocksTransactorSession struct {
	Contract     *IAnchorBlocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAnchorBlocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAnchorBlocksRaw struct {
	Contract *IAnchorBlocks // Generic contract binding to access the raw methods on
}

// IAnchorBlocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAnchorBlocksCallerRaw struct {
	Contract *IAnchorBlocksCaller // Generic read-only contract binding to access the raw methods on
}

// IAnchorBlocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAnchorBlocksTransactorRaw struct {
	Contract *IAnchorBlocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAnchorBlocks creates a new instance of IAnchorBlocks, bound to a specific deployed contract.
func NewIAnchorBlocks(address common.Address, backend bind.ContractBackend) (*IAnchorBlocks, error) {
	contract, err := bindIAnchorBlocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAnchorBlocks{IAnchorBlocksCaller: IAnchorBlocksCaller{contract: contract}, IAnchorBlocksTransactor: IAnchorBlocksTransactor{contract: contract}, IAnchorBlocksFilterer: IAnchorBlocksFilterer{contract: contract}}, nil
}

// NewIAnchorBlocksCaller creates a new read-only instance of IAnchorBlocks, bound to a specific deployed contract.
func NewIAnchorBlocksCaller(address common.Address, caller bind.ContractCaller) (*IAnchorBlocksCaller, error) {
	contract, err := bindIAnchorBlocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAnchorBlocksCaller{contract: contract}, nil
}

// NewIAnchorBlocksTransactor creates a new write-only instance of IAnchorBlocks, bound to a specific deployed contract.
func NewIAnchorBlocksTransactor(address common.Address, transactor bind.ContractTransactor) (*IAnchorBlocksTransactor, error) {
	contract, err := bindIAnchorBlocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAnchorBlocksTransactor{contract: contract}, nil
}

// NewIAnchorBlocksFilterer creates a new log filterer instance of IAnchorBlocks, bound to a specific deployed contract.
func NewIAnchorBlocksFilterer(address common.Address, filterer bind.ContractFilterer) (*IAnchorBlocksFilterer, error) {
	contract, err := bindIAnchorBlocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAnchorBlocksFilterer{contract: contract}, nil
}

// bindIAnchorBlocks binds a generic wrapper to an already deployed contract.
func bindIAnchorBlocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAnchorBlocksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAnchorBlocks *IAnchorBlocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAnchorBlocks.Contract.IAnchorBlocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAnchorBlocks *IAnchorBlocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAnchorBlocks.Contract.IAnchorBlocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAnchorBlocks *IAnchorBlocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAnchorBlocks.Contract.IAnchorBlocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAnchorBlocks *IAnchorBlocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAnchorBlocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAnchorBlocks *IAnchorBlocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAnchorBlocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAnchorBlocks *IAnchorBlocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAnchorBlocks.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 blockNum) view returns(bytes32)
func (_IAnchorBlocks *IAnchorBlocksCaller) Blocks(opts *bind.CallOpts, blockNum *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IAnchorBlocks.contract.Call(opts, &out, "blocks", blockNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 blockNum) view returns(bytes32)
func (_IAnchorBlocks *IAnchorBlocksSession) Blocks(blockNum *big.Int) ([32]byte, error) {
	return _IAnchorBlocks.Contract.Blocks(&_IAnchorBlocks.CallOpts, blockNum)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 blockNum) view returns(bytes32)
func (_IAnchorBlocks *IAnchorBlocksCallerSession) Blocks(blockNum *big.Int) ([32]byte, error) {
	return _IAnchorBlocks.Contract.Blocks(&_IAnchorBlocks.CallOpts, blockNum)
}

// IBeaconVerifierMetaData contains all meta data concerning the IBeaconVerifier contract.
var IBeaconVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signingRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteePoseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"participation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"verifySignatureProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sszRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"verifySyncCommitteeRootMappingProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBeaconVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconVerifierMetaData.ABI instead.
var IBeaconVerifierABI = IBeaconVerifierMetaData.ABI

// IBeaconVerifier is an auto generated Go binding around an Ethereum contract.
type IBeaconVerifier struct {
	IBeaconVerifierCaller     // Read-only binding to the contract
	IBeaconVerifierTransactor // Write-only binding to the contract
	IBeaconVerifierFilterer   // Log filterer for contract events
}

// IBeaconVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBeaconVerifierSession struct {
	Contract     *IBeaconVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBeaconVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBeaconVerifierCallerSession struct {
	Contract *IBeaconVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IBeaconVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBeaconVerifierTransactorSession struct {
	Contract     *IBeaconVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IBeaconVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBeaconVerifierRaw struct {
	Contract *IBeaconVerifier // Generic contract binding to access the raw methods on
}

// IBeaconVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBeaconVerifierCallerRaw struct {
	Contract *IBeaconVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBeaconVerifierTransactorRaw struct {
	Contract *IBeaconVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeaconVerifier creates a new instance of IBeaconVerifier, bound to a specific deployed contract.
func NewIBeaconVerifier(address common.Address, backend bind.ContractBackend) (*IBeaconVerifier, error) {
	contract, err := bindIBeaconVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeaconVerifier{IBeaconVerifierCaller: IBeaconVerifierCaller{contract: contract}, IBeaconVerifierTransactor: IBeaconVerifierTransactor{contract: contract}, IBeaconVerifierFilterer: IBeaconVerifierFilterer{contract: contract}}, nil
}

// NewIBeaconVerifierCaller creates a new read-only instance of IBeaconVerifier, bound to a specific deployed contract.
func NewIBeaconVerifierCaller(address common.Address, caller bind.ContractCaller) (*IBeaconVerifierCaller, error) {
	contract, err := bindIBeaconVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconVerifierCaller{contract: contract}, nil
}

// NewIBeaconVerifierTransactor creates a new write-only instance of IBeaconVerifier, bound to a specific deployed contract.
func NewIBeaconVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconVerifierTransactor, error) {
	contract, err := bindIBeaconVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconVerifierTransactor{contract: contract}, nil
}

// NewIBeaconVerifierFilterer creates a new log filterer instance of IBeaconVerifier, bound to a specific deployed contract.
func NewIBeaconVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconVerifierFilterer, error) {
	contract, err := bindIBeaconVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconVerifierFilterer{contract: contract}, nil
}

// bindIBeaconVerifier binds a generic wrapper to an already deployed contract.
func bindIBeaconVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBeaconVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconVerifier *IBeaconVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconVerifier.Contract.IBeaconVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconVerifier *IBeaconVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconVerifier.Contract.IBeaconVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconVerifier *IBeaconVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconVerifier.Contract.IBeaconVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconVerifier *IBeaconVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconVerifier *IBeaconVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconVerifier *IBeaconVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifySignatureProof is a free data retrieval call binding the contract method 0x52356da0.
//
// Solidity: function verifySignatureProof(bytes32 signingRoot, bytes32 syncCommitteePoseidonRoot, uint256 participation, uint256 commitment, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_IBeaconVerifier *IBeaconVerifierCaller) VerifySignatureProof(opts *bind.CallOpts, signingRoot [32]byte, syncCommitteePoseidonRoot [32]byte, participation *big.Int, commitment *big.Int, p IBeaconVerifierProof) (bool, error) {
	var out []interface{}
	err := _IBeaconVerifier.contract.Call(opts, &out, "verifySignatureProof", signingRoot, syncCommitteePoseidonRoot, participation, commitment, p)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignatureProof is a free data retrieval call binding the contract method 0x52356da0.
//
// Solidity: function verifySignatureProof(bytes32 signingRoot, bytes32 syncCommitteePoseidonRoot, uint256 participation, uint256 commitment, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_IBeaconVerifier *IBeaconVerifierSession) VerifySignatureProof(signingRoot [32]byte, syncCommitteePoseidonRoot [32]byte, participation *big.Int, commitment *big.Int, p IBeaconVerifierProof) (bool, error) {
	return _IBeaconVerifier.Contract.VerifySignatureProof(&_IBeaconVerifier.CallOpts, signingRoot, syncCommitteePoseidonRoot, participation, commitment, p)
}

// VerifySignatureProof is a free data retrieval call binding the contract method 0x52356da0.
//
// Solidity: function verifySignatureProof(bytes32 signingRoot, bytes32 syncCommitteePoseidonRoot, uint256 participation, uint256 commitment, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_IBeaconVerifier *IBeaconVerifierCallerSession) VerifySignatureProof(signingRoot [32]byte, syncCommitteePoseidonRoot [32]byte, participation *big.Int, commitment *big.Int, p IBeaconVerifierProof) (bool, error) {
	return _IBeaconVerifier.Contract.VerifySignatureProof(&_IBeaconVerifier.CallOpts, signingRoot, syncCommitteePoseidonRoot, participation, commitment, p)
}

// VerifySyncCommitteeRootMappingProof is a free data retrieval call binding the contract method 0x7a538781.
//
// Solidity: function verifySyncCommitteeRootMappingProof(bytes32 sszRoot, bytes32 poseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_IBeaconVerifier *IBeaconVerifierCaller) VerifySyncCommitteeRootMappingProof(opts *bind.CallOpts, sszRoot [32]byte, poseidonRoot [32]byte, p IBeaconVerifierProof) (bool, error) {
	var out []interface{}
	err := _IBeaconVerifier.contract.Call(opts, &out, "verifySyncCommitteeRootMappingProof", sszRoot, poseidonRoot, p)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySyncCommitteeRootMappingProof is a free data retrieval call binding the contract method 0x7a538781.
//
// Solidity: function verifySyncCommitteeRootMappingProof(bytes32 sszRoot, bytes32 poseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_IBeaconVerifier *IBeaconVerifierSession) VerifySyncCommitteeRootMappingProof(sszRoot [32]byte, poseidonRoot [32]byte, p IBeaconVerifierProof) (bool, error) {
	return _IBeaconVerifier.Contract.VerifySyncCommitteeRootMappingProof(&_IBeaconVerifier.CallOpts, sszRoot, poseidonRoot, p)
}

// VerifySyncCommitteeRootMappingProof is a free data retrieval call binding the contract method 0x7a538781.
//
// Solidity: function verifySyncCommitteeRootMappingProof(bytes32 sszRoot, bytes32 poseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) p) view returns(bool)
func (_IBeaconVerifier *IBeaconVerifierCallerSession) VerifySyncCommitteeRootMappingProof(sszRoot [32]byte, poseidonRoot [32]byte, p IBeaconVerifierProof) (bool, error) {
	return _IBeaconVerifier.Contract.VerifySyncCommitteeRootMappingProof(&_IBeaconVerifier.CallOpts, sszRoot, poseidonRoot, p)
}

// IBlockChunksMetaData contains all meta data concerning the IBlockChunks contract.
var IBlockChunksMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numFinal\",\"type\":\"uint32\"}],\"name\":\"UpdateEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"startBlockNumber\",\"type\":\"uint32\"}],\"name\":\"historicalRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"claimedBlkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"numFinal\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[7]\",\"name\":\"merkleProof\",\"type\":\"bytes32[7]\"}],\"internalType\":\"structIBlockChunks.BlockHashWitness\",\"name\":\"witness\",\"type\":\"tuple\"}],\"name\":\"isBlockHashValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nextRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nextNumFinal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"}],\"name\":\"updateOld\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"}],\"name\":\"updateRecent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBlockChunksABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlockChunksMetaData.ABI instead.
var IBlockChunksABI = IBlockChunksMetaData.ABI

// IBlockChunks is an auto generated Go binding around an Ethereum contract.
type IBlockChunks struct {
	IBlockChunksCaller     // Read-only binding to the contract
	IBlockChunksTransactor // Write-only binding to the contract
	IBlockChunksFilterer   // Log filterer for contract events
}

// IBlockChunksCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlockChunksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockChunksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlockChunksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockChunksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlockChunksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockChunksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlockChunksSession struct {
	Contract     *IBlockChunks     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlockChunksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlockChunksCallerSession struct {
	Contract *IBlockChunksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBlockChunksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlockChunksTransactorSession struct {
	Contract     *IBlockChunksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBlockChunksRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlockChunksRaw struct {
	Contract *IBlockChunks // Generic contract binding to access the raw methods on
}

// IBlockChunksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlockChunksCallerRaw struct {
	Contract *IBlockChunksCaller // Generic read-only contract binding to access the raw methods on
}

// IBlockChunksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlockChunksTransactorRaw struct {
	Contract *IBlockChunksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlockChunks creates a new instance of IBlockChunks, bound to a specific deployed contract.
func NewIBlockChunks(address common.Address, backend bind.ContractBackend) (*IBlockChunks, error) {
	contract, err := bindIBlockChunks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlockChunks{IBlockChunksCaller: IBlockChunksCaller{contract: contract}, IBlockChunksTransactor: IBlockChunksTransactor{contract: contract}, IBlockChunksFilterer: IBlockChunksFilterer{contract: contract}}, nil
}

// NewIBlockChunksCaller creates a new read-only instance of IBlockChunks, bound to a specific deployed contract.
func NewIBlockChunksCaller(address common.Address, caller bind.ContractCaller) (*IBlockChunksCaller, error) {
	contract, err := bindIBlockChunks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlockChunksCaller{contract: contract}, nil
}

// NewIBlockChunksTransactor creates a new write-only instance of IBlockChunks, bound to a specific deployed contract.
func NewIBlockChunksTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlockChunksTransactor, error) {
	contract, err := bindIBlockChunks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlockChunksTransactor{contract: contract}, nil
}

// NewIBlockChunksFilterer creates a new log filterer instance of IBlockChunks, bound to a specific deployed contract.
func NewIBlockChunksFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlockChunksFilterer, error) {
	contract, err := bindIBlockChunks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlockChunksFilterer{contract: contract}, nil
}

// bindIBlockChunks binds a generic wrapper to an already deployed contract.
func bindIBlockChunks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBlockChunksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlockChunks *IBlockChunksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlockChunks.Contract.IBlockChunksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlockChunks *IBlockChunksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockChunks.Contract.IBlockChunksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlockChunks *IBlockChunksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlockChunks.Contract.IBlockChunksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlockChunks *IBlockChunksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlockChunks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlockChunks *IBlockChunksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockChunks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlockChunks *IBlockChunksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlockChunks.Contract.contract.Transact(opts, method, params...)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0xb551a187.
//
// Solidity: function historicalRoots(uint64 chainId, uint32 startBlockNumber) view returns(bytes32)
func (_IBlockChunks *IBlockChunksCaller) HistoricalRoots(opts *bind.CallOpts, chainId uint64, startBlockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _IBlockChunks.contract.Call(opts, &out, "historicalRoots", chainId, startBlockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HistoricalRoots is a free data retrieval call binding the contract method 0xb551a187.
//
// Solidity: function historicalRoots(uint64 chainId, uint32 startBlockNumber) view returns(bytes32)
func (_IBlockChunks *IBlockChunksSession) HistoricalRoots(chainId uint64, startBlockNumber uint32) ([32]byte, error) {
	return _IBlockChunks.Contract.HistoricalRoots(&_IBlockChunks.CallOpts, chainId, startBlockNumber)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0xb551a187.
//
// Solidity: function historicalRoots(uint64 chainId, uint32 startBlockNumber) view returns(bytes32)
func (_IBlockChunks *IBlockChunksCallerSession) HistoricalRoots(chainId uint64, startBlockNumber uint32) ([32]byte, error) {
	return _IBlockChunks.Contract.HistoricalRoots(&_IBlockChunks.CallOpts, chainId, startBlockNumber)
}

// IsBlockHashValid is a free data retrieval call binding the contract method 0x544f73a4.
//
// Solidity: function isBlockHashValid((uint64,uint32,bytes32,bytes32,uint32,bytes32[7]) witness) view returns(bool)
func (_IBlockChunks *IBlockChunksCaller) IsBlockHashValid(opts *bind.CallOpts, witness IBlockChunksBlockHashWitness) (bool, error) {
	var out []interface{}
	err := _IBlockChunks.contract.Call(opts, &out, "isBlockHashValid", witness)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockHashValid is a free data retrieval call binding the contract method 0x544f73a4.
//
// Solidity: function isBlockHashValid((uint64,uint32,bytes32,bytes32,uint32,bytes32[7]) witness) view returns(bool)
func (_IBlockChunks *IBlockChunksSession) IsBlockHashValid(witness IBlockChunksBlockHashWitness) (bool, error) {
	return _IBlockChunks.Contract.IsBlockHashValid(&_IBlockChunks.CallOpts, witness)
}

// IsBlockHashValid is a free data retrieval call binding the contract method 0x544f73a4.
//
// Solidity: function isBlockHashValid((uint64,uint32,bytes32,bytes32,uint32,bytes32[7]) witness) view returns(bool)
func (_IBlockChunks *IBlockChunksCallerSession) IsBlockHashValid(witness IBlockChunksBlockHashWitness) (bool, error) {
	return _IBlockChunks.Contract.IsBlockHashValid(&_IBlockChunks.CallOpts, witness)
}

// UpdateOld is a paid mutator transaction binding the contract method 0x52cfc560.
//
// Solidity: function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes proofData) returns()
func (_IBlockChunks *IBlockChunksTransactor) UpdateOld(opts *bind.TransactOpts, chainId uint64, nextRoot [32]byte, nextNumFinal uint32, proofData []byte) (*types.Transaction, error) {
	return _IBlockChunks.contract.Transact(opts, "updateOld", chainId, nextRoot, nextNumFinal, proofData)
}

// UpdateOld is a paid mutator transaction binding the contract method 0x52cfc560.
//
// Solidity: function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes proofData) returns()
func (_IBlockChunks *IBlockChunksSession) UpdateOld(chainId uint64, nextRoot [32]byte, nextNumFinal uint32, proofData []byte) (*types.Transaction, error) {
	return _IBlockChunks.Contract.UpdateOld(&_IBlockChunks.TransactOpts, chainId, nextRoot, nextNumFinal, proofData)
}

// UpdateOld is a paid mutator transaction binding the contract method 0x52cfc560.
//
// Solidity: function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes proofData) returns()
func (_IBlockChunks *IBlockChunksTransactorSession) UpdateOld(chainId uint64, nextRoot [32]byte, nextNumFinal uint32, proofData []byte) (*types.Transaction, error) {
	return _IBlockChunks.Contract.UpdateOld(&_IBlockChunks.TransactOpts, chainId, nextRoot, nextNumFinal, proofData)
}

// UpdateRecent is a paid mutator transaction binding the contract method 0x50a26d48.
//
// Solidity: function updateRecent(uint64 chainId, bytes proofData) returns()
func (_IBlockChunks *IBlockChunksTransactor) UpdateRecent(opts *bind.TransactOpts, chainId uint64, proofData []byte) (*types.Transaction, error) {
	return _IBlockChunks.contract.Transact(opts, "updateRecent", chainId, proofData)
}

// UpdateRecent is a paid mutator transaction binding the contract method 0x50a26d48.
//
// Solidity: function updateRecent(uint64 chainId, bytes proofData) returns()
func (_IBlockChunks *IBlockChunksSession) UpdateRecent(chainId uint64, proofData []byte) (*types.Transaction, error) {
	return _IBlockChunks.Contract.UpdateRecent(&_IBlockChunks.TransactOpts, chainId, proofData)
}

// UpdateRecent is a paid mutator transaction binding the contract method 0x50a26d48.
//
// Solidity: function updateRecent(uint64 chainId, bytes proofData) returns()
func (_IBlockChunks *IBlockChunksTransactorSession) UpdateRecent(chainId uint64, proofData []byte) (*types.Transaction, error) {
	return _IBlockChunks.Contract.UpdateRecent(&_IBlockChunks.TransactOpts, chainId, proofData)
}

// IBlockChunksUpdateEventIterator is returned from FilterUpdateEvent and is used to iterate over the raw logs and unpacked data for UpdateEvent events raised by the IBlockChunks contract.
type IBlockChunksUpdateEventIterator struct {
	Event *IBlockChunksUpdateEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBlockChunksUpdateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlockChunksUpdateEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBlockChunksUpdateEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBlockChunksUpdateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlockChunksUpdateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlockChunksUpdateEvent represents a UpdateEvent event raised by the IBlockChunks contract.
type IBlockChunksUpdateEvent struct {
	ChainId          uint64
	StartBlockNumber uint32
	PrevHash         [32]byte
	Root             [32]byte
	NumFinal         uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateEvent is a free log retrieval operation binding the contract event 0xa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e.
//
// Solidity: event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal)
func (_IBlockChunks *IBlockChunksFilterer) FilterUpdateEvent(opts *bind.FilterOpts) (*IBlockChunksUpdateEventIterator, error) {

	logs, sub, err := _IBlockChunks.contract.FilterLogs(opts, "UpdateEvent")
	if err != nil {
		return nil, err
	}
	return &IBlockChunksUpdateEventIterator{contract: _IBlockChunks.contract, event: "UpdateEvent", logs: logs, sub: sub}, nil
}

// WatchUpdateEvent is a free log subscription operation binding the contract event 0xa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e.
//
// Solidity: event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal)
func (_IBlockChunks *IBlockChunksFilterer) WatchUpdateEvent(opts *bind.WatchOpts, sink chan<- *IBlockChunksUpdateEvent) (event.Subscription, error) {

	logs, sub, err := _IBlockChunks.contract.WatchLogs(opts, "UpdateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlockChunksUpdateEvent)
				if err := _IBlockChunks.contract.UnpackLog(event, "UpdateEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateEvent is a log parse operation binding the contract event 0xa946b3d79b3150fec8e4d9ccc6100de98308a472f1ba57d23aeda162bb86e33e.
//
// Solidity: event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal)
func (_IBlockChunks *IBlockChunksFilterer) ParseUpdateEvent(log types.Log) (*IBlockChunksUpdateEvent, error) {
	event := new(IBlockChunksUpdateEvent)
	if err := _IBlockChunks.contract.UnpackLog(event, "UpdateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthereumLightClientMetaData contains all meta data concerning the IEthereumLightClient contract.
var IEthereumLightClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"finalizedExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"}],\"name\":\"verifyCommitteeSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IEthereumLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthereumLightClientMetaData.ABI instead.
var IEthereumLightClientABI = IEthereumLightClientMetaData.ABI

// IEthereumLightClient is an auto generated Go binding around an Ethereum contract.
type IEthereumLightClient struct {
	IEthereumLightClientCaller     // Read-only binding to the contract
	IEthereumLightClientTransactor // Write-only binding to the contract
	IEthereumLightClientFilterer   // Log filterer for contract events
}

// IEthereumLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthereumLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthereumLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthereumLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthereumLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthereumLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthereumLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthereumLightClientSession struct {
	Contract     *IEthereumLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IEthereumLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthereumLightClientCallerSession struct {
	Contract *IEthereumLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IEthereumLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthereumLightClientTransactorSession struct {
	Contract     *IEthereumLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IEthereumLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthereumLightClientRaw struct {
	Contract *IEthereumLightClient // Generic contract binding to access the raw methods on
}

// IEthereumLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthereumLightClientCallerRaw struct {
	Contract *IEthereumLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// IEthereumLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthereumLightClientTransactorRaw struct {
	Contract *IEthereumLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthereumLightClient creates a new instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClient(address common.Address, backend bind.ContractBackend) (*IEthereumLightClient, error) {
	contract, err := bindIEthereumLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClient{IEthereumLightClientCaller: IEthereumLightClientCaller{contract: contract}, IEthereumLightClientTransactor: IEthereumLightClientTransactor{contract: contract}, IEthereumLightClientFilterer: IEthereumLightClientFilterer{contract: contract}}, nil
}

// NewIEthereumLightClientCaller creates a new read-only instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClientCaller(address common.Address, caller bind.ContractCaller) (*IEthereumLightClientCaller, error) {
	contract, err := bindIEthereumLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClientCaller{contract: contract}, nil
}

// NewIEthereumLightClientTransactor creates a new write-only instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthereumLightClientTransactor, error) {
	contract, err := bindIEthereumLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClientTransactor{contract: contract}, nil
}

// NewIEthereumLightClientFilterer creates a new log filterer instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthereumLightClientFilterer, error) {
	contract, err := bindIEthereumLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClientFilterer{contract: contract}, nil
}

// bindIEthereumLightClient binds a generic wrapper to an already deployed contract.
func bindIEthereumLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEthereumLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthereumLightClient *IEthereumLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthereumLightClient.Contract.IEthereumLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthereumLightClient *IEthereumLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.IEthereumLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthereumLightClient *IEthereumLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.IEthereumLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthereumLightClient *IEthereumLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthereumLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthereumLightClient *IEthereumLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthereumLightClient *IEthereumLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.contract.Transact(opts, method, params...)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientCaller) FinalizedExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _IEthereumLightClient.contract.Call(opts, &out, "finalizedExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _IEthereumLightClient.Contract.FinalizedExecutionStateRootAndSlot(&_IEthereumLightClient.CallOpts)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientCallerSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _IEthereumLightClient.Contract.FinalizedExecutionStateRootAndSlot(&_IEthereumLightClient.CallOpts)
}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientCaller) OptimisticExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _IEthereumLightClient.contract.Call(opts, &out, "optimisticExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientSession) OptimisticExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _IEthereumLightClient.Contract.OptimisticExecutionStateRootAndSlot(&_IEthereumLightClient.CallOpts)
}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientCallerSession) OptimisticExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _IEthereumLightClient.Contract.OptimisticExecutionStateRootAndSlot(&_IEthereumLightClient.CallOpts)
}

// VerifyCommitteeSignature is a free data retrieval call binding the contract method 0x3dd3f4aa.
//
// Solidity: function verifyCommitteeSignature(uint64 signatureSlot, (uint64,uint64,bytes32,bytes32,bytes32) header, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate) view returns()
func (_IEthereumLightClient *IEthereumLightClientCaller) VerifyCommitteeSignature(opts *bind.CallOpts, signatureSlot uint64, header BeaconBlockHeader, syncAggregate SyncAggregate) error {
	var out []interface{}
	err := _IEthereumLightClient.contract.Call(opts, &out, "verifyCommitteeSignature", signatureSlot, header, syncAggregate)

	if err != nil {
		return err
	}

	return err

}

// VerifyCommitteeSignature is a free data retrieval call binding the contract method 0x3dd3f4aa.
//
// Solidity: function verifyCommitteeSignature(uint64 signatureSlot, (uint64,uint64,bytes32,bytes32,bytes32) header, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate) view returns()
func (_IEthereumLightClient *IEthereumLightClientSession) VerifyCommitteeSignature(signatureSlot uint64, header BeaconBlockHeader, syncAggregate SyncAggregate) error {
	return _IEthereumLightClient.Contract.VerifyCommitteeSignature(&_IEthereumLightClient.CallOpts, signatureSlot, header, syncAggregate)
}

// VerifyCommitteeSignature is a free data retrieval call binding the contract method 0x3dd3f4aa.
//
// Solidity: function verifyCommitteeSignature(uint64 signatureSlot, (uint64,uint64,bytes32,bytes32,bytes32) header, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate) view returns()
func (_IEthereumLightClient *IEthereumLightClientCallerSession) VerifyCommitteeSignature(signatureSlot uint64, header BeaconBlockHeader, syncAggregate SyncAggregate) error {
	return _IEthereumLightClient.Contract.VerifyCommitteeSignature(&_IEthereumLightClient.CallOpts, signatureSlot, header, syncAggregate)
}

// IReceiptVerifierMetaData contains all meta data concerning the IReceiptVerifier contract.
var IReceiptVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIReceiptVerifier.LogInfo[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIReceiptVerifier.ReceiptInfo\",\"name\":\"receiptInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptAndLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIReceiptVerifier.LogInfo[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIReceiptVerifier.ReceiptInfo\",\"name\":\"receiptInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IReceiptVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IReceiptVerifierMetaData.ABI instead.
var IReceiptVerifierABI = IReceiptVerifierMetaData.ABI

// IReceiptVerifier is an auto generated Go binding around an Ethereum contract.
type IReceiptVerifier struct {
	IReceiptVerifierCaller     // Read-only binding to the contract
	IReceiptVerifierTransactor // Write-only binding to the contract
	IReceiptVerifierFilterer   // Log filterer for contract events
}

// IReceiptVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReceiptVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiptVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReceiptVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiptVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReceiptVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiptVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReceiptVerifierSession struct {
	Contract     *IReceiptVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReceiptVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReceiptVerifierCallerSession struct {
	Contract *IReceiptVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IReceiptVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReceiptVerifierTransactorSession struct {
	Contract     *IReceiptVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IReceiptVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReceiptVerifierRaw struct {
	Contract *IReceiptVerifier // Generic contract binding to access the raw methods on
}

// IReceiptVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReceiptVerifierCallerRaw struct {
	Contract *IReceiptVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IReceiptVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReceiptVerifierTransactorRaw struct {
	Contract *IReceiptVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReceiptVerifier creates a new instance of IReceiptVerifier, bound to a specific deployed contract.
func NewIReceiptVerifier(address common.Address, backend bind.ContractBackend) (*IReceiptVerifier, error) {
	contract, err := bindIReceiptVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReceiptVerifier{IReceiptVerifierCaller: IReceiptVerifierCaller{contract: contract}, IReceiptVerifierTransactor: IReceiptVerifierTransactor{contract: contract}, IReceiptVerifierFilterer: IReceiptVerifierFilterer{contract: contract}}, nil
}

// NewIReceiptVerifierCaller creates a new read-only instance of IReceiptVerifier, bound to a specific deployed contract.
func NewIReceiptVerifierCaller(address common.Address, caller bind.ContractCaller) (*IReceiptVerifierCaller, error) {
	contract, err := bindIReceiptVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiptVerifierCaller{contract: contract}, nil
}

// NewIReceiptVerifierTransactor creates a new write-only instance of IReceiptVerifier, bound to a specific deployed contract.
func NewIReceiptVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IReceiptVerifierTransactor, error) {
	contract, err := bindIReceiptVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiptVerifierTransactor{contract: contract}, nil
}

// NewIReceiptVerifierFilterer creates a new log filterer instance of IReceiptVerifier, bound to a specific deployed contract.
func NewIReceiptVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IReceiptVerifierFilterer, error) {
	contract, err := bindIReceiptVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReceiptVerifierFilterer{contract: contract}, nil
}

// bindIReceiptVerifier binds a generic wrapper to an already deployed contract.
func bindIReceiptVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReceiptVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiptVerifier *IReceiptVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiptVerifier.Contract.IReceiptVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiptVerifier *IReceiptVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiptVerifier.Contract.IReceiptVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiptVerifier *IReceiptVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiptVerifier.Contract.IReceiptVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiptVerifier *IReceiptVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiptVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiptVerifier *IReceiptVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiptVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiptVerifier *IReceiptVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiptVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x68ac2f78.
//
// Solidity: function verifyReceipt(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) receiptInfo)
func (_IReceiptVerifier *IReceiptVerifierCaller) VerifyReceipt(opts *bind.CallOpts, receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (IReceiptVerifierReceiptInfo, error) {
	var out []interface{}
	err := _IReceiptVerifier.contract.Call(opts, &out, "verifyReceipt", receiptRaw, proofData, auxiBlkVerifyInfo)

	if err != nil {
		return *new(IReceiptVerifierReceiptInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IReceiptVerifierReceiptInfo)).(*IReceiptVerifierReceiptInfo)

	return out0, err

}

// VerifyReceipt is a free data retrieval call binding the contract method 0x68ac2f78.
//
// Solidity: function verifyReceipt(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) receiptInfo)
func (_IReceiptVerifier *IReceiptVerifierSession) VerifyReceipt(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (IReceiptVerifierReceiptInfo, error) {
	return _IReceiptVerifier.Contract.VerifyReceipt(&_IReceiptVerifier.CallOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x68ac2f78.
//
// Solidity: function verifyReceipt(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) receiptInfo)
func (_IReceiptVerifier *IReceiptVerifierCallerSession) VerifyReceipt(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (IReceiptVerifierReceiptInfo, error) {
	return _IReceiptVerifier.Contract.VerifyReceipt(&_IReceiptVerifier.CallOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceiptAndLog is a paid mutator transaction binding the contract method 0x3996da7a.
//
// Solidity: function verifyReceiptAndLog(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) receiptInfo)
func (_IReceiptVerifier *IReceiptVerifierTransactor) VerifyReceiptAndLog(opts *bind.TransactOpts, receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _IReceiptVerifier.contract.Transact(opts, "verifyReceiptAndLog", receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceiptAndLog is a paid mutator transaction binding the contract method 0x3996da7a.
//
// Solidity: function verifyReceiptAndLog(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) receiptInfo)
func (_IReceiptVerifier *IReceiptVerifierSession) VerifyReceiptAndLog(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _IReceiptVerifier.Contract.VerifyReceiptAndLog(&_IReceiptVerifier.TransactOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceiptAndLog is a paid mutator transaction binding the contract method 0x3996da7a.
//
// Solidity: function verifyReceiptAndLog(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) receiptInfo)
func (_IReceiptVerifier *IReceiptVerifierTransactorSession) VerifyReceiptAndLog(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _IReceiptVerifier.Contract.VerifyReceiptAndLog(&_IReceiptVerifier.TransactOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// ISMTMetaData contains all meta data concerning the ISMT contract.
var ISMTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"}],\"name\":\"isSmtRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newSmtRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"endBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextChunkMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"knowledgeProof\",\"type\":\"uint256[2]\"}],\"internalType\":\"structISMT.SmtUpdate\",\"name\":\"u\",\"type\":\"tuple\"}],\"name\":\"updateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISMTABI is the input ABI used to generate the binding from.
// Deprecated: Use ISMTMetaData.ABI instead.
var ISMTABI = ISMTMetaData.ABI

// ISMT is an auto generated Go binding around an Ethereum contract.
type ISMT struct {
	ISMTCaller     // Read-only binding to the contract
	ISMTTransactor // Write-only binding to the contract
	ISMTFilterer   // Log filterer for contract events
}

// ISMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISMTSession struct {
	Contract     *ISMT             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISMTCallerSession struct {
	Contract *ISMTCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISMTTransactorSession struct {
	Contract     *ISMTTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISMTRaw struct {
	Contract *ISMT // Generic contract binding to access the raw methods on
}

// ISMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISMTCallerRaw struct {
	Contract *ISMTCaller // Generic read-only contract binding to access the raw methods on
}

// ISMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISMTTransactorRaw struct {
	Contract *ISMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISMT creates a new instance of ISMT, bound to a specific deployed contract.
func NewISMT(address common.Address, backend bind.ContractBackend) (*ISMT, error) {
	contract, err := bindISMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISMT{ISMTCaller: ISMTCaller{contract: contract}, ISMTTransactor: ISMTTransactor{contract: contract}, ISMTFilterer: ISMTFilterer{contract: contract}}, nil
}

// NewISMTCaller creates a new read-only instance of ISMT, bound to a specific deployed contract.
func NewISMTCaller(address common.Address, caller bind.ContractCaller) (*ISMTCaller, error) {
	contract, err := bindISMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISMTCaller{contract: contract}, nil
}

// NewISMTTransactor creates a new write-only instance of ISMT, bound to a specific deployed contract.
func NewISMTTransactor(address common.Address, transactor bind.ContractTransactor) (*ISMTTransactor, error) {
	contract, err := bindISMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISMTTransactor{contract: contract}, nil
}

// NewISMTFilterer creates a new log filterer instance of ISMT, bound to a specific deployed contract.
func NewISMTFilterer(address common.Address, filterer bind.ContractFilterer) (*ISMTFilterer, error) {
	contract, err := bindISMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISMTFilterer{contract: contract}, nil
}

// bindISMT binds a generic wrapper to an already deployed contract.
func bindISMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISMT *ISMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISMT.Contract.ISMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISMT *ISMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISMT.Contract.ISMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISMT *ISMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISMT.Contract.ISMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISMT *ISMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISMT *ISMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISMT *ISMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISMT.Contract.contract.Transact(opts, method, params...)
}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_ISMT *ISMTCaller) IsSmtRootValid(opts *bind.CallOpts, chainId uint64, smtRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _ISMT.contract.Call(opts, &out, "isSmtRootValid", chainId, smtRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_ISMT *ISMTSession) IsSmtRootValid(chainId uint64, smtRoot [32]byte) (bool, error) {
	return _ISMT.Contract.IsSmtRootValid(&_ISMT.CallOpts, chainId, smtRoot)
}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_ISMT *ISMTCallerSession) IsSmtRootValid(chainId uint64, smtRoot [32]byte) (bool, error) {
	return _ISMT.Contract.IsSmtRootValid(&_ISMT.CallOpts, chainId, smtRoot)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_ISMT *ISMTTransactor) UpdateRoot(opts *bind.TransactOpts, chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _ISMT.contract.Transact(opts, "updateRoot", chainId, u)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_ISMT *ISMTSession) UpdateRoot(chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _ISMT.Contract.UpdateRoot(&_ISMT.TransactOpts, chainId, u)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_ISMT *ISMTTransactorSession) UpdateRoot(chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _ISMT.Contract.UpdateRoot(&_ISMT.TransactOpts, chainId, u)
}

// ISigsVerifierMetaData contains all meta data concerning the ISigsVerifier contract.
var ISigsVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISigsVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ISigsVerifierMetaData.ABI instead.
var ISigsVerifierABI = ISigsVerifierMetaData.ABI

// ISigsVerifier is an auto generated Go binding around an Ethereum contract.
type ISigsVerifier struct {
	ISigsVerifierCaller     // Read-only binding to the contract
	ISigsVerifierTransactor // Write-only binding to the contract
	ISigsVerifierFilterer   // Log filterer for contract events
}

// ISigsVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISigsVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISigsVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISigsVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISigsVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISigsVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISigsVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISigsVerifierSession struct {
	Contract     *ISigsVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISigsVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISigsVerifierCallerSession struct {
	Contract *ISigsVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISigsVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISigsVerifierTransactorSession struct {
	Contract     *ISigsVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISigsVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISigsVerifierRaw struct {
	Contract *ISigsVerifier // Generic contract binding to access the raw methods on
}

// ISigsVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISigsVerifierCallerRaw struct {
	Contract *ISigsVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ISigsVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISigsVerifierTransactorRaw struct {
	Contract *ISigsVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISigsVerifier creates a new instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifier(address common.Address, backend bind.ContractBackend) (*ISigsVerifier, error) {
	contract, err := bindISigsVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifier{ISigsVerifierCaller: ISigsVerifierCaller{contract: contract}, ISigsVerifierTransactor: ISigsVerifierTransactor{contract: contract}, ISigsVerifierFilterer: ISigsVerifierFilterer{contract: contract}}, nil
}

// NewISigsVerifierCaller creates a new read-only instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifierCaller(address common.Address, caller bind.ContractCaller) (*ISigsVerifierCaller, error) {
	contract, err := bindISigsVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifierCaller{contract: contract}, nil
}

// NewISigsVerifierTransactor creates a new write-only instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ISigsVerifierTransactor, error) {
	contract, err := bindISigsVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifierTransactor{contract: contract}, nil
}

// NewISigsVerifierFilterer creates a new log filterer instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ISigsVerifierFilterer, error) {
	contract, err := bindISigsVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifierFilterer{contract: contract}, nil
}

// bindISigsVerifier binds a generic wrapper to an already deployed contract.
func bindISigsVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISigsVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISigsVerifier *ISigsVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISigsVerifier.Contract.ISigsVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISigsVerifier *ISigsVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.ISigsVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISigsVerifier *ISigsVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.ISigsVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISigsVerifier *ISigsVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISigsVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISigsVerifier *ISigsVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISigsVerifier *ISigsVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ISigsVerifier *ISigsVerifierCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	var out []interface{}
	err := _ISigsVerifier.contract.Call(opts, &out, "verifySigs", _msg, _sigs, _signers, _powers)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ISigsVerifier *ISigsVerifierSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _ISigsVerifier.Contract.VerifySigs(&_ISigsVerifier.CallOpts, _msg, _sigs, _signers, _powers)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ISigsVerifier *ISigsVerifierCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _ISigsVerifier.Contract.VerifySigs(&_ISigsVerifier.CallOpts, _msg, _sigs, _signers, _powers)
}

// ISlotValueVerifierMetaData contains all meta data concerning the ISlotValueVerifier contract.
var ISlotValueVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifySlotValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"addrHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slotKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slotValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"}],\"internalType\":\"structISlotValueVerifier.SlotInfo\",\"name\":\"slotInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISlotValueVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlotValueVerifierMetaData.ABI instead.
var ISlotValueVerifierABI = ISlotValueVerifierMetaData.ABI

// ISlotValueVerifier is an auto generated Go binding around an Ethereum contract.
type ISlotValueVerifier struct {
	ISlotValueVerifierCaller     // Read-only binding to the contract
	ISlotValueVerifierTransactor // Write-only binding to the contract
	ISlotValueVerifierFilterer   // Log filterer for contract events
}

// ISlotValueVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlotValueVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlotValueVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlotValueVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlotValueVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlotValueVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlotValueVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlotValueVerifierSession struct {
	Contract     *ISlotValueVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISlotValueVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlotValueVerifierCallerSession struct {
	Contract *ISlotValueVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ISlotValueVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlotValueVerifierTransactorSession struct {
	Contract     *ISlotValueVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ISlotValueVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlotValueVerifierRaw struct {
	Contract *ISlotValueVerifier // Generic contract binding to access the raw methods on
}

// ISlotValueVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlotValueVerifierCallerRaw struct {
	Contract *ISlotValueVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ISlotValueVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlotValueVerifierTransactorRaw struct {
	Contract *ISlotValueVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlotValueVerifier creates a new instance of ISlotValueVerifier, bound to a specific deployed contract.
func NewISlotValueVerifier(address common.Address, backend bind.ContractBackend) (*ISlotValueVerifier, error) {
	contract, err := bindISlotValueVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlotValueVerifier{ISlotValueVerifierCaller: ISlotValueVerifierCaller{contract: contract}, ISlotValueVerifierTransactor: ISlotValueVerifierTransactor{contract: contract}, ISlotValueVerifierFilterer: ISlotValueVerifierFilterer{contract: contract}}, nil
}

// NewISlotValueVerifierCaller creates a new read-only instance of ISlotValueVerifier, bound to a specific deployed contract.
func NewISlotValueVerifierCaller(address common.Address, caller bind.ContractCaller) (*ISlotValueVerifierCaller, error) {
	contract, err := bindISlotValueVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlotValueVerifierCaller{contract: contract}, nil
}

// NewISlotValueVerifierTransactor creates a new write-only instance of ISlotValueVerifier, bound to a specific deployed contract.
func NewISlotValueVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlotValueVerifierTransactor, error) {
	contract, err := bindISlotValueVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlotValueVerifierTransactor{contract: contract}, nil
}

// NewISlotValueVerifierFilterer creates a new log filterer instance of ISlotValueVerifier, bound to a specific deployed contract.
func NewISlotValueVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlotValueVerifierFilterer, error) {
	contract, err := bindISlotValueVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlotValueVerifierFilterer{contract: contract}, nil
}

// bindISlotValueVerifier binds a generic wrapper to an already deployed contract.
func bindISlotValueVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlotValueVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlotValueVerifier *ISlotValueVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlotValueVerifier.Contract.ISlotValueVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlotValueVerifier *ISlotValueVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotValueVerifier.Contract.ISlotValueVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlotValueVerifier *ISlotValueVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlotValueVerifier.Contract.ISlotValueVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlotValueVerifier *ISlotValueVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlotValueVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlotValueVerifier *ISlotValueVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotValueVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlotValueVerifier *ISlotValueVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlotValueVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifySlotValue is a free data retrieval call binding the contract method 0x0b885e53.
//
// Solidity: function verifySlotValue(uint64 chainId, bytes proofData, bytes blkVerifyInfo) view returns((uint64,bytes32,bytes32,bytes32,bytes32,uint32) slotInfo)
func (_ISlotValueVerifier *ISlotValueVerifierCaller) VerifySlotValue(opts *bind.CallOpts, chainId uint64, proofData []byte, blkVerifyInfo []byte) (ISlotValueVerifierSlotInfo, error) {
	var out []interface{}
	err := _ISlotValueVerifier.contract.Call(opts, &out, "verifySlotValue", chainId, proofData, blkVerifyInfo)

	if err != nil {
		return *new(ISlotValueVerifierSlotInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlotValueVerifierSlotInfo)).(*ISlotValueVerifierSlotInfo)

	return out0, err

}

// VerifySlotValue is a free data retrieval call binding the contract method 0x0b885e53.
//
// Solidity: function verifySlotValue(uint64 chainId, bytes proofData, bytes blkVerifyInfo) view returns((uint64,bytes32,bytes32,bytes32,bytes32,uint32) slotInfo)
func (_ISlotValueVerifier *ISlotValueVerifierSession) VerifySlotValue(chainId uint64, proofData []byte, blkVerifyInfo []byte) (ISlotValueVerifierSlotInfo, error) {
	return _ISlotValueVerifier.Contract.VerifySlotValue(&_ISlotValueVerifier.CallOpts, chainId, proofData, blkVerifyInfo)
}

// VerifySlotValue is a free data retrieval call binding the contract method 0x0b885e53.
//
// Solidity: function verifySlotValue(uint64 chainId, bytes proofData, bytes blkVerifyInfo) view returns((uint64,bytes32,bytes32,bytes32,bytes32,uint32) slotInfo)
func (_ISlotValueVerifier *ISlotValueVerifierCallerSession) VerifySlotValue(chainId uint64, proofData []byte, blkVerifyInfo []byte) (ISlotValueVerifierSlotInfo, error) {
	return _ISlotValueVerifier.Contract.VerifySlotValue(&_ISlotValueVerifier.CallOpts, chainId, proofData, blkVerifyInfo)
}

// ITxVerifierMetaData contains all meta data concerning the ITxVerifier contract.
var ITxVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasTipCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFeeCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"}],\"internalType\":\"structITxVerifier.TxInfo\",\"name\":\"txInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyTxAndLog\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasTipCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFeeCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"}],\"internalType\":\"structITxVerifier.TxInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITxVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ITxVerifierMetaData.ABI instead.
var ITxVerifierABI = ITxVerifierMetaData.ABI

// ITxVerifier is an auto generated Go binding around an Ethereum contract.
type ITxVerifier struct {
	ITxVerifierCaller     // Read-only binding to the contract
	ITxVerifierTransactor // Write-only binding to the contract
	ITxVerifierFilterer   // Log filterer for contract events
}

// ITxVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITxVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITxVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITxVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITxVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITxVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITxVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITxVerifierSession struct {
	Contract     *ITxVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITxVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITxVerifierCallerSession struct {
	Contract *ITxVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ITxVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITxVerifierTransactorSession struct {
	Contract     *ITxVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ITxVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITxVerifierRaw struct {
	Contract *ITxVerifier // Generic contract binding to access the raw methods on
}

// ITxVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITxVerifierCallerRaw struct {
	Contract *ITxVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ITxVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITxVerifierTransactorRaw struct {
	Contract *ITxVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITxVerifier creates a new instance of ITxVerifier, bound to a specific deployed contract.
func NewITxVerifier(address common.Address, backend bind.ContractBackend) (*ITxVerifier, error) {
	contract, err := bindITxVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITxVerifier{ITxVerifierCaller: ITxVerifierCaller{contract: contract}, ITxVerifierTransactor: ITxVerifierTransactor{contract: contract}, ITxVerifierFilterer: ITxVerifierFilterer{contract: contract}}, nil
}

// NewITxVerifierCaller creates a new read-only instance of ITxVerifier, bound to a specific deployed contract.
func NewITxVerifierCaller(address common.Address, caller bind.ContractCaller) (*ITxVerifierCaller, error) {
	contract, err := bindITxVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITxVerifierCaller{contract: contract}, nil
}

// NewITxVerifierTransactor creates a new write-only instance of ITxVerifier, bound to a specific deployed contract.
func NewITxVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ITxVerifierTransactor, error) {
	contract, err := bindITxVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITxVerifierTransactor{contract: contract}, nil
}

// NewITxVerifierFilterer creates a new log filterer instance of ITxVerifier, bound to a specific deployed contract.
func NewITxVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ITxVerifierFilterer, error) {
	contract, err := bindITxVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITxVerifierFilterer{contract: contract}, nil
}

// bindITxVerifier binds a generic wrapper to an already deployed contract.
func bindITxVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITxVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITxVerifier *ITxVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITxVerifier.Contract.ITxVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITxVerifier *ITxVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITxVerifier.Contract.ITxVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITxVerifier *ITxVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITxVerifier.Contract.ITxVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITxVerifier *ITxVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITxVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITxVerifier *ITxVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITxVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITxVerifier *ITxVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITxVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyTx is a free data retrieval call binding the contract method 0xa8da8d69.
//
// Solidity: function verifyTx(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) txInfo)
func (_ITxVerifier *ITxVerifierCaller) VerifyTx(opts *bind.CallOpts, txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (ITxVerifierTxInfo, error) {
	var out []interface{}
	err := _ITxVerifier.contract.Call(opts, &out, "verifyTx", txRaw, proofData, auxiBlkVerifyInfo)

	if err != nil {
		return *new(ITxVerifierTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITxVerifierTxInfo)).(*ITxVerifierTxInfo)

	return out0, err

}

// VerifyTx is a free data retrieval call binding the contract method 0xa8da8d69.
//
// Solidity: function verifyTx(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) txInfo)
func (_ITxVerifier *ITxVerifierSession) VerifyTx(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (ITxVerifierTxInfo, error) {
	return _ITxVerifier.Contract.VerifyTx(&_ITxVerifier.CallOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTx is a free data retrieval call binding the contract method 0xa8da8d69.
//
// Solidity: function verifyTx(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) txInfo)
func (_ITxVerifier *ITxVerifierCallerSession) VerifyTx(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (ITxVerifierTxInfo, error) {
	return _ITxVerifier.Contract.VerifyTx(&_ITxVerifier.CallOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTxAndLog is a paid mutator transaction binding the contract method 0x361108de.
//
// Solidity: function verifyTxAndLog(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_ITxVerifier *ITxVerifierTransactor) VerifyTxAndLog(opts *bind.TransactOpts, txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _ITxVerifier.contract.Transact(opts, "verifyTxAndLog", txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTxAndLog is a paid mutator transaction binding the contract method 0x361108de.
//
// Solidity: function verifyTxAndLog(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_ITxVerifier *ITxVerifierSession) VerifyTxAndLog(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _ITxVerifier.Contract.VerifyTxAndLog(&_ITxVerifier.TransactOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTxAndLog is a paid mutator transaction binding the contract method 0x361108de.
//
// Solidity: function verifyTxAndLog(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_ITxVerifier *ITxVerifierTransactorSession) VerifyTxAndLog(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _ITxVerifier.Contract.VerifyTxAndLog(&_ITxVerifier.TransactOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// IVerifierMetaData contains all meta data concerning the IVerifier contract.
var IVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"knowledgeProof\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IVerifierMetaData.ABI instead.
var IVerifierABI = IVerifierMetaData.ABI

// IVerifier is an auto generated Go binding around an Ethereum contract.
type IVerifier struct {
	IVerifierCaller     // Read-only binding to the contract
	IVerifierTransactor // Write-only binding to the contract
	IVerifierFilterer   // Log filterer for contract events
}

// IVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVerifierSession struct {
	Contract     *IVerifier        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVerifierCallerSession struct {
	Contract *IVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVerifierTransactorSession struct {
	Contract     *IVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVerifierRaw struct {
	Contract *IVerifier // Generic contract binding to access the raw methods on
}

// IVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVerifierCallerRaw struct {
	Contract *IVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVerifierTransactorRaw struct {
	Contract *IVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVerifier creates a new instance of IVerifier, bound to a specific deployed contract.
func NewIVerifier(address common.Address, backend bind.ContractBackend) (*IVerifier, error) {
	contract, err := bindIVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVerifier{IVerifierCaller: IVerifierCaller{contract: contract}, IVerifierTransactor: IVerifierTransactor{contract: contract}, IVerifierFilterer: IVerifierFilterer{contract: contract}}, nil
}

// NewIVerifierCaller creates a new read-only instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierCaller(address common.Address, caller bind.ContractCaller) (*IVerifierCaller, error) {
	contract, err := bindIVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierCaller{contract: contract}, nil
}

// NewIVerifierTransactor creates a new write-only instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IVerifierTransactor, error) {
	contract, err := bindIVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierTransactor{contract: contract}, nil
}

// NewIVerifierFilterer creates a new log filterer instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IVerifierFilterer, error) {
	contract, err := bindIVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVerifierFilterer{contract: contract}, nil
}

// bindIVerifier binds a generic wrapper to an already deployed contract.
func bindIVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifier *IVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifier.Contract.IVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifier *IVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifier.Contract.IVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifier *IVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifier.Contract.IVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifier *IVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifier *IVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifier *IVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commit, uint256[2] knowledgeProof, uint256[9] input) view returns(bool r)
func (_IVerifier *IVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commit [2]*big.Int, knowledgeProof [2]*big.Int, input [9]*big.Int) (bool, error) {
	var out []interface{}
	err := _IVerifier.contract.Call(opts, &out, "verifyProof", proof, commit, knowledgeProof, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commit, uint256[2] knowledgeProof, uint256[9] input) view returns(bool r)
func (_IVerifier *IVerifierSession) VerifyProof(proof [8]*big.Int, commit [2]*big.Int, knowledgeProof [2]*big.Int, input [9]*big.Int) (bool, error) {
	return _IVerifier.Contract.VerifyProof(&_IVerifier.CallOpts, proof, commit, knowledgeProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commit, uint256[2] knowledgeProof, uint256[9] input) view returns(bool r)
func (_IVerifier *IVerifierCallerSession) VerifyProof(proof [8]*big.Int, commit [2]*big.Int, knowledgeProof [2]*big.Int, input [9]*big.Int) (bool, error) {
	return _IVerifier.Contract.VerifyProof(&_IVerifier.CallOpts, proof, commit, knowledgeProof, input)
}

// IZkpVerifierMetaData contains all meta data concerning the IZkpVerifier contract.
var IZkpVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"}],\"name\":\"verifyRaw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IZkpVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IZkpVerifierMetaData.ABI instead.
var IZkpVerifierABI = IZkpVerifierMetaData.ABI

// IZkpVerifier is an auto generated Go binding around an Ethereum contract.
type IZkpVerifier struct {
	IZkpVerifierCaller     // Read-only binding to the contract
	IZkpVerifierTransactor // Write-only binding to the contract
	IZkpVerifierFilterer   // Log filterer for contract events
}

// IZkpVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZkpVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkpVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZkpVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkpVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZkpVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkpVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZkpVerifierSession struct {
	Contract     *IZkpVerifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZkpVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZkpVerifierCallerSession struct {
	Contract *IZkpVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IZkpVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZkpVerifierTransactorSession struct {
	Contract     *IZkpVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IZkpVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZkpVerifierRaw struct {
	Contract *IZkpVerifier // Generic contract binding to access the raw methods on
}

// IZkpVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZkpVerifierCallerRaw struct {
	Contract *IZkpVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IZkpVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZkpVerifierTransactorRaw struct {
	Contract *IZkpVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZkpVerifier creates a new instance of IZkpVerifier, bound to a specific deployed contract.
func NewIZkpVerifier(address common.Address, backend bind.ContractBackend) (*IZkpVerifier, error) {
	contract, err := bindIZkpVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZkpVerifier{IZkpVerifierCaller: IZkpVerifierCaller{contract: contract}, IZkpVerifierTransactor: IZkpVerifierTransactor{contract: contract}, IZkpVerifierFilterer: IZkpVerifierFilterer{contract: contract}}, nil
}

// NewIZkpVerifierCaller creates a new read-only instance of IZkpVerifier, bound to a specific deployed contract.
func NewIZkpVerifierCaller(address common.Address, caller bind.ContractCaller) (*IZkpVerifierCaller, error) {
	contract, err := bindIZkpVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZkpVerifierCaller{contract: contract}, nil
}

// NewIZkpVerifierTransactor creates a new write-only instance of IZkpVerifier, bound to a specific deployed contract.
func NewIZkpVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IZkpVerifierTransactor, error) {
	contract, err := bindIZkpVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZkpVerifierTransactor{contract: contract}, nil
}

// NewIZkpVerifierFilterer creates a new log filterer instance of IZkpVerifier, bound to a specific deployed contract.
func NewIZkpVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IZkpVerifierFilterer, error) {
	contract, err := bindIZkpVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZkpVerifierFilterer{contract: contract}, nil
}

// bindIZkpVerifier binds a generic wrapper to an already deployed contract.
func bindIZkpVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZkpVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkpVerifier *IZkpVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkpVerifier.Contract.IZkpVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkpVerifier *IZkpVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkpVerifier.Contract.IZkpVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkpVerifier *IZkpVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkpVerifier.Contract.IZkpVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkpVerifier *IZkpVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkpVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkpVerifier *IZkpVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkpVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkpVerifier *IZkpVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkpVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyRaw is a free data retrieval call binding the contract method 0x457726e6.
//
// Solidity: function verifyRaw(bytes proofData) view returns(bool r)
func (_IZkpVerifier *IZkpVerifierCaller) VerifyRaw(opts *bind.CallOpts, proofData []byte) (bool, error) {
	var out []interface{}
	err := _IZkpVerifier.contract.Call(opts, &out, "verifyRaw", proofData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRaw is a free data retrieval call binding the contract method 0x457726e6.
//
// Solidity: function verifyRaw(bytes proofData) view returns(bool r)
func (_IZkpVerifier *IZkpVerifierSession) VerifyRaw(proofData []byte) (bool, error) {
	return _IZkpVerifier.Contract.VerifyRaw(&_IZkpVerifier.CallOpts, proofData)
}

// VerifyRaw is a free data retrieval call binding the contract method 0x457726e6.
//
// Solidity: function verifyRaw(bytes proofData) view returns(bool r)
func (_IZkpVerifier *IZkpVerifierCallerSession) VerifyRaw(proofData []byte) (bool, error) {
	return _IZkpVerifier.Contract.VerifyRaw(&_IZkpVerifier.CallOpts, proofData)
}

// LightClientStoreMetaData contains all meta data concerning the LightClientStore contract.
var LightClientStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"bestValidUpdate\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"attestedHeader\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"beacon\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"stateRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"blockNumber\",\"type\":\"tuple\"}],\"internalType\":\"structExecutionPayload\",\"name\":\"execution\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"branch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structLeafWithBranch\",\"name\":\"executionRoot\",\"type\":\"tuple\"}],\"internalType\":\"structHeaderWithExecution\",\"name\":\"finalizedHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"nextSyncCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextSyncCommitteePoseidonRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"nextSyncCommitteeRootMappingProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"participation\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"poseidonRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitment\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBeaconVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structSyncAggregate\",\"name\":\"syncAggregate\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"signatureSlot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSyncCommitteePoseidonRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSyncCommitteeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedExecutionStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forkEpochs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forkVersions\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSyncCommitteePoseidonRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSyncCommitteeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticExecutionStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkVerifier\",\"outputs\":[{\"internalType\":\"contractIBeaconVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LightClientStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use LightClientStoreMetaData.ABI instead.
var LightClientStoreABI = LightClientStoreMetaData.ABI

// LightClientStore is an auto generated Go binding around an Ethereum contract.
type LightClientStore struct {
	LightClientStoreCaller     // Read-only binding to the contract
	LightClientStoreTransactor // Write-only binding to the contract
	LightClientStoreFilterer   // Log filterer for contract events
}

// LightClientStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightClientStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightClientStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightClientStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightClientStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightClientStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightClientStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightClientStoreSession struct {
	Contract     *LightClientStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightClientStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightClientStoreCallerSession struct {
	Contract *LightClientStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LightClientStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightClientStoreTransactorSession struct {
	Contract     *LightClientStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LightClientStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightClientStoreRaw struct {
	Contract *LightClientStore // Generic contract binding to access the raw methods on
}

// LightClientStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightClientStoreCallerRaw struct {
	Contract *LightClientStoreCaller // Generic read-only contract binding to access the raw methods on
}

// LightClientStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightClientStoreTransactorRaw struct {
	Contract *LightClientStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightClientStore creates a new instance of LightClientStore, bound to a specific deployed contract.
func NewLightClientStore(address common.Address, backend bind.ContractBackend) (*LightClientStore, error) {
	contract, err := bindLightClientStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightClientStore{LightClientStoreCaller: LightClientStoreCaller{contract: contract}, LightClientStoreTransactor: LightClientStoreTransactor{contract: contract}, LightClientStoreFilterer: LightClientStoreFilterer{contract: contract}}, nil
}

// NewLightClientStoreCaller creates a new read-only instance of LightClientStore, bound to a specific deployed contract.
func NewLightClientStoreCaller(address common.Address, caller bind.ContractCaller) (*LightClientStoreCaller, error) {
	contract, err := bindLightClientStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightClientStoreCaller{contract: contract}, nil
}

// NewLightClientStoreTransactor creates a new write-only instance of LightClientStore, bound to a specific deployed contract.
func NewLightClientStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*LightClientStoreTransactor, error) {
	contract, err := bindLightClientStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightClientStoreTransactor{contract: contract}, nil
}

// NewLightClientStoreFilterer creates a new log filterer instance of LightClientStore, bound to a specific deployed contract.
func NewLightClientStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*LightClientStoreFilterer, error) {
	contract, err := bindLightClientStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightClientStoreFilterer{contract: contract}, nil
}

// bindLightClientStore binds a generic wrapper to an already deployed contract.
func bindLightClientStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightClientStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightClientStore *LightClientStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightClientStore.Contract.LightClientStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightClientStore *LightClientStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightClientStore.Contract.LightClientStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightClientStore *LightClientStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightClientStore.Contract.LightClientStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightClientStore *LightClientStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightClientStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightClientStore *LightClientStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightClientStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightClientStore *LightClientStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightClientStore.Contract.contract.Transact(opts, method, params...)
}

// BestValidUpdate is a free data retrieval call binding the contract method 0xba67ee48.
//
// Solidity: function bestValidUpdate() view returns(((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) attestedHeader, ((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) finalizedHeader, bytes32 nextSyncCommitteeRoot, bytes32 nextSyncCommitteePoseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) nextSyncCommitteeRootMappingProof, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate, uint64 signatureSlot)
func (_LightClientStore *LightClientStoreCaller) BestValidUpdate(opts *bind.CallOpts) (struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "bestValidUpdate")

	outstruct := new(struct {
		AttestedHeader                    HeaderWithExecution
		FinalizedHeader                   HeaderWithExecution
		NextSyncCommitteeRoot             [32]byte
		NextSyncCommitteePoseidonRoot     [32]byte
		NextSyncCommitteeRootMappingProof IBeaconVerifierProof
		SyncAggregate                     SyncAggregate
		SignatureSlot                     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestedHeader = *abi.ConvertType(out[0], new(HeaderWithExecution)).(*HeaderWithExecution)
	outstruct.FinalizedHeader = *abi.ConvertType(out[1], new(HeaderWithExecution)).(*HeaderWithExecution)
	outstruct.NextSyncCommitteeRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.NextSyncCommitteePoseidonRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.NextSyncCommitteeRootMappingProof = *abi.ConvertType(out[4], new(IBeaconVerifierProof)).(*IBeaconVerifierProof)
	outstruct.SyncAggregate = *abi.ConvertType(out[5], new(SyncAggregate)).(*SyncAggregate)
	outstruct.SignatureSlot = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// BestValidUpdate is a free data retrieval call binding the contract method 0xba67ee48.
//
// Solidity: function bestValidUpdate() view returns(((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) attestedHeader, ((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) finalizedHeader, bytes32 nextSyncCommitteeRoot, bytes32 nextSyncCommitteePoseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) nextSyncCommitteeRootMappingProof, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate, uint64 signatureSlot)
func (_LightClientStore *LightClientStoreSession) BestValidUpdate() (struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}, error) {
	return _LightClientStore.Contract.BestValidUpdate(&_LightClientStore.CallOpts)
}

// BestValidUpdate is a free data retrieval call binding the contract method 0xba67ee48.
//
// Solidity: function bestValidUpdate() view returns(((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) attestedHeader, ((uint64,uint64,bytes32,bytes32,bytes32),((bytes32,bytes32[]),(bytes32,bytes32[]),(bytes32,bytes32[])),(bytes32,bytes32[])) finalizedHeader, bytes32 nextSyncCommitteeRoot, bytes32 nextSyncCommitteePoseidonRoot, (uint256[2],uint256[2][2],uint256[2],uint256[2]) nextSyncCommitteeRootMappingProof, (uint64,bytes32,uint256,(uint256[2],uint256[2][2],uint256[2],uint256[2])) syncAggregate, uint64 signatureSlot)
func (_LightClientStore *LightClientStoreCallerSession) BestValidUpdate() (struct {
	AttestedHeader                    HeaderWithExecution
	FinalizedHeader                   HeaderWithExecution
	NextSyncCommitteeRoot             [32]byte
	NextSyncCommitteePoseidonRoot     [32]byte
	NextSyncCommitteeRootMappingProof IBeaconVerifierProof
	SyncAggregate                     SyncAggregate
	SignatureSlot                     uint64
}, error) {
	return _LightClientStore.Contract.BestValidUpdate(&_LightClientStore.CallOpts)
}

// CurrentSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0x65e700de.
//
// Solidity: function currentSyncCommitteePoseidonRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCaller) CurrentSyncCommitteePoseidonRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "currentSyncCommitteePoseidonRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0x65e700de.
//
// Solidity: function currentSyncCommitteePoseidonRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreSession) CurrentSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _LightClientStore.Contract.CurrentSyncCommitteePoseidonRoot(&_LightClientStore.CallOpts)
}

// CurrentSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0x65e700de.
//
// Solidity: function currentSyncCommitteePoseidonRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCallerSession) CurrentSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _LightClientStore.Contract.CurrentSyncCommitteePoseidonRoot(&_LightClientStore.CallOpts)
}

// CurrentSyncCommitteeRoot is a free data retrieval call binding the contract method 0xa4059e07.
//
// Solidity: function currentSyncCommitteeRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCaller) CurrentSyncCommitteeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "currentSyncCommitteeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentSyncCommitteeRoot is a free data retrieval call binding the contract method 0xa4059e07.
//
// Solidity: function currentSyncCommitteeRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreSession) CurrentSyncCommitteeRoot() ([32]byte, error) {
	return _LightClientStore.Contract.CurrentSyncCommitteeRoot(&_LightClientStore.CallOpts)
}

// CurrentSyncCommitteeRoot is a free data retrieval call binding the contract method 0xa4059e07.
//
// Solidity: function currentSyncCommitteeRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCallerSession) CurrentSyncCommitteeRoot() ([32]byte, error) {
	return _LightClientStore.Contract.CurrentSyncCommitteeRoot(&_LightClientStore.CallOpts)
}

// FinalizedExecutionStateRoot is a free data retrieval call binding the contract method 0xc5190436.
//
// Solidity: function finalizedExecutionStateRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCaller) FinalizedExecutionStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "finalizedExecutionStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedExecutionStateRoot is a free data retrieval call binding the contract method 0xc5190436.
//
// Solidity: function finalizedExecutionStateRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreSession) FinalizedExecutionStateRoot() ([32]byte, error) {
	return _LightClientStore.Contract.FinalizedExecutionStateRoot(&_LightClientStore.CallOpts)
}

// FinalizedExecutionStateRoot is a free data retrieval call binding the contract method 0xc5190436.
//
// Solidity: function finalizedExecutionStateRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCallerSession) FinalizedExecutionStateRoot() ([32]byte, error) {
	return _LightClientStore.Contract.FinalizedExecutionStateRoot(&_LightClientStore.CallOpts)
}

// FinalizedSlot is a free data retrieval call binding the contract method 0xd1802369.
//
// Solidity: function finalizedSlot() view returns(uint64)
func (_LightClientStore *LightClientStoreCaller) FinalizedSlot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "finalizedSlot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FinalizedSlot is a free data retrieval call binding the contract method 0xd1802369.
//
// Solidity: function finalizedSlot() view returns(uint64)
func (_LightClientStore *LightClientStoreSession) FinalizedSlot() (uint64, error) {
	return _LightClientStore.Contract.FinalizedSlot(&_LightClientStore.CallOpts)
}

// FinalizedSlot is a free data retrieval call binding the contract method 0xd1802369.
//
// Solidity: function finalizedSlot() view returns(uint64)
func (_LightClientStore *LightClientStoreCallerSession) FinalizedSlot() (uint64, error) {
	return _LightClientStore.Contract.FinalizedSlot(&_LightClientStore.CallOpts)
}

// ForkEpochs is a free data retrieval call binding the contract method 0xbcbaf770.
//
// Solidity: function forkEpochs(uint256 ) view returns(uint64)
func (_LightClientStore *LightClientStoreCaller) ForkEpochs(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "forkEpochs", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkEpochs is a free data retrieval call binding the contract method 0xbcbaf770.
//
// Solidity: function forkEpochs(uint256 ) view returns(uint64)
func (_LightClientStore *LightClientStoreSession) ForkEpochs(arg0 *big.Int) (uint64, error) {
	return _LightClientStore.Contract.ForkEpochs(&_LightClientStore.CallOpts, arg0)
}

// ForkEpochs is a free data retrieval call binding the contract method 0xbcbaf770.
//
// Solidity: function forkEpochs(uint256 ) view returns(uint64)
func (_LightClientStore *LightClientStoreCallerSession) ForkEpochs(arg0 *big.Int) (uint64, error) {
	return _LightClientStore.Contract.ForkEpochs(&_LightClientStore.CallOpts, arg0)
}

// ForkVersions is a free data retrieval call binding the contract method 0xbaa94ea2.
//
// Solidity: function forkVersions(uint256 ) view returns(bytes4)
func (_LightClientStore *LightClientStoreCaller) ForkVersions(opts *bind.CallOpts, arg0 *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "forkVersions", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ForkVersions is a free data retrieval call binding the contract method 0xbaa94ea2.
//
// Solidity: function forkVersions(uint256 ) view returns(bytes4)
func (_LightClientStore *LightClientStoreSession) ForkVersions(arg0 *big.Int) ([4]byte, error) {
	return _LightClientStore.Contract.ForkVersions(&_LightClientStore.CallOpts, arg0)
}

// ForkVersions is a free data retrieval call binding the contract method 0xbaa94ea2.
//
// Solidity: function forkVersions(uint256 ) view returns(bytes4)
func (_LightClientStore *LightClientStoreCallerSession) ForkVersions(arg0 *big.Int) ([4]byte, error) {
	return _LightClientStore.Contract.ForkVersions(&_LightClientStore.CallOpts, arg0)
}

// NextSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0xe1861b08.
//
// Solidity: function nextSyncCommitteePoseidonRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCaller) NextSyncCommitteePoseidonRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "nextSyncCommitteePoseidonRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0xe1861b08.
//
// Solidity: function nextSyncCommitteePoseidonRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreSession) NextSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _LightClientStore.Contract.NextSyncCommitteePoseidonRoot(&_LightClientStore.CallOpts)
}

// NextSyncCommitteePoseidonRoot is a free data retrieval call binding the contract method 0xe1861b08.
//
// Solidity: function nextSyncCommitteePoseidonRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCallerSession) NextSyncCommitteePoseidonRoot() ([32]byte, error) {
	return _LightClientStore.Contract.NextSyncCommitteePoseidonRoot(&_LightClientStore.CallOpts)
}

// NextSyncCommitteeRoot is a free data retrieval call binding the contract method 0x67b49cc7.
//
// Solidity: function nextSyncCommitteeRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCaller) NextSyncCommitteeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "nextSyncCommitteeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextSyncCommitteeRoot is a free data retrieval call binding the contract method 0x67b49cc7.
//
// Solidity: function nextSyncCommitteeRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreSession) NextSyncCommitteeRoot() ([32]byte, error) {
	return _LightClientStore.Contract.NextSyncCommitteeRoot(&_LightClientStore.CallOpts)
}

// NextSyncCommitteeRoot is a free data retrieval call binding the contract method 0x67b49cc7.
//
// Solidity: function nextSyncCommitteeRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCallerSession) NextSyncCommitteeRoot() ([32]byte, error) {
	return _LightClientStore.Contract.NextSyncCommitteeRoot(&_LightClientStore.CallOpts)
}

// OptimisticExecutionStateRoot is a free data retrieval call binding the contract method 0x39536c8f.
//
// Solidity: function optimisticExecutionStateRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCaller) OptimisticExecutionStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "optimisticExecutionStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OptimisticExecutionStateRoot is a free data retrieval call binding the contract method 0x39536c8f.
//
// Solidity: function optimisticExecutionStateRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreSession) OptimisticExecutionStateRoot() ([32]byte, error) {
	return _LightClientStore.Contract.OptimisticExecutionStateRoot(&_LightClientStore.CallOpts)
}

// OptimisticExecutionStateRoot is a free data retrieval call binding the contract method 0x39536c8f.
//
// Solidity: function optimisticExecutionStateRoot() view returns(bytes32)
func (_LightClientStore *LightClientStoreCallerSession) OptimisticExecutionStateRoot() ([32]byte, error) {
	return _LightClientStore.Contract.OptimisticExecutionStateRoot(&_LightClientStore.CallOpts)
}

// OptimisticSlot is a free data retrieval call binding the contract method 0x3cf5ea9e.
//
// Solidity: function optimisticSlot() view returns(uint64)
func (_LightClientStore *LightClientStoreCaller) OptimisticSlot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "optimisticSlot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OptimisticSlot is a free data retrieval call binding the contract method 0x3cf5ea9e.
//
// Solidity: function optimisticSlot() view returns(uint64)
func (_LightClientStore *LightClientStoreSession) OptimisticSlot() (uint64, error) {
	return _LightClientStore.Contract.OptimisticSlot(&_LightClientStore.CallOpts)
}

// OptimisticSlot is a free data retrieval call binding the contract method 0x3cf5ea9e.
//
// Solidity: function optimisticSlot() view returns(uint64)
func (_LightClientStore *LightClientStoreCallerSession) OptimisticSlot() (uint64, error) {
	return _LightClientStore.Contract.OptimisticSlot(&_LightClientStore.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_LightClientStore *LightClientStoreCaller) ZkVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightClientStore.contract.Call(opts, &out, "zkVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_LightClientStore *LightClientStoreSession) ZkVerifier() (common.Address, error) {
	return _LightClientStore.Contract.ZkVerifier(&_LightClientStore.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_LightClientStore *LightClientStoreCallerSession) ZkVerifier() (common.Address, error) {
	return _LightClientStore.Contract.ZkVerifier(&_LightClientStore.CallOpts)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairingMetaData contains all meta data concerning the Pairing contract.
var PairingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea264697066735822122008b5dc48e05bf7f7f7240b47fe3611cea1ed078a366ec0239182386e6a59575664736f6c63430008140033",
}

// PairingABI is the input ABI used to generate the binding from.
// Deprecated: Use PairingMetaData.ABI instead.
var PairingABI = PairingMetaData.ABI

// PairingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PairingMetaData.Bin instead.
var PairingBin = PairingMetaData.Bin

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// RLPReaderMetaData contains all meta data concerning the RLPReader contract.
var RLPReaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea264697066735822122085da79539df330f84295e55f40fba7b9a17c7398871a75704eb00e47f3e92a0d64736f6c63430008140033",
}

// RLPReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPReaderMetaData.ABI instead.
var RLPReaderABI = RLPReaderMetaData.ABI

// RLPReaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RLPReaderMetaData.Bin instead.
var RLPReaderBin = RLPReaderMetaData.Bin

// DeployRLPReader deploys a new Ethereum contract, binding an instance of RLPReader to it.
func DeployRLPReader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPReader, error) {
	parsed, err := RLPReaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RLPReaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// RLPReader is an auto generated Go binding around an Ethereum contract.
type RLPReader struct {
	RLPReaderCaller     // Read-only binding to the contract
	RLPReaderTransactor // Write-only binding to the contract
	RLPReaderFilterer   // Log filterer for contract events
}

// RLPReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPReaderSession struct {
	Contract     *RLPReader        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPReaderCallerSession struct {
	Contract *RLPReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPReaderTransactorSession struct {
	Contract     *RLPReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPReaderRaw struct {
	Contract *RLPReader // Generic contract binding to access the raw methods on
}

// RLPReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPReaderCallerRaw struct {
	Contract *RLPReaderCaller // Generic read-only contract binding to access the raw methods on
}

// RLPReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPReaderTransactorRaw struct {
	Contract *RLPReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPReader creates a new instance of RLPReader, bound to a specific deployed contract.
func NewRLPReader(address common.Address, backend bind.ContractBackend) (*RLPReader, error) {
	contract, err := bindRLPReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// NewRLPReaderCaller creates a new read-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderCaller(address common.Address, caller bind.ContractCaller) (*RLPReaderCaller, error) {
	contract, err := bindRLPReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderCaller{contract: contract}, nil
}

// NewRLPReaderTransactor creates a new write-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPReaderTransactor, error) {
	contract, err := bindRLPReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderTransactor{contract: contract}, nil
}

// NewRLPReaderFilterer creates a new log filterer instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPReaderFilterer, error) {
	contract, err := bindRLPReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPReaderFilterer{contract: contract}, nil
}

// bindRLPReader binds a generic wrapper to an already deployed contract.
func bindRLPReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.RLPReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transact(opts, method, params...)
}

// ReceiptVerifierMetaData contains all meta data concerning the ReceiptVerifier contract.
var ReceiptVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockChunks\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateBlockChunks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateVerifierAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"}],\"name\":\"VerifiedReceipt\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockChunks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptRaw\",\"type\":\"bytes\"}],\"name\":\"decodeReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIReceiptVerifier.LogInfo[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIReceiptVerifier.ReceiptInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockChunks\",\"type\":\"address\"}],\"name\":\"updateBlockChunks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"}],\"name\":\"updateVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"verifierAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIReceiptVerifier.LogInfo[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIReceiptVerifier.ReceiptInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptAndLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIReceiptVerifier.LogInfo[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIReceiptVerifier.ReceiptInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346100a457601f611a4b38819003918201601f19168301916001600160401b038311848410176100a8578084926020946040528339810103126100a457516001600160a01b0390818116908190036100a4575f5460018060a01b03199033828216175f55604051933391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3600254161760025561198e90816100bd8239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f803560e01c9081631eeb86da146100b7575080633996da7a146100b257806368ac2f78146100ad578063715018a6146100a8578063724796ed146100a35780638da5cb5b1461009e578063c109ae5d14610099578063ec4ffc5214610094578063f2fde38b1461008f5763f5cec6af1461008a575f80fd5b61062a565b61057e565b6104e9565b610495565b610470565b61044a565b6103ed565b6103c1565b610333565b34610127576020366003190112610127577f0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a9860206100f361012a565b6001600160a01b039061010a82865416331461066c565b16806001600160a01b03196002541617600255604051908152a180f35b80fd5b600435906001600160a01b038216820361014057565b5f80fd5b9181601f840112156101405782359167ffffffffffffffff8311610140576020838186019501011161014057565b9060606003198301126101405767ffffffffffffffff600435818111610140578361019f91600401610144565b9390939260243583811161014057826101ba91600401610144565b93909392604435918211610140576101d491600401610144565b9091565b5f5b8381106101e95750505f910152565b81810151838201526020016101da565b602080825260e0820190835115158184015267ffffffffffffffff928382860151166040908183015260a081870151966060978885015263ffffffff8882015116966080978886015287820151168285015201519560c08084015286518095526101009184838501938760051b8601019801965f9081955b888710610285575050505050505050505090565b9091929394959697989960ff19828203018752888b51878301926001600160a01b0382511681528783830151928a8584015283518096528488840194019589905b808210610316575050600195509184939184930151908a8184039101526102f8815180928185528580860191016101d8565b601f01601f191601019c019a99989190910196019493929190610271565b87518652968601968f9695909501948b93506001909101906102c6565b34610140576103bd61039f7fa5db3bb7a25cc2804c7835ab71b15513b8c68585e3593c3fdee39a8837547366604061038761036d36610172565b9261037f99959996949692919261073d565b50868a610988565b9467ffffffffffffffff60208701511692369161079a565b6020815191012082519182526020820152a1604051918291826101f9565b0390f35b34610140576103bd6103e16103d536610172565b94939093929192610988565b604051918291826101f9565b34610140575f80600319360112610127578080546001600160a01b03196001600160a01b0382169161042033841461066c565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b34610140575f3660031901126101405760206001600160a01b0360025416604051908152f35b34610140575f3660031901126101405760206001600160a01b035f5416604051908152f35b346101405760203660031901126101405760043567ffffffffffffffff8111610140576103e16104cc6103bd923690600401610144565b90611039565b6004359067ffffffffffffffff8216820361014057565b34610140576040366003190112610140576105026104d2565b602435906001600160a01b03918281168091036101405767ffffffffffffffff6040926105547ffd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f955f5416331461066c565b1690815f526001602052825f20816001600160a01b031982541617905582519182526020820152a1005b346101405760203660031901126101405761059761012a565b6001600160a01b036105ad815f5416331461066c565b8116156105bf576105bd906114c2565b005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b346101405760203660031901126101405767ffffffffffffffff61064c6104d2565b165f52600160205260206001600160a01b0360405f205416604051908152f35b1561067357565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff8211176106e757604052565b6106b7565b90601f8019910116810190811067ffffffffffffffff8211176106e757604052565b6040519060c0820182811067ffffffffffffffff8211176106e757604052565b6040519061073b826106cb565b565b6040519060c0820182811067ffffffffffffffff8211176106e757604052606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b67ffffffffffffffff81116106e757601f01601f191660200190565b9291926107a68261077e565b916107b460405193846106ec565b829481845281830111610140578281602093845f960137010152565b156107d757565b60405162461bcd60e51b815260206004820152600f60248201527f70726f6f66206e6f742076616c696400000000000000000000000000000000006044820152606490fd5b6020908361073b93959495604051968361083f89955180928880890191016101d8565b84019185830137015f838201520380855201836106ec565b1561085e57565b60405162461bcd60e51b815260206004820152601260248201527f6c65616648617368206e6f74206d6174636800000000000000000000000000006044820152606490fd5b90816020910312610140575180151581036101405790565b91909161018081019267ffffffffffffffff815116825260a08063ffffffff926020938085830151168587015260408201516040870152606082015160608701526080820151166080860152015192015f905b6007821061091c5750505050565b8280600192865181520194019101909261090e565b6040513d5f823e3d90fd5b1561094357565b60405162461bcd60e51b815260206004820152600f60248201527f696e76616c696420626c6b4861736800000000000000000000000000000000006044820152606490fd5b9594610aae9493929161099961073d565b506109cc6109c76109aa86866113b6565b9560808701956109c2875167ffffffffffffffff1690565b61124c565b6107d0565b6109fa6109de828a60a088015161081c565b928351966109f56020988980970120885114610857565b610dac565b85519294919267ffffffffffffffff1690610a6a6040890194610a21865163ffffffff1690565b97858b0198610a538a5191610a47610a3761070e565b67ffffffffffffffff9099168952565b63ffffffff1687890152565b6040860152606085015263ffffffff166080840152565b60a0820152610a93610a87610a876002546001600160a01b031690565b6001600160a01b031690565b60405180809b8194631513dce960e21b8352600483016108bb565b03915afa918215610b7a57610b076060610b3d97610b2f96610af1610b189560409f610aec61073b9f9a8f9b610b229c5f92610b4d575b505061093c565b611039565b9d8e9151910152015167ffffffffffffffff1690565b67ffffffffffffffff1660808b0152565b5163ffffffff1690565b63ffffffff166060880152565b5167ffffffffffffffff1690565b67ffffffffffffffff1690840152565b610b6c9250803d10610b73575b610b6481836106ec565b8101906108a3565b8f80610ae5565b503d610b5a565b610931565b15610b8657565b60405162461bcd60e51b815260206004820152601b60248201527f696e636f72726563742061757869426c6b566572696679496e666f00000000006044820152606490fd5b906020116101405790602090565b906024116101405760200190600490565b90610104116101405760e40190602090565b909291928360011161014057831161014057600101915f190190565b9061018011610140576101700190601090565b906101a011610140576101800190602090565b906101c011610140576101b00190601090565b906101e011610140576101dc0190600490565b9061020011610140576101f80190600890565b9061020811610140576102000190600890565b92919261020891848311610140578411610140570191610207190190565b90939293848311610140578411610140578101920390565b359060208110610cce575090565b5f199060200360031b1b1690565b7fffffffff000000000000000000000000000000000000000000000000000000009035818116939260048110610d1157505050565b60040360031b82901b16169150565b634e487b7160e01b5f52601160045260245ffd5b60ff1660ff8114610d455760010190565b610d20565b60051b90611fe060e0831692168203610d4557565b60ff166024019060ff8211610d4557565b60ff60019116019060ff8211610d4557565b634e487b7160e01b5f52603260045260245ffd5b906007811015610da75760051b0190565b610d82565b916040519060e0820182811067ffffffffffffffff8211176106e75760405260e036833781610dde6101048514610b7f565b610df1610deb8587610bcb565b90610cc0565b93610e0e610e08610e028389610bd9565b90610cdc565b60e01c90565b945f5b60ff808216906006821015610e735790610e68610e61610deb85948d89610e42610e3d610e6e9a610d4a565b610d5f565b9280610e58610e3d610e538c610d70565b610d4a565b16931691610ca8565b9189610d96565b52610d34565b610e11565b50505095610deb60c092610e8992969496610bea565b910152565b9015610da75790565b15610e9e57565b60405162461bcd60e51b815260206004820152601660248201527f6e6f7420612044796e616d6963466565547854797065000000000000000000006044820152606490fd5b805115610da75760200190565b805160031015610da75760800190565b805160011015610da75760400190565b805160021015610da75760600190565b8051821015610da75760209160051b010190565b906020825192015160ff60f81b908181169360018110610f5357505050565b60010360031b82901b16169150565b67ffffffffffffffff81116106e75760051b60200190565b90610f8482610f62565b604090610f93825191826106ec565b8381528093610fa4601f1991610f62565b01905f92835b838110610fb8575050505050565b8151906060918281019281841067ffffffffffffffff8511176106e75760209385528782528390808284015285830152828601015201610faa565b90610ffd82610f62565b61100a60405191826106ec565b828152809261101b601f1991610f62565b0190602036910137565b602081519101519060208110610cce575090565b61110d6111076110b76110b26110ad6110a661105361073d565b978060ff986110a160028b61109a61109461106e8787610e8e565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b1614610e97565b610bfc565b369161079a565b61151f565b61156e565b600160f81b7fff000000000000000000000000000000000000000000000000000000000000006110f76110f26110ec85610ee3565b5161173b565b610f34565b161461121c575b94939294610ef0565b5161156e565b926111188451610f7a565b915f80945b8651918487169283101561120e57611138611107848a610f20565b9761116761114e6111488b610ee3565b51611663565b611158868a610f20565b51906001600160a01b03169052565b6111736111078a610f00565b9261117e8451610ff3565b956020968761118d888c610f20565b510152835b8551898216908110156111d657908a610e68838b6111c98c6111c36111be6110ec8f9a6111d19b610f20565b611025565b95610f20565b510151610f20565b611192565b5050989295509861120593506111fd6111f66110ec604093979497610f10565b9289610f20565b510152610d34565b9491909561111d565b5094505093505060a0830152565b600185526110fe565b90918060409360208452816020850152848401375f828201840152601f01601f1916010190565b909167ffffffffffffffff82165f5260016020526001600160a01b0360405f205416156112f8576112c8926112ab610a87610a8761129e60209667ffffffffffffffff165f52600160205260405f2090565b546001600160a01b031690565b906040518095819482936322bb937360e11b845260048401611225565b03915afa908115610b7a575f916112dd575090565b6112f5915060203d8111610b7357610b6481836106ec565b90565b60405162461bcd60e51b815260206004820152601660248201527f636861696e207665726966696572206e6f7420736574000000000000000000006044820152606490fd5b6fffffffffffffffffffffffffffffffff19903581811693926010811061136357505050565b60100360031b82901b16169150565b7fffffffffffffffff00000000000000000000000000000000000000000000000090358181169392600881106113a757505050565b60080360031b82901b16169150565b91906113c061073d565b928161016011610140578180826114076113f26113ec6113e66114bb986110a698610c18565b9061133d565b60801c90565b6fffffffffffffffffffffffffffffffff1690565b61014082013560801b178752611429611423610deb8484610c2b565b60801b90565b61143c6113f26113ec6113e68686610c3e565b176020880152611462611455610e08610e028585610c51565b63ffffffff166040890152565b61149261148161147b6114758585610c64565b90611372565b60c01c90565b67ffffffffffffffff166060890152565b6114b66114a561147b6114758585610c77565b67ffffffffffffffff166080890152565b610c8a565b60a0830152565b5f54906001600160a01b0380911691826001600160a01b03198216175f55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b60405190611514826106cb565b5f6020838281520152565b611527611507565b5060208151916040519261153a846106cb565b835201602082015290565b9060018201809211610d4557565b91908201809211610d4557565b5f198114610d455760010190565b6115778161163f565b156101405761158581611788565b61158e81610f62565b9161159c60405193846106ec565b818352601f196115ab83610f62565b015f5b8181106116285750506115cf6020809201516115c98161186e565b90611553565b5f905b8382106115e0575050505090565b61161c816115f0611622936117e9565b906115f961072e565b828152818782015261160b868a610f20565b526116168589610f20565b50611553565b91611560565b906115d2565b602090611633611507565b828288010152016115ae565b80511561165e57602060c0910151515f1a1061165a57600190565b5f90565b505f90565b8051601581036101405780151590816116ab575b50156101405761168e6001600160a01b039161170f565b9051906020811061169e57501690565b6020036101000a90041690565b6021915011155f611677565b60bf19810191908211610d4557565b607f19810191908211610d4557565b6020039060208211610d4557565b5f19810191908211610d4557565b60f619810191908211610d4557565b60b619810191908211610d4557565b90602082019161171f835161186e565b925190838201809211610d455751928303928311610d45579190565b8051156101405761174e6112f59161170f565b61175a8193929361077e565b9261176860405194856106ec565b818452601f196117778361077e565b0136602086013783602001906118db565b80511561165e575f90602081019081516117a18161186e565b8101809111610d4557915190518101809111610d455791905b8281106117c75750905090565b6117d0816117e9565b8101809111610d45576117e39091611560565b906117ba565b80515f1a9060808210156117fe575050600190565b60b882101561181957506118146112f5916116c6565b611545565b9060c081101561183d5760b51991600160b783602003016101000a91015104010190565b9060f882101561185457506118146112f5916116b7565b60010151602082900360f7016101000a90040160f5190190565b515f1a608081101561187f57505f90565b60b8811080156118b6575b156118955750600190565b60c08110156118aa576118146112f591611700565b6118146112f5916116f1565b5060c0811015801561188a575060f8811061188a565b601f8111610d45576101000a90565b9290919283156119525792915b60209384841061191d5780518252848101809111610d4557938101809111610d455791601f198101908111610d4557916118e8565b919350918061192b57505050565b61193f61193a611944926116d5565b6118cc565b6116e3565b905182518216911916179052565b5091505056fea26469706673582212208e8af8864b40215585f68244ee45749e12931db3e49d2b7134f845425bd8321f64736f6c63430008140033",
}

// ReceiptVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptVerifierMetaData.ABI instead.
var ReceiptVerifierABI = ReceiptVerifierMetaData.ABI

// ReceiptVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiptVerifierMetaData.Bin instead.
var ReceiptVerifierBin = ReceiptVerifierMetaData.Bin

// DeployReceiptVerifier deploys a new Ethereum contract, binding an instance of ReceiptVerifier to it.
func DeployReceiptVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _blockChunks common.Address) (common.Address, *types.Transaction, *ReceiptVerifier, error) {
	parsed, err := ReceiptVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiptVerifierBin), backend, _blockChunks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiptVerifier{ReceiptVerifierCaller: ReceiptVerifierCaller{contract: contract}, ReceiptVerifierTransactor: ReceiptVerifierTransactor{contract: contract}, ReceiptVerifierFilterer: ReceiptVerifierFilterer{contract: contract}}, nil
}

// ReceiptVerifier is an auto generated Go binding around an Ethereum contract.
type ReceiptVerifier struct {
	ReceiptVerifierCaller     // Read-only binding to the contract
	ReceiptVerifierTransactor // Write-only binding to the contract
	ReceiptVerifierFilterer   // Log filterer for contract events
}

// ReceiptVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptVerifierSession struct {
	Contract     *ReceiptVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptVerifierCallerSession struct {
	Contract *ReceiptVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReceiptVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptVerifierTransactorSession struct {
	Contract     *ReceiptVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReceiptVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptVerifierRaw struct {
	Contract *ReceiptVerifier // Generic contract binding to access the raw methods on
}

// ReceiptVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptVerifierCallerRaw struct {
	Contract *ReceiptVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptVerifierTransactorRaw struct {
	Contract *ReceiptVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptVerifier creates a new instance of ReceiptVerifier, bound to a specific deployed contract.
func NewReceiptVerifier(address common.Address, backend bind.ContractBackend) (*ReceiptVerifier, error) {
	contract, err := bindReceiptVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifier{ReceiptVerifierCaller: ReceiptVerifierCaller{contract: contract}, ReceiptVerifierTransactor: ReceiptVerifierTransactor{contract: contract}, ReceiptVerifierFilterer: ReceiptVerifierFilterer{contract: contract}}, nil
}

// NewReceiptVerifierCaller creates a new read-only instance of ReceiptVerifier, bound to a specific deployed contract.
func NewReceiptVerifierCaller(address common.Address, caller bind.ContractCaller) (*ReceiptVerifierCaller, error) {
	contract, err := bindReceiptVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierCaller{contract: contract}, nil
}

// NewReceiptVerifierTransactor creates a new write-only instance of ReceiptVerifier, bound to a specific deployed contract.
func NewReceiptVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptVerifierTransactor, error) {
	contract, err := bindReceiptVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierTransactor{contract: contract}, nil
}

// NewReceiptVerifierFilterer creates a new log filterer instance of ReceiptVerifier, bound to a specific deployed contract.
func NewReceiptVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptVerifierFilterer, error) {
	contract, err := bindReceiptVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierFilterer{contract: contract}, nil
}

// bindReceiptVerifier binds a generic wrapper to an already deployed contract.
func bindReceiptVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiptVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptVerifier *ReceiptVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptVerifier.Contract.ReceiptVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptVerifier *ReceiptVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.ReceiptVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptVerifier *ReceiptVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.ReceiptVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptVerifier *ReceiptVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptVerifier *ReceiptVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptVerifier *ReceiptVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.contract.Transact(opts, method, params...)
}

// BlockChunks is a free data retrieval call binding the contract method 0x724796ed.
//
// Solidity: function blockChunks() view returns(address)
func (_ReceiptVerifier *ReceiptVerifierCaller) BlockChunks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiptVerifier.contract.Call(opts, &out, "blockChunks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockChunks is a free data retrieval call binding the contract method 0x724796ed.
//
// Solidity: function blockChunks() view returns(address)
func (_ReceiptVerifier *ReceiptVerifierSession) BlockChunks() (common.Address, error) {
	return _ReceiptVerifier.Contract.BlockChunks(&_ReceiptVerifier.CallOpts)
}

// BlockChunks is a free data retrieval call binding the contract method 0x724796ed.
//
// Solidity: function blockChunks() view returns(address)
func (_ReceiptVerifier *ReceiptVerifierCallerSession) BlockChunks() (common.Address, error) {
	return _ReceiptVerifier.Contract.BlockChunks(&_ReceiptVerifier.CallOpts)
}

// DecodeReceipt is a free data retrieval call binding the contract method 0xc109ae5d.
//
// Solidity: function decodeReceipt(bytes receiptRaw) pure returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierCaller) DecodeReceipt(opts *bind.CallOpts, receiptRaw []byte) (IReceiptVerifierReceiptInfo, error) {
	var out []interface{}
	err := _ReceiptVerifier.contract.Call(opts, &out, "decodeReceipt", receiptRaw)

	if err != nil {
		return *new(IReceiptVerifierReceiptInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IReceiptVerifierReceiptInfo)).(*IReceiptVerifierReceiptInfo)

	return out0, err

}

// DecodeReceipt is a free data retrieval call binding the contract method 0xc109ae5d.
//
// Solidity: function decodeReceipt(bytes receiptRaw) pure returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierSession) DecodeReceipt(receiptRaw []byte) (IReceiptVerifierReceiptInfo, error) {
	return _ReceiptVerifier.Contract.DecodeReceipt(&_ReceiptVerifier.CallOpts, receiptRaw)
}

// DecodeReceipt is a free data retrieval call binding the contract method 0xc109ae5d.
//
// Solidity: function decodeReceipt(bytes receiptRaw) pure returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierCallerSession) DecodeReceipt(receiptRaw []byte) (IReceiptVerifierReceiptInfo, error) {
	return _ReceiptVerifier.Contract.DecodeReceipt(&_ReceiptVerifier.CallOpts, receiptRaw)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReceiptVerifier *ReceiptVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiptVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReceiptVerifier *ReceiptVerifierSession) Owner() (common.Address, error) {
	return _ReceiptVerifier.Contract.Owner(&_ReceiptVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReceiptVerifier *ReceiptVerifierCallerSession) Owner() (common.Address, error) {
	return _ReceiptVerifier.Contract.Owner(&_ReceiptVerifier.CallOpts)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_ReceiptVerifier *ReceiptVerifierCaller) VerifierAddresses(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _ReceiptVerifier.contract.Call(opts, &out, "verifierAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_ReceiptVerifier *ReceiptVerifierSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _ReceiptVerifier.Contract.VerifierAddresses(&_ReceiptVerifier.CallOpts, arg0)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_ReceiptVerifier *ReceiptVerifierCallerSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _ReceiptVerifier.Contract.VerifierAddresses(&_ReceiptVerifier.CallOpts, arg0)
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x68ac2f78.
//
// Solidity: function verifyReceipt(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierCaller) VerifyReceipt(opts *bind.CallOpts, receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (IReceiptVerifierReceiptInfo, error) {
	var out []interface{}
	err := _ReceiptVerifier.contract.Call(opts, &out, "verifyReceipt", receiptRaw, proofData, auxiBlkVerifyInfo)

	if err != nil {
		return *new(IReceiptVerifierReceiptInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IReceiptVerifierReceiptInfo)).(*IReceiptVerifierReceiptInfo)

	return out0, err

}

// VerifyReceipt is a free data retrieval call binding the contract method 0x68ac2f78.
//
// Solidity: function verifyReceipt(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierSession) VerifyReceipt(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (IReceiptVerifierReceiptInfo, error) {
	return _ReceiptVerifier.Contract.VerifyReceipt(&_ReceiptVerifier.CallOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x68ac2f78.
//
// Solidity: function verifyReceipt(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierCallerSession) VerifyReceipt(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (IReceiptVerifierReceiptInfo, error) {
	return _ReceiptVerifier.Contract.VerifyReceipt(&_ReceiptVerifier.CallOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReceiptVerifier *ReceiptVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReceiptVerifier *ReceiptVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.RenounceOwnership(&_ReceiptVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReceiptVerifier *ReceiptVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.RenounceOwnership(&_ReceiptVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReceiptVerifier *ReceiptVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReceiptVerifier *ReceiptVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.TransferOwnership(&_ReceiptVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReceiptVerifier *ReceiptVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.TransferOwnership(&_ReceiptVerifier.TransactOpts, newOwner)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _blockChunks) returns()
func (_ReceiptVerifier *ReceiptVerifierTransactor) UpdateBlockChunks(opts *bind.TransactOpts, _blockChunks common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.contract.Transact(opts, "updateBlockChunks", _blockChunks)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _blockChunks) returns()
func (_ReceiptVerifier *ReceiptVerifierSession) UpdateBlockChunks(_blockChunks common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.UpdateBlockChunks(&_ReceiptVerifier.TransactOpts, _blockChunks)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _blockChunks) returns()
func (_ReceiptVerifier *ReceiptVerifierTransactorSession) UpdateBlockChunks(_blockChunks common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.UpdateBlockChunks(&_ReceiptVerifier.TransactOpts, _blockChunks)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_ReceiptVerifier *ReceiptVerifierTransactor) UpdateVerifierAddress(opts *bind.TransactOpts, _chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.contract.Transact(opts, "updateVerifierAddress", _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_ReceiptVerifier *ReceiptVerifierSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.UpdateVerifierAddress(&_ReceiptVerifier.TransactOpts, _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_ReceiptVerifier *ReceiptVerifierTransactorSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.UpdateVerifierAddress(&_ReceiptVerifier.TransactOpts, _chainId, _verifierAddress)
}

// VerifyReceiptAndLog is a paid mutator transaction binding the contract method 0x3996da7a.
//
// Solidity: function verifyReceiptAndLog(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierTransactor) VerifyReceiptAndLog(opts *bind.TransactOpts, receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _ReceiptVerifier.contract.Transact(opts, "verifyReceiptAndLog", receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceiptAndLog is a paid mutator transaction binding the contract method 0x3996da7a.
//
// Solidity: function verifyReceiptAndLog(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierSession) VerifyReceiptAndLog(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.VerifyReceiptAndLog(&_ReceiptVerifier.TransactOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyReceiptAndLog is a paid mutator transaction binding the contract method 0x3996da7a.
//
// Solidity: function verifyReceiptAndLog(bytes receiptRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((bool,uint64,bytes32,uint32,uint64,(address,bytes32[],bytes)[]) info)
func (_ReceiptVerifier *ReceiptVerifierTransactorSession) VerifyReceiptAndLog(receiptRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _ReceiptVerifier.Contract.VerifyReceiptAndLog(&_ReceiptVerifier.TransactOpts, receiptRaw, proofData, auxiBlkVerifyInfo)
}

// ReceiptVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReceiptVerifier contract.
type ReceiptVerifierOwnershipTransferredIterator struct {
	Event *ReceiptVerifierOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReceiptVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptVerifierOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReceiptVerifierOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReceiptVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the ReceiptVerifier contract.
type ReceiptVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReceiptVerifier *ReceiptVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReceiptVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReceiptVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierOwnershipTransferredIterator{contract: _ReceiptVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReceiptVerifier *ReceiptVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReceiptVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReceiptVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptVerifierOwnershipTransferred)
				if err := _ReceiptVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReceiptVerifier *ReceiptVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*ReceiptVerifierOwnershipTransferred, error) {
	event := new(ReceiptVerifierOwnershipTransferred)
	if err := _ReceiptVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptVerifierUpdateBlockChunksIterator is returned from FilterUpdateBlockChunks and is used to iterate over the raw logs and unpacked data for UpdateBlockChunks events raised by the ReceiptVerifier contract.
type ReceiptVerifierUpdateBlockChunksIterator struct {
	Event *ReceiptVerifierUpdateBlockChunks // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReceiptVerifierUpdateBlockChunksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptVerifierUpdateBlockChunks)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReceiptVerifierUpdateBlockChunks)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReceiptVerifierUpdateBlockChunksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptVerifierUpdateBlockChunksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptVerifierUpdateBlockChunks represents a UpdateBlockChunks event raised by the ReceiptVerifier contract.
type ReceiptVerifierUpdateBlockChunks struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateBlockChunks is a free log retrieval operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_ReceiptVerifier *ReceiptVerifierFilterer) FilterUpdateBlockChunks(opts *bind.FilterOpts) (*ReceiptVerifierUpdateBlockChunksIterator, error) {

	logs, sub, err := _ReceiptVerifier.contract.FilterLogs(opts, "UpdateBlockChunks")
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierUpdateBlockChunksIterator{contract: _ReceiptVerifier.contract, event: "UpdateBlockChunks", logs: logs, sub: sub}, nil
}

// WatchUpdateBlockChunks is a free log subscription operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_ReceiptVerifier *ReceiptVerifierFilterer) WatchUpdateBlockChunks(opts *bind.WatchOpts, sink chan<- *ReceiptVerifierUpdateBlockChunks) (event.Subscription, error) {

	logs, sub, err := _ReceiptVerifier.contract.WatchLogs(opts, "UpdateBlockChunks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptVerifierUpdateBlockChunks)
				if err := _ReceiptVerifier.contract.UnpackLog(event, "UpdateBlockChunks", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateBlockChunks is a log parse operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_ReceiptVerifier *ReceiptVerifierFilterer) ParseUpdateBlockChunks(log types.Log) (*ReceiptVerifierUpdateBlockChunks, error) {
	event := new(ReceiptVerifierUpdateBlockChunks)
	if err := _ReceiptVerifier.contract.UnpackLog(event, "UpdateBlockChunks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptVerifierUpdateVerifierAddressIterator is returned from FilterUpdateVerifierAddress and is used to iterate over the raw logs and unpacked data for UpdateVerifierAddress events raised by the ReceiptVerifier contract.
type ReceiptVerifierUpdateVerifierAddressIterator struct {
	Event *ReceiptVerifierUpdateVerifierAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReceiptVerifierUpdateVerifierAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptVerifierUpdateVerifierAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReceiptVerifierUpdateVerifierAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReceiptVerifierUpdateVerifierAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptVerifierUpdateVerifierAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptVerifierUpdateVerifierAddress represents a UpdateVerifierAddress event raised by the ReceiptVerifier contract.
type ReceiptVerifierUpdateVerifierAddress struct {
	ChainId    uint64
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifierAddress is a free log retrieval operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_ReceiptVerifier *ReceiptVerifierFilterer) FilterUpdateVerifierAddress(opts *bind.FilterOpts) (*ReceiptVerifierUpdateVerifierAddressIterator, error) {

	logs, sub, err := _ReceiptVerifier.contract.FilterLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierUpdateVerifierAddressIterator{contract: _ReceiptVerifier.contract, event: "UpdateVerifierAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifierAddress is a free log subscription operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_ReceiptVerifier *ReceiptVerifierFilterer) WatchUpdateVerifierAddress(opts *bind.WatchOpts, sink chan<- *ReceiptVerifierUpdateVerifierAddress) (event.Subscription, error) {

	logs, sub, err := _ReceiptVerifier.contract.WatchLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptVerifierUpdateVerifierAddress)
				if err := _ReceiptVerifier.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateVerifierAddress is a log parse operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_ReceiptVerifier *ReceiptVerifierFilterer) ParseUpdateVerifierAddress(log types.Log) (*ReceiptVerifierUpdateVerifierAddress, error) {
	event := new(ReceiptVerifierUpdateVerifierAddress)
	if err := _ReceiptVerifier.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptVerifierVerifiedReceiptIterator is returned from FilterVerifiedReceipt and is used to iterate over the raw logs and unpacked data for VerifiedReceipt events raised by the ReceiptVerifier contract.
type ReceiptVerifierVerifiedReceiptIterator struct {
	Event *ReceiptVerifierVerifiedReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReceiptVerifierVerifiedReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptVerifierVerifiedReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReceiptVerifierVerifiedReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReceiptVerifierVerifiedReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptVerifierVerifiedReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptVerifierVerifiedReceipt represents a VerifiedReceipt event raised by the ReceiptVerifier contract.
type ReceiptVerifierVerifiedReceipt struct {
	ChainId     uint64
	ReceiptHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerifiedReceipt is a free log retrieval operation binding the contract event 0xa5db3bb7a25cc2804c7835ab71b15513b8c68585e3593c3fdee39a8837547366.
//
// Solidity: event VerifiedReceipt(uint64 chainId, bytes32 receiptHash)
func (_ReceiptVerifier *ReceiptVerifierFilterer) FilterVerifiedReceipt(opts *bind.FilterOpts) (*ReceiptVerifierVerifiedReceiptIterator, error) {

	logs, sub, err := _ReceiptVerifier.contract.FilterLogs(opts, "VerifiedReceipt")
	if err != nil {
		return nil, err
	}
	return &ReceiptVerifierVerifiedReceiptIterator{contract: _ReceiptVerifier.contract, event: "VerifiedReceipt", logs: logs, sub: sub}, nil
}

// WatchVerifiedReceipt is a free log subscription operation binding the contract event 0xa5db3bb7a25cc2804c7835ab71b15513b8c68585e3593c3fdee39a8837547366.
//
// Solidity: event VerifiedReceipt(uint64 chainId, bytes32 receiptHash)
func (_ReceiptVerifier *ReceiptVerifierFilterer) WatchVerifiedReceipt(opts *bind.WatchOpts, sink chan<- *ReceiptVerifierVerifiedReceipt) (event.Subscription, error) {

	logs, sub, err := _ReceiptVerifier.contract.WatchLogs(opts, "VerifiedReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptVerifierVerifiedReceipt)
				if err := _ReceiptVerifier.contract.UnpackLog(event, "VerifiedReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifiedReceipt is a log parse operation binding the contract event 0xa5db3bb7a25cc2804c7835ab71b15513b8c68585e3593c3fdee39a8837547366.
//
// Solidity: event VerifiedReceipt(uint64 chainId, bytes32 receiptHash)
func (_ReceiptVerifier *ReceiptVerifierFilterer) ParseVerifiedReceipt(log types.Log) (*ReceiptVerifierVerifiedReceipt, error) {
	event := new(ReceiptVerifierVerifiedReceipt)
	if err := _ReceiptVerifier.contract.UnpackLog(event, "VerifiedReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SMTMetaData contains all meta data concerning the SMT contract.
var SMTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"_anchorProviders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_initRoots\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"anchorProvider\",\"type\":\"address\"}],\"name\":\"AnchorProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endBlockNum\",\"type\":\"uint64\"}],\"name\":\"SmtRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"anchorProviders\",\"outputs\":[{\"internalType\":\"contractIAnchorBlocks\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getLatestRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"}],\"name\":\"isSmtRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"latestRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"anchorProvider\",\"type\":\"address\"}],\"name\":\"setAnchorProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"smtRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newSmtRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"endBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextChunkMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"knowledgeProof\",\"type\":\"uint256[2]\"}],\"internalType\":\"structISMT.SmtUpdate\",\"name\":\"u\",\"type\":\"tuple\"}],\"name\":\"updateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60406080815234620002795762000ef4803803806200001e8162000297565b928339810191608082840312620002795781516001600160401b0391908281116200027957830184601f820112156200027957805194620000696200006387620002d1565b62000297565b9182968084526020808095019160051b8301019183831162000279578401905b8282106200027d5750505081850151848111620002795781620000ae918701620002e9565b9383860151818111620002795782620000c9918801620002e9565b956060810151908282116200027957019180601f8401121562000279578251620000f76200006382620002d1565b93858086848152019260051b820101928311620002795785809101915b838310620002685750505f8054336001600160a01b031980831682178455929a92956001600160a01b0395509293509084167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08b80a3620001798a5189511462000357565b620001888a5183511462000357565b620001978a5186511462000357565b885b8a51811015620002595781620001b0828d62000393565b511684620001bf838c62000393565b5116818c52600190818a528a8d20908882541617905585620001e2848762000393565b5116828d5260028a528a8d20908882541617905560038952898c2062000209848a62000393565b518d528952898c209060ff1982541617905562000227828862000393565b51908b5260048852888b20555f198114620002455760010162000199565b634e487b7160e01b8a52601160045260248afd5b8751610b379081620003bd8239f35b825181529181019186910162000114565b5f80fd5b815187811681036200027957815290840190840162000089565b6040519190601f01601f191682016001600160401b03811183821017620002bd57604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b038111620002bd5760051b60200190565b9080601f830112156200027957815190620003086200006383620002d1565b9182938184526020808095019260051b82010192831162000279578301905b82821062000336575050505090565b81516001600160a01b03811681036200027957815290830190830162000327565b156200035f57565b60405162461bcd60e51b815260206004820152600c60248201526b0d8cadc40dad2e6dac2e8c6d60a31b6044820152606490fd5b8051821015620003a85760209160051b010190565b634e487b7160e01b5f52603260045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c9081631019b616146109df575080633870253214610995578063479aa6da146108ac5780635ca32bd8146108e55780636ae3e080146108ac578063715018a6146108515780638195408d1461080f5780638da5cb5b146107ea57806397c7c309146102785780639c8413c5146101c3578063afe8154b146101815763f2fde38b146100a0575f80fd5b3461017d57602036600319011261017d576004356001600160a01b0380821680920361017d575f54908116906100d7338314610ab6565b8215610112576001600160a01b0319839116175f557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b5f80fd5b3461017d57602036600319011261017d5767ffffffffffffffff6101a3610a24565b165f52600160205260206001600160a01b0360405f205416604051908152f35b3461017d57604036600319011261017d577fb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be81626101fd610a24565b610205610a3b565b906001600160a01b0361021c815f54163314610ab6565b67ffffffffffffffff82165f52600260205260405f209083166001600160a01b0319825416179055610273604051928392839092916001600160a01b0360209167ffffffffffffffff604085019616845216910152565b0390a1005b3461017d5761022036600319011261017d57610292610a24565b61020036602319011261017d576040519060e0820182811067ffffffffffffffff82111761063c57604052602435825260443567ffffffffffffffff8116810361017d576020830152606435604083015260843560608301523660c3121561017d57604051610100810181811067ffffffffffffffff82111761063c57604052806101a49136831161017d5760a4905b8382106107da5750506080840152366101c3121561017d576040519061034782610a51565b8190366101e41161017d57905b6101e482106107ca57505060a083015236610203121561017d5760405161037a81610a51565b80366102241161017d576101e4905b61022482106107ba57505060c0830152606082015115610695575b67ffffffffffffffff81165f52600460205260405f205460026020526001600160a01b0360405f205416801561065057604051918261012081011067ffffffffffffffff6101208501111761063c5761012083016040526101203684378060801c83526fffffffffffffffffffffffffffffffff8091166020840152845160801c6040840152808551166060840152604085015160801c60808401528060408601511660a084015267ffffffffffffffff60208601511660c0840152606085015160801c60e0840152606085015116610100830152608084015160a08501519260c0860151604051948593633072c1a360e11b8552600485015f905b6008821061062257505050906104be6104c992610104860190610a8f565b610144840190610a8f565b5f61018483015b60098210610608575050506102a4816020935afa9081156105fd575f916105c2575b501561057d577f05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7918167ffffffffffffffff604093165f526003602052825f2082515f52602052825f20600160ff1982541617905567ffffffffffffffff825191165f526004602052825f205567ffffffffffffffff602082519201511682519182526020820152a1005b60405162461bcd60e51b815260206004820152601060248201527f696e76616c6964207a6b2070726f6f66000000000000000000000000000000006044820152606490fd5b90506020813d6020116105f5575b816105dd60209383610a6d565b8101031261017d5751801515810361017d57836104f2565b3d91506105d0565b6040513d5f823e3d90fd5b8293506020809160019394518152019301910184926104d0565b8251815288965060209283019260019290920191016104a0565b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152601760248201527f6e6f20766572696669657220666f7220636861696e49640000000000000000006044820152606490fd5b67ffffffffffffffff81165f5260016020526001600160a01b0360405f205416801561077557602067ffffffffffffffff818501511660246040518094819363f25b3f9960e01b835260048301525afa9081156105fd575f91610743575b506040830151146103a45760405162461bcd60e51b815260206004820152601360248201527f616e63686f7220636865636b206661696c6564000000000000000000000000006044820152606490fd5b90506020813d60201161076d575b8161075e60209383610a6d565b8101031261017d5751836106f3565b3d9150610751565b60405162461bcd60e51b815260206004820152601760248201527f756e6b6e6f776e20616e63686f722070726f76696465720000000000000000006044820152606490fd5b8135815260209182019101610389565b8135815260209182019101610354565b8135815260209182019101610322565b3461017d575f36600319011261017d5760206001600160a01b035f5416604051908152f35b3461017d57602036600319011261017d5767ffffffffffffffff610831610a24565b165f52600260205260206001600160a01b0360405f205416604051908152f35b3461017d575f36600319011261017d575f80546001600160a01b03196001600160a01b03821691610883338414610ab6565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461017d57602036600319011261017d5767ffffffffffffffff6108ce610a24565b165f526004602052602060405f2054604051908152f35b3461017d57604036600319011261017d577fd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c809061091f610a24565b610927610a3b565b906001600160a01b0361093e815f54163314610ab6565b67ffffffffffffffff82165f52600160205260405f209083166001600160a01b0319825416179055610273604051928392839092916001600160a01b0360209167ffffffffffffffff604085019616845216910152565b3461017d57604036600319011261017d5767ffffffffffffffff6109b7610a24565b165f52600360205260405f206024355f52602052602060ff60405f2054166040519015158152f35b3461017d57604036600319011261017d5760209067ffffffffffffffff610a04610a24565b165f526003825260405f206024355f52825260ff60405f20541615158152f35b6004359067ffffffffffffffff8216820361017d57565b602435906001600160a01b038216820361017d57565b6040810190811067ffffffffffffffff82111761063c57604052565b90601f8019910116810190811067ffffffffffffffff82111761063c57604052565b5f915b60028310610a9f57505050565b600190825181526020809101920192019190610a92565b15610abd57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfea26469706673582212206f75ec3cb662004b1536186a4b0618689944ce8f522f015c5dbe28f5233111b164736f6c63430008140033",
}

// SMTABI is the input ABI used to generate the binding from.
// Deprecated: Use SMTMetaData.ABI instead.
var SMTABI = SMTMetaData.ABI

// SMTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SMTMetaData.Bin instead.
var SMTBin = SMTMetaData.Bin

// DeploySMT deploys a new Ethereum contract, binding an instance of SMT to it.
func DeploySMT(auth *bind.TransactOpts, backend bind.ContractBackend, _chainIds []uint64, _anchorProviders []common.Address, _verifiers []common.Address, _initRoots [][32]byte) (common.Address, *types.Transaction, *SMT, error) {
	parsed, err := SMTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SMTBin), backend, _chainIds, _anchorProviders, _verifiers, _initRoots)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SMT{SMTCaller: SMTCaller{contract: contract}, SMTTransactor: SMTTransactor{contract: contract}, SMTFilterer: SMTFilterer{contract: contract}}, nil
}

// SMT is an auto generated Go binding around an Ethereum contract.
type SMT struct {
	SMTCaller     // Read-only binding to the contract
	SMTTransactor // Write-only binding to the contract
	SMTFilterer   // Log filterer for contract events
}

// SMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type SMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SMTSession struct {
	Contract     *SMT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SMTCallerSession struct {
	Contract *SMTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SMTTransactorSession struct {
	Contract     *SMTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type SMTRaw struct {
	Contract *SMT // Generic contract binding to access the raw methods on
}

// SMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SMTCallerRaw struct {
	Contract *SMTCaller // Generic read-only contract binding to access the raw methods on
}

// SMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SMTTransactorRaw struct {
	Contract *SMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSMT creates a new instance of SMT, bound to a specific deployed contract.
func NewSMT(address common.Address, backend bind.ContractBackend) (*SMT, error) {
	contract, err := bindSMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SMT{SMTCaller: SMTCaller{contract: contract}, SMTTransactor: SMTTransactor{contract: contract}, SMTFilterer: SMTFilterer{contract: contract}}, nil
}

// NewSMTCaller creates a new read-only instance of SMT, bound to a specific deployed contract.
func NewSMTCaller(address common.Address, caller bind.ContractCaller) (*SMTCaller, error) {
	contract, err := bindSMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SMTCaller{contract: contract}, nil
}

// NewSMTTransactor creates a new write-only instance of SMT, bound to a specific deployed contract.
func NewSMTTransactor(address common.Address, transactor bind.ContractTransactor) (*SMTTransactor, error) {
	contract, err := bindSMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SMTTransactor{contract: contract}, nil
}

// NewSMTFilterer creates a new log filterer instance of SMT, bound to a specific deployed contract.
func NewSMTFilterer(address common.Address, filterer bind.ContractFilterer) (*SMTFilterer, error) {
	contract, err := bindSMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SMTFilterer{contract: contract}, nil
}

// bindSMT binds a generic wrapper to an already deployed contract.
func bindSMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMT *SMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SMT.Contract.SMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMT *SMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMT.Contract.SMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMT *SMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMT.Contract.SMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMT *SMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMT *SMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMT *SMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMT.Contract.contract.Transact(opts, method, params...)
}

// AnchorProviders is a free data retrieval call binding the contract method 0xafe8154b.
//
// Solidity: function anchorProviders(uint64 ) view returns(address)
func (_SMT *SMTCaller) AnchorProviders(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "anchorProviders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnchorProviders is a free data retrieval call binding the contract method 0xafe8154b.
//
// Solidity: function anchorProviders(uint64 ) view returns(address)
func (_SMT *SMTSession) AnchorProviders(arg0 uint64) (common.Address, error) {
	return _SMT.Contract.AnchorProviders(&_SMT.CallOpts, arg0)
}

// AnchorProviders is a free data retrieval call binding the contract method 0xafe8154b.
//
// Solidity: function anchorProviders(uint64 ) view returns(address)
func (_SMT *SMTCallerSession) AnchorProviders(arg0 uint64) (common.Address, error) {
	return _SMT.Contract.AnchorProviders(&_SMT.CallOpts, arg0)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0x479aa6da.
//
// Solidity: function getLatestRoot(uint64 chainId) view returns(bytes32)
func (_SMT *SMTCaller) GetLatestRoot(opts *bind.CallOpts, chainId uint64) ([32]byte, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "getLatestRoot", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestRoot is a free data retrieval call binding the contract method 0x479aa6da.
//
// Solidity: function getLatestRoot(uint64 chainId) view returns(bytes32)
func (_SMT *SMTSession) GetLatestRoot(chainId uint64) ([32]byte, error) {
	return _SMT.Contract.GetLatestRoot(&_SMT.CallOpts, chainId)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0x479aa6da.
//
// Solidity: function getLatestRoot(uint64 chainId) view returns(bytes32)
func (_SMT *SMTCallerSession) GetLatestRoot(chainId uint64) ([32]byte, error) {
	return _SMT.Contract.GetLatestRoot(&_SMT.CallOpts, chainId)
}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_SMT *SMTCaller) IsSmtRootValid(opts *bind.CallOpts, chainId uint64, smtRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "isSmtRootValid", chainId, smtRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_SMT *SMTSession) IsSmtRootValid(chainId uint64, smtRoot [32]byte) (bool, error) {
	return _SMT.Contract.IsSmtRootValid(&_SMT.CallOpts, chainId, smtRoot)
}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_SMT *SMTCallerSession) IsSmtRootValid(chainId uint64, smtRoot [32]byte) (bool, error) {
	return _SMT.Contract.IsSmtRootValid(&_SMT.CallOpts, chainId, smtRoot)
}

// LatestRoots is a free data retrieval call binding the contract method 0x6ae3e080.
//
// Solidity: function latestRoots(uint64 ) view returns(bytes32)
func (_SMT *SMTCaller) LatestRoots(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "latestRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRoots is a free data retrieval call binding the contract method 0x6ae3e080.
//
// Solidity: function latestRoots(uint64 ) view returns(bytes32)
func (_SMT *SMTSession) LatestRoots(arg0 uint64) ([32]byte, error) {
	return _SMT.Contract.LatestRoots(&_SMT.CallOpts, arg0)
}

// LatestRoots is a free data retrieval call binding the contract method 0x6ae3e080.
//
// Solidity: function latestRoots(uint64 ) view returns(bytes32)
func (_SMT *SMTCallerSession) LatestRoots(arg0 uint64) ([32]byte, error) {
	return _SMT.Contract.LatestRoots(&_SMT.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SMT *SMTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SMT *SMTSession) Owner() (common.Address, error) {
	return _SMT.Contract.Owner(&_SMT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SMT *SMTCallerSession) Owner() (common.Address, error) {
	return _SMT.Contract.Owner(&_SMT.CallOpts)
}

// SmtRoots is a free data retrieval call binding the contract method 0x38702532.
//
// Solidity: function smtRoots(uint64 , bytes32 ) view returns(bool)
func (_SMT *SMTCaller) SmtRoots(opts *bind.CallOpts, arg0 uint64, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "smtRoots", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SmtRoots is a free data retrieval call binding the contract method 0x38702532.
//
// Solidity: function smtRoots(uint64 , bytes32 ) view returns(bool)
func (_SMT *SMTSession) SmtRoots(arg0 uint64, arg1 [32]byte) (bool, error) {
	return _SMT.Contract.SmtRoots(&_SMT.CallOpts, arg0, arg1)
}

// SmtRoots is a free data retrieval call binding the contract method 0x38702532.
//
// Solidity: function smtRoots(uint64 , bytes32 ) view returns(bool)
func (_SMT *SMTCallerSession) SmtRoots(arg0 uint64, arg1 [32]byte) (bool, error) {
	return _SMT.Contract.SmtRoots(&_SMT.CallOpts, arg0, arg1)
}

// Verifiers is a free data retrieval call binding the contract method 0x8195408d.
//
// Solidity: function verifiers(uint64 ) view returns(address)
func (_SMT *SMTCaller) Verifiers(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _SMT.contract.Call(opts, &out, "verifiers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0x8195408d.
//
// Solidity: function verifiers(uint64 ) view returns(address)
func (_SMT *SMTSession) Verifiers(arg0 uint64) (common.Address, error) {
	return _SMT.Contract.Verifiers(&_SMT.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0x8195408d.
//
// Solidity: function verifiers(uint64 ) view returns(address)
func (_SMT *SMTCallerSession) Verifiers(arg0 uint64) (common.Address, error) {
	return _SMT.Contract.Verifiers(&_SMT.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SMT *SMTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SMT *SMTSession) RenounceOwnership() (*types.Transaction, error) {
	return _SMT.Contract.RenounceOwnership(&_SMT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SMT *SMTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SMT.Contract.RenounceOwnership(&_SMT.TransactOpts)
}

// SetAnchorProvider is a paid mutator transaction binding the contract method 0x5ca32bd8.
//
// Solidity: function setAnchorProvider(uint64 chainId, address anchorProvider) returns()
func (_SMT *SMTTransactor) SetAnchorProvider(opts *bind.TransactOpts, chainId uint64, anchorProvider common.Address) (*types.Transaction, error) {
	return _SMT.contract.Transact(opts, "setAnchorProvider", chainId, anchorProvider)
}

// SetAnchorProvider is a paid mutator transaction binding the contract method 0x5ca32bd8.
//
// Solidity: function setAnchorProvider(uint64 chainId, address anchorProvider) returns()
func (_SMT *SMTSession) SetAnchorProvider(chainId uint64, anchorProvider common.Address) (*types.Transaction, error) {
	return _SMT.Contract.SetAnchorProvider(&_SMT.TransactOpts, chainId, anchorProvider)
}

// SetAnchorProvider is a paid mutator transaction binding the contract method 0x5ca32bd8.
//
// Solidity: function setAnchorProvider(uint64 chainId, address anchorProvider) returns()
func (_SMT *SMTTransactorSession) SetAnchorProvider(chainId uint64, anchorProvider common.Address) (*types.Transaction, error) {
	return _SMT.Contract.SetAnchorProvider(&_SMT.TransactOpts, chainId, anchorProvider)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x9c8413c5.
//
// Solidity: function setVerifier(uint64 chainId, address verifier) returns()
func (_SMT *SMTTransactor) SetVerifier(opts *bind.TransactOpts, chainId uint64, verifier common.Address) (*types.Transaction, error) {
	return _SMT.contract.Transact(opts, "setVerifier", chainId, verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x9c8413c5.
//
// Solidity: function setVerifier(uint64 chainId, address verifier) returns()
func (_SMT *SMTSession) SetVerifier(chainId uint64, verifier common.Address) (*types.Transaction, error) {
	return _SMT.Contract.SetVerifier(&_SMT.TransactOpts, chainId, verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x9c8413c5.
//
// Solidity: function setVerifier(uint64 chainId, address verifier) returns()
func (_SMT *SMTTransactorSession) SetVerifier(chainId uint64, verifier common.Address) (*types.Transaction, error) {
	return _SMT.Contract.SetVerifier(&_SMT.TransactOpts, chainId, verifier)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SMT *SMTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SMT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SMT *SMTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SMT.Contract.TransferOwnership(&_SMT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SMT *SMTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SMT.Contract.TransferOwnership(&_SMT.TransactOpts, newOwner)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_SMT *SMTTransactor) UpdateRoot(opts *bind.TransactOpts, chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _SMT.contract.Transact(opts, "updateRoot", chainId, u)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_SMT *SMTSession) UpdateRoot(chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _SMT.Contract.UpdateRoot(&_SMT.TransactOpts, chainId, u)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_SMT *SMTTransactorSession) UpdateRoot(chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _SMT.Contract.UpdateRoot(&_SMT.TransactOpts, chainId, u)
}

// SMTAnchorProviderUpdatedIterator is returned from FilterAnchorProviderUpdated and is used to iterate over the raw logs and unpacked data for AnchorProviderUpdated events raised by the SMT contract.
type SMTAnchorProviderUpdatedIterator struct {
	Event *SMTAnchorProviderUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMTAnchorProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMTAnchorProviderUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMTAnchorProviderUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMTAnchorProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMTAnchorProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMTAnchorProviderUpdated represents a AnchorProviderUpdated event raised by the SMT contract.
type SMTAnchorProviderUpdated struct {
	ChainId        uint64
	AnchorProvider common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAnchorProviderUpdated is a free log retrieval operation binding the contract event 0xd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c8090.
//
// Solidity: event AnchorProviderUpdated(uint64 chainId, address anchorProvider)
func (_SMT *SMTFilterer) FilterAnchorProviderUpdated(opts *bind.FilterOpts) (*SMTAnchorProviderUpdatedIterator, error) {

	logs, sub, err := _SMT.contract.FilterLogs(opts, "AnchorProviderUpdated")
	if err != nil {
		return nil, err
	}
	return &SMTAnchorProviderUpdatedIterator{contract: _SMT.contract, event: "AnchorProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchAnchorProviderUpdated is a free log subscription operation binding the contract event 0xd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c8090.
//
// Solidity: event AnchorProviderUpdated(uint64 chainId, address anchorProvider)
func (_SMT *SMTFilterer) WatchAnchorProviderUpdated(opts *bind.WatchOpts, sink chan<- *SMTAnchorProviderUpdated) (event.Subscription, error) {

	logs, sub, err := _SMT.contract.WatchLogs(opts, "AnchorProviderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMTAnchorProviderUpdated)
				if err := _SMT.contract.UnpackLog(event, "AnchorProviderUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnchorProviderUpdated is a log parse operation binding the contract event 0xd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c8090.
//
// Solidity: event AnchorProviderUpdated(uint64 chainId, address anchorProvider)
func (_SMT *SMTFilterer) ParseAnchorProviderUpdated(log types.Log) (*SMTAnchorProviderUpdated, error) {
	event := new(SMTAnchorProviderUpdated)
	if err := _SMT.contract.UnpackLog(event, "AnchorProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SMTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SMT contract.
type SMTOwnershipTransferredIterator struct {
	Event *SMTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMTOwnershipTransferred represents a OwnershipTransferred event raised by the SMT contract.
type SMTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SMT *SMTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SMTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SMT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SMTOwnershipTransferredIterator{contract: _SMT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SMT *SMTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SMTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SMT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMTOwnershipTransferred)
				if err := _SMT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SMT *SMTFilterer) ParseOwnershipTransferred(log types.Log) (*SMTOwnershipTransferred, error) {
	event := new(SMTOwnershipTransferred)
	if err := _SMT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SMTSmtRootUpdatedIterator is returned from FilterSmtRootUpdated and is used to iterate over the raw logs and unpacked data for SmtRootUpdated events raised by the SMT contract.
type SMTSmtRootUpdatedIterator struct {
	Event *SMTSmtRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMTSmtRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMTSmtRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMTSmtRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMTSmtRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMTSmtRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMTSmtRootUpdated represents a SmtRootUpdated event raised by the SMT contract.
type SMTSmtRootUpdated struct {
	SmtRoot     [32]byte
	EndBlockNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSmtRootUpdated is a free log retrieval operation binding the contract event 0x05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7.
//
// Solidity: event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum)
func (_SMT *SMTFilterer) FilterSmtRootUpdated(opts *bind.FilterOpts) (*SMTSmtRootUpdatedIterator, error) {

	logs, sub, err := _SMT.contract.FilterLogs(opts, "SmtRootUpdated")
	if err != nil {
		return nil, err
	}
	return &SMTSmtRootUpdatedIterator{contract: _SMT.contract, event: "SmtRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSmtRootUpdated is a free log subscription operation binding the contract event 0x05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7.
//
// Solidity: event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum)
func (_SMT *SMTFilterer) WatchSmtRootUpdated(opts *bind.WatchOpts, sink chan<- *SMTSmtRootUpdated) (event.Subscription, error) {

	logs, sub, err := _SMT.contract.WatchLogs(opts, "SmtRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMTSmtRootUpdated)
				if err := _SMT.contract.UnpackLog(event, "SmtRootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSmtRootUpdated is a log parse operation binding the contract event 0x05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7.
//
// Solidity: event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum)
func (_SMT *SMTFilterer) ParseSmtRootUpdated(log types.Log) (*SMTSmtRootUpdated, error) {
	event := new(SMTSmtRootUpdated)
	if err := _SMT.contract.UnpackLog(event, "SmtRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SMTVerifierUpdatedIterator is returned from FilterVerifierUpdated and is used to iterate over the raw logs and unpacked data for VerifierUpdated events raised by the SMT contract.
type SMTVerifierUpdatedIterator struct {
	Event *SMTVerifierUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMTVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMTVerifierUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMTVerifierUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMTVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMTVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMTVerifierUpdated represents a VerifierUpdated event raised by the SMT contract.
type SMTVerifierUpdated struct {
	ChainId  uint64
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierUpdated is a free log retrieval operation binding the contract event 0xb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be8162.
//
// Solidity: event VerifierUpdated(uint64 chainId, address verifier)
func (_SMT *SMTFilterer) FilterVerifierUpdated(opts *bind.FilterOpts) (*SMTVerifierUpdatedIterator, error) {

	logs, sub, err := _SMT.contract.FilterLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &SMTVerifierUpdatedIterator{contract: _SMT.contract, event: "VerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchVerifierUpdated is a free log subscription operation binding the contract event 0xb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be8162.
//
// Solidity: event VerifierUpdated(uint64 chainId, address verifier)
func (_SMT *SMTFilterer) WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *SMTVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _SMT.contract.WatchLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMTVerifierUpdated)
				if err := _SMT.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifierUpdated is a log parse operation binding the contract event 0xb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be8162.
//
// Solidity: event VerifierUpdated(uint64 chainId, address verifier)
func (_SMT *SMTFilterer) ParseVerifierUpdated(log types.Log) (*SMTVerifierUpdated, error) {
	event := new(SMTVerifierUpdated)
	if err := _SMT.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SameChainAnchorBlocksMetaData contains all meta data concerning the SameChainAnchorBlocks contract.
var SameChainAnchorBlocksMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005a575f8054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a361028e908161005f8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c8063715018a6146101965780638da5cb5b14610163578063f25b3f99146101445763f2fde38b14610046575f80fd5b346101415760203660031901126101415760043573ffffffffffffffffffffffffffffffffffffffff80821680920361013d5782549081169061008a33831461020d565b82156100d25773ffffffffffffffffffffffffffffffffffffffff1916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b8280fd5b80fd5b5034610141576020366003190112610141576020604051600435408152f35b503461014157806003193601126101415773ffffffffffffffffffffffffffffffffffffffff6020915416604051908152f35b503461014157806003193601126101415780805473ffffffffffffffffffffffffffffffffffffffff1973ffffffffffffffffffffffffffffffffffffffff8216916101e333841461020d565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b1561021457565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfea26469706673582212200fa6aa1d9eb09c17c6ea4780195a9605ab73a60ee9d03e6a24b185a3faa756a664736f6c63430008140033",
}

// SameChainAnchorBlocksABI is the input ABI used to generate the binding from.
// Deprecated: Use SameChainAnchorBlocksMetaData.ABI instead.
var SameChainAnchorBlocksABI = SameChainAnchorBlocksMetaData.ABI

// SameChainAnchorBlocksBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SameChainAnchorBlocksMetaData.Bin instead.
var SameChainAnchorBlocksBin = SameChainAnchorBlocksMetaData.Bin

// DeploySameChainAnchorBlocks deploys a new Ethereum contract, binding an instance of SameChainAnchorBlocks to it.
func DeploySameChainAnchorBlocks(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SameChainAnchorBlocks, error) {
	parsed, err := SameChainAnchorBlocksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SameChainAnchorBlocksBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SameChainAnchorBlocks{SameChainAnchorBlocksCaller: SameChainAnchorBlocksCaller{contract: contract}, SameChainAnchorBlocksTransactor: SameChainAnchorBlocksTransactor{contract: contract}, SameChainAnchorBlocksFilterer: SameChainAnchorBlocksFilterer{contract: contract}}, nil
}

// SameChainAnchorBlocks is an auto generated Go binding around an Ethereum contract.
type SameChainAnchorBlocks struct {
	SameChainAnchorBlocksCaller     // Read-only binding to the contract
	SameChainAnchorBlocksTransactor // Write-only binding to the contract
	SameChainAnchorBlocksFilterer   // Log filterer for contract events
}

// SameChainAnchorBlocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type SameChainAnchorBlocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SameChainAnchorBlocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SameChainAnchorBlocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SameChainAnchorBlocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SameChainAnchorBlocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SameChainAnchorBlocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SameChainAnchorBlocksSession struct {
	Contract     *SameChainAnchorBlocks // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SameChainAnchorBlocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SameChainAnchorBlocksCallerSession struct {
	Contract *SameChainAnchorBlocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SameChainAnchorBlocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SameChainAnchorBlocksTransactorSession struct {
	Contract     *SameChainAnchorBlocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SameChainAnchorBlocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type SameChainAnchorBlocksRaw struct {
	Contract *SameChainAnchorBlocks // Generic contract binding to access the raw methods on
}

// SameChainAnchorBlocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SameChainAnchorBlocksCallerRaw struct {
	Contract *SameChainAnchorBlocksCaller // Generic read-only contract binding to access the raw methods on
}

// SameChainAnchorBlocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SameChainAnchorBlocksTransactorRaw struct {
	Contract *SameChainAnchorBlocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSameChainAnchorBlocks creates a new instance of SameChainAnchorBlocks, bound to a specific deployed contract.
func NewSameChainAnchorBlocks(address common.Address, backend bind.ContractBackend) (*SameChainAnchorBlocks, error) {
	contract, err := bindSameChainAnchorBlocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SameChainAnchorBlocks{SameChainAnchorBlocksCaller: SameChainAnchorBlocksCaller{contract: contract}, SameChainAnchorBlocksTransactor: SameChainAnchorBlocksTransactor{contract: contract}, SameChainAnchorBlocksFilterer: SameChainAnchorBlocksFilterer{contract: contract}}, nil
}

// NewSameChainAnchorBlocksCaller creates a new read-only instance of SameChainAnchorBlocks, bound to a specific deployed contract.
func NewSameChainAnchorBlocksCaller(address common.Address, caller bind.ContractCaller) (*SameChainAnchorBlocksCaller, error) {
	contract, err := bindSameChainAnchorBlocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SameChainAnchorBlocksCaller{contract: contract}, nil
}

// NewSameChainAnchorBlocksTransactor creates a new write-only instance of SameChainAnchorBlocks, bound to a specific deployed contract.
func NewSameChainAnchorBlocksTransactor(address common.Address, transactor bind.ContractTransactor) (*SameChainAnchorBlocksTransactor, error) {
	contract, err := bindSameChainAnchorBlocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SameChainAnchorBlocksTransactor{contract: contract}, nil
}

// NewSameChainAnchorBlocksFilterer creates a new log filterer instance of SameChainAnchorBlocks, bound to a specific deployed contract.
func NewSameChainAnchorBlocksFilterer(address common.Address, filterer bind.ContractFilterer) (*SameChainAnchorBlocksFilterer, error) {
	contract, err := bindSameChainAnchorBlocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SameChainAnchorBlocksFilterer{contract: contract}, nil
}

// bindSameChainAnchorBlocks binds a generic wrapper to an already deployed contract.
func bindSameChainAnchorBlocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SameChainAnchorBlocksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SameChainAnchorBlocks *SameChainAnchorBlocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SameChainAnchorBlocks.Contract.SameChainAnchorBlocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SameChainAnchorBlocks *SameChainAnchorBlocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.SameChainAnchorBlocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SameChainAnchorBlocks *SameChainAnchorBlocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.SameChainAnchorBlocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SameChainAnchorBlocks *SameChainAnchorBlocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SameChainAnchorBlocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SameChainAnchorBlocks *SameChainAnchorBlocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SameChainAnchorBlocks *SameChainAnchorBlocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 blockNum) view returns(bytes32)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksCaller) Blocks(opts *bind.CallOpts, blockNum *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SameChainAnchorBlocks.contract.Call(opts, &out, "blocks", blockNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 blockNum) view returns(bytes32)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksSession) Blocks(blockNum *big.Int) ([32]byte, error) {
	return _SameChainAnchorBlocks.Contract.Blocks(&_SameChainAnchorBlocks.CallOpts, blockNum)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 blockNum) view returns(bytes32)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksCallerSession) Blocks(blockNum *big.Int) ([32]byte, error) {
	return _SameChainAnchorBlocks.Contract.Blocks(&_SameChainAnchorBlocks.CallOpts, blockNum)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SameChainAnchorBlocks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksSession) Owner() (common.Address, error) {
	return _SameChainAnchorBlocks.Contract.Owner(&_SameChainAnchorBlocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksCallerSession) Owner() (common.Address, error) {
	return _SameChainAnchorBlocks.Contract.Owner(&_SameChainAnchorBlocks.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SameChainAnchorBlocks *SameChainAnchorBlocksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SameChainAnchorBlocks *SameChainAnchorBlocksSession) RenounceOwnership() (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.RenounceOwnership(&_SameChainAnchorBlocks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SameChainAnchorBlocks *SameChainAnchorBlocksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.RenounceOwnership(&_SameChainAnchorBlocks.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SameChainAnchorBlocks *SameChainAnchorBlocksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SameChainAnchorBlocks *SameChainAnchorBlocksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.TransferOwnership(&_SameChainAnchorBlocks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SameChainAnchorBlocks *SameChainAnchorBlocksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SameChainAnchorBlocks.Contract.TransferOwnership(&_SameChainAnchorBlocks.TransactOpts, newOwner)
}

// SameChainAnchorBlocksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SameChainAnchorBlocks contract.
type SameChainAnchorBlocksOwnershipTransferredIterator struct {
	Event *SameChainAnchorBlocksOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SameChainAnchorBlocksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SameChainAnchorBlocksOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SameChainAnchorBlocksOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SameChainAnchorBlocksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SameChainAnchorBlocksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SameChainAnchorBlocksOwnershipTransferred represents a OwnershipTransferred event raised by the SameChainAnchorBlocks contract.
type SameChainAnchorBlocksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SameChainAnchorBlocksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SameChainAnchorBlocks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SameChainAnchorBlocksOwnershipTransferredIterator{contract: _SameChainAnchorBlocks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SameChainAnchorBlocksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SameChainAnchorBlocks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SameChainAnchorBlocksOwnershipTransferred)
				if err := _SameChainAnchorBlocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SameChainAnchorBlocks *SameChainAnchorBlocksFilterer) ParseOwnershipTransferred(log types.Log) (*SameChainAnchorBlocksOwnershipTransferred, error) {
	event := new(SameChainAnchorBlocksOwnershipTransferred)
	if err := _SameChainAnchorBlocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotValueVerifierMetaData contains all meta data concerning the SlotValueVerifier contract.
var SlotValueVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blocChunks\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateBlockChunks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateVerifierAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BlockChunks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_BlockChunks\",\"type\":\"address\"}],\"name\":\"updateBlockChunks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"}],\"name\":\"updateVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"verifierAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifySlotValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"addrHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slotKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slotValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"}],\"internalType\":\"structISlotValueVerifier.SlotInfo\",\"name\":\"slotInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100a457601f610b3538819003918201601f19168301916001600160401b038311848410176100a8578084926020946040528339810103126100a457516001600160a01b0390818116908190036100a4575f5460018060a01b03199033828216175f55604051933391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a36002541617600255610a7890816100bd8239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c9081630215d013146108fd575080630b885e53146103195780631eeb86da146102aa578063715018a61461024f5780638da5cb5b1461022a578063ec4ffc5214610195578063f2fde38b146100ba5763f5cec6af14610074575f80fd5b346100b65760203660031901126100b65767ffffffffffffffff610096610920565b165f52600160205260206001600160a01b0360405f205416604051908152f35b5f80fd5b346100b65760203660031901126100b6576100d3610965565b5f54906001600160a01b03808316916100ed33841461097b565b1691821561012a576001600160a01b0319839116175f557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b346100b65760403660031901126100b6576101ae610920565b602435906001600160a01b03918281168091036100b65767ffffffffffffffff6040926102007ffd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f955f5416331461097b565b1690815f526001602052825f20816001600160a01b031982541617905582519182526020820152a1005b346100b6575f3660031901126100b65760206001600160a01b035f5416604051908152f35b346100b6575f3660031901126100b6575f80546001600160a01b03196001600160a01b0382169161028133841461097b565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100b65760203660031901126100b6577f0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a9860206102e6610965565b6001600160a01b03906102fd825f5416331461097b565b16806001600160a01b03196002541617600255604051908152a1005b346100b65760603660031901126100b657610332610920565b60243567ffffffffffffffff81116100b657610352903690600401610937565b9160443567ffffffffffffffff81116100b657610373903690600401610937565b929060405194610382866109c6565b5f86525f60208701525f60408701525f60608701525f60808701525f60a087015267ffffffffffffffff84165f5260016020526001600160a01b0360405f205416156108b85767ffffffffffffffff84165f5260016020526001600160a01b0360405f205416602060405180926322bb937360e11b8252826004830152846024830152848760448401375f6044868401015281604481601f19601f89011681010301915afa90811561081b575f91610899575b50156100b6576040519360e0850185811067ffffffffffffffff8211176108405760405260e0368637610104860361085457856020116100b657856024116100b6575f5b60ff6006818316101561053f5761049761049283610a1c565b610a31565b90600181841601818111610500576104b26104928392610a1c565b1680828416116100b6578981116100b65791811686810135920360208110610528575b506007908316101561051457611fe08260051b1687015260ff80911690811461050057600101610479565b634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f1960209190910360031b1b9091169060076104d5565b50508684868589610104116100b65760e481013560c08301526040519560a0870187811067ffffffffffffffff8211176108405760405280610160116100b65780610180116100b65761017084013560801c61014085013560801b178752806101a0116100b657806101c0116100b6576101b084013560801c61018085013560801b176020880152806101e0116100b65780610200116100b6576101f084013560801c6101c085013560801b17604088015280610220116100b65780610240116100b65761023084013560801c61020085013560801b176060880152610260116100b657906020839261025c60a095013560e01c608089015260405193610645856109c6565b67ffffffffffffffff8816855261025c81013560e01c8386015261014061017082013560801c91013560801b17604085015280356060850152013560e01c6080830152828201526001600160a01b036002541690604051928391631513dce960e21b835267ffffffffffffffff815116600484015263ffffffff6020820151166024840152604081015160448401526060810151606484015263ffffffff6080820151166084840152015160a482015f905b6007821061082657505050610184816020935afa90811561081b575f916107ec575b50156107a75760a063ffffffff91606060c09567ffffffffffffffff86168352805160408401526020810151602084015284608082015116848401526040810151828401520151608082015267ffffffffffffffff604051941684526020810151602085015260408101516040850152606081015160608501526080810151608085015201511660a0820152f35b60405162461bcd60e51b815260206004820152600f60248201527f696e76616c696420626c6b4861736800000000000000000000000000000000006044820152606490fd5b61080e915060203d602011610814575b61080681836109e2565b810190610a04565b84610719565b503d6107fc565b6040513d5f823e3d90fd5b8293506020809160019394518152019301910184926106f7565b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152601760248201527f696e636f727265637420626c6b566572696679496e666f0000000000000000006044820152606490fd5b6108b2915060203d6020116108145761080681836109e2565b87610435565b60405162461bcd60e51b815260206004820152601660248201527f636861696e207665726966696572206e6f7420736574000000000000000000006044820152606490fd5b346100b6575f3660031901126100b6576020906001600160a01b03600254168152f35b6004359067ffffffffffffffff821682036100b657565b9181601f840112156100b65782359167ffffffffffffffff83116100b657602083818601950101116100b657565b600435906001600160a01b03821682036100b657565b1561098257565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60c0810190811067ffffffffffffffff82111761084057604052565b90601f8019910116810190811067ffffffffffffffff82111761084057604052565b908160209103126100b6575180151581036100b65790565b60051b90611fe060e083169216820361050057565b60ff166024019060ff82116105005756fea264697066735822122040824e79cd1f55281ec50dbb73269d37cc40708a75b388dd5261f665b1fd6c2264736f6c63430008140033",
}

// SlotValueVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use SlotValueVerifierMetaData.ABI instead.
var SlotValueVerifierABI = SlotValueVerifierMetaData.ABI

// SlotValueVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlotValueVerifierMetaData.Bin instead.
var SlotValueVerifierBin = SlotValueVerifierMetaData.Bin

// DeploySlotValueVerifier deploys a new Ethereum contract, binding an instance of SlotValueVerifier to it.
func DeploySlotValueVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _blocChunks common.Address) (common.Address, *types.Transaction, *SlotValueVerifier, error) {
	parsed, err := SlotValueVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlotValueVerifierBin), backend, _blocChunks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlotValueVerifier{SlotValueVerifierCaller: SlotValueVerifierCaller{contract: contract}, SlotValueVerifierTransactor: SlotValueVerifierTransactor{contract: contract}, SlotValueVerifierFilterer: SlotValueVerifierFilterer{contract: contract}}, nil
}

// SlotValueVerifier is an auto generated Go binding around an Ethereum contract.
type SlotValueVerifier struct {
	SlotValueVerifierCaller     // Read-only binding to the contract
	SlotValueVerifierTransactor // Write-only binding to the contract
	SlotValueVerifierFilterer   // Log filterer for contract events
}

// SlotValueVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlotValueVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotValueVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlotValueVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotValueVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlotValueVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotValueVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlotValueVerifierSession struct {
	Contract     *SlotValueVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SlotValueVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlotValueVerifierCallerSession struct {
	Contract *SlotValueVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SlotValueVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlotValueVerifierTransactorSession struct {
	Contract     *SlotValueVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SlotValueVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlotValueVerifierRaw struct {
	Contract *SlotValueVerifier // Generic contract binding to access the raw methods on
}

// SlotValueVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlotValueVerifierCallerRaw struct {
	Contract *SlotValueVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// SlotValueVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlotValueVerifierTransactorRaw struct {
	Contract *SlotValueVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlotValueVerifier creates a new instance of SlotValueVerifier, bound to a specific deployed contract.
func NewSlotValueVerifier(address common.Address, backend bind.ContractBackend) (*SlotValueVerifier, error) {
	contract, err := bindSlotValueVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifier{SlotValueVerifierCaller: SlotValueVerifierCaller{contract: contract}, SlotValueVerifierTransactor: SlotValueVerifierTransactor{contract: contract}, SlotValueVerifierFilterer: SlotValueVerifierFilterer{contract: contract}}, nil
}

// NewSlotValueVerifierCaller creates a new read-only instance of SlotValueVerifier, bound to a specific deployed contract.
func NewSlotValueVerifierCaller(address common.Address, caller bind.ContractCaller) (*SlotValueVerifierCaller, error) {
	contract, err := bindSlotValueVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifierCaller{contract: contract}, nil
}

// NewSlotValueVerifierTransactor creates a new write-only instance of SlotValueVerifier, bound to a specific deployed contract.
func NewSlotValueVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SlotValueVerifierTransactor, error) {
	contract, err := bindSlotValueVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifierTransactor{contract: contract}, nil
}

// NewSlotValueVerifierFilterer creates a new log filterer instance of SlotValueVerifier, bound to a specific deployed contract.
func NewSlotValueVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SlotValueVerifierFilterer, error) {
	contract, err := bindSlotValueVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifierFilterer{contract: contract}, nil
}

// bindSlotValueVerifier binds a generic wrapper to an already deployed contract.
func bindSlotValueVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlotValueVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlotValueVerifier *SlotValueVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlotValueVerifier.Contract.SlotValueVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlotValueVerifier *SlotValueVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.SlotValueVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlotValueVerifier *SlotValueVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.SlotValueVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlotValueVerifier *SlotValueVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlotValueVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlotValueVerifier *SlotValueVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlotValueVerifier *SlotValueVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.contract.Transact(opts, method, params...)
}

// BlockChunks is a free data retrieval call binding the contract method 0x0215d013.
//
// Solidity: function BlockChunks() view returns(address)
func (_SlotValueVerifier *SlotValueVerifierCaller) BlockChunks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlotValueVerifier.contract.Call(opts, &out, "BlockChunks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockChunks is a free data retrieval call binding the contract method 0x0215d013.
//
// Solidity: function BlockChunks() view returns(address)
func (_SlotValueVerifier *SlotValueVerifierSession) BlockChunks() (common.Address, error) {
	return _SlotValueVerifier.Contract.BlockChunks(&_SlotValueVerifier.CallOpts)
}

// BlockChunks is a free data retrieval call binding the contract method 0x0215d013.
//
// Solidity: function BlockChunks() view returns(address)
func (_SlotValueVerifier *SlotValueVerifierCallerSession) BlockChunks() (common.Address, error) {
	return _SlotValueVerifier.Contract.BlockChunks(&_SlotValueVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlotValueVerifier *SlotValueVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlotValueVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlotValueVerifier *SlotValueVerifierSession) Owner() (common.Address, error) {
	return _SlotValueVerifier.Contract.Owner(&_SlotValueVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlotValueVerifier *SlotValueVerifierCallerSession) Owner() (common.Address, error) {
	return _SlotValueVerifier.Contract.Owner(&_SlotValueVerifier.CallOpts)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_SlotValueVerifier *SlotValueVerifierCaller) VerifierAddresses(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _SlotValueVerifier.contract.Call(opts, &out, "verifierAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_SlotValueVerifier *SlotValueVerifierSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _SlotValueVerifier.Contract.VerifierAddresses(&_SlotValueVerifier.CallOpts, arg0)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_SlotValueVerifier *SlotValueVerifierCallerSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _SlotValueVerifier.Contract.VerifierAddresses(&_SlotValueVerifier.CallOpts, arg0)
}

// VerifySlotValue is a free data retrieval call binding the contract method 0x0b885e53.
//
// Solidity: function verifySlotValue(uint64 chainId, bytes proofData, bytes blkVerifyInfo) view returns((uint64,bytes32,bytes32,bytes32,bytes32,uint32) slotInfo)
func (_SlotValueVerifier *SlotValueVerifierCaller) VerifySlotValue(opts *bind.CallOpts, chainId uint64, proofData []byte, blkVerifyInfo []byte) (ISlotValueVerifierSlotInfo, error) {
	var out []interface{}
	err := _SlotValueVerifier.contract.Call(opts, &out, "verifySlotValue", chainId, proofData, blkVerifyInfo)

	if err != nil {
		return *new(ISlotValueVerifierSlotInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlotValueVerifierSlotInfo)).(*ISlotValueVerifierSlotInfo)

	return out0, err

}

// VerifySlotValue is a free data retrieval call binding the contract method 0x0b885e53.
//
// Solidity: function verifySlotValue(uint64 chainId, bytes proofData, bytes blkVerifyInfo) view returns((uint64,bytes32,bytes32,bytes32,bytes32,uint32) slotInfo)
func (_SlotValueVerifier *SlotValueVerifierSession) VerifySlotValue(chainId uint64, proofData []byte, blkVerifyInfo []byte) (ISlotValueVerifierSlotInfo, error) {
	return _SlotValueVerifier.Contract.VerifySlotValue(&_SlotValueVerifier.CallOpts, chainId, proofData, blkVerifyInfo)
}

// VerifySlotValue is a free data retrieval call binding the contract method 0x0b885e53.
//
// Solidity: function verifySlotValue(uint64 chainId, bytes proofData, bytes blkVerifyInfo) view returns((uint64,bytes32,bytes32,bytes32,bytes32,uint32) slotInfo)
func (_SlotValueVerifier *SlotValueVerifierCallerSession) VerifySlotValue(chainId uint64, proofData []byte, blkVerifyInfo []byte) (ISlotValueVerifierSlotInfo, error) {
	return _SlotValueVerifier.Contract.VerifySlotValue(&_SlotValueVerifier.CallOpts, chainId, proofData, blkVerifyInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlotValueVerifier *SlotValueVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotValueVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlotValueVerifier *SlotValueVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.RenounceOwnership(&_SlotValueVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlotValueVerifier *SlotValueVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.RenounceOwnership(&_SlotValueVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlotValueVerifier *SlotValueVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlotValueVerifier *SlotValueVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.TransferOwnership(&_SlotValueVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlotValueVerifier *SlotValueVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.TransferOwnership(&_SlotValueVerifier.TransactOpts, newOwner)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _BlockChunks) returns()
func (_SlotValueVerifier *SlotValueVerifierTransactor) UpdateBlockChunks(opts *bind.TransactOpts, _BlockChunks common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.contract.Transact(opts, "updateBlockChunks", _BlockChunks)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _BlockChunks) returns()
func (_SlotValueVerifier *SlotValueVerifierSession) UpdateBlockChunks(_BlockChunks common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.UpdateBlockChunks(&_SlotValueVerifier.TransactOpts, _BlockChunks)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _BlockChunks) returns()
func (_SlotValueVerifier *SlotValueVerifierTransactorSession) UpdateBlockChunks(_BlockChunks common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.UpdateBlockChunks(&_SlotValueVerifier.TransactOpts, _BlockChunks)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_SlotValueVerifier *SlotValueVerifierTransactor) UpdateVerifierAddress(opts *bind.TransactOpts, _chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.contract.Transact(opts, "updateVerifierAddress", _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_SlotValueVerifier *SlotValueVerifierSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.UpdateVerifierAddress(&_SlotValueVerifier.TransactOpts, _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_SlotValueVerifier *SlotValueVerifierTransactorSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _SlotValueVerifier.Contract.UpdateVerifierAddress(&_SlotValueVerifier.TransactOpts, _chainId, _verifierAddress)
}

// SlotValueVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SlotValueVerifier contract.
type SlotValueVerifierOwnershipTransferredIterator struct {
	Event *SlotValueVerifierOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlotValueVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotValueVerifierOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlotValueVerifierOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlotValueVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotValueVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotValueVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the SlotValueVerifier contract.
type SlotValueVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlotValueVerifier *SlotValueVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SlotValueVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SlotValueVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifierOwnershipTransferredIterator{contract: _SlotValueVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlotValueVerifier *SlotValueVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SlotValueVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SlotValueVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotValueVerifierOwnershipTransferred)
				if err := _SlotValueVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlotValueVerifier *SlotValueVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*SlotValueVerifierOwnershipTransferred, error) {
	event := new(SlotValueVerifierOwnershipTransferred)
	if err := _SlotValueVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotValueVerifierUpdateBlockChunksIterator is returned from FilterUpdateBlockChunks and is used to iterate over the raw logs and unpacked data for UpdateBlockChunks events raised by the SlotValueVerifier contract.
type SlotValueVerifierUpdateBlockChunksIterator struct {
	Event *SlotValueVerifierUpdateBlockChunks // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlotValueVerifierUpdateBlockChunksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotValueVerifierUpdateBlockChunks)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlotValueVerifierUpdateBlockChunks)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlotValueVerifierUpdateBlockChunksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotValueVerifierUpdateBlockChunksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotValueVerifierUpdateBlockChunks represents a UpdateBlockChunks event raised by the SlotValueVerifier contract.
type SlotValueVerifierUpdateBlockChunks struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateBlockChunks is a free log retrieval operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_SlotValueVerifier *SlotValueVerifierFilterer) FilterUpdateBlockChunks(opts *bind.FilterOpts) (*SlotValueVerifierUpdateBlockChunksIterator, error) {

	logs, sub, err := _SlotValueVerifier.contract.FilterLogs(opts, "UpdateBlockChunks")
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifierUpdateBlockChunksIterator{contract: _SlotValueVerifier.contract, event: "UpdateBlockChunks", logs: logs, sub: sub}, nil
}

// WatchUpdateBlockChunks is a free log subscription operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_SlotValueVerifier *SlotValueVerifierFilterer) WatchUpdateBlockChunks(opts *bind.WatchOpts, sink chan<- *SlotValueVerifierUpdateBlockChunks) (event.Subscription, error) {

	logs, sub, err := _SlotValueVerifier.contract.WatchLogs(opts, "UpdateBlockChunks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotValueVerifierUpdateBlockChunks)
				if err := _SlotValueVerifier.contract.UnpackLog(event, "UpdateBlockChunks", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateBlockChunks is a log parse operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_SlotValueVerifier *SlotValueVerifierFilterer) ParseUpdateBlockChunks(log types.Log) (*SlotValueVerifierUpdateBlockChunks, error) {
	event := new(SlotValueVerifierUpdateBlockChunks)
	if err := _SlotValueVerifier.contract.UnpackLog(event, "UpdateBlockChunks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotValueVerifierUpdateVerifierAddressIterator is returned from FilterUpdateVerifierAddress and is used to iterate over the raw logs and unpacked data for UpdateVerifierAddress events raised by the SlotValueVerifier contract.
type SlotValueVerifierUpdateVerifierAddressIterator struct {
	Event *SlotValueVerifierUpdateVerifierAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlotValueVerifierUpdateVerifierAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotValueVerifierUpdateVerifierAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlotValueVerifierUpdateVerifierAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlotValueVerifierUpdateVerifierAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotValueVerifierUpdateVerifierAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotValueVerifierUpdateVerifierAddress represents a UpdateVerifierAddress event raised by the SlotValueVerifier contract.
type SlotValueVerifierUpdateVerifierAddress struct {
	ChainId    uint64
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifierAddress is a free log retrieval operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_SlotValueVerifier *SlotValueVerifierFilterer) FilterUpdateVerifierAddress(opts *bind.FilterOpts) (*SlotValueVerifierUpdateVerifierAddressIterator, error) {

	logs, sub, err := _SlotValueVerifier.contract.FilterLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return &SlotValueVerifierUpdateVerifierAddressIterator{contract: _SlotValueVerifier.contract, event: "UpdateVerifierAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifierAddress is a free log subscription operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_SlotValueVerifier *SlotValueVerifierFilterer) WatchUpdateVerifierAddress(opts *bind.WatchOpts, sink chan<- *SlotValueVerifierUpdateVerifierAddress) (event.Subscription, error) {

	logs, sub, err := _SlotValueVerifier.contract.WatchLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotValueVerifierUpdateVerifierAddress)
				if err := _SlotValueVerifier.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateVerifierAddress is a log parse operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_SlotValueVerifier *SlotValueVerifierFilterer) ParseUpdateVerifierAddress(log types.Log) (*SlotValueVerifierUpdateVerifierAddress, error) {
	event := new(SlotValueVerifierUpdateVerifierAddress)
	if err := _SlotValueVerifier.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSMTMetaData contains all meta data concerning the TestSMT contract.
var TestSMTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"_anchorProviders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_initRoots\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"anchorProvider\",\"type\":\"address\"}],\"name\":\"AnchorProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endBlockNum\",\"type\":\"uint64\"}],\"name\":\"SmtRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNum\",\"type\":\"uint64\"}],\"name\":\"addRootForTesting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"anchorProviders\",\"outputs\":[{\"internalType\":\"contractIAnchorBlocks\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getLatestRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"}],\"name\":\"isSmtRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"latestRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"anchorProvider\",\"type\":\"address\"}],\"name\":\"setAnchorProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"smtRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newSmtRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"endBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextChunkMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"knowledgeProof\",\"type\":\"uint256[2]\"}],\"internalType\":\"structISMT.SmtUpdate\",\"name\":\"u\",\"type\":\"tuple\"}],\"name\":\"updateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60406080815234620002795762000fbb803803806200001e8162000297565b928339810191608082840312620002795781516001600160401b0391908281116200027957830184601f820112156200027957805194620000696200006387620002d1565b62000297565b9182968084526020808095019160051b8301019183831162000279578401905b8282106200027d5750505081850151848111620002795781620000ae918701620002e9565b9383860151818111620002795782620000c9918801620002e9565b956060810151908282116200027957019180601f8401121562000279578251620000f76200006382620002d1565b93858086848152019260051b820101928311620002795785809101915b838310620002685750505f8054336001600160a01b031980831682178455929a92956001600160a01b0395509293509084167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08b80a3620001798a5189511462000357565b620001888a5183511462000357565b620001978a5186511462000357565b885b8a51811015620002595781620001b0828d62000393565b511684620001bf838c62000393565b5116818c52600190818a528a8d20908882541617905585620001e2848762000393565b5116828d5260028a528a8d20908882541617905560038952898c2062000209848a62000393565b518d528952898c209060ff1982541617905562000227828862000393565b51908b5260048852888b20555f198114620002455760010162000199565b634e487b7160e01b8a52601160045260248afd5b8751610bfe9081620003bd8239f35b825181529181019186910162000114565b5f80fd5b815187811681036200027957815290840190840162000089565b6040519190601f01601f191682016001600160401b03811183821017620002bd57604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b038111620002bd5760051b60200190565b9080601f830112156200027957815190620003086200006383620002d1565b9182938184526020808095019260051b82010192831162000279578301905b82821062000336575050505090565b81516001600160a01b03811681036200027957815290830190830162000327565b156200035f57565b60405162461bcd60e51b815260206004820152600c60248201526b0d8cadc40dad2e6dac2e8c6d60a31b6044820152606490fd5b8051821015620003a85760209160051b010190565b634e487b7160e01b5f52603260045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c9081631019b61614610a8f575080633870253214610a45578063479aa6da1461095c5780635ca32bd8146109955780636ae3e0801461095c578063715018a6146109015780638195408d146108bf5780638da5cb5b1461089a57806397c7c3091461033c5780639c8413c51461028c578063afe8154b1461024a578063e31476da1461018c5763f2fde38b146100ab575f80fd5b34610188576020366003190112610188576004356001600160a01b03808216809203610188575f54908116906100e2338314610b7d565b821561011d576001600160a01b0319839116175f557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b5f80fd5b34610188576060366003190112610188577f05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a76101c6610ad4565b60243567ffffffffffffffff6101da610aeb565b926101f06001600160a01b035f54163314610b7d565b16805f52600360205260405f20825f5260205260405f20600160ff198254161790555f5260046020528060405f20556102456040519283928390929167ffffffffffffffff6020916040840195845216910152565b0390a1005b346101885760203660031901126101885767ffffffffffffffff61026c610ad4565b165f52600160205260206001600160a01b0360405f205416604051908152f35b34610188576040366003190112610188577fb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be81626102c6610ad4565b6102ce610b02565b906001600160a01b036102e5815f54163314610b7d565b67ffffffffffffffff82165f52600260205260405f209083166001600160a01b0319825416179055610245604051928392839092916001600160a01b0360209167ffffffffffffffff604085019616845216910152565b346101885761022036600319011261018857610356610ad4565b6102003660231901126101885760405160e0810181811067ffffffffffffffff8211176106ec57604052602435815261038d610aeb565b6020820152606435604082015260843560608201523660c3121561018857604051610100810181811067ffffffffffffffff8211176106ec57604052806101a4913683116101885760a4905b83821061088a5750506080830152366101c3121561018857604051906103fe82610b18565b8190366101e41161018857905b6101e4821061087a57505060a08201523661020312156101885760405161043181610b18565b803661022411610188576101e4905b610224821061086a57505060c0820152606081015115610745575b67ffffffffffffffff82165f52600460205260405f205460026020526001600160a01b0360405f205416801561070057604051918261012081011067ffffffffffffffff610120850111176106ec5761012083016040526101203684378060801c83526fffffffffffffffffffffffffffffffff8091166020840152835160801c6040840152808451166060840152604084015160801c60808401528060408501511660a084015267ffffffffffffffff60208501511660c0840152606084015160801c60e0840152606084015116610100830152608083015160a08401519260c0850151604051948593633072c1a360e11b8552600485015f905b600882106106d2575050509061057561058092610104860190610b56565b610144840190610b56565b5f61018483015b600982106106b8575050506102a4816020935afa9081156106ad575f91610672575b501561062d5767ffffffffffffffff9182165f818152600360209081526040808320855184528252808320805460ff1916600117905584519383526004825291829020929092558251928201518151938452909316908201527f05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a79181908101610245565b60405162461bcd60e51b815260206004820152601060248201527f696e76616c6964207a6b2070726f6f66000000000000000000000000000000006044820152606490fd5b90506020813d6020116106a5575b8161068d60209383610b34565b810103126101885751801515810361018857836105a9565b3d9150610680565b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610587565b825181528896506020928301926001929092019101610557565b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152601760248201527f6e6f20766572696669657220666f7220636861696e49640000000000000000006044820152606490fd5b67ffffffffffffffff82165f5260016020526001600160a01b0360405f205416801561082557602067ffffffffffffffff818401511660246040518094819363f25b3f9960e01b835260048301525afa9081156106ad575f916107f3575b5060408201511461045b5760405162461bcd60e51b815260206004820152601360248201527f616e63686f7220636865636b206661696c6564000000000000000000000000006044820152606490fd5b90506020813d60201161081d575b8161080e60209383610b34565b810103126101885751836107a3565b3d9150610801565b60405162461bcd60e51b815260206004820152601760248201527f756e6b6e6f776e20616e63686f722070726f76696465720000000000000000006044820152606490fd5b8135815260209182019101610440565b813581526020918201910161040b565b81358152602091820191016103d9565b34610188575f3660031901126101885760206001600160a01b035f5416604051908152f35b346101885760203660031901126101885767ffffffffffffffff6108e1610ad4565b165f52600260205260206001600160a01b0360405f205416604051908152f35b34610188575f366003190112610188575f80546001600160a01b03196001600160a01b03821691610933338414610b7d565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346101885760203660031901126101885767ffffffffffffffff61097e610ad4565b165f526004602052602060405f2054604051908152f35b34610188576040366003190112610188577fd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c80906109cf610ad4565b6109d7610b02565b906001600160a01b036109ee815f54163314610b7d565b67ffffffffffffffff82165f52600160205260405f209083166001600160a01b0319825416179055610245604051928392839092916001600160a01b0360209167ffffffffffffffff604085019616845216910152565b346101885760403660031901126101885767ffffffffffffffff610a67610ad4565b165f52600360205260405f206024355f52602052602060ff60405f2054166040519015158152f35b346101885760403660031901126101885760209067ffffffffffffffff610ab4610ad4565b165f526003825260405f206024355f52825260ff60405f20541615158152f35b6004359067ffffffffffffffff8216820361018857565b6044359067ffffffffffffffff8216820361018857565b602435906001600160a01b038216820361018857565b6040810190811067ffffffffffffffff8211176106ec57604052565b90601f8019910116810190811067ffffffffffffffff8211176106ec57604052565b5f915b60028310610b6657505050565b600190825181526020809101920192019190610b59565b15610b8457565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfea2646970667358221220fccc4eb631af4446280bfe22556eb2612569da142008a15c6c980719fa05c4d764736f6c63430008140033",
}

// TestSMTABI is the input ABI used to generate the binding from.
// Deprecated: Use TestSMTMetaData.ABI instead.
var TestSMTABI = TestSMTMetaData.ABI

// TestSMTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestSMTMetaData.Bin instead.
var TestSMTBin = TestSMTMetaData.Bin

// DeployTestSMT deploys a new Ethereum contract, binding an instance of TestSMT to it.
func DeployTestSMT(auth *bind.TransactOpts, backend bind.ContractBackend, _chainIds []uint64, _anchorProviders []common.Address, _verifiers []common.Address, _initRoots [][32]byte) (common.Address, *types.Transaction, *TestSMT, error) {
	parsed, err := TestSMTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestSMTBin), backend, _chainIds, _anchorProviders, _verifiers, _initRoots)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestSMT{TestSMTCaller: TestSMTCaller{contract: contract}, TestSMTTransactor: TestSMTTransactor{contract: contract}, TestSMTFilterer: TestSMTFilterer{contract: contract}}, nil
}

// TestSMT is an auto generated Go binding around an Ethereum contract.
type TestSMT struct {
	TestSMTCaller     // Read-only binding to the contract
	TestSMTTransactor // Write-only binding to the contract
	TestSMTFilterer   // Log filterer for contract events
}

// TestSMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestSMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestSMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestSMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSMTSession struct {
	Contract     *TestSMT          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestSMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestSMTCallerSession struct {
	Contract *TestSMTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestSMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestSMTTransactorSession struct {
	Contract     *TestSMTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestSMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestSMTRaw struct {
	Contract *TestSMT // Generic contract binding to access the raw methods on
}

// TestSMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestSMTCallerRaw struct {
	Contract *TestSMTCaller // Generic read-only contract binding to access the raw methods on
}

// TestSMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestSMTTransactorRaw struct {
	Contract *TestSMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestSMT creates a new instance of TestSMT, bound to a specific deployed contract.
func NewTestSMT(address common.Address, backend bind.ContractBackend) (*TestSMT, error) {
	contract, err := bindTestSMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestSMT{TestSMTCaller: TestSMTCaller{contract: contract}, TestSMTTransactor: TestSMTTransactor{contract: contract}, TestSMTFilterer: TestSMTFilterer{contract: contract}}, nil
}

// NewTestSMTCaller creates a new read-only instance of TestSMT, bound to a specific deployed contract.
func NewTestSMTCaller(address common.Address, caller bind.ContractCaller) (*TestSMTCaller, error) {
	contract, err := bindTestSMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestSMTCaller{contract: contract}, nil
}

// NewTestSMTTransactor creates a new write-only instance of TestSMT, bound to a specific deployed contract.
func NewTestSMTTransactor(address common.Address, transactor bind.ContractTransactor) (*TestSMTTransactor, error) {
	contract, err := bindTestSMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestSMTTransactor{contract: contract}, nil
}

// NewTestSMTFilterer creates a new log filterer instance of TestSMT, bound to a specific deployed contract.
func NewTestSMTFilterer(address common.Address, filterer bind.ContractFilterer) (*TestSMTFilterer, error) {
	contract, err := bindTestSMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestSMTFilterer{contract: contract}, nil
}

// bindTestSMT binds a generic wrapper to an already deployed contract.
func bindTestSMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestSMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSMT *TestSMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSMT.Contract.TestSMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSMT *TestSMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSMT.Contract.TestSMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSMT *TestSMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSMT.Contract.TestSMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSMT *TestSMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSMT *TestSMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSMT *TestSMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSMT.Contract.contract.Transact(opts, method, params...)
}

// AnchorProviders is a free data retrieval call binding the contract method 0xafe8154b.
//
// Solidity: function anchorProviders(uint64 ) view returns(address)
func (_TestSMT *TestSMTCaller) AnchorProviders(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "anchorProviders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnchorProviders is a free data retrieval call binding the contract method 0xafe8154b.
//
// Solidity: function anchorProviders(uint64 ) view returns(address)
func (_TestSMT *TestSMTSession) AnchorProviders(arg0 uint64) (common.Address, error) {
	return _TestSMT.Contract.AnchorProviders(&_TestSMT.CallOpts, arg0)
}

// AnchorProviders is a free data retrieval call binding the contract method 0xafe8154b.
//
// Solidity: function anchorProviders(uint64 ) view returns(address)
func (_TestSMT *TestSMTCallerSession) AnchorProviders(arg0 uint64) (common.Address, error) {
	return _TestSMT.Contract.AnchorProviders(&_TestSMT.CallOpts, arg0)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0x479aa6da.
//
// Solidity: function getLatestRoot(uint64 chainId) view returns(bytes32)
func (_TestSMT *TestSMTCaller) GetLatestRoot(opts *bind.CallOpts, chainId uint64) ([32]byte, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "getLatestRoot", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestRoot is a free data retrieval call binding the contract method 0x479aa6da.
//
// Solidity: function getLatestRoot(uint64 chainId) view returns(bytes32)
func (_TestSMT *TestSMTSession) GetLatestRoot(chainId uint64) ([32]byte, error) {
	return _TestSMT.Contract.GetLatestRoot(&_TestSMT.CallOpts, chainId)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0x479aa6da.
//
// Solidity: function getLatestRoot(uint64 chainId) view returns(bytes32)
func (_TestSMT *TestSMTCallerSession) GetLatestRoot(chainId uint64) ([32]byte, error) {
	return _TestSMT.Contract.GetLatestRoot(&_TestSMT.CallOpts, chainId)
}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_TestSMT *TestSMTCaller) IsSmtRootValid(opts *bind.CallOpts, chainId uint64, smtRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "isSmtRootValid", chainId, smtRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_TestSMT *TestSMTSession) IsSmtRootValid(chainId uint64, smtRoot [32]byte) (bool, error) {
	return _TestSMT.Contract.IsSmtRootValid(&_TestSMT.CallOpts, chainId, smtRoot)
}

// IsSmtRootValid is a free data retrieval call binding the contract method 0x1019b616.
//
// Solidity: function isSmtRootValid(uint64 chainId, bytes32 smtRoot) view returns(bool)
func (_TestSMT *TestSMTCallerSession) IsSmtRootValid(chainId uint64, smtRoot [32]byte) (bool, error) {
	return _TestSMT.Contract.IsSmtRootValid(&_TestSMT.CallOpts, chainId, smtRoot)
}

// LatestRoots is a free data retrieval call binding the contract method 0x6ae3e080.
//
// Solidity: function latestRoots(uint64 ) view returns(bytes32)
func (_TestSMT *TestSMTCaller) LatestRoots(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "latestRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRoots is a free data retrieval call binding the contract method 0x6ae3e080.
//
// Solidity: function latestRoots(uint64 ) view returns(bytes32)
func (_TestSMT *TestSMTSession) LatestRoots(arg0 uint64) ([32]byte, error) {
	return _TestSMT.Contract.LatestRoots(&_TestSMT.CallOpts, arg0)
}

// LatestRoots is a free data retrieval call binding the contract method 0x6ae3e080.
//
// Solidity: function latestRoots(uint64 ) view returns(bytes32)
func (_TestSMT *TestSMTCallerSession) LatestRoots(arg0 uint64) ([32]byte, error) {
	return _TestSMT.Contract.LatestRoots(&_TestSMT.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestSMT *TestSMTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestSMT *TestSMTSession) Owner() (common.Address, error) {
	return _TestSMT.Contract.Owner(&_TestSMT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestSMT *TestSMTCallerSession) Owner() (common.Address, error) {
	return _TestSMT.Contract.Owner(&_TestSMT.CallOpts)
}

// SmtRoots is a free data retrieval call binding the contract method 0x38702532.
//
// Solidity: function smtRoots(uint64 , bytes32 ) view returns(bool)
func (_TestSMT *TestSMTCaller) SmtRoots(opts *bind.CallOpts, arg0 uint64, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "smtRoots", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SmtRoots is a free data retrieval call binding the contract method 0x38702532.
//
// Solidity: function smtRoots(uint64 , bytes32 ) view returns(bool)
func (_TestSMT *TestSMTSession) SmtRoots(arg0 uint64, arg1 [32]byte) (bool, error) {
	return _TestSMT.Contract.SmtRoots(&_TestSMT.CallOpts, arg0, arg1)
}

// SmtRoots is a free data retrieval call binding the contract method 0x38702532.
//
// Solidity: function smtRoots(uint64 , bytes32 ) view returns(bool)
func (_TestSMT *TestSMTCallerSession) SmtRoots(arg0 uint64, arg1 [32]byte) (bool, error) {
	return _TestSMT.Contract.SmtRoots(&_TestSMT.CallOpts, arg0, arg1)
}

// Verifiers is a free data retrieval call binding the contract method 0x8195408d.
//
// Solidity: function verifiers(uint64 ) view returns(address)
func (_TestSMT *TestSMTCaller) Verifiers(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _TestSMT.contract.Call(opts, &out, "verifiers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0x8195408d.
//
// Solidity: function verifiers(uint64 ) view returns(address)
func (_TestSMT *TestSMTSession) Verifiers(arg0 uint64) (common.Address, error) {
	return _TestSMT.Contract.Verifiers(&_TestSMT.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0x8195408d.
//
// Solidity: function verifiers(uint64 ) view returns(address)
func (_TestSMT *TestSMTCallerSession) Verifiers(arg0 uint64) (common.Address, error) {
	return _TestSMT.Contract.Verifiers(&_TestSMT.CallOpts, arg0)
}

// AddRootForTesting is a paid mutator transaction binding the contract method 0xe31476da.
//
// Solidity: function addRootForTesting(uint64 chainId, bytes32 newRoot, uint64 endBlockNum) returns()
func (_TestSMT *TestSMTTransactor) AddRootForTesting(opts *bind.TransactOpts, chainId uint64, newRoot [32]byte, endBlockNum uint64) (*types.Transaction, error) {
	return _TestSMT.contract.Transact(opts, "addRootForTesting", chainId, newRoot, endBlockNum)
}

// AddRootForTesting is a paid mutator transaction binding the contract method 0xe31476da.
//
// Solidity: function addRootForTesting(uint64 chainId, bytes32 newRoot, uint64 endBlockNum) returns()
func (_TestSMT *TestSMTSession) AddRootForTesting(chainId uint64, newRoot [32]byte, endBlockNum uint64) (*types.Transaction, error) {
	return _TestSMT.Contract.AddRootForTesting(&_TestSMT.TransactOpts, chainId, newRoot, endBlockNum)
}

// AddRootForTesting is a paid mutator transaction binding the contract method 0xe31476da.
//
// Solidity: function addRootForTesting(uint64 chainId, bytes32 newRoot, uint64 endBlockNum) returns()
func (_TestSMT *TestSMTTransactorSession) AddRootForTesting(chainId uint64, newRoot [32]byte, endBlockNum uint64) (*types.Transaction, error) {
	return _TestSMT.Contract.AddRootForTesting(&_TestSMT.TransactOpts, chainId, newRoot, endBlockNum)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestSMT *TestSMTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSMT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestSMT *TestSMTSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestSMT.Contract.RenounceOwnership(&_TestSMT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestSMT *TestSMTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestSMT.Contract.RenounceOwnership(&_TestSMT.TransactOpts)
}

// SetAnchorProvider is a paid mutator transaction binding the contract method 0x5ca32bd8.
//
// Solidity: function setAnchorProvider(uint64 chainId, address anchorProvider) returns()
func (_TestSMT *TestSMTTransactor) SetAnchorProvider(opts *bind.TransactOpts, chainId uint64, anchorProvider common.Address) (*types.Transaction, error) {
	return _TestSMT.contract.Transact(opts, "setAnchorProvider", chainId, anchorProvider)
}

// SetAnchorProvider is a paid mutator transaction binding the contract method 0x5ca32bd8.
//
// Solidity: function setAnchorProvider(uint64 chainId, address anchorProvider) returns()
func (_TestSMT *TestSMTSession) SetAnchorProvider(chainId uint64, anchorProvider common.Address) (*types.Transaction, error) {
	return _TestSMT.Contract.SetAnchorProvider(&_TestSMT.TransactOpts, chainId, anchorProvider)
}

// SetAnchorProvider is a paid mutator transaction binding the contract method 0x5ca32bd8.
//
// Solidity: function setAnchorProvider(uint64 chainId, address anchorProvider) returns()
func (_TestSMT *TestSMTTransactorSession) SetAnchorProvider(chainId uint64, anchorProvider common.Address) (*types.Transaction, error) {
	return _TestSMT.Contract.SetAnchorProvider(&_TestSMT.TransactOpts, chainId, anchorProvider)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x9c8413c5.
//
// Solidity: function setVerifier(uint64 chainId, address verifier) returns()
func (_TestSMT *TestSMTTransactor) SetVerifier(opts *bind.TransactOpts, chainId uint64, verifier common.Address) (*types.Transaction, error) {
	return _TestSMT.contract.Transact(opts, "setVerifier", chainId, verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x9c8413c5.
//
// Solidity: function setVerifier(uint64 chainId, address verifier) returns()
func (_TestSMT *TestSMTSession) SetVerifier(chainId uint64, verifier common.Address) (*types.Transaction, error) {
	return _TestSMT.Contract.SetVerifier(&_TestSMT.TransactOpts, chainId, verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x9c8413c5.
//
// Solidity: function setVerifier(uint64 chainId, address verifier) returns()
func (_TestSMT *TestSMTTransactorSession) SetVerifier(chainId uint64, verifier common.Address) (*types.Transaction, error) {
	return _TestSMT.Contract.SetVerifier(&_TestSMT.TransactOpts, chainId, verifier)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestSMT *TestSMTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestSMT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestSMT *TestSMTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestSMT.Contract.TransferOwnership(&_TestSMT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestSMT *TestSMTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestSMT.Contract.TransferOwnership(&_TestSMT.TransactOpts, newOwner)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_TestSMT *TestSMTTransactor) UpdateRoot(opts *bind.TransactOpts, chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _TestSMT.contract.Transact(opts, "updateRoot", chainId, u)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_TestSMT *TestSMTSession) UpdateRoot(chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _TestSMT.Contract.UpdateRoot(&_TestSMT.TransactOpts, chainId, u)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x97c7c309.
//
// Solidity: function updateRoot(uint64 chainId, (bytes32,uint64,bytes32,bytes32,uint256[8],uint256[2],uint256[2]) u) returns()
func (_TestSMT *TestSMTTransactorSession) UpdateRoot(chainId uint64, u ISMTSmtUpdate) (*types.Transaction, error) {
	return _TestSMT.Contract.UpdateRoot(&_TestSMT.TransactOpts, chainId, u)
}

// TestSMTAnchorProviderUpdatedIterator is returned from FilterAnchorProviderUpdated and is used to iterate over the raw logs and unpacked data for AnchorProviderUpdated events raised by the TestSMT contract.
type TestSMTAnchorProviderUpdatedIterator struct {
	Event *TestSMTAnchorProviderUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestSMTAnchorProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSMTAnchorProviderUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestSMTAnchorProviderUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestSMTAnchorProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSMTAnchorProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSMTAnchorProviderUpdated represents a AnchorProviderUpdated event raised by the TestSMT contract.
type TestSMTAnchorProviderUpdated struct {
	ChainId        uint64
	AnchorProvider common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAnchorProviderUpdated is a free log retrieval operation binding the contract event 0xd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c8090.
//
// Solidity: event AnchorProviderUpdated(uint64 chainId, address anchorProvider)
func (_TestSMT *TestSMTFilterer) FilterAnchorProviderUpdated(opts *bind.FilterOpts) (*TestSMTAnchorProviderUpdatedIterator, error) {

	logs, sub, err := _TestSMT.contract.FilterLogs(opts, "AnchorProviderUpdated")
	if err != nil {
		return nil, err
	}
	return &TestSMTAnchorProviderUpdatedIterator{contract: _TestSMT.contract, event: "AnchorProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchAnchorProviderUpdated is a free log subscription operation binding the contract event 0xd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c8090.
//
// Solidity: event AnchorProviderUpdated(uint64 chainId, address anchorProvider)
func (_TestSMT *TestSMTFilterer) WatchAnchorProviderUpdated(opts *bind.WatchOpts, sink chan<- *TestSMTAnchorProviderUpdated) (event.Subscription, error) {

	logs, sub, err := _TestSMT.contract.WatchLogs(opts, "AnchorProviderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSMTAnchorProviderUpdated)
				if err := _TestSMT.contract.UnpackLog(event, "AnchorProviderUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnchorProviderUpdated is a log parse operation binding the contract event 0xd621c244f07f12e1f37bb9c40d61e278041fc4f2859a6736794b26f4297c8090.
//
// Solidity: event AnchorProviderUpdated(uint64 chainId, address anchorProvider)
func (_TestSMT *TestSMTFilterer) ParseAnchorProviderUpdated(log types.Log) (*TestSMTAnchorProviderUpdated, error) {
	event := new(TestSMTAnchorProviderUpdated)
	if err := _TestSMT.contract.UnpackLog(event, "AnchorProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSMTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestSMT contract.
type TestSMTOwnershipTransferredIterator struct {
	Event *TestSMTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestSMTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSMTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestSMTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestSMTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSMTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSMTOwnershipTransferred represents a OwnershipTransferred event raised by the TestSMT contract.
type TestSMTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestSMT *TestSMTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestSMTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestSMT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestSMTOwnershipTransferredIterator{contract: _TestSMT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestSMT *TestSMTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestSMTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestSMT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSMTOwnershipTransferred)
				if err := _TestSMT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestSMT *TestSMTFilterer) ParseOwnershipTransferred(log types.Log) (*TestSMTOwnershipTransferred, error) {
	event := new(TestSMTOwnershipTransferred)
	if err := _TestSMT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSMTSmtRootUpdatedIterator is returned from FilterSmtRootUpdated and is used to iterate over the raw logs and unpacked data for SmtRootUpdated events raised by the TestSMT contract.
type TestSMTSmtRootUpdatedIterator struct {
	Event *TestSMTSmtRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestSMTSmtRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSMTSmtRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestSMTSmtRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestSMTSmtRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSMTSmtRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSMTSmtRootUpdated represents a SmtRootUpdated event raised by the TestSMT contract.
type TestSMTSmtRootUpdated struct {
	SmtRoot     [32]byte
	EndBlockNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSmtRootUpdated is a free log retrieval operation binding the contract event 0x05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7.
//
// Solidity: event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum)
func (_TestSMT *TestSMTFilterer) FilterSmtRootUpdated(opts *bind.FilterOpts) (*TestSMTSmtRootUpdatedIterator, error) {

	logs, sub, err := _TestSMT.contract.FilterLogs(opts, "SmtRootUpdated")
	if err != nil {
		return nil, err
	}
	return &TestSMTSmtRootUpdatedIterator{contract: _TestSMT.contract, event: "SmtRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSmtRootUpdated is a free log subscription operation binding the contract event 0x05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7.
//
// Solidity: event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum)
func (_TestSMT *TestSMTFilterer) WatchSmtRootUpdated(opts *bind.WatchOpts, sink chan<- *TestSMTSmtRootUpdated) (event.Subscription, error) {

	logs, sub, err := _TestSMT.contract.WatchLogs(opts, "SmtRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSMTSmtRootUpdated)
				if err := _TestSMT.contract.UnpackLog(event, "SmtRootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSmtRootUpdated is a log parse operation binding the contract event 0x05712e4e1ed94f6f2706270c5ab0cbecd31613e0534ea4878ec4a16dc2f532a7.
//
// Solidity: event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum)
func (_TestSMT *TestSMTFilterer) ParseSmtRootUpdated(log types.Log) (*TestSMTSmtRootUpdated, error) {
	event := new(TestSMTSmtRootUpdated)
	if err := _TestSMT.contract.UnpackLog(event, "SmtRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSMTVerifierUpdatedIterator is returned from FilterVerifierUpdated and is used to iterate over the raw logs and unpacked data for VerifierUpdated events raised by the TestSMT contract.
type TestSMTVerifierUpdatedIterator struct {
	Event *TestSMTVerifierUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestSMTVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSMTVerifierUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestSMTVerifierUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestSMTVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSMTVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSMTVerifierUpdated represents a VerifierUpdated event raised by the TestSMT contract.
type TestSMTVerifierUpdated struct {
	ChainId  uint64
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierUpdated is a free log retrieval operation binding the contract event 0xb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be8162.
//
// Solidity: event VerifierUpdated(uint64 chainId, address verifier)
func (_TestSMT *TestSMTFilterer) FilterVerifierUpdated(opts *bind.FilterOpts) (*TestSMTVerifierUpdatedIterator, error) {

	logs, sub, err := _TestSMT.contract.FilterLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &TestSMTVerifierUpdatedIterator{contract: _TestSMT.contract, event: "VerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchVerifierUpdated is a free log subscription operation binding the contract event 0xb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be8162.
//
// Solidity: event VerifierUpdated(uint64 chainId, address verifier)
func (_TestSMT *TestSMTFilterer) WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *TestSMTVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _TestSMT.contract.WatchLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSMTVerifierUpdated)
				if err := _TestSMT.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifierUpdated is a log parse operation binding the contract event 0xb78ea0eaf11776732556ef6189312ceb60eab6b3177526b3a12d966e37be8162.
//
// Solidity: event VerifierUpdated(uint64 chainId, address verifier)
func (_TestSMT *TestSMTFilterer) ParseVerifierUpdated(log types.Log) (*TestSMTVerifierUpdated, error) {
	event := new(TestSMTVerifierUpdated)
	if err := _TestSMT.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TxVerifierMetaData contains all meta data concerning the TxVerifier contract.
var TxVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockChunks\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateBlockChunks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateVerifierAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"VerifiedTx\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockChunks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txRaw\",\"type\":\"bytes\"}],\"name\":\"decodeTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasTipCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFeeCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"}],\"internalType\":\"structITxVerifier.TxInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockChunks\",\"type\":\"address\"}],\"name\":\"updateBlockChunks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"}],\"name\":\"updateVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"verifierAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasTipCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFeeCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"}],\"internalType\":\"structITxVerifier.TxInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auxiBlkVerifyInfo\",\"type\":\"bytes\"}],\"name\":\"verifyTxAndLog\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasTipCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFeeCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"blkNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"blkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blkTime\",\"type\":\"uint64\"}],\"internalType\":\"structITxVerifier.TxInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346100a457601f611d3a38819003918201601f19168301916001600160401b038311848410176100a8578084926020946040528339810103126100a457516001600160a01b0390818116908190036100a4575f5460018060a01b03199033828216175f55604051933391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a36002541617600255611c7d90816100bd8239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f803560e01c9081631eeb86da146100b757508063361108de146100b2578063715018a6146100ad578063724796ed146100a85780638da5cb5b146100a3578063a8da8d691461009e578063dae029d314610099578063ec4ffc5214610094578063f2fde38b1461008f5763f5cec6af1461008a575f80fd5b6105fd565b610551565b6104bc565b610468565b61043c565b610417565b6103f1565b610394565b610309565b34610127576020366003190112610127577f0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a9860206100f361012a565b6001600160a01b039061010a82865416331461063f565b16806001600160a01b03196002541617600255604051908152a180f35b80fd5b600435906001600160a01b038216820361014057565b5f80fd5b9181601f840112156101405782359167ffffffffffffffff8311610140576020838186019501011161014057565b9060606003198301126101405767ffffffffffffffff600435818111610140578361019f91600401610144565b9390939260243583811161014057826101ba91600401610144565b93909392604435918211610140576101d491600401610144565b9091565b5f5b8381106101e95750505f910152565b81810151838201526020016101da565b90602091610212815180928185528580860191016101d8565b601f01601f1916010190565b610306906020815261023d60208201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660408201526040830151606082015260608301516080820152608083015160a082015261028960a084015160c08301906001600160a01b03169052565b60c083015160e082015260e083015192610180906102b46101009583878601526101a08501906101f9565b948101516102d061012091828601906001600160a01b03169052565b8101516102e8610140918286019063ffffffff169052565b81015161016084810191909152015167ffffffffffffffff16910152565b90565b34610140576103906103727fe1df3a08ea1a2c110c3f833d615f7e02814f32a3418b98f011888d0516669888604061035d61034336610172565b92610355999599969496929192610710565b50868a610993565b9467ffffffffffffffff865116923691610798565b6020815191012082519182526020820152a16040519182918261021e565b0390f35b34610140575f80600319360112610127578080546001600160a01b03196001600160a01b038216916103c733841461063f565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b34610140575f3660031901126101405760206001600160a01b0360025416604051908152f35b34610140575f3660031901126101405760206001600160a01b035f5416604051908152f35b346101405761039061045c61045036610172565b94939093929192610993565b6040519182918261021e565b346101405760203660031901126101405760043567ffffffffffffffff81116101405761045c61049f610390923690600401610144565b90611193565b6004359067ffffffffffffffff8216820361014057565b34610140576040366003190112610140576104d56104a5565b602435906001600160a01b03918281168091036101405767ffffffffffffffff6040926105277ffd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f955f5416331461063f565b1690815f526001602052825f20816001600160a01b031982541617905582519182526020820152a1005b346101405760203660031901126101405761056a61012a565b6001600160a01b03610580815f5416331461063f565b81161561059257610590906117e5565b005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b346101405760203660031901126101405767ffffffffffffffff61061f6104a5565b165f52600160205260206001600160a01b0360405f205416604051908152f35b1561064657565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff8211176106ba57604052565b61068a565b90601f8019910116810190811067ffffffffffffffff8211176106ba57604052565b6040519060c0820182811067ffffffffffffffff8211176106ba57604052565b6040519061070e8261069e565b565b60405190610180820182811067ffffffffffffffff8211176106ba57604052816101605f918281528260208201528260408201528260608201528260808201528260a08201528260c0820152606060e08201528261010082015282610120820152826101408201520152565b67ffffffffffffffff81116106ba57601f01601f191660200190565b9291926107a48261077c565b916107b260405193846106bf565b829481845281830111610140578281602093845f960137010152565b156107d557565b60405162461bcd60e51b815260206004820152600f60248201527f70726f6f66206e6f742076616c696400000000000000000000000000000000006044820152606490fd5b908092918237015f815290565b6020908361070e93959495604051968361084a89955180928880890191016101d8565b84019185830137015f838201520380855201836106bf565b1561086957565b60405162461bcd60e51b815260206004820152601260248201527f6c65616648617368206e6f74206d6174636800000000000000000000000000006044820152606490fd5b90816020910312610140575180151581036101405790565b91909161018081019267ffffffffffffffff815116825260a08063ffffffff926020938085830151168587015260408201516040870152606082015160608701526080820151166080860152015192015f905b600782106109275750505050565b82806001928651815201940191019092610919565b6040513d5f823e3d90fd5b1561094e57565b60405162461bcd60e51b815260206004820152600f60248201527f696e76616c696420626c6b4861736800000000000000000000000000000000006044820152606490fd5b909391959492956109a2610710565b506109ad8583611193565b96818189516109c39067ffffffffffffffff1690565b916109cd92611560565b6109d6906107ce565b6109df916116d3565b936080850151916109ef92610827565b918251602080940120845114610a0490610862565b610a0d91610dff565b86519293919267ffffffffffffffff169060408601938451610a329063ffffffff1690565b9584880196875190610a426106e1565b67ffffffffffffffff909616865263ffffffff16858701526040850152606084015263ffffffff16608083015260a082015260025482906001600160a01b03166001600160a01b03166001600160a01b0316604051631513dce960e21b81529283919082908190610ab690600483016108c6565b03915afa908115610b525761070e95610b1795610ae5610afb94606094610b0d975f92610b25575b5050610947565b516101408a0152015167ffffffffffffffff1690565b67ffffffffffffffff16610160870152565b5163ffffffff1690565b63ffffffff16610120840152565b610b449250803d10610b4b575b610b3c81836106bf565b8101906108ae565b5f80610ade565b503d610b32565b61093c565b15610b5e57565b60405162461bcd60e51b815260206004820152601b60248201527f696e636f72726563742061757869426c6b566572696679496e666f00000000006044820152606490fd5b906020116101405790602090565b906024116101405760200190600490565b90610104116101405760e40190602090565b909291928360011161014057831161014057600101915f190190565b909291928360021161014057831161014057600201916001190190565b906003116101405760020190600190565b906002116101405790600290565b909291928360031161014057831161014057600301916002190190565b9061016011610140576101400190602090565b9061018011610140576101700190601090565b906101a011610140576101800190602090565b906101c011610140576101b00190601090565b906101e011610140576101dc0190600490565b9061020011610140576101f80190600890565b929192610200918483116101405784116101405701916101ff190190565b90939293848311610140578411610140578101920390565b359060208110610cff575090565b5f199060200360031b1b1690565b7fffffffff000000000000000000000000000000000000000000000000000000009035818116939260048110610d4257505050565b60040360031b82901b16169150565b634e487b7160e01b5f52601160045260245ffd5b60ff1660ff8114610d765760010190565b610d51565b60051b90611fe060e0831692168203610d7657565b60ff166024019060ff8211610d7657565b60ff60019116019060ff8211610d7657565b60ff166002019060ff8211610d7657565b60ff1660c0019060ff8211610d7657565b634e487b7160e01b5f52603260045260245ffd5b906007811015610dfa5760051b0190565b610dd5565b916040519060e0820182811067ffffffffffffffff8211176106ba5760405260e036833781610e316101048514610b57565b610e44610e3e8587610ba3565b90610cf1565b93610e61610e5b610e558389610bb1565b90610d0d565b60e01c90565b945f5b60ff808216906006821015610ec65790610ebb610eb4610e3e85948d89610e95610e90610ec19a610d7b565b610d90565b9280610eab610e90610ea68c610da1565b610d7b565b16931691610cd9565b9189610de9565b52610d65565b610e64565b50505095610e3e60c092610edc92969496610bc2565b910152565b9015610dfa5790565b9060011015610dfa5760010190565b15610f0057565b60405162461bcd60e51b815260206004820152601660248201527f6e6f7420612044796e616d6963466565547854797065000000000000000000006044820152606490fd5b805115610dfa5760200190565b805160011015610dfa5760400190565b805160021015610dfa5760600190565b805160031015610dfa5760800190565b805160041015610dfa5760a00190565b805160051015610dfa5760c00190565b805160061015610dfa5760e00190565b805160071015610dfa576101000190565b805160091015610dfa576101400190565b8051600a1015610dfa576101600190565b8051600b1015610dfa576101800190565b8051821015610dfa5760209160051b010190565b602081519101519060208110610cff575090565b60ff60f6199116019060ff8211610d7657565b60ff6042199116019060ff8211610d7657565b6001600160f01b0319903581811693926002811061106157505050565b60020360031b82901b16169150565b61ffff9081166042190191908211610d7657565b6001600160f81b031990358181169392600181106110a157505050565b60010360031b82901b16169150565b604219810191908211610d7657565b60bf19810191908211610d7657565b607f19810191908211610d7657565b6020039060208211610d7657565b5f19810191908211610d7657565b60f619810191908211610d7657565b60b619810191908211610d7657565b6001600160f81b0319909116815260f960f81b60018201526001600160f01b031990911660028201526004929182908483013701015f815290565b6001600160f81b03199182168152601f60fb1b6001820152911660028201526003929182908483013701015f815290565b805160011015610dfa5760210190565b91906113e961070e916111a4610710565b9460ff906111d96002836111d26111cc6111be8887610ee1565b356001600160f81b03191690565b60f81c90565b1614610ef9565b6111fe6111f96111f46111ed868086610bd4565b3691610798565b611842565b6118a9565b9061123261122461121761121185610f45565b516119bc565b67ffffffffffffffff1690565b67ffffffffffffffff168952565b61125561124461121761121185610f52565b67ffffffffffffffff1660208a0152565b61126161121183610f62565b604089015261127261121183610f72565b606089015261128361121183610f82565b60808901526112ad61129d61129784610f92565b5161199e565b6001600160a01b031660a08a0152565b6112b961121183610fa2565b60c08901526112d06112ca83610fb2565b51611a2a565b60e08901526112ea6112e461121184610fc3565b60ff1690565b9361130e6113036112ca6113086113036112ca88610fd4565b61100a565b95610fe5565b936113276113226111cc6111be8587610eea565b61101e565b928184166001036114a8576113536112e461134e6111cc6113488786610c0d565b90611084565b611031565b935b61ffff8516603781116113fa5750506113c36113db926113bb836113ad611394611382896113d398610c1e565b93909961138e816110b0565b91610c2c565b6113a76040519a8b95602087019161081a565b9161081a565b03601f1981018752866106bf565b849516610dc4565b60f81b6001600160f81b03191690565b5f1a91611183565b535b602081519101206114de565b6001600160a01b0316610100840152565b9094929390841061146c576114679382826114276114216111be6114399661145998610ee1565b98610db3565b92611431826110b0565b931691610cd9565b60405195869491929160f81b6001600160f81b0319169060208601611152565b03601f1981018352826106bf565b6113dd565b6114679382826114276114216111be6114889661145998610ee1565b60405195869491929160f01b6001600160f01b0319169060208601611117565b6114d86114d36114cd6114c76114c06112e489610db3565b8786610bf0565b90611044565b60f01c90565b611070565b93611355565b919260ff8116601b811061151f575b509160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610b52575f5190565b601b9150929192019060ff8211610d7657919060206114ed565b90918060409360208452816020850152848401375f828201840152601f01601f1916010190565b909167ffffffffffffffff82165f5260016020526001600160a01b0360405f20541615611615576115e8926115cb6115bf6115bf6115b260209667ffffffffffffffff165f52600160205260405f2090565b546001600160a01b031690565b6001600160a01b031690565b906040518095819482936322bb937360e11b845260048401611539565b03915afa908115610b52575f916115fd575090565b610306915060203d8111610b4b57610b3c81836106bf565b60405162461bcd60e51b815260206004820152601660248201527f636861696e207665726966696572206e6f7420736574000000000000000000006044820152606490fd5b6fffffffffffffffffffffffffffffffff19903581811693926010811061168057505050565b60100360031b82901b16169150565b7fffffffffffffffff00000000000000000000000000000000000000000000000090358181169392600881106116c457505050565b60080360031b82901b16169150565b604051929160a0840167ffffffffffffffff8111858210176106ba576117e1916111ed916040525f86528060208701945f86526117dc604089015f81526117b260608b01915f835260808c019960608b528c61173b611735610e3e8a8a610c49565b60801b90565b61176f61175a61175461174e8c8c610c5c565b9061165a565b60801c90565b6fffffffffffffffffffffffffffffffff1690565b179052611782611735610e3e8989610c6f565b61179561175a61175461174e8b8b610c82565b1790526117a8610e5b610e558888610c95565b63ffffffff169052565b6117ce6117c86117c28686610ca8565b9061168f565b60c01c90565b67ffffffffffffffff169052565b610cbb565b9052565b5f54906001600160a01b0380911691826001600160a01b03198216175f55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b604051906118378261069e565b5f6020838281520152565b61184a61182a565b5060208151916040519261185d8461069e565b835201602082015290565b67ffffffffffffffff81116106ba5760051b60200190565b9060018201809211610d7657565b91908201809211610d7657565b5f198114610d765760010190565b6118b28161197a565b15610140576118c081611a77565b6118c981611868565b916118d760405193846106bf565b818352601f196118e683611868565b015f5b81811061196357505061190a60208092015161190481611b5d565b9061188e565b5f905b83821061191b575050505090565b6119578161192b61195d93611ad8565b90611934610701565b8281528187820152611946868a610ff6565b526119518589610ff6565b5061188e565b9161189b565b9061190d565b60209061196e61182a565b828288010152016118e9565b80511561199957602060c0910151515f1a1061199557600190565b5f90565b505f90565b6015815103610140576119b86001600160a01b03916119bc565b1690565b805180151590816119f2575b5015610140576119d7906119fe565b905190602081106119e6575090565b6020036101000a900490565b6021915011155f6119c8565b906020820191611a0e8351611b5d565b925190838201809211610d765751928303928311610d76579190565b80511561014057611a3d610306916119fe565b611a498193929361077c565b92611a5760405194856106bf565b818452601f19611a668361077c565b013660208601378360200190611bca565b805115611999575f9060208101908151611a9081611b5d565b8101809111610d7657915190518101809111610d765791905b828110611ab65750905090565b611abf81611ad8565b8101809111610d7657611ad2909161189b565b90611aa9565b80515f1a906080821015611aed575050600190565b60b8821015611b085750611b03610306916110ce565b611880565b9060c0811015611b2c5760b51991600160b783602003016101000a91015104010190565b9060f8821015611b435750611b03610306916110bf565b60010151602082900360f7016101000a90040160f5190190565b515f1a6080811015611b6e57505f90565b60b881108015611ba5575b15611b845750600190565b60c0811015611b9957611b0361030691611108565b611b03610306916110f9565b5060c08110158015611b79575060f88110611b79565b601f8111610d76576101000a90565b929091928315611c415792915b602093848410611c0c5780518252848101809111610d7657938101809111610d765791601f198101908111610d765791611bd7565b9193509180611c1a57505050565b611c2e611c29611c33926110dd565b611bbb565b6110eb565b905182518216911916179052565b5091505056fea264697066735822122000a53020841a446876a1689f600ffb2904cce2c0c184bf67b90b6c4e0506f32764736f6c63430008140033",
}

// TxVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use TxVerifierMetaData.ABI instead.
var TxVerifierABI = TxVerifierMetaData.ABI

// TxVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TxVerifierMetaData.Bin instead.
var TxVerifierBin = TxVerifierMetaData.Bin

// DeployTxVerifier deploys a new Ethereum contract, binding an instance of TxVerifier to it.
func DeployTxVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _blockChunks common.Address) (common.Address, *types.Transaction, *TxVerifier, error) {
	parsed, err := TxVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TxVerifierBin), backend, _blockChunks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TxVerifier{TxVerifierCaller: TxVerifierCaller{contract: contract}, TxVerifierTransactor: TxVerifierTransactor{contract: contract}, TxVerifierFilterer: TxVerifierFilterer{contract: contract}}, nil
}

// TxVerifier is an auto generated Go binding around an Ethereum contract.
type TxVerifier struct {
	TxVerifierCaller     // Read-only binding to the contract
	TxVerifierTransactor // Write-only binding to the contract
	TxVerifierFilterer   // Log filterer for contract events
}

// TxVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxVerifierSession struct {
	Contract     *TxVerifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxVerifierCallerSession struct {
	Contract *TxVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TxVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxVerifierTransactorSession struct {
	Contract     *TxVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TxVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxVerifierRaw struct {
	Contract *TxVerifier // Generic contract binding to access the raw methods on
}

// TxVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxVerifierCallerRaw struct {
	Contract *TxVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// TxVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxVerifierTransactorRaw struct {
	Contract *TxVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxVerifier creates a new instance of TxVerifier, bound to a specific deployed contract.
func NewTxVerifier(address common.Address, backend bind.ContractBackend) (*TxVerifier, error) {
	contract, err := bindTxVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxVerifier{TxVerifierCaller: TxVerifierCaller{contract: contract}, TxVerifierTransactor: TxVerifierTransactor{contract: contract}, TxVerifierFilterer: TxVerifierFilterer{contract: contract}}, nil
}

// NewTxVerifierCaller creates a new read-only instance of TxVerifier, bound to a specific deployed contract.
func NewTxVerifierCaller(address common.Address, caller bind.ContractCaller) (*TxVerifierCaller, error) {
	contract, err := bindTxVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxVerifierCaller{contract: contract}, nil
}

// NewTxVerifierTransactor creates a new write-only instance of TxVerifier, bound to a specific deployed contract.
func NewTxVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*TxVerifierTransactor, error) {
	contract, err := bindTxVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxVerifierTransactor{contract: contract}, nil
}

// NewTxVerifierFilterer creates a new log filterer instance of TxVerifier, bound to a specific deployed contract.
func NewTxVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*TxVerifierFilterer, error) {
	contract, err := bindTxVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxVerifierFilterer{contract: contract}, nil
}

// bindTxVerifier binds a generic wrapper to an already deployed contract.
func bindTxVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TxVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxVerifier *TxVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TxVerifier.Contract.TxVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxVerifier *TxVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxVerifier.Contract.TxVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxVerifier *TxVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxVerifier.Contract.TxVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxVerifier *TxVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TxVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxVerifier *TxVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxVerifier *TxVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxVerifier.Contract.contract.Transact(opts, method, params...)
}

// BlockChunks is a free data retrieval call binding the contract method 0x724796ed.
//
// Solidity: function blockChunks() view returns(address)
func (_TxVerifier *TxVerifierCaller) BlockChunks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxVerifier.contract.Call(opts, &out, "blockChunks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockChunks is a free data retrieval call binding the contract method 0x724796ed.
//
// Solidity: function blockChunks() view returns(address)
func (_TxVerifier *TxVerifierSession) BlockChunks() (common.Address, error) {
	return _TxVerifier.Contract.BlockChunks(&_TxVerifier.CallOpts)
}

// BlockChunks is a free data retrieval call binding the contract method 0x724796ed.
//
// Solidity: function blockChunks() view returns(address)
func (_TxVerifier *TxVerifierCallerSession) BlockChunks() (common.Address, error) {
	return _TxVerifier.Contract.BlockChunks(&_TxVerifier.CallOpts)
}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txRaw) pure returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierCaller) DecodeTx(opts *bind.CallOpts, txRaw []byte) (ITxVerifierTxInfo, error) {
	var out []interface{}
	err := _TxVerifier.contract.Call(opts, &out, "decodeTx", txRaw)

	if err != nil {
		return *new(ITxVerifierTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITxVerifierTxInfo)).(*ITxVerifierTxInfo)

	return out0, err

}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txRaw) pure returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierSession) DecodeTx(txRaw []byte) (ITxVerifierTxInfo, error) {
	return _TxVerifier.Contract.DecodeTx(&_TxVerifier.CallOpts, txRaw)
}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txRaw) pure returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierCallerSession) DecodeTx(txRaw []byte) (ITxVerifierTxInfo, error) {
	return _TxVerifier.Contract.DecodeTx(&_TxVerifier.CallOpts, txRaw)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TxVerifier *TxVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TxVerifier *TxVerifierSession) Owner() (common.Address, error) {
	return _TxVerifier.Contract.Owner(&_TxVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TxVerifier *TxVerifierCallerSession) Owner() (common.Address, error) {
	return _TxVerifier.Contract.Owner(&_TxVerifier.CallOpts)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_TxVerifier *TxVerifierCaller) VerifierAddresses(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _TxVerifier.contract.Call(opts, &out, "verifierAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_TxVerifier *TxVerifierSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _TxVerifier.Contract.VerifierAddresses(&_TxVerifier.CallOpts, arg0)
}

// VerifierAddresses is a free data retrieval call binding the contract method 0xf5cec6af.
//
// Solidity: function verifierAddresses(uint64 ) view returns(address)
func (_TxVerifier *TxVerifierCallerSession) VerifierAddresses(arg0 uint64) (common.Address, error) {
	return _TxVerifier.Contract.VerifierAddresses(&_TxVerifier.CallOpts, arg0)
}

// VerifyTx is a free data retrieval call binding the contract method 0xa8da8d69.
//
// Solidity: function verifyTx(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierCaller) VerifyTx(opts *bind.CallOpts, txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (ITxVerifierTxInfo, error) {
	var out []interface{}
	err := _TxVerifier.contract.Call(opts, &out, "verifyTx", txRaw, proofData, auxiBlkVerifyInfo)

	if err != nil {
		return *new(ITxVerifierTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITxVerifierTxInfo)).(*ITxVerifierTxInfo)

	return out0, err

}

// VerifyTx is a free data retrieval call binding the contract method 0xa8da8d69.
//
// Solidity: function verifyTx(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierSession) VerifyTx(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (ITxVerifierTxInfo, error) {
	return _TxVerifier.Contract.VerifyTx(&_TxVerifier.CallOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTx is a free data retrieval call binding the contract method 0xa8da8d69.
//
// Solidity: function verifyTx(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) view returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierCallerSession) VerifyTx(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (ITxVerifierTxInfo, error) {
	return _TxVerifier.Contract.VerifyTx(&_TxVerifier.CallOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TxVerifier *TxVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TxVerifier *TxVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _TxVerifier.Contract.RenounceOwnership(&_TxVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TxVerifier *TxVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TxVerifier.Contract.RenounceOwnership(&_TxVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TxVerifier *TxVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TxVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TxVerifier *TxVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TxVerifier.Contract.TransferOwnership(&_TxVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TxVerifier *TxVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TxVerifier.Contract.TransferOwnership(&_TxVerifier.TransactOpts, newOwner)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _blockChunks) returns()
func (_TxVerifier *TxVerifierTransactor) UpdateBlockChunks(opts *bind.TransactOpts, _blockChunks common.Address) (*types.Transaction, error) {
	return _TxVerifier.contract.Transact(opts, "updateBlockChunks", _blockChunks)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _blockChunks) returns()
func (_TxVerifier *TxVerifierSession) UpdateBlockChunks(_blockChunks common.Address) (*types.Transaction, error) {
	return _TxVerifier.Contract.UpdateBlockChunks(&_TxVerifier.TransactOpts, _blockChunks)
}

// UpdateBlockChunks is a paid mutator transaction binding the contract method 0x1eeb86da.
//
// Solidity: function updateBlockChunks(address _blockChunks) returns()
func (_TxVerifier *TxVerifierTransactorSession) UpdateBlockChunks(_blockChunks common.Address) (*types.Transaction, error) {
	return _TxVerifier.Contract.UpdateBlockChunks(&_TxVerifier.TransactOpts, _blockChunks)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_TxVerifier *TxVerifierTransactor) UpdateVerifierAddress(opts *bind.TransactOpts, _chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _TxVerifier.contract.Transact(opts, "updateVerifierAddress", _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_TxVerifier *TxVerifierSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _TxVerifier.Contract.UpdateVerifierAddress(&_TxVerifier.TransactOpts, _chainId, _verifierAddress)
}

// UpdateVerifierAddress is a paid mutator transaction binding the contract method 0xec4ffc52.
//
// Solidity: function updateVerifierAddress(uint64 _chainId, address _verifierAddress) returns()
func (_TxVerifier *TxVerifierTransactorSession) UpdateVerifierAddress(_chainId uint64, _verifierAddress common.Address) (*types.Transaction, error) {
	return _TxVerifier.Contract.UpdateVerifierAddress(&_TxVerifier.TransactOpts, _chainId, _verifierAddress)
}

// VerifyTxAndLog is a paid mutator transaction binding the contract method 0x361108de.
//
// Solidity: function verifyTxAndLog(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierTransactor) VerifyTxAndLog(opts *bind.TransactOpts, txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _TxVerifier.contract.Transact(opts, "verifyTxAndLog", txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTxAndLog is a paid mutator transaction binding the contract method 0x361108de.
//
// Solidity: function verifyTxAndLog(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierSession) VerifyTxAndLog(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _TxVerifier.Contract.VerifyTxAndLog(&_TxVerifier.TransactOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// VerifyTxAndLog is a paid mutator transaction binding the contract method 0x361108de.
//
// Solidity: function verifyTxAndLog(bytes txRaw, bytes proofData, bytes auxiBlkVerifyInfo) returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address,uint32,bytes32,uint64) info)
func (_TxVerifier *TxVerifierTransactorSession) VerifyTxAndLog(txRaw []byte, proofData []byte, auxiBlkVerifyInfo []byte) (*types.Transaction, error) {
	return _TxVerifier.Contract.VerifyTxAndLog(&_TxVerifier.TransactOpts, txRaw, proofData, auxiBlkVerifyInfo)
}

// TxVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TxVerifier contract.
type TxVerifierOwnershipTransferredIterator struct {
	Event *TxVerifierOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TxVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxVerifierOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TxVerifierOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TxVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the TxVerifier contract.
type TxVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TxVerifier *TxVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TxVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TxVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TxVerifierOwnershipTransferredIterator{contract: _TxVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TxVerifier *TxVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TxVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TxVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxVerifierOwnershipTransferred)
				if err := _TxVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TxVerifier *TxVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*TxVerifierOwnershipTransferred, error) {
	event := new(TxVerifierOwnershipTransferred)
	if err := _TxVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TxVerifierUpdateBlockChunksIterator is returned from FilterUpdateBlockChunks and is used to iterate over the raw logs and unpacked data for UpdateBlockChunks events raised by the TxVerifier contract.
type TxVerifierUpdateBlockChunksIterator struct {
	Event *TxVerifierUpdateBlockChunks // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TxVerifierUpdateBlockChunksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxVerifierUpdateBlockChunks)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TxVerifierUpdateBlockChunks)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TxVerifierUpdateBlockChunksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxVerifierUpdateBlockChunksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxVerifierUpdateBlockChunks represents a UpdateBlockChunks event raised by the TxVerifier contract.
type TxVerifierUpdateBlockChunks struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateBlockChunks is a free log retrieval operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_TxVerifier *TxVerifierFilterer) FilterUpdateBlockChunks(opts *bind.FilterOpts) (*TxVerifierUpdateBlockChunksIterator, error) {

	logs, sub, err := _TxVerifier.contract.FilterLogs(opts, "UpdateBlockChunks")
	if err != nil {
		return nil, err
	}
	return &TxVerifierUpdateBlockChunksIterator{contract: _TxVerifier.contract, event: "UpdateBlockChunks", logs: logs, sub: sub}, nil
}

// WatchUpdateBlockChunks is a free log subscription operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_TxVerifier *TxVerifierFilterer) WatchUpdateBlockChunks(opts *bind.WatchOpts, sink chan<- *TxVerifierUpdateBlockChunks) (event.Subscription, error) {

	logs, sub, err := _TxVerifier.contract.WatchLogs(opts, "UpdateBlockChunks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxVerifierUpdateBlockChunks)
				if err := _TxVerifier.contract.UnpackLog(event, "UpdateBlockChunks", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateBlockChunks is a log parse operation binding the contract event 0x0addee9cb6aa9328bbfbe8282179a6737af344619320736b6918af70b6a94a98.
//
// Solidity: event UpdateBlockChunks(address newAddress)
func (_TxVerifier *TxVerifierFilterer) ParseUpdateBlockChunks(log types.Log) (*TxVerifierUpdateBlockChunks, error) {
	event := new(TxVerifierUpdateBlockChunks)
	if err := _TxVerifier.contract.UnpackLog(event, "UpdateBlockChunks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TxVerifierUpdateVerifierAddressIterator is returned from FilterUpdateVerifierAddress and is used to iterate over the raw logs and unpacked data for UpdateVerifierAddress events raised by the TxVerifier contract.
type TxVerifierUpdateVerifierAddressIterator struct {
	Event *TxVerifierUpdateVerifierAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TxVerifierUpdateVerifierAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxVerifierUpdateVerifierAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TxVerifierUpdateVerifierAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TxVerifierUpdateVerifierAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxVerifierUpdateVerifierAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxVerifierUpdateVerifierAddress represents a UpdateVerifierAddress event raised by the TxVerifier contract.
type TxVerifierUpdateVerifierAddress struct {
	ChainId    uint64
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifierAddress is a free log retrieval operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_TxVerifier *TxVerifierFilterer) FilterUpdateVerifierAddress(opts *bind.FilterOpts) (*TxVerifierUpdateVerifierAddressIterator, error) {

	logs, sub, err := _TxVerifier.contract.FilterLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return &TxVerifierUpdateVerifierAddressIterator{contract: _TxVerifier.contract, event: "UpdateVerifierAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifierAddress is a free log subscription operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_TxVerifier *TxVerifierFilterer) WatchUpdateVerifierAddress(opts *bind.WatchOpts, sink chan<- *TxVerifierUpdateVerifierAddress) (event.Subscription, error) {

	logs, sub, err := _TxVerifier.contract.WatchLogs(opts, "UpdateVerifierAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxVerifierUpdateVerifierAddress)
				if err := _TxVerifier.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateVerifierAddress is a log parse operation binding the contract event 0xfd4bb2421fdcb098a8b68c19410f433a24b805b98f148eb28b7cc384b0a0e65f.
//
// Solidity: event UpdateVerifierAddress(uint64 chainId, address newAddress)
func (_TxVerifier *TxVerifierFilterer) ParseUpdateVerifierAddress(log types.Log) (*TxVerifierUpdateVerifierAddress, error) {
	event := new(TxVerifierUpdateVerifierAddress)
	if err := _TxVerifier.contract.UnpackLog(event, "UpdateVerifierAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TxVerifierVerifiedTxIterator is returned from FilterVerifiedTx and is used to iterate over the raw logs and unpacked data for VerifiedTx events raised by the TxVerifier contract.
type TxVerifierVerifiedTxIterator struct {
	Event *TxVerifierVerifiedTx // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TxVerifierVerifiedTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxVerifierVerifiedTx)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TxVerifierVerifiedTx)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TxVerifierVerifiedTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxVerifierVerifiedTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxVerifierVerifiedTx represents a VerifiedTx event raised by the TxVerifier contract.
type TxVerifierVerifiedTx struct {
	ChainId uint64
	TxHash  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVerifiedTx is a free log retrieval operation binding the contract event 0xe1df3a08ea1a2c110c3f833d615f7e02814f32a3418b98f011888d0516669888.
//
// Solidity: event VerifiedTx(uint64 chainId, bytes32 txHash)
func (_TxVerifier *TxVerifierFilterer) FilterVerifiedTx(opts *bind.FilterOpts) (*TxVerifierVerifiedTxIterator, error) {

	logs, sub, err := _TxVerifier.contract.FilterLogs(opts, "VerifiedTx")
	if err != nil {
		return nil, err
	}
	return &TxVerifierVerifiedTxIterator{contract: _TxVerifier.contract, event: "VerifiedTx", logs: logs, sub: sub}, nil
}

// WatchVerifiedTx is a free log subscription operation binding the contract event 0xe1df3a08ea1a2c110c3f833d615f7e02814f32a3418b98f011888d0516669888.
//
// Solidity: event VerifiedTx(uint64 chainId, bytes32 txHash)
func (_TxVerifier *TxVerifierFilterer) WatchVerifiedTx(opts *bind.WatchOpts, sink chan<- *TxVerifierVerifiedTx) (event.Subscription, error) {

	logs, sub, err := _TxVerifier.contract.WatchLogs(opts, "VerifiedTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxVerifierVerifiedTx)
				if err := _TxVerifier.contract.UnpackLog(event, "VerifiedTx", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifiedTx is a log parse operation binding the contract event 0xe1df3a08ea1a2c110c3f833d615f7e02814f32a3418b98f011888d0516669888.
//
// Solidity: event VerifiedTx(uint64 chainId, bytes32 txHash)
func (_TxVerifier *TxVerifierFilterer) ParseVerifiedTx(log types.Log) (*TxVerifierVerifiedTx, error) {
	event := new(TxVerifierVerifiedTx)
	if err := _TxVerifier.contract.UnpackLog(event, "VerifiedTx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
