package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gauss/gauss/v6/x/fixedprice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Orders(goCtx context.Context, req *types.QueryOrdersRequest) (*types.QueryOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var prefixKey []byte
	prefixKey = types.KeyPrefix(types.PrefixOrder + req.PoolAddress)
	return k.getOrders(goCtx, req, prefixKey)
}

func (k Keeper) getOrders(goCtx context.Context, req *types.QueryOrdersRequest, prefixKey []byte) (*types.QueryOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var orders []*types.Order
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, prefixKey)

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var order types.Order
		if err := k.cdc.Unmarshal(value, &order); err != nil {
			return err
		}

		orders = append(orders, &order)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryOrdersResponse{
		Orders:     orders,
		Pagination: pageRes,
	}, nil
}
