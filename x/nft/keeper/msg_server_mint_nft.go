package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/nft/types"
)

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nft := types.Nft{
		TokenId:       msg.TokenId,
		CateId:        msg.CateId,
		Owner:         msg.Recipient,
		Creator:       msg.Recipient,
		CompanyUri:    msg.CompanyUri,
		TokenUri:      msg.TokenUri,
		ValueAddedTax: msg.ValueAddedTax,
		CopyrightTax:  msg.CopyrightTax,
		TokenStatus:   types.NFTFree,
		Name:          msg.Name,
		Components:    msg.Components,
	}
	// nft mint fee
	//sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	//params := k.GetParams(ctx)
	//var createFee sdk.Coin
	//var isLess bool
	//var balance sdk.Coin
	//for _, coin := range params.NftCreationFee {
	//	balance = k.bankKeeper.GetBalance(ctx, sender, coin.Denom)
	//	createFee = coin
	//	if balance.IsLT(coin) {
	//		isLess = true
	//	} else {
	//		isLess = false
	//	}
	//}
	//// is less
	//if isLess {
	//	return nil, sdkerrors.Wrapf(
	//		types.ErrInsufficientFunds, "%s is smaller than %s", balance, createFee)
	//}
	// deal nft mint fee

	// mint
	if k.Exists(ctx, msg.TokenId) {
		return nil, sdkerrors.Wrap(types.ErrNFTAlreadyExists, "tokenId already exists "+msg.TokenId)
	}
	// 保存
	err := k.save(ctx, nft)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintNftResponse{}, nil
}

// 保存
func (k msgServer) save(ctx sdk.Context, nft types.Nft) error {
	err := k.SaveNFT(ctx, &nft)
	if err != nil {
		return err
	}
	err = k.SaveCate(ctx, nft.CateId)
	if err != nil {
		return err
	}
	err = k.SaveTokenID(ctx, nft.CateId, nft.TokenId)
	if err != nil {
		return err
	}
	err = k.SetOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
	if err != nil {
		return err
	}
	err = k.SetCreatorNFT(ctx, nft.CateId, nft.TokenId, nft.Creator)
	if err != nil {
		return err
	}
	err = k.SetNameTokenId(ctx, nft.Name, nft.TokenId, nft.Owner)
	if err != nil {
		return err
	}
	return nil
}
