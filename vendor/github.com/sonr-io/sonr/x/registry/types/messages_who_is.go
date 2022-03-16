package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWhoIs = "create_who_is"
	TypeMsgUpdateWhoIs = "update_who_is"
	TypeMsgDeleteWhoIs = "delete_who_is"
)

var _ sdk.Msg = &MsgCreateWhoIs{}

func NewMsgCreateWhoIs(
	creator string,
	index string,
	did string,
	value string,

) *MsgCreateWhoIs {
	return &MsgCreateWhoIs{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgCreateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhoIs) Type() string {
	return TypeMsgCreateWhoIs
}

func (msg *MsgCreateWhoIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhoIs{}

func NewMsgUpdateWhoIs(
	creator string,
	index string,
	did string,
	value string,

) *MsgUpdateWhoIs {
	return &MsgUpdateWhoIs{
		Creator: creator,
		Index:   index,
		Did:     did,
		Value:   value,
	}
}

func (msg *MsgUpdateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhoIs) Type() string {
	return TypeMsgUpdateWhoIs
}

func (msg *MsgUpdateWhoIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhoIs{}

func NewMsgDeleteWhoIs(
	creator string,
	index string,

) *MsgDeleteWhoIs {
	return &MsgDeleteWhoIs{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhoIs) Type() string {
	return TypeMsgDeleteWhoIs
}

func (msg *MsgDeleteWhoIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
