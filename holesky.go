package pyecon

import (
	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-pyecon/bnet"
)

// HoleskyClient returns a client for the pyecon contract on the holesky blockchain-network.
func HoleskyClient(rpcurl string) Client {
	chainid, contractAddress, contractFromBlockNumber := bnet.Holesky()

	return Client{
		chainid:chainid,
		contractAddress:contractAddress,
		contractFromBlockNumber:contractFromBlockNumber,
		rpcurl:opt.Something(rpcurl),
	}
}
