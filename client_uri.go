package pyecon

import (
	"context"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-pyecon/tokenid"
)

// URI calls the "uri" method on the pyecon contract, and return the URI to the associate metadata.
func (receiver Client) URI(tokenID tokenid.TokenID) (string, error) {
	const methodname string = "uri"

	var contractAddress ethcommon.Address
	{
		var something bool

		contractAddress, something = receiver.contractAddress.Get()
		if !something {
			return "", errMissingContractAddress
		}
	}

	var rpcurl string
	{
		var something bool

		rpcurl, something = receiver.rpcurl.Get()
		if !something {
			return "", errMissingRPCURL
		}
	}

	var callData []byte
	{
		var err error

		callData, err = abi.Pack(methodname, tokenID.BigInt())
		if nil != err {
			return "", erorr.Errorf("problem packing method-name %q for call-data: %s", methodname, err)
		}
		if nil == callData {
			return "", errNilCallData
		}
	}

	var callMsg = ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	var client *ethclient.Client
	{
		rpcclient, err := ethrpc.Dial(rpcurl)
		if err != nil {
			return "", erorr.Errorf("pyecon: problem creating RPC client: %s", err)
		}

		client = ethclient.NewClient(rpcclient)
		if nil == client {
			return "", errNilRPCClient
		}
	}

	var result []byte
	{
		var err error

		result, err = client.CallContract(context.Background(), callMsg, nil)
		if nil != err {
			return "", erorr.Errorf("pyecon: problem calling contract (0x%x) method %q: (%T) %s", contractAddress[:], methodname, err, err)
		}
	}

	var uri string
        {

		values, err := abi.Unpack(methodname, result)
		if nil != err {
			return "", erorr.Errorf("pyecon: problem unpacking result from call to contract (0x%x) method %q: %s", contractAddress[:], methodname, err)
		}

		var value0 interface{}
		{
			const expected = 1
			var actual int = len(values)

			if expected != actual {
				return "", erorr.Errorf("pyecon: expected call to contract (0x%x) method %q to return %d result but actually got %d results", contractAddress[:], methodname, expected, actual)
			}

			value0 = values[0]
		}

		{
			var casted bool

			uri, casted = value0.(string)
			if !casted {
				return "", erorr.Errorf("pyecon: expected call to contract (0x%x) method %q to return result of type %T but actually was %T", contractAddress[:], methodname, uri, value0)
			}
		}
	}

	return uri, nil
}
