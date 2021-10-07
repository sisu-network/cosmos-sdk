package tx

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	typestx "github.com/cosmos/cosmos-sdk/types/tx"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

const aminoNonCriticalFieldsError = "protobuf transaction contains unknown non-critical fields. This is a transaction malleability issue and SIGN_MODE_LEGACY_AMINO_JSON cannot be used."

var _ signing.SignModeHandler = signModeLegacyAminoJSONHandler{}

// signModeLegacyAminoJSONHandler defines the SIGN_MODE_LEGACY_AMINO_JSON
// SignModeHandler.
type signModeLegacyAminoJSONHandler struct{}

func (s signModeLegacyAminoJSONHandler) DefaultMode() signingtypes.SignMode {
	return signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
}

func (s signModeLegacyAminoJSONHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON}
}

func (s signModeLegacyAminoJSONHandler) GetSignBytes(mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	if mode != signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON {
		return nil, fmt.Errorf("expected %s, got %s", signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, mode)
	}

	protoTx, ok := tx.(*wrapper)
	if !ok {
		return nil, fmt.Errorf("can only handle a protobuf Tx, got %T", tx)
	}

	if protoTx.txBodyHasUnknownNonCriticals {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, aminoNonCriticalFieldsError)
	}

	body := protoTx.tx.Body

	if len(body.ExtensionOptions) != 0 || len(body.NonCriticalExtensionOptions) != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "SIGN_MODE_LEGACY_AMINO_JSON does not support protobuf extension options")
	}

	// We set a convention that if the tipper signs with LEGACY_AMINO_JSON, then
	// they sign over empty fees and 0 gas.
	var isTipper bool
	addr := data.Address
	if addr == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "got empty address in SIGN_MODE_LEGACY_AMINO_JSON handler")
	}
	if tipTx, ok := tx.(typestx.TipTx); ok && tipTx.GetTip() != nil {
		isTipper = tipTx.GetTip().Tipper == addr.String()
	}

	if isTipper {
		return legacytx.StdSignBytes(
			data.ChainID, data.AccountNumber, data.Sequence, protoTx.GetTimeoutHeight(),
			// The tipper signs over 0 fee and 0 gas, by convention.
			legacytx.StdFee{Amount: sdk.Coins{}, Gas: 0},
			tx.GetMsgs(), protoTx.GetMemo(),
		), nil
	}

	return legacytx.StdSignBytes(
		data.ChainID, data.AccountNumber, data.Sequence, protoTx.GetTimeoutHeight(),
		legacytx.StdFee{Amount: protoTx.GetFee(), Gas: protoTx.GetGas()},
		tx.GetMsgs(), protoTx.GetMemo(),
	), nil
}
