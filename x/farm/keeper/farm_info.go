package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gauss/gauss/v6/x/farm/types"
)

// GetFarmer return the specified farmer
func (k Keeper) GetFarmerInfo(ctx sdk.Context, poolName, address string) (info types.Farmer, exist bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyFarmerInfo(address, poolName))
	if len(bz) == 0 {
		return info, false
	}

	k.cdc.MustUnmarshal(bz, &info)

	return info, true
}

func (k Keeper) IteratorFarmInfo(ctx sdk.Context, address string, fun func(farmer types.Farmer)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixFarmInfo(address))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var farmer types.Farmer
		k.cdc.MustUnmarshal(iterator.Value(), &farmer)
		fun(farmer)
	}
}

func (k Keeper) IteratorAllFarmInfo(ctx sdk.Context, fun func(farmer types.Farmer)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.FarmerKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var farmer types.Farmer
		k.cdc.MustUnmarshal(iterator.Value(), &farmer)
		fun(farmer)
	}
}

// SetFarmer save the farmer information
func (k Keeper) SetFarmerInfo(ctx sdk.Context, farmer types.Farmer) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&farmer)
	store.Set(types.KeyFarmerInfo(farmer.FarmerAddress, farmer.PoolName), bz)
}

func (k Keeper) DeleteFarmerInfo(ctx sdk.Context, poolName, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyFarmerInfo(address, poolName))
}

func (k Keeper) GetPoolFarmersCount(ctx sdk.Context) map[string]uint64 {
	m := make(map[string]uint64)
	k.IteratorAllFarmInfo(ctx, func(farmer types.Farmer) {
		m[farmer.PoolName]++
	})
	return m
}
