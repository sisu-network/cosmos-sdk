package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sisu-network/cosmos-sdk/codec"
	codectypes "github.com/sisu-network/cosmos-sdk/codec/types"
	"github.com/sisu-network/cosmos-sdk/std"
	"github.com/sisu-network/cosmos-sdk/testutil/testdata"
	sdk "github.com/sisu-network/cosmos-sdk/types"
	"github.com/sisu-network/cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
