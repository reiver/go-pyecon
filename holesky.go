package pyecon

import (
	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-pyecon/addr"
)

// HoleskyClient returns a client for the pyecon contract on the holesky blockchain-network.
func HoleskyClient(rpcurl string) Client {
	contractAddress := addr.Holesky()

	return Client{
		contractAddress:contractAddress,
		rpcurl:opt.Something(rpcurl),
	}
}
