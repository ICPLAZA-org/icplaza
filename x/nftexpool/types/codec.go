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
	cdc.RegisterConcrete(&MsgUndelegate{}, "pool/Undelegate", nil)

	cdc.RegisterConcrete(&MsgDelegate{}, "pool/Delegate", nil)

	cdc.RegisterConcrete(&MsgUpdatePool{}, "pool/UpdatePool", nil)

	cdc.RegisterConcrete(&MsgCreatePool{}, "pool/CreatePool", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndelegate{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDelegate{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePool{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
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
