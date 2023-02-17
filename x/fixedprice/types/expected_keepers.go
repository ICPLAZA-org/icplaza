package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	pooltypes "github.com/gauss/gauss/v6/x/nftexpool/types"
	"github.com/gauss/gauss/v6/x/nft/types"
)

type NftKeeper interface {
	SaveNFT(ctx sdk.Context, nft *types.Nft) error
	SetOwnerNFT(ctx sdk.Context, cateId string, tokenId string, owner string) error
	DeleteOwnerNFT(ctx sdk.Context, cateId string, tokenId string, owner string)
	GetNFT(ctx sdk.Context, cateId, tokenId string) (*types.Nft, error)
	LockNFT(ctx sdk.Context, cateId, tokenId string) (err error)
	FreeNFT(ctx sdk.Context, cateId, tokenId string) error
	CheckNFT(nft types.Nft) error
	GetLastPrice(ctx sdk.Context, tokenId string, denom string) sdk.Coin
	SaveLastPrice(ctx sdk.Context, tokenId string, price sdk.Coin) error
	Transfer(ctx sdk.Context, from, to string, nft types.Nft, transferType uint64) error
}

type PoolKeeper interface {
	PoolExists(ctx sdk.Context, poolAddress string) bool
	CheckPool(ctx sdk.Context, poolAddress string) error
	GetPool(ctx sdk.Context, poolAddress string) (*pooltypes.Pool, error)
}

type BankKeeper interface {
	InputOutputCoins(ctx sdk.Context, inputs []banktypes.Input, outputs []banktypes.Output) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}
