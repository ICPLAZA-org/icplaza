package keeper

import (
	"fmt"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gauss/gauss/v6/x/nftexpool/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		paramSpace    paramstypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	paramSpace paramstypes.Subspace,
// this line is used by starport scaffolding # ibc/keeper/parameter

) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		paramSpace:    paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams gets the parameters for the liquidity module.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams sets the parameters for the liquidity module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) SetPools(ctx sdk.Context, pools []*types.Pool) {
	for _, pool := range pools {
		err := k.savePool(ctx, *pool)
		if err != nil {
			panic(err)
		}
	}
}

func (k Keeper) SetDelegations(ctx sdk.Context, logs []*types.Delegation) {
	store := ctx.KVStore(k.storeKey)
	for _, log := range logs {
		key := types.KeyPrefix(types.PrefixDelegate + log.PoolAddress + log.Delegator)
		bz, err := k.cdc.Marshal(log)
		if err != nil {
			panic(err)
		}
		store.Set(key, bz)
	}
}
