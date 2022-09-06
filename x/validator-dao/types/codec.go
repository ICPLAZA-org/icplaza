package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)


func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgReqAuthorization{}, "validatordao/ReqAuthorization", nil)
	cdc.RegisterConcrete(&MsgWithdrawAuthorization{}, "validatordao/WithdrawAuthorization", nil)
	cdc.RegisterConcrete(&MsgAddAuthBiz{}, "validatordao/AddAuthBiz", nil)
	cdc.RegisterConcrete(&MsgUpdateAuthBiz{}, "validatordao/UpdateAuthBiz", nil)
	cdc.RegisterConcrete(&MsgRemoveAuthBiz{}, "validatordao/RemoveAuthBiz", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReqAuthorization{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawAuthorization{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAuthBiz{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateAuthBiz{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveAuthBiz{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	// ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
