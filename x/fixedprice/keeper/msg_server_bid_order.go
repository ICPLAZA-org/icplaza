package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gauss/gauss/v6/x/fixedprice/types"
)

func (k msgServer) BidOrder(goCtx context.Context, msg *types.MsgBidOrder) (*types.MsgBidOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, err := k.getOrder(ctx, msg.PoolAddress, msg.TokenId)
	if err != nil {
		return nil, err
	}
	// check balance
	err = k.checkBalance(ctx, msg.Sender, *msg.Price)
	if err != nil {
		return nil, err
	}
	// check price
	if msg.Price.Denom != order.Price.Denom {
		return nil, types.ErrDenomMatch
	}
	if !msg.Price.Equal(order.Price) {
		return nil, sdkerrors.Wrap(types.ErrInvalidBidPrice, fmt.Sprintf("bid price(%s) should equal order price(%s)", msg.Price, order.Price))
	}
	//
	pool, err := k.poolKeeper.GetPool(ctx, order.PoolAddress)
	if err != nil {
		return nil, err
	}
	taxAddr, _ := sdk.AccAddressFromBech32(pool.ValueAddedTaxAddress)
	commissionAddr, _ := sdk.AccAddressFromBech32(pool.CommissionAddress)
	commissionRate := pool.CommissionRate
	creator, _ := sdk.AccAddressFromBech32(order.Nft.Creator)
	taxRate := order.Nft.ValueAddedTax
	copyrightTax := order.Nft.CopyrightTax

	//// change owner
	//k.nftKeeper.DeleteOwnerNFT(ctx, order.Nft.CateId, order.Nft.TokenId, order.Nft.Owner)
	//// seller
	seller, _ := sdk.AccAddressFromBech32(order.Nft.Owner)
	//// save new
	//order.Nft.Owner = msg.Sender
	//err = k.nftKeeper.SaveNFT(ctx, order.Nft)
	//if err != nil {
	//	return nil, err
	//}
	//k.nftKeeper.SetOwnerNFT(ctx, order.Nft.CateId, order.Nft.TokenId, order.Nft.Owner)

	// -------------------------- 使用封装后的 start
	err = k.nftKeeper.Transfer(ctx, order.Nft.Owner, msg.Sender, *order.Nft, 1)
	if err != nil {
		return nil, err
	}
	// -------------------------- 使用封装后的 end
	// 转账 start
	denom := order.Price.Denom
	// 手续费
	commissionAmount := commissionRate.MulInt(order.Price.Amount).RoundInt()
	commisionCoin := sdk.NewCoin(denom, commissionAmount)
	// 增值税
	lastPrice := k.nftKeeper.GetLastPrice(ctx, order.Nft.TokenId, msg.Price.Denom)
	taxAmount := sdk.NewInt(0)
	taxCoin := sdk.NewCoin(denom, sdk.NewInt(0))

	differPrice := sdk.Coin{Amount: sdk.NewInt(0), Denom: msg.Price.Denom}
	if lastPrice.IsLT(*msg.Price) {
		differPrice = msg.Price.Sub(lastPrice)
	}
	if differPrice.IsPositive() {
		taxAmount = taxRate.MulInt(differPrice.Amount).RoundInt()
		taxCoin = sdk.NewCoin(denom, taxAmount)
	}
	// 版税
	copyrightAmount := copyrightTax.MulInt(order.Price.Amount).RoundInt()
	copyrightCoin := sdk.NewCoin(denom, copyrightAmount)
	// 出售人

	sellerAmount := order.Price.Amount.Sub(commissionAmount).Sub(taxAmount).Sub(copyrightAmount)
	sellerCoin := sdk.NewCoin(denom, sellerAmount)

	var inputs []banktypes.Input
	var outputs []banktypes.Output
	accSender, _ := sdk.AccAddressFromBech32(msg.Sender)
	inputs = append(inputs, banktypes.NewInput(accSender, sdk.Coins{*msg.Price}))
	if commisionCoin.IsPositive() {
		outputs = append(outputs, banktypes.NewOutput(commissionAddr, sdk.Coins{commisionCoin}))
	}
	if copyrightCoin.IsPositive() {
		outputs = append(outputs, banktypes.NewOutput(creator, sdk.Coins{copyrightCoin}))
	}
	if taxCoin.IsPositive() {
		outputs = append(outputs, banktypes.NewOutput(taxAddr, sdk.Coins{taxCoin}))
	}
	outputs = append(outputs, banktypes.NewOutput(seller, sdk.Coins{sellerCoin}))

	err = k.bankKeeper.InputOutputCoins(ctx, inputs, outputs)
	if err != nil {
		return nil, err
	}
	// 转账 end
	// 删除订单
	k.RemoveOrder(ctx, order.PoolAddress, order.TokenId)
	// 保存最后成交价格
	lastPrice = sdk.NewCoin(denom, msg.Price.Amount)
	err = k.nftKeeper.SaveLastPrice(ctx, order.Nft.TokenId, lastPrice)
	if err != nil {
		return nil, err
	}

	return &types.MsgBidOrderResponse{}, nil
}

// check balance
func (k Keeper) checkBalance(ctx sdk.Context, senderAddr string, bidPrice sdk.Coin) error {
	sender, _ := sdk.AccAddressFromBech32(senderAddr)
	balance := k.bankKeeper.GetBalance(ctx, sender, bidPrice.Denom)
	if balance.IsLT(bidPrice) {
		return sdkerrors.Wrapf(types.ErrInsufficientFunds, "balance is smaller than %s", bidPrice)
	}
	return nil
}
