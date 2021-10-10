package tmservice

import (
	"context"

	ctypes "github.com/sisu-network/tendermint/rpc/core/types"

	"github.com/sisu-network/cosmos-sdk/client"
)

func getBlock(clientCtx client.Context, height *int64) (*ctypes.ResultBlock, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Block(context.Background(), height)
}
