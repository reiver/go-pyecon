package pyecon

import (
	"math/big"
	"time"

	"github.com/reiver/go-ethaddr"
)

type Token interface {
	TokenID() *big.Int
	PositionAddress() ethaddr.Address
	IsPrincipalToken() bool
	IsYieldToken() bool
	TokenType() TokenType
	TokenName() string

	Chain10Exponent() (uint64, bool)
	ChainCode() string
	ChainID() uint64
	ChainName() string

	Maturity() (time.Time, error)
	PrincipalAmount() *big.Int
	PrincipalTokenYieldPercentage() *big.Int
}
