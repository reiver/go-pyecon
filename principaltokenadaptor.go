package pyecon

import (
	"fmt"
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

func (receiver internalPrincipalTokenAdaptor) TokenName() string {
	var datecode string
	{
		maturity, err := receiver.Maturity()
		if nil == err {
			datecode = maturity.Format("02Jan2006")
		}
	}

	return fmt.Sprintf("PT-pye%s-%s", receiver.ChainCode(), datecode)
}



func (receiver internalPrincipalTokenAdaptor) Chain10Exponent() (uint64, bool) {
	return receiver.bitoken.Chain10Exponent()
}

func (receiver internalPrincipalTokenAdaptor) ChainCode() string {
	return receiver.bitoken.ChainCode()
}

func (receiver internalPrincipalTokenAdaptor) ChainID() uint64 {
	return receiver.bitoken.ChainID()
}

func (receiver internalPrincipalTokenAdaptor) ChainName() string {
	return receiver.bitoken.ChainName()
}



func (receiver internalPrincipalTokenAdaptor) Maturity() (time.Time, error) {
	return receiver.bitoken.Maturity()
}

func (receiver internalPrincipalTokenAdaptor) PrincipalAmount() *big.Int {
	return receiver.bitoken.PrincipalAmount()
}

func (receiver internalPrincipalTokenAdaptor) PrincipalTokenYieldPercentage() *big.Int {
	return receiver.bitoken.PrincipalTokenYieldPercentage()
}
