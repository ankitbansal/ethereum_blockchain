// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RatingABI is the input ABI used to generate the binding from.
const RatingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rating\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"providerName\",\"type\":\"string\"},{\"name\":\"customerName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"providerName\",\"type\":\"string\"},{\"name\":\"customerName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// RatingBin is the compiled bytecode used for deploying new contracts.
const RatingBin = `0x608060405234801561001057600080fd5b506040516103cd3803806103cd83398101604081815282516020808501518386015160608601855283865290860182860181905295019284018390526000828155855192959492909161006891600191870190610090565b5060408201518051610084916002840191602090910190610090565b5090505050505061012b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100d157805160ff19168380011785556100fe565b828001600101855582156100fe579182015b828111156100fe5782518255916020019190600101906100e3565b5061010a92915061010e565b5090565b61012891905b8082111561010a5760008155600101610114565b90565b6102938061013a6000396000f3006080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166380f474f68114610045575b600080fd5b34801561005157600080fd5b5061005a61013f565b604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156100a1578181015183820152602001610089565b50505050905090810190601f1680156100ce5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156101015781810151838201526020016100e9565b50505050905090810190601f16801561012e5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6000805460018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815293949392918301828280156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b50505060028085018054604080516020601f600019610100600187161502019094169590950492830185900485028101850190915281815295969594509092509083018282801561025d5780601f106102325761010080835404028352916020019161025d565b820191906000526020600020905b81548152906001019060200180831161024057829003601f168201915b50505050509050835600a165627a7a7230582089ab44d723fbca774758dc62a05354aa8b3e8e5e18c84e3c298bcfba22f38b560029`

// DeployRating deploys a new Ethereum contract, binding an instance of Rating to it.
func DeployRating(auth *bind.TransactOpts, backend bind.ContractBackend, value *big.Int, providerName string, customerName string) (common.Address, *types.Transaction, *Rating, error) {
	parsed, err := abi.JSON(strings.NewReader(RatingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RatingBin), backend, value, providerName, customerName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rating{RatingCaller: RatingCaller{contract: contract}, RatingTransactor: RatingTransactor{contract: contract}, RatingFilterer: RatingFilterer{contract: contract}}, nil
}

// Rating is an auto generated Go binding around an Ethereum contract.
type Rating struct {
	RatingCaller     // Read-only binding to the contract
	RatingTransactor // Write-only binding to the contract
	RatingFilterer   // Log filterer for contract events
}

// RatingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RatingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RatingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RatingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RatingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RatingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RatingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RatingSession struct {
	Contract     *Rating           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RatingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RatingCallerSession struct {
	Contract *RatingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RatingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RatingTransactorSession struct {
	Contract     *RatingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RatingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RatingRaw struct {
	Contract *Rating // Generic contract binding to access the raw methods on
}

// RatingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RatingCallerRaw struct {
	Contract *RatingCaller // Generic read-only contract binding to access the raw methods on
}

// RatingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RatingTransactorRaw struct {
	Contract *RatingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRating creates a new instance of Rating, bound to a specific deployed contract.
func NewRating(address common.Address, backend bind.ContractBackend) (*Rating, error) {
	contract, err := bindRating(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rating{RatingCaller: RatingCaller{contract: contract}, RatingTransactor: RatingTransactor{contract: contract}, RatingFilterer: RatingFilterer{contract: contract}}, nil
}

// NewRatingCaller creates a new read-only instance of Rating, bound to a specific deployed contract.
func NewRatingCaller(address common.Address, caller bind.ContractCaller) (*RatingCaller, error) {
	contract, err := bindRating(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RatingCaller{contract: contract}, nil
}

// NewRatingTransactor creates a new write-only instance of Rating, bound to a specific deployed contract.
func NewRatingTransactor(address common.Address, transactor bind.ContractTransactor) (*RatingTransactor, error) {
	contract, err := bindRating(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RatingTransactor{contract: contract}, nil
}

// NewRatingFilterer creates a new log filterer instance of Rating, bound to a specific deployed contract.
func NewRatingFilterer(address common.Address, filterer bind.ContractFilterer) (*RatingFilterer, error) {
	contract, err := bindRating(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RatingFilterer{contract: contract}, nil
}

// bindRating binds a generic wrapper to an already deployed contract.
func bindRating(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RatingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rating *RatingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rating.Contract.RatingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rating *RatingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rating.Contract.RatingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rating *RatingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rating.Contract.RatingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rating *RatingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rating.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rating *RatingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rating.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rating *RatingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rating.Contract.contract.Transact(opts, method, params...)
}

// Rating is a free data retrieval call binding the contract method 0x80f474f6.
//
// Solidity: function rating() constant returns(value uint256, providerName string, customerName string)
func (_Rating *RatingCaller) Rating(opts *bind.CallOpts) (struct {
	Value        *big.Int
	ProviderName string
	CustomerName string
}, error) {
	ret := new(struct {
		Value        *big.Int
		ProviderName string
		CustomerName string
	})
	out := ret
	err := _Rating.contract.Call(opts, out, "rating")
	return *ret, err
}

// Rating is a free data retrieval call binding the contract method 0x80f474f6.
//
// Solidity: function rating() constant returns(value uint256, providerName string, customerName string)
func (_Rating *RatingSession) Rating() (struct {
	Value        *big.Int
	ProviderName string
	CustomerName string
}, error) {
	return _Rating.Contract.Rating(&_Rating.CallOpts)
}

// Rating is a free data retrieval call binding the contract method 0x80f474f6.
//
// Solidity: function rating() constant returns(value uint256, providerName string, customerName string)
func (_Rating *RatingCallerSession) Rating() (struct {
	Value        *big.Int
	ProviderName string
	CustomerName string
}, error) {
	return _Rating.Contract.Rating(&_Rating.CallOpts)
}
