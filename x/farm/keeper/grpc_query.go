package keeper

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/gauss/gauss/v6/x/farm/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) FarmPools(goctx context.Context, request *types.QueryFarmPoolsRequest) (*types.QueryFarmPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	farmersCount := k.GetPoolFarmersCount(ctx)
	var list []*types.FarmPoolEntry
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.FarmPoolKey)
	pageRes, err := query.Paginate(prefixStore, request.Pagination, func(_ []byte, value []byte) error {
		var pool types.FarmPool
		k.cdc.MustUnmarshal(value, &pool)

		var remainingReward sdk.Coins
		farmersRewards := k.GetFarmersRewards(ctx, pool.Name)
		if farmersRewards != nil {
			for _, r := range farmersRewards.FarmersRewards {
				remainingReward = remainingReward.Add(r.RemainingRewards)
			}
		}
		// farmer count
		pool.FarmerCount = farmersCount[pool.Name]
		list = append(list, &types.FarmPoolEntry{
			Pool:             pool,
			RemainingRewards: remainingReward,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryFarmPoolsResponse{
		Pools:      list,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) FarmPool(goctx context.Context,
	request *types.QueryFarmPoolRequest) (*types.QueryFarmPoolResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if len(request.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool name can not be empty")
	}
	ctx := sdk.UnwrapSDKContext(goctx)

	pool, exist := k.GetPool(ctx, request.Name)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, request.Name)
	}

	var remainingReward sdk.Coins
	farmersRewards := k.GetFarmersRewards(ctx, pool.Name)
	if farmersRewards != nil {
		for _, r := range farmersRewards.FarmersRewards {
			remainingReward.Add(r.RemainingRewards)
		}
	}
	// 计算池子委托人数
	farmersCount := k.GetPoolFarmersCount(ctx)
	pool.FarmerCount = farmersCount[pool.Name]

	poolEntry := &types.FarmPoolEntry{
		Pool:             pool,
		RemainingRewards: remainingReward,
	}
	return &types.QueryFarmPoolResponse{Pool: poolEntry}, nil
}

func (k Keeper) Farmer(goctx context.Context, request *types.QueryFarmerRequest) (*types.QueryFarmerResponse, error) {
	var list []*types.LockedList
	var err error
	var farmInfos []types.Farmer

	ctx := sdk.UnwrapSDKContext(goctx)
	cacheCtx, _ := ctx.CacheContext()
	if len(request.PoolName) == 0 {
		k.IteratorFarmInfo(cacheCtx, request.Farmer, func(farmInfo types.Farmer) {
			farmInfos = append(farmInfos, farmInfo)
		})
	} else {
		farmInfo, existed := k.GetFarmerInfo(cacheCtx, request.PoolName, request.Farmer)
		if existed {
			farmInfos = append(farmInfos, farmInfo)
		}
	}
	if len(farmInfos) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrFarmerNotFound, "not found farmer: %s", request.Farmer)
	}
	for _, farmer := range farmInfos {
		pool, exist := k.GetPool(cacheCtx, farmer.PoolName)
		if !exist {
			return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, farmer.PoolName)
		}

		//The farm pool has not started, no reward
		if pool.StartHeight > ctx.BlockHeight() {
			list = append(list, &types.LockedList{
				Farmer:         &farmer,
				PendingRewards: sdk.Coins{},
			})
			continue
		}

		if !k.Expired(ctx, pool) {
			pool, _, err = k.updatePool(cacheCtx, pool, sdk.ZeroInt(), false)
			if err != nil {
				return nil, err
			}
		}

		debets := k.GetFarmerRewards(cacheCtx, farmer.PoolName, request.Farmer)
		farmersRewards := k.GetFarmersRewards(cacheCtx, farmer.PoolName)

		rewards, _ := farmersRewards.CaclRewards(farmer, sdk.ZeroInt(), debets.Debts)
		// farmer赋值到tmpFarmer，用于指针引用
		tmpFarmer := farmer
		list = append(list, &types.LockedList{
			Farmer:         &tmpFarmer,
			PendingRewards: rewards,
		})
	}

	return &types.QueryFarmerResponse{
		List:   list,
		Height: ctx.BlockHeight(),
	}, nil
}

func (k Keeper) Params(goctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
