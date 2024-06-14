package pyecon

import (
	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-pyecon/addr"
)

// Client is a client for a pyecon contract on some blockchain-network.
type Client struct {
	contractAddress addr.Address
	rpcurl          opt.Optional[string]
}
