package wasmvm

import (
	"math/big"
	"os"
	"sync/atomic"

	"github.com/iuscript/wasmvm/common"
	"github.com/iuscript/wasmvm/state"
)

const maxCallDepth = 1024

const MaxCodeSize = 24576

// run runs the given contract and takes care of running precompiles with a fallback to the byte code interpreter.
func run(evm *EVM, contract *Contract, input []byte) ([]byte, error) {

	if evm.interpreter.CanRun(contract.Code) {

		return evm.interpreter.Run(contract, input)
	}
	return nil, ErrNoCompatibleInterpreter
}

// Config are the configuration options for the Interpreter
type Config struct {
	// Debug enabled debugging Interpreter options
	Debug bool
	// Tracer is the op code logger
	Tracer Tracer
	// NoRecursion disabled Interpreter call, callcode,
	// delegate call and create.
	NoRecursion bool
	// Enable recording of SHA3/keccak preimages
	EnablePreimageRecording bool
	// Type of the EWASM interpreter
	EWASMInterpreter string
	// Type of the EVM interpreter
	EVMInterpreter string
}

type Interpreter interface {
	Run(contract *Contract, input []byte) ([]byte, error)

	CanRun([]byte) bool

	IsReadOnly() bool

	SetReadOnly(bool)
}

type EVM struct {
	Context

	StateDB state.StateDB

	depth int

	vmConfig Config

	interpreter Interpreter

	abort int32
}

// NewEVM returns a new EVM. The returned EVM is not thread safe and should
// only ever be used *once*.
func NewEVM(ctx Context, StateDB state.StateDB, vmConfig Config) *EVM {
	evm := &EVM{
		Context:  ctx,
		StateDB:  StateDB,
		vmConfig: vmConfig,
	}

	evm.interpreter = NewWasmIntptr(evm)

	return evm
}

func (evm *EVM) DB() state.StateDB {
	return evm.StateDB
}

// Coinbase returns the address of block producer
func (evm *EVM) Coinbase() common.Address {
	return evm.Context.Coinbase
}

// Cancel cancels any running EVM operation. This may be called concurrently and
// it's safe to be called multiple times.
func (evm *EVM) Cancel() {
	atomic.StoreInt32(&evm.abort, 1)
}

// Interpreter returns the current interpreter
func (evm *EVM) Interpreter() Interpreter {
	return evm.interpreter
}

// Call executes the contract associated with the addr with the given input as
// parameters. It also handles any necessary value transfer required and takes
// the necessary steps to create accounts and reverses the state in case of an
// execution error or failed value transfer.
func (evm *EVM) Call(caller ContractRef, addr common.Address, input []byte, value *big.Int) (ret []byte, err error) {
	if evm.vmConfig.NoRecursion && evm.depth > 0 {
		return nil, ErrRecursion
	}

	// Fail if we're trying to execute above the call depth limit
	if evm.depth > int(maxCallDepth) {
		return nil, ErrDepth
	}
	// Fail if we're trying to transfer more than the available balance
	if !evm.Context.CanTransfer(evm.StateDB, caller.Address(), value) {
		return nil, ErrInsufficientBalance
	}

	var (
		to = AccountRef(addr)
	)

	// evm.Transfer(evm.StateDB, caller.Address(), to.Address(), value)

	// Initialise a new contract and set the code that is to be used by the EVM.
	// The contract is a scoped environment for this execution context only.
	contract := NewContract(caller, to, value)
	contract.SetCallCode(&addr, evm.StateDB.GetCodeHash(addr), evm.StateDB.GetCode(addr))

	ret, err = run(evm, contract, input)

	return ret, err
}

func (evm *EVM) Save(caller ContractRef, code []byte) (common.Address, error) {
	codeHash := CodeHash(code)
	contractAddress := CreateAddress(caller.Address(), codeHash)
	object := AccountRef(contractAddress)

	contract := NewContract(caller, object, big.NewInt(0))
	contract.SetCode(codeHash, code)

	evm.StateDB.SetCode(contractAddress, code)

	// fmt.Println(contract)
	os.Exit(0)

	return contractAddress, nil
}
