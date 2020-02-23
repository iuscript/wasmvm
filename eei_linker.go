package wasmvm

import (
	"fmt"

	"github.com/iuscript/wasmvm/wagon/wasm"
)

func moduleResolver(w *WasmIntptr, name string) (*wasm.Module, error) {
	if name == "sea" {
		m := wasm.NewModule()
		m.Types.Entries = w.eeiFuncSet.entries
		m.FunctionIndexSpace = w.eeiFuncSet.funcs
		m.Export.Entries = w.eeiFuncSet.exports

		return m, nil
	}

	if w.debug() && name == "debug" {
		m := wasm.NewModule()
		m.Types.Entries = w.debugFuncSet.entries
		m.FunctionIndexSpace = w.debugFuncSet.funcs
		m.Export.Entries = w.debugFuncSet.exports

		return m, nil
	}

	return nil, fmt.Errorf("unknow module name %s", name)
}

func ModuleResolver(w *WasmIntptr) wasm.ResolveFunc {
	return func(name string) (*wasm.Module, error) {
		return moduleResolver(w, name)
	}
}
