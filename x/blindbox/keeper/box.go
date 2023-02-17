package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/blindbox/types"
	"math/rand"
	"time"
)

func (k Keeper) CreateBoxes(ctx sdk.Context, boxes []*types.Box) error {
	store := ctx.KVStore(k.storeKey)
	for i := 0; i < len(boxes); i++ {
		box := boxes[i]
		key := types.KeyOfBox(box.GroupId, box.BoxId)
		bz, err := k.cdc.Marshal(box)
		if err != nil {
			return err
		}
		store.Set(key, bz)
	}
	return nil
}

func (k Keeper) RemoveBox(ctx sdk.Context, box types.Box) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfBox(box.GroupId, box.BoxId)
	store.Delete(key)
}

func (k Keeper) RemoveBoxes(ctx sdk.Context, boxes []*types.Box) {
	store := ctx.KVStore(k.storeKey)
	for _, box := range boxes {
		key := types.KeyOfBox(box.GroupId, box.BoxId)
		store.Delete(key)
	}
}

func (k Keeper) OpenBox(ctx sdk.Context, groupId string, boxIds []string, creator string) error {
	store := ctx.KVStore(k.storeKey)
	var nfts []string
	for i := 0; i < len(boxIds); i++ {
		boxId := boxIds[i]
		key := types.KeyOfBox(groupId, boxId)
		bz := store.Get(key)
		if bz == nil {
			return sdkerrors.Wrapf(types.InvalidParam, "box(groupId:%s,boxId: %s) not exists", groupId, boxId)
		}
		var box types.Box
		err := k.cdc.Unmarshal(bz, &box)
		if err != nil {
			return err
		}
		if box.Opened {
			return sdkerrors.Wrapf(types.InvalidParam, "box(groupId:%s,boxId: %s) already opened", groupId, boxId)
		}
		// get group info
		group, err := k.GetGroupInfo(ctx, groupId)
		if err != nil {
			return err
		}

		if time.Now().Before(group.StartTime) {
			return sdkerrors.Wrapf(types.ErrNotStart, "box must open after %s", group.StartTime.Format("2006-01-02T15:04:05Z"))
		}
		if group.EndTime.Before(time.Now()) {
			return sdkerrors.Wrapf(types.ErrExpired, "box end at %s", group.EndTime.Format("2006-01-02T15:04:05Z"))
		}
		if group.OpenSize == group.BoxSize {
			return types.ErrAllOpened
		}
		// check balance
		fromAddr, _ := sdk.AccAddressFromBech32(creator)
		balance := k.bankKeeper.GetBalance(ctx, fromAddr, group.BoxPrice.Denom)
		if balance.IsLT(*group.BoxPrice) {
			return sdkerrors.Wrapf(types.ErrInsufficientFunds, "balance is smaller than %s", group.BoxPrice)
		}
		// open
		// create box
		fixedNFTs := group.FixedNfts
		leftFixedNFTs := group.LeftFixedNfts
		perSize := len(fixedNFTs) / int(group.BoxSize)
		if perSize > 0 {
			box.Nfts = append(box.Nfts, leftFixedNFTs[0:perSize]...)
			leftFixedNFTs = leftFixedNFTs[perSize:]
		}
		leftRandomNFTs := group.LeftRandomNfts
		randomAlloc := 0
		if len(leftRandomNFTs) > 0 {
			randomAlloc = random(int(group.BoxSize-group.OpenSize), int(group.RandomMin), int(group.RandomMax), len(leftRandomNFTs))
			box.Nfts = append(box.Nfts, leftRandomNFTs[0:randomAlloc]...)
			leftRandomNFTs = leftRandomNFTs[randomAlloc:]
		}
		nfts = append(nfts, box.Nfts...)
		// update box
		box.Opened = true
		err = k.saveBox(ctx, box)
		if err != nil {
			return err
		}
		// update group
		group.LeftFixedNfts = leftFixedNFTs
		group.LeftRandomNfts = leftRandomNFTs
		group.OpenSize += 1
		err = k.SaveGroup(ctx, *group)
		if err != nil {
			return err
		}
		if group.BoxPrice.IsPositive() {
			// send fee to feeAddr
			leftCoin := *group.BoxPrice
			found, pool := k.GetPoolById(ctx, group.PoolId)
			if !found {
				return sdkerrors.Wrapf(types.ErrNotFound, "pool(%s) not found", group.PoolId)
			}
			if pool.FeeRate.IsPositive() {
				feeAmt := pool.FeeRate.MulInt(group.BoxPrice.Amount).TruncateInt()
				if feeAmt.GT(sdk.ZeroInt()) {
					feeCoin := sdk.NewCoin(group.BoxPrice.Denom, feeAmt)
					feeAddr, _ := sdk.AccAddressFromBech32(pool.FeeAddress)
					k.bankKeeper.SendCoins(ctx, fromAddr, feeAddr, sdk.Coins{feeCoin})
					leftCoin = leftCoin.Sub(feeCoin)
				}
			}
			// send coin to seller
			toAddr, _ := sdk.AccAddressFromBech32(group.Creator)
			err = k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, sdk.Coins{leftCoin})
			if err != nil {
				return err
			}
		}
	}
	// send nft to buyer
	for _, token := range nfts {
		nft, err := k.nftKeeper.GetNFT(ctx, "", token)
		if err != nil {
			return err
		}
		// set free
		nft.TokenStatus = 0
		// then send
		err = k.nftKeeper.Transfer(ctx, nft.Owner, creator, *nft, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) saveBox(ctx sdk.Context, box types.Box) error {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfBox(box.GroupId, box.BoxId)
	bz, err := k.cdc.Marshal(&box)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

func (k Keeper) GetBoxInfo(ctx sdk.Context, groupId string, boxId string) (*types.Box, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyOfBox(groupId, boxId)
	bz := store.Get(key)
	var box types.Box
	err := k.cdc.Unmarshal(bz, &box)
	if err != nil {
		return nil, err
	}
	return &box, nil
}

func (k Keeper) GetBoxesByGroupId(ctx sdk.Context, groupId string) []*types.Box {
	store := ctx.KVStore(k.storeKey)

	str := fmt.Sprintf(types.PrefixKeyOfBox, groupId, "")

	var list []*types.Box

	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(str))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var item types.Box
		k.cdc.MustUnmarshal(iterator.Value(), &item)
		list = append(list, &item)
	}
	return list
}


/*
计算随机中奖数量
1.计算中奖概率
2.计算中多少个
@param boxSize 盒子数量
@param minPerBox 最少
@param maxPerBox 最多
@param total 分配数量
*/
func random(boxSize int, minPerBox int, maxPerBox int, total int) int {
	// 需要一定概率才能抽中
	win := 0
	if total < boxSize {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(boxSize)
		// 抽中
		if r <= total {
			// 随机个数
			rand.Seed(time.Now().Add(time.Second * 10).UnixNano())
			r = rand.Intn(maxPerBox - minPerBox + 1)
			win = r + minPerBox
		}
	} else {
		// 肯定会中
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(maxPerBox - minPerBox + 1)
		win = r + minPerBox
	}
	// 修正结果，保证剩余能分完total
	for {
		if total-win > (boxSize-1)*maxPerBox && boxSize > 0 && win <= maxPerBox {
			win++
		} else {
			break
		}
	}
	// 修正结果，保证不能分多
	for {
		if win > total && win >= minPerBox {
			win--
		} else {
			break
		}
	}
	// 保证结果达到最小值
	if win < minPerBox {
		win = minPerBox
	}
	// 保证结果不超过最大值
	if win > maxPerBox {
		win = maxPerBox
	}
	return win
}
