// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MortalABI is the input ABI used to generate the binding from.
const MortalABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// MortalBin is the compiled bytecode used for deploying new contracts.
const MortalBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b60bb8061003b6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b58114603c575b600080fd5b3415604657600080fd5b604c604e565b005b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415608c5760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5b5600a165627a7a72305820fa5e06a1dfec01fb1811be57e2277466de1de6ca327174fffc91243471605a8b0029`

// DeployMortal deploys a new Ethereum contract, binding an instance of Mortal to it.
func DeployMortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mortal, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}}, nil
}

// Mortal is an auto generated Go binding around an Ethereum contract.
type Mortal struct {
	MortalCaller     // Read-only binding to the contract
	MortalTransactor // Write-only binding to the contract
}

// MortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type MortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MortalSession struct {
	Contract     *Mortal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MortalCallerSession struct {
	Contract *MortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MortalTransactorSession struct {
	Contract     *MortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type MortalRaw struct {
	Contract *Mortal // Generic contract binding to access the raw methods on
}

// MortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MortalCallerRaw struct {
	Contract *MortalCaller // Generic read-only contract binding to access the raw methods on
}

// MortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MortalTransactorRaw struct {
	Contract *MortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMortal creates a new instance of Mortal, bound to a specific deployed contract.
func NewMortal(address common.Address, backend bind.ContractBackend) (*Mortal, error) {
	contract, err := bindMortal(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}}, nil
}

// NewMortalCaller creates a new read-only instance of Mortal, bound to a specific deployed contract.
func NewMortalCaller(address common.Address, caller bind.ContractCaller) (*MortalCaller, error) {
	contract, err := bindMortal(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MortalCaller{contract: contract}, nil
}

// NewMortalTransactor creates a new write-only instance of Mortal, bound to a specific deployed contract.
func NewMortalTransactor(address common.Address, transactor bind.ContractTransactor) (*MortalTransactor, error) {
	contract, err := bindMortal(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MortalTransactor{contract: contract}, nil
}

// bindMortal binds a generic wrapper to an already deployed contract.
func bindMortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.MortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transact(opts, method, params...)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactorSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// PopcontractABI is the input ABI used to generate the binding from.
const PopcontractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_nameOfParty\",\"type\":\"string\"},{\"name\":\"place\",\"type\":\"string\"},{\"name\":\"organizers\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"address[]\"},{\"name\":\"durationInMinutes\",\"type\":\"uint256\"}],\"name\":\"setConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalKeySet\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signWholeConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrganizersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locationOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isOrganizer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignedConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"configSignOrganizers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"publicKeyConsensus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signedConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"keySetAdded\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizersAdresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSets\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicKeySet\",\"type\":\"bytes32[]\"}],\"name\":\"depositPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// PopcontractBin is the compiled bytecode used for deploying new contracts.
const PopcontractBin = `0x6060604052600b8054600160a060020a0319169055341561001f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b600480546000919060ff19166001835b02179055505b5b6112ee806100636000396000f300606060405236156101045763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c01b2ae81146101095780630c3f6acf146101fb5780631ae95ec414610232578063210790c81461025a57806323266dfe14610281578063232a6b9d146102e857806330756ec31461030f57806341676f151461039a57806341c0e1b5146103cd5780635a52ecf6146103e257806372ee91c2146104495780638e7d0f44146104785780639f638dbf1461049f578063a9950fbc146104d1578063b9bcc99f146104f6578063c1b8471b14610528578063c323f8841461055a578063d7e3a4d3146105bd578063eff64cb3146105e2575b600080fd5b341561011457600080fd5b6101e760046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650509335935061066d92505050565b604051901515815260200160405180910390f35b341561020657600080fd5b61020e610760565b6040518082600481111561021e57fe5b60ff16815260200191505060405180910390f35b341561023d57600080fd5b610248600435610769565b60405190815260200160405180910390f35b341561026557600080fd5b6101e761078c565b604051901515815260200160405180910390f35b341561028c57600080fd5b6102946108ad565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102d45780820151818401525b6020016102bb565b505050509050019250505060405180910390f35b34156102f357600080fd5b6101e7610916565b604051901515815260200160405180910390f35b341561031a57600080fd5b610322610937565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561035f5780820151818401525b602001610346565b50505050905090810190601f16801561038c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103a557600080fd5b6101e7600160a060020a03600435166109d5565b604051901515815260200160405180910390f35b34156103d857600080fd5b6103e0610a85565b005b34156103ed57600080fd5b610294610aad565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102d45780820151818401525b6020016102bb565b505050509050019250505060405180910390f35b341561045457600080fd5b61045c610b16565b604051600160a060020a03909116815260200160405180910390f35b341561048357600080fd5b6101e7610c2f565b604051901515815260200160405180910390f35b34156104aa57600080fd5b61045c600435610d25565b604051600160a060020a03909116815260200160405180910390f35b34156104dc57600080fd5b610248610d57565b60405190815260200160405180910390f35b341561050157600080fd5b61045c600435610d5d565b604051600160a060020a03909116815260200160405180910390f35b341561053357600080fd5b61045c600435610d8f565b604051600160a060020a03909116815260200160405180910390f35b341561056557600080fd5b6101e76004602481358181019083013580602081810201604051908101604052809392919081815260200183836020028082843750949650610dbe95505050505050565b604051901515815260200160405180910390f35b34156105c857600080fd5b610248610ee5565b60405190815260200160405180910390f35b34156105ed57600080fd5b610322610eeb565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561035f5780820151818401525b602001610346565b50505050905090810190601f16801561038c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60048054600091829160ff169081111561068357fe5b81600481111561068f57fe5b14156101045760005433600160a060020a03908116911614156101045760058780516106bf929160200190610f89565b5042603c84020160075560068680516106dc929160200190610f89565b506008859055846106ee600982611008565b506009848051610702929160200190611032565b50600480546001919060ff191682805b02179055506000198501610727600a82611008565b506001610735600282611008565b506000196003556001915061074a565b600080fd5b5b610755565b600080fd5b5b5095945050505050565b60045460ff1681565b600280548290811061077757fe5b906000526020600020900160005b5054905081565b600480546000918291829160019160ff16908111156107a757fe5b8160048111156107b357fe5b1415610104576000805490935033600160a060020a039081169116141561089a57600091505b600a5482101561087a57600b54600a8054600160a060020a03909216918490811061080057fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a0316141561083457600080fd5b600b805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600192505b5b6001909101906107d9565b821561089a57600480546002919060ff19166001835b0217905550600193505b5b5b6108a6565b600080fd5b5b50505090565b6108b56110d1565b600980548060200260200160405190810160405280929190818152602001828054801561090b57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108ed575b505050505090505b90565b600b5474010000000000000000000000000000000000000000900460ff1681565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109cd5780601f106109a2576101008083540402835291602001916109cd565b820191906000526020600020905b8154815290600101906020018083116109b057829003601f168201915b505050505081565b60048054600091829160029160ff909116908111156109f057fe5b8160048111156109fc57fe5b141561010457600091505b600854821015610a6e5783600160a060020a0316600983815481101515610a2a57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03161415610a625760019250610a73565b5b600190910190610a07565b600092505b610a7d565b600080fd5b5b5050919050565b60005433600160a060020a0390811691161415610aaa57600054600160a060020a0316ff5b5b565b610ab56110d1565b600a80548060200260200160405190810160405280929190818152602001828054801561090b57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108ed575b505050505090505b90565b60048054600091829160019160ff90911690811115610b3157fe5b816004811115610b3d57fe5b141561010457600091505b600854821015610c10576009805483908110610b6057fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031633600160a060020a03161415610c0457600a805460018101610bab8382611008565b916000526020600020900160005b6009805486908110610bc757fe5b906000526020600020900160005b9054835461010093840a600160a060020a039390940a9091048216830292909102191617905550339250610c1f565b5b600190910190610b48565b600b54600160a060020a031692505b610c29565b600080fd5b5b505090565b6004805460009160039160ff1690811115610c4657fe5b816004811115610c5257fe5b1415610104576007544211610d145760005433600160a060020a039081169116148015610c83575060035460001914155b15610d0f57600b805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600354600180549091908110610ccf57fe5b906000526020600020906002020160005b506001018054610cf29160029161110d565b5060048054819060ff19166001825b021790555060019150610d14565b600091505b5b5b610d20565b600080fd5b5b5090565b600a805482908110610d3357fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b60035481565b6009805482908110610d3357fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b6001805482908110610d9d57fe5b906000526020600020906002020160005b5054600160a060020a0316905081565b6004805460009160029160ff1690811115610dd557fe5b816004811115610de157fe5b1415610104576007544211610ed257610df9336109d5565b15610ecd5760018054808201610e0f838261115e565b916000526020600020906002020160005b60408051908101604052600160a060020a0333168152602081018790529190508151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101908051610e84929160200190611190565b50506003805460010181559150610e989050565b6004805460ff1690811115610ea957fe5b14610ec457600480546003919060ff19166001835b02179055505b60019150610ed2565b600091505b5b5b610ede565b600080fd5b5b50919050565b60075481565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109cd5780601f106109a2576101008083540402835291602001916109cd565b820191906000526020600020905b8154815290600101906020018083116109b057829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610fca57805160ff1916838001178555610ff7565b82800160010185558215610ff7579182015b82811115610ff7578251825591602001919060010190610fdc565b5b50610d209291506111de565b5090565b81548183558181151161102c5760008381526020902061102c9181019083016111de565b5b505050565b828054828255906000526020600020908101928215611096579160200282015b82811115611096578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617825560209290920191600190910190611052565b5b50610d209291506111ff565b5090565b81548183558181151161102c5760008381526020902061102c9181019083016111de565b5b505050565b60206040519081016040526000815290565b81548183558181151161102c5760008381526020902061102c9181019083016111de565b5b505050565b828054828255906000526020600020908101928215610ff75760005260206000209182015b82811115610ff7578254825591600101919060010190611132565b5b50610d209291506111de565b5090565b81548183558181151161102c5760020281600202836000526020600020918201910161102c9190611258565b5b505050565b828054828255906000526020600020908101928215610ff7579160200282015b82811115610ff757825182556020909201916001909101906111b0565b5b50610d209291506111de565b5090565b61091391905b80821115610d2057600081556001016111e4565b5090565b90565b61091391905b80821115610d2057805473ffffffffffffffffffffffffffffffffffffffff19168155600101611205565b5090565b90565b61091391905b80821115610d2057600081556001016111e4565b5090565b90565b61091391905b80821115610d2057805473ffffffffffffffffffffffffffffffffffffffff19168155600061129060018301826112a0565b5060020161125e565b5090565b90565b50805460008255906000526020600020908101906112be91906111de565b5b505600a165627a7a7230582051b5bbca1be5b1e33847def16c38035e6253a66b5cf188c2834124ad08687cf80029`

// DeployPopcontract deploys a new Ethereum contract, binding an instance of Popcontract to it.
func DeployPopcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Popcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(PopcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PopcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Popcontract{PopcontractCaller: PopcontractCaller{contract: contract}, PopcontractTransactor: PopcontractTransactor{contract: contract}}, nil
}

// Popcontract is an auto generated Go binding around an Ethereum contract.
type Popcontract struct {
	PopcontractCaller     // Read-only binding to the contract
	PopcontractTransactor // Write-only binding to the contract
}

// PopcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PopcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PopcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PopcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PopcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PopcontractSession struct {
	Contract     *Popcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PopcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PopcontractCallerSession struct {
	Contract *PopcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PopcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PopcontractTransactorSession struct {
	Contract     *PopcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PopcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PopcontractRaw struct {
	Contract *Popcontract // Generic contract binding to access the raw methods on
}

// PopcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PopcontractCallerRaw struct {
	Contract *PopcontractCaller // Generic read-only contract binding to access the raw methods on
}

// PopcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PopcontractTransactorRaw struct {
	Contract *PopcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPopcontract creates a new instance of Popcontract, bound to a specific deployed contract.
func NewPopcontract(address common.Address, backend bind.ContractBackend) (*Popcontract, error) {
	contract, err := bindPopcontract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Popcontract{PopcontractCaller: PopcontractCaller{contract: contract}, PopcontractTransactor: PopcontractTransactor{contract: contract}}, nil
}

// NewPopcontractCaller creates a new read-only instance of Popcontract, bound to a specific deployed contract.
func NewPopcontractCaller(address common.Address, caller bind.ContractCaller) (*PopcontractCaller, error) {
	contract, err := bindPopcontract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PopcontractCaller{contract: contract}, nil
}

// NewPopcontractTransactor creates a new write-only instance of Popcontract, bound to a specific deployed contract.
func NewPopcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*PopcontractTransactor, error) {
	contract, err := bindPopcontract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PopcontractTransactor{contract: contract}, nil
}

// bindPopcontract binds a generic wrapper to an already deployed contract.
func bindPopcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PopcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Popcontract *PopcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Popcontract.Contract.PopcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Popcontract *PopcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.Contract.PopcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Popcontract *PopcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Popcontract.Contract.PopcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Popcontract *PopcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Popcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Popcontract *PopcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Popcontract *PopcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Popcontract.Contract.contract.Transact(opts, method, params...)
}

// AllSets is a free data retrieval call binding the contract method 0xc1b8471b.
//
// Solidity: function allSets( uint256) constant returns(sender address)
func (_Popcontract *PopcontractCaller) AllSets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "allSets", arg0)
	return *ret0, err
}

// AllSets is a free data retrieval call binding the contract method 0xc1b8471b.
//
// Solidity: function allSets( uint256) constant returns(sender address)
func (_Popcontract *PopcontractSession) AllSets(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.AllSets(&_Popcontract.CallOpts, arg0)
}

// AllSets is a free data retrieval call binding the contract method 0xc1b8471b.
//
// Solidity: function allSets( uint256) constant returns(sender address)
func (_Popcontract *PopcontractCallerSession) AllSets(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.AllSets(&_Popcontract.CallOpts, arg0)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() constant returns(uint8)
func (_Popcontract *PopcontractCaller) CurrentState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "currentState")
	return *ret0, err
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() constant returns(uint8)
func (_Popcontract *PopcontractSession) CurrentState() (uint8, error) {
	return _Popcontract.Contract.CurrentState(&_Popcontract.CallOpts)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() constant returns(uint8)
func (_Popcontract *PopcontractCallerSession) CurrentState() (uint8, error) {
	return _Popcontract.Contract.CurrentState(&_Popcontract.CallOpts)
}

// EndOfParty is a free data retrieval call binding the contract method 0xd7e3a4d3.
//
// Solidity: function endOfParty() constant returns(uint256)
func (_Popcontract *PopcontractCaller) EndOfParty(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "endOfParty")
	return *ret0, err
}

// EndOfParty is a free data retrieval call binding the contract method 0xd7e3a4d3.
//
// Solidity: function endOfParty() constant returns(uint256)
func (_Popcontract *PopcontractSession) EndOfParty() (*big.Int, error) {
	return _Popcontract.Contract.EndOfParty(&_Popcontract.CallOpts)
}

// EndOfParty is a free data retrieval call binding the contract method 0xd7e3a4d3.
//
// Solidity: function endOfParty() constant returns(uint256)
func (_Popcontract *PopcontractCallerSession) EndOfParty() (*big.Int, error) {
	return _Popcontract.Contract.EndOfParty(&_Popcontract.CallOpts)
}

// FinalKeySet is a free data retrieval call binding the contract method 0x1ae95ec4.
//
// Solidity: function finalKeySet( uint256) constant returns(bytes32)
func (_Popcontract *PopcontractCaller) FinalKeySet(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "finalKeySet", arg0)
	return *ret0, err
}

// FinalKeySet is a free data retrieval call binding the contract method 0x1ae95ec4.
//
// Solidity: function finalKeySet( uint256) constant returns(bytes32)
func (_Popcontract *PopcontractSession) FinalKeySet(arg0 *big.Int) ([32]byte, error) {
	return _Popcontract.Contract.FinalKeySet(&_Popcontract.CallOpts, arg0)
}

// FinalKeySet is a free data retrieval call binding the contract method 0x1ae95ec4.
//
// Solidity: function finalKeySet( uint256) constant returns(bytes32)
func (_Popcontract *PopcontractCallerSession) FinalKeySet(arg0 *big.Int) ([32]byte, error) {
	return _Popcontract.Contract.FinalKeySet(&_Popcontract.CallOpts, arg0)
}

// GetOrganizersAddresses is a free data retrieval call binding the contract method 0x23266dfe.
//
// Solidity: function getOrganizersAddresses() constant returns(address[])
func (_Popcontract *PopcontractCaller) GetOrganizersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "getOrganizersAddresses")
	return *ret0, err
}

// GetOrganizersAddresses is a free data retrieval call binding the contract method 0x23266dfe.
//
// Solidity: function getOrganizersAddresses() constant returns(address[])
func (_Popcontract *PopcontractSession) GetOrganizersAddresses() ([]common.Address, error) {
	return _Popcontract.Contract.GetOrganizersAddresses(&_Popcontract.CallOpts)
}

// GetOrganizersAddresses is a free data retrieval call binding the contract method 0x23266dfe.
//
// Solidity: function getOrganizersAddresses() constant returns(address[])
func (_Popcontract *PopcontractCallerSession) GetOrganizersAddresses() ([]common.Address, error) {
	return _Popcontract.Contract.GetOrganizersAddresses(&_Popcontract.CallOpts)
}

// GetSignedConfiguration is a free data retrieval call binding the contract method 0x5a52ecf6.
//
// Solidity: function getSignedConfiguration() constant returns(address[])
func (_Popcontract *PopcontractCaller) GetSignedConfiguration(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "getSignedConfiguration")
	return *ret0, err
}

// GetSignedConfiguration is a free data retrieval call binding the contract method 0x5a52ecf6.
//
// Solidity: function getSignedConfiguration() constant returns(address[])
func (_Popcontract *PopcontractSession) GetSignedConfiguration() ([]common.Address, error) {
	return _Popcontract.Contract.GetSignedConfiguration(&_Popcontract.CallOpts)
}

// GetSignedConfiguration is a free data retrieval call binding the contract method 0x5a52ecf6.
//
// Solidity: function getSignedConfiguration() constant returns(address[])
func (_Popcontract *PopcontractCallerSession) GetSignedConfiguration() ([]common.Address, error) {
	return _Popcontract.Contract.GetSignedConfiguration(&_Popcontract.CallOpts)
}

// KeySetAdded is a free data retrieval call binding the contract method 0xa9950fbc.
//
// Solidity: function keySetAdded() constant returns(int256)
func (_Popcontract *PopcontractCaller) KeySetAdded(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "keySetAdded")
	return *ret0, err
}

// KeySetAdded is a free data retrieval call binding the contract method 0xa9950fbc.
//
// Solidity: function keySetAdded() constant returns(int256)
func (_Popcontract *PopcontractSession) KeySetAdded() (*big.Int, error) {
	return _Popcontract.Contract.KeySetAdded(&_Popcontract.CallOpts)
}

// KeySetAdded is a free data retrieval call binding the contract method 0xa9950fbc.
//
// Solidity: function keySetAdded() constant returns(int256)
func (_Popcontract *PopcontractCallerSession) KeySetAdded() (*big.Int, error) {
	return _Popcontract.Contract.KeySetAdded(&_Popcontract.CallOpts)
}

// LocationOfParty is a free data retrieval call binding the contract method 0x30756ec3.
//
// Solidity: function locationOfParty() constant returns(string)
func (_Popcontract *PopcontractCaller) LocationOfParty(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "locationOfParty")
	return *ret0, err
}

// LocationOfParty is a free data retrieval call binding the contract method 0x30756ec3.
//
// Solidity: function locationOfParty() constant returns(string)
func (_Popcontract *PopcontractSession) LocationOfParty() (string, error) {
	return _Popcontract.Contract.LocationOfParty(&_Popcontract.CallOpts)
}

// LocationOfParty is a free data retrieval call binding the contract method 0x30756ec3.
//
// Solidity: function locationOfParty() constant returns(string)
func (_Popcontract *PopcontractCallerSession) LocationOfParty() (string, error) {
	return _Popcontract.Contract.LocationOfParty(&_Popcontract.CallOpts)
}

// NameOfParty is a free data retrieval call binding the contract method 0xeff64cb3.
//
// Solidity: function nameOfParty() constant returns(string)
func (_Popcontract *PopcontractCaller) NameOfParty(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "nameOfParty")
	return *ret0, err
}

// NameOfParty is a free data retrieval call binding the contract method 0xeff64cb3.
//
// Solidity: function nameOfParty() constant returns(string)
func (_Popcontract *PopcontractSession) NameOfParty() (string, error) {
	return _Popcontract.Contract.NameOfParty(&_Popcontract.CallOpts)
}

// NameOfParty is a free data retrieval call binding the contract method 0xeff64cb3.
//
// Solidity: function nameOfParty() constant returns(string)
func (_Popcontract *PopcontractCallerSession) NameOfParty() (string, error) {
	return _Popcontract.Contract.NameOfParty(&_Popcontract.CallOpts)
}

// OrganizersAdresses is a free data retrieval call binding the contract method 0xb9bcc99f.
//
// Solidity: function organizersAdresses( uint256) constant returns(address)
func (_Popcontract *PopcontractCaller) OrganizersAdresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "organizersAdresses", arg0)
	return *ret0, err
}

// OrganizersAdresses is a free data retrieval call binding the contract method 0xb9bcc99f.
//
// Solidity: function organizersAdresses( uint256) constant returns(address)
func (_Popcontract *PopcontractSession) OrganizersAdresses(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.OrganizersAdresses(&_Popcontract.CallOpts, arg0)
}

// OrganizersAdresses is a free data retrieval call binding the contract method 0xb9bcc99f.
//
// Solidity: function organizersAdresses( uint256) constant returns(address)
func (_Popcontract *PopcontractCallerSession) OrganizersAdresses(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.OrganizersAdresses(&_Popcontract.CallOpts, arg0)
}

// Signed is a free data retrieval call binding the contract method 0x232a6b9d.
//
// Solidity: function signed() constant returns(bool)
func (_Popcontract *PopcontractCaller) Signed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "signed")
	return *ret0, err
}

// Signed is a free data retrieval call binding the contract method 0x232a6b9d.
//
// Solidity: function signed() constant returns(bool)
func (_Popcontract *PopcontractSession) Signed() (bool, error) {
	return _Popcontract.Contract.Signed(&_Popcontract.CallOpts)
}

// Signed is a free data retrieval call binding the contract method 0x232a6b9d.
//
// Solidity: function signed() constant returns(bool)
func (_Popcontract *PopcontractCallerSession) Signed() (bool, error) {
	return _Popcontract.Contract.Signed(&_Popcontract.CallOpts)
}

// SignedConfiguration is a free data retrieval call binding the contract method 0x9f638dbf.
//
// Solidity: function signedConfiguration( uint256) constant returns(address)
func (_Popcontract *PopcontractCaller) SignedConfiguration(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "signedConfiguration", arg0)
	return *ret0, err
}

// SignedConfiguration is a free data retrieval call binding the contract method 0x9f638dbf.
//
// Solidity: function signedConfiguration( uint256) constant returns(address)
func (_Popcontract *PopcontractSession) SignedConfiguration(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.SignedConfiguration(&_Popcontract.CallOpts, arg0)
}

// SignedConfiguration is a free data retrieval call binding the contract method 0x9f638dbf.
//
// Solidity: function signedConfiguration( uint256) constant returns(address)
func (_Popcontract *PopcontractCallerSession) SignedConfiguration(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.SignedConfiguration(&_Popcontract.CallOpts, arg0)
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(address)
func (_Popcontract *PopcontractTransactor) ConfigSignOrganizers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "configSignOrganizers")
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(address)
func (_Popcontract *PopcontractSession) ConfigSignOrganizers() (*types.Transaction, error) {
	return _Popcontract.Contract.ConfigSignOrganizers(&_Popcontract.TransactOpts)
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(address)
func (_Popcontract *PopcontractTransactorSession) ConfigSignOrganizers() (*types.Transaction, error) {
	return _Popcontract.Contract.ConfigSignOrganizers(&_Popcontract.TransactOpts)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0xc323f884.
//
// Solidity: function depositPublicKeys(_publicKeySet bytes32[]) returns(bool)
func (_Popcontract *PopcontractTransactor) DepositPublicKeys(opts *bind.TransactOpts, _publicKeySet [][32]byte) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "depositPublicKeys", _publicKeySet)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0xc323f884.
//
// Solidity: function depositPublicKeys(_publicKeySet bytes32[]) returns(bool)
func (_Popcontract *PopcontractSession) DepositPublicKeys(_publicKeySet [][32]byte) (*types.Transaction, error) {
	return _Popcontract.Contract.DepositPublicKeys(&_Popcontract.TransactOpts, _publicKeySet)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0xc323f884.
//
// Solidity: function depositPublicKeys(_publicKeySet bytes32[]) returns(bool)
func (_Popcontract *PopcontractTransactorSession) DepositPublicKeys(_publicKeySet [][32]byte) (*types.Transaction, error) {
	return _Popcontract.Contract.DepositPublicKeys(&_Popcontract.TransactOpts, _publicKeySet)
}

// IsOrganizer is a paid mutator transaction binding the contract method 0x41676f15.
//
// Solidity: function isOrganizer(sender address) returns(bool)
func (_Popcontract *PopcontractTransactor) IsOrganizer(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "isOrganizer", sender)
}

// IsOrganizer is a paid mutator transaction binding the contract method 0x41676f15.
//
// Solidity: function isOrganizer(sender address) returns(bool)
func (_Popcontract *PopcontractSession) IsOrganizer(sender common.Address) (*types.Transaction, error) {
	return _Popcontract.Contract.IsOrganizer(&_Popcontract.TransactOpts, sender)
}

// IsOrganizer is a paid mutator transaction binding the contract method 0x41676f15.
//
// Solidity: function isOrganizer(sender address) returns(bool)
func (_Popcontract *PopcontractTransactorSession) IsOrganizer(sender common.Address) (*types.Transaction, error) {
	return _Popcontract.Contract.IsOrganizer(&_Popcontract.TransactOpts, sender)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Popcontract *PopcontractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Popcontract *PopcontractSession) Kill() (*types.Transaction, error) {
	return _Popcontract.Contract.Kill(&_Popcontract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Popcontract *PopcontractTransactorSession) Kill() (*types.Transaction, error) {
	return _Popcontract.Contract.Kill(&_Popcontract.TransactOpts)
}

// PublicKeyConsensus is a paid mutator transaction binding the contract method 0x8e7d0f44.
//
// Solidity: function publicKeyConsensus() returns(bool)
func (_Popcontract *PopcontractTransactor) PublicKeyConsensus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "publicKeyConsensus")
}

// PublicKeyConsensus is a paid mutator transaction binding the contract method 0x8e7d0f44.
//
// Solidity: function publicKeyConsensus() returns(bool)
func (_Popcontract *PopcontractSession) PublicKeyConsensus() (*types.Transaction, error) {
	return _Popcontract.Contract.PublicKeyConsensus(&_Popcontract.TransactOpts)
}

// PublicKeyConsensus is a paid mutator transaction binding the contract method 0x8e7d0f44.
//
// Solidity: function publicKeyConsensus() returns(bool)
func (_Popcontract *PopcontractTransactorSession) PublicKeyConsensus() (*types.Transaction, error) {
	return _Popcontract.Contract.PublicKeyConsensus(&_Popcontract.TransactOpts)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0x0c01b2ae.
//
// Solidity: function setConfiguration(_nameOfParty string, place string, organizers uint256, data address[], durationInMinutes uint256) returns(bool)
func (_Popcontract *PopcontractTransactor) SetConfiguration(opts *bind.TransactOpts, _nameOfParty string, place string, organizers *big.Int, data []common.Address, durationInMinutes *big.Int) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "setConfiguration", _nameOfParty, place, organizers, data, durationInMinutes)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0x0c01b2ae.
//
// Solidity: function setConfiguration(_nameOfParty string, place string, organizers uint256, data address[], durationInMinutes uint256) returns(bool)
func (_Popcontract *PopcontractSession) SetConfiguration(_nameOfParty string, place string, organizers *big.Int, data []common.Address, durationInMinutes *big.Int) (*types.Transaction, error) {
	return _Popcontract.Contract.SetConfiguration(&_Popcontract.TransactOpts, _nameOfParty, place, organizers, data, durationInMinutes)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0x0c01b2ae.
//
// Solidity: function setConfiguration(_nameOfParty string, place string, organizers uint256, data address[], durationInMinutes uint256) returns(bool)
func (_Popcontract *PopcontractTransactorSession) SetConfiguration(_nameOfParty string, place string, organizers *big.Int, data []common.Address, durationInMinutes *big.Int) (*types.Transaction, error) {
	return _Popcontract.Contract.SetConfiguration(&_Popcontract.TransactOpts, _nameOfParty, place, organizers, data, durationInMinutes)
}

// SignWholeConfiguration is a paid mutator transaction binding the contract method 0x210790c8.
//
// Solidity: function signWholeConfiguration() returns(bool)
func (_Popcontract *PopcontractTransactor) SignWholeConfiguration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "signWholeConfiguration")
}

// SignWholeConfiguration is a paid mutator transaction binding the contract method 0x210790c8.
//
// Solidity: function signWholeConfiguration() returns(bool)
func (_Popcontract *PopcontractSession) SignWholeConfiguration() (*types.Transaction, error) {
	return _Popcontract.Contract.SignWholeConfiguration(&_Popcontract.TransactOpts)
}

// SignWholeConfiguration is a paid mutator transaction binding the contract method 0x210790c8.
//
// Solidity: function signWholeConfiguration() returns(bool)
func (_Popcontract *PopcontractTransactorSession) SignWholeConfiguration() (*types.Transaction, error) {
	return _Popcontract.Contract.SignWholeConfiguration(&_Popcontract.TransactOpts)
}
