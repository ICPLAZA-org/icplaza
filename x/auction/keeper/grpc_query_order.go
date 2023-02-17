package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/auction/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Order(goCtx context.Context, req *types.QueryOrderRequest) (*types.QueryOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	poolAddr, err := k.getPoolAddress(ctx, req.TokenId)
	if err != nil {
		return nil, err
	}
	order, err := k.getOrder(ctx, req.TokenId, poolAddr)
	if err != nil {
		return nil, err
	}
	if order != nil {
		return &types.QueryOrderResponse{Order: order}, nil
	}
	return &types.QueryOrderResponse{Order: nil}, nil
}

// get pool address
func (k Keeper) getPoolAddress(ctx sdk.Context, tokenId string) (string, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrderTokenIdToPool + tokenId)
	bz := store.Get(key)
	if bz == nil {
		return "", nil
	}
	var pool types.PoolAddress
	err := k.cdc.Unmarshal(bz, &pool)
	if err != nil {
		return "", err
	}
	return pool.Address, nil
}

// get order
func (k Keeper) getOrder(ctx sdk.Context, tokenId, poolAddr string) (*types.Order, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrder + poolAddr + tokenId)
	bz := store.Get(key)
	if bz == nil {
		return nil, nil
	}
	var order types.Order
	err := k.cdc.Unmarshal(bz, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
