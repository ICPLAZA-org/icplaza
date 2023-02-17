package keeper

// DONTCOVER

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gauss/gauss/v6/x/farm/types"
)

// RegisterInvariants registers all invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "reward", RewardInvariant(k))
}

// AllInvariants runs all invariants of the farm module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return RewardInvariant(k)(ctx)
	}
}

// RewardInvariant checks whether the amount of module account is consistent with the recorded in the farm
func RewardInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		expectedBalances := sdk.Coins{}
		balances := k.bk.GetAllBalances(ctx, k.ak.GetModuleAddress(types.ModuleName))

		k.IteratorAllPools(ctx, func(pool types.FarmPool) {
			expectedBalances = expectedBalances.Add(sdk.NewCoin(pool.MinStaking.Denom, pool.Tokens))
		
			farmersRewards := k.GetFarmersRewards(ctx, pool.Name)
			for _, rr := range farmersRewards.FarmersRewards {
				expectedBalances = expectedBalances.Add(rr.RemainingRewards)
			}

			// k.IteratorRewardRules(ctx, pool.Name, func(r types.FarmRewardRule) {
			//	farmersRewards := k.GetFarmersRewards(ctx, pool.Name)
			//	tmpCoin := sdk.NewCoin(r.TotalRewards.Denom, sdk.ZeroInt())
			//	for _, rr := range farmersRewards.FarmersRewards {
			//		if rr.RemainingRewards.Denom == r.TotalRewards.Denom {
			//			tmpCoin = rr.RemainingRewards
			//		}
			//	}
			//	expectedBalance = expectedBalance.Add(tmpCoin)
			// })
		})

		broken := !expectedBalances.IsEqual(balances)
		return sdk.FormatInvariant(
			types.ModuleName,
			"module account balance",
			fmt.Sprintf(
				"\tsum of accounts coins: %v\n"+
					"\tbalance:          %v\n",
				expectedBalances, balances,
			),
		), broken
	}
}
