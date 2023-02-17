package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestSetupMsgServer(t *testing.T) {
	a := sdk.Coin{Amount: sdk.NewInt(1000), Denom: "uunique"}
	b := sdk.Coin{Amount: sdk.NewInt(900), Denom: "uunique"}
	fmt.Println(a.Sub(b))
	d := a.Sub(b)
	c := int64(200)
	dd := d.Amount.Quo(sdk.NewInt(c))
	fmt.Println(dd)

}

//func (k Keeper) GetLastPrice(ctx sdk.Context, tokenId string, denom string) sdk.Coin {
//	store := ctx.KVStore(k.storeKey)
//	key := types.KeyPrefix(types.PrefixLastPrice + tokenId + denom)
//	var price sdk.Coin
//	bz := store.Get(key)
//	if bz == nil {
//		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
//	}
//	err := k.cdc.Unmarshal(bz, &price)
//	if err != nil {
//		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
//	}
//	err = k.cdc.Unmarshal(bz, &price)
//	if err != nil {
//		return sdk.Coin{Denom: denom, Amount: sdk.NewInt(0)}
//	}
//	return price
//}