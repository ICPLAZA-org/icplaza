package keeper

import (
	"fmt"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/auction/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramSpace paramstypes.Subspace
		nftKeeper  types.NftKeeper
		poolKeeper types.PoolKeeper
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	paramSpace paramstypes.Subspace,
	nft types.NftKeeper,
	pool types.PoolKeeper,
	bank types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: paramSpace,
		nftKeeper:  nft,
		poolKeeper: pool,
		bankKeeper: bank,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// delete tokenId to poolAddress
func (k Keeper) deleteOrderTokenIdToPoolAddress(ctx sdk.Context, tokenId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrderTokenIdToPool + tokenId)
	store.Delete(key)
}
