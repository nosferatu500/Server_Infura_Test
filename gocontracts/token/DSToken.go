// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DSAuthABI is the input ABI used to generate the binding from.
const DSAuthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// DSAuthBin is the compiled bytecode used for deploying new contracts.
const DSAuthBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a2610384806100686000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af403581146100665780637a9e5e4b146100875780638da5cb5b146100a6578063bf7e214f146100d5575b600080fd5b341561007157600080fd5b610085600160a060020a03600435166100e8565b005b341561009257600080fd5b610085600160a060020a0360043516610179565b34156100b157600080fd5b6100b9610205565b604051600160a060020a03909116815260200160405180910390f35b34156100e057600080fd5b6100b9610214565b61011b610116336000357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916610223565b610349565b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b6101a7610116336000357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916610223565b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600154600160a060020a031681565b600054600160a060020a031681565b600030600160a060020a031683600160a060020a0316141561024757506001610343565b600154600160a060020a038481169116141561026557506001610343565b600054600160a060020a0316151561027f57506000610343565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a0393841660048201529190921660248201527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199091166044820152606401602060405180830381600087803b151561032657600080fd5b6102c65a03f1151561033757600080fd5b50505060405180519150505b92915050565b80151561035557600080fd5b505600a165627a7a723058202a003806fef6a7fab2267c5c4af5a97df5b6f110d82d8a84189824d23111c8cc0029`

// DeployDSAuth deploys a new Ethereum contract, binding an instance of DSAuth to it.
func DeployDSAuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSAuth, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSAuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSAuth{DSAuthCaller: DSAuthCaller{contract: contract}, DSAuthTransactor: DSAuthTransactor{contract: contract}}, nil
}

// DSAuth is an auto generated Go binding around an Ethereum contract.
type DSAuth struct {
	DSAuthCaller     // Read-only binding to the contract
	DSAuthTransactor // Write-only binding to the contract
}

// DSAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSAuthSession struct {
	Contract     *DSAuth           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSAuthCallerSession struct {
	Contract *DSAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSAuthTransactorSession struct {
	Contract     *DSAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSAuthRaw struct {
	Contract *DSAuth // Generic contract binding to access the raw methods on
}

// DSAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSAuthCallerRaw struct {
	Contract *DSAuthCaller // Generic read-only contract binding to access the raw methods on
}

// DSAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSAuthTransactorRaw struct {
	Contract *DSAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSAuth creates a new instance of DSAuth, bound to a specific deployed contract.
func NewDSAuth(address common.Address, backend bind.ContractBackend) (*DSAuth, error) {
	contract, err := bindDSAuth(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSAuth{DSAuthCaller: DSAuthCaller{contract: contract}, DSAuthTransactor: DSAuthTransactor{contract: contract}}, nil
}

// NewDSAuthCaller creates a new read-only instance of DSAuth, bound to a specific deployed contract.
func NewDSAuthCaller(address common.Address, caller bind.ContractCaller) (*DSAuthCaller, error) {
	contract, err := bindDSAuth(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthCaller{contract: contract}, nil
}

// NewDSAuthTransactor creates a new write-only instance of DSAuth, bound to a specific deployed contract.
func NewDSAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*DSAuthTransactor, error) {
	contract, err := bindDSAuth(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSAuthTransactor{contract: contract}, nil
}

// bindDSAuth binds a generic wrapper to an already deployed contract.
func bindDSAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuth *DSAuthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuth.Contract.DSAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuth *DSAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuth.Contract.DSAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuth *DSAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuth.Contract.DSAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuth *DSAuthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuth *DSAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuth *DSAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuth.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSAuth *DSAuthCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSAuth.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSAuth *DSAuthSession) Authority() (common.Address, error) {
	return _DSAuth.Contract.Authority(&_DSAuth.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSAuth *DSAuthCallerSession) Authority() (common.Address, error) {
	return _DSAuth.Contract.Authority(&_DSAuth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSAuth *DSAuthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSAuth.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSAuth *DSAuthSession) Owner() (common.Address, error) {
	return _DSAuth.Contract.Owner(&_DSAuth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSAuth *DSAuthCallerSession) Owner() (common.Address, error) {
	return _DSAuth.Contract.Owner(&_DSAuth.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSAuth *DSAuthTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DSAuth.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSAuth *DSAuthSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetAuthority(&_DSAuth.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSAuth *DSAuthTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetAuthority(&_DSAuth.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSAuth *DSAuthTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DSAuth.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSAuth *DSAuthSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetOwner(&_DSAuth.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSAuth *DSAuthTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetOwner(&_DSAuth.TransactOpts, owner_)
}

// DSAuthEventsABI is the input ABI used to generate the binding from.
const DSAuthEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// DSAuthEventsBin is the compiled bytecode used for deploying new contracts.
const DSAuthEventsBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820bb53fac702c9dec5d14b025a44c134244d08f3a3522a049608c5473b1940b7870029`

