package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReadObject = "read_object"

var _ sdk.Msg = &MsgReadObject{}

func NewMsgReadObject(creator string, did string) *MsgReadObject {
	return &MsgReadObject{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgReadObject) Route() string {
	return RouterKey
}

func (msg *MsgReadObject) Type() string {
	return TypeMsgReadObject
}

func (msg *MsgReadObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReadObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReadObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
