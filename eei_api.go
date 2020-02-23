package wasmvm

import (
	"github.com/iuscript/wasmvm/wagon/exec"
)

const (
	// EEICallSuccess is the return value in case of a successful contract execution
	EEICallSuccess = 0
	// ErrEEICallFailure is the return value in case of a contract execution failture
	ErrEEICallFailure = 1
	// ErrEEICallRevert is the return value in case a contract calls `revert`
	ErrEEICallRevert = 2
)

type eeiApi struct{}

func (*eeiApi) getAddress(p *exec.Process, w *WasmIntptr, resultOffset int32) {
	writeToMem(p, w.contract.Address().Bytes(), resultOffset)
}

func (*eeiApi) getCaller(p *exec.Process, w *WasmIntptr, resultOffset int32) {

	addr := w.contract.CallerAddress
	writeToMem(p, addr.Bytes(), resultOffset)
}

func (*eeiApi) getCallValue(p *exec.Process, w *WasmIntptr, resultOffset int32) {

	writeToMem(p, w.contract.Value().Bytes(), resultOffset)
}

func (*eeiApi) codeCopy(p *exec.Process, w *WasmIntptr, resultOffset, codeOffset, length int32) {

	writeToMem(p, w.contract.Code[codeOffset:codeOffset+length], resultOffset)
}

func (*eeiApi) getCodeSize(p *exec.Process, w *WasmIntptr) int32 {

	return int32(len(w.contract.Code))
}

func (*eeiApi) getBlockCoinbase(p *exec.Process, w *WasmIntptr, resultOffset int32) {

	writeToMem(p, w.evm.Coinbase().Bytes(), resultOffset)
}

func (*eeiApi) getBlockDifficulty(p *exec.Process, w *WasmIntptr, resultOffset int32) {

}

func (*eeiApi) finish(p *exec.Process, w *WasmIntptr, dataOffset, length int32) {
	w.returnData = loadFromMem(p, dataOffset, length)
	w.terminateType = TerminateFinish
	p.Terminate()
}

func (*eeiApi) revert(p *exec.Process, w *WasmIntptr, dataOffset, length int32) {
	w.returnData = loadFromMem(p, dataOffset, length)
	w.terminateType = TerminateRevert
	p.Terminate()
}

func (*eeiApi) getReturnDataSize(p *exec.Process, w *WasmIntptr) int32 {

	return int32(len(w.returnData))
}

func (*eeiApi) returnDataCopy(p *exec.Process, w *WasmIntptr, resultOffset, dataOffset, length int32) {

	writeToMem(p, w.returnData[dataOffset:dataOffset+length], resultOffset)
}

// swapEndian swap big endian to little endian or reverse.
func swapEndian(src []byte) []byte {
	rect := make([]byte, len(src))
	for i, b := range src {
		rect[len(src)-i-1] = b
	}
	return rect
}

func loadFromMem(p *exec.Process, offset int32, size int32) []byte {
	b := make([]byte, size)
	p.ReadAt(b, int64(offset))
	return swapEndian(b)
}

func writeToMem(p *exec.Process, data []byte, offset int32) (int, error) {
	return p.WriteAt(swapEndian(data), int64(offset))
}
