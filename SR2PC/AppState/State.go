// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AppState

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

// CrossMessage is an auto generated low-level Go binding around an user-defined struct.
type CrossMessage struct {
	SourceChainId  *big.Int
	TargetChainId  *big.Int
	Phase          uint8
	SourceHeight   *big.Int
	Nonce          *big.Int
	Payload        [][]byte
	CmInputHash    [32]byte
	CmInputHeight  *big.Int
	RespPayload    [32]byte
	ExpectedHeight *big.Int
}

// StateShadowLock is an auto generated low-level Go binding around an user-defined struct.
type StateShadowLock struct {
	ChainId *big.Int
	Height  *big.Int
}

// StateShadowTranMeta is an auto generated low-level Go binding around an user-defined struct.
type StateShadowTranMeta struct {
	TransactionHash [32]byte
	Root            [32]byte
	Success         bool
}

// AppStateMetaData contains all meta data concerning the AppState contract.
var AppStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"appValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structState.ShadowLock\",\"name\":\"sdl\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"}],\"name\":\"getShadowLock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structState.ShadowLock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getShadowTranMeta\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structState.ShadowTranMeta\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getShadowValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"initValues\",\"type\":\"uint256[]\"}],\"name\":\"initValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"notifyUnconfirmedPrepare\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_tranHashesSuccess\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_tranHashesFailed\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"prepare\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"newCM\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"execSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"prepareUnconfirmed\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"newCM\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"execSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"rollback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611752806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c9e7263a11610071578063c9e7263a146101a2578063cbb1d90c146101b3578063d70f22ee146101ea578063edcfe2bb14610241578063f8cad932146102a2578063febc046f146102c3576100a9565b80630ff4c916146100ae5780631f71a643146100d45780639815bef114610148578063a5a587401461015d578063b2a4f5d41461017f575b600080fd5b6100c16100bc36600461135a565b6102d6565b6040519081526020015b60405180910390f35b61011e6100e236600461135a565b60006020818152918152604090819020815180830190925280548252600181015492820192909252600382015460059092015490919060ff1683565b604080518451815260209485015194810194909452830191909152151560608201526080016100cb565b61015b61015636600461129b565b6102ee565b005b61017061016b366004611372565b610365565b6040516100cb93929190611553565b61019261018d366004611372565b6106a3565b60405190151581526020016100cb565b61015b6101b0366004611372565b50565b6101c66101c13660046113f0565b6107c9565b604080518251815260208084015190820152918101511515908201526060016100cb565b6102346101f836600461135a565b60408051808201909152600080825260208201525060009081526020818152604091829020825180840190935280548352600101549082015290565b6040516100cb91906115fc565b6100c161024f3660046113f0565b600091825260208281526040808420805460018201548351808601929092528184015260608082019590955282518082039095018552608001825283519383019390932084526002909201905290205490565b6102b56102b0366004611411565b61085b565b6040516100cb929190611525565b6101706102d13660046113ad565b610ceb565b6000818152602081905260409020600301545b919050565b805160015560005b81518110156103615781818151811061031f57634e487b7160e01b600052603260045260246000fd5b60209081029190910181015160008381529182905260409091206003810191909155600501805460ff1916600117905580610359816116d5565b9150506102f6565b5050565b61036d6111cb565b600080808061037f60a0870187611613565b600081811061039e57634e487b7160e01b600052603260045260246000fd5b90506020028101906103b09190611661565b8101906103bd91906113f0565b600082815260208190526040902060050154919350915060ff166103ea57506000925082915061069c9050565b60008281526020819052604090205415801590610417575060008281526020819052604090206001015415155b1561042c5750600092506001915061069c9050565b600082815260208190526040902060030154600194508111156104525760009350610479565b600082815260208190526040812060030180548392906104739084906116be565b90915550505b6040805160038082526080820190925290816020015b606081526020019060019003908161048f57505060a0868101919091526104b890870187611613565b60008181106104d757634e487b7160e01b600052603260045260246000fd5b90506020028101906104e99190611661565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060a08a015180519094509092501515905061054457634e487b7160e01b600052603260045260246000fd5b602090810291909101015261055c60a0870187611613565b600181811061057b57634e487b7160e01b600052603260045260246000fd5b905060200281019061058d9190611661565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a0860151805160019081106105e757634e487b7160e01b600052603260045260246000fd5b60209081029190910101526105ff60a0870187611613565b600281811061061e57634e487b7160e01b600052603260045260246000fd5b90506020028101906106309190611661565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08601518051600290811061068a57634e487b7160e01b600052603260045260246000fd5b60209081029190910101525060009150505b9193909250565b600080806106b460a0850185611613565b60008181106106d357634e487b7160e01b600052603260045260246000fd5b90506020028101906106e59190611661565b8101906106f291906113f0565b600082815260208190526040902060050154919350915060ff1661075c5760405162461bcd60e51b815260206004820152601a60248201527f526f6c6c4261636b57697468496e76616c696456616c75654964000000000000604482015260640160405180910390fd5b60008281526020819052604090205415801590610789575060008281526020819052604090206001015415155b15610799576001925050506102e9565b600082815260208190526040812060030180548392906107ba9084906116a6565b90915550600095945050505050565b60408051606081018252600080825260208083018290528284018290528582528190529190912060040180548390811061081357634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805160608101825260039093029091018054835260018101549383019390935260029092015460ff1615159181019190915290505b92915050565b60608060008060005b60015481101561099c576000818152602081905260409020548814801561089b575060008181526020819052604090206001015487145b1561098a5760005b600082815260208190526040902060040154811015610988578660008084815260200190815260200160002060040182815481106108f157634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160010154141561097657600082815260208190526040902060040180548290811061093a57634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff16156109685783610960816116d5565b945050610976565b82610972816116d5565b9350505b80610980816116d5565b9150506108a3565b505b80610994816116d5565b915050610864565b508167ffffffffffffffff8111156109c457634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156109ed578160200160208202803683370190505b5093508067ffffffffffffffff811115610a1757634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a40578160200160208202803683370190505b509250600091506000905060005b600154811015610ce05760008181526020819052604090205488148015610a85575060008181526020819052604090206001015487145b15610cce576000818152602081815260408083208381556001810184905581518084018d90528083018c905260608082018c90528351808303909101815260809091018352805190840120808552600282018452918420548585529284905260030191909155905b600083815260208190526040902060040154811015610cae57876000808581526020019081526020016000206004018281548110610b3b57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610c9c576000838152602081905260409020600401805482908110610b8457634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610c20576000838152602081905260409020600401805482908110610bd157634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154878680610bef906116d5565b975081518110610c0f57634e487b7160e01b600052603260045260246000fd5b602002602001018181525050610c9c565b6000838152602081905260409020600401805482908110610c5157634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154868580610c6f906116d5565b965081518110610c8f57634e487b7160e01b600052603260045260246000fd5b6020026020010181815250505b80610ca6816116d5565b915050610aed565b506000828152602081905260408120610ccc9160049091019061123c565b505b80610cd8816116d5565b915050610a4e565b505050935093915050565b610cf36111cb565b60008080610d0460a0870187611613565b6001818110610d2357634e487b7160e01b600052603260045260246000fd5b9050602002810190610d359190611661565b810190610d42919061135a565b9050600080610d5460a0890189611613565b6000818110610d7357634e487b7160e01b600052603260045260246000fd5b9050602002810190610d859190611661565b810190610d9291906113f0565b600082815260208190526040902060050154919350915060ff16610dc05750600093508392506111c4915050565b600082815260208190526040902054158015610deb5750600082815260208190526040902060010154155b15610e52576000828152602081815260408083208b358082556101208d0135600183018190558351808601929092528184015260608082018d9052835180830390910181526080909101835280519084012060038201549085526002909101909252909120555b6000828152602081905260409020548835141580610e86575060008281526020819052604090206001015461012089013514155b15610e9c575060009350600192506111c4915050565b60408051893560208201526101208a0135918101919091526060810188905260009060800160408051601f1981840301815291815281516020928301206000868152808452828120828252600201909352912054909150600190831115610f0557506000610f35565b60008481526020818152604080832085845260020190915281208054859290610f2f9084906116be565b90915550505b6040805160038082526080820190925290816020015b6060815260200190600190039081610f4b57505060a089810191909152610f74908b018b611613565b6000818110610f9357634e487b7160e01b600052603260045260246000fd5b9050602002810190610fa59190611661565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060a08d015180519094509092501515905061100057634e487b7160e01b600052603260045260246000fd5b602090810291909101015261101860a08b018b611613565b600181811061103757634e487b7160e01b600052603260045260246000fd5b90506020028101906110499190611661565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a0890151805160019081106110a357634e487b7160e01b600052603260045260246000fd5b60209081029190910101526110bb60a08b018b611613565b60028181106110da57634e487b7160e01b600052603260045260246000fd5b90506020028101906110ec9190611661565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08901518051600290811061114657634e487b7160e01b600052603260045260246000fd5b6020908102919091018101919091526000948552848152604080862081516060810183529788528783018c8152841515928901928352600490910180546001808201835591895293882098516003909402909801928355519682019690965594516002909501805460ff191695151595909517909455509193509150505b9250925092565b60405180610140016040528060008152602001600081526020016000600481111561120657634e487b7160e01b600052602160045260246000fd5b8152600060208201819052604082018190526060808301526080820181905260a0820181905260c0820181905260e09091015290565b50805460008255600302906000526020600020908101906101b091905b8082111561127f576000808255600182015560028101805460ff19169055600301611259565b5090565b60006101408284031215611295578081fd5b50919050565b600060208083850312156112ad578182fd5b823567ffffffffffffffff808211156112c4578384fd5b818501915085601f8301126112d7578384fd5b8135818111156112e9576112e9611706565b838102604051601f19603f8301168101818110858211171561130d5761130d611706565b604052828152858101935084860182860187018a101561132b578788fd5b8795505b8386101561134d57803585526001959095019493860193860161132f565b5098975050505050505050565b60006020828403121561136b578081fd5b5035919050565b600060208284031215611383578081fd5b813567ffffffffffffffff811115611399578182fd5b6113a584828501611283565b949350505050565b600080604083850312156113bf578081fd5b823567ffffffffffffffff8111156113d5578182fd5b6113e185828601611283565b95602094909401359450505050565b60008060408385031215611402578182fd5b50508035926020909101359150565b600080600060608486031215611425578081fd5b505081359360208301359350604090920135919050565b6000815180845260208085019450808401835b8381101561146b5781518752958201959082019060010161144f565b509495945050505050565b60008282518085526020808601955080818302840101818601855b848110156114f657601f1980878503018a5282518051808652895b818110156114c75782810188015187820189015287016114ac565b818111156114d7578a8883890101525b509a86019a601f01909116939093018401925090830190600101611491565b5090979650505050505050565b6005811061152157634e487b7160e01b600052602160045260246000fd5b9052565b600060408252611538604083018561143c565b828103602084015261154a818561143c565b95945050505050565b6000606082528451606083015260208501516080830152604085015161157c60a0840182611503565b50606085015160c0830152608085015160e083015260a085015161014061010081818601526115af6101a0860184611476565b60c08901516101208781019190915260e08a0151938701939093529088015161016086015290870151610180850152851515602085015291506115ef9050565b82151560408301526113a5565b815181526020808301519082015260408101610855565b6000808335601e19843603018112611629578283fd5b83018035915067ffffffffffffffff821115611643578283fd5b602090810192508102360382131561165a57600080fd5b9250929050565b6000808335601e19843603018112611677578283fd5b83018035915067ffffffffffffffff821115611691578283fd5b60200191503681900382131561165a57600080fd5b600082198211156116b9576116b96116f0565b500190565b6000828210156116d0576116d06116f0565b500390565b60006000198214156116e9576116e96116f0565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220a91b95349970b3882d610958db66a167ca6c35a886fcffd1c7d38f123fc45fbc64736f6c63430008020033",
}

