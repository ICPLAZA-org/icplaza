package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v11/x/auction/keeper"
	"github.com/evmos/evmos/v11/x/auction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	orders := k.GetAllOrders(ctx)
	params := types.Params{AutoAgreePeriod: k.GetParams(ctx).AutoAgreePeriod, Orders: orders}
	return types.NewGenesis(params)
}
