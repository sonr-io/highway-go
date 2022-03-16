package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterService = "register_service"

var _ sdk.Msg = &MsgRegisterService{}

func NewMsgRegisterService(creator string, serviceName string, jwt string) *MsgRegisterService {
	return &MsgRegisterService{
		Creator:     creator,
		ServiceName: serviceName,
		Jwt:         jwt,
	}
}

func (msg *MsgRegisterService) Route() string {
	return RouterKey
}

func (msg *MsgRegisterService) Type() string {
	return TypeMsgRegisterService
}

func (msg *MsgRegisterService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
