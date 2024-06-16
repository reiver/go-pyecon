package addr

import (
	"github.com/reiver/go-ethaddr"
)

const holeskyAddress string = "0x762a204437f0648821e0460fcae62c072a0fa27d"

var holesky ethaddr.Address = ethaddr.ParseStringElsePanic(holeskyAddress)

func Holesky() ethaddr.Address {
	return holesky
}
