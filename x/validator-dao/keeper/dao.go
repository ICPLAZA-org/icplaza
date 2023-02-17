package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gauss/gauss/v6/x/validator-dao/types"
)

func (k Keeper) ConsumeAuthorization(ctx sdk.Context, granteeAddr sdk.AccAddress, bizName string) (bool) {
	rc := k.IsAuthorizer(ctx, granteeAddr)
	if rc {
		return rc
	}

	granteeAuthBizs := k.GetGranteeAuthBizs(ctx, granteeAddr)
	for i, authBiz := range granteeAuthBizs.Bizs {
		if bizName == authBiz.BizName {
			if rc = k.consumeAuthorization(ctx, &authBiz); !rc {
				continue
			}
			if authBiz.Amount.IsZero() {
				granteeAuthBizs.Bizs = append(granteeAuthBizs.Bizs[:i], granteeAuthBizs.Bizs[i+1:]...)
				if len(granteeAuthBizs.Bizs) == 0 {
					k.removeGranteeAuthBizs(ctx, granteeAddr)
				}
			} else {
				granteeAuthBizs.Bizs[i].Amount = authBiz.Amount
			}
			k.SetGranteeAuthBizs(ctx, granteeAddr, granteeAuthBizs)
			break;
		}	
	}

	return rc
}

func (k Keeper) consumeAuthorization(ctx sdk.Context, authBiz *types.AcqBiz) (bool) {
	rc := false

	fee := authBiz.Price

	count := authBiz.Amount.Amount.Quo(authBiz.Price.Amount)
	if count.LT(sdk.OneInt()) {
		return rc
	}
	if count.Equal(sdk.OneInt()) {
		fee = authBiz.Amount
	}
	authBiz.Amount = authBiz.Amount.Sub(fee)

	authorizerAddr, err := sdk.AccAddressFromBech32(authBiz.From)
	if err != nil {
		return rc
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, authorizerAddr, sdk.NewCoins(fee)); err != nil {
		return rc
	}

	rc = true
	return rc
}

func (k Keeper) ReqAuthorization(ctx sdk.Context, granteeAddr, authorizerAddr sdk.AccAddress, bizName string, fee sdk.Coin) error {
	// only once
	granteeAuthBizs := k.GetGranteeAuthBizs(ctx, granteeAddr)
	for _, authBiz := range granteeAuthBizs.Bizs {
		if (bizName == authBiz.BizName) && (authorizerAddr.String() == authBiz.From) {
			return sdkerrors.Wrapf(types.ErrAuthorizationFound, "biz(%s)'s authorization from %s does exist.", bizName, authBiz.From)
		}
	}
	
	// fee
	price := k.getAuthorizerBizPrice(ctx, authorizerAddr, bizName, fee.Denom)
	if price.IsZero() {
		return sdkerrors.Wrapf(types.ErrNoBizFound, "biz(%s) does not exist.", bizName)
	}
	if fee.IsLT(price) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "got: %s, expect: %s", fee.String(), price.String())
	}

	// send
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, granteeAddr, types.ModuleName, sdk.NewCoins(fee)); err != nil {
		return err
	}
	
	// authorize
	granteeAuthBizs.Bizs = append(granteeAuthBizs.Bizs, types.NewAcqBiz(authorizerAddr.String(), bizName, fee, price))
	k.SetGranteeAuthBizs(ctx, granteeAddr, granteeAuthBizs)

	return nil
}


func (k Keeper) WithdrawAuthorization(ctx sdk.Context, granteeAddr, authorizerAddr sdk.AccAddress, bizName string) (sdk.Coin, error) {
	return sdk.NewInt64Coin(sdk.DefaultBondDenom, 0), sdkerrors.Wrap(sdkerrors.ErrNotSupported, "Withdrawing authorization do not implement")
}

func (k Keeper) AddAuthBiz(ctx sdk.Context, authorizerAddr sdk.AccAddress, bizName string, fee sdk.Coin) error {
	if err := k.validateBiz(ctx, bizName, fee); err != nil {
		return err
	}

	authorizerBizs := k.GetAuthorizerBizs(ctx, authorizerAddr)
	for _, biz := range authorizerBizs.Bizs {
		if biz.Name == bizName {
			return sdkerrors.Wrapf(types.ErrBizFound, "biz(%s) already exist.", bizName)
		}
	}
	
	authorizerBizs.Bizs = append(authorizerBizs.Bizs, types.NewDaoBiz(bizName, fee))
        k.SetAuthorizerBizs(ctx, authorizerAddr, authorizerBizs)

	return nil
}

