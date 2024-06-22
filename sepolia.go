package pyecon

import (
	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-pyecon/bnet"
)

// SepoliaClient returns a client for the pyecon contract on the sepolia blockchain-network.
func SepoliaClient(rpcurl string) Client {
	chainid, contractAddress, contractFromBlockNumber := bnet.Sepolia()

	return Client{
		chainid:chainid,
		contractAddress:contractAddress,
		contractFromBlockNumber:contractFromBlockNumber,
		rpcurl:opt.Something(rpcurl),
	}
}
