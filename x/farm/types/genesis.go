package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	defaultGenesisState = GenesisState{
		Params: Params{
			CreatingPoolFee:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(5000)),
			MaxRewardCategories: 2,
		},
	}
)

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(params Params, pools []FarmPool, farmers []Farmer, famersRewards []FarmersRewards, farmerRewards []FarmerRewards) *GenesisState {
	return &GenesisState{
		params, pools, farmers, famersRewards, farmerRewards,
	}
}

// DefaultGenesisState gets the default genesis state for testing
func DefaultGenesisState() *GenesisState {
	return &defaultGenesisState
}

func SetDefaultGenesisState(state GenesisState) {
	defaultGenesisState = state
}

// ValidateGenesis validates the provided farm genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	for _, pool := range data.FarmPools {
		if err := ValidatePoolName(pool.Name); err != nil {
			return err
		}

		if err := ValidateDescription(pool.Description); err != nil {
			return err
		}

		if err := ValidateAddress(pool.Creator); err != nil {
			return err
		}

		if err := ValidateCoins("TotalLptLocked", pool.MinStaking); err != nil {
			return err
		}

		for _, r := range pool.FarmRewardRules {
			if err := ValidateLpTokenDenom(r.TotalRewards.Denom); err != nil {
				return err
			}

			if !r.TotalRewards.IsPositive() {
				return fmt.Errorf("totalReward must be positive, but got %s", r.TotalRewards.String())
			}

			if !r.RewardsPerBlock.IsPositive() {
				return fmt.Errorf("rewardPerBlock must be positive, but got %s", r.RewardsPerBlock.String())
			}
		}
	}

	for _, farmersRewards := range data.FarmersRewards {
		for _, r := range farmersRewards.FarmersRewards {
			if r.RemainingRewards.IsNegative() {
				return fmt.Errorf("temainingReward must be greater than zero, but got %s", r.RemainingRewards.String())
			}
			if r.RewardsPerShare.IsNegative() {
				return fmt.Errorf("rewardPerShare must be positive, but got %s", r.RewardsPerShare.String())
			}
		}
	}

	for _, tt := range data.FarmerRewards {
		for _, r := range tt.Debts {
			if r.IsNegative() {
				return fmt.Errorf("Debts must be positive, but got %s", r.String())
			}
		}
	}

	for _, info := range data.Farmers {
		if err := ValidatePoolName(info.PoolName); err != nil {
			return err
		}

		if err := ValidateAddress(info.FarmerAddress); err != nil {
			return err
		}

		if info.Tokens.IsNegative() {
			return fmt.Errorf("locked must be positive, but got %s", info.Tokens.String())
		}
	}

	return ValidateCoins("CreatePoolFee", data.Params.CreatingPoolFee)
}
