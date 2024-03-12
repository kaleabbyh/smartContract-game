// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// RefundEscrowMetaData contains all meta data concerning the RefundEscrow contract.
var RefundEscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"beneficiary_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RefundsClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RefundsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_whitelistedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"beneficiaryTokenWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiaryWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"depositsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRefunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumRefundEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"tokenDepositsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWithdrawalAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"whitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"whitelistedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawalAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160045534801562000015575f80fd5b506040516200244b3803806200244b83398181016040528101906200003b919062000311565b80805f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000b0575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000a7919062000367565b60405180910390fd5b620000c181620001a960201b60201c565b50505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000134576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200012b9062000406565b60405180910390fd5b81600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f60055f6101000a81548160ff021916908360028111156200019c576200019b62000426565b5b0217905550505062000453565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000299826200026e565b9050919050565b620002ab816200028d565b8114620002b6575f80fd5b50565b5f81519050620002c981620002a0565b92915050565b5f620002db826200026e565b9050919050565b620002ed81620002cf565b8114620002f8575f80fd5b50565b5f815190506200030b81620002e2565b92915050565b5f80604083850312156200032a57620003296200026a565b5b5f6200033985828601620002b9565b92505060206200034c85828601620002fb565b9150509250929050565b6200036181620002cf565b82525050565b5f6020820190506200037c5f83018462000356565b92915050565b5f82825260208201905092915050565b7f526566756e64457363726f773a2062656e6566696369617279206973207468655f8201527f207a65726f206164647265737300000000000000000000000000000000000000602082015250565b5f620003ee602d8362000382565b9150620003fb8262000392565b604082019050919050565b5f6020820190508181035f8301526200041f81620003e0565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b611fea80620004615f395ff3fe60806040526004361061011e575f3560e01c80638c52dc411161009f578063db0a1ce511610063578063db0a1ce514610374578063e3a9db1a146103b0578063f2fde38b146103ec578063f340fa0114610414578063f8e4b8d6146104305761011e565b80638c52dc41146102cc5780638da5cb5b146102e25780639af6549a1461030c578063bdbc144014610322578063c19d93fb1461034a5761011e565b8063467580f9116100e6578063467580f9146101ee57806351cff8d914610216578063685ca1941461023e578063715018a61461027a57806383672600146102905761011e565b8063069fdaae1461012257806338af3eed1461014a57806343c35fc31461017457806343d726d6146101b057806344b1fc75146101c6575b5f80fd5b34801561012d575f80fd5b5061014860048036038101906101439190611655565b61046c565b005b348015610155575f80fd5b5061015e6104c7565b60405161016b91906116b3565b60405180910390f35b34801561017f575f80fd5b5061019a600480360381019061019591906116cc565b6104f0565b6040516101a79190611706565b60405180910390f35b3480156101bb575f80fd5b506101c4610520565b005b3480156101d1575f80fd5b506101ec60048036038101906101e79190611655565b6105ed565b005b3480156101f9575f80fd5b50610214600480360381019061020f9190611752565b610644565b005b348015610221575f80fd5b5061023c600480360381019061023791906117cc565b6106c1565b005b348015610249575f80fd5b50610264600480360381019061025f91906117f7565b610715565b604051610271919061183c565b60405180910390f35b348015610285575f80fd5b5061028e61074c565b005b34801561029b575f80fd5b506102b660048036038101906102b191906116cc565b61075f565b6040516102c39190611706565b60405180910390f35b3480156102d7575f80fd5b506102e0610798565b005b3480156102ed575f80fd5b506102f6610865565b6040516103039190611706565b60405180910390f35b348015610317575f80fd5b5061032061088c565b005b34801561032d575f80fd5b50610348600480360381019061034391906116cc565b61090c565b005b348015610355575f80fd5b5061035e610adb565b60405161036b91906118c8565b60405180910390f35b34801561037f575f80fd5b5061039a600480360381019061039591906118e1565b610af0565b6040516103a7919061192e565b60405180910390f35b3480156103bb575f80fd5b506103d660048036038101906103d191906117f7565b610b46565b6040516103e3919061192e565b60405180910390f35b3480156103f7575f80fd5b50610412600480360381019061040d91906117f7565b610b8c565b005b61042e600480360381019061042991906117f7565b610c10565b005b34801561043b575f80fd5b5061045660048036038101906104519190611655565b610c89565b604051610463919061183c565b60405180910390f35b610474610cc1565b8060025f8481526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b5f600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6002602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610528610cc1565b5f600281111561053b5761053a611855565b5b610543610adb565b600281111561055557610554611855565b5b14610595576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058c906119c7565b60405180910390fd5b600260055f6101000a81548160ff021916908360028111156105ba576105b9611855565b5b02179055507f088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f60405160405180910390a1565b6105f78282610c89565b610636576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062d90611a55565b60405180910390fd5b6106408282610d48565b5050565b5f600281111561065757610656611855565b5b61065f610adb565b600281111561067157610670611855565b5b146106b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a890611ae3565b60405180910390fd5b6106bc838383610ef0565b505050565b6106ca81610715565b610709576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070090611b71565b60405180910390fd5b610712816111d8565b50565b5f6001600281111561072a57610729611855565b5b610732610adb565b600281111561074457610743611855565b5b149050919050565b610754610cc1565b61075d5f6112de565b565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6107a0610cc1565b5f60028111156107b3576107b2611855565b5b6107bb610adb565b60028111156107cd576107cc611855565b5b1461080d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080490611bff565b60405180910390fd5b600160055f6101000a81548160ff0219169083600281111561083257610831611855565b5b02179055507f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8960405160405180910390a1565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60028081111561089f5761089e611855565b5b6108a7610adb565b60028111156108b9576108b8611855565b5b146108f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f090611c8d565b60405180910390fd5b61090a6109046104c7565b4761139f565b565b60028081111561091f5761091e611855565b5b610927610adb565b600281111561093957610938611855565b5b14610979576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097090611c8d565b60405180910390fd5b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016109e39190611706565b602060405180830381865afa1580156109fe573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a229190611cbf565b905060025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb610a786104c7565b836040518363ffffffff1660e01b8152600401610a96929190611d45565b6020604051808303815f875af1158015610ab2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ad69190611d96565b505050565b5f60055f9054906101000a900460ff16905090565b5f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610b94610cc1565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c04575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610bfb9190611706565b60405180910390fd5b610c0d816112de565b50565b5f6002811115610c2357610c22611855565b5b610c2b610adb565b6002811115610c3d57610c3c611855565b5b14610c7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7490611ae3565b60405180910390fd5b610c8681611488565b50565b5f60016002811115610c9e57610c9d611855565b5b610ca6610adb565b6002811115610cb857610cb7611855565b5b14905092915050565b610cc96115bd565b73ffffffffffffffffffffffffffffffffffffffff16610ce7610865565b73ffffffffffffffffffffffffffffffffffffffff1614610d4657610d0a6115bd565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610d3d9190611706565b60405180910390fd5b565b610d50610cc1565b5f60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205490505f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f208190555060025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610e5c929190611dc1565b6020604051808303815f875af1158015610e78573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e9c9190611d96565b508173ffffffffffffffffffffffffffffffffffffffff167f525f755cf37e950a7b04b93cee8178d9d0b97d9ca82540787d26d5802f6046fc82604051610ee3919061192e565b60405180910390a2505050565b600160045414610f35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2c90611e32565b60405180910390fd5b60026004819055508260025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b8152600401610fa79190611706565b602060405180830381865afa158015610fc2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fe69190611cbf565b1015611027576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101e90611ec0565b60405180910390fd5b60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8230866040518463ffffffff1660e01b815260040161109493929190611ede565b6020604051808303815f875af11580156110b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110d49190611d96565b508260035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205461112d9190611f40565b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f20819055508073ffffffffffffffffffffffffffffffffffffffff167f0a3641841709764e7fddb9e250125a825976f5b18478824f0895a64677e90561846040516111c3919061192e565b60405180910390a26001600481905550505050565b6111e0610cc1565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555061128c818373ffffffffffffffffffffffffffffffffffffffff1661139f90919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5826040516112d2919061192e565b60405180910390a25050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b804710156113e457306040517fcd7860590000000000000000000000000000000000000000000000000000000081526004016113db9190611706565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff168260405161140990611fa0565b5f6040518083038185875af1925050503d805f8114611443576040519150601f19603f3d011682016040523d82523d5f602084013e611448565b606091505b5050905080611483576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b6001600454146114cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c490611e32565b60405180910390fd5b60026004819055505f3490508060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546115229190611f40565b60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508173ffffffffffffffffffffffffffffffffffffffff167f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4826040516115a9919061192e565b60405180910390a250600160048190555050565b5f33905090565b5f80fd5b5f819050919050565b6115da816115c8565b81146115e4575f80fd5b50565b5f813590506115f5816115d1565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611624826115fb565b9050919050565b6116348161161a565b811461163e575f80fd5b50565b5f8135905061164f8161162b565b92915050565b5f806040838503121561166b5761166a6115c4565b5b5f611678858286016115e7565b925050602061168985828601611641565b9150509250929050565b5f61169d826115fb565b9050919050565b6116ad81611693565b82525050565b5f6020820190506116c65f8301846116a4565b92915050565b5f602082840312156116e1576116e06115c4565b5b5f6116ee848285016115e7565b91505092915050565b6117008161161a565b82525050565b5f6020820190506117195f8301846116f7565b92915050565b5f819050919050565b6117318161171f565b811461173b575f80fd5b50565b5f8135905061174c81611728565b92915050565b5f805f60608486031215611769576117686115c4565b5b5f6117768682870161173e565b9350506020611787868287016115e7565b925050604061179886828701611641565b9150509250925092565b6117ab81611693565b81146117b5575f80fd5b50565b5f813590506117c6816117a2565b92915050565b5f602082840312156117e1576117e06115c4565b5b5f6117ee848285016117b8565b91505092915050565b5f6020828403121561180c5761180b6115c4565b5b5f61181984828501611641565b91505092915050565b5f8115159050919050565b61183681611822565b82525050565b5f60208201905061184f5f83018461182d565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003811061189357611892611855565b5b50565b5f8190506118a382611882565b919050565b5f6118b282611896565b9050919050565b6118c2816118a8565b82525050565b5f6020820190506118db5f8301846118b9565b92915050565b5f80604083850312156118f7576118f66115c4565b5b5f61190485828601611641565b9250506020611915858286016115e7565b9150509250929050565b6119288161171f565b82525050565b5f6020820190506119415f83018461191f565b92915050565b5f82825260208201905092915050565b7f526566756e64457363726f773a2063616e206f6e6c7920636c6f7365207768695f8201527f6c65206163746976650000000000000000000000000000000000000000000000602082015250565b5f6119b1602983611947565b91506119bc82611957565b604082019050919050565b5f6020820190508181035f8301526119de816119a5565b9050919050565b7f436f6e646974696f6e616c457363726f773a207061796565206973206e6f74205f8201527f616c6c6f77656420746f20776974686472617720746f6b656e00000000000000602082015250565b5f611a3f603983611947565b9150611a4a826119e5565b604082019050919050565b5f6020820190508181035f830152611a6c81611a33565b9050919050565b7f526566756e64457363726f773a2063616e206f6e6c79206465706f73697420775f8201527f68696c6520616374697665000000000000000000000000000000000000000000602082015250565b5f611acd602b83611947565b9150611ad882611a73565b604082019050919050565b5f6020820190508181035f830152611afa81611ac1565b9050919050565b7f436f6e646974696f6e616c457363726f773a207061796565206973206e6f74205f8201527f616c6c6f77656420746f20776974686472617700000000000000000000000000602082015250565b5f611b5b603383611947565b9150611b6682611b01565b604082019050919050565b5f6020820190508181035f830152611b8881611b4f565b9050919050565b7f526566756e64457363726f773a2063616e206f6e6c7920656e61626c652072655f8201527f66756e6473207768696c65206163746976650000000000000000000000000000602082015250565b5f611be9603283611947565b9150611bf482611b8f565b604082019050919050565b5f6020820190508181035f830152611c1681611bdd565b9050919050565b7f526566756e64457363726f773a2062656e65666963696172792063616e206f6e5f8201527f6c79207769746864726177207768696c6520636c6f7365640000000000000000602082015250565b5f611c77603883611947565b9150611c8282611c1d565b604082019050919050565b5f6020820190508181035f830152611ca481611c6b565b9050919050565b5f81519050611cb981611728565b92915050565b5f60208284031215611cd457611cd36115c4565b5b5f611ce184828501611cab565b91505092915050565b5f819050919050565b5f611d0d611d08611d03846115fb565b611cea565b6115fb565b9050919050565b5f611d1e82611cf3565b9050919050565b5f611d2f82611d14565b9050919050565b611d3f81611d25565b82525050565b5f604082019050611d585f830185611d36565b611d65602083018461191f565b9392505050565b611d7581611822565b8114611d7f575f80fd5b50565b5f81519050611d9081611d6c565b92915050565b5f60208284031215611dab57611daa6115c4565b5b5f611db884828501611d82565b91505092915050565b5f604082019050611dd45f8301856116f7565b611de1602083018461191f565b9392505050565b7f5245454e5452414e4359000000000000000000000000000000000000000000005f82015250565b5f611e1c600a83611947565b9150611e2782611de8565b602082019050919050565b5f6020820190508181035f830152611e4981611e10565b9050919050565b7f457363726f773a2063616e206f6e6c79206465706f73697420617661696c61625f8201527f6c6520616d6f756e740000000000000000000000000000000000000000000000602082015250565b5f611eaa602983611947565b9150611eb582611e50565b604082019050919050565b5f6020820190508181035f830152611ed781611e9e565b9050919050565b5f606082019050611ef15f8301866116f7565b611efe60208301856116f7565b611f0b604083018461191f565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611f4a8261171f565b9150611f558361171f565b9250828201905080821115611f6d57611f6c611f13565b5b92915050565b5f81905092915050565b50565b5f611f8b5f83611f73565b9150611f9682611f7d565b5f82019050919050565b5f611faa82611f80565b915081905091905056fea2646970667358221220dbe5217a26bdf3de570e1d2f7b8c3dfa83814c83c8e400e7b915c17b725dbeff64736f6c63430008160033",
}

// RefundEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use RefundEscrowMetaData.ABI instead.
var RefundEscrowABI = RefundEscrowMetaData.ABI

// RefundEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RefundEscrowMetaData.Bin instead.
var RefundEscrowBin = RefundEscrowMetaData.Bin

// DeployRefundEscrow deploys a new Ethereum contract, binding an instance of RefundEscrow to it.
func DeployRefundEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, beneficiary_ common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *RefundEscrow, error) {
	parsed, err := RefundEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RefundEscrowBin), backend, beneficiary_, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RefundEscrow{RefundEscrowCaller: RefundEscrowCaller{contract: contract}, RefundEscrowTransactor: RefundEscrowTransactor{contract: contract}, RefundEscrowFilterer: RefundEscrowFilterer{contract: contract}}, nil
}

// RefundEscrow is an auto generated Go binding around an Ethereum contract.
type RefundEscrow struct {
	RefundEscrowCaller     // Read-only binding to the contract
	RefundEscrowTransactor // Write-only binding to the contract
	RefundEscrowFilterer   // Log filterer for contract events
}

// RefundEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefundEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundEscrowSession struct {
	Contract     *RefundEscrow     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefundEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundEscrowCallerSession struct {
	Contract *RefundEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RefundEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundEscrowTransactorSession struct {
	Contract     *RefundEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RefundEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundEscrowRaw struct {
	Contract *RefundEscrow // Generic contract binding to access the raw methods on
}

// RefundEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundEscrowCallerRaw struct {
	Contract *RefundEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// RefundEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundEscrowTransactorRaw struct {
	Contract *RefundEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundEscrow creates a new instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrow(address common.Address, backend bind.ContractBackend) (*RefundEscrow, error) {
	contract, err := bindRefundEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundEscrow{RefundEscrowCaller: RefundEscrowCaller{contract: contract}, RefundEscrowTransactor: RefundEscrowTransactor{contract: contract}, RefundEscrowFilterer: RefundEscrowFilterer{contract: contract}}, nil
}

// NewRefundEscrowCaller creates a new read-only instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrowCaller(address common.Address, caller bind.ContractCaller) (*RefundEscrowCaller, error) {
	contract, err := bindRefundEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowCaller{contract: contract}, nil
}

// NewRefundEscrowTransactor creates a new write-only instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundEscrowTransactor, error) {
	contract, err := bindRefundEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowTransactor{contract: contract}, nil
}

