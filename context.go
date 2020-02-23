package wasmvm

import (
	"math/big"

	"github.com/iuscript/wasmvm/common"
	"github.com/iuscript/wasmvm/state"
)

type (
	// CanTransferFunc is the signature of a transfer guard function
	CanTransferFunc func(state.StateDB, common.Address, *big.Int) bool
	// TransferFunc is the signature of a transfer function
	TransferFunc func(state.StateDB, common.Address, common.Address, *big.Int)
	// GetHashFunc returns the nth block hash in the blockchain
	// and is used by the BLOCKHASH EVM op code.
	GetHashFunc func() common.Hash
)

// Context provides the EVM with auxiliary information. Once provided
// it shouldn't be modified.
type Context struct {
	// CanTransfer returns whether the account contains
	// sufficient ether to transfer the value
	CanTransfer CanTransferFunc
	// Transfer transfers ether from one account to the other
	Transfer TransferFunc
	// GetHash returns the hash corresponding to n
	GetHash GetHashFunc

	// Message information
	Origin common.Address // Provides information for ORIGIN

	// Block information
	Coinbase    common.Address // Provides information for COINBASE
	BlockHeight *big.Int       // Provides information for HEIGHT
	Time        *big.Int       // Provides information for TIME

}

// NewEVMContext creates a new context for use in the EVM.
func NewEVMContext(from *common.Address, author *common.Address) Context {
	// If we don't have an explicit author (i.e. not mining), extract from the header
	var beneficiary common.Address

	beneficiary = *author

	return Context{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(),
		Origin:      *from,
		Coinbase:    beneficiary,
		BlockHeight: new(big.Int).SetUint64(45),
		//Difficulty:  new(big.Int).Set(header.Difficulty),

	}
}

// GetHashFn returns a GetHashFunc which retrieves header hashes by number
func GetHashFn() func() common.Hash {
	return func() common.Hash {
		return common.Hash{}
	}
}

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db state.StateDB, addr common.Address, amount *big.Int) bool {
	// return db.GetBalance(addr).Cmp(amount) >= 0
	return true
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db state.StateDB, sender, recipient common.Address, amount *big.Int) {
	// db.SubBalance(sender, amount)
	// db.AddBalance(recipient, amount)
}
