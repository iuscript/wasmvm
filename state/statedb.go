package state

import (
	"fmt"
	"io/ioutil"

	"github.com/iuscript/wasmvm/common"
)

type StateDB struct{}

func (*StateDB) CreateAccount(addr common.Address) {

}

func (*StateDB) GetCodeHash(addr common.Address) common.Hash {
	return common.Hash{}
}

func (*StateDB) GetCode(addr common.Address) []byte {
	code, err := ioutil.ReadFile("testdata/testI.wasm")
	if err != nil {
		fmt.Println(err)
	}
	return code
}

func (*StateDB) SetCode(addr common.Address, code []byte) {
	filename := addr.String()
	ioutil.WriteFile("testdata/"+filename, code, 0644)
}
