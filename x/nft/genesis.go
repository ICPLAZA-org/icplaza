package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/keeper"
	"github.com/gauss/gauss/v6/x/nft/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetLastPrices(ctx, genState.LastPrices)
	k.SetNfts(ctx, genState.Nfts)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	lastPrices := k.GetLastPrices(ctx)
	nfts := k.GetNfts(ctx)
	return types.NewGenesisState(params, lastPrices, nfts)
}
