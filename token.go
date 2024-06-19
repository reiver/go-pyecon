package pyecon

import (
	"math/big"
	"time"
)

type Token interface {
	TokenID() *big.Int
	IsPrincipalToken() bool
	IsYieldToken() bool
	TokenType() TokenType

	Maturity() (time.Time, error)
	NetworkName() string
	PrincipalAmount() *big.Int
	PrincipalTokenYieldPercentage() *big.Int
	VaultAPY() *big.Int
}
