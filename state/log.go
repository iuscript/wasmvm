package state

import (
	"github.com/iuscript/wasmvm/common"
)

type Log struct {
	Address     common.Address
	Topics      []common.Hash
	Data        []byte
	BlockHeight uint64
}
