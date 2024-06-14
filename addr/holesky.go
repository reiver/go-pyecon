package addr

const holeskyAddress string = "0x762a204437f0648821e0460fcae62c072a0fa27d"

var holesky Address

func init() {
	const src string = holeskyAddress

	var buffer [len(src)]byte
	copy(buffer[:], src)

	var compiled [(len(src)-2)/2]byte

	err := compile(compiled[:], buffer[:])
	if nil != err {
		panic(err)
	}

	holesky = Something(compiled)
}

func Holesky() Address {
	return holesky
}
