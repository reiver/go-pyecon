package pyecon

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errMissingContractAddress = erorr.Error("pyecon: missing contract-address")
	errMissingRPCURL          = erorr.Error("pyecon: missing rpc-url")
	errNilCallData            = erorr.Error("pyecon: nil call-datao")
	errNilReceiver            = erorr.Error("pyecon: nil receiver")
	errNilRPCClient           = erorr.Error("pyecon: nil rpc-client")
)
