package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/gauss/gauss/v6/x/farm/types"
)

// Keeper of the farm store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         storetypes.StoreKey
	paramSpace       paramstypes.Subspace
	validateLPToken  types.ValidateLPToken
	bk               types.BankKeeper
	ak               types.AccountKeeper
	feeCollectorName string // name of the fee collector
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bk types.BankKeeper,
	ak types.AccountKeeper,
	paramSpace paramstypes.Subspace,
	feeCollectorName string,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ParamKeyTable())
	}

	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// ensure farm module accounts are set
	if addr := ak.GetModuleAddress(types.RewardCollector); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.RewardCollector))
	}

	return Keeper{
		storeKey:         storeKey,
		cdc:              cdc,
		bk:               bk,
		ak:               ak,
		validateLPToken:  validateLPToken,
		paramSpace:       paramSpace,
		feeCollectorName: feeCollectorName,
	}
}

// CreatePool creates an new farm pool
func (k Keeper) SetPool(ctx sdk.Context, pool types.FarmPool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pool)
	store.Set(types.KeyFarmPool(pool.Name), bz)
}

// GetPool return the specified farm pool
func (k Keeper) GetPool(ctx sdk.Context, poolName string) (types.FarmPool, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyFarmPool(poolName))
	if len(bz) == 0 {
		return types.FarmPool{}, false
	}

	var pool types.FarmPool
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}


func (k Keeper) SetRewardRules(ctx sdk.Context, poolName string, rules types.FarmRewardRules) {
	for _, r := range rules {
		k.SetRewardRule(ctx, poolName, r)
	}
}

func (k Keeper) SetRewardRule(ctx sdk.Context, poolName string, rule types.FarmRewardRule) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rule)

	store.Set(types.KeyRewardRule(poolName, rule.TotalRewards.Denom), bz)
}

func (k Keeper) GetRewardRules(ctx sdk.Context, poolName string) (rules types.FarmRewardRules) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixRewardRule(poolName))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.FarmRewardRule
		k.cdc.MustUnmarshal(iterator.Value(), &r)
		rules = append(rules, r)
	}
	return
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "irismod/farm")
}

func (k Keeper) IteratorRewardRules(ctx sdk.Context, poolName string, fun func(r types.FarmRewardRule)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixRewardRule(poolName))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.FarmRewardRule
		k.cdc.MustUnmarshal(iterator.Value(), &r)
		fun(r)
	}
}

func validateLPToken(ctx sdk.Context, denom string) error {
	return nil
}
