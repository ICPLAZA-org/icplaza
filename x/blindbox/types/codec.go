package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "blindbox/CreatePool", nil)
	cdc.RegisterConcrete(&MsgUpdatePool{}, "blindbox/UpdatePool", nil)
	cdc.RegisterConcrete(&MsgRemovePool{}, "blindbox/RemovePool", nil)
	cdc.RegisterConcrete(&MsgCreateBox{}, "blindbox/CreateBox", nil)
	cdc.RegisterConcrete(&MsgRevokeBoxGroup{}, "blindbox/RevokeBoxGroup", nil)
	cdc.RegisterConcrete(&MsgOpenBox{}, "blindbox/OpenBox", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemovePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBox{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeBoxGroup{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenBox{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
	//ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}