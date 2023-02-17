package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidParam         = sdkerrors.Register(ModuleName, 1, "invalid param")
	ErrInvalidPool          = sdkerrors.Register(ModuleName, 2, "invalid pool")
	ErrInvalidPoolExists    = sdkerrors.Register(ModuleName, 3, "pool already exists")
	ErrInvalidPoolNotExists = sdkerrors.Register(ModuleName, 4, "pool not exists")
	ErrNoDelegation         = sdkerrors.Register(ModuleName, 5, "no delegation")
	ErrInsufficientFunds    = sdkerrors.Register(ModuleName, 6, "insufficient funds")
	ErrInvalidCreator       = sdkerrors.Register(ModuleName, 7, "invalid creator")
	ErrInvalidPoolCreator   = sdkerrors.Register(ModuleName, 8, "invalid pool creator")
	ErrOutOfDelegation      = sdkerrors.Register(ModuleName, 9, "out of delegate amount")
)
