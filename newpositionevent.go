package pyecon

import (
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/reiver/go-ethaddr"
	"github.com/reiver/go-ethdigest"
	"sourcecode.social/reiver/go-erorr"
)

const newPositionEventName string = "NewPositionEvent"

type internalNewPositionEvent struct {
	PrincipalAmount    *big.Int
	Maturity           *big.Int
	PtYieldPercent     *big.Int
	PTokenId           *big.Int
	MintedPtShares     *big.Int
	YTokenId           *big.Int
	MintedYtShares     *big.Int
	VaultApy           *big.Int
	PositionContract    ethcommon.Address
}

// NewPositionEvent represents a "NewPositionEvent" in the logs of the contract from the network.
type NewPositionEvent struct {
	Log
	internal internalNewPositionEvent
	networkName string
}

var _ BiToken = NewPositionEvent{}

func (receiver *NewPositionEvent) setFromLog(log *ethtypes.Log, networkName string) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == log {
		return errNilLog
	}

	err := receiver.Log.setFromLog(log)
	if nil != err {
		return err
	}

	err = abi.UnpackIntoInterface(&receiver.internal, newPositionEventName, log.Data)
	if nil != err {
		return erorr.Errorf("pyecon: problem unpacking log-data from log from contract into a %T: %w", receiver.internal, err)
	}

	receiver.networkName = networkName

	return nil
}

func (receiver NewPositionEvent) NetworkName() string {
	return receiver.networkName
}

func (receiver NewPositionEvent) PrincipalAmount() *big.Int {
	var bigint *big.Int = receiver.internal.PrincipalAmount
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) Maturity() (time.Time, error) {
	var bigint *big.Int = receiver.internal.Maturity

	if nil == bigint  {
		return time.Time{}, errNilMaturity
	}
	if !bigint.IsInt64() {
		return time.Time{}, errMaturityOverflow
	}

	var unixtimestamp int64 = bigint.Int64()

	return time.Unix(unixtimestamp, 0), nil
}

func (receiver NewPositionEvent) PrincipalTokenYieldPercentage() *big.Int {
	var bigint *big.Int = receiver.internal.PtYieldPercent
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) PrincipalTokenID() *big.Int {
	var bigint *big.Int = receiver.internal.PTokenId
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) MintedPrincipalTokenShares() *big.Int {
	var bigint *big.Int = receiver.internal.MintedPtShares
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) YieldTokenID() *big.Int {
	var bigint *big.Int = receiver.internal.YTokenId
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) MintedYieldTokenShares() *big.Int {
	var bigint *big.Int = receiver.internal.MintedYtShares
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) VaultAPY() *big.Int {
	var bigint *big.Int = receiver.internal.VaultApy
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) PositionContract() ethaddr.Address {
	return ethaddr.Something(receiver.internal.PositionContract)
}







func (receiver NewPositionEvent) ContractAddress() ethaddr.Address {
	return receiver.contractAddress
}

func (receiver NewPositionEvent) BlockDigest() ethdigest.Digest {
	return receiver.blockDigest
}

func (receiver NewPositionEvent) BlockNumber() uint64 {
	return receiver.blockNumber
}

func (receiver NewPositionEvent) Index() uint {
	return receiver.index
}

func (receiver NewPositionEvent) Removed() bool {
	return receiver.removed
}

func (receiver NewPositionEvent) Topics() []ethdigest.Digest {
	return append([]ethdigest.Digest(nil), receiver.topics...)
}

func (receiver NewPositionEvent) TxDigest() ethdigest.Digest {
	return receiver.txDigest
}

func (receiver NewPositionEvent) TxIndex() uint {
	return receiver.txIndex
}