// AppStateABI is the input ABI used to generate the binding from.
// Deprecated: Use AppStateMetaData.ABI instead.
var AppStateABI = AppStateMetaData.ABI

// AppStateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AppStateMetaData.Bin instead.
var AppStateBin = AppStateMetaData.Bin

// DeployAppState deploys a new Ethereum contract, binding an instance of AppState to it.
func DeployAppState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AppState, error) {
	parsed, err := AppStateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AppStateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AppState{AppStateCaller: AppStateCaller{contract: contract}, AppStateTransactor: AppStateTransactor{contract: contract}, AppStateFilterer: AppStateFilterer{contract: contract}}, nil
}

// AppState is an auto generated Go binding around an Ethereum contract.
type AppState struct {
	AppStateCaller     // Read-only binding to the contract
	AppStateTransactor // Write-only binding to the contract
	AppStateFilterer   // Log filterer for contract events
}

// AppStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppStateSession struct {
	Contract     *AppState         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppStateCallerSession struct {
	Contract *AppStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AppStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppStateTransactorSession struct {
	Contract     *AppStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AppStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppStateRaw struct {
	Contract *AppState // Generic contract binding to access the raw methods on
}

// AppStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppStateCallerRaw struct {
	Contract *AppStateCaller // Generic read-only contract binding to access the raw methods on
}

// AppStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppStateTransactorRaw struct {
	Contract *AppStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppState creates a new instance of AppState, bound to a specific deployed contract.
func NewAppState(address common.Address, backend bind.ContractBackend) (*AppState, error) {
	contract, err := bindAppState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppState{AppStateCaller: AppStateCaller{contract: contract}, AppStateTransactor: AppStateTransactor{contract: contract}, AppStateFilterer: AppStateFilterer{contract: contract}}, nil
}

// NewAppStateCaller creates a new read-only instance of AppState, bound to a specific deployed contract.
func NewAppStateCaller(address common.Address, caller bind.ContractCaller) (*AppStateCaller, error) {
	contract, err := bindAppState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppStateCaller{contract: contract}, nil
}

// NewAppStateTransactor creates a new write-only instance of AppState, bound to a specific deployed contract.
func NewAppStateTransactor(address common.Address, transactor bind.ContractTransactor) (*AppStateTransactor, error) {
	contract, err := bindAppState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppStateTransactor{contract: contract}, nil
}

// NewAppStateFilterer creates a new log filterer instance of AppState, bound to a specific deployed contract.
func NewAppStateFilterer(address common.Address, filterer bind.ContractFilterer) (*AppStateFilterer, error) {
	contract, err := bindAppState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppStateFilterer{contract: contract}, nil
}

// bindAppState binds a generic wrapper to an already deployed contract.
func bindAppState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppState *AppStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppState.Contract.AppStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppState *AppStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppState.Contract.AppStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppState *AppStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppState.Contract.AppStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppState *AppStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppState *AppStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppState *AppStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppState.Contract.contract.Transact(opts, method, params...)
}

