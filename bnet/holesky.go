package bnet

import (
	"github.com/reiver/go-ethaddr"
)

const holeskyContractAddressHex string  = "0x762a204437f0648821e0460fcae62c072a0fa27d"
const holeskyFromBlockNumber            = 1727830

var holeskyContractAddress ethaddr.Address = ethaddr.ParseStringElsePanic(holeskyContractAddressHex)

func Holesky() (contractAddress ethaddr.Address, fromBlockNumber uint64) {
	return holeskyContractAddress, holeskyFromBlockNumber
}
