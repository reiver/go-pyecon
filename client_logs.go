package pyecon

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"sourcecode.social/reiver/go-erorr"
)

// Logs returns the logs for the pyecon contract on the network.
func (receiver Client) Logs(fromBlockNumber *big.Int, toBlockNumber *big.Int) (Logs, error) {

	var contractAddress ethcommon.Address
	{
		var something bool

		contractAddress, something = receiver.contractAddress.Get()
		if !something {
			return Logs{}, errMissingContractAddress
		}
	}

	var contractFromBlockNumber uint64 = receiver.contractFromBlockNumber

	var rpcurl string
	{
		var something bool

		rpcurl, something = receiver.rpcurl.Get()
		if !something {
			return Logs{}, errMissingRPCURL
		}
	}

	var client *ethclient.Client
	{
		rpcclient, err := ethrpc.Dial(rpcurl)
		if err != nil {
			return Logs{}, erorr.Errorf("pyecon: problem creating RPC client: %s", err)
		}
		defer rpcclient.Close()

		client = ethclient.NewClient(rpcclient)
		if nil == client {
			return Logs{}, errNilRPCClient
		}
		defer client.Close()
	}

//	var currentBlockNumber *big.Int
//	{
//		header, err := client.HeaderByNumber(context.Background(), nil)
//		if err != nil {
//			return Logs{}, erorr.Errorf("pyecon: problem getting the current (latest) block-number: %s", err)
//		}
//
//		currentBlockNumber = header.Number
//	}

	{
		if nil == fromBlockNumber {
			fromBlockNumber = big.NewInt(0).SetUint64(contractFromBlockNumber)
		}
	}

	query := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{contractAddress},
		FromBlock: fromBlockNumber,
		ToBlock:   toBlockNumber,
	}

	var logs []ethtypes.Log
	{
		var err error

		logs, err = client.FilterLogs(context.Background(), query)
		if nil != err {
			return Logs{}, erorr.Errorf("pyecon: problem getting logs for contract (0x%x): %s", contractAddress[:], err)
		}

	}

	return Logs{logs}, nil
}
