package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/evmos/evmos/v11/x/fixedprice/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramSpace paramstypes.Subspace
		nftKeeper  types.NftKeeper
		poolKeeper types.PoolKeeper
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
	nft types.NftKeeper,
	pool types.PoolKeeper,
	bank types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: paramSpace,
		nftKeeper:  nft,
		poolKeeper: pool,
		bankKeeper: bank,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

func (k Keeper) GetAllOrders(ctx sdk.Context) []*types.Order {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixOrder))
	var orders []*types.Order
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var order types.Order
		err := k.cdc.Unmarshal(iterator.Value(), &order)
		if err != nil {
			panic(err)
		}
		orders = append(orders, &order)
	}
	return orders
}

// delete tokenId to poolAddress
func (k Keeper) deleteOrderTokenIdToPoolAddress(ctx sdk.Context, tokenId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrderTokenIdToPool + tokenId)
	store.Delete(key)
}
