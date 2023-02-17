package blindbox

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/blindbox/keeper"
	"github.com/gauss/gauss/v6/x/blindbox/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.ImportPools(ctx, genState.Pools)
	k.ImportGroupBoxes(ctx, genState.GroupBoxes)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.GroupBoxes = k.ExportGroupBoxes(ctx)
	genesis.Pools = k.ExportPools(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
