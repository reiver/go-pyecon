package bnet

import (
	gochainid "github.com/reiver/go-chainid"
	"github.com/reiver/go-ethaddr"
)

const holeskyChainID            = gochainid.Holesky
const holeskyContractAddressHex = "0x762a204437f0648821e0460fcae62c072a0fa27d"
const holeskyFromBlockNumber    = 1727830

var holeskyContractAddress ethaddr.Address = ethaddr.ParseStringElsePanic(holeskyContractAddressHex)

func Holesky() (chainid uint64, contractAddress ethaddr.Address, fromBlockNumber uint64) {
	return holeskyChainID, holeskyContractAddress, holeskyFromBlockNumber
}
