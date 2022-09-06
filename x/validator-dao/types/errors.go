package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/validatordao module sentinel errors
var (
	ErrAuthorizationFound  = sdkerrors.Register(ModuleName, 1, "authorization already exist")
	ErrBizFound            = sdkerrors.Register(ModuleName, 2, "biz already exist")
	ErrInvalidAuthorizer   = sdkerrors.Register(ModuleName, 3, "invalid authorizer")
	ErrInvalidBizFee       = sdkerrors.Register(ModuleName, 4, "invalid biz fee")
	ErrInvalidBizName      = sdkerrors.Register(ModuleName, 5, "invalid biz name")
	ErrInvalidParams       = sdkerrors.Register(ModuleName, 6, "invalid params")
	ErrNoBizFound          = sdkerrors.Register(ModuleName, 7, "biz does not exist")
	ErrNoValidatorFound    = sdkerrors.Register(ModuleName, 8, "validator does not exist")
)
