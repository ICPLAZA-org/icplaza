package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/auction/types"
)

func (k Keeper) RemoveOrder(ctx sdk.Context, poolAddress, tokenId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOrder + poolAddress + tokenId)
	store.Delete(key)
	// delete tokenId to poolAddress
	k.deleteOrderTokenIdToPoolAddress(ctx, tokenId)
}

// 处理全部到期应结算订单，限制处理100条
func (k Keeper) CloseExpired(ctx sdk.Context) error {
	count := 100
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixOrder))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		if count == 0 {
			break
		}
		var order types.Order
		err := k.cdc.Unmarshal(iterator.Value(), &order)
		if err != nil {
			return err
		}
		// expired
		if order.EndTime.Before(ctx.BlockTime()) || (!order.MinEndTime.Equal(time.Time{}) && order.MinEndTime.Before(ctx.BlockTime())) {
			if order.Bidder == "" {
				// just remove
				// 1 unlock nft
				err := k.nftKeeper.FreeNFT(ctx, order.Nft.CateId, order.Nft.TokenId)
				if err != nil {
					return err
				}
				// 2 remove order
				k.RemoveOrder(ctx, order.PoolAddress, order.Nft.TokenId)
				// emit event
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						types.EventTypeExpiredOrderClose,
						sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%s", order.TokenId)),
						sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%s", "0")),
					),
				)
			} else {
				err = k.dealOrder(ctx, order)
				if err != nil {
					return err
				}
			}
			count--
		}
	}
	return nil
}

// deal order
func (k Keeper) dealOrder(ctx sdk.Context, order types.Order) error {
	pool, err := k.poolKeeper.GetPool(ctx, order.PoolAddress)
	if err != nil {
		return err
	}
	denom := order.BidPrice.Denom
	// 最终结算价格
	finalPrice := order.BidPrice

	// 历史成交价格
	lastPrice := k.nftKeeper.GetLastPrice(ctx, order.Nft.TokenId, denom)
	differPrice := sdk.Coin{Amount: sdk.NewInt(0), Denom: denom}
	if lastPrice.IsLT(finalPrice) {
		differPrice = finalPrice.Sub(lastPrice)
	}
	// 接收手续费地址
	commissionAddr, _ := sdk.AccAddressFromBech32(pool.CommissionAddress)
	// 手续费
	commissionRate := pool.CommissionRate
	commissionAmount := commissionRate.MulInt(finalPrice.Amount).RoundInt()
	commissionCoin := sdk.NewCoin(denom, commissionAmount)

	// 增值税
	valueAddedTax := order.Nft.ValueAddedTax

	valueAddedTaxAmount := sdk.NewInt(0)
	valueAddedTaxCoin := sdk.NewCoin(denom, valueAddedTaxAmount)

	if differPrice.IsPositive() {
		valueAddedTaxAmount = valueAddedTax.MulInt(differPrice.Amount).RoundInt()
		valueAddedTaxCoin = sdk.NewCoin(denom, valueAddedTaxAmount)
	}

	// 接收增值税地址
	valueAddedTaxAddr, _ := sdk.AccAddressFromBech32(pool.ValueAddedTaxAddress)

	// 版税比例
	copyrightTax := order.Nft.CopyrightTax
	copyrightAmount := copyrightTax.MulInt(finalPrice.Amount).RoundInt()
	copyrightCoin := sdk.NewCoin(denom, copyrightAmount)
	// 接收版税地址
	copyrightTaxAddr, _ := sdk.AccAddressFromBech32(order.Nft.Creator)

	// 卖家
	var sellerCoin sdk.Coin
	sellerAmount := order.BidPrice.Amount.Sub(commissionAmount).Sub(valueAddedTaxAmount).Sub(copyrightAmount)
	if sellerAmount.IsPositive() {
		sellerCoin = sdk.NewCoin(denom, sellerAmount)
	} else {
		sellerCoin = sdk.NewCoin(denom, sdk.NewInt(0))
	}
	sellerAddr, _ := sdk.AccAddressFromBech32(order.Nft.Owner)

	// 发送token
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, commissionAddr, sdk.Coins{commissionCoin})
	if valueAddedTaxCoin.IsPositive() {
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, valueAddedTaxAddr, sdk.Coins{valueAddedTaxCoin})
		// 增值税
		ctx.EventManager().EmitEvent(
			sdk.NewEvent("vat_fee_transfer",
				sdk.NewAttribute("from", types.ModuleName),
				sdk.NewAttribute("to", valueAddedTaxAddr.String()),
				sdk.NewAttribute("amount", valueAddedTaxCoin.String()),
			),
		)
	}
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, copyrightTaxAddr, sdk.Coins{copyrightCoin})
	// 版权税
	if copyrightCoin.IsPositive() {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent("copyright_fee_transfer",
				sdk.NewAttribute("from", types.ModuleName),
				sdk.NewAttribute("to", copyrightTaxAddr.String()),
				sdk.NewAttribute("amount", copyrightCoin.String()),
			),
		)
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sellerAddr, sdk.Coins{sellerCoin})
	// 事件
	if sellerCoin.IsPositive() {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent("token_transfer",
				sdk.NewAttribute("from", types.ModuleName),
				sdk.NewAttribute("to", sellerAddr.String()),
				sdk.NewAttribute("amount", sellerCoin.String()),
			),
		)
	}
	// set nft info
	// delete older info
	//k.nftKeeper.DeleteOwnerNFT(ctx, order.Nft.CateId, order.Nft.TokenId, order.Nft.Owner)
	// save new nft
	//order.Nft.Owner = order.Bidder
	//err = k.nftKeeper.SaveNFT(ctx, order.Nft)
	//if err != nil {
	//	return err
	//}
	//// set new nft info
	//err = k.nftKeeper.SetOwnerNFT(ctx, order.Nft.CateId, order.Nft.TokenId, order.Nft.Owner)
	//if err != nil {
	//	return err
	//}
	// -------------------------- 使用封装后的 start
	_ = k.nftKeeper.Transfer(ctx, order.Nft.Owner, order.Bidder, *order.Nft, 1)
	// nft转移事件
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("nft_transfer",
			sdk.NewAttribute("from", order.Nft.Owner),
			sdk.NewAttribute("to", order.Bidder),
			sdk.NewAttribute(types.AttributeKeyTokenID, order.Nft.TokenId),
		),
	)
	// -------------------------- 使用封装后的 end
	// 删除订单
	k.RemoveOrder(ctx, order.PoolAddress, order.Nft.TokenId)
	// 保存最后成交价格
	err = k.nftKeeper.SaveLastPrice(ctx, order.Nft.TokenId, finalPrice)
	if err != nil {
		return err
	}
	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeOrderClose,
			sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%s", order.TokenId)),
			sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%s", finalPrice.String())),
		),
	)
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
	if order.StartTime.After(ctx.BlockTime()) {
		return types.ErrOrderNotStart
	}
	if !order.MinEndTime.Equal(time.Time{}) && order.MinEndTime.Before(ctx.BlockTime()) {
		return types.ErrOrderExpired
	}
	return nil
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
