package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Collection(goCtx context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.PrefixNFT+req.CateId))
	var nfts []*types.Nft
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var nft types.Nft
		if err := k.cdc.Unmarshal(value, &nft); err != nil {
			return err
		}

		nfts = append(nfts, &nft)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//collection, found := k.GetCollection(ctx, req.CateId)
	//if found {
	//	return &types.QueryCollectionResponse{Collection: &collection}, nil
	//}
	return &types.QueryCollectionResponse{Nfts: nfts, Pagination: pageRes}, nil
}
