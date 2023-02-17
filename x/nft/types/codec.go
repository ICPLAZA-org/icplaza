package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2

	cdc.RegisterConcrete(&MsgMintBatch{}, "nft/MintBatch", nil)

	cdc.RegisterConcrete(&MsgFrozenNft{}, "nft/FrozenNft", nil)

	cdc.RegisterConcrete(&MsgTransferNft{}, "nft/TransferNft", nil)

	cdc.RegisterConcrete(&MsgMintNft{}, "nft/MintNft", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintBatch{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFrozenNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintNft{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	//ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
	// 修改为amino编码方式，不然无法通过/txs接口签名验证
	ModuleCdc = codec.NewAminoCodec(amino)
)
//
func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}