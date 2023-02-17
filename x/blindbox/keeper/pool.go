package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/blindbox/types"
)

func (k Keeper) CreatePool(ctx sdk.Context, msg types.MsgCreatePool) error {
	found, _ := k.GetPoolById(ctx, msg.PoolId)
	if found {
		return sdkerrors.Wrapf(types.InvalidParam, "pool(%s) exists", msg.PoolId)
	}
	pool := types.Pool{
		PoolId:     msg.PoolId,
		FeeRate:    msg.FeeRate,
		FeeAddress: msg.FeeAddress,
		Creator:    msg.Creator,
	}
	return k.savePool(ctx, pool)
}

func (k Keeper) PoolExists(ctx sdk.Context, poolId string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfPool(poolId)
	return store.Has(key)
}

func (k Keeper) GetPoolById(ctx sdk.Context, poolId string) (bool, *types.Pool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfPool(poolId)
	bz := store.Get(key)
	if bz == nil {
		return false, nil
	}
	var pool types.Pool
	k.cdc.MustUnmarshal(bz, &pool)
	return true, &pool
}

func (k Keeper) UpdatePool(ctx sdk.Context, msg types.MsgUpdatePool) error {
	found, pool := k.GetPoolById(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(types.InvalidParam, "pool(%s) not exists", msg.PoolId)
	}
	if msg.Creator != pool.Creator {
		return sdkerrors.Wrapf(types.InvalidParam, "pool(%s) not owned by ", msg.PoolId, msg.Creator)
	}
	pool.FeeRate = msg.FeeRate
	pool.FeeAddress = msg.FeeAddress
	return k.savePool(ctx, *pool)
}

// TODO test
func (k Keeper) RemovePool(ctx sdk.Context, msg types.MsgRemovePool) error {
	found, pool := k.GetPoolById(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(types.InvalidParam, "pool(%s) not exists", msg.PoolId)
	}
	if pool.Creator != msg.Creator {
		return sdkerrors.Wrapf(types.InvalidParam, "invalid creator %s ", msg.Creator)
	}
	// check pool groups
	store := ctx.KVStore(k.storeKey)
	str := fmt.Sprintf(types.PrefixKeyOfPoolIdToGroupId, msg.PoolId, "")
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(str))
	defer iterator.Close()
	if iterator.Valid() {
		return sdkerrors.Wrapf(types.ErrRemove, "not blank pool")
	}
	// delete pool
	key := types.KeyOfPool(msg.PoolId)
	store.Delete(key)
	return nil
}

func (k Keeper) savePool(ctx sdk.Context, pool types.Pool) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfPool(pool.PoolId)
	bz, err := k.cdc.Marshal(&pool)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	k.saveCreatorToPoolId(ctx, pool.Creator, pool.PoolId)
	return nil
}

func (k Keeper) saveCreatorToPoolId(ctx sdk.Context, creator string, poolId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfCreatorToPoolId(creator, poolId)
	bz := []byte(poolId)
	store.Set(key, bz)
}
