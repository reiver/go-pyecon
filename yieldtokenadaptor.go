package pyecon

import (
	"fmt"
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

func (receiver internalYieldTokenAdaptor) TokenName() string {
	var datecode string
	{
		maturity, err := receiver.Maturity()
		if nil == err {
			datecode = maturity.Format("02Jan2006")
		}
	}

	return fmt.Sprintf("YT-pye%s-%s", receiver.ChainCode(), datecode)
}



func (receiver internalYieldTokenAdaptor) Chain10Exponent() (uint64, bool) {
	return receiver.bitoken.Chain10Exponent()
}

func (receiver internalYieldTokenAdaptor) ChainCode() string {
	return receiver.bitoken.ChainCode()
}

func (receiver internalYieldTokenAdaptor) ChainID() uint64 {
	return receiver.bitoken.ChainID()
}

func (receiver internalYieldTokenAdaptor) ChainName() string {
	return receiver.bitoken.ChainName()
}



func (receiver internalYieldTokenAdaptor) Maturity() (time.Time, error) {
	return receiver.bitoken.Maturity()
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
