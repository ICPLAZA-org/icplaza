package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/farm/types"
)

func (k Keeper) SetFarmerRewards(ctx sdk.Context, poolName string, sender string, farmerRewards types.FarmerRewards) {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixFarmerRewards(poolName, sender)

	bz := k.cdc.MustMarshal(&farmerRewards)
	store.Set(key, bz)
}

func (k Keeper) GetFarmerRewards(ctx sdk.Context, poolName string, sender string) *types.FarmerRewards {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixFarmerRewards(poolName, sender)

	bz := store.Get(key)
	if bz == nil {
		return &types.FarmerRewards{Debts: sdk.Coins{}}
	}
	var farmerRewards types.FarmerRewards
	k.cdc.MustUnmarshal(bz, &farmerRewards)
	return &farmerRewards
}

func (k Keeper) IteratorAllFarmerRewards(ctx sdk.Context, fun func(farmer types.FarmerRewards)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.FarmerRewardsKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var farmer types.FarmerRewards
		k.cdc.MustUnmarshal(iterator.Value(), &farmer)
		fun(farmer)
	}
}
