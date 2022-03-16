package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteObject = "delete_object"

var _ sdk.Msg = &MsgDeleteObject{}

func NewMsgDeleteObject(creator string, did string, publicKey string) *MsgDeleteObject {
	return &MsgDeleteObject{
		Creator:   creator,
		Did:       did,
		PublicKey: publicKey,
	}
}

func (msg *MsgDeleteObject) Route() string {
	return RouterKey
}

func (msg *MsgDeleteObject) Type() string {
	return TypeMsgDeleteObject
}

func (msg *MsgDeleteObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
