package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
)

func (k msgServer) TransferNft(goCtx context.Context, msg *types.MsgTransferNft) (*types.MsgTransferNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cateId := msg.CateId
	if cateId == "" {
		cateId = k.GetCateIdByTokenId(ctx, msg.TokenId)
	}
	if cateId == "" {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, msg.CateId)
	}
	nft, err := k.GetNFT(ctx, cateId, msg.TokenId)
	if err != nil {
		return nil, err
	}
	// check owner
	if nft.Owner != msg.Sender {
		return nil, sdkerrors.Wrap(types.ErrInvalidOwner, msg.Sender)
	}
	// check status
	// if nft.TokenStatus == types.NFTLocked {
	//	return nil, sdkerrors.Wrap(types.ErrNFTLocked, msg.Sender)
	// }
	// if nft.TokenStatus == types.NFTFrozen {
	//	return nil, sdkerrors.Wrap(types.ErrNFTFrozen, msg.Sender)
	// }
	if nft.TokenStatus != types.NFTFree {
		return nil, sdkerrors.Wrap(types.ErrNotFree, nft.TokenId)
	}
	// delete old
	k.DeleteOwnerNFT(ctx, cateId, msg.TokenId, nft.Owner)
	// name tokenId
	k.DeleteNameTokenId(ctx, nft.Name, nft.Owner)
	// save new
	nft.Owner = msg.Recipient
	err = k.SaveNFT(ctx, nft)
	if err != nil {
		return nil, err
	}
	// save new owner info
	err = k.SetOwnerNFT(ctx, cateId, msg.TokenId, msg.Recipient)
	if err != nil {
		return nil, err
	}
	// save new name => tokenId
	err = k.SetNameTokenId(ctx, nft.Name, nft.TokenId, nft.Owner)
	if err != nil {
		return nil, err
	}
	// transfer
	var isCNFT = false
	if len(nft.Components) > 0 {
		isCNFT = true
	}
	// tag transfered cnft
	if isCNFT {
		k.tagTransferedCNFT(ctx, nft.TokenId, 0)
	}
	return &types.MsgTransferNftResponse{}, nil
}
