package auth_test

import (
	"testing"

	abcitypes "github.com/sisu-network/tendermint/abci/types"
	tmproto "github.com/sisu-network/tendermint/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	"github.com/sisu-network/cosmos-sdk/simapp"
	"github.com/sisu-network/cosmos-sdk/x/auth/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	acc := app.AccountKeeper.GetAccount(ctx, types.NewModuleAddress(types.FeeCollectorName))
	require.NotNil(t, acc)
}
