package wasmvm

import (
	"reflect"

	"github.com/iuscript/wasmvm/wagon/wasm"
)

func goType2WasmType(kind reflect.Kind) wasm.ValueType {
	switch kind {
	case reflect.Int, reflect.Int32, reflect.Struct:
		return wasm.ValueTypeI32
	case reflect.Ptr, reflect.Uint, reflect.Int64, reflect.Uint64:
		return wasm.ValueTypeI64
	case reflect.Float32:
		return wasm.ValueTypeF32
	case reflect.Float64:
		return wasm.ValueTypeF64
	default:
		return wasm.ValueTypeI64
	}
}