// NewRefundEscrowFilterer creates a new log filterer instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*RefundEscrowFilterer, error) {
	contract, err := bindRefundEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowFilterer{contract: contract}, nil
}

// bindRefundEscrow binds a generic wrapper to an already deployed contract.
func bindRefundEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RefundEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundEscrow *RefundEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundEscrow.Contract.RefundEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundEscrow *RefundEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.Contract.RefundEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundEscrow *RefundEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundEscrow.Contract.RefundEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundEscrow *RefundEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundEscrow *RefundEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundEscrow *RefundEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundEscrow.Contract.contract.Transact(opts, method, params...)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0x43c35fc3.
//
// Solidity: function _whitelistedTokens(bytes32 ) view returns(address)
func (_RefundEscrow *RefundEscrowCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "_whitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0x43c35fc3.
//
// Solidity: function _whitelistedTokens(bytes32 ) view returns(address)
func (_RefundEscrow *RefundEscrowSession) WhitelistedTokens(arg0 [32]byte) (common.Address, error) {
	return _RefundEscrow.Contract.WhitelistedTokens(&_RefundEscrow.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0x43c35fc3.
//
// Solidity: function _whitelistedTokens(bytes32 ) view returns(address)
func (_RefundEscrow *RefundEscrowCallerSession) WhitelistedTokens(arg0 [32]byte) (common.Address, error) {
	return _RefundEscrow.Contract.WhitelistedTokens(&_RefundEscrow.CallOpts, arg0)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RefundEscrow *RefundEscrowCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RefundEscrow *RefundEscrowSession) Beneficiary() (common.Address, error) {
	return _RefundEscrow.Contract.Beneficiary(&_RefundEscrow.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RefundEscrow *RefundEscrowCallerSession) Beneficiary() (common.Address, error) {
	return _RefundEscrow.Contract.Beneficiary(&_RefundEscrow.CallOpts)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_RefundEscrow *RefundEscrowCaller) DepositsOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "depositsOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_RefundEscrow *RefundEscrowSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _RefundEscrow.Contract.DepositsOf(&_RefundEscrow.CallOpts, payee)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_RefundEscrow *RefundEscrowCallerSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _RefundEscrow.Contract.DepositsOf(&_RefundEscrow.CallOpts, payee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefundEscrow *RefundEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefundEscrow *RefundEscrowSession) Owner() (common.Address, error) {
	return _RefundEscrow.Contract.Owner(&_RefundEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefundEscrow *RefundEscrowCallerSession) Owner() (common.Address, error) {
	return _RefundEscrow.Contract.Owner(&_RefundEscrow.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_RefundEscrow *RefundEscrowCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_RefundEscrow *RefundEscrowSession) State() (uint8, error) {
	return _RefundEscrow.Contract.State(&_RefundEscrow.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_RefundEscrow *RefundEscrowCallerSession) State() (uint8, error) {
	return _RefundEscrow.Contract.State(&_RefundEscrow.CallOpts)
}

// TokenDepositsOf is a free data retrieval call binding the contract method 0xdb0a1ce5.
//
// Solidity: function tokenDepositsOf(address payee, bytes32 symbol) view returns(uint256)
func (_RefundEscrow *RefundEscrowCaller) TokenDepositsOf(opts *bind.CallOpts, payee common.Address, symbol [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "tokenDepositsOf", payee, symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDepositsOf is a free data retrieval call binding the contract method 0xdb0a1ce5.
//
// Solidity: function tokenDepositsOf(address payee, bytes32 symbol) view returns(uint256)
func (_RefundEscrow *RefundEscrowSession) TokenDepositsOf(payee common.Address, symbol [32]byte) (*big.Int, error) {
	return _RefundEscrow.Contract.TokenDepositsOf(&_RefundEscrow.CallOpts, payee, symbol)
}

// TokenDepositsOf is a free data retrieval call binding the contract method 0xdb0a1ce5.
//
// Solidity: function tokenDepositsOf(address payee, bytes32 symbol) view returns(uint256)
func (_RefundEscrow *RefundEscrowCallerSession) TokenDepositsOf(payee common.Address, symbol [32]byte) (*big.Int, error) {
	return _RefundEscrow.Contract.TokenDepositsOf(&_RefundEscrow.CallOpts, payee, symbol)
}

// TokenWithdrawalAllowed is a free data retrieval call binding the contract method 0xf8e4b8d6.
//
// Solidity: function tokenWithdrawalAllowed(bytes32 , address ) view returns(bool)
func (_RefundEscrow *RefundEscrowCaller) TokenWithdrawalAllowed(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "tokenWithdrawalAllowed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenWithdrawalAllowed is a free data retrieval call binding the contract method 0xf8e4b8d6.
//
// Solidity: function tokenWithdrawalAllowed(bytes32 , address ) view returns(bool)
func (_RefundEscrow *RefundEscrowSession) TokenWithdrawalAllowed(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _RefundEscrow.Contract.TokenWithdrawalAllowed(&_RefundEscrow.CallOpts, arg0, arg1)
}

// TokenWithdrawalAllowed is a free data retrieval call binding the contract method 0xf8e4b8d6.
//
// Solidity: function tokenWithdrawalAllowed(bytes32 , address ) view returns(bool)
func (_RefundEscrow *RefundEscrowCallerSession) TokenWithdrawalAllowed(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _RefundEscrow.Contract.TokenWithdrawalAllowed(&_RefundEscrow.CallOpts, arg0, arg1)
}

// WhitelistedToken is a free data retrieval call binding the contract method 0x83672600.
//
// Solidity: function whitelistedToken(bytes32 symbol) view returns(address)
func (_RefundEscrow *RefundEscrowCaller) WhitelistedToken(opts *bind.CallOpts, symbol [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "whitelistedToken", symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedToken is a free data retrieval call binding the contract method 0x83672600.
//
// Solidity: function whitelistedToken(bytes32 symbol) view returns(address)
func (_RefundEscrow *RefundEscrowSession) WhitelistedToken(symbol [32]byte) (common.Address, error) {
	return _RefundEscrow.Contract.WhitelistedToken(&_RefundEscrow.CallOpts, symbol)
}

// WhitelistedToken is a free data retrieval call binding the contract method 0x83672600.
//
// Solidity: function whitelistedToken(bytes32 symbol) view returns(address)
func (_RefundEscrow *RefundEscrowCallerSession) WhitelistedToken(symbol [32]byte) (common.Address, error) {
	return _RefundEscrow.Contract.WhitelistedToken(&_RefundEscrow.CallOpts, symbol)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address ) view returns(bool)
func (_RefundEscrow *RefundEscrowCaller) WithdrawalAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "withdrawalAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address ) view returns(bool)
func (_RefundEscrow *RefundEscrowSession) WithdrawalAllowed(arg0 common.Address) (bool, error) {
	return _RefundEscrow.Contract.WithdrawalAllowed(&_RefundEscrow.CallOpts, arg0)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address ) view returns(bool)
func (_RefundEscrow *RefundEscrowCallerSession) WithdrawalAllowed(arg0 common.Address) (bool, error) {
	return _RefundEscrow.Contract.WithdrawalAllowed(&_RefundEscrow.CallOpts, arg0)
}

// BeneficiaryTokenWithdraw is a paid mutator transaction binding the contract method 0xbdbc1440.
//
// Solidity: function beneficiaryTokenWithdraw(bytes32 symbol) returns()
func (_RefundEscrow *RefundEscrowTransactor) BeneficiaryTokenWithdraw(opts *bind.TransactOpts, symbol [32]byte) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "beneficiaryTokenWithdraw", symbol)
}

// BeneficiaryTokenWithdraw is a paid mutator transaction binding the contract method 0xbdbc1440.
//
// Solidity: function beneficiaryTokenWithdraw(bytes32 symbol) returns()
func (_RefundEscrow *RefundEscrowSession) BeneficiaryTokenWithdraw(symbol [32]byte) (*types.Transaction, error) {
	return _RefundEscrow.Contract.BeneficiaryTokenWithdraw(&_RefundEscrow.TransactOpts, symbol)
}

// BeneficiaryTokenWithdraw is a paid mutator transaction binding the contract method 0xbdbc1440.
//
// Solidity: function beneficiaryTokenWithdraw(bytes32 symbol) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) BeneficiaryTokenWithdraw(symbol [32]byte) (*types.Transaction, error) {
	return _RefundEscrow.Contract.BeneficiaryTokenWithdraw(&_RefundEscrow.TransactOpts, symbol)
}

// BeneficiaryWithdraw is a paid mutator transaction binding the contract method 0x9af6549a.
//
// Solidity: function beneficiaryWithdraw() returns()
func (_RefundEscrow *RefundEscrowTransactor) BeneficiaryWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "beneficiaryWithdraw")
}

// BeneficiaryWithdraw is a paid mutator transaction binding the contract method 0x9af6549a.
//
// Solidity: function beneficiaryWithdraw() returns()
func (_RefundEscrow *RefundEscrowSession) BeneficiaryWithdraw() (*types.Transaction, error) {
	return _RefundEscrow.Contract.BeneficiaryWithdraw(&_RefundEscrow.TransactOpts)
}

// BeneficiaryWithdraw is a paid mutator transaction binding the contract method 0x9af6549a.
//
// Solidity: function beneficiaryWithdraw() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) BeneficiaryWithdraw() (*types.Transaction, error) {
	return _RefundEscrow.Contract.BeneficiaryWithdraw(&_RefundEscrow.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundEscrow *RefundEscrowTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundEscrow *RefundEscrowSession) Close() (*types.Transaction, error) {
	return _RefundEscrow.Contract.Close(&_RefundEscrow.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) Close() (*types.Transaction, error) {
	return _RefundEscrow.Contract.Close(&_RefundEscrow.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address refundee) payable returns()
func (_RefundEscrow *RefundEscrowTransactor) Deposit(opts *bind.TransactOpts, refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "deposit", refundee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address refundee) payable returns()
func (_RefundEscrow *RefundEscrowSession) Deposit(refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Deposit(&_RefundEscrow.TransactOpts, refundee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address refundee) payable returns()
func (_RefundEscrow *RefundEscrowTransactorSession) Deposit(refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Deposit(&_RefundEscrow.TransactOpts, refundee)
}

// DepositToken is a paid mutator transaction binding the contract method 0x467580f9.
//
// Solidity: function depositToken(uint256 amount, bytes32 symbol, address refundee) returns()
func (_RefundEscrow *RefundEscrowTransactor) DepositToken(opts *bind.TransactOpts, amount *big.Int, symbol [32]byte, refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "depositToken", amount, symbol, refundee)
}

// DepositToken is a paid mutator transaction binding the contract method 0x467580f9.
//
// Solidity: function depositToken(uint256 amount, bytes32 symbol, address refundee) returns()
func (_RefundEscrow *RefundEscrowSession) DepositToken(amount *big.Int, symbol [32]byte, refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.DepositToken(&_RefundEscrow.TransactOpts, amount, symbol, refundee)
}

// DepositToken is a paid mutator transaction binding the contract method 0x467580f9.
//
// Solidity: function depositToken(uint256 amount, bytes32 symbol, address refundee) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) DepositToken(amount *big.Int, symbol [32]byte, refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.DepositToken(&_RefundEscrow.TransactOpts, amount, symbol, refundee)
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundEscrow *RefundEscrowTransactor) EnableRefunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "enableRefunds")
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundEscrow *RefundEscrowSession) EnableRefunds() (*types.Transaction, error) {
	return _RefundEscrow.Contract.EnableRefunds(&_RefundEscrow.TransactOpts)
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) EnableRefunds() (*types.Transaction, error) {
	return _RefundEscrow.Contract.EnableRefunds(&_RefundEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefundEscrow *RefundEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefundEscrow *RefundEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _RefundEscrow.Contract.RenounceOwnership(&_RefundEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RefundEscrow.Contract.RenounceOwnership(&_RefundEscrow.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefundEscrow *RefundEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefundEscrow *RefundEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.TransferOwnership(&_RefundEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.TransferOwnership(&_RefundEscrow.TransactOpts, newOwner)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x069fdaae.
//
// Solidity: function whitelistToken(bytes32 symbol, address tokenAddress) returns()
func (_RefundEscrow *RefundEscrowTransactor) WhitelistToken(opts *bind.TransactOpts, symbol [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "whitelistToken", symbol, tokenAddress)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x069fdaae.
//
// Solidity: function whitelistToken(bytes32 symbol, address tokenAddress) returns()
func (_RefundEscrow *RefundEscrowSession) WhitelistToken(symbol [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.WhitelistToken(&_RefundEscrow.TransactOpts, symbol, tokenAddress)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x069fdaae.
//
// Solidity: function whitelistToken(bytes32 symbol, address tokenAddress) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) WhitelistToken(symbol [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.WhitelistToken(&_RefundEscrow.TransactOpts, symbol, tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_RefundEscrow *RefundEscrowTransactor) Withdraw(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "withdraw", payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_RefundEscrow *RefundEscrowSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Withdraw(&_RefundEscrow.TransactOpts, payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Withdraw(&_RefundEscrow.TransactOpts, payee)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x44b1fc75.
//
// Solidity: function withdrawToken(bytes32 symbol, address payee) returns()
func (_RefundEscrow *RefundEscrowTransactor) WithdrawToken(opts *bind.TransactOpts, symbol [32]byte, payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "withdrawToken", symbol, payee)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x44b1fc75.
//
// Solidity: function withdrawToken(bytes32 symbol, address payee) returns()
func (_RefundEscrow *RefundEscrowSession) WithdrawToken(symbol [32]byte, payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.WithdrawToken(&_RefundEscrow.TransactOpts, symbol, payee)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x44b1fc75.
//
// Solidity: function withdrawToken(bytes32 symbol, address payee) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) WithdrawToken(symbol [32]byte, payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.WithdrawToken(&_RefundEscrow.TransactOpts, symbol, payee)
}

// RefundEscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the RefundEscrow contract.
type RefundEscrowDepositedIterator struct {
	Event *RefundEscrowDeposited // Event containing the contract specifics and raw log

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
func (it *RefundEscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowDeposited)
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
		it.Event = new(RefundEscrowDeposited)
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
func (it *RefundEscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowDeposited represents a Deposited event raised by the RefundEscrow contract.
type RefundEscrowDeposited struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) FilterDeposited(opts *bind.FilterOpts, payee []common.Address) (*RefundEscrowDepositedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowDepositedIterator{contract: _RefundEscrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RefundEscrowDeposited, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowDeposited)
				if err := _RefundEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) ParseDeposited(log types.Log) (*RefundEscrowDeposited, error) {
	event := new(RefundEscrowDeposited)
	if err := _RefundEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowDepositedTokenIterator is returned from FilterDepositedToken and is used to iterate over the raw logs and unpacked data for DepositedToken events raised by the RefundEscrow contract.
type RefundEscrowDepositedTokenIterator struct {
	Event *RefundEscrowDepositedToken // Event containing the contract specifics and raw log

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
func (it *RefundEscrowDepositedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowDepositedToken)
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
		it.Event = new(RefundEscrowDepositedToken)
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
func (it *RefundEscrowDepositedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowDepositedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowDepositedToken represents a DepositedToken event raised by the RefundEscrow contract.
type RefundEscrowDepositedToken struct {
	Payee       common.Address
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositedToken is a free log retrieval operation binding the contract event 0x0a3641841709764e7fddb9e250125a825976f5b18478824f0895a64677e90561.
//
// Solidity: event DepositedToken(address indexed payee, uint256 tokenAmount)
func (_RefundEscrow *RefundEscrowFilterer) FilterDepositedToken(opts *bind.FilterOpts, payee []common.Address) (*RefundEscrowDepositedTokenIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "DepositedToken", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowDepositedTokenIterator{contract: _RefundEscrow.contract, event: "DepositedToken", logs: logs, sub: sub}, nil
}

// WatchDepositedToken is a free log subscription operation binding the contract event 0x0a3641841709764e7fddb9e250125a825976f5b18478824f0895a64677e90561.
//
// Solidity: event DepositedToken(address indexed payee, uint256 tokenAmount)
func (_RefundEscrow *RefundEscrowFilterer) WatchDepositedToken(opts *bind.WatchOpts, sink chan<- *RefundEscrowDepositedToken, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "DepositedToken", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowDepositedToken)
				if err := _RefundEscrow.contract.UnpackLog(event, "DepositedToken", log); err != nil {
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

// ParseDepositedToken is a log parse operation binding the contract event 0x0a3641841709764e7fddb9e250125a825976f5b18478824f0895a64677e90561.
//
// Solidity: event DepositedToken(address indexed payee, uint256 tokenAmount)
func (_RefundEscrow *RefundEscrowFilterer) ParseDepositedToken(log types.Log) (*RefundEscrowDepositedToken, error) {
	event := new(RefundEscrowDepositedToken)
	if err := _RefundEscrow.contract.UnpackLog(event, "DepositedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RefundEscrow contract.
type RefundEscrowOwnershipTransferredIterator struct {
	Event *RefundEscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RefundEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowOwnershipTransferred)
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
		it.Event = new(RefundEscrowOwnershipTransferred)
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
func (it *RefundEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the RefundEscrow contract.
type RefundEscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefundEscrow *RefundEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RefundEscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowOwnershipTransferredIterator{contract: _RefundEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefundEscrow *RefundEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RefundEscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowOwnershipTransferred)
				if err := _RefundEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RefundEscrow *RefundEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*RefundEscrowOwnershipTransferred, error) {
	event := new(RefundEscrowOwnershipTransferred)
	if err := _RefundEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowRefundsClosedIterator is returned from FilterRefundsClosed and is used to iterate over the raw logs and unpacked data for RefundsClosed events raised by the RefundEscrow contract.
type RefundEscrowRefundsClosedIterator struct {
	Event *RefundEscrowRefundsClosed // Event containing the contract specifics and raw log

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
func (it *RefundEscrowRefundsClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowRefundsClosed)
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
		it.Event = new(RefundEscrowRefundsClosed)
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
func (it *RefundEscrowRefundsClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowRefundsClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowRefundsClosed represents a RefundsClosed event raised by the RefundEscrow contract.
type RefundEscrowRefundsClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefundsClosed is a free log retrieval operation binding the contract event 0x088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f.
//
// Solidity: event RefundsClosed()
func (_RefundEscrow *RefundEscrowFilterer) FilterRefundsClosed(opts *bind.FilterOpts) (*RefundEscrowRefundsClosedIterator, error) {

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "RefundsClosed")
	if err != nil {
		return nil, err
	}
	return &RefundEscrowRefundsClosedIterator{contract: _RefundEscrow.contract, event: "RefundsClosed", logs: logs, sub: sub}, nil
}

// WatchRefundsClosed is a free log subscription operation binding the contract event 0x088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f.
//
// Solidity: event RefundsClosed()
func (_RefundEscrow *RefundEscrowFilterer) WatchRefundsClosed(opts *bind.WatchOpts, sink chan<- *RefundEscrowRefundsClosed) (event.Subscription, error) {

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "RefundsClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowRefundsClosed)
				if err := _RefundEscrow.contract.UnpackLog(event, "RefundsClosed", log); err != nil {
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

// ParseRefundsClosed is a log parse operation binding the contract event 0x088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f.
//
// Solidity: event RefundsClosed()
func (_RefundEscrow *RefundEscrowFilterer) ParseRefundsClosed(log types.Log) (*RefundEscrowRefundsClosed, error) {
	event := new(RefundEscrowRefundsClosed)
	if err := _RefundEscrow.contract.UnpackLog(event, "RefundsClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowRefundsEnabledIterator is returned from FilterRefundsEnabled and is used to iterate over the raw logs and unpacked data for RefundsEnabled events raised by the RefundEscrow contract.
type RefundEscrowRefundsEnabledIterator struct {
	Event *RefundEscrowRefundsEnabled // Event containing the contract specifics and raw log

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
func (it *RefundEscrowRefundsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowRefundsEnabled)
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
		it.Event = new(RefundEscrowRefundsEnabled)
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
func (it *RefundEscrowRefundsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowRefundsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowRefundsEnabled represents a RefundsEnabled event raised by the RefundEscrow contract.
type RefundEscrowRefundsEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefundsEnabled is a free log retrieval operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundEscrow *RefundEscrowFilterer) FilterRefundsEnabled(opts *bind.FilterOpts) (*RefundEscrowRefundsEnabledIterator, error) {

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "RefundsEnabled")
	if err != nil {
		return nil, err
	}
	return &RefundEscrowRefundsEnabledIterator{contract: _RefundEscrow.contract, event: "RefundsEnabled", logs: logs, sub: sub}, nil
}

// WatchRefundsEnabled is a free log subscription operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundEscrow *RefundEscrowFilterer) WatchRefundsEnabled(opts *bind.WatchOpts, sink chan<- *RefundEscrowRefundsEnabled) (event.Subscription, error) {

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "RefundsEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowRefundsEnabled)
				if err := _RefundEscrow.contract.UnpackLog(event, "RefundsEnabled", log); err != nil {
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

// ParseRefundsEnabled is a log parse operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundEscrow *RefundEscrowFilterer) ParseRefundsEnabled(log types.Log) (*RefundEscrowRefundsEnabled, error) {
	event := new(RefundEscrowRefundsEnabled)
	if err := _RefundEscrow.contract.UnpackLog(event, "RefundsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RefundEscrow contract.
type RefundEscrowWithdrawnIterator struct {
	Event *RefundEscrowWithdrawn // Event containing the contract specifics and raw log

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
func (it *RefundEscrowWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowWithdrawn)
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
		it.Event = new(RefundEscrowWithdrawn)
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
func (it *RefundEscrowWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowWithdrawn represents a Withdrawn event raised by the RefundEscrow contract.
type RefundEscrowWithdrawn struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) FilterWithdrawn(opts *bind.FilterOpts, payee []common.Address) (*RefundEscrowWithdrawnIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowWithdrawnIterator{contract: _RefundEscrow.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RefundEscrowWithdrawn, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowWithdrawn)
				if err := _RefundEscrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) ParseWithdrawn(log types.Log) (*RefundEscrowWithdrawn, error) {
	event := new(RefundEscrowWithdrawn)
	if err := _RefundEscrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowWithdrawnTokenIterator is returned from FilterWithdrawnToken and is used to iterate over the raw logs and unpacked data for WithdrawnToken events raised by the RefundEscrow contract.
type RefundEscrowWithdrawnTokenIterator struct {
	Event *RefundEscrowWithdrawnToken // Event containing the contract specifics and raw log

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
func (it *RefundEscrowWithdrawnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowWithdrawnToken)
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
		it.Event = new(RefundEscrowWithdrawnToken)
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
func (it *RefundEscrowWithdrawnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowWithdrawnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowWithdrawnToken represents a WithdrawnToken event raised by the RefundEscrow contract.
type RefundEscrowWithdrawnToken struct {
	Payee       common.Address
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnToken is a free log retrieval operation binding the contract event 0x525f755cf37e950a7b04b93cee8178d9d0b97d9ca82540787d26d5802f6046fc.
//
// Solidity: event WithdrawnToken(address indexed payee, uint256 tokenAmount)
func (_RefundEscrow *RefundEscrowFilterer) FilterWithdrawnToken(opts *bind.FilterOpts, payee []common.Address) (*RefundEscrowWithdrawnTokenIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "WithdrawnToken", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowWithdrawnTokenIterator{contract: _RefundEscrow.contract, event: "WithdrawnToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawnToken is a free log subscription operation binding the contract event 0x525f755cf37e950a7b04b93cee8178d9d0b97d9ca82540787d26d5802f6046fc.
//
// Solidity: event WithdrawnToken(address indexed payee, uint256 tokenAmount)
func (_RefundEscrow *RefundEscrowFilterer) WatchWithdrawnToken(opts *bind.WatchOpts, sink chan<- *RefundEscrowWithdrawnToken, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "WithdrawnToken", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowWithdrawnToken)
				if err := _RefundEscrow.contract.UnpackLog(event, "WithdrawnToken", log); err != nil {
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

// ParseWithdrawnToken is a log parse operation binding the contract event 0x525f755cf37e950a7b04b93cee8178d9d0b97d9ca82540787d26d5802f6046fc.
//
// Solidity: event WithdrawnToken(address indexed payee, uint256 tokenAmount)
func (_RefundEscrow *RefundEscrowFilterer) ParseWithdrawnToken(log types.Log) (*RefundEscrowWithdrawnToken, error) {
	event := new(RefundEscrowWithdrawnToken)
	if err := _RefundEscrow.contract.UnpackLog(event, "WithdrawnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}