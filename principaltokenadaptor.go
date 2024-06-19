package pyecon

import (
	"math/big"
	"time"
)

type internalPrincipalTokenAdaptor struct {
	bitoken BiToken
}

var _ Token = internalPrincipalTokenAdaptor{}
var _ PrincipalToken = internalPrincipalTokenAdaptor{}



func (receiver internalPrincipalTokenAdaptor) MintedPrincipalTokenShares() *big.Int {
	return receiver.bitoken.MintedPrincipalTokenShares()
}



func (receiver internalPrincipalTokenAdaptor) TokenID() *big.Int {
	return receiver.bitoken.PrincipalTokenID()
}

func (receiver internalPrincipalTokenAdaptor) IsPrincipalToken() bool {
	return true
}

func (receiver internalPrincipalTokenAdaptor) IsYieldToken() bool {
	return false
}

func (receiver internalPrincipalTokenAdaptor) TokenType() TokenType {
	return TokenTypePrincipalToken
}



func (receiver internalPrincipalTokenAdaptor) Maturity() (time.Time, error) {
	return receiver.bitoken.Maturity()
}

func (receiver internalPrincipalTokenAdaptor) NetworkName() string {
	return receiver.bitoken.NetworkName()
}

func (receiver internalPrincipalTokenAdaptor) PrincipalAmount() *big.Int {
	return receiver.bitoken.PrincipalAmount()
}

func (receiver internalPrincipalTokenAdaptor) PrincipalTokenYieldPercentage() *big.Int {
	return receiver.bitoken.PrincipalTokenYieldPercentage()
}

func (receiver internalPrincipalTokenAdaptor) VaultAPY() *big.Int {
	return receiver.bitoken.VaultAPY()
}
