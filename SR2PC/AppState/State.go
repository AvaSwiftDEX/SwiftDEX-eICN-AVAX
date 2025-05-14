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
	Bin: "0x608060405234801561001057600080fd5b506117e9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c9e7263a11610071578063c9e7263a146101a2578063cbb1d90c146101b3578063d70f22ee146101ea578063edcfe2bb14610241578063f8cad932146102a2578063febc046f146102c3576100a9565b80630ff4c916146100ae5780631f71a643146100d45780639815bef114610148578063a5a587401461015d578063b2a4f5d41461017f575b600080fd5b6100c16100bc3660046114b0565b6102d6565b6040519081526020015b60405180910390f35b61011e6100e23660046114b0565b60006020818152918152604090819020815180830190925280548252600181015492820192909252600382015460059092015490919060ff1683565b604080518451815260209485015194810194909452830191909152151560608201526080016100cb565b61015b61015636600461138a565b6102ee565b005b61017061016b366004611432565b610365565b6040516100cb93929190611628565b61019261018d366004611432565b61064d565b60405190151581526020016100cb565b61015b6101b0366004611432565b50565b6101c66101c13660046114c8565b6107ac565b604080518251815260208084015190820152918101511515908201526060016100cb565b6102346101f83660046114b0565b60408051808201909152600080825260208201525060009081526020818152604091829020825180840190935280548352600101549082015290565b6040516100cb91906116d1565b6100c161024f3660046114c8565b600091825260208281526040808420805460018201548351808601929092528184015260608082019590955282518082039095018552608001825283519383019390932084526002909201905290205490565b6102b56102b03660046114e9565b61083e565b6040516100cb9291906115fa565b6101706102d136600461146d565b610cce565b6000818152602081905260409020600301545b919050565b805160015560005b81518110156103615781818151811061031f57634e487b7160e01b600052603260045260246000fd5b60209081029190910181015160008381529182905260409091206003810191909155600501805460ff19166001179055806103598161176c565b9150506102f6565b5050565b61036d611153565b60008060008460a0015160028151811061039757634e487b7160e01b600052603260045260246000fd5b60200260200101518060200190518101906103b2919061141a565b60008181526020819052604090206005015490915060ff166103db575060009150819050610646565b60008181526020819052604090205415801590610408575060008181526020819052604090206001015415155b1561041b57506000915060019050610646565b6001925060008560a0015160008151811061044657634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610461919061141a565b60008381526020819052604090206003015490915081111561048657600093506104ad565b600082815260208190526040812060030180548392906104a7908490611755565b90915550505b60408051600480825260a0820190925290816020015b60608152602001906001900390816104c35790505060a08601526040805160208101839052016040516020818303038152906040528560a0015160008151811061051d57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508560a0015160018151811061054d57634e487b7160e01b600052603260045260246000fd5b60200260200101518560a0015160018151811061057a57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508560a001516002815181106105aa57634e487b7160e01b600052603260045260246000fd5b60200260200101518560a001516002815181106105d757634e487b7160e01b600052603260045260246000fd5b60200260200101819052508560a0015160038151811061060757634e487b7160e01b600052603260045260246000fd5b60200260200101518560a0015160038151811061063457634e487b7160e01b600052603260045260246000fd5b60209081029190910101525060009150505b9193909250565b6000808260a0015160028151811061067557634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610690919061141a565b60008181526020819052604090206005015490915060ff166106f85760405162461bcd60e51b815260206004820152601a60248201527f526f6c6c4261636b57697468496e76616c696456616c75654964000000000000604482015260640160405180910390fd5b60008181526020819052604090205415801590610725575060008181526020819052604090206001015415155b156107345760019150506102e9565b60008360a0015160008151811061075b57634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610776919061141a565b905080600080848152602001908152602001600020600301600082825461079d919061173d565b90915550600095945050505050565b6040805160608101825260008082526020808301829052828401829052858252819052919091206004018054839081106107f657634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805160608101825260039093029091018054835260018101549383019390935260029092015460ff1615159181019190915290505b92915050565b60608060008060005b60015481101561097f576000818152602081905260409020548814801561087e575060008181526020819052604090206001015487145b1561096d5760005b60008281526020819052604090206004015481101561096b578660008084815260200190815260200160002060040182815481106108d457634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160010154141561095957600082815260208190526040902060040180548290811061091d57634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff161561094b57836109438161176c565b945050610959565b826109558161176c565b9350505b806109638161176c565b915050610886565b505b806109778161176c565b915050610847565b508167ffffffffffffffff8111156109a757634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156109d0578160200160208202803683370190505b5093508067ffffffffffffffff8111156109fa57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a23578160200160208202803683370190505b509250600091506000905060005b600154811015610cc35760008181526020819052604090205488148015610a68575060008181526020819052604090206001015487145b15610cb1576000818152602081815260408083208381556001810184905581518084018d90528083018c905260608082018c90528351808303909101815260809091018352805190840120808552600282018452918420548585529284905260030191909155905b600083815260208190526040902060040154811015610c9157876000808581526020019081526020016000206004018281548110610b1e57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610c7f576000838152602081905260409020600401805482908110610b6757634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610c03576000838152602081905260409020600401805482908110610bb457634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154878680610bd29061176c565b975081518110610bf257634e487b7160e01b600052603260045260246000fd5b602002602001018181525050610c7f565b6000838152602081905260409020600401805482908110610c3457634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154868580610c529061176c565b965081518110610c7257634e487b7160e01b600052603260045260246000fd5b6020026020010181815250505b80610c898161176c565b915050610ad0565b506000828152602081905260408120610caf916004909101906111c4565b505b80610cbb8161176c565b915050610a31565b505050935093915050565b610cd6611153565b60008060008560a00151600181518110610d0057634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610d1b919061141a565b905060008660a00151600281518110610d4457634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610d5f919061141a565b60008181526020819052604090206005015490915060ff16610d8a57506000925082915061114c9050565b600081815260208190526040902054158015610db55750600081815260208190526040902060010154155b15610e205786516000828152602081815260408083209384556101208b0151600185018190558b518251808501919091528083019190915260608082018c90528251808303909101815260809091018252805190830120600385015490845260029094019091529020555b8651600082815260208190526040902054141580610e54575061012087015160008281526020819052604090206001015414155b15610e695750600092506001915061114c9050565b86516101208801516040805160208101939093528201526060810187905260009060800160405160208183030381529060405280519060200120905060008860a00151600081518110610ecc57634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610ee7919061141a565b600084815260208181526040808320868452600201909152902054909150600190821115610f1757506000610f47565b60008481526020818152604080832086845260020190915281208054849290610f41908490611755565b90915550505b60408051600480825260a0820190925290816020015b6060815260200190600190039081610f5d5790505060a08901526040805160208101849052016040516020818303038152906040528860a00151600081518110610fb757634e487b7160e01b600052603260045260246000fd5b60200260200101819052508960a00151600181518110610fe757634e487b7160e01b600052603260045260246000fd5b60200260200101518860a0015160018151811061101457634e487b7160e01b600052603260045260246000fd5b60200260200101819052508960a0015160028151811061104457634e487b7160e01b600052603260045260246000fd5b60200260200101518860a0015160028151811061107157634e487b7160e01b600052603260045260246000fd5b60200260200101819052508960a001516003815181106110a157634e487b7160e01b600052603260045260246000fd5b60200260200101518860a001516003815181106110ce57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101919091526000948552848152604080862081516060810183529788528783018c8152841515928901928352600490910180546001808201835591895293882098516003909402909801928355519682019690965594516002909501805460ff191695151595909517909455509193509150505b9250925092565b60405180610140016040528060008152602001600081526020016000600481111561118e57634e487b7160e01b600052602160045260246000fd5b8152600060208201819052604082018190526060808301526080820181905260a0820181905260c0820181905260e09091015290565b50805460008255600302906000526020600020908101906101b091905b80821115611207576000808255600182015560028101805460ff191690556003016111e1565b5090565b6000601f838184011261121c578182fd5b8235602061123161122c83611719565b6116e8565b82815281810190868301865b858110156112bd57813589018a603f820112611257578889fd5b85810135604067ffffffffffffffff8211156112755761127561179d565b611286828b01601f191689016116e8565b8281528d82848601011115611299578b8cfd5b828285018a83013791820188018b905250855250928401929084019060010161123d565b509098975050505050505050565b8035600581106102e957600080fd5b60006101408083850312156112ed578182fd5b6112f6816116e8565b9150508135815260208201356020820152611313604083016112cb565b6040820152606082013560608201526080820135608082015260a082013567ffffffffffffffff81111561134657600080fd5b6113528482850161120b565b60a08301525060c082013560c082015260e082013560e082015261010080830135818301525061012080830135818301525092915050565b6000602080838503121561139c578182fd5b823567ffffffffffffffff8111156113b2578283fd5b8301601f810185136113c2578283fd5b80356113d061122c82611719565b81815283810190838501858402850186018910156113ec578687fd5b8694505b8385101561140e5780358352600194909401939185019185016113f0565b50979650505050505050565b60006020828403121561142b578081fd5b5051919050565b600060208284031215611443578081fd5b813567ffffffffffffffff811115611459578182fd5b611465848285016112da565b949350505050565b6000806040838503121561147f578081fd5b823567ffffffffffffffff811115611495578182fd5b6114a1858286016112da565b95602094909401359450505050565b6000602082840312156114c1578081fd5b5035919050565b600080604083850312156114da578182fd5b50508035926020909101359150565b6000806000606084860312156114fd578081fd5b505081359360208301359350604090920135919050565b6000815180845260208085019450808401835b8381101561154357815187529582019590820190600101611527565b509495945050505050565b6000815180845260208085018081965082840281019150828601855b858110156115cb57828403895281518051808652885b8181101561159b578281018801518782018901528701611580565b818111156115ab57898883890101525b5099860199601f01601f191694909401850193509084019060010161156a565b5091979650505050505050565b600581106115f657634e487b7160e01b600052602160045260246000fd5b9052565b60006040825261160d6040830185611514565b828103602084015261161f8185611514565b95945050505050565b6000606082528451606083015260208501516080830152604085015161165160a08401826115d8565b50606085015160c0830152608085015160e083015260a085015161014061010081818601526116846101a086018461154e565b60c08901516101208781019190915260e08a0151938701939093529088015161016086015290870151610180850152851515602085015291506116c49050565b8215156040830152611465565b815181526020808301519082015260408101610838565b604051601f8201601f1916810167ffffffffffffffff811182821017156117115761171161179d565b604052919050565b600067ffffffffffffffff8211156117335761173361179d565b5060209081020190565b6000821982111561175057611750611787565b500190565b60008282101561176757611767611787565b500390565b600060001982141561178057611780611787565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220caf0f7d138750da839607b4651da89491e2f427c3205e36a3f56c760bacb7a5d64736f6c63430008020033",
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
