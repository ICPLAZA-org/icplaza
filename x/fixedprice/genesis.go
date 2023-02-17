package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/fixedprice/keeper"
	"github.com/gauss/gauss/v6/x/fixedprice/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	orders := k.GetAllOrders(ctx)
	params := types.Params{AdjustPricePeriod: k.GetParams(ctx).AdjustPricePeriod, Orders: orders}
	return types.NewGenesis(params)
}
