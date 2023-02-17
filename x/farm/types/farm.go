package types

import (
	math "math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (pool FarmPool) ExpiredHeight() (int64, error) {
	var targetInteval = int64(math.MaxInt64)
	for _, r := range pool.FarmRewardRules {
		inteval := r.TotalRewards.Amount.Quo(r.RewardsPerBlock).Int64()
		if targetInteval > inteval {
			targetInteval = inteval
		}
	}
	if int64(math.MaxInt64)-pool.StartHeight < targetInteval {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidHeight, "endheight overflow")
	}
	return pool.StartHeight + targetInteval, nil
}


func (frs *FarmersRewards) CaclRewards(farmInfo Farmer, deltaAmt sdk.Int, debets sdk.Coins) (rewards, rewardDebt sdk.Coins) {
	for _, r := range frs.FarmersRewards {
		if farmInfo.Tokens.GT(sdk.ZeroInt()) {
			pendingRewardTotal := r.RewardsPerShare.MulInt(farmInfo.Tokens).TruncateInt()
			pendingReward := pendingRewardTotal.Sub(debets.AmountOf(r.RemainingRewards.Denom))

			rewards = rewards.Add(sdk.NewCoin(r.RemainingRewards.Denom, pendingReward))
		}
		locked := farmInfo.Tokens.Add(deltaAmt)
		debt := sdk.NewCoin(r.RemainingRewards.Denom, r.RewardsPerShare.MulInt(locked).TruncateInt())
		rewardDebt = rewardDebt.Add(debt)
	}
	return rewards, rewardDebt
}

type FarmRewardRules []FarmRewardRule

func (rs FarmRewardRules) Contains(reward sdk.Coins) bool {
	if len(rs) < len(reward) {
		return false
	}
	var allRewards sdk.Coins
	for _, r := range rs {
		allRewards = allRewards.Add(r.TotalRewards)
	}
	return reward.DenomsSubsetOf(allRewards)
}

func (rs FarmRewardRules) UpdateWith(rewardsPerBlock sdk.Coins) FarmRewardRules {
	for i := range rs {
		rewardAmt := rewardsPerBlock.AmountOf(rs[i].TotalRewards.Denom)
		if rewardAmt.IsPositive() {
			rs[i].RewardsPerBlock = rewardAmt
		}
	}
	return rs
}

func (rs FarmRewardRules) RewardsPerBlock() (coins sdk.Coins) {
	for _, r := range rs {
		coins = coins.Add(sdk.NewCoin(r.TotalRewards.Denom, r.RewardsPerBlock))
	}
	return coins
}
