package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
)

// 冻结NFT可以只传递tokenId
func (k msgServer) FrozenNft(goCtx context.Context, msg *types.MsgFrozenNft) (*types.MsgFrozenNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cateId := msg.CateId
	if cateId == "" {
		cateId = k.GetCateIdByTokenId(ctx, msg.TokenId)
	}
	if cateId == "" {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, msg.CateId)
	}
	// check nft owner
	nft, err := k.GetNFT(ctx, cateId, msg.TokenId)
	if err != nil {
		return nil, err
	}
	if nft.Owner != msg.Sender {
		return nil, sdkerrors.Wrap(types.ErrInvalidOwner, msg.Sender+" is not the nft["+msg.TokenId+"] owner ")
	}
	// frozen
	err = k.FrozenNFT(ctx, cateId, msg.TokenId)
	if err != nil {
		return nil, err
	}
	return &types.MsgFrozenNftResponse{}, nil
}
