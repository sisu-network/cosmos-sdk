package tmservice

import (
	"context"

	ctypes "github.com/sisu-network/tendermint/rpc/core/types"

	"github.com/sisu-network/cosmos-sdk/client"
)

func getNodeStatus(clientCtx client.Context) (*ctypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}
	return node.Status(context.Background())
}
