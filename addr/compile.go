package addr

import (
	gohex "encoding/hex"
	"bytes"

	"sourcecode.social/reiver/go-erorr"
)

func compile(dst []byte, hex []byte) error {
	if nil == dst {
		return errNilDestinationByteSlice
	}

	{
		var prefix = [...]byte{'0','x'}

		if bytes.HasPrefix(hex, prefix[:]) {
			hex = hex[len(prefix):]
		}
	}

	{
		if 1 == len(hex) % 2 {
			hex = append([]byte{'0'}, hex...)
		}
	}

	{
		var lenDst   int = len(dst)
		var lenHex int = len(hex)

		if lenDst < (lenHex / 2) {
			return errDestinationByteSliceTooShort
		}
	}

	{
		_, err := gohex.Decode(dst, hex)
		if nil != err {
			return erorr.Errorf("pyecon: problem decoding hexadecial: %w", err)
		}
	}

	return nil
}

