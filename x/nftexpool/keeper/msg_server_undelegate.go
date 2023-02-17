package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

func (k msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, err := k.GetPool(ctx, msg.PoolAddress)
	if err != nil {
		return nil, err
	}

	if pool.Creator != msg.Sender {
		return nil, types.ErrInvalidPoolCreator
	}

	delegator, _ := sdk.AccAddressFromBech32(msg.Delegator)
	balance, err := k.GetDelegationBalance(ctx, msg.PoolAddress, msg.Delegator, msg.Amount.Denom)
	if err != nil {
		return nil, err
	}
	if balance.IsLT(msg.Amount) {
		return nil, types.ErrOutOfDelegation
	}

	// 从模块解除委托
	err = k.bankKeeper.UndelegateCoinsFromModuleToAccount(ctx, types.ModuleName, delegator, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}
	// 记录
	err = k.LogSubCoin(ctx, msg.PoolAddress, msg.Delegator, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgUndelegateResponse{}, nil
}
