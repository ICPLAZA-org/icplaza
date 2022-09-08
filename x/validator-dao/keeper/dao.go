package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/icplaza/icplaza/v6/x/validator-dao/types"
)

// ConsumeAuthorization pay for power
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

// consumeAuthorization pay charges
func (k Keeper) consumeAuthorization(ctx sdk.Context, authBiz *types.AcqBiz) (bool) {
	rc := false
	fee := authBiz.Price

	count := authBiz.Amount.Amount.Quo(authBiz.Price.Amount)
	if count.LT(sdk.OneInt()) {
		return rc
	}
	authBiz.Amount = authBiz.Amount.Sub(fee)

	// pay charges
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

// ReqAuthorization prepay fees to purchase permission at any time
func (k Keeper) ReqAuthorization(ctx sdk.Context, granteeAddr, authorizerAddr sdk.AccAddress, bizName string, fee sdk.Coin) error {
	// grantee can only purchase the same service to a certain validator once at the same time
	granteeAuthBizs := k.GetGranteeAuthBizs(ctx, granteeAddr)
	for _, authBiz := range granteeAuthBizs.Bizs {
		if (bizName == authBiz.BizName) && (authorizerAddr.String() == authBiz.From) {
			return sdkerrors.Wrapf(types.ErrAuthorizationFound, "biz(%s)'s authorization from %s does exist.", bizName, authBiz.From)
		}
	}
	
	// check: fees >= price
	price := k.getAuthorizerBizPrice(ctx, authorizerAddr, bizName, fee.Denom)
	if price.IsZero() {
		return sdkerrors.Wrapf(types.ErrNoBizFound, "biz(%s) does not exist.", bizName)
	}
	if fee.IsLT(price) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "got: %s, expect: %s", fee.String(), price.String())
	}

	// escrow to the module address
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, granteeAddr, types.ModuleName, sdk.NewCoins(fee)); err != nil {
		return err
	}
	
	granteeAuthBizs.Bizs = append(granteeAuthBizs.Bizs, types.NewAcqBiz(authorizerAddr.String(), bizName, fee, price))
	k.SetGranteeAuthBizs(ctx, granteeAddr, granteeAuthBizs)

	return nil
}

// WithdrawAuthorization withdraw grantee's escrow coins
func (k Keeper) WithdrawAuthorization(ctx sdk.Context, granteeAddr, authorizerAddr sdk.AccAddress, bizName string) (sdk.Coin, error) {
	err := nil
	rc := sdk.ZeroCoin()

	granteeAuthBizs := k.GetGranteeAuthBizs(ctx, granteeAddr)
	for i, authBiz := range granteeAuthBizs.Bizs {
		if (bizName == authBiz.BizName) && (authorizerAddr.String() == authBiz.From) {
			rc = authBiz.Amount

			granteeAuthBizs.Bizs = append(granteeAuthBizs.Bizs[:i], granteeAuthBizs.Bizs[i+1:]...)
			k.SetGranteeAuthBizs(ctx, granteeAddr, granteeAuthBizs)

			break;
		}
	}

	if rc.IsGT(sdk.ZeroCoin()) {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, granteeAddr, sdk.NewCoins(rc))
	}

	return rc, err
}


// AddAuthBiz add service into validator's service list
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

// UpdateAuthBiz update price of the validator's service
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

// RemoveAuthBiz remove service from validator's service list
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

// getAuthorizerBizPrice return the price of a centain service from the validator price list or the network price list
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
