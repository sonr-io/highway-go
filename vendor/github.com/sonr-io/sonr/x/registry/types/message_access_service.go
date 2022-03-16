package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAccessService = "access_service"

var _ sdk.Msg = &MsgAccessService{}

func NewMsgAccessService(creator string, did string) *MsgAccessService {
	return &MsgAccessService{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgAccessService) Route() string {
	return RouterKey
}

func (msg *MsgAccessService) Type() string {
	return TypeMsgAccessService
}

func (msg *MsgAccessService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAccessService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAccessService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
