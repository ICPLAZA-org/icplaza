package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Nftoken(goCtx context.Context, req *types.QueryNftokenRequest) (*types.QueryNftokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cateId := k.GetCateIdByTokenId(ctx, req.TokenId)
	if cateId == "" {
		return &types.QueryNftokenResponse{Nft: nil}, nil
	} else {
		nft, err := k.GetNFT(ctx, cateId, req.TokenId)
		if err != nil {
			return nil, err
		}
		return &types.QueryNftokenResponse{Nft: nft}, nil
	}
}
