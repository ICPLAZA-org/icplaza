package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/nft/types"
)

// 检测NFT是否已经存在
func (k Keeper) Exists(ctx sdk.Context, tokenId string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixTokenToCate + tokenId)
	return store.Has(key)
}

// 保存NFT
func (k Keeper) SaveNFT(ctx sdk.Context, nft *types.Nft) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixNFT + nft.CateId + nft.TokenId)
	bz, err := k.cdc.Marshal(nft)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// 保存分类
func (k Keeper) SaveCate(ctx sdk.Context, cateId string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixCate + cateId)
	var cate types.Cate
	cate.CateId = cateId
	bz, err := k.cdc.Marshal(&cate)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// 保存tokenID 用于查询时使用
func (k Keeper) SaveTokenID(ctx sdk.Context, cateId string, tokenId string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixTokenToCate + tokenId)
	var cate types.Cate
	cate.CateId = cateId
	bz, err := k.cdc.Marshal(&cate)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// 保存到token拥有人
func (k Keeper) SetOwnerNFT(ctx sdk.Context, cateId string, tokenId string, owner string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOwner + owner + cateId + tokenId)
	token := types.NfToken{ TokenId: tokenId }
	bz, err := k.cdc.Marshal(&token)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// 删除拥有人
func (k Keeper) DeleteOwnerNFT(ctx sdk.Context, cateId string, tokenId string, owner string) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixOwner + owner + cateId + tokenId)
	store.Delete(key)
}

// 设置Owner-Name
func (k Keeper) SetNameTokenId(ctx sdk.Context, name string, tokenId string, owner string) error {
	if name == "" {
		return nil
	}
	store := ctx.KVStore(k.storeKey)
	key := types.GetKeyOfTokenName(owner, name)
	token := types.NfToken{ TokenId: tokenId }
	bz, err := k.cdc.Marshal(&token)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// 删除Owner-Name
func (k Keeper) DeleteNameTokenId(ctx sdk.Context, name string, owner string) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKeyOfTokenName(owner, name)
	store.Delete(key)
}

// 创建人所有tokenId
func (k Keeper) SetCreatorNFT(ctx sdk.Context, cateId string, tokenId string, address string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixCreator + address + cateId + tokenId)

	token := types.NfToken{TokenId:tokenId}
	bz, err := k.cdc.Marshal(&token)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

// 查询NFT
func (k Keeper) GetNFT(ctx sdk.Context, cateId, tokenId string) (*types.Nft, error) {
	if cateId == "" {
		cateId = k.GetCateIdByTokenId(ctx, tokenId)
	}
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixNFT + cateId + tokenId)
	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidNFT, fmt.Sprintf("cateId #%s, tokenID #%s not exist", cateId, tokenId))
	}
	var nft types.Nft
	err := k.cdc.Unmarshal(bz, &nft)
	return &nft, err
}

// swap owner
func (k Keeper) SwapOwner(ctx sdk.Context, newOwner string, nft *types.Nft) error {
	// delete old
	k.DeleteOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
	k.DeleteNameTokenId(ctx, nft.Name, nft.Owner)
	// set new
	nft.Owner = newOwner
	err := k.SetOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
	if err != nil {
		return err
	}
	err = k.SetNameTokenId(ctx, nft.Name, nft.TokenId, nft.Owner)
	if err != nil {
		return err
	}
	err = k.SaveNFT(ctx, nft)
	if err != nil {
		return err
	}
	return nil
}

//
func (k Keeper) GetNFTByTokenId(ctx sdk.Context, tokenId string) (*types.Nft, error) {
	cateId := k.GetCateIdByTokenId(ctx, tokenId)
	return k.GetNFT(ctx, cateId, tokenId)
}

// 锁定NFT
func (k Keeper) LockNFT(ctx sdk.Context, cateId, tokenId string) (err error) {
	nft, err := k.GetNFT(ctx, cateId, tokenId)
	if err != nil {
		return err
	}
	if nft.TokenStatus == types.NFTFree {
		nft.TokenStatus = types.NFTLocked
		return k.SaveNFT(ctx, nft)
	}
	return sdkerrors.Wrap(types.ErrCannotLockOrFrozen, "")
}

// free nft
func (k Keeper) FreeNFT(ctx sdk.Context, cateId, tokenId string) error {
	nft, err := k.GetNFT(ctx, cateId, tokenId)
	if err != nil {
		return err
	}
	if nft.TokenStatus != types.NFTFrozen {
		nft.TokenStatus = types.NFTFree
		return k.SaveNFT(ctx, nft)
	}
	return sdkerrors.Wrap(types.ErrNFTFrozen, tokenId)
}

