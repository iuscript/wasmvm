package wasmvm

import (
	"github.com/iuscript/wasmvm/common"
	"github.com/seaio-co/crypto/sha3"
)

func CreateAddress(callerAddr common.Address, codeHash common.Hash) common.Address {
	return common.BytesToAddress(sha3.Keccak256(callerAddr.Bytes(), codeHash.Bytes()))
}

func CodeHash(code []byte) common.Hash {
	return common.BytesToHash(sha3.Keccak256(code))
}
