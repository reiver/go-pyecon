package pyecon

import (
	"math/big"
)

type PrincipalToken interface {
	Token
	MintedPrincipalTokenShares() *big.Int
}
