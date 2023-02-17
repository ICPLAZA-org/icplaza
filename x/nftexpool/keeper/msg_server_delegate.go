package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := k.GetPool(ctx, msg.PoolAddress)
	if err != nil {
		return nil, err
	}
	delegator, _ := sdk.AccAddressFromBech32(msg.Delegator)
	balance := k.bankKeeper.GetBalance(ctx, delegator, msg.Amount.Denom)
	if balance.IsLT(msg.Amount) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientFunds, balance.String()+" is less than "+msg.Amount.String())
	}
	// 委托到模块
	err = k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, delegator, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}
	// 记录
	err = k.LogAddCoin(ctx, msg.PoolAddress, msg.Delegator, msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgDelegateResponse{}, nil
}
