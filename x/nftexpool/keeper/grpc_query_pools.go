package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QPools(goCtx context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var pools []*types.Pool
	ctx := sdk.UnwrapSDKContext(goCtx)

	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrefixPool))

	pageRes, err := query.Paginate(prefixStore, req.Pagination, func(key []byte, value []byte) error {

		var pool types.Pool
		if err := k.cdc.Unmarshal(value, &pool); err != nil {
			return err
		}

		pools = append(pools, &pool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPoolsResponse{
		Pools:      pools,
		Pagination: pageRes,
	}, nil
}