// AppValue is a free data retrieval call binding the contract method 0x1f71a643.
//
// Solidity: function appValue(uint256 ) view returns((uint256,uint256) sdl, uint256 value, bool exist)
func (_AppState *AppStateCaller) AppValue(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sdl   StateShadowLock
	Value *big.Int
	Exist bool
}, error) {
	var out []interface{}
	err := _AppState.contract.Call(opts, &out, "appValue", arg0)

	outstruct := new(struct {
		Sdl   StateShadowLock
		Value *big.Int
		Exist bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sdl = *abi.ConvertType(out[0], new(StateShadowLock)).(*StateShadowLock)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exist = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// AppValue is a free data retrieval call binding the contract method 0x1f71a643.
//
// Solidity: function appValue(uint256 ) view returns((uint256,uint256) sdl, uint256 value, bool exist)
func (_AppState *AppStateSession) AppValue(arg0 *big.Int) (struct {
	Sdl   StateShadowLock
	Value *big.Int
	Exist bool
}, error) {
	return _AppState.Contract.AppValue(&_AppState.CallOpts, arg0)
}

// AppValue is a free data retrieval call binding the contract method 0x1f71a643.
//
// Solidity: function appValue(uint256 ) view returns((uint256,uint256) sdl, uint256 value, bool exist)
func (_AppState *AppStateCallerSession) AppValue(arg0 *big.Int) (struct {
	Sdl   StateShadowLock
	Value *big.Int
	Exist bool
}, error) {
	return _AppState.Contract.AppValue(&_AppState.CallOpts, arg0)
}

// GetShadowLock is a free data retrieval call binding the contract method 0xd70f22ee.
//
// Solidity: function getShadowLock(uint256 valueId) view returns((uint256,uint256))
func (_AppState *AppStateCaller) GetShadowLock(opts *bind.CallOpts, valueId *big.Int) (StateShadowLock, error) {
	var out []interface{}
	err := _AppState.contract.Call(opts, &out, "getShadowLock", valueId)

	if err != nil {
		return *new(StateShadowLock), err
	}

	out0 := *abi.ConvertType(out[0], new(StateShadowLock)).(*StateShadowLock)

	return out0, err

}

// GetShadowLock is a free data retrieval call binding the contract method 0xd70f22ee.
//
// Solidity: function getShadowLock(uint256 valueId) view returns((uint256,uint256))
func (_AppState *AppStateSession) GetShadowLock(valueId *big.Int) (StateShadowLock, error) {
	return _AppState.Contract.GetShadowLock(&_AppState.CallOpts, valueId)
}

// GetShadowLock is a free data retrieval call binding the contract method 0xd70f22ee.
//
// Solidity: function getShadowLock(uint256 valueId) view returns((uint256,uint256))
func (_AppState *AppStateCallerSession) GetShadowLock(valueId *big.Int) (StateShadowLock, error) {
	return _AppState.Contract.GetShadowLock(&_AppState.CallOpts, valueId)
}

// GetShadowTranMeta is a free data retrieval call binding the contract method 0xcbb1d90c.
//
// Solidity: function getShadowTranMeta(uint256 valueId, uint256 index) view returns((bytes32,bytes32,bool))
func (_AppState *AppStateCaller) GetShadowTranMeta(opts *bind.CallOpts, valueId *big.Int, index *big.Int) (StateShadowTranMeta, error) {
	var out []interface{}
	err := _AppState.contract.Call(opts, &out, "getShadowTranMeta", valueId, index)

	if err != nil {
		return *new(StateShadowTranMeta), err
	}

	out0 := *abi.ConvertType(out[0], new(StateShadowTranMeta)).(*StateShadowTranMeta)

	return out0, err

}

// GetShadowTranMeta is a free data retrieval call binding the contract method 0xcbb1d90c.
//
// Solidity: function getShadowTranMeta(uint256 valueId, uint256 index) view returns((bytes32,bytes32,bool))
func (_AppState *AppStateSession) GetShadowTranMeta(valueId *big.Int, index *big.Int) (StateShadowTranMeta, error) {
	return _AppState.Contract.GetShadowTranMeta(&_AppState.CallOpts, valueId, index)
}

// GetShadowTranMeta is a free data retrieval call binding the contract method 0xcbb1d90c.
//
// Solidity: function getShadowTranMeta(uint256 valueId, uint256 index) view returns((bytes32,bytes32,bool))
func (_AppState *AppStateCallerSession) GetShadowTranMeta(valueId *big.Int, index *big.Int) (StateShadowTranMeta, error) {
	return _AppState.Contract.GetShadowTranMeta(&_AppState.CallOpts, valueId, index)
}

// GetShadowValue is a free data retrieval call binding the contract method 0xedcfe2bb.
//
// Solidity: function getShadowValue(uint256 valueId, bytes32 root) view returns(uint256)
func (_AppState *AppStateCaller) GetShadowValue(opts *bind.CallOpts, valueId *big.Int, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AppState.contract.Call(opts, &out, "getShadowValue", valueId, root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShadowValue is a free data retrieval call binding the contract method 0xedcfe2bb.
//
// Solidity: function getShadowValue(uint256 valueId, bytes32 root) view returns(uint256)
func (_AppState *AppStateSession) GetShadowValue(valueId *big.Int, root [32]byte) (*big.Int, error) {
	return _AppState.Contract.GetShadowValue(&_AppState.CallOpts, valueId, root)
}

// GetShadowValue is a free data retrieval call binding the contract method 0xedcfe2bb.
//
// Solidity: function getShadowValue(uint256 valueId, bytes32 root) view returns(uint256)
func (_AppState *AppStateCallerSession) GetShadowValue(valueId *big.Int, root [32]byte) (*big.Int, error) {
	return _AppState.Contract.GetShadowValue(&_AppState.CallOpts, valueId, root)
}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 valueId) view returns(uint256)
func (_AppState *AppStateCaller) GetValue(opts *bind.CallOpts, valueId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AppState.contract.Call(opts, &out, "getValue", valueId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 valueId) view returns(uint256)
func (_AppState *AppStateSession) GetValue(valueId *big.Int) (*big.Int, error) {
	return _AppState.Contract.GetValue(&_AppState.CallOpts, valueId)
}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 valueId) view returns(uint256)
func (_AppState *AppStateCallerSession) GetValue(valueId *big.Int) (*big.Int, error) {
	return _AppState.Contract.GetValue(&_AppState.CallOpts, valueId)
}

// Commit is a paid mutator transaction binding the contract method 0xc9e7263a.
//
// Solidity: function commit((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns()
func (_AppState *AppStateTransactor) Commit(opts *bind.TransactOpts, cm CrossMessage) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "commit", cm)
}

// Commit is a paid mutator transaction binding the contract method 0xc9e7263a.
//
// Solidity: function commit((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns()
func (_AppState *AppStateSession) Commit(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Commit(&_AppState.TransactOpts, cm)
}

// Commit is a paid mutator transaction binding the contract method 0xc9e7263a.
//
// Solidity: function commit((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns()
func (_AppState *AppStateTransactorSession) Commit(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Commit(&_AppState.TransactOpts, cm)
}

// InitValue is a paid mutator transaction binding the contract method 0x9815bef1.
//
// Solidity: function initValue(uint256[] initValues) returns()
func (_AppState *AppStateTransactor) InitValue(opts *bind.TransactOpts, initValues []*big.Int) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "initValue", initValues)
}

