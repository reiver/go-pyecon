package addr

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

// Address represents a address.
// Note that Address is an optional-type.
type Address struct {
	optional opt.Optional[[20]byte]
}

func Nothing() Address {
	return Address{}
}

func Something(value [20]byte) Address {
	return Address{
		optional:opt.Something(value),
	}
}

func (receiver Address) Get() ([20]byte, bool) {
	return receiver.optional.Get()
}

func (receiver Address) String() string {
	value, something := receiver.optional.Get()
	if !something {
		return ""
	}

	return fmt.Sprintf("0x%040x", value[:])
}
