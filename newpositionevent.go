package pyecon

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

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
}

func (receiver *NewPositionEvent) setFromLog(log *ethtypes.Log) error {
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

	return nil
}

func (receiver NewPositionEvent) PrincipalAmount() *big.Int {
	var bigint *big.Int = receiver.internal.PrincipalAmount
	return big.NewInt(0).Set(bigint)
}

func (receiver NewPositionEvent) Maturity() *big.Int {
	var bigint *big.Int = receiver.internal.Maturity
	return big.NewInt(0).Set(bigint)
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
