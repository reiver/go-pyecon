package addr

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errDestinationByteSliceTooShort = erorr.Error("pyecon: destination byte-slice too short")
	errNilDestinationByteSlice      = erorr.Error("pyecon: nil destination byte-slice")
)
