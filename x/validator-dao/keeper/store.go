package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/icplaza/icplaza/v6/x/validator-dao/types"
)

func (k Keeper) SetAuthorizerBizs(ctx sdk.Context, authorizerAddr sdk.AccAddress, bizs types.AuthorizerBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bizs)
	store.Set(types.GetAuthorizerBizsKey(authorizerAddr), bz)
}

func (k Keeper) GetAuthorizerBizs(ctx sdk.Context, authorizerAddr sdk.AccAddress) (authorizerBizs types.AuthorizerBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetAuthorizerBizsKey(authorizerAddr))
	if bz == nil {
		return types.AuthorizerBizs{}
	}
	k.cdc.Unmarshal(bz, &authorizerBizs)
	return authorizerBizs
}

func (k Keeper) RemoveAuthorizerBizs(ctx sdk.Context, authorizerAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAuthorizerBizsKey(authorizerAddr))
}

func (k Keeper) SetGranteeAuthBizs(ctx sdk.Context, granteeAddr sdk.AccAddress, granteeAuthBizs types.GranteeBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&granteeAuthBizs)
	store.Set(types.GetGranteeAuthBizsKey(granteeAddr), bz)
}

func (k Keeper) GetGranteeAuthBizs(ctx sdk.Context, granteeAddr sdk.AccAddress) (granteeAuthBizs types.GranteeBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetGranteeAuthBizsKey(granteeAddr))
	if bz == nil {
		return types.GranteeBizs{Grantee: granteeAddr.String()}
	}
	k.cdc.Unmarshal(bz, &granteeAuthBizs)
	return granteeAuthBizs
}

func (k Keeper) removeGranteeAuthBizs(ctx sdk.Context, granteeAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetGranteeAuthBizsKey(granteeAddr))
}

func (k Keeper) GetAllAuthorizerBizs(ctx sdk.Context) []types.AuthorizerBizs {
	var rc []types.AuthorizerBizs

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AuthorizerBizsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.AuthorizerBizs
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		rc = append(rc, data)
	}
	return rc
}

func (k Keeper) GetAllGranteeAuthBizs(ctx sdk.Context) []types.GranteeBizs {
	var rc []types.GranteeBizs

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GranteeAuthBizsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.GranteeBizs
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		rc = append(rc, data)
	}
	return rc
}
