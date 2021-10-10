package bank

import (
	"strings"
	"testing"

	tmproto "github.com/sisu-network/tendermint/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	"github.com/sisu-network/cosmos-sdk/testutil/testdata"
	sdk "github.com/sisu-network/cosmos-sdk/types"
	sdkerrors "github.com/sisu-network/cosmos-sdk/types/errors"
)

func TestInvalidMsg(t *testing.T) {
	h := NewHandler(nil)

	res, err := h(sdk.NewContext(nil, tmproto.Header{}, false, nil), testdata.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)

	_, _, log := sdkerrors.ABCIInfo(err, false)
	require.True(t, strings.Contains(log, "unrecognized bank message type"))
}
