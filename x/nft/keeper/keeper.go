package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    sdk.StoreKey
		memKey      sdk.StoreKey
		bankKeeper  types.BankKeeper
		distrKeeper types.DistributionKeeper
		// this line is used by starport scaffolding # ibc/keeper/attribute
		accountKeeper types.AccountKeeper
		paramSpace    paramstypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistributionKeeper,
	accountKeeper types.AccountKeeper,
	paramSpace paramstypes.Subspace,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		bankKeeper:    bankKeeper,
		distrKeeper:   distrKeeper,
		accountKeeper: accountKeeper,
		paramSpace:    paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams gets the parameters for the liquidity module.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams sets the parameters for the liquidity module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetLastPrices gets all token's last price
func (k Keeper) GetLastPrices(ctx sdk.Context) []*types.LastPrice {
	var lastPrices []*types.LastPrice
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixLastPrice))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var lastPrice types.LastPrice
		k.cdc.MustUnmarshal(iterator.Value(), &lastPrice)
		lastPrices = append(lastPrices, &lastPrice)
	}
	return lastPrices
}

// SetLastPrices sets all token's last price
func (k Keeper) SetLastPrices(ctx sdk.Context, lastPrices []*types.LastPrice) {
	store := ctx.KVStore(k.storeKey)
	for _, lastPrice := range lastPrices {
		key := types.KeyPrefix(types.PrefixLastPrice + lastPrice.TokenId)
		bz, err := k.cdc.Marshal(lastPrice)
		if err != nil {
			store.Set(key, bz)
		}
	}
}

// GetNfts gets all nft
func (k Keeper) GetNfts(ctx sdk.Context) []*types.Nft {
	var nfts []*types.Nft
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixNFT))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var nft types.Nft
		k.cdc.MustUnmarshal(iterator.Value(), &nft)
		nfts = append(nfts, &nft)
	}
	return nfts
}

// SetNfts sets all nfts
func (k Keeper) SetNfts(ctx sdk.Context, nfts []*types.Nft) {
	for _, nft := range nfts {
		// save nft
		err := k.SaveNFT(ctx, nft)
		if err != nil {
			panic(err)
		}
		// save cate
		err = k.SaveCate(ctx, nft.CateId)
		if err != nil {
			panic(err)
		}
		// save TokenID
		err = k.SaveTokenID(ctx, nft.CateId, nft.TokenId)
		if err != nil {
			panic(err)
		}
		// save owner nft
		err = k.SetOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
		if err != nil {
			panic(err)
		}
		// save SetCreatorNFT
		err = k.SetCreatorNFT(ctx, nft.CateId, nft.TokenId, nft.Creator)
		if err != nil {
			panic(err)
		}
	}
}

