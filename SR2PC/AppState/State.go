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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"appValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structState.ShadowLock\",\"name\":\"sdl\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"}],\"name\":\"getShadowLock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structState.ShadowLock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getShadowTranMeta\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structState.ShadowTranMeta\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getShadowValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"initValues\",\"type\":\"uint256[]\"}],\"name\":\"initValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"notifyUnconfirmedPrepare\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_tranHashesSuccess\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_tranHashesFailed\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_lockHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"prepare\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"newCM\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"execSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"lockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"prepareUnconfirmed\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"newCM\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"execSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"lockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"rollback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"lockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061193c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c9e7263a11610071578063c9e7263a146101ba578063cbb1d90c146101cb578063d70f22ee14610202578063edcfe2bb14610259578063f8cad932146102ba578063febc046f146102dc576100a9565b80630ff4c916146100ae5780631f71a643146100e45780639815bef114610158578063a5a587401461016d578063b2a4f5d414610190575b600080fd5b6100d16100bc366004611529565b60009081526020819052604090206003015490565b6040519081526020015b60405180910390f35b61012e6100f2366004611529565b60006020818152918152604090819020815180830190925280548252600181015492820192909252600382015460059092015490919060ff1683565b604080518451815260209485015194810194909452830191909152151560608201526080016100db565b61016b61016636600461146a565b6102ef565b005b61018061017b366004611541565b610366565b6040516100db9493929190611737565b6101a361019e366004611541565b6106f3565b6040805192151583526020830191909152016100db565b61016b6101c8366004611541565b50565b6101de6101d93660046115bf565b610866565b604080518251815260208084015190820152918101511515908201526060016100db565b61024c610210366004611529565b60408051808201909152600080825260208201525060009081526020818152604091829020825180840190935280548352600101549082015290565b6040516100db91906117e6565b6100d16102673660046115bf565b600091825260208281526040808420805460018201548351808601929092528184015260608082019590955282518082039095018552608001825283519383019390932084526002909201905290205490565b6102cd6102c83660046115e0565b6108f8565b6040516100db939291906116f4565b6101806102ea36600461157c565b610e68565b805160015560005b81518110156103625781818151811061032057634e487b7160e01b600052603260045260246000fd5b60209081029190910181015160008381529182905260409091206003810191909155600501805460ff191660011790558061035a816118bf565b9150506102f7565b5050565b61036e61139a565b60008080808061038160a08801886117fd565b60008181106103a057634e487b7160e01b600052603260045260246000fd5b90506020028101906103b2919061184b565b8101906103bf91906115bf565b600082815260208190526040902060050154919350915060ff166103ef5750600093508392508291506106ec9050565b6000828152602081905260409020541580159061041c575060008281526020819052604090206001015415155b156104795760008281526020818152604091829020805460019182015484519384018790529383015260608201929092529094506080016040516020818303038152906040528051906020012092508560009550955050506106ec565b6000828152602081905260409020600301546001955081111561049f57600094506104c6565b600082815260208190526040812060030180548392906104c09084906118a8565b90915550505b6040805160038082526080820190925290816020015b60608152602001906001900390816104dc57505060a087810191909152610505908801886117fd565b600081811061052457634e487b7160e01b600052603260045260246000fd5b9050602002810190610536919061184b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060a08b015180519094509092501515905061059157634e487b7160e01b600052603260045260246000fd5b60209081029190910101526105a960a08801886117fd565b60018181106105c857634e487b7160e01b600052603260045260246000fd5b90506020028101906105da919061184b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08701518051600190811061063457634e487b7160e01b600052603260045260246000fd5b602090810291909101015261064c60a08801886117fd565b600281811061066b57634e487b7160e01b600052603260045260246000fd5b905060200281019061067d919061184b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a0870151805160029081106106d757634e487b7160e01b600052603260045260246000fd5b60209081029190910101525060009250829150505b9193509193565b600080808061070560a08601866117fd565b600081811061072457634e487b7160e01b600052603260045260246000fd5b9050602002810190610736919061184b565b81019061074391906115bf565b600082815260208190526040902060050154919350915060ff166107ad5760405162461bcd60e51b815260206004820152601a60248201527f526f6c6c4261636b57697468496e76616c696456616c75654964000000000000604482015260640160405180910390fd5b600082815260208190526040902054158015906107da575060008281526020819052604090206001015415155b15610831576000828152602081815260409182902080546001909101548351928301869052928201526060810191909152608001604051602081830303815290604052805190602001209250600193505050610861565b60008281526020819052604081206003018054839290610852908490611890565b90915550600094508493505050505b915091565b6040805160608101825260008082526020808301829052828401829052858252819052919091206004018054839081106108b057634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805160608101825260039093029091018054835260018101549383019390935260029092015460ff1615159181019190915290505b92915050565b60608060606000806000805b600154811015610a49576000818152602081905260409020548a14801561093b575060008181526020819052604090206001015489145b15610a375761094b600183611890565b915060005b600082815260208190526040902060040154811015610a355788600080848152602001908152602001600020600401828154811061099e57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610a235760008281526020819052604090206004018054829081106109e757634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610a155784610a0d816118bf565b955050610a23565b83610a1f816118bf565b9450505b80610a2d816118bf565b915050610950565b505b80610a41816118bf565b915050610904565b508267ffffffffffffffff811115610a7157634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a9a578160200160208202803683370190505b5095508167ffffffffffffffff811115610ac457634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610aed578160200160208202803683370190505b5094508067ffffffffffffffff811115610b1757634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610b40578160200160208202803683370190505b50935060009250600091506000905060005b600154811015610e5b576000818152602081905260409020548a148015610b89575060008181526020819052604090206001015489145b15610e49576000818152602081815260408083208381556001019290925581519081018390529081018b9052606081018a905260800160405160208183030381529060405280519060200120858380610be1906118bf565b945081518110610c0157634e487b7160e01b600052603260045260246000fd5b602090810291909101810191909152604080518083018d90528082018c905260608082018c90528251808303909101815260809091018252805190830120600084815280845282812082825260028101855292812054858252938190526003909201929092555b600083815260208190526040902060040154811015610e2957896000808581526020019081526020016000206004018281548110610cb657634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610e17576000838152602081905260409020600401805482908110610cff57634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610d9b576000838152602081905260409020600401805482908110610d4c57634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154898780610d6a906118bf565b985081518110610d8a57634e487b7160e01b600052603260045260246000fd5b602002602001018181525050610e17565b6000838152602081905260409020600401805482908110610dcc57634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154888680610dea906118bf565b975081518110610e0a57634e487b7160e01b600052603260045260246000fd5b6020026020010181815250505b80610e21816118bf565b915050610c68565b506000828152602081905260408120610e479160049091019061140b565b505b80610e53816118bf565b915050610b52565b5050505093509350939050565b610e7061139a565b6000808080610e8260a08801886117fd565b6001818110610ea157634e487b7160e01b600052603260045260246000fd5b9050602002810190610eb3919061184b565b810190610ec09190611529565b9050600080610ed260a08a018a6117fd565b6000818110610ef157634e487b7160e01b600052603260045260246000fd5b9050602002810190610f03919061184b565b810190610f1091906115bf565b600082815260208190526040902060050154919350915060ff16610f41575060009450849350839250611391915050565b600082815260208190526040902054158015610f6c5750600082815260208190526040902060010154155b15610fd3576000828152602081815260408083208c358082556101208e0135600183018190558351808601929092528184015260608082018e9052835180830390910181526080909101835280519084012060038201549085526002909101909252909120555b600082815260208190526040902054893514158061100757506000828152602081905260409020600101546101208a013514155b156110665760008281526020818152604091829020805460019091015483519283018690529282015260608101919091526080016040516020818303038152906040528051906020012093508660006001965096509650505050611391565b604080518a3560208201526101208b0135918101919091526060810189905260009060800160408051601f19818403018152918152815160209283012060008681528084528281208282526002019093529120549091506001908311156110cf575060006110ff565b600084815260208181526040808320858452600201909152812080548592906110f99084906118a8565b90915550505b6040805160038082526080820190925290816020015b606081526020019060019003908161111557505060a08a81019190915261113e908c018c6117fd565b600081811061115d57634e487b7160e01b600052603260045260246000fd5b905060200281019061116f919061184b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060a08e01518051909450909250151590506111ca57634e487b7160e01b600052603260045260246000fd5b60209081029190910101526111e260a08c018c6117fd565b600181811061120157634e487b7160e01b600052603260045260246000fd5b9050602002810190611213919061184b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08a01518051600190811061126d57634e487b7160e01b600052603260045260246000fd5b602090810291909101015261128560a08c018c6117fd565b60028181106112a457634e487b7160e01b600052603260045260246000fd5b90506020028101906112b6919061184b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08a01518051600290811061131057634e487b7160e01b600052603260045260246000fd5b6020908102919091018101919091526000948552848152604080862081516060810183529788528783018d8152841515928901928352600490910180546001808201835591895293882098516003909402909801928355519682019690965594516002909501805460ff191695151595909517909455509194509250829150505b92959194509250565b6040518061014001604052806000815260200160008152602001600060048111156113d557634e487b7160e01b600052602160045260246000fd5b8152600060208201819052604082018190526060808301526080820181905260a0820181905260c0820181905260e09091015290565b50805460008255600302906000526020600020908101906101c891905b8082111561144e576000808255600182015560028101805460ff19169055600301611428565b5090565b60006101408284031215611464578081fd5b50919050565b6000602080838503121561147c578182fd5b823567ffffffffffffffff80821115611493578384fd5b818501915085601f8301126114a6578384fd5b8135818111156114b8576114b86118f0565b838102604051601f19603f830116810181811085821117156114dc576114dc6118f0565b604052828152858101935084860182860187018a10156114fa578788fd5b8795505b8386101561151c5780358552600195909501949386019386016114fe565b5098975050505050505050565b60006020828403121561153a578081fd5b5035919050565b600060208284031215611552578081fd5b813567ffffffffffffffff811115611568578182fd5b61157484828501611452565b949350505050565b6000806040838503121561158e578081fd5b823567ffffffffffffffff8111156115a4578182fd5b6115b085828601611452565b95602094909401359450505050565b600080604083850312156115d1578182fd5b50508035926020909101359150565b6000806000606084860312156115f4578081fd5b505081359360208301359350604090920135919050565b6000815180845260208085019450808401835b8381101561163a5781518752958201959082019060010161161e565b509495945050505050565b60008282518085526020808601955080818302840101818601855b848110156116c557601f1980878503018a5282518051808652895b8181101561169657828101880151878201890152870161167b565b818111156116a6578a8883890101525b509a86019a601f01909116939093018401925090830190600101611660565b5090979650505050505050565b600581106116f057634e487b7160e01b600052602160045260246000fd5b9052565b600060608252611707606083018661160b565b8281036020840152611719818661160b565b9050828103604084015261172d818561160b565b9695505050505050565b60006080825285516080830152602086015160a0830152604086015161176060c08401826116d2565b50606086015160e08301526080860151610100818185015260a0880151915061014061012081818701526117986101c0870185611645565b60c08b01519287019290925260e08a015161016087015291890151610180860152908801516101a0850152861515602085015291506117d49050565b92151560408201526060015292915050565b8151815260208083015190820152604081016108f2565b6000808335601e19843603018112611813578283fd5b83018035915067ffffffffffffffff82111561182d578283fd5b602090810192508102360382131561184457600080fd5b9250929050565b6000808335601e19843603018112611861578283fd5b83018035915067ffffffffffffffff82111561187b578283fd5b60200191503681900382131561184457600080fd5b600082198211156118a3576118a36118da565b500190565b6000828210156118ba576118ba6118da565b500390565b60006000198214156118d3576118d36118da565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212201e65e3d7882597f339315ff24607c07e111b0a148bd65f6065cd195f3e111dc764736f6c63430008020033",
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
// Solidity: function notifyUnconfirmedPrepare(uint256 _chainId, uint256 _height, bytes32 _root) returns(bytes32[] _tranHashesSuccess, bytes32[] _tranHashesFailed, bytes32[] _lockHashes)
func (_AppState *AppStateTransactor) NotifyUnconfirmedPrepare(opts *bind.TransactOpts, _chainId *big.Int, _height *big.Int, _root [32]byte) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "notifyUnconfirmedPrepare", _chainId, _height, _root)
}

