package pyecon

import (
	"math/big"
	"time"

	"github.com/reiver/go-ethaddr"
	"github.com/reiver/go-ethdigest"
)

// BiToken represents a bi-token.
// A bi-token represents the combined principal-token and yield-token.
//
// The principal-token and yield-token pair are stored on the blockchain-network together.
// And when you 'fetch' the information off of the blockchain-network, you get both.
// BiToken represents that.
type BiToken interface {
	BlockDigest() ethdigest.Digest
	BlockNumber() uint64
	ContractAddress() ethaddr.Address
	Index() uint
	Maturity() (time.Time, error)
	MintedPrincipalTokenShares() *big.Int
	MintedYieldTokenShares() *big.Int
	NetworkName() string
	PositionContract() ethaddr.Address
	PrincipalAmount() *big.Int
	PrincipalTokenID() *big.Int
	PrincipalTokenYieldPercentage() *big.Int
	Removed() bool
	Topics() []ethdigest.Digest
	TxDigest() ethdigest.Digest
	TxIndex() uint
	VaultAPY() *big.Int
	YieldTokenID() *big.Int
}
