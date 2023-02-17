package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gauss/gauss/v6/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Creator(goCtx context.Context, req *types.QueryCreatorRequest) (*types.QueryCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err
	}

	var key []byte
	store := ctx.KVStore(k.storeKey)
	if req.CateId == "" {
		key = types.KeyPrefix(types.PrefixOwner + req.Creator)
	} else {
		key = types.KeyPrefix(types.PrefixOwner + req.Creator + req.CateId)
	}
	postStore := prefix.NewStore(store, key)

	var nfts []*types.Nft
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.NfToken
		if err := k.cdc.Unmarshal(value, &token); err != nil {
			return err
		}
		cateId := ""
		if req.CateId == "" {
			cateId = k.GetCateIdByTokenId(ctx, token.TokenId)
		} else {
			cateId = req.CateId
		}
		if cateId != "" {
			var nft *types.Nft
			nft, err := k.GetNFT(ctx, cateId, token.TokenId)
			if err == nil {
				nfts = append(nfts, nft)
			}
		}
		return nil
	})

	return &types.QueryCreatorResponse{Nfts: nfts, Pagination: pageRes}, nil
}
