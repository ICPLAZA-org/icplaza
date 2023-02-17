package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

// check exists
func (k Keeper) PoolExists(ctx sdk.Context, poolAddress string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixPool + poolAddress)
	bz := store.Get(key)
	if bz == nil {
		return false
	}
	return true
}

// check pool
func (k Keeper) CheckPool(ctx sdk.Context, poolAddress string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixPool + poolAddress)
	bz := store.Get(key)
	if bz != nil {
		return types.ErrInvalidPoolExists
	}
	return nil
}

// save pool
func (k Keeper) createPool(ctx sdk.Context, pool types.Pool) error {
	// check
	err := k.CheckPool(ctx, pool.PoolAddress)
	if err != nil {
		return err
	}
	err = k.savePool(ctx, pool)
	if err != nil {
		return err
	}
	return nil
}

// save pool
func (k Keeper) savePool(ctx sdk.Context, pool types.Pool) error {
	// save
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixPool + pool.PoolAddress)
	// marshal
	bz, err := k.cdc.Marshal(&pool)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// update pool
func (k Keeper) updatePool(ctx sdk.Context, pool types.Pool) error {
	oldPool, err := k.GetPool(ctx, pool.PoolAddress)
	if err != nil {
		return err
	}
	if oldPool.Creator != pool.Creator {
		return types.ErrInvalidCreator
	}
	// save
	err = k.savePool(ctx, pool)
	if err != nil {
		return err
	}
	return nil
}

// get pool
func (k Keeper) GetPool(ctx sdk.Context, poolAddress string) (*types.Pool, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixPool + poolAddress)
	bz := store.Get(key)
	if bz == nil {
		return nil, types.ErrInvalidPoolNotExists
	}
	var pool types.Pool
	err := k.cdc.Unmarshal(bz, &pool)
	if err != nil {
		return nil, err
	}
	return &pool, nil
}

// add log
func (k msgServer) LogAddCoin(ctx sdk.Context, poolAddr string, delegator string, coin sdk.Coin) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixDelegate + poolAddr + delegator)
	bz := store.Get(key)
	var delegation types.Delegation
	if bz == nil {
		var coins []sdk.Coin
		coins = append(coins, coin)
		delegation = types.Delegation{Coins: coins, PoolAddress: poolAddr, Delegator: delegator}
	} else {
		var found bool
		err := k.cdc.Unmarshal(bz, &delegation)
		if err != nil {
			return err
		}
		for index, coinV := range delegation.Coins {
			if coinV.Denom == coin.Denom {
				delegation.Coins[index] = delegation.Coins[index].Add(coin)
				found = true
			}
		}
		if !found {
			delegation.Coins = append(delegation.Coins, coin)
		}
	}
	// save
	bz, err := k.cdc.Marshal(&delegation)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// log
func (k msgServer) LogSubCoin(ctx sdk.Context, poolAddr string, delegator string, coin sdk.Coin) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixDelegate + poolAddr + delegator)
	bz := store.Get(key)
	var delegation types.Delegation
	if bz == nil {
		return types.ErrNoDelegation
	} else {
		var found bool
		k.cdc.MustUnmarshal(bz, &delegation)
		for index, coinV := range delegation.Coins {
			if coinV.Denom == coin.Denom {
				if coinV.IsLT(coin) {
					return types.ErrInsufficientFunds
				}
				delegation.Coins[index] = delegation.Coins[index].Sub(coin)
				found = true
			}
		}
		if !found {
			return sdkerrors.Wrap(types.ErrNoDelegation, coin.Denom)
		}
	}
	// save
	bz, err := k.cdc.Marshal(&delegation)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// delegation
func (k msgServer) GetDelegationBalance(ctx sdk.Context, poolAddr string, delegator string, denom string) (sdk.Coin, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixDelegate + poolAddr + delegator)
	bz := store.Get(key)
	var delegation types.Delegation
	if bz == nil {
		return sdk.Coin{}, types.ErrNoDelegation
	} else {
		k.cdc.MustUnmarshal(bz, &delegation)
		for _, coinV := range delegation.Coins {
			if coinV.Denom == denom {
				return coinV, nil
			}
		}
	}
	return sdk.Coin{}, types.ErrNoDelegation
}

// all pools
func (k Keeper) AllPools(ctx sdk.Context) []*types.Pool {
	var pools []*types.Pool
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixPool))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var pool types.Pool
		k.cdc.MustUnmarshal(iterator.Value(), &pool)
		pools = append(pools, &pool)
	}
	return pools
}

// all delegate logs
func (k Keeper) AllDelegations(ctx sdk.Context) []*types.Delegation {
	var logs []*types.Delegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixDelegate))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var log types.Delegation
		var delegation types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		logs = append(logs, &log)
	}
	return logs
}
