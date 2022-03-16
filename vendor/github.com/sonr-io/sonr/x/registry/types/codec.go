package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterService{}, "registry/RegisterService", nil)
	cdc.RegisterConcrete(&MsgRegisterName{}, "registry/RegisterName", nil)
	cdc.RegisterConcrete(&MsgAccessName{}, "registry/AccessName", nil)
	cdc.RegisterConcrete(&MsgUpdateName{}, "registry/UpdateName", nil)
	cdc.RegisterConcrete(&MsgAccessService{}, "registry/AccessService", nil)
	cdc.RegisterConcrete(&MsgUpdateService{}, "registry/UpdateService", nil)
	cdc.RegisterConcrete(&MsgCreateWhoIs{}, "registry/CreateWhoIs", nil)
	cdc.RegisterConcrete(&MsgUpdateWhoIs{}, "registry/UpdateWhoIs", nil)
	cdc.RegisterConcrete(&MsgDeleteWhoIs{}, "registry/DeleteWhoIs", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhoIs{},
		&MsgUpdateWhoIs{},
		&MsgDeleteWhoIs{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
