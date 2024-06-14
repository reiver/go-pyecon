package tokenid

import (
	"math/big"

	"sourcecode.social/reiver/go-opt"
)

// TokenID represents a token-id.
// Note that TokenID is an optional-type.
type TokenID struct {
	optional opt.Optional[[32]byte]
}

func Nothing() TokenID {
	return TokenID{}
}

func Uint64(value uint64) TokenID {
	return TokenID{
		optional:opt.Something([32]byte{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,

			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,

			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,

			byte((value & 0xff00000000000000) >> (7 * 8)),
			byte((value & 0x00ff000000000000) >> (6 * 8)),
			byte((value & 0x0000ff0000000000) >> (5 * 8)),
			byte((value & 0x000000ff00000000) >> (4 * 8)),
			byte((value & 0x00000000ff000000) >> (3 * 8)),
			byte((value & 0x0000000000ff0000) >> (2 * 8)),
			byte((value & 0x000000000000ff00) >> (1 * 8)),
			byte((value & 0x00000000000000ff) >> (0 * 8)),
		}),
	}
}

func (receiver TokenID) BigInt() *big.Int {
	bytes, something := receiver.optional.Get()
	if !something {
		return nil
	}

	var bigint big.Int
	bigint.SetBytes(bytes[:])

	return &bigint
}

func (receiver TokenID) String() string {
	bigint := receiver.BigInt()
	if nil == bigint {
		return ""
	}

	return bigint.String()
}
