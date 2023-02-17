package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gauss/gauss/v6/x/farm/types"
)

// Stake is responsible for the user to mortgage the lp token to the system and get back the reward accumulated before then
func (k Keeper) Stake(
	ctx sdk.Context,
	poolName string,
	lpToken sdk.Coin,
	sender sdk.AccAddress,
) (reward sdk.Coins, err error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return reward, sdkerrors.Wrapf(types.ErrPoolNotFound, poolName)
	}

	if pool.StartHeight > ctx.BlockHeight() {
		return reward, sdkerrors.Wrapf(
			types.ErrPoolNotStart,
			"farm pool [%s] will start at height [%d], current [%d]",
			poolName, pool.StartHeight, ctx.BlockHeight(),
		)
	}

	if k.Expired(ctx, pool) {
		return reward, sdkerrors.Wrapf(
			types.ErrPoolExpired,
			"pool [%s] has expired at height [%d], current [%d]",
			poolName, pool.EndHeight, ctx.BlockHeight(),
		)
	}

	if lpToken.Denom != pool.MinStaking.Denom {
		return reward, sdkerrors.Wrapf(
			types.ErrNotMatch,
			"pool [%s] only accept [%s] token, but got [%s]",
			poolName, pool.MinStaking.Denom, lpToken.Denom,
		)
	}
	// check amount
	if lpToken.IsLT(pool.MinStaking) {
		return reward, sdkerrors.Wrapf(
			types.ErrMinStaking,
			"staking amount [%s] should greater than [%s]",
			lpToken, pool.MinStaking,
		)
	}
	if err := k.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(lpToken)); err != nil {
		return reward, err
	}

	//update pool reward shards
	pool, _, err = k.updatePool(ctx, pool, lpToken.Amount, false)
	if err != nil {
		return nil, err
	}

	farmInfo, exist := k.GetFarmerInfo(ctx, poolName, sender.String())
	if !exist {
		farmInfo = types.Farmer{
			PoolName:      poolName,
			FarmerAddress: sender.String(),
			Tokens:        sdk.ZeroInt(),
		}
	}
	//
	farmerRewards := k.GetFarmerRewards(ctx, poolName, sender.String())
	//
	farmersRewards := k.GetFarmersRewards(ctx, poolName)

	rewards, rewardDebt := farmersRewards.CaclRewards(farmInfo, lpToken.Amount, farmerRewards.Debts)
	//reward users
	if rewards.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.RewardCollector, sender, rewards); err != nil {
			return reward, err
		}
	}
	farmInfo.Tokens = farmInfo.Tokens.Add(lpToken.Amount)

	k.SetFarmerInfo(ctx, farmInfo)
	// save farmer rewards
	farmerRewards.Debts = rewardDebt
	farmerRewards.FarmerAddress = sender.String()
	farmerRewards.PoolName = poolName
	k.SetFarmerRewards(ctx, poolName, sender.String(), *farmerRewards)
	return rewards, nil
}

