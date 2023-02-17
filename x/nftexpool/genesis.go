package pool

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nftexpool/keeper"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	k.SetPools(ctx, genState.Pools)
	k.SetDelegations(ctx, genState.Delegations)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	pools := k.AllPools(ctx)
	logs := k.AllDelegations(ctx)

	return types.NewGenesis(params, pools, logs)
}
