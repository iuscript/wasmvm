package main

import (
	"fmt"

	"github.com/iuscript/wasmvm"
	"github.com/iuscript/wasmvm/common"
	"github.com/iuscript/wasmvm/state"
)

func main() {
	from := common.HexToAddress("0x57e9b7576ff33fba1c059971292dfb968c58de4c")
	// to := common.HexToAddress("0x43365fbfc22d8eaf328cc58cea0ab2eec7a27550")
	author := common.HexToAddress("0x26a5d142602b6d94957a5e5918af8ce02337b4a4")
	caller := wasmvm.AccountRef(from)

	input := []byte(`main,12345`)

	ctx := wasmvm.NewEVMContext(&from, &author)

	std := state.StateDB{}

	vmconfig := wasmvm.Config{Debug: true}

	xvm := wasmvm.NewEVM(ctx, std, vmconfig)

	// ret, err := xvm.Call(caller, to, input, big.NewInt(0))
	ret, err := xvm.Save(caller, input)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
