package farm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gauss/gauss/v6/x/farm/keeper"
	"github.com/gauss/gauss/v6/x/farm/types"
)

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}
	// save farm pools
	for _, pool := range data.FarmPools {
		k.SetPool(ctx, pool)
		if !k.Expired(ctx, pool) {
			k.EnqueueActivePool(ctx, pool.Name, pool.EndHeight)
		}
	}
	// save farmer
	for _, famer := range data.Farmers {
		k.SetFarmerInfo(ctx, famer)
	}
	// save famers rewards
	for _, fsr := range data.FarmersRewards {
		k.SetFarmersRewards(ctx, fsr.PoolName, fsr)
	}
	// save farmer debets
	for _, fr := range data.FarmerRewards {
		k.SetFarmerRewards(ctx, fr.PoolName, fr.FarmerAddress, fr)
	}
	// save params
	k.SetParams(ctx, data.Params)
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// get params
	params := k.GetParams(ctx)
	// get farm pools
	var pools []types.FarmPool
	var farmers []types.Farmer
	var farmersRewards []types.FarmersRewards
	var farmerRewards []types.FarmerRewards
	k.IteratorAllPools(ctx, func(pool types.FarmPool) {
		pools = append(pools, pool)
	})
	// get farmers
	k.IteratorAllFarmInfo(ctx, func(farmInfo types.Farmer) {
		farmers = append(farmers, farmInfo)
	})
	// get all farmers rewards
	k.IteratorAllFarmersRewards(ctx, func(r types.FarmersRewards) {
		farmersRewards = append(farmersRewards, r)
	})
	// get all farmer rewards
	k.IteratorAllFarmerRewards(ctx, func(r types.FarmerRewards) {
		farmerRewards = append(farmerRewards, r)
	})
	return &types.GenesisState{
		Params:         params,
		FarmPools:      pools,
		Farmers:        farmers,
		FarmersRewards: farmersRewards,
		FarmerRewards:  farmerRewards,
	}
}
