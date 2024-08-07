package pyecon

import (
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/reiver/go-chain10"
	"github.com/reiver/go-chaincode"
	"github.com/reiver/go-chainname"
	"github.com/reiver/go-ethaddr"
	"github.com/reiver/go-ethdigest"
	"sourcecode.social/reiver/go-erorr"
)

const newPositionEventName string = "NewPositionEvent"

type internalNewPositionEvent struct {
	PrincipalAmount    *big.Int
	Maturity           *big.Int
	PtYieldPercent     *big.Int
	PtId               *big.Int
	PtAmount           *big.Int
	YtId               *big.Int
	YtAmount           *big.Int
	PositionContract    ethcommon.Address
}

// newPositionEvent represents a "NewPositionEvent" in the logs of the contract from the network.
type newPositionEvent struct {
	Log
	internal internalNewPositionEvent
	chainid uint64
}

var _ BiToken = newPositionEvent{}

func (receiver *newPositionEvent) setFromLog(log *ethtypes.Log, chainid uint64) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == log {
		return errNilLog
	}

	err := receiver.Log.setFromEthereumLog(log)
	if nil != err {
		return err
	}

	err = abi.UnpackIntoInterface(&receiver.internal, newPositionEventName, log.Data)
	if nil != err {
		return erorr.Errorf("pyecon: problem unpacking log-data from log from contract into a %T: %w", receiver.internal, err)
	}

	receiver.chainid = chainid

	return nil
}

func (receiver newPositionEvent) Chain10Exponent() (uint64, bool) {
	return chain10.Exponent(receiver.chainid)
}

func (receiver newPositionEvent) ChainCode() string {
	return chaincode.ChainCode(receiver.chainid)
}

func (receiver newPositionEvent) ChainID() uint64 {
	return receiver.chainid
}

func (receiver newPositionEvent) ChainName() string {
	return chainname.ChainName(receiver.chainid)
}

func (receiver newPositionEvent) PrincipalAmount() *big.Int {
	var bigint *big.Int = receiver.internal.PrincipalAmount
	return big.NewInt(0).Set(bigint)
}

func (receiver newPositionEvent) Maturity() (time.Time, error) {
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

func (receiver newPositionEvent) PrincipalTokenYieldPercentage() *big.Int {
	var bigint *big.Int = receiver.internal.PtYieldPercent
	return big.NewInt(0).Set(bigint)
}

func (receiver newPositionEvent) PrincipalTokenID() *big.Int {
	var bigint *big.Int = receiver.internal.PtId
	return big.NewInt(0).Set(bigint)
}

func (receiver newPositionEvent) MintedPrincipalTokenShares() *big.Int {
	var bigint *big.Int = receiver.internal.PtAmount
	return big.NewInt(0).Set(bigint)
}

func (receiver newPositionEvent) YieldTokenID() *big.Int {
	var bigint *big.Int = receiver.internal.YtId
	return big.NewInt(0).Set(bigint)
}

func (receiver newPositionEvent) MintedYieldTokenShares() *big.Int {
	var bigint *big.Int = receiver.internal.YtAmount
	return big.NewInt(0).Set(bigint)
}

func (receiver newPositionEvent) PositionAddress() ethaddr.Address {
	return ethaddr.Something(receiver.internal.PositionContract)
}







func (receiver newPositionEvent) ContractAddress() ethaddr.Address {
	return receiver.contractAddress
}

func (receiver newPositionEvent) BlockDigest() ethdigest.Digest {
	return receiver.blockDigest
}

func (receiver newPositionEvent) BlockNumber() uint64 {
	return receiver.blockNumber
}

func (receiver newPositionEvent) Index() uint {
	return receiver.index
}

func (receiver newPositionEvent) Removed() bool {
	return receiver.removed
}

func (receiver newPositionEvent) Topics() []ethdigest.Digest {
	return append([]ethdigest.Digest(nil), receiver.topics...)
}

func (receiver newPositionEvent) TxDigest() ethdigest.Digest {
	return receiver.txDigest
}

func (receiver newPositionEvent) TxIndex() uint {
	return receiver.txIndex
}



func (receiver newPositionEvent) PrincipalToken() PrincipalToken {
	return internalPrincipalTokenAdaptor{
		bitoken: &receiver,
	}
}

func (receiver newPositionEvent) YieldToken() YieldToken {
	return internalYieldTokenAdaptor{
		bitoken: &receiver,
	}
}
