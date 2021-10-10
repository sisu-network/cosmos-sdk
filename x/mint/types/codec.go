package types

import (
	"github.com/sisu-network/cosmos-sdk/codec"
	cryptocodec "github.com/sisu-network/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
