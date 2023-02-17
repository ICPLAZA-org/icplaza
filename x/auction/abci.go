package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/auction/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	_ = k.CloseExpired(ctx)
}