func (k Keeper) UpdateAuthBiz(ctx sdk.Context, authorizerAddr sdk.AccAddress, bizName string, fee sdk.Coin) error {
	if err := k.validateBiz(ctx, bizName, fee); err != nil {
		return err
	}

	var found bool = false
	authorizerBizs := k.GetAuthorizerBizs(ctx, authorizerAddr)
	for i, biz := range authorizerBizs.Bizs {
		if biz.Name == bizName {
			authorizerBizs.Bizs[i].Fee = fee
			found = true
			break
		}
	}
	if !found {
		return sdkerrors.Wrapf(types.ErrNoBizFound, "biz(%s) does not exist.", bizName)
	} else {
        	k.SetAuthorizerBizs(ctx, authorizerAddr, authorizerBizs)
	}

	return nil
}

func (k Keeper) RemoveAuthBiz(ctx sdk.Context, authorizerAddr sdk.AccAddress, bizName string) error {
	var found bool = false
	authorizerBizs := k.GetAuthorizerBizs(ctx, authorizerAddr)
	for i, biz := range authorizerBizs.Bizs {
		if biz.Name == bizName {
			authorizerBizs.Bizs = append(authorizerBizs.Bizs[:i], authorizerBizs.Bizs[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return sdkerrors.Wrapf(types.ErrNoBizFound, "biz(%s) does not found", bizName)
	} else {
		k.SetAuthorizerBizs(ctx, authorizerAddr, authorizerBizs)
	}

	return nil
}

func (k Keeper) SetAuthorizerBizs(ctx sdk.Context, authorizerAddr sdk.AccAddress, bizs types.AuthorizerBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bizs)
	store.Set(types.GetAuthorizerBizsKey(authorizerAddr), bz)
}

func (k Keeper) GetAuthorizerBizs(ctx sdk.Context, authorizerAddr sdk.AccAddress) (authorizerBizs types.AuthorizerBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetAuthorizerBizsKey(authorizerAddr))
	if bz == nil {
		return types.AuthorizerBizs{}
	}
	k.cdc.Unmarshal(bz, &authorizerBizs)
	return authorizerBizs
}

func (k Keeper) RemoveAuthorizerBizs(ctx sdk.Context, authorizerAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAuthorizerBizsKey(authorizerAddr))
}

func (k Keeper) SetGranteeAuthBizs(ctx sdk.Context, granteeAddr sdk.AccAddress, granteeAuthBizs types.GranteeBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&granteeAuthBizs)
	store.Set(types.GetGranteeAuthBizsKey(granteeAddr), bz)
}

func (k Keeper) GetGranteeAuthBizs(ctx sdk.Context, granteeAddr sdk.AccAddress) (granteeAuthBizs types.GranteeBizs) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetGranteeAuthBizsKey(granteeAddr))
	if bz == nil {
		return types.GranteeBizs{Grantee: granteeAddr.String()}
	}
	k.cdc.Unmarshal(bz, &granteeAuthBizs)
	return granteeAuthBizs
}

func (k Keeper) removeGranteeAuthBizs(ctx sdk.Context, granteeAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetGranteeAuthBizsKey(granteeAddr))
}

func (k Keeper) GetAllAuthorizerBizs(ctx sdk.Context) []types.AuthorizerBizs {
	var rc []types.AuthorizerBizs

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AuthorizerBizsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.AuthorizerBizs
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		rc = append(rc, data)
	}
	return rc
}

func (k Keeper) GetAllGranteeAuthBizs(ctx sdk.Context) []types.GranteeBizs {
	var rc []types.GranteeBizs

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GranteeAuthBizsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.GranteeBizs
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		rc = append(rc, data)
	}
	return rc
}

func (k Keeper) validateBiz(ctx sdk.Context, bizName string, fee sdk.Coin) error {
	params := k.GetParams(ctx)

	var err error = nil
	var found bool = false
	for _, biz := range params.AuthBizs {
		if biz.Name == bizName {
			if fee.IsLT(biz.Fee) {
                		err = sdkerrors.Wrapf(types.ErrInvalidBizFee, "biz fee(%s) should be >= (%s in params)",  fee.String(), biz.Fee.String())
			}
			found = true
			break;
		}
	}
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.Wrapf(types.ErrNoBizFound, "biz(%s) is not found.", bizName)
	}

	return nil
}


func (k Keeper) getAuthorizerBizPrice(ctx sdk.Context, authorizerAddr sdk.AccAddress, bizName, denom string) sdk.Coin {
	var feeL sdk.Coin = sdk.NewCoin(denom, sdk.ZeroInt())
	
	authorizerBizs := k.GetAuthorizerBizs(ctx, authorizerAddr)
	for _, biz := range authorizerBizs.Bizs {
		if biz.Name == bizName {
			feeL = biz.Fee
			break;
		}
	}
	if feeL.IsZero() {
		params := k.GetParams(ctx)
		for _, biz := range params.AuthBizs {
			if biz.Name == bizName {
				feeL = biz.Fee
				break;
			}
		}
	}

	return feeL
}
