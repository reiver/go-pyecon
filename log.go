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
	topics        []ethdigest.Digest
	txDigest        ethdigest.Digest
	txIndex         uint
}

func (receiver *Log) setFromEthereumLog(log *ethtypes.Log) error {
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

	for _,topicHash := range log.Topics {

		var topicDigest ethdigest.Digest = ethdigest.Something(topicHash)

		receiver.topics = append(receiver.topics, topicDigest)
	}

	return nil
}
