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
	Bin: "0x608060405234801561001057600080fd5b50611852806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c9e7263a11610071578063c9e7263a146101a2578063cbb1d90c146101b3578063d70f22ee146101ea578063edcfe2bb14610241578063f8cad932146102a2578063febc046f146102c3576100a9565b80630ff4c916146100ae5780631f71a643146100d45780639815bef114610148578063a5a587401461015d578063b2a4f5d41461017f575b600080fd5b6100c16100bc366004611473565b6102d6565b6040519081526020015b60405180910390f35b61011e6100e2366004611473565b60006020818152918152604090819020815180830190925280548252600181015492820192909252600382015460059092015490919060ff1683565b604080518451815260209485015194810194909452830191909152151560608201526080016100cb565b61015b610156366004611404565b6102ee565b005b61017061016b36600461148b565b61035f565b6040516100cb93929190611669565b61019261018d36600461148b565b61072e565b60405190151581526020016100cb565b61015b6101b036600461148b565b50565b6101c66101c1366004611509565b6108a3565b604080518251815260208084015190820152918101511515908201526060016100cb565b6102346101f8366004611473565b60408051808201909152600080825260208201525060009081526020818152604091829020825180840190935280548352600101549082015290565b6040516100cb9190611712565b6100c161024f366004611509565b600091825260208281526040808420805460018201548351808601929092528184015260608082019590955282518082039095018552608001825283519383019390932084526002909201905290205490565b6102b56102b036600461152a565b610935565b6040516100cb92919061163b565b6101706102d13660046114c6565b610dc5565b6000818152602081905260409020600301545b919050565b600181905560005b8181101561035a5782828281811061031e57634e487b7160e01b600052603260045260246000fd5b60008481526020818152604090912091029290920135600383015550600501805460ff1916600117905580610352816117eb565b9150506102f6565b505050565b610367611334565b6000808061037860a0860186611729565b600281811061039757634e487b7160e01b600052603260045260246000fd5b90506020028101906103a99190611777565b8101906103b69190611473565b60008181526020819052604090206005015490915060ff166103df575060009150819050610727565b6000818152602081905260409020541580159061040c575060008181526020819052604090206001015415155b1561041f57506000915060019050610727565b60019250600061043260a0870187611729565b600081811061045157634e487b7160e01b600052603260045260246000fd5b90506020028101906104639190611777565b8101906104709190611473565b60008381526020819052604090206003015490915081111561049557600093506104bc565b600082815260208190526040812060030180548392906104b69084906117d4565b90915550505b60408051600480825260a0820190925290816020015b60608152602001906001900390816104d25790505060a08601526040805160208101839052016040516020818303038152906040528560a0015160008151811061052c57634e487b7160e01b600052603260045260246000fd5b602090810291909101015261054460a0870187611729565b600181811061056357634e487b7160e01b600052603260045260246000fd5b90506020028101906105759190611777565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a0860151805160019081106105cf57634e487b7160e01b600052603260045260246000fd5b60209081029190910101526105e760a0870187611729565b600281811061060657634e487b7160e01b600052603260045260246000fd5b90506020028101906106189190611777565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08601518051600290811061067257634e487b7160e01b600052603260045260246000fd5b602090810291909101015261068a60a0870187611729565b60038181106106a957634e487b7160e01b600052603260045260246000fd5b90506020028101906106bb9190611777565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08601518051600390811061071557634e487b7160e01b600052603260045260246000fd5b60209081029190910101525060009150505b9193909250565b60008061073e60a0840184611729565b600281811061075d57634e487b7160e01b600052603260045260246000fd5b905060200281019061076f9190611777565b81019061077c9190611473565b60008181526020819052604090206005015490915060ff166107e45760405162461bcd60e51b815260206004820152601a60248201527f526f6c6c4261636b57697468496e76616c696456616c75654964000000000000604482015260640160405180910390fd5b60008181526020819052604090205415801590610811575060008181526020819052604090206001015415155b156108205760019150506102e9565b600061082f60a0850185611729565b600081811061084e57634e487b7160e01b600052603260045260246000fd5b90506020028101906108609190611777565b81019061086d9190611473565b905080600080848152602001908152602001600020600301600082825461089491906117bc565b90915550600095945050505050565b6040805160608101825260008082526020808301829052828401829052858252819052919091206004018054839081106108ed57634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805160608101825260039093029091018054835260018101549383019390935260029092015460ff1615159181019190915290505b92915050565b60608060008060005b600154811015610a765760008181526020819052604090205488148015610975575060008181526020819052604090206001015487145b15610a645760005b600082815260208190526040902060040154811015610a62578660008084815260200190815260200160002060040182815481106109cb57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610a50576000828152602081905260409020600401805482908110610a1457634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610a425783610a3a816117eb565b945050610a50565b82610a4c816117eb565b9350505b80610a5a816117eb565b91505061097d565b505b80610a6e816117eb565b91505061093e565b508167ffffffffffffffff811115610a9e57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610ac7578160200160208202803683370190505b5093508067ffffffffffffffff811115610af157634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610b1a578160200160208202803683370190505b509250600091506000905060005b600154811015610dba5760008181526020819052604090205488148015610b5f575060008181526020819052604090206001015487145b15610da8576000818152602081815260408083208381556001810184905581518084018d90528083018c905260608082018c90528351808303909101815260809091018352805190840120808552600282018452918420548585529284905260030191909155905b600083815260208190526040902060040154811015610d8857876000808581526020019081526020016000206004018281548110610c1557634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610d76576000838152602081905260409020600401805482908110610c5e57634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610cfa576000838152602081905260409020600401805482908110610cab57634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154878680610cc9906117eb565b975081518110610ce957634e487b7160e01b600052603260045260246000fd5b602002602001018181525050610d76565b6000838152602081905260409020600401805482908110610d2b57634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154868580610d49906117eb565b965081518110610d6957634e487b7160e01b600052603260045260246000fd5b6020026020010181815250505b80610d80816117eb565b915050610bc7565b506000828152602081905260408120610da6916004909101906113a5565b505b80610db2816117eb565b915050610b28565b505050935093915050565b610dcd611334565b60008080610dde60a0870187611729565b6001818110610dfd57634e487b7160e01b600052603260045260246000fd5b9050602002810190610e0f9190611777565b810190610e1c9190611473565b90506000610e2d60a0880188611729565b6002818110610e4c57634e487b7160e01b600052603260045260246000fd5b9050602002810190610e5e9190611777565b810190610e6b9190611473565b60008181526020819052604090206005015490915060ff16610e9657506000925082915061132d9050565b600081815260208190526040902054158015610ec15750600081815260208190526040902060010154155b15610f28576000818152602081815260408083208a358082556101208c0135600183018190558351808601929092528184015260608082018c9052835180830390910181526080909101835280519084012060038201549085526002909101909252909120555b6000818152602081905260409020548735141580610f5c575060008181526020819052604090206001015461012088013514155b15610f715750600092506001915061132d9050565b6040805188356020808301919091526101208a01358284015260608083018a905283518084039091018152608090920190925280519101206000610fb860a08a018a611729565b6000818110610fd757634e487b7160e01b600052603260045260246000fd5b9050602002810190610fe99190611777565b810190610ff69190611473565b60008481526020818152604080832086845260020190915290205490915060019082111561102657506000611056565b600084815260208181526040808320868452600201909152812080548492906110509084906117d4565b90915550505b60408051600480825260a0820190925290816020015b606081526020019060019003908161106c5790505060a08901526040805160208101849052016040516020818303038152906040528860a001516000815181106110c657634e487b7160e01b600052603260045260246000fd5b60209081029190910101526110de60a08b018b611729565b60018181106110fd57634e487b7160e01b600052603260045260246000fd5b905060200281019061110f9190611777565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08901518051600190811061116957634e487b7160e01b600052603260045260246000fd5b602090810291909101015261118160a08b018b611729565b60028181106111a057634e487b7160e01b600052603260045260246000fd5b90506020028101906111b29190611777565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08901518051600290811061120c57634e487b7160e01b600052603260045260246000fd5b602090810291909101015261122460a08b018b611729565b600381811061124357634e487b7160e01b600052603260045260246000fd5b90506020028101906112559190611777565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a0890151805160039081106112af57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101919091526000948552848152604080862081516060810183529788528783018c8152841515928901928352600490910180546001808201835591895293882098516003909402909801928355519682019690965594516002909501805460ff191695151595909517909455509193509150505b9250925092565b60405180610140016040528060008152602001600081526020016000600481111561136f57634e487b7160e01b600052602160045260246000fd5b8152600060208201819052604082018190526060808301526080820181905260a0820181905260c0820181905260e09091015290565b50805460008255600302906000526020600020908101906101b091905b808211156113e8576000808255600182015560028101805460ff191690556003016113c2565b5090565b600061014082840312156113fe578081fd5b50919050565b60008060208385031215611416578182fd5b823567ffffffffffffffff8082111561142d578384fd5b818501915085601f830112611440578384fd5b81358181111561144e578485fd5b8660208083028501011115611461578485fd5b60209290920196919550909350505050565b600060208284031215611484578081fd5b5035919050565b60006020828403121561149c578081fd5b813567ffffffffffffffff8111156114b2578182fd5b6114be848285016113ec565b949350505050565b600080604083850312156114d8578182fd5b823567ffffffffffffffff8111156114ee578283fd5b6114fa858286016113ec565b95602094909401359450505050565b6000806040838503121561151b578182fd5b50508035926020909101359150565b60008060006060848603121561153e578081fd5b505081359360208301359350604090920135919050565b6000815180845260208085019450808401835b8381101561158457815187529582019590820190600101611568565b509495945050505050565b6000815180845260208085018081965082840281019150828601855b8581101561160c57828403895281518051808652885b818110156115dc5782810188015187820189015287016115c1565b818111156115ec57898883890101525b5099860199601f01601f19169490940185019350908401906001016115ab565b5091979650505050505050565b6005811061163757634e487b7160e01b600052602160045260246000fd5b9052565b60006040825261164e6040830185611555565b82810360208401526116608185611555565b95945050505050565b6000606082528451606083015260208501516080830152604085015161169260a0840182611619565b50606085015160c0830152608085015160e083015260a085015161014061010081818601526116c56101a086018461158f565b60c08901516101208781019190915260e08a0151938701939093529088015161016086015290870151610180850152851515602085015291506117059050565b82151560408301526114be565b81518152602080830151908201526040810161092f565b6000808335601e1984360301811261173f578283fd5b83018035915067ffffffffffffffff821115611759578283fd5b602090810192508102360382131561177057600080fd5b9250929050565b6000808335601e1984360301811261178d578283fd5b83018035915067ffffffffffffffff8211156117a7578283fd5b60200191503681900382131561177057600080fd5b600082198211156117cf576117cf611806565b500190565b6000828210156117e6576117e6611806565b500390565b60006000198214156117ff576117ff611806565b5060010190565b634e487b7160e01b600052601160045260246000fdfea26469706673582212203ef361c42eb568553436912278044870daf4a55381bbfe638451800601e9e94164736f6c63430008020033",
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
