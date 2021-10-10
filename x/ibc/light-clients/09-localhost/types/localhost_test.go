package types_test

import (
	"testing"

	tmproto "github.com/sisu-network/tendermint/proto/tendermint/types"
	"github.com/stretchr/testify/suite"

	"github.com/sisu-network/cosmos-sdk/codec"
	"github.com/sisu-network/cosmos-sdk/simapp"
	sdk "github.com/sisu-network/cosmos-sdk/types"
	clienttypes "github.com/sisu-network/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/sisu-network/cosmos-sdk/x/ibc/core/exported"
)

const (
	height = 4
)

var (
	clientHeight = clienttypes.NewHeight(0, 10)
)

type LocalhostTestSuite struct {
	suite.Suite

	cdc   codec.Marshaler
	ctx   sdk.Context
	store sdk.KVStore
}

func (suite *LocalhostTestSuite) SetupTest() {
	isCheckTx := false
	app := simapp.Setup(isCheckTx)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{Height: 1, ChainID: "ibc-chain"})
	suite.store = app.IBCKeeper.ClientKeeper.ClientStore(suite.ctx, exported.Localhost)
}

func TestLocalhostTestSuite(t *testing.T) {
	suite.Run(t, new(LocalhostTestSuite))
}
