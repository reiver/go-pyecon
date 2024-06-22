package bnet

import (
	gochainid "github.com/reiver/go-chainid"
	"github.com/reiver/go-ethaddr"
)

const sepoliaChainID            = gochainid.Sepolia
const sepoliaContractAddressHex = "0xbee9c33615e574D2070043c61282127CA2943B97"
const sepoliaFromBlockNumber    = 6162714

var sepoliaContractAddress ethaddr.Address = ethaddr.ParseStringElsePanic(sepoliaContractAddressHex)

func Sepolia() (chainid uint64, contractAddress ethaddr.Address, fromBlockNumber uint64) {
	return sepoliaChainID, sepoliaContractAddress, sepoliaFromBlockNumber
}
