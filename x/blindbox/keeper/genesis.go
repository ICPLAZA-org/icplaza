package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/blindbox/types"
)

// use for genesis
func (k Keeper) ExportGroupBoxes(ctx sdk.Context) []*types.GroupBox {
	store := ctx.KVStore(k.storeKey)
	str := fmt.Sprintf(types.PrefixKeyOfGroup, "")
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(str))
	var list []*types.GroupBox
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var group types.Group
		err := k.cdc.Unmarshal(iterator.Value(), &group)
		if err != nil {
			panic(err)
		}
		boxes := k.GetBoxesByGroupId(ctx, group.GroupId)
		item := types.GroupBox{
			Group: &group,
			Boxes: boxes,
		}
		list = append(list, &item)
	}
	return list
}

func (k Keeper) ImportGroupBoxes(ctx sdk.Context, list []*types.GroupBox) {
	for _, item := range list {
		// save group
		err := k.SaveGroup(ctx, *item.Group)
		if err != nil {
			panic(err)
		}
		k.saveCreatorToGroupId(ctx, item.Group.Creator, item.Group.GroupId)
		k.savePoolIdToGroupId(ctx, item.Group.PoolId, item.Group.GroupId)
		// save boxes
		err = k.CreateBoxes(ctx, item.Boxes)
		if err != nil {
			panic(err)
		}
	}
}

func (k Keeper) ExportPools(ctx sdk.Context) []*types.Pool {
	store := ctx.KVStore(k.storeKey)
	str := fmt.Sprintf(types.PrefixKeyOfPool, "")
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(str))
	var list []*types.Pool
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var item types.Pool
		err := k.cdc.Unmarshal(iterator.Value(), &item)
		if err != nil {
			panic(err)
		}
		list = append(list, &item)
	}
	return list
}

func (k Keeper) ImportPools(ctx sdk.Context, list []*types.Pool) {
	for _, item := range list {
		err := k.savePool(ctx, *item)
		if err != nil {
			panic(err)
		}
	}
}
