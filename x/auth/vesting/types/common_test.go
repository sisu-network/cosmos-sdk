package types_test

import (
	"github.com/sisu-network/cosmos-sdk/simapp"
)

var (
	app         = simapp.Setup(false)
	appCodec, _ = simapp.MakeCodecs()
)
