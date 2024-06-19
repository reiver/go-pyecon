package pyecon

import (
	"math/big"
	"time"
)

type internalYieldTokenAdaptor struct {
	bitoken BiToken
}

var _ Token = internalYieldTokenAdaptor{}
var _ YieldToken = internalYieldTokenAdaptor{}



func (receiver internalYieldTokenAdaptor) MintedYieldTokenShares() *big.Int {
	return receiver.bitoken.MintedYieldTokenShares()
}



func (receiver internalYieldTokenAdaptor) TokenID() *big.Int {
	return receiver.bitoken.YieldTokenID()
}

func (receiver internalYieldTokenAdaptor) IsPrincipalToken() bool {
	return false
}

func (receiver internalYieldTokenAdaptor) IsYieldToken() bool {
	return true
}

func (receiver internalYieldTokenAdaptor) TokenType() TokenType {
	return TokenTypeYieldToken
}



func (receiver internalYieldTokenAdaptor) Maturity() (time.Time, error) {
	return receiver.bitoken.Maturity()
}

func (receiver internalYieldTokenAdaptor) NetworkName() string {
	return receiver.bitoken.NetworkName()
}

func (receiver internalYieldTokenAdaptor) PrincipalAmount() *big.Int {
	return receiver.bitoken.PrincipalAmount()
}

func (receiver internalYieldTokenAdaptor) PrincipalTokenYieldPercentage() *big.Int {
	return receiver.bitoken.PrincipalTokenYieldPercentage()
}

func (receiver internalYieldTokenAdaptor) VaultAPY() *big.Int {
	return receiver.bitoken.VaultAPY()
}
