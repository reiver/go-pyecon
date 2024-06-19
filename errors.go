package pyecon

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errMaturityOverflow       = erorr.Error("pyecon: 'maturity' overflow")
	errMissingContractAddress = erorr.Error("pyecon: missing contract-address")
	errMissingRPCURL          = erorr.Error("pyecon: missing rpc-url")
	errNilCallData            = erorr.Error("pyecon: nil call-datao")
	errNilLog                 = erorr.Error("pyecon: nil log")
	errNilMaturity            = erorr.Error("pyecon: nil 'maturity'")
	errNilReceiver            = erorr.Error("pyecon: nil receiver")
	errNilRPCClient           = erorr.Error("pyecon: nil rpc-client")
)