// DeployDSAuthEvents deploys a new Ethereum contract, binding an instance of DSAuthEvents to it.
func DeployDSAuthEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSAuthEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSAuthEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSAuthEvents{DSAuthEventsCaller: DSAuthEventsCaller{contract: contract}, DSAuthEventsTransactor: DSAuthEventsTransactor{contract: contract}}, nil
}

// DSAuthEvents is an auto generated Go binding around an Ethereum contract.
type DSAuthEvents struct {
	DSAuthEventsCaller     // Read-only binding to the contract
	DSAuthEventsTransactor // Write-only binding to the contract
}

// DSAuthEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSAuthEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSAuthEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSAuthEventsSession struct {
	Contract     *DSAuthEvents     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSAuthEventsCallerSession struct {
	Contract *DSAuthEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DSAuthEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSAuthEventsTransactorSession struct {
	Contract     *DSAuthEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DSAuthEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSAuthEventsRaw struct {
	Contract *DSAuthEvents // Generic contract binding to access the raw methods on
}

// DSAuthEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSAuthEventsCallerRaw struct {
	Contract *DSAuthEventsCaller // Generic read-only contract binding to access the raw methods on
}

// DSAuthEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSAuthEventsTransactorRaw struct {
	Contract *DSAuthEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSAuthEvents creates a new instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEvents(address common.Address, backend bind.ContractBackend) (*DSAuthEvents, error) {
	contract, err := bindDSAuthEvents(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSAuthEvents{DSAuthEventsCaller: DSAuthEventsCaller{contract: contract}, DSAuthEventsTransactor: DSAuthEventsTransactor{contract: contract}}, nil
}

// NewDSAuthEventsCaller creates a new read-only instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEventsCaller(address common.Address, caller bind.ContractCaller) (*DSAuthEventsCaller, error) {
	contract, err := bindDSAuthEvents(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsCaller{contract: contract}, nil
}

// NewDSAuthEventsTransactor creates a new write-only instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*DSAuthEventsTransactor, error) {
	contract, err := bindDSAuthEvents(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsTransactor{contract: contract}, nil
}

// bindDSAuthEvents binds a generic wrapper to an already deployed contract.
func bindDSAuthEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthEvents *DSAuthEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthEvents.Contract.DSAuthEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthEvents *DSAuthEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.DSAuthEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthEvents *DSAuthEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.DSAuthEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthEvents *DSAuthEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthEvents *DSAuthEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthEvents *DSAuthEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.contract.Transact(opts, method, params...)
}

// DSAuthorityABI is the input ABI used to generate the binding from.
const DSAuthorityABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DSAuthorityBin is the compiled bytecode used for deploying new contracts.
const DSAuthorityBin = `0x`

// DeployDSAuthority deploys a new Ethereum contract, binding an instance of DSAuthority to it.
func DeployDSAuthority(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSAuthority, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthorityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSAuthorityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSAuthority{DSAuthorityCaller: DSAuthorityCaller{contract: contract}, DSAuthorityTransactor: DSAuthorityTransactor{contract: contract}}, nil
}

// DSAuthority is an auto generated Go binding around an Ethereum contract.
type DSAuthority struct {
	DSAuthorityCaller     // Read-only binding to the contract
	DSAuthorityTransactor // Write-only binding to the contract
}

// DSAuthorityCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSAuthorityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthorityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSAuthorityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthoritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSAuthoritySession struct {
	Contract     *DSAuthority      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthorityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSAuthorityCallerSession struct {
	Contract *DSAuthorityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DSAuthorityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSAuthorityTransactorSession struct {
	Contract     *DSAuthorityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DSAuthorityRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSAuthorityRaw struct {
	Contract *DSAuthority // Generic contract binding to access the raw methods on
}

// DSAuthorityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSAuthorityCallerRaw struct {
	Contract *DSAuthorityCaller // Generic read-only contract binding to access the raw methods on
}

// DSAuthorityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSAuthorityTransactorRaw struct {
	Contract *DSAuthorityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSAuthority creates a new instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthority(address common.Address, backend bind.ContractBackend) (*DSAuthority, error) {
	contract, err := bindDSAuthority(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSAuthority{DSAuthorityCaller: DSAuthorityCaller{contract: contract}, DSAuthorityTransactor: DSAuthorityTransactor{contract: contract}}, nil
}

// NewDSAuthorityCaller creates a new read-only instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthorityCaller(address common.Address, caller bind.ContractCaller) (*DSAuthorityCaller, error) {
	contract, err := bindDSAuthority(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthorityCaller{contract: contract}, nil
}

// NewDSAuthorityTransactor creates a new write-only instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthorityTransactor(address common.Address, transactor bind.ContractTransactor) (*DSAuthorityTransactor, error) {
	contract, err := bindDSAuthority(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSAuthorityTransactor{contract: contract}, nil
}

// bindDSAuthority binds a generic wrapper to an already deployed contract.
func bindDSAuthority(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthorityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthority *DSAuthorityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthority.Contract.DSAuthorityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthority *DSAuthorityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthority.Contract.DSAuthorityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthority *DSAuthorityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthority.Contract.DSAuthorityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthority *DSAuthorityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthority.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthority *DSAuthorityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthority.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthority *DSAuthorityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthority.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_DSAuthority *DSAuthorityCaller) CanCall(opts *bind.CallOpts, src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSAuthority.contract.Call(opts, out, "canCall", src, dst, sig)
	return *ret0, err
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_DSAuthority *DSAuthoritySession) CanCall(src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	return _DSAuthority.Contract.CanCall(&_DSAuthority.CallOpts, src, dst, sig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_DSAuthority *DSAuthorityCallerSession) CanCall(src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	return _DSAuthority.Contract.CanCall(&_DSAuthority.CallOpts, src, dst, sig)
}

// DSMathABI is the input ABI used to generate the binding from.
const DSMathABI = "[]"

// DSMathBin is the compiled bytecode used for deploying new contracts.
const DSMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058200e7e9f6f7e3fd85a11ec377aae2a46724a14660c46f5d3cf0e3d3f35ee5e94200029`

// DeployDSMath deploys a new Ethereum contract, binding an instance of DSMath to it.
func DeployDSMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSMath, error) {
	parsed, err := abi.JSON(strings.NewReader(DSMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSMath{DSMathCaller: DSMathCaller{contract: contract}, DSMathTransactor: DSMathTransactor{contract: contract}}, nil
}

// DSMath is an auto generated Go binding around an Ethereum contract.
type DSMath struct {
	DSMathCaller     // Read-only binding to the contract
	DSMathTransactor // Write-only binding to the contract
}

// DSMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSMathSession struct {
	Contract     *DSMath           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSMathCallerSession struct {
	Contract *DSMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSMathTransactorSession struct {
	Contract     *DSMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSMathRaw struct {
	Contract *DSMath // Generic contract binding to access the raw methods on
}

// DSMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSMathCallerRaw struct {
	Contract *DSMathCaller // Generic read-only contract binding to access the raw methods on
}

// DSMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSMathTransactorRaw struct {
	Contract *DSMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSMath creates a new instance of DSMath, bound to a specific deployed contract.
func NewDSMath(address common.Address, backend bind.ContractBackend) (*DSMath, error) {
	contract, err := bindDSMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSMath{DSMathCaller: DSMathCaller{contract: contract}, DSMathTransactor: DSMathTransactor{contract: contract}}, nil
}

// NewDSMathCaller creates a new read-only instance of DSMath, bound to a specific deployed contract.
func NewDSMathCaller(address common.Address, caller bind.ContractCaller) (*DSMathCaller, error) {
	contract, err := bindDSMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSMathCaller{contract: contract}, nil
}

// NewDSMathTransactor creates a new write-only instance of DSMath, bound to a specific deployed contract.
func NewDSMathTransactor(address common.Address, transactor bind.ContractTransactor) (*DSMathTransactor, error) {
	contract, err := bindDSMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSMathTransactor{contract: contract}, nil
}

// bindDSMath binds a generic wrapper to an already deployed contract.
func bindDSMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSMath *DSMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSMath.Contract.DSMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSMath *DSMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSMath.Contract.DSMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSMath *DSMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSMath.Contract.DSMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSMath *DSMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSMath *DSMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSMath *DSMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSMath.Contract.contract.Transact(opts, method, params...)
}

// DSNoteABI is the input ABI used to generate the binding from.
const DSNoteABI = "[{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// DSNoteBin is the compiled bytecode used for deploying new contracts.
const DSNoteBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058207d9fa5ffcc41fc49cbfdac972a9325355e7602566bf02f7898ebe1ce08be92090029`

// DeployDSNote deploys a new Ethereum contract, binding an instance of DSNote to it.
func DeployDSNote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSNote, error) {
	parsed, err := abi.JSON(strings.NewReader(DSNoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSNoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSNote{DSNoteCaller: DSNoteCaller{contract: contract}, DSNoteTransactor: DSNoteTransactor{contract: contract}}, nil
}

// DSNote is an auto generated Go binding around an Ethereum contract.
type DSNote struct {
	DSNoteCaller     // Read-only binding to the contract
	DSNoteTransactor // Write-only binding to the contract
}

// DSNoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSNoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSNoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSNoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSNoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSNoteSession struct {
	Contract     *DSNote           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSNoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSNoteCallerSession struct {
	Contract *DSNoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSNoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSNoteTransactorSession struct {
	Contract     *DSNoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSNoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSNoteRaw struct {
	Contract *DSNote // Generic contract binding to access the raw methods on
}

// DSNoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSNoteCallerRaw struct {
	Contract *DSNoteCaller // Generic read-only contract binding to access the raw methods on
}

// DSNoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSNoteTransactorRaw struct {
	Contract *DSNoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSNote creates a new instance of DSNote, bound to a specific deployed contract.
func NewDSNote(address common.Address, backend bind.ContractBackend) (*DSNote, error) {
	contract, err := bindDSNote(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSNote{DSNoteCaller: DSNoteCaller{contract: contract}, DSNoteTransactor: DSNoteTransactor{contract: contract}}, nil
}

// NewDSNoteCaller creates a new read-only instance of DSNote, bound to a specific deployed contract.
func NewDSNoteCaller(address common.Address, caller bind.ContractCaller) (*DSNoteCaller, error) {
	contract, err := bindDSNote(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSNoteCaller{contract: contract}, nil
}

// NewDSNoteTransactor creates a new write-only instance of DSNote, bound to a specific deployed contract.
func NewDSNoteTransactor(address common.Address, transactor bind.ContractTransactor) (*DSNoteTransactor, error) {
	contract, err := bindDSNote(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSNoteTransactor{contract: contract}, nil
}

// bindDSNote binds a generic wrapper to an already deployed contract.
func bindDSNote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSNoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSNote *DSNoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSNote.Contract.DSNoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSNote *DSNoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSNote.Contract.DSNoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSNote *DSNoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSNote.Contract.DSNoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSNote *DSNoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSNote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSNote *DSNoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSNote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSNote *DSNoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSNote.Contract.contract.Transact(opts, method, params...)
}

// DSStopABI is the input ABI used to generate the binding from.
const DSStopABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// DSStopBin is the compiled bytecode used for deploying new contracts.
const DSStopBin = `0x6060604090815260018054600160a060020a03191633600160a060020a0316908117909155907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a261050d8061005f6000396000f3006060604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307da68f5811461008757806313af40351461009c57806375f12b21146100bb5780637a9e5e4b146100e25780638da5cb5b14610101578063be9a655514610130578063bf7e214f14610143575b600080fd5b341561009257600080fd5b61009a610156565b005b34156100a757600080fd5b61009a600160a060020a0360043516610203565b34156100c657600080fd5b6100ce61027a565b604051901515815260200160405180910390f35b34156100ed57600080fd5b61009a600160a060020a036004351661029b565b341561010c57600080fd5b610114610312565b604051600160a060020a03909116815260200160405180910390f35b341561013b57600080fd5b61009a610321565b341561014e57600080fd5b6101146103b2565b61017461016f33600035600160e060020a0319166103c1565b6104d2565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506001805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b61021c61016f33600035600160e060020a0319166103c1565b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b60015474010000000000000000000000000000000000000000900460ff1681565b6102b461016f33600035600160e060020a0319166103c1565b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600154600160a060020a031681565b61033a61016f33600035600160e060020a0319166103c1565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506001805474ff000000000000000000000000000000000000000019169055565b600054600160a060020a031681565b600030600160a060020a031683600160a060020a031614156103e5575060016104cc565b600154600160a060020a0384811691161415610403575060016104cc565b600054600160a060020a0316151561041d575060006104cc565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b15156104af57600080fd5b6102c65a03f115156104c057600080fd5b50505060405180519150505b92915050565b8015156104de57600080fd5b505600a165627a7a72305820a317e4cca6cf713c18fdb204966acae3fe569223aa733f3fa113f5fdfeaa0bf30029`

// DeployDSStop deploys a new Ethereum contract, binding an instance of DSStop to it.
func DeployDSStop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSStop, error) {
	parsed, err := abi.JSON(strings.NewReader(DSStopABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSStopBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSStop{DSStopCaller: DSStopCaller{contract: contract}, DSStopTransactor: DSStopTransactor{contract: contract}}, nil
}

// DSStop is an auto generated Go binding around an Ethereum contract.
type DSStop struct {
	DSStopCaller     // Read-only binding to the contract
	DSStopTransactor // Write-only binding to the contract
}

// DSStopCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSStopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSStopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSStopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSStopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSStopSession struct {
	Contract     *DSStop           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSStopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSStopCallerSession struct {
	Contract *DSStopCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSStopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSStopTransactorSession struct {
	Contract     *DSStopTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSStopRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSStopRaw struct {
	Contract *DSStop // Generic contract binding to access the raw methods on
}

// DSStopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSStopCallerRaw struct {
	Contract *DSStopCaller // Generic read-only contract binding to access the raw methods on
}

// DSStopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSStopTransactorRaw struct {
	Contract *DSStopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSStop creates a new instance of DSStop, bound to a specific deployed contract.
func NewDSStop(address common.Address, backend bind.ContractBackend) (*DSStop, error) {
	contract, err := bindDSStop(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSStop{DSStopCaller: DSStopCaller{contract: contract}, DSStopTransactor: DSStopTransactor{contract: contract}}, nil
}

// NewDSStopCaller creates a new read-only instance of DSStop, bound to a specific deployed contract.
func NewDSStopCaller(address common.Address, caller bind.ContractCaller) (*DSStopCaller, error) {
	contract, err := bindDSStop(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSStopCaller{contract: contract}, nil
}

// NewDSStopTransactor creates a new write-only instance of DSStop, bound to a specific deployed contract.
func NewDSStopTransactor(address common.Address, transactor bind.ContractTransactor) (*DSStopTransactor, error) {
	contract, err := bindDSStop(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSStopTransactor{contract: contract}, nil
}

// bindDSStop binds a generic wrapper to an already deployed contract.
func bindDSStop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSStopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSStop *DSStopRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSStop.Contract.DSStopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSStop *DSStopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.Contract.DSStopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSStop *DSStopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSStop.Contract.DSStopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSStop *DSStopCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSStop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSStop *DSStopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSStop *DSStopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSStop.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSStop *DSStopCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSStop.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSStop *DSStopSession) Authority() (common.Address, error) {
	return _DSStop.Contract.Authority(&_DSStop.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSStop *DSStopCallerSession) Authority() (common.Address, error) {
	return _DSStop.Contract.Authority(&_DSStop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSStop *DSStopCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSStop.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSStop *DSStopSession) Owner() (common.Address, error) {
	return _DSStop.Contract.Owner(&_DSStop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSStop *DSStopCallerSession) Owner() (common.Address, error) {
	return _DSStop.Contract.Owner(&_DSStop.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSStop *DSStopCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSStop.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSStop *DSStopSession) Stopped() (bool, error) {
	return _DSStop.Contract.Stopped(&_DSStop.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSStop *DSStopCallerSession) Stopped() (bool, error) {
	return _DSStop.Contract.Stopped(&_DSStop.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSStop *DSStopTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSStop *DSStopSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetAuthority(&_DSStop.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSStop *DSStopTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetAuthority(&_DSStop.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSStop *DSStopTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSStop *DSStopSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetOwner(&_DSStop.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSStop *DSStopTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetOwner(&_DSStop.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSStop *DSStopTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSStop *DSStopSession) Start() (*types.Transaction, error) {
	return _DSStop.Contract.Start(&_DSStop.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSStop *DSStopTransactorSession) Start() (*types.Transaction, error) {
	return _DSStop.Contract.Start(&_DSStop.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSStop *DSStopTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSStop *DSStopSession) Stop() (*types.Transaction, error) {
	return _DSStop.Contract.Stop(&_DSStop.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSStop *DSStopTransactorSession) Stop() (*types.Transaction, error) {
	return _DSStop.Contract.Stop(&_DSStop.TransactOpts)
}

// DSTokenABI is the input ABI used to generate the binding from.
const DSTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"push\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// DSTokenBin is the compiled bytecode used for deploying new contracts.
const DSTokenBin = `0x606060405260126006556000600755341561001957600080fd5b604051602080610e8783398101604052808051600160a060020a03331660008181526001602052604080822082905590805560048054600160a060020a0319168317905591935091507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a2600555610de88061009f6000396000f30060606040526004361061011c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461012157806307da68f514610146578063095ea7b31461015b57806313af40351461019157806318160ddd146101b057806323b872dd146101c3578063313ce567146101eb5780633452f51d146101fe5780635ac801fe1461022957806369d3e20e1461023f57806370a082311461025e57806375f12b211461027d5780637a9e5e4b146102905780638402181f146102af5780638da5cb5b146102da57806390bc16931461030957806395d89b4114610328578063a9059cbb1461033b578063be9a65551461035d578063bf7e214f14610370578063dd62ed3e14610383575b600080fd5b341561012c57600080fd5b6101346103a8565b60405190815260200160405180910390f35b341561015157600080fd5b6101596103ae565b005b341561016657600080fd5b61017d600160a060020a036004351660243561044a565b604051901515815260200160405180910390f35b341561019c57600080fd5b610159600160a060020a03600435166104cd565b34156101bb57600080fd5b610134610544565b34156101ce57600080fd5b61017d600160a060020a036004358116906024351660443561054a565b34156101f657600080fd5b6101346105cf565b341561020957600080fd5b61017d600160a060020a03600435166001608060020a03602435166105d5565b341561023457600080fd5b6101596004356105f3565b341561024a57600080fd5b6101596001608060020a0360043516610611565b341561026957600080fd5b610134600160a060020a03600435166106fb565b341561028857600080fd5b61017d610716565b341561029b57600080fd5b610159600160a060020a0360043516610726565b34156102ba57600080fd5b61017d600160a060020a03600435166001608060020a036024351661079d565b34156102e557600080fd5b6102ed6107b3565b604051600160a060020a03909116815260200160405180910390f35b341561031457600080fd5b6101596001608060020a03600435166107c2565b341561033357600080fd5b6101346108a4565b341561034657600080fd5b61017d600160a060020a03600435166024356108aa565b341561036857600080fd5b610159610924565b341561037b57600080fd5b6102ed6109b5565b341561038e57600080fd5b610134600160a060020a03600435811690602435166109c4565b60075481565b6103cc6103c733600035600160e060020a0319166109ef565b610afb565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff0000000000000000000000000000000000000000191660a060020a179055565b6004546000906104649060a060020a900460ff1615610afb565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46104c48585610b0a565b95945050505050565b6104e66103c733600035600160e060020a0319166109ef565b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b60005490565b6004546000906105649060a060020a900460ff1615610afb565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46105c5868686610b76565b9695505050505050565b60065481565b60006105ea83836001608060020a03166108aa565b90505b92915050565b61060c6103c733600035600160e060020a0319166109ef565b600755565b61062a6103c733600035600160e060020a0319166109ef565b6004546106419060a060020a900460ff1615610afb565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600160a060020a0333166000908152600160205260409020546106c3906001608060020a038516610ccc565b600160a060020a033316600090815260016020526040812091909155546106f3906001608060020a038516610ccc565b600055505050565b600160a060020a031660009081526001602052604090205490565b60045460a060020a900460ff1681565b61073f6103c733600035600160e060020a0319166109ef565b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b60006105ea8333846001608060020a031661054a565b600454600160a060020a031681565b6107db6103c733600035600160e060020a0319166109ef565b6004546107f29060a060020a900460ff1615610afb565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600160a060020a033316600090815260016020526040902054610874906001608060020a038516610cd9565b600160a060020a033316600090815260016020526040812091909155546106f3906001608060020a038516610cd9565b60055481565b6004546000906108c49060a060020a900460ff1615610afb565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46104c48585610ce6565b61093d6103c733600035600160e060020a0319166109ef565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff000000000000000000000000000000000000000019169055565b600354600160a060020a031681565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600030600160a060020a031683600160a060020a03161415610a13575060016105ed565b600454600160a060020a0384811691161415610a31575060016105ed565b600354600160a060020a03161515610a4b575060006105ed565b600354600160a060020a031663b70096138430856000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b1515610ad957600080fd5b6102c65a03f11515610aea57600080fd5b5050506040518051905090506105ed565b801515610b0757600080fd5b50565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b600160a060020a03831660009081526001602052604081205482901015610b9957fe5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482901015610bca57fe5b600160a060020a0380851660009081526002602090815260408083203390941683529290522054610bfb9083610cd9565b600160a060020a038086166000818152600260209081526040808320339095168352938152838220949094559081526001909252902054610c3c9083610cd9565b600160a060020a038086166000908152600160205260408082209390935590851681522054610c6b9083610ccc565b600160a060020a03808516600081815260016020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b808201828110156105ed57fe5b808203828111156105ed57fe5b600160a060020a03331660009081526001602052604081205482901015610d0957fe5b600160a060020a033316600090815260016020526040902054610d2c9083610cd9565b600160a060020a033381166000908152600160205260408082209390935590851681522054610d5b9083610ccc565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001929150505600a165627a7a723058205501255b0aee05a2eed8b6f0d09521208d0c804568252cd984f458030b1f7b770029`

// DeployDSToken deploys a new Ethereum contract, binding an instance of DSToken to it.
func DeployDSToken(auth *bind.TransactOpts, backend bind.ContractBackend, symbol_ [32]byte) (common.Address, *types.Transaction, *DSToken, error) {
	parsed, err := abi.JSON(strings.NewReader(DSTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSTokenBin), backend, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSToken{DSTokenCaller: DSTokenCaller{contract: contract}, DSTokenTransactor: DSTokenTransactor{contract: contract}}, nil
}

// DSToken is an auto generated Go binding around an Ethereum contract.
type DSToken struct {
	DSTokenCaller     // Read-only binding to the contract
	DSTokenTransactor // Write-only binding to the contract
}

// DSTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSTokenSession struct {
	Contract     *DSToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSTokenCallerSession struct {
	Contract *DSTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DSTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSTokenTransactorSession struct {
	Contract     *DSTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DSTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSTokenRaw struct {
	Contract *DSToken // Generic contract binding to access the raw methods on
}

// DSTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSTokenCallerRaw struct {
	Contract *DSTokenCaller // Generic read-only contract binding to access the raw methods on
}

// DSTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSTokenTransactorRaw struct {
	Contract *DSTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSToken creates a new instance of DSToken, bound to a specific deployed contract.
func NewDSToken(address common.Address, backend bind.ContractBackend) (*DSToken, error) {
	contract, err := bindDSToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSToken{DSTokenCaller: DSTokenCaller{contract: contract}, DSTokenTransactor: DSTokenTransactor{contract: contract}}, nil
}

// NewDSTokenCaller creates a new read-only instance of DSToken, bound to a specific deployed contract.
func NewDSTokenCaller(address common.Address, caller bind.ContractCaller) (*DSTokenCaller, error) {
	contract, err := bindDSToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSTokenCaller{contract: contract}, nil
}

// NewDSTokenTransactor creates a new write-only instance of DSToken, bound to a specific deployed contract.
func NewDSTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DSTokenTransactor, error) {
	contract, err := bindDSToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSTokenTransactor{contract: contract}, nil
}

// bindDSToken binds a generic wrapper to an already deployed contract.
func bindDSToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSToken *DSTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSToken.Contract.DSTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSToken *DSTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSToken.Contract.DSTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSToken *DSTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSToken.Contract.DSTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSToken *DSTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSToken *DSTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSToken *DSTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSToken *DSTokenCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSToken *DSTokenSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _DSToken.Contract.Allowance(&_DSToken.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSToken *DSTokenCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _DSToken.Contract.Allowance(&_DSToken.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSToken *DSTokenCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSToken *DSTokenSession) Authority() (common.Address, error) {
	return _DSToken.Contract.Authority(&_DSToken.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSToken *DSTokenCallerSession) Authority() (common.Address, error) {
	return _DSToken.Contract.Authority(&_DSToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSToken *DSTokenCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSToken *DSTokenSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _DSToken.Contract.BalanceOf(&_DSToken.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSToken *DSTokenCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _DSToken.Contract.BalanceOf(&_DSToken.CallOpts, src)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_DSToken *DSTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_DSToken *DSTokenSession) Decimals() (*big.Int, error) {
	return _DSToken.Contract.Decimals(&_DSToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_DSToken *DSTokenCallerSession) Decimals() (*big.Int, error) {
	return _DSToken.Contract.Decimals(&_DSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_DSToken *DSTokenCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_DSToken *DSTokenSession) Name() ([32]byte, error) {
	return _DSToken.Contract.Name(&_DSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_DSToken *DSTokenCallerSession) Name() ([32]byte, error) {
	return _DSToken.Contract.Name(&_DSToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSToken *DSTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSToken *DSTokenSession) Owner() (common.Address, error) {
	return _DSToken.Contract.Owner(&_DSToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSToken *DSTokenCallerSession) Owner() (common.Address, error) {
	return _DSToken.Contract.Owner(&_DSToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSToken *DSTokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSToken *DSTokenSession) Stopped() (bool, error) {
	return _DSToken.Contract.Stopped(&_DSToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSToken *DSTokenCallerSession) Stopped() (bool, error) {
	return _DSToken.Contract.Stopped(&_DSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_DSToken *DSTokenCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_DSToken *DSTokenSession) Symbol() ([32]byte, error) {
	return _DSToken.Contract.Symbol(&_DSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_DSToken *DSTokenCallerSession) Symbol() ([32]byte, error) {
	return _DSToken.Contract.Symbol(&_DSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSToken *DSTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSToken *DSTokenSession) TotalSupply() (*big.Int, error) {
	return _DSToken.Contract.TotalSupply(&_DSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSToken *DSTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DSToken.Contract.TotalSupply(&_DSToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSToken *DSTokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSToken *DSTokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Approve(&_DSToken.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSToken *DSTokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Approve(&_DSToken.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_DSToken *DSTokenTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_DSToken *DSTokenSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Burn(&_DSToken.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_DSToken *DSTokenTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Burn(&_DSToken.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_DSToken *DSTokenTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_DSToken *DSTokenSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Mint(&_DSToken.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_DSToken *DSTokenTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Mint(&_DSToken.TransactOpts, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_DSToken *DSTokenTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_DSToken *DSTokenSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Pull(&_DSToken.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_DSToken *DSTokenTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Pull(&_DSToken.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_DSToken *DSTokenTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_DSToken *DSTokenSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Push(&_DSToken.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_DSToken *DSTokenTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Push(&_DSToken.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSToken *DSTokenTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSToken *DSTokenSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSToken.Contract.SetAuthority(&_DSToken.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSToken *DSTokenTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSToken.Contract.SetAuthority(&_DSToken.TransactOpts, authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(name_ bytes32) returns()
func (_DSToken *DSTokenTransactor) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(name_ bytes32) returns()
func (_DSToken *DSTokenSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _DSToken.Contract.SetName(&_DSToken.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(name_ bytes32) returns()
func (_DSToken *DSTokenTransactorSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _DSToken.Contract.SetName(&_DSToken.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSToken *DSTokenTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSToken *DSTokenSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSToken.Contract.SetOwner(&_DSToken.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSToken *DSTokenTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSToken.Contract.SetOwner(&_DSToken.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSToken *DSTokenTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSToken *DSTokenSession) Start() (*types.Transaction, error) {
	return _DSToken.Contract.Start(&_DSToken.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSToken *DSTokenTransactorSession) Start() (*types.Transaction, error) {
	return _DSToken.Contract.Start(&_DSToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSToken *DSTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSToken *DSTokenSession) Stop() (*types.Transaction, error) {
	return _DSToken.Contract.Stop(&_DSToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSToken *DSTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _DSToken.Contract.Stop(&_DSToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSToken *DSTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSToken *DSTokenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Transfer(&_DSToken.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSToken *DSTokenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.Transfer(&_DSToken.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSToken *DSTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSToken *DSTokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.TransferFrom(&_DSToken.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSToken *DSTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSToken.Contract.TransferFrom(&_DSToken.TransactOpts, src, dst, wad)
}

// DSTokenBaseABI is the input ABI used to generate the binding from.
const DSTokenBaseABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// DSTokenBaseBin is the compiled bytecode used for deploying new contracts.
const DSTokenBaseBin = `0x6060604052341561000f57600080fd5b6040516020806104df83398101604052808051600160a060020a03331660009081526001602052604081208290555550506104908061004f6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100b257806323b872dd146100d757806370a08231146100ff578063a9059cbb1461011e578063dd62ed3e14610140575b600080fd5b341561008757600080fd5b61009e600160a060020a0360043516602435610165565b604051901515815260200160405180910390f35b34156100bd57600080fd5b6100c56101d2565b60405190815260200160405180910390f35b34156100e257600080fd5b61009e600160a060020a03600435811690602435166044356101d8565b341561010a57600080fd5b6100c5600160a060020a036004351661032e565b341561012957600080fd5b61009e600160a060020a0360043516602435610349565b341561014b57600080fd5b6100c5600160a060020a036004358116906024351661041f565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005490565b600160a060020a038316600090815260016020526040812054829010156101fb57fe5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548290101561022c57fe5b600160a060020a038085166000908152600260209081526040808320339094168352929052205461025d908361044a565b600160a060020a03808616600081815260026020908152604080832033909516835293815283822094909455908152600190925290205461029e908361044a565b600160a060020a0380861660009081526001602052604080822093909355908516815220546102cd9083610457565b600160a060020a03808516600081815260016020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600160a060020a0333166000908152600160205260408120548290101561036c57fe5b600160a060020a03331660009081526001602052604090205461038f908361044a565b600160a060020a0333811660009081526001602052604080822093909355908516815220546103be9083610457565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b808203828111156101cc57fe5b808201828110156101cc57fe00a165627a7a72305820e8c5b0bad73c82b88667142daffd95695c45d3016241b98f66ab256cec8f9de30029`

// DeployDSTokenBase deploys a new Ethereum contract, binding an instance of DSTokenBase to it.
func DeployDSTokenBase(auth *bind.TransactOpts, backend bind.ContractBackend, supply *big.Int) (common.Address, *types.Transaction, *DSTokenBase, error) {
	parsed, err := abi.JSON(strings.NewReader(DSTokenBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSTokenBaseBin), backend, supply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSTokenBase{DSTokenBaseCaller: DSTokenBaseCaller{contract: contract}, DSTokenBaseTransactor: DSTokenBaseTransactor{contract: contract}}, nil
}

// DSTokenBase is an auto generated Go binding around an Ethereum contract.
type DSTokenBase struct {
	DSTokenBaseCaller     // Read-only binding to the contract
	DSTokenBaseTransactor // Write-only binding to the contract
}

// DSTokenBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSTokenBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSTokenBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSTokenBaseSession struct {
	Contract     *DSTokenBase      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSTokenBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSTokenBaseCallerSession struct {
	Contract *DSTokenBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DSTokenBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSTokenBaseTransactorSession struct {
	Contract     *DSTokenBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DSTokenBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSTokenBaseRaw struct {
	Contract *DSTokenBase // Generic contract binding to access the raw methods on
}

// DSTokenBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSTokenBaseCallerRaw struct {
	Contract *DSTokenBaseCaller // Generic read-only contract binding to access the raw methods on
}

// DSTokenBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSTokenBaseTransactorRaw struct {
	Contract *DSTokenBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSTokenBase creates a new instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBase(address common.Address, backend bind.ContractBackend) (*DSTokenBase, error) {
	contract, err := bindDSTokenBase(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSTokenBase{DSTokenBaseCaller: DSTokenBaseCaller{contract: contract}, DSTokenBaseTransactor: DSTokenBaseTransactor{contract: contract}}, nil
}

// NewDSTokenBaseCaller creates a new read-only instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBaseCaller(address common.Address, caller bind.ContractCaller) (*DSTokenBaseCaller, error) {
	contract, err := bindDSTokenBase(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseCaller{contract: contract}, nil
}

// NewDSTokenBaseTransactor creates a new write-only instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*DSTokenBaseTransactor, error) {
	contract, err := bindDSTokenBase(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseTransactor{contract: contract}, nil
}

// bindDSTokenBase binds a generic wrapper to an already deployed contract.
func bindDSTokenBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSTokenBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSTokenBase *DSTokenBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSTokenBase.Contract.DSTokenBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSTokenBase *DSTokenBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSTokenBase.Contract.DSTokenBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSTokenBase *DSTokenBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSTokenBase.Contract.DSTokenBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSTokenBase *DSTokenBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSTokenBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSTokenBase *DSTokenBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSTokenBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSTokenBase *DSTokenBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSTokenBase.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSTokenBase.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.Allowance(&_DSTokenBase.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.Allowance(&_DSTokenBase.CallOpts, src, guy)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSTokenBase.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.BalanceOf(&_DSTokenBase.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.BalanceOf(&_DSTokenBase.CallOpts, src)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSTokenBase.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSTokenBase *DSTokenBaseSession) TotalSupply() (*big.Int, error) {
	return _DSTokenBase.Contract.TotalSupply(&_DSTokenBase.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCallerSession) TotalSupply() (*big.Int, error) {
	return _DSTokenBase.Contract.TotalSupply(&_DSTokenBase.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Approve(&_DSTokenBase.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Approve(&_DSTokenBase.TransactOpts, guy, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Transfer(&_DSTokenBase.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Transfer(&_DSTokenBase.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.TransferFrom(&_DSTokenBase.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.TransferFrom(&_DSTokenBase.TransactOpts, src, dst, wad)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}
