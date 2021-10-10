package simulation

import (
	"math/rand"

	simtypes "github.com/sisu-network/cosmos-sdk/types/simulation"
	"github.com/sisu-network/cosmos-sdk/x/ibc/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
