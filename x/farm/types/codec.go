package types

import (
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "gauss/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(&MsgDestroyPool{}, "gauss/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(&MsgStake{}, "gauss/farm/MsgStake", nil)
	cdc.RegisterConcrete(&MsgUnbond{}, "gauss/farm/MsgUnbond", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "gauss/farm/MsgWithdraw", nil)
}

// RegisterInterfaces registers the interface
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreatePool{},
		&MsgDestroyPool{},
		&MsgStake{},
		&MsgUnbond{},
		&MsgWithdraw{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// MustUnMarshalPoolName return the poolName protobuf code
func MustMarshalPoolName(cdc codec.BinaryCodec, poolName string) []byte {
	poolNameWrap := gogotypes.StringValue{Value: poolName}
	return cdc.MustMarshal(&poolNameWrap)
}

// MustUnMarshalPoolName return the poolName
func MustUnMarshalPoolName(cdc codec.BinaryCodec, poolName []byte) string {
	var poolNameWrap gogotypes.StringValue
	cdc.MustUnmarshal(poolName, &poolNameWrap)
	return poolNameWrap.Value
}
