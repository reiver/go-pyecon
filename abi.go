package pyecon

import (
	_ "embed"
	"io"
	"strings"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"sourcecode.social/reiver/go-erorr"
)

//go:embed abi.json
var abijson string

var abi ethabi.ABI

func init () {
	var reader io.Reader = strings.NewReader(abijson)

	var err error

	abi, err = ethabi.JSON(reader)
	if nil != err {
		var e error = erorr.Errorf("pyecon: problem parsing pyecon contract JSON ABI: %w", err)
		panic(e)
	}
}
