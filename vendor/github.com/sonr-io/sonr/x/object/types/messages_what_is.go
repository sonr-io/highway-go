package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWhatIs = "create_what_is"
	TypeMsgUpdateWhatIs = "update_what_is"
	TypeMsgDeleteWhatIs = "delete_what_is"
)

var _ sdk.Msg = &MsgCreateWhatIs{}

func NewMsgCreateWhatIs(
	creator string,
	index string,
	did string,
	value []byte,

) *MsgCreateWhatIs {
	return &MsgCreateWhatIs{
		Creator:  creator,
		Index:    index,
		Did:      did,
		Document: value,
	}
}

func (msg *MsgCreateWhatIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhatIs) Type() string {
	return TypeMsgCreateWhatIs
}

func (msg *MsgCreateWhatIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhatIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhatIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhatIs{}

func NewMsgUpdateWhatIs(
	creator string,
	index string,
	did string,
	value []byte,

) *MsgUpdateWhatIs {
	return &MsgUpdateWhatIs{
		Creator:  creator,
		Index:    index,
		Did:      did,
		Document: value,
	}
}

func (msg *MsgUpdateWhatIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhatIs) Type() string {
	return TypeMsgUpdateWhatIs
}

func (msg *MsgUpdateWhatIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhatIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhatIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhatIs{}

func NewMsgDeleteWhatIs(
	creator string,
	index string,

) *MsgDeleteWhatIs {
	return &MsgDeleteWhatIs{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteWhatIs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhatIs) Type() string {
	return TypeMsgDeleteWhatIs
}

func (msg *MsgDeleteWhatIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhatIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhatIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
