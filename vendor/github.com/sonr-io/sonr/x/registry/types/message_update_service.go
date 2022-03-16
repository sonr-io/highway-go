package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateService = "update_service"

var _ sdk.Msg = &MsgUpdateService{}

func NewMsgUpdateService(creator string, did string) *MsgUpdateService {
	return &MsgUpdateService{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgUpdateService) Route() string {
	return RouterKey
}

func (msg *MsgUpdateService) Type() string {
	return TypeMsgUpdateService
}

func (msg *MsgUpdateService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
