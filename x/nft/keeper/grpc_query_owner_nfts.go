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

func (k Keeper) OwnerNfts(goCtx context.Context, req *types.QueryNftsByNameOrTokenRequest) (*types.QueryNftsByNameOrTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.TokenId == "" && req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request tokenId and Name cannot be all empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	if req.TokenId != "" {
		nft, err := k.GetNFTByTokenId(ctx, req.TokenId)
		if err != nil {
			return nil, err
		}
		var nfts []*types.Nft
		nfts = append(nfts, nft)
		return &types.QueryNftsByNameOrTokenResponse{Nfts: nfts}, nil
	} else {
		return k.getNftsByName(ctx, req)
	}
}

func (k Keeper) getNftsByName(ctx sdk.Context, req *types.QueryNftsByNameOrTokenRequest) (*types.QueryNftsByNameOrTokenResponse, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKeyOfTokenName(req.Owner, req.Name)
	postStore := prefix.NewStore(store, key)
	var nfts []*types.Nft
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.NfToken
		if err := k.cdc.Unmarshal(value, &token); err != nil {
			return err
		}
		var nft *types.Nft
		nft, err := k.GetNFTByTokenId(ctx, token.TokenId)
		if err == nil {
			nfts = append(nfts, nft)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryNftsByNameOrTokenResponse{Nfts: nfts, Pagination: pageRes}, nil
}