// NotifyUnconfirmedPrepare is a paid mutator transaction binding the contract method 0xf8cad932.
//
// Solidity: function notifyUnconfirmedPrepare(uint256 _chainId, uint256 _height, bytes32 _root) returns(bytes32[] _tranHashesSuccess, bytes32[] _tranHashesFailed, bytes32[] _lockHashes)
func (_AppState *AppStateSession) NotifyUnconfirmedPrepare(_chainId *big.Int, _height *big.Int, _root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.NotifyUnconfirmedPrepare(&_AppState.TransactOpts, _chainId, _height, _root)
}

// NotifyUnconfirmedPrepare is a paid mutator transaction binding the contract method 0xf8cad932.
//
// Solidity: function notifyUnconfirmedPrepare(uint256 _chainId, uint256 _height, bytes32 _root) returns(bytes32[] _tranHashesSuccess, bytes32[] _tranHashesFailed, bytes32[] _lockHashes)
func (_AppState *AppStateTransactorSession) NotifyUnconfirmedPrepare(_chainId *big.Int, _height *big.Int, _root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.NotifyUnconfirmedPrepare(&_AppState.TransactOpts, _chainId, _height, _root)
}

// Prepare is a paid mutator transaction binding the contract method 0xa5a58740.
//
// Solidity: function prepare((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry, bytes32 lockHash)
func (_AppState *AppStateTransactor) Prepare(opts *bind.TransactOpts, cm CrossMessage) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "prepare", cm)
}

