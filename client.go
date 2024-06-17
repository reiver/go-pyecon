package pyecon

import (
	"github.com/reiver/go-ethaddr"
	"sourcecode.social/reiver/go-opt"

//	"github.com/reiver/go-pyecon/blknum"
)

// Client is a client for a pyecon contract on some blockchain-network.
type Client struct {
	contractAddress         ethaddr.Address
	contractFromBlockNumber uint64
	rpcurl                  opt.Optional[string]
}

