package wasmvm

import (
	"math/big"
	"time"

	"github.com/iuscript/wasmvm/common"
)

type Tracer interface {
	CaptureStart(from common.Address, to common.Address, call bool, input []byte, gas uint64, value *big.Int) error
	CaptureState(env *EVM, pc uint64, cost uint64) error
	CaptureFault(env *EVM, pc uint64, cost uint64) error
	CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error
}