// InitValue is a paid mutator transaction binding the contract method 0x9815bef1.
//
// Solidity: function initValue(uint256[] initValues) returns()
func (_AppState *AppStateSession) InitValue(initValues []*big.Int) (*types.Transaction, error) {
	return _AppState.Contract.InitValue(&_AppState.TransactOpts, initValues)
}

// InitValue is a paid mutator transaction binding the contract method 0x9815bef1.
//
// Solidity: function initValue(uint256[] initValues) returns()
func (_AppState *AppStateTransactorSession) InitValue(initValues []*big.Int) (*types.Transaction, error) {
	return _AppState.Contract.InitValue(&_AppState.TransactOpts, initValues)
}

// NotifyUnconfirmedPrepare is a paid mutator transaction binding the contract method 0xf8cad932.
//
// Solidity: function notifyUnconfirmedPrepare(uint256 _chainId, uint256 _height, bytes32 _root) returns(bytes32[] _tranHashesSuccess, bytes32[] _tranHashesFailed)
func (_AppState *AppStateTransactor) NotifyUnconfirmedPrepare(opts *bind.TransactOpts, _chainId *big.Int, _height *big.Int, _root [32]byte) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "notifyUnconfirmedPrepare", _chainId, _height, _root)
}

// NotifyUnconfirmedPrepare is a paid mutator transaction binding the contract method 0xf8cad932.
//
// Solidity: function notifyUnconfirmedPrepare(uint256 _chainId, uint256 _height, bytes32 _root) returns(bytes32[] _tranHashesSuccess, bytes32[] _tranHashesFailed)
func (_AppState *AppStateSession) NotifyUnconfirmedPrepare(_chainId *big.Int, _height *big.Int, _root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.NotifyUnconfirmedPrepare(&_AppState.TransactOpts, _chainId, _height, _root)
}

