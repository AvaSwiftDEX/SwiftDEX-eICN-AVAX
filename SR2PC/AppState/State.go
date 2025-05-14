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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"appValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structState.ShadowLock\",\"name\":\"sdl\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"}],\"name\":\"getShadowLock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structState.ShadowLock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getShadowTranMeta\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structState.ShadowTranMeta\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getShadowValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueId\",\"type\":\"uint256\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"initValues\",\"type\":\"uint256[]\"}],\"name\":\"initValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"notifyUnconfirmedPrepare\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_tranHashesSuccess\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_tranHashesFailed\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"prepare\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"newCM\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"execSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"prepareUnconfirmed\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"newCM\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"execSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumCMPhase\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"payload\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"cmInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cmInputHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"respPayload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"internalType\":\"structCrossMessage\",\"name\":\"cm\",\"type\":\"tuple\"}],\"name\":\"rollback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"retry\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061170f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c9e7263a11610071578063c9e7263a1461018d578063cbb1d90c1461019e578063d70f22ee146101d5578063edcfe2bb1461022c578063f8cad9321461028d578063febc046f146102ae576100a9565b80630ff4c916146100ae5780631f71a643146100d45780639815bef114610133578063a5a5874014610148578063b2a4f5d41461016a575b600080fd5b6100c16100bc3660046113d6565b6102c1565b6040519081526020015b60405180910390f35b6101136100e23660046113d6565b6000602081815291815260409081902081518083019092528054825260018101549282019290925260039091015482565b6040805183518152602093840151938101939093528201526060016100cb565b6101466101413660046112b0565b6102d9565b005b61015b610156366004611358565b61033e565b6040516100cb9392919061154e565b61017d610178366004611358565b610600565b60405190151581526020016100cb565b61014661019b366004611358565b50565b6101b16101ac3660046113ee565b6106fa565b604080518251815260208084015190820152918101511515908201526060016100cb565b61021f6101e33660046113d6565b60408051808201909152600080825260208201525060009081526020818152604091829020825180840190935280548352600101549082015290565b6040516100cb91906115f7565b6100c161023a3660046113ee565b600091825260208281526040808420805460018201548351808601929092528184015260608082019590955282518082039095018552608001825283519383019390932084526002909201905290205490565b6102a061029b36600461140f565b61078c565b6040516100cb929190611520565b61015b6102bc366004611393565b610c1c565b6000818152602081905260409020600301545b919050565b805160015560005b815181101561033a5781818151811061030a57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516000838152918290526040909120600301558061033281611692565b9150506102e1565b5050565b610346611079565b60008060008460a0015160028151811061037057634e487b7160e01b600052603260045260246000fd5b602002602001015180602001905181019061038b9190611340565b600081815260208190526040902054909150158015906103bb575060008181526020819052604090206001015415155b156103ce575060009150600190506105f9565b6001925060008560a001516000815181106103f957634e487b7160e01b600052603260045260246000fd5b60200260200101518060200190518101906104149190611340565b6000838152602081905260409020600301549091508111156104395760009350610460565b6000828152602081905260408120600301805483929061045a90849061167b565b90915550505b60408051600480825260a0820190925290816020015b60608152602001906001900390816104765790505060a08601526040805160208101839052016040516020818303038152906040528560a001516000815181106104d057634e487b7160e01b600052603260045260246000fd5b60200260200101819052508560a0015160018151811061050057634e487b7160e01b600052603260045260246000fd5b60200260200101518560a0015160018151811061052d57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508560a0015160028151811061055d57634e487b7160e01b600052603260045260246000fd5b60200260200101518560a0015160028151811061058a57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508560a001516003815181106105ba57634e487b7160e01b600052603260045260246000fd5b60200260200101518560a001516003815181106105e757634e487b7160e01b600052603260045260246000fd5b60209081029190910101525060009150505b9193909250565b6000808260a0015160028151811061062857634e487b7160e01b600052603260045260246000fd5b60200260200101518060200190518101906106439190611340565b60008181526020819052604090205490915015801590610673575060008181526020819052604090206001015415155b156106825760019150506102d4565b60008360a001516000815181106106a957634e487b7160e01b600052603260045260246000fd5b60200260200101518060200190518101906106c49190611340565b90508060008084815260200190815260200160002060030160008282546106eb9190611663565b90915550600095945050505050565b60408051606081018252600080825260208083018290528284018290528582528190529190912060040180548390811061074457634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805160608101825260039093029091018054835260018101549383019390935260029092015460ff1615159181019190915290505b92915050565b60608060008060005b6001548110156108cd57600081815260208190526040902054881480156107cc575060008181526020819052604090206001015487145b156108bb5760005b6000828152602081905260409020600401548110156108b95786600080848152602001908152602001600020600401828154811061082257634e487b7160e01b600052603260045260246000fd5b90600052602060002090600302016001015414156108a757600082815260208190526040902060040180548290811061086b57634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610899578361089181611692565b9450506108a7565b826108a381611692565b9350505b806108b181611692565b9150506107d4565b505b806108c581611692565b915050610795565b508167ffffffffffffffff8111156108f557634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561091e578160200160208202803683370190505b5093508067ffffffffffffffff81111561094857634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610971578160200160208202803683370190505b509250600091506000905060005b600154811015610c1157600081815260208190526040902054881480156109b6575060008181526020819052604090206001015487145b15610bff576000818152602081815260408083208381556001810184905581518084018d90528083018c905260608082018c90528351808303909101815260809091018352805190840120808552600282018452918420548585529284905260030191909155905b600083815260208190526040902060040154811015610bdf57876000808581526020019081526020016000206004018281548110610a6c57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060030201600101541415610bcd576000838152602081905260409020600401805482908110610ab557634e487b7160e01b600052603260045260246000fd5b600091825260209091206002600390920201015460ff1615610b51576000838152602081905260409020600401805482908110610b0257634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154878680610b2090611692565b975081518110610b4057634e487b7160e01b600052603260045260246000fd5b602002602001018181525050610bcd565b6000838152602081905260409020600401805482908110610b8257634e487b7160e01b600052603260045260246000fd5b906000526020600020906003020160000154868580610ba090611692565b965081518110610bc057634e487b7160e01b600052603260045260246000fd5b6020026020010181815250505b80610bd781611692565b915050610a1e565b506000828152602081905260408120610bfd916004909101906110ea565b505b80610c0981611692565b91505061097f565b505050935093915050565b610c24611079565b60008060008560a00151600181518110610c4e57634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610c699190611340565b905060008660a00151600281518110610c9257634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610cad9190611340565b600081815260208190526040902054909150158015610cdb5750600081815260208190526040902060010154155b15610d465786516000828152602081815260408083209384556101208b0151600185018190558b518251808501919091528083019190915260608082018c90528251808303909101815260809091018252805190830120600385015490845260029094019091529020555b8651600082815260208190526040902054141580610d7a575061012087015160008281526020819052604090206001015414155b15610d8f575060009250600191506110729050565b86516101208801516040805160208101939093528201526060810187905260009060800160405160208183030381529060405280519060200120905060008860a00151600081518110610df257634e487b7160e01b600052603260045260246000fd5b6020026020010151806020019051810190610e0d9190611340565b600084815260208181526040808320868452600201909152902054909150600190821115610e3d57506000610e6d565b60008481526020818152604080832086845260020190915281208054849290610e6790849061167b565b90915550505b60408051600480825260a0820190925290816020015b6060815260200190600190039081610e835790505060a08901526040805160208101849052016040516020818303038152906040528860a00151600081518110610edd57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508960a00151600181518110610f0d57634e487b7160e01b600052603260045260246000fd5b60200260200101518860a00151600181518110610f3a57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508960a00151600281518110610f6a57634e487b7160e01b600052603260045260246000fd5b60200260200101518860a00151600281518110610f9757634e487b7160e01b600052603260045260246000fd5b60200260200101819052508960a00151600381518110610fc757634e487b7160e01b600052603260045260246000fd5b60200260200101518860a00151600381518110610ff457634e487b7160e01b600052603260045260246000fd5b6020908102919091018101919091526000948552848152604080862081516060810183529788528783018c8152841515928901928352600490910180546001808201835591895293882098516003909402909801928355519682019690965594516002909501805460ff191695151595909517909455509193509150505b9250925092565b6040518061014001604052806000815260200160008152602001600060048111156110b457634e487b7160e01b600052602160045260246000fd5b8152600060208201819052604082018190526060808301526080820181905260a0820181905260c0820181905260e09091015290565b508054600082556003029060005260206000209081019061019b91905b8082111561112d576000808255600182015560028101805460ff19169055600301611107565b5090565b6000601f8381840112611142578182fd5b823560206111576111528361163f565b61160e565b82815281810190868301865b858110156111e357813589018a603f82011261117d578889fd5b85810135604067ffffffffffffffff82111561119b5761119b6116c3565b6111ac828b01601f1916890161160e565b8281528d828486010111156111bf578b8cfd5b828285018a83013791820188018b9052508552509284019290840190600101611163565b509098975050505050505050565b8035600581106102d457600080fd5b6000610140808385031215611213578182fd5b61121c8161160e565b9150508135815260208201356020820152611239604083016111f1565b6040820152606082013560608201526080820135608082015260a082013567ffffffffffffffff81111561126c57600080fd5b61127884828501611131565b60a08301525060c082013560c082015260e082013560e082015261010080830135818301525061012080830135818301525092915050565b600060208083850312156112c2578182fd5b823567ffffffffffffffff8111156112d8578283fd5b8301601f810185136112e8578283fd5b80356112f66111528261163f565b8181528381019083850185840285018601891015611312578687fd5b8694505b83851015611334578035835260019490940193918501918501611316565b50979650505050505050565b600060208284031215611351578081fd5b5051919050565b600060208284031215611369578081fd5b813567ffffffffffffffff81111561137f578182fd5b61138b84828501611200565b949350505050565b600080604083850312156113a5578081fd5b823567ffffffffffffffff8111156113bb578182fd5b6113c785828601611200565b95602094909401359450505050565b6000602082840312156113e7578081fd5b5035919050565b60008060408385031215611400578182fd5b50508035926020909101359150565b600080600060608486031215611423578081fd5b505081359360208301359350604090920135919050565b6000815180845260208085019450808401835b838110156114695781518752958201959082019060010161144d565b509495945050505050565b6000815180845260208085018081965082840281019150828601855b858110156114f157828403895281518051808652885b818110156114c15782810188015187820189015287016114a6565b818111156114d157898883890101525b5099860199601f01601f1916949094018501935090840190600101611490565b5091979650505050505050565b6005811061151c57634e487b7160e01b600052602160045260246000fd5b9052565b600060408252611533604083018561143a565b8281036020840152611545818561143a565b95945050505050565b6000606082528451606083015260208501516080830152604085015161157760a08401826114fe565b50606085015160c0830152608085015160e083015260a085015161014061010081818601526115aa6101a0860184611474565b60c08901516101208781019190915260e08a0151938701939093529088015161016086015290870151610180850152851515602085015291506115ea9050565b821515604083015261138b565b815181526020808301519082015260408101610786565b604051601f8201601f1916810167ffffffffffffffff81118282101715611637576116376116c3565b604052919050565b600067ffffffffffffffff821115611659576116596116c3565b5060209081020190565b60008219821115611676576116766116ad565b500190565b60008282101561168d5761168d6116ad565b500390565b60006000198214156116a6576116a66116ad565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122071f544fc2775f6c671809dd0d5dfb55d7a0384008cb77a8aec87b714c66db96d64736f6c63430008020033",
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
// Solidity: function appValue(uint256 ) view returns((uint256,uint256) sdl, uint256 value)
func (_AppState *AppStateCaller) AppValue(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sdl   StateShadowLock
	Value *big.Int
}, error) {
	var out []interface{}
	err := _AppState.contract.Call(opts, &out, "appValue", arg0)

	outstruct := new(struct {
		Sdl   StateShadowLock
		Value *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sdl = *abi.ConvertType(out[0], new(StateShadowLock)).(*StateShadowLock)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AppValue is a free data retrieval call binding the contract method 0x1f71a643.
//
// Solidity: function appValue(uint256 ) view returns((uint256,uint256) sdl, uint256 value)
func (_AppState *AppStateSession) AppValue(arg0 *big.Int) (struct {
	Sdl   StateShadowLock
	Value *big.Int
}, error) {
	return _AppState.Contract.AppValue(&_AppState.CallOpts, arg0)
}

// AppValue is a free data retrieval call binding the contract method 0x1f71a643.
//
// Solidity: function appValue(uint256 ) view returns((uint256,uint256) sdl, uint256 value)
func (_AppState *AppStateCallerSession) AppValue(arg0 *big.Int) (struct {
	Sdl   StateShadowLock
	Value *big.Int
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