// Prepare is a paid mutator transaction binding the contract method 0xa5a58740.
//
// Solidity: function prepare((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry, bytes32 lockHash)
func (_AppState *AppStateSession) Prepare(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Prepare(&_AppState.TransactOpts, cm)
}

// Prepare is a paid mutator transaction binding the contract method 0xa5a58740.
//
// Solidity: function prepare((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry, bytes32 lockHash)
func (_AppState *AppStateTransactorSession) Prepare(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Prepare(&_AppState.TransactOpts, cm)
}

// PrepareUnconfirmed is a paid mutator transaction binding the contract method 0xfebc046f.
//
// Solidity: function prepareUnconfirmed((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm, bytes32 root) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry, bytes32 lockHash)
func (_AppState *AppStateTransactor) PrepareUnconfirmed(opts *bind.TransactOpts, cm CrossMessage, root [32]byte) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "prepareUnconfirmed", cm, root)
}

// PrepareUnconfirmed is a paid mutator transaction binding the contract method 0xfebc046f.
//
// Solidity: function prepareUnconfirmed((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm, bytes32 root) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry, bytes32 lockHash)
func (_AppState *AppStateSession) PrepareUnconfirmed(cm CrossMessage, root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.PrepareUnconfirmed(&_AppState.TransactOpts, cm, root)
}