// 冻结
func (k Keeper) FrozenNFT(ctx sdk.Context, cateId, tokenId string) (err error) {
	nft, err := k.GetNFT(ctx, cateId, tokenId)
	if err != nil {
		return err
	}
	if nft.TokenStatus == types.NFTFree {
		nft.TokenStatus = types.NFTFrozen
		return k.SaveNFT(ctx, nft)
	}
	return sdkerrors.Wrap(types.ErrCannotLockOrFrozen, "")
}

// 出租：自由或者出租中都可以，冲突时间由其他模块制定
func (k Keeper) RentNFT(ctx sdk.Context, tokenId string) error {
	nft, err := k.GetNFT(ctx, "", tokenId)
	if err != nil {
		return err
	}
	if nft.TokenStatus == types.NFTFree {
		nft.TokenStatus = types.NFTRent
		return k.SaveNFT(ctx, nft)
	} else if nft.TokenStatus == types.NFTRent {
		return nil
	}
	return sdkerrors.Wrap(types.ErrNotFree, tokenId)
}

// 其他更改状态
func (k Keeper) ChangeNFTStatus(ctx sdk.Context, tokenId string, status uint32) error {
	nft, err := k.GetNFT(ctx, "", tokenId)
	if err != nil {
		return err
	}
	if nft.TokenStatus == types.NFTFree {
		nft.TokenStatus = status
		return k.SaveNFT(ctx, nft)
	}
	return sdkerrors.Wrap(types.ErrNotFree, tokenId)
}

// 解除冻结状态
func (k Keeper) UnFrozenNFT(ctx sdk.Context, cateId, tokenId string) error {
	nft, err := k.GetNFT(ctx, cateId, tokenId)
	if err != nil {
		return err
	}
	if nft.TokenStatus == types.NFTFrozen {
		nft.TokenStatus = types.NFTFree
		return k.SaveNFT(ctx, nft)
	}
	return sdkerrors.Wrap(types.ErrInvalidNFT, fmt.Sprintf("nft(%s) is not a frozen nft", tokenId))
}

// 通过tokenId获取cateId
func (k Keeper) GetCateIdByTokenId(ctx sdk.Context, tokenId string) string {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixTokenToCate + tokenId)
	bz := store.Get(key)
	if bz == nil {
		return ""
	}
	var cate types.Cate
	err := k.cdc.Unmarshal(bz, &cate)
	if err != nil {
		return ""
	}
	return cate.CateId
}

// 检查NFT状态
func (k Keeper) CheckNFT(nft types.Nft) error {
	// if nft.TokenStatus == types.NFTLocked {
	//	return sdkerrors.Wrap(types.ErrNFTLocked, nft.TokenId)
	// } else if nft.TokenStatus == types.NFTFrozen {
	//	return sdkerrors.Wrap(types.ErrNFTFrozen, nft.TokenId)
	// }
	if nft.TokenStatus != types.NFTFree {
		return sdkerrors.Wrap(types.ErrNotFree, nft.TokenId)
	}
	return nil
}

//// 历史最后成交价
func (k Keeper) GetLastPrice(ctx sdk.Context, tokenId string, denom string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixLastPrice + tokenId)
	var lastPrice types.LastPrice
	bz := store.Get(key)
	if bz == nil {
		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
	}
	err := k.cdc.Unmarshal(bz, &lastPrice)
	if err != nil {
		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
	}
	if lastPrice.Coin.Denom != denom {
		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
	}
	return *lastPrice.Coin
}
func (k Keeper) GetTokenLastPrice(ctx sdk.Context, tokenId string) *sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixLastPrice + tokenId)
	var lastPrice types.LastPrice
	bz := store.Get(key)
	if bz != nil {
		err := k.cdc.Unmarshal(bz, &lastPrice)
		if err != nil {
			return lastPrice.Coin
		}
	}
	return nil
}

//func (k Keeper) GetLastPrice(ctx sdk.Context, tokenId string, denom string) sdk.Coin {
//	store := ctx.KVStore(k.storeKey)
//	key := types.KeyPrefix(types.PrefixLastPrice + tokenId)
//	var lastPrice types.LastPrice
//	bz := store.Get(key)
//	if bz == nil {
//		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
//	}
//	err := k.cdc.Unmarshal(bz, &lastPrice)
//	if err != nil {
//		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
//	}
//	for _, coin := range lastPrice.Coins {
//		if coin.Denom == denom {
//			return sdk.Coin{Denom: coin.Denom, Amount: coin.Amount}
//		}
//	}
//	return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
//}

