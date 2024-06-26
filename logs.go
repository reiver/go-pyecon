package pyecon

import (
	"io"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type Logs struct {
	chainid uint64
	logs []ethtypes.Log
}

func (receiver *Logs) BiToken() (BiToken, error) {
	const name string = newPositionEventName
	var event newPositionEvent

	if nil == receiver {
		return nil, errNilReceiver
	}
	if receiver.Len() < 1 {
		return nil, io.EOF
	}

	var log ethtypes.Log
	{
		log, receiver.logs = receiver.logs[0], receiver.logs[1:]
	}

	{
		err := event.setFromLog(&log, receiver.chainid)
		if nil != err {
			return nil, err
		}
	}

	return &event, nil
}

func (receiver *Logs) Len() int {
	if nil == receiver {
		return 0
	}

	return len(receiver.logs)
}

func (receiver *Logs) Next() bool {
	if nil == receiver {
		return false
	}

	return 0 < receiver.Len()
}
