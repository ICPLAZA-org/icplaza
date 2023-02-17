package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QDelegate(goCtx context.Context, req *types.QueryDelegateRequest) (*types.QueryDelegateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixDelegate + req.PoolAddress + req.Delegator)
	bz := store.Get(key)
	if bz == nil {
		return &types.QueryDelegateResponse{}, nil
	}

	var delegation types.Delegation
	k.cdc.MustUnmarshal(bz, &delegation)

	var coins []*sdk.Coin
	for _, coin := range delegation.Coins {
		coins = append(coins, &coin)
	}

	return &types.QueryDelegateResponse{Coins: coins}, nil
}
