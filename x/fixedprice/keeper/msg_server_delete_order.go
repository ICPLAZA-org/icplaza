package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/fixedprice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteOrder(goCtx context.Context, msg *types.MsgDeleteOrder) (*types.MsgDeleteOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, err := k.getOrder(ctx, msg.PoolAddress, msg.TokenId)
	if err != nil {
		return nil, err
	}
	// check order creator
	if order.Nft.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrInvalidParam, msg.Creator+" is not order creator")
	}
	// unlock nft and remove order
	err = k.nftKeeper.FreeNFT(ctx, order.Nft.CateId, msg.TokenId)
	if err != nil {
		return nil, err
	}
	k.RemoveOrder(ctx, msg.PoolAddress, msg.TokenId)
	return &types.MsgDeleteOrderResponse{}, nil
}