// NotifyUnconfirmedPrepare is a paid mutator transaction binding the contract method 0xf8cad932.
//
// Solidity: function notifyUnconfirmedPrepare(uint256 _chainId, uint256 _height, bytes32 _root) returns(bytes32[] _tranHashesSuccess, bytes32[] _tranHashesFailed)
func (_AppState *AppStateTransactorSession) NotifyUnconfirmedPrepare(_chainId *big.Int, _height *big.Int, _root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.NotifyUnconfirmedPrepare(&_AppState.TransactOpts, _chainId, _height, _root)
}

// Prepare is a paid mutator transaction binding the contract method 0xa5a58740.
//
// Solidity: function prepare((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry)
func (_AppState *AppStateTransactor) Prepare(opts *bind.TransactOpts, cm CrossMessage) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "prepare", cm)
}

// Prepare is a paid mutator transaction binding the contract method 0xa5a58740.
//
// Solidity: function prepare((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry)
func (_AppState *AppStateSession) Prepare(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Prepare(&_AppState.TransactOpts, cm)
}

// Prepare is a paid mutator transaction binding the contract method 0xa5a58740.
//
// Solidity: function prepare((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry)
func (_AppState *AppStateTransactorSession) Prepare(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Prepare(&_AppState.TransactOpts, cm)
}

