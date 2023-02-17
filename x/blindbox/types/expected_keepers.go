package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	nfttype "github.com/gauss/gauss/v6/x/nft/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	// Methods imported from bank should be defined here
}

type NftKeeper interface {
	SaveNFT(ctx sdk.Context, nft *nfttype.Nft) error
	SetOwnerNFT(ctx sdk.Context, cateId string, tokenId string, owner string) error
	DeleteOwnerNFT(ctx sdk.Context, cateId string, tokenId string, owner string)
	GetNFT(ctx sdk.Context, cateId, tokenId string) (*nfttype.Nft, error)
	LockNFT(ctx sdk.Context, cateId, tokenId string) (err error)
	FreeNFT(ctx sdk.Context, cateId, tokenId string) error
	CheckNFT(nft nfttype.Nft) error
	GetLastPrice(ctx sdk.Context, tokenId string, denom string) sdk.Coin
	SaveLastPrice(ctx sdk.Context, tokenId string, price sdk.Coin) error
	Transfer(ctx sdk.Context, from, to string, nft nfttype.Nft, transferType uint64) error
}
