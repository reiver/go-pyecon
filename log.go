package pyecon

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/reiver/go-ethaddr"
	"github.com/reiver/go-ethdigest"
)

type Log struct {
	contractAddress ethaddr.Address
	blockDigest     ethdigest.Digest
	blockNumber     uint64
	index           uint
	removed         bool
	txDigest        ethdigest.Digest
	txIndex         uint
}

func (receiver *Log) setFromLog(log *ethtypes.Log) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == log {
		return errNilLog
	}

	receiver.contractAddress = ethaddr.Something(log.Address)
	receiver.blockDigest     = ethdigest.Something(log.BlockHash)
	receiver.blockNumber     = log.BlockNumber
	receiver.index           = log.Index
	receiver.removed         = log.Removed
	receiver.txDigest        = ethdigest.Something(log.TxHash)
	receiver.txIndex         = log.TxIndex

	return nil
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

func (receiver NewPositionEvent) TxDigest() ethdigest.Digest {
	return receiver.txDigest
}

func (receiver NewPositionEvent) TxIndex() uint {
	return receiver.txIndex
}
