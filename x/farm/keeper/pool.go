package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gauss/gauss/v6/x/farm/types"
)

// CreatePool creates an new farm pool
func (k Keeper) CreatePool(
	ctx sdk.Context,
	name string,
	description string,
	minStaking sdk.Coin,
	startHeight int64,
	rewardsPerBlock sdk.Coins,
	totalRewards sdk.Coins,
	creator sdk.AccAddress,
) error {
	//Escrow total reward
	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, types.ModuleName, totalRewards); err != nil {
		return err
	}

	//send CreatePoolFee to feeCollectorName
	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, k.feeCollectorName, sdk.NewCoins(k.CreatePoolFee(ctx))); err != nil {
		return err
	}

	pool := types.FarmPool{
		Name:            name,
		Creator:         creator.String(),
		Description:     description,
		StartHeight:     startHeight,
		MinStaking:      minStaking,
		Tokens:          sdk.ZeroInt(),
		FarmRewardRules: []types.FarmRewardRule{},
	}

	var farmersRewards []types.FarmersRewardEntity
	//save farm rule
	for _, total := range totalRewards {
		rewardRule := types.FarmRewardRule{
			TotalRewards:    total,
			RewardsPerBlock: rewardsPerBlock.AmountOf(total.Denom),
		}
		pool.FarmRewardRules = append(pool.FarmRewardRules, rewardRule)

		farmersRewardEntity := types.FarmersRewardEntity{
			RemainingRewards: total,
			RewardsPerShare:  sdk.NewDec(0),
		}
		farmersRewards = append(farmersRewards, farmersRewardEntity)
	}

	endHeight, err := pool.ExpiredHeight()
	if err != nil {
		return err
	}
	//save farm pool
	pool.EndHeight = endHeight

	k.SetPool(ctx, pool)
	// put to expired farm pool queue
	k.EnqueueActivePool(ctx, name, pool.EndHeight)
	// save farms rewards
	var tfarmersRewards = types.FarmersRewards{
		PreviousHeight: pool.StartHeight,
		FarmersRewards: farmersRewards,
		PoolName:       pool.Name,
	}
	k.SetFarmersRewards(ctx, pool.Name, tfarmersRewards)
	return nil
}

// Destroy destroy an exist farm pool
func (k Keeper) DestroyPool(ctx sdk.Context, poolName string, creator sdk.AccAddress) (sdk.Coins, error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, poolName)
	}

	if creator.String() != pool.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "creator [%s] is not the creator of the pool", creator.String())
	}

	if k.Expired(ctx, pool) {
		return nil, sdkerrors.Wrapf(types.ErrPoolExpired,
			"pool [%s] has expired at height [%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}
	return k.Refund(ctx, pool)
}

// updatePool is responsible for updating the status information of the farm pool, including the total accumulated bonus from the last time the bonus was distributed to the present, the current remaining bonus in the farm pool, and the ratio of the current liquidity token to the bonus.

// Note that when multiple transactions at the same block height trigger the farm pool update at the same time, only the first transaction will trigger the `RewardPerShare` update operation

// updatePool returns the updated farm pool and the reward collected in this period
// 更新交易池token数量，计算当前区块per_share
func (k Keeper) updatePool(
	ctx sdk.Context,
	pool types.FarmPool,
	amount sdk.Int,
	isDestroy bool,
) (types.FarmPool, sdk.Coins, error) {
	height := ctx.BlockHeight()

	// 销毁未开始的farm pool
	if isDestroy && (pool.StartHeight > height) {
		pool.StartHeight = height
		pool.EndHeight = height
		k.SetPool(ctx, pool)
		return pool, nil, nil
	}

	farmersRewards := k.GetFarmersRewards(ctx, pool.Name)

	if height < farmersRewards.PreviousHeight {
		return pool, nil, sdkerrors.Wrapf(
			types.ErrExpiredHeight,
			"invalid height: [%d], last distribution height: [%d]",
			height, farmersRewards.PreviousHeight,
		)
	}
	rules := pool.FarmRewardRules
	//rules := k.GetRewardRules(ctx, pool.Name)
	//if len(rules) == 0 {
	//	return pool, nil, sdkerrors.Wrapf(types.ErrPoolNotFound, pool.Name)
	//}

	var rewardTotal sdk.Coins
	//when there are multiple farm operations in the same block, the value needs to be updated once
	if height > farmersRewards.PreviousHeight &&
		pool.Tokens.GT(sdk.ZeroInt()) {
		blockInterval := height - farmersRewards.PreviousHeight
		for i := range rules {
			rewardCollected := rules[i].RewardsPerBlock.MulRaw(blockInterval)
			coinCollected := sdk.NewCoin(rules[i].TotalRewards.Denom, rewardCollected)
			for j := range farmersRewards.FarmersRewards {
				if farmersRewards.FarmersRewards[j].RemainingRewards.Denom == rules[i].TotalRewards.Denom {
					if farmersRewards.FarmersRewards[j].RemainingRewards.Amount.LT(rewardCollected) {
						k.Logger(ctx).Error(
							"The remaining amount is not enough to pay the bonus",
							"poolName", pool.Name,
							"remainingReward", farmersRewards.FarmersRewards[j].RemainingRewards.Amount.String(),
							"rewardCollected", rewardCollected.String(),
						)
						return pool, nil, sdkerrors.Wrapf(
							sdkerrors.ErrInsufficientFunds,
							"the remaining reward of the pool [%s] is [%s], but got [%s]",
							pool.Name, farmersRewards.FarmersRewards[j].RemainingRewards, coinCollected,
						)
					}
					newRewardPerShare := sdk.NewDecFromInt(rewardCollected).QuoInt(pool.Tokens)
					farmersRewards.FarmersRewards[j].RewardsPerShare = farmersRewards.FarmersRewards[j].RewardsPerShare.Add(newRewardPerShare)
					farmersRewards.FarmersRewards[j].RemainingRewards = farmersRewards.FarmersRewards[j].RemainingRewards.Sub(coinCollected)
				}
			}
			rewardTotal = rewardTotal.Add(coinCollected)
		}
	}
	//escrow the collected rewards to the `RewardCollector` account
	if rewardTotal.IsAllPositive() {
		if err := k.bk.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.RewardCollector, rewardTotal); err != nil {
			return pool, rewardTotal, err
		}
	}
	pool.Tokens = pool.Tokens.Add(amount)

	farmersRewards.PreviousHeight = ctx.BlockHeight()
	if isDestroy {
		pool.EndHeight = ctx.BlockHeight()
		if pool.StartHeight > pool.EndHeight {
			pool.StartHeight = pool.EndHeight
		}
	}
	k.SetPool(ctx, pool)
	k.SetFarmersRewards(ctx, pool.Name, *farmersRewards)
	return pool, rewardTotal, nil
}
