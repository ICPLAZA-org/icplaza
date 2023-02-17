package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/blindbox/types"
)

func (k Keeper) CreateGroup(ctx sdk.Context, msg types.MsgCreateBox) error {
	if !k.PoolExists(ctx, msg.PoolId) {
		return sdkerrors.Wrapf(types.InvalidParam, "pool(%s) not exists", msg.PoolId)
	}
	if k.GroupExists(ctx, msg.GroupId) {
		return sdkerrors.Wrapf(types.InvalidParam, "group(%s) exists", msg.GroupId)
	}
	// save new group
	data := types.Group{
		GroupId:        msg.GroupId,
		BoxSize:        msg.BoxSize,
		BoxPrice:       msg.BoxPrice,
		OpenSize:       0,
		StartTime:      msg.StartTime,
		EndTime:        msg.EndTime,
		Creator:        msg.Creator,
		PoolId:         msg.PoolId,
		RandomMin:      msg.RandomMin,
		RandomMax:      msg.RandomMax,
		RandomNfts:     msg.RandomNfts,
		FixedNfts:      msg.FixedNfts,
		LeftRandomNfts: msg.RandomNfts,
		LeftFixedNfts:  msg.FixedNfts,
	}
	err := k.SaveGroup(ctx, data)
	if err != nil {
		return err
	}
	k.saveCreatorToGroupId(ctx, msg.Creator, msg.GroupId)
	k.savePoolIdToGroupId(ctx, msg.PoolId, msg.GroupId)
	return nil
}

func (k Keeper) GetGroupInfo(ctx sdk.Context, groupId string) (*types.Group, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfGroup(groupId)
	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.InvalidParam, "group(%s) not exists", groupId)
	}
	var group types.Group
	err := k.cdc.Unmarshal(bz, &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (k Keeper) GroupExists(ctx sdk.Context, groupId string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfGroup(groupId)
	return store.Has(key)
}

func (k Keeper) SaveGroup(ctx sdk.Context, data types.Group) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfGroup(data.GroupId)
	bz, err := k.cdc.Marshal(&data)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

func (k Keeper) saveCreatorToGroupId(ctx sdk.Context, creator string, groupId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfCreatorToGroupId(creator, groupId)
	bz := []byte(groupId)
	store.Set(key, bz)
}

func (k Keeper) removeCreatorToGroupId(ctx sdk.Context, creator string, groupId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfCreatorToGroupId(creator, groupId)
	store.Delete(key)
}

func (k Keeper) savePoolIdToGroupId(ctx sdk.Context, poolId string, groupId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfPoolIdToGroupId(poolId, groupId)
	bz := []byte(groupId)
	store.Set(key, bz)
}

func (k Keeper) removePoolIdToGroupId(ctx sdk.Context, poolId string, groupId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfPoolIdToGroupId(poolId, groupId)
	store.Delete(key)
}

func (k Keeper) IncreaseGroupOpenedSize(ctx sdk.Context, groupId string, addAmt uint64) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfGroup(groupId)
	bz := store.Get(key)
	if bz == nil {
		return sdkerrors.Wrapf(types.InvalidParam, "group(%s) not exists", groupId)
	}
	var group types.Group
	err := k.cdc.Unmarshal(bz, &group)
	if err != nil {
		return err
	}
	group.OpenSize += addAmt
	bz, err = k.cdc.Marshal(&group)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

//func (k Keeper) RemoveGroup(ctx sdk.Context, groupId string) error {
//	store := ctx.KVStore(k.storeKey)
//	key := types.KeyOfGroup(groupId)
//	bz := store.Get(key)
//	if bz == nil {
//		return sdkerrors.Wrapf(types.InvalidParam, "groupId(%s) not exits", groupId)
//	}
//	var group types.Group
//	err := k.cdc.Unmarshal(bz, &group)
//	if err != nil {
//		return err
//	}
//	k.removeCreatorToGroupId(ctx, group.Creator, group.GroupId)
//	k.removePoolIdToGroupId(ctx, group.PoolId, group.GroupId)
//	if len(group.LeftFixedNfts) > 0 {
//		err = k.freeNFTs(ctx, group.LeftFixedNfts)
//		if err != nil {
//			return err
//		}
//	}
//	if len(group.LeftRandomNfts) > 0 {
//		err = k.freeNFTs(ctx, group.LeftRandomNfts)
//		if err != nil {
//			return err
//		}
//	}
//	store.Delete(key)
//	return nil
//}

func (k Keeper) freeNFTs(ctx sdk.Context, tokens []string) error {
	for _, token := range tokens {
		err := k.nftKeeper.FreeNFT(ctx, "", token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) RevokeGroup(ctx sdk.Context, group *types.Group) error {
	group.OpenSize = group.BoxSize
	err := k.SaveGroup(ctx, *group)
	if err != nil {
		return err
	}
	boxes := k.GetBoxesByGroupId(ctx, group.GroupId)
	for _, box := range boxes {
		if box.Opened == false {
			box.Opened = true
			err := k.saveBox(ctx, *box)
			if err != nil {
				return err
			}
		}
	}
	// free left nft
	if len(group.LeftFixedNfts) > 0 {
		err = k.freeNFTs(ctx, group.LeftFixedNfts)
		if err != nil {
			return err
		}
	}
	if len(group.LeftRandomNfts) > 0 {
		err = k.freeNFTs(ctx, group.LeftRandomNfts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) RevokeGroupById(ctx sdk.Context, id string) error {
	groupInfo, err := k.GetGroupInfo(ctx, id)
	if err != nil {
		return err
	}
	return k.RevokeGroup(ctx, groupInfo)
}
