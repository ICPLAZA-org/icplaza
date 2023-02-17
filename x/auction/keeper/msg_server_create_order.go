package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/auction/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	if msg.AutoAgreePeriod.Seconds() < params.GetAutoAgreePeriod().Seconds() {
		return nil, sdkerrors.Wrap(types.ErrInvalidParam, fmt.Sprintf("AutoAgreePeriod(%s) should greater or equal than %s", msg.AutoAgreePeriod, params.GetAutoAgreePeriod()))
	}
	// get nft
	nft, err := k.nftKeeper.GetNFT(ctx, "", msg.TokenId)
	if err != nil {
		return nil, err
	}
	// check nft status
	err = k.nftKeeper.CheckNFT(*nft)
	if err != nil {
		return nil, err
	}
	// check nft owner
	if nft.Owner != msg.Sender {
		return nil, sdkerrors.Wrap(types.ErrInvalidNftOwner, msg.Sender)
	}
	// check pool
	exists := k.poolKeeper.PoolExists(ctx, msg.PoolAddress)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrInvalidPool, msg.PoolAddress)
	}
	// create order
	// 1 lock nft
	err = k.nftKeeper.LockNFT(ctx, nft.CateId, nft.TokenId)
	if err != nil {
		return nil, err
	}
	// 2 save order info
	var order = types.Order{
		TokenId:         nft.TokenId,
		Price:           &msg.StartPrice,
		BidPrice:        sdk.Coin{},
		Nft:             nft,
		Bidder:          "",
		StartTime:       msg.StartTime,
		EndTime:         msg.EndTime,
		MinEndTime:      time.Time{},
		MinStep:         msg.MinStep,
		PoolAddress:     msg.PoolAddress,
		AutoAgreePeriod: msg.AutoAgreePeriod,
	}
	err = k.saveOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	// save tokenId to poolAddress
	err = k.saveTokenIdToPoolAddress(ctx, nft.TokenId, msg.PoolAddress)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateOrderResponse{}, nil
}

// save tokenId to poolAddress
func (k msgServer) saveTokenIdToPoolAddress(ctx sdk.Context, tokenId, poolAddress string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrderTokenIdToPool + tokenId)
	poolAddr := &types.PoolAddress{
		Address: poolAddress,
	}
	bz, err := k.cdc.Marshal(poolAddr)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}
