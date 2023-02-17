package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gauss/gauss/v6/x/blindbox/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetBox(goCtx context.Context, req *types.QueryGetBoxRequest) (*types.QueryGetBoxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	group, _ := k.GetGroupInfo(ctx, req.GroupId)
	box, _ := k.GetBoxInfo(ctx, req.GroupId, req.BoxId)
	if box != nil {
		return &types.QueryGetBoxResponse{
			Box:   box,
			Group: group,
		}, nil
	}

	return &types.QueryGetBoxResponse{}, nil
}
func (k Keeper) GetGroup(goCtx context.Context, req *types.QueryGetGroupRequest) (*types.QueryGetGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	group, err := k.GetGroupInfo(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	boxes := k.GetBoxesByGroupId(ctx, req.GroupId)

	if group != nil {
		return &types.QueryGetGroupResponse{
			Group: group,
			Boxes: boxes,
		}, nil
	}

	return &types.QueryGetGroupResponse{}, nil
}
func (k Keeper) GetGroups(goCtx context.Context, req *types.QueryGetGroupsRequest) (*types.QueryGetGroupsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	str := fmt.Sprintf(types.PrefixKeyOfPoolIdToGroupId, req.PoolId, "")
	postStore := prefix.NewStore(store, []byte(str))
	var list []*types.QueryGetGroupResponse
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		groupId := string(value)
		group, _ := k.GetGroupInfo(ctx, groupId)
		boxes := k.GetBoxesByGroupId(ctx, groupId)
		item := types.QueryGetGroupResponse{
			Group: group,
			Boxes: boxes,
		}
		list = append(list, &item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryGetGroupsResponse{
		Groups:     list,
		Pagination: pageRes,
	}, nil
}
func (k Keeper) GetPool(goCtx context.Context, req *types.QueryGetPoolRequest) (*types.QueryGetPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	found, pool := k.GetPoolById(ctx, req.PoolId)
	if !found {
		return &types.QueryGetPoolResponse{}, nil
	}

	return &types.QueryGetPoolResponse{
		Pool: pool,
	}, nil
}
func (k Keeper) GetPools(goCtx context.Context, req *types.QueryGetPoolsRequest) (*types.QueryGetPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	str := fmt.Sprintf(types.PrefixKeyOfPool, "")
	postStore := prefix.NewStore(store, []byte(str))
	var list []*types.Pool
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var item types.Pool
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}
		list = append(list, &item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetPoolsResponse{
		Pools:      list,
		Pagination: pageRes,
	}, nil
}
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
