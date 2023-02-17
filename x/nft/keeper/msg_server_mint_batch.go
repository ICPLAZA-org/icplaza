package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/nft/types"
)

func (k msgServer) MintBatch(goCtx context.Context, msg *types.MsgMintBatch) (*types.MsgMintBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cates := make(map[string]string)

	var nfts []types.Nft
	for _, item := range msg.Items {
		cates[item.CateId] = item.CateId
		nft := types.Nft{
			TokenId:       item.TokenId,
			CateId:        item.CateId,
			Owner:         msg.Recipient,
			Creator:       msg.Recipient,
			CompanyUri:    item.CompanyUri,
			TokenUri:      item.TokenUri,
			ValueAddedTax: item.ValueAddedTax,
			CopyrightTax:  item.CopyrightTax,
			TokenStatus:   types.NFTFree,
			Name:          item.Name,
			Components:    item.Components,
		}
		nfts = append(nfts, nft)
	}
	err := k.saveBatch(ctx, nfts, cates)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintBatchResponse{}, nil
}

func (k msgServer) saveBatch(ctx sdk.Context, nfts []types.Nft, cates map[string]string) error {
	// save some gas
	for _, cate := range cates {
		err := k.SaveCate(ctx, cate)
		if err != nil {
			return err
		}
	}
	for _, nft := range nfts {
		if k.Exists(ctx, nft.TokenId) {
			return sdkerrors.Wrap(types.ErrNFTAlreadyExists, "tokenId="+nft.TokenId+" already exists ")
		}
		err := k.SaveNFT(ctx, &nft)
		if err != nil {
			return err
		}
		err = k.SaveTokenID(ctx, nft.CateId, nft.TokenId)
		if err != nil {
			return err
		}
		err = k.SetOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
		if err != nil {
			return err
		}
		err = k.SetCreatorNFT(ctx, nft.CateId, nft.TokenId, nft.Creator)
		if err != nil {
			return err
		}
		err = k.SetNameTokenId(ctx, nft.Name, nft.TokenId, nft.Owner)
		if err != nil {
			return err
		}
	}
	return nil
}