// PrepareUnconfirmed is a paid mutator transaction binding the contract method 0xfebc046f.
//
// Solidity: function prepareUnconfirmed((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm, bytes32 root) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry)
func (_AppState *AppStateTransactor) PrepareUnconfirmed(opts *bind.TransactOpts, cm CrossMessage, root [32]byte) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "prepareUnconfirmed", cm, root)
}

// PrepareUnconfirmed is a paid mutator transaction binding the contract method 0xfebc046f.
//
// Solidity: function prepareUnconfirmed((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm, bytes32 root) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry)
func (_AppState *AppStateSession) PrepareUnconfirmed(cm CrossMessage, root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.PrepareUnconfirmed(&_AppState.TransactOpts, cm, root)
}

// PrepareUnconfirmed is a paid mutator transaction binding the contract method 0xfebc046f.
//
// Solidity: function prepareUnconfirmed((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm, bytes32 root) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry)
func (_AppState *AppStateTransactorSession) PrepareUnconfirmed(cm CrossMessage, root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.PrepareUnconfirmed(&_AppState.TransactOpts, cm, root)
}

// Rollback is a paid mutator transaction binding the contract method 0xb2a4f5d4.
//
// Solidity: function rollback((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns(bool retry)
func (_AppState *AppStateTransactor) Rollback(opts *bind.TransactOpts, cm CrossMessage) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "rollback", cm)
}

// Rollback is a paid mutator transaction binding the contract method 0xb2a4f5d4.
//
// Solidity: function rollback((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns(bool retry)
func (_AppState *AppStateSession) Rollback(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Rollback(&_AppState.TransactOpts, cm)
}

// Rollback is a paid mutator transaction binding the contract method 0xb2a4f5d4.
//
// Solidity: function rollback((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns(bool retry)
func (_AppState *AppStateTransactorSession) Rollback(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Rollback(&_AppState.TransactOpts, cm)
}
