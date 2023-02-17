package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
)

// InitGenesis initializes the nft module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := k.ValidateGenesis(ctx, genState); err != nil {
		panic(err)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the liquidity module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	return &types.GenesisState{Params: params}
}

// ValidateGenesis validates the liquidity module's genesis state.
func (k Keeper) ValidateGenesis(ctx sdk.Context, genState types.GenesisState) error {
	if err := genState.Params.Validate(); err != nil {
		return err
	}

	cc, _ := ctx.CacheContext()
	k.SetParams(cc, genState.Params)

	return nil
}
