package bnet

import (
	gochainid "github.com/reiver/go-chainid"
	"github.com/reiver/go-ethaddr"
)

const holeskyChainID            = gochainid.Holesky
const holeskyContractAddressHex = "0x25ae93d88fc634e6cd9c61f04e3a33026eb40112"
const holeskyFromBlockNumber    = 1787730

var holeskyContractAddress ethaddr.Address = ethaddr.ParseStringElsePanic(holeskyContractAddressHex)

func Holesky() (chainid uint64, contractAddress ethaddr.Address, fromBlockNumber uint64) {
	return holeskyChainID, holeskyContractAddress, holeskyFromBlockNumber
}
