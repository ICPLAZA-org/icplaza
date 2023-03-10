package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v11/x/nftexpool/types"
)

func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pool = types.Pool{
		Creator:              msg.Creator,
		PoolAddress:          msg.PoolAddress,
		CommissionRate:       msg.CommissionRate,
		CommissionAddress:    msg.CommissionAddress,
		ValueAddedTaxAddress: msg.ValueAddedTaxAddress,
	}
	err := k.updatePool(ctx, pool)
	if err != nil {
		return nil, err
	}
	return &types.MsgUpdatePoolResponse{}, nil
}
