package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/validator-dao/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) AuthorizerBizs(goCtx context.Context, req *types.QueryAuthorizerBizsRequest) (*types.QueryAuthorizerBizsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	authorierAddr, err := sdk.AccAddressFromBech32(req.Authorizer)
	if err != nil {
		return nil, err
	}

	authBizs := k.GetAuthorizerBizs(ctx, authorierAddr)

	return &types.QueryAuthorizerBizsResponse{
		Bizs: authBizs.Bizs,
	}, nil
}

func (k Keeper) GranteeAuthBizs(goCtx context.Context, req *types.QueryGranteeAuthBizsRequest) (*types.QueryGranteeAuthBizsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	granteeAddr, err := sdk.AccAddressFromBech32(req.Grantee)
	if err != nil {
		return nil, err
	}

	authBizs := k.GetGranteeAuthBizs(ctx, granteeAddr)

	return &types.QueryGranteeAuthBizsResponse{
		Bizs: authBizs.Bizs,
	}, nil
}
