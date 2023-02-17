package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/farm/types"
)

func (k Keeper) SetFarmersRewards(ctx sdk.Context, poolName string, farmersRewards types.FarmersRewards) {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixFarmersRewards(poolName)

	bz := k.cdc.MustMarshal(&farmersRewards)
	store.Set(key, bz)
}

func (k Keeper) GetFarmersRewards(ctx sdk.Context, poolName string) *types.FarmersRewards {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixFarmersRewards(poolName)

	bz := store.Get(key)
	if bz == nil {
		return nil
	}
	var farmersRewards types.FarmersRewards
	k.cdc.MustUnmarshal(bz, &farmersRewards)
	return &farmersRewards
}

func (k Keeper) IteratorAllFarmersRewards(ctx sdk.Context, fun func(farmer types.FarmersRewards)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.FarmersRewardsKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var farmer types.FarmersRewards
		k.cdc.MustUnmarshal(iterator.Value(), &farmer)
		fun(farmer)
	}
}
