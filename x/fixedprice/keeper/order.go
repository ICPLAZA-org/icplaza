package keeper

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/fixedprice/types"
)

func (k Keeper) RemoveOrder(ctx sdk.Context, poolAddress, tokenId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrder + poolAddress + tokenId)
	store.Delete(key)
	// delete tokenId to poolAddress
	k.deleteOrderTokenIdToPoolAddress(ctx, tokenId)
}

// 处理全部到期应结算订单，限制处理100条
func (k Keeper) CloseExpiredAndAdjustPrice(ctx sdk.Context) error {
	counter := 100
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixOrder))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		if counter == 0 {
			break
		}
		var order types.Order
		err := k.cdc.Unmarshal(iterator.Value(), &order)
		if err != nil {
			return err
		}
		// expired
		if order.EndTime.Before(ctx.BlockTime()) {
			// unlock nft
			err := k.nftKeeper.FreeNFT(ctx, order.Nft.CateId, order.Nft.TokenId)
			if err != nil {
				return err
			}
			// remove order
			k.RemoveOrder(ctx, order.PoolAddress, order.Nft.TokenId)
			// emit event
			//err = ctx.EventManager().EmitTypedEvent(&types.OrderClose{TokenId: order.TokenId})
			//if err != nil {
			//	panic(err)
			//}
			//ctx.EventManager().EmitEvent(
			//	sdk.NewEvent(
			//		types.EventTypeOrderClose,
			//		sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%s", order.TokenId)),
			//		sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%s", "0")),
			//	),
			//)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeExpiredOrderClose,
					sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%s", order.TokenId)),
					sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%s", "0")),
				),
			)
			counter--
		} else {
			if order.NextTimeAdjustPrice.Before(ctx.BlockTime()) {
				if order.Price.Equal(order.EndPrice) {
					continue
				}
				// reset next time
				params := k.GetParams(ctx)
				order.NextTimeAdjustPrice = order.NextTimeAdjustPrice.Add(params.AdjustPricePeriod)
				// calc new price
				d := order.EndTime.Sub(order.StartTime).Seconds()
				c := int64(math.Ceil(d/params.AdjustPricePeriod.Seconds())) - 1
				if c > 0 {
					adjustPrice := sdk.NewCoin(order.EndPrice.Denom, sdk.NewInt(0))
					if order.EndPrice.IsGTE(*order.StartPrice) {
						differ := order.EndPrice.Sub(*order.StartPrice).Amount
						adjustAmount := differ.Quo(sdk.NewInt(c))
						adjustPrice = sdk.NewCoin(order.EndPrice.Denom, adjustAmount)
						if order.Price.Add(adjustPrice).IsLT(*order.EndPrice) {
							t := *order.Price
							t = t.Add(adjustPrice)
							order.Price = &t
						} else {
							order.Price = order.EndPrice
						}
					} else {
						differ := order.StartPrice.Sub(*order.EndPrice).Amount
						adjustAmount := differ.Quo(sdk.NewInt(c))
						adjustPrice = sdk.NewCoin(order.EndPrice.Denom, adjustAmount)
						if order.Price.IsLT(adjustPrice) {
							order.Price = order.EndPrice
						} else {
							t := *order.Price
							if t.Sub(adjustPrice).IsGTE(*order.EndPrice) {
								t := t.Sub(adjustPrice)
								order.Price = &t
							} else {
								order.Price = order.EndPrice
							}
						}
					}
					// save order
					k.saveOrder(ctx, order)
					// emit event
					ctx.EventManager().EmitTypedEvent(&types.OrderPriceAdjust{TokenId: order.TokenId, Price: order.Price})
					ctx.EventManager().EmitEvent(
						sdk.NewEvent(
							types.EventTypeAdjustPrice,
							sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%s", order.TokenId)),
							sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%s", order.Price.String())),
						),
					)
				}
				counter--
			}
		}
	}
	return nil
}

func (k Keeper) saveOrder(ctx sdk.Context, order types.Order) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrder + order.PoolAddress + order.Nft.TokenId)
	bz, err := k.cdc.Marshal(&order)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// save order
func (k msgServer) saveOrder(ctx sdk.Context, order types.Order) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrder + order.PoolAddress + order.Nft.TokenId)
	bz, err := k.cdc.Marshal(&order)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// get order
func (k msgServer) getOrder(ctx sdk.Context, poolAddress, tokenId string) (*types.Order, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrder + poolAddress + tokenId)
	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidOrder, "tokenId="+tokenId+",poolAddress="+poolAddress)
	}
	var order types.Order
	err := k.cdc.Unmarshal(bz, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// check order
func (k Keeper) checkOrder(ctx sdk.Context, order types.Order) error {
	if order.EndTime.Before(ctx.BlockTime()) {
		return types.ErrOrderExpired
	}
	if order.StartTime.Before(ctx.BlockTime()) {
		return types.ErrOrderNotStart
	}
	return nil
}
