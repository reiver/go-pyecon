package pyecon

import (
	"context"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-pyecon/addr"
	"github.com/reiver/go-pyecon/tokenid"
)

// PositionAddress calls the "positionAddress" method on the pyecon contract, and return the address of the position-contract.
func (receiver Client) PositionAddress(tokenID tokenid.TokenID) (addr.Address, error) {
	const methodname string = "positionAddress"

	var contractAddress ethcommon.Address
	{
		var something bool

		contractAddress, something = receiver.contractAddress.Get()
		if !something {
			return addr.Nothing(), errMissingContractAddress
		}
	}

	var rpcurl string
	{
		var something bool

		rpcurl, something = receiver.rpcurl.Get()
		if !something {
			return addr.Nothing(), errMissingRPCURL
		}
	}

	var callData []byte
	{
		var err error

		callData, err = abi.Pack(methodname, tokenID.BigInt())
		if nil != err {
			return addr.Nothing(), erorr.Errorf("problem packing method-name %q for call-data: %s", methodname, err)
		}
		if nil == callData {
			return addr.Nothing(), errNilCallData
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
			return addr.Nothing(), erorr.Errorf("pyecon: problem creating RPC client: %s", err)
		}

		client = ethclient.NewClient(rpcclient)
		if nil == client {
			return addr.Nothing(), errNilRPCClient
		}
	}

	var result []byte
	{
		var err error

		result, err = client.CallContract(context.Background(), callMsg, nil)
		if nil != err {
			return addr.Nothing(), erorr.Errorf("pyecon: problem calling contract (0x%x) method %q: (%T) %s", contractAddress[:], methodname, err, err)
		}
	}

	var address addr.Address
        {

		values, err := abi.Unpack(methodname, result)
		if nil != err {
	 		return addr.Nothing(), erorr.Errorf("pyecon: problem unpacking result from call to contract (0x%x) method %q: %s", contractAddress[:], methodname, err)
		}

		var value0 interface{}
		{
			const expected = 1
			var actual int = len(values)

			if expected != actual {
				return addr.Nothing(), erorr.Errorf("pyecon: expected call to contract (0x%x) method %q to return %d result but actually got %d results", contractAddress[:], methodname, expected, actual)
			}

			value0 = values[0]
		}

		var value0commonAddress ethcommon.Address
		{
			var casted bool

			value0commonAddress, casted = value0.(ethcommon.Address)
			if !casted {
				return addr.Nothing(), erorr.Errorf("pyecon: expected call to contract (0x%x) method %q to return result of type %T but actually was %T", contractAddress[:], methodname, value0commonAddress, value0)
			}
		}

		var value0address [20]byte = [20]byte(value0commonAddress)

		address = addr.Something(value0address)
        }

	return address, nil
}