// PrepareUnconfirmed is a paid mutator transaction binding the contract method 0xfebc046f.
//
// Solidity: function prepareUnconfirmed((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm, bytes32 root) returns((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) newCM, bool execSuccess, bool retry, bytes32 lockHash)
func (_AppState *AppStateTransactorSession) PrepareUnconfirmed(cm CrossMessage, root [32]byte) (*types.Transaction, error) {
	return _AppState.Contract.PrepareUnconfirmed(&_AppState.TransactOpts, cm, root)
}

// Rollback is a paid mutator transaction binding the contract method 0xb2a4f5d4.
//
// Solidity: function rollback((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns(bool retry, bytes32 lockHash)
func (_AppState *AppStateTransactor) Rollback(opts *bind.TransactOpts, cm CrossMessage) (*types.Transaction, error) {
	return _AppState.contract.Transact(opts, "rollback", cm)
}

// Rollback is a paid mutator transaction binding the contract method 0xb2a4f5d4.
//
// Solidity: function rollback((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns(bool retry, bytes32 lockHash)
func (_AppState *AppStateSession) Rollback(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Rollback(&_AppState.TransactOpts, cm)
}

// Rollback is a paid mutator transaction binding the contract method 0xb2a4f5d4.
//
// Solidity: function rollback((uint256,uint256,uint8,uint256,uint256,bytes[],bytes32,uint256,bytes32,uint256) cm) returns(bool retry, bytes32 lockHash)
func (_AppState *AppStateTransactorSession) Rollback(cm CrossMessage) (*types.Transaction, error) {
	return _AppState.Contract.Rollback(&_AppState.TransactOpts, cm)
}
