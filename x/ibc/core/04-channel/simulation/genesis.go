package simulation

import (
	"math/rand"

	simtypes "github.com/sisu-network/cosmos-sdk/types/simulation"
	"github.com/sisu-network/cosmos-sdk/x/ibc/core/04-channel/types"
)

// GenChannelGenesis returns the default channel genesis state.
func GenChannelGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
