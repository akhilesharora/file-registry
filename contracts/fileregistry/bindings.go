// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fileregistry

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

// FileRegistryMetaData contains all meta data concerning the FileRegistry contract.
var FileRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107458061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063693ec85e14610038578063962939b814610068575b5f5ffd5b610052600480360381019061004d91906102ad565b610084565b60405161005f9190610354565b60405180910390f35b610082600480360381019061007d9190610374565b610131565b005b60605f826040516100959190610424565b908152602001604051809103902080546100ae90610467565b80601f01602080910402602001604051908101604052809291908181526020018280546100da90610467565b80156101255780601f106100fc57610100808354040283529160200191610125565b820191905f5260205f20905b81548152906001019060200180831161010857829003601f168201915b50505050509050919050565b805f836040516101419190610424565b9081526020016040518091039020908161015b9190610640565b505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101bf82610179565b810181811067ffffffffffffffff821117156101de576101dd610189565b5b80604052505050565b5f6101f0610160565b90506101fc82826101b6565b919050565b5f67ffffffffffffffff82111561021b5761021a610189565b5b61022482610179565b9050602081019050919050565b828183375f83830152505050565b5f61025161024c84610201565b6101e7565b90508281526020810184848401111561026d5761026c610175565b5b610278848285610231565b509392505050565b5f82601f83011261029457610293610171565b5b81356102a484826020860161023f565b91505092915050565b5f602082840312156102c2576102c1610169565b5b5f82013567ffffffffffffffff8111156102df576102de61016d565b5b6102eb84828501610280565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610326826102f4565b61033081856102fe565b935061034081856020860161030e565b61034981610179565b840191505092915050565b5f6020820190508181035f83015261036c818461031c565b905092915050565b5f5f6040838503121561038a57610389610169565b5b5f83013567ffffffffffffffff8111156103a7576103a661016d565b5b6103b385828601610280565b925050602083013567ffffffffffffffff8111156103d4576103d361016d565b5b6103e085828601610280565b9150509250929050565b5f81905092915050565b5f6103fe826102f4565b61040881856103ea565b935061041881856020860161030e565b80840191505092915050565b5f61042f82846103f4565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061047e57607f821691505b6020821081036104915761049061043a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104f37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826104b8565b6104fd86836104b8565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61054161053c61053784610515565b61051e565b610515565b9050919050565b5f819050919050565b61055a83610527565b61056e61056682610548565b8484546104c4565b825550505050565b5f5f905090565b610585610576565b610590818484610551565b505050565b5b818110156105b3576105a85f8261057d565b600181019050610596565b5050565b601f8211156105f8576105c981610497565b6105d2846104a9565b810160208510156105e1578190505b6105f56105ed856104a9565b830182610595565b50505b505050565b5f82821c905092915050565b5f6106185f19846008026105fd565b1980831691505092915050565b5f6106308383610609565b9150826002028217905092915050565b610649826102f4565b67ffffffffffffffff81111561066257610661610189565b5b61066c8254610467565b6106778282856105b7565b5f60209050601f8311600181146106a8575f8415610696578287015190505b6106a08582610625565b865550610707565b601f1984166106b686610497565b5f5b828110156106dd578489015182556001820191506020850194506020810190506106b8565b868310156106fa57848901516106f6601f891682610609565b8355505b6001600288020188555050505b50505050505056fea264697066735822122007847a58f076a55a2af4f781f40d0153555f38f1d818a5dbc19395b724434fc464736f6c634300081c0033",
}

// FileRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FileRegistryMetaData.ABI instead.
var FileRegistryABI = FileRegistryMetaData.ABI

// FileRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FileRegistryMetaData.Bin instead.
var FileRegistryBin = FileRegistryMetaData.Bin

// DeployFileRegistry deploys a new Ethereum contract, binding an instance of FileRegistry to it.
func DeployFileRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FileRegistry, error) {
	parsed, err := FileRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FileRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileRegistry{FileRegistryCaller: FileRegistryCaller{contract: contract}, FileRegistryTransactor: FileRegistryTransactor{contract: contract}, FileRegistryFilterer: FileRegistryFilterer{contract: contract}}, nil
}

// FileRegistry is an auto generated Go binding around an Ethereum contract.
type FileRegistry struct {
	FileRegistryCaller     // Read-only binding to the contract
	FileRegistryTransactor // Write-only binding to the contract
	FileRegistryFilterer   // Log filterer for contract events
}

// FileRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileRegistrySession struct {
	Contract     *FileRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileRegistryCallerSession struct {
	Contract *FileRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FileRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileRegistryTransactorSession struct {
	Contract     *FileRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FileRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileRegistryRaw struct {
	Contract *FileRegistry // Generic contract binding to access the raw methods on
}

// FileRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileRegistryCallerRaw struct {
	Contract *FileRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FileRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileRegistryTransactorRaw struct {
	Contract *FileRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileRegistry creates a new instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistry(address common.Address, backend bind.ContractBackend) (*FileRegistry, error) {
	contract, err := bindFileRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileRegistry{FileRegistryCaller: FileRegistryCaller{contract: contract}, FileRegistryTransactor: FileRegistryTransactor{contract: contract}, FileRegistryFilterer: FileRegistryFilterer{contract: contract}}, nil
}

// NewFileRegistryCaller creates a new read-only instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistryCaller(address common.Address, caller bind.ContractCaller) (*FileRegistryCaller, error) {
	contract, err := bindFileRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileRegistryCaller{contract: contract}, nil
}

// NewFileRegistryTransactor creates a new write-only instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FileRegistryTransactor, error) {
	contract, err := bindFileRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileRegistryTransactor{contract: contract}, nil
}

// NewFileRegistryFilterer creates a new log filterer instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FileRegistryFilterer, error) {
	contract, err := bindFileRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileRegistryFilterer{contract: contract}, nil
}

// bindFileRegistry binds a generic wrapper to an already deployed contract.
func bindFileRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FileRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileRegistry *FileRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileRegistry.Contract.FileRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileRegistry *FileRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileRegistry.Contract.FileRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileRegistry *FileRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileRegistry.Contract.FileRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileRegistry *FileRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileRegistry *FileRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileRegistry *FileRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileRegistry.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_FileRegistry *FileRegistryCaller) Get(opts *bind.CallOpts, filePath string) (string, error) {
	var out []interface{}
	err := _FileRegistry.contract.Call(opts, &out, "get", filePath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_FileRegistry *FileRegistrySession) Get(filePath string) (string, error) {
	return _FileRegistry.Contract.Get(&_FileRegistry.CallOpts, filePath)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_FileRegistry *FileRegistryCallerSession) Get(filePath string) (string, error) {
	return _FileRegistry.Contract.Get(&_FileRegistry.CallOpts, filePath)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_FileRegistry *FileRegistryTransactor) Save(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error) {
	return _FileRegistry.contract.Transact(opts, "save", filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_FileRegistry *FileRegistrySession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _FileRegistry.Contract.Save(&_FileRegistry.TransactOpts, filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_FileRegistry *FileRegistryTransactorSession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _FileRegistry.Contract.Save(&_FileRegistry.TransactOpts, filePath, cid)
}