// Unstake withdraw lp token from farm pool
func (k Keeper) Unbond(ctx sdk.Context, poolName string, lpToken sdk.Coin, sender sdk.AccAddress) (_ sdk.Coins, err error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, poolName)
	}

	//lpToken demon must be same as pool.TotalLptLocked.Denom
	if lpToken.Denom != pool.MinStaking.Denom {
		return nil, sdkerrors.Wrapf(
			types.ErrNotMatch,
			"pool [%s] only accept [%s] token, but got [%s]",
			poolName, pool.MinStaking.Denom, lpToken.Denom,
		)
	}

	//farmInfo must be exist
	farmInfo, exist := k.GetFarmerInfo(ctx, poolName, sender.String())
	if !exist {
		return nil, sdkerrors.Wrapf(
			types.ErrFarmerNotFound,
			"farmer [%s] not found in pool [%s]",
			sender.String(), poolName,
		)
	}

	//the lp token unstaked must be less than staked
	if farmInfo.Tokens.LT(lpToken.Amount) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"farmer locked lp token [%s], but unstake [%s]",
			farmInfo.Tokens.String(), lpToken.Amount.String(),
		)
	}

	//the lp token unstaked must be less than pool
	if pool.Tokens.LT(lpToken.Amount) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"farmer locked token [%s], but farm pool total: [%s]",
			farmInfo.Tokens.String(), pool.Tokens.String(),
		)
	}

	if k.Expired(ctx, pool) {
		//If the farm has ended, the reward rules cannot be updated
		pool.Tokens = pool.Tokens.Sub(lpToken.Amount)
		k.SetPool(ctx, pool)
	} else {
		//update pool reward shards
		pool, _, err = k.updatePool(ctx, pool, lpToken.Amount.Neg(), false)
		if err != nil {
			return nil, err
		}
	}

	//unstake lpToken to sender account
	if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(lpToken)); err != nil {
		return nil, err
	}
	//compute farmer rewards
	farmerRewards := k.GetFarmerRewards(ctx, poolName, sender.String())
	//
	farmersRewards := k.GetFarmersRewards(ctx, poolName)
	rewards, rewardDebt := farmersRewards.CaclRewards(farmInfo, lpToken.Amount.Neg(), farmerRewards.Debts)
	if rewards.IsAllPositive() {
		//distribute reward
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.RewardCollector, sender, rewards); err != nil {
			return nil, err
		}
	}
	farmInfo.Tokens = farmInfo.Tokens.Sub(lpToken.Amount)
	if farmInfo.Tokens.IsZero() {
		k.DeleteFarmerInfo(ctx, poolName, sender.String())
		return rewards, nil
	}
	k.SetFarmerInfo(ctx, farmInfo)
	// save farmer rewards
	farmerRewards.Debts = rewardDebt
	k.SetFarmerRewards(ctx, poolName, sender.String(), *farmerRewards)
	return rewards, nil
}

// Harvest creates an new farm pool
func (k Keeper) Withdraw(ctx sdk.Context, poolName string, sender sdk.AccAddress) (sdk.Coins, error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, poolName)
	}

	if k.Expired(ctx, pool) {
		return nil, sdkerrors.Wrapf(
			types.ErrPoolExpired,
			"pool [%s] has expired at height [%d], current [%d]",
			poolName, pool.EndHeight, ctx.BlockHeight(),
		)
	}

	farmInfo, exist := k.GetFarmerInfo(ctx, poolName, sender.String())
	if !exist {
		return nil, sdkerrors.Wrapf(
			types.ErrFarmerNotFound,
			"farmer [%s] not found in pool [%s]",
			sender.String(), poolName,
		)
	}

	amtAdded := sdk.ZeroInt()
	//update pool reward shards
	pool, _, err := k.updatePool(ctx, pool, amtAdded, false)
	if err != nil {
		return nil, err
	}
	//compute farmer rewards
	farmerRewards := k.GetFarmerRewards(ctx, poolName, sender.String())
	//
	farmersRewards := k.GetFarmersRewards(ctx, poolName)
	rewards, rewardDebt := farmersRewards.CaclRewards(farmInfo, sdk.ZeroInt(), farmerRewards.Debts)
	//reward users
	if rewards.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.RewardCollector, sender, rewards); err != nil {
			return nil, err
		}
	}

	// save farmer rewards
	farmerRewards.Debts = rewardDebt
	k.SetFarmerRewards(ctx, poolName, sender.String(), *farmerRewards)

	return rewards, nil
}

// Refund refund the remaining reward to pool creator
func (k Keeper) Refund(ctx sdk.Context, pool types.FarmPool) (sdk.Coins, error) {
	//remove from active Pool
	k.DequeueActivePool(ctx, pool.Name, pool.EndHeight)
	pool, _, err := k.updatePool(ctx, pool, sdk.ZeroInt(), true)
	if err != nil {
		return nil, err
	}

	creator, err := sdk.AccAddressFromBech32(pool.Creator)
	if err != nil {
		return nil, err
	}
	//
	var refundTotal sdk.Coins
	farmersRewards := k.GetFarmersRewards(ctx, pool.Name)
	for index, r := range farmersRewards.FarmersRewards {
		refundTotal = refundTotal.Add(r.RemainingRewards)
		// TODO 需要?
		farmersRewards.FarmersRewards[index].RemainingRewards = sdk.NewCoin(r.RemainingRewards.Denom, sdk.ZeroInt())
	}
	// TODO 需要?
	k.SetFarmersRewards(ctx, pool.Name, *farmersRewards)
	if refundTotal.IsAllPositive() {
		//refund the total remaining reward to creator
		if err := k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, refundTotal); err != nil {
			return nil, err
		}
	}
	return refundTotal, nil
}