// 保存历史价格
// 仅保存上次价格
func (k Keeper) SaveLastPrice(ctx sdk.Context, tokenId string, price sdk.Coin) error {
	var lastPrice = types.LastPrice{
		TokenId: tokenId,
		Coin:    &price,
	}
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(types.PrefixLastPrice + tokenId)
	bz, err := k.cdc.Marshal(&lastPrice)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

//func (k Keeper) SaveLastPrice(ctx sdk.Context, tokenId string, price sdk.Coin) error {
//	var lastPrice types.LastPrice
//	store := ctx.KVStore(k.storeKey)
//	key := types.KeyPrefix(types.PrefixLastPrice + tokenId)
//	bz := store.Get(key)
//	if bz == nil {
//		lastPrice.TokenId = tokenId
//		lastPrice.Coins = append(lastPrice.Coins, &price)
//		bz, err := k.cdc.Marshal(&lastPrice)
//		if err != nil {
//			return err
//		}
//		store.Set(key, bz)
//	} else {
//		err := k.cdc.Unmarshal(bz, &lastPrice)
//		if err != nil {
//			return err
//		}
//		found := false
//		for index, coin := range lastPrice.Coins {
//			//if coin.Denom == price.Denom && coin.Amount.LT(price.Amount) {
//			if coin.Denom == price.Denom {
//				lastPrice.Coins[index] = &price
//				found = true
//			}
//		}
//		if !found {
//			lastPrice.Coins = append(lastPrice.Coins, &price)
//		}
//		bz, err := k.cdc.Marshal(&lastPrice)
//		if err != nil {
//			return err
//		}
//		store.Set(key, bz)
//	}
//	return nil
//}

// transfer from one to other
func (k Keeper) Transfer(ctx sdk.Context, from, to string, nft types.Nft, transferType uint64) error {
	_, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(to)
	if err != nil {
		return err
	}
	if nft.Owner != from {
		return sdkerrors.Wrap(types.ErrInvalidNFTOwner, fmt.Sprintf("(%s) is not the nft(%s) owner", from, nft.TokenId))
	}
	// if nft.TokenStatus == types.NFTFrozen {
	if nft.TokenStatus != types.NFTFree {
		return sdkerrors.Wrap(types.ErrNFTFrozen, fmt.Sprintf("nft(%s) it not allow transfer", nft.TokenId))
	}
	// transfer
	var isCNFT = false
	if len(nft.Components) > 0 {
		isCNFT = true
	}
	// tag transfered cnft
	if isCNFT {
		k.tagTransferedCNFT(ctx, nft.TokenId, transferType)
	}
	// do transfer
	// delete old owner
	k.DeleteOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
	// delete old name => tokenId relationship
	k.DeleteNameTokenId(ctx, nft.Name, nft.Owner)
	nft.Owner = to
	// save new owner
	err = k.SetOwnerNFT(ctx, nft.CateId, nft.TokenId, nft.Owner)
	if err != nil {
		return err
	}
	// save new name => tokenId relationship
	err = k.SetNameTokenId(ctx, nft.Name, nft.TokenId, nft.Owner)
	if err != nil {
		return err
	}
	// save nft
	err = k.SaveNFT(ctx, &nft)
	if err != nil {
		return err
	}
	return nil
}

// tag cnft was transfered
func (k Keeper) tagTransferedCNFT(ctx sdk.Context, tokenId string, transferType uint64) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKeyOfTranferedCNFT(tokenId)

	var transferedCNFT = types.TransferedCNFT{
		TokenId: tokenId,
		Type:    transferType,
	}
	bz, err := k.cdc.Marshal(&transferedCNFT)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

//
func (k Keeper) RemoveTransferedCNFT(ctx sdk.Context, tokenId string) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKeyOfTranferedCNFT(tokenId)
	store.Delete(key)
}

func (k Keeper) GetTransferedCNFT(ctx sdk.Context) []types.TransferedCNFT {
	var transferedCNFTs []types.TransferedCNFT
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PrefixTranferedCNFT))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var transferedCNFT types.TransferedCNFT
		err := k.cdc.Unmarshal(iterator.Value(), &transferedCNFT)
		if err == nil {
			transferedCNFTs = append(transferedCNFTs, transferedCNFT)
		}
	}
	return transferedCNFTs
}
