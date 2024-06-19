package pyecon

import (
	"math/big"
)

type YieldToken interface {
	Token
	MintedYieldTokenShares() *big.Int
}
