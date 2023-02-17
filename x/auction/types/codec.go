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
	// this line is used by starport scaffolding # 2

	cdc.RegisterConcrete(&MsgCreateOrder{}, "auction/CreateOrder", nil)
	cdc.RegisterConcrete(&MsgDeleteOrder{}, "auction/DeleteOrder", nil)
	cdc.RegisterConcrete(&MsgBidOrder{}, "auction/BidOrder", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBidOrder{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteOrder{},
	)

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