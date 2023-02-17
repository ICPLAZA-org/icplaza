package validator_dao

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/validator-dao/keeper"
	"github.com/gauss/gauss/v6/x/validator-dao/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, authorizerBizs := range genState.AuthorizerBizs {
		authorizerAddr, err := sdk.AccAddressFromBech32(authorizerBizs.Authorizer)
		if err != nil {
			panic(err)
		}
		k.SetAuthorizerBizs(ctx, authorizerAddr, authorizerBizs)
	}
	for _, granteeAuthBizs := range genState.GranteeAuthBizs {
		granteeAddr, err := sdk.AccAddressFromBech32(granteeAuthBizs.Grantee)
		if err != nil {
			panic(err)
		}
		k.SetGranteeAuthBizs(ctx, granteeAddr, granteeAuthBizs)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Params = k.GetParams(ctx)
	genesis.AuthorizerBizs = k.GetAllAuthorizerBizs(ctx)
	genesis.GranteeAuthBizs = k.GetAllGranteeAuthBizs(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
