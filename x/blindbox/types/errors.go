package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/blindbox module sentinel errors
var (
	InvalidParam         = sdkerrors.Register(ModuleName, 2, "invalid param")
	ErrExpired           = sdkerrors.Register(ModuleName, 3, "expired")
	ErrAllOpened         = sdkerrors.Register(ModuleName, 4, "all of boxes are opened")
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 5, "insufficient funds")
	ErrNotFound          = sdkerrors.Register(ModuleName, 6, "not found")
	ErrNotStart          = sdkerrors.Register(ModuleName, 7, "not start")
	ErrInvalidCreator    = sdkerrors.Register(ModuleName, 8, "invalid creator")
	ErrInvalidNFT        = sdkerrors.Register(ModuleName, 9, "invalid nft")
	ErrInvalidNFTOwner   = sdkerrors.Register(ModuleName, 10, "invalid nft Owner")
	ErrRemove            = sdkerrors.Register(ModuleName, 11, "remove error")
	ErrRevoke            = sdkerrors.Register(ModuleName, 12, "Revoke error")
)
