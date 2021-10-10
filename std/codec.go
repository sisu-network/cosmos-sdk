package std

import (
	"github.com/sisu-network/cosmos-sdk/codec"
	"github.com/sisu-network/cosmos-sdk/codec/types"
	cryptocodec "github.com/sisu-network/cosmos-sdk/crypto/codec"
	sdk "github.com/sisu-network/cosmos-sdk/types"
	txtypes "github.com/sisu-network/cosmos-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
