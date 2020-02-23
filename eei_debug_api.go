package wasmvm

import (
	"fmt"

	"github.com/iuscript/wasmvm/wagon/exec"
)

type eeiDebugApi struct{}

func (*eeiDebugApi) print32(p *exec.Process, w *WasmIntptr, value int32) {
	fmt.Println(value)
}

func (*eeiDebugApi) print64(p *exec.Process, w *WasmIntptr, value int64) {
	fmt.Println(value)
}

func (*eeiDebugApi) printMem(p *exec.Process, w *WasmIntptr, offset, len int32) {
	fmt.Println(loadFromMem(p, offset, len))
}
