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

func (k Keeper) Collections(goCtx context.Context, req *types.QueryCollectionsRequest) (*types.QueryCollectionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nfts []*types.Nft
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.PrefixNFT))

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
	//var collections = k.GetCollections(ctx)
	//
	//var collections2 []*types.Collection
	//
	//for i := 0; i < len(collections); i++ {
	//	var collection types.Collection
	//	collection = types.Collection{
	//		CateId: collections[i].CateId,
	//		NFTs:   collections[i].NFTs,
	//	}
	//	collections2 = append(collections2, &collection)
	//}

	return &types.QueryCollectionsResponse{Nfts: nfts, Pagination: pageRes}, nil
}
