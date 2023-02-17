package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrInvalidValueAddedTax = sdkerrors.Register(ModuleName, 1, "invalid ValueAddedTax value")
	ErrInvalidCopyrightTax  = sdkerrors.Register(ModuleName, 2, "invalid CopyrightTax value")
	ErrInvalidSum           = sdkerrors.Register(ModuleName, 3, "sum of CopyrightTax and CopyrightTax must less than 1")
	ErrInvalidCollection    = sdkerrors.Register(ModuleName, 4, "invalid NFT collection")
	ErrUnknownCollection    = sdkerrors.Register(ModuleName, 5, "unknown NFT collection")
	ErrInvalidNFT           = sdkerrors.Register(ModuleName, 6, "invalid NFT")
	ErrUnknownNFT           = sdkerrors.Register(ModuleName, 7, "unknown NFT")
	ErrNFTAlreadyExists     = sdkerrors.Register(ModuleName, 8, "NFT already exists")
	ErrEmptyMetadata        = sdkerrors.Register(ModuleName, 9, "NFT metadata can't be empty")
	ErrInvalidOwner         = sdkerrors.Register(ModuleName, 10, "invalid owner")
	ErrNFTLocked            = sdkerrors.Register(ModuleName, 11, "nft is locked")
	ErrNFTFrozen            = sdkerrors.Register(ModuleName, 12, "nft is frozen")
	ErrPoolCountOver        = sdkerrors.Register(ModuleName, 13, "already have pool with poolType 1")
	ErrPoolExists           = sdkerrors.Register(ModuleName, 14, "pool already exists")
	ErrInvalidPoolType      = sdkerrors.Register(ModuleName, 15, "poolType should be 0 or 1")
	ErrInvalidPollAddress   = sdkerrors.Register(ModuleName, 16, "invalid pool")
	ErrPoolNotExists        = sdkerrors.Register(ModuleName, 17, "pool not exists")
	ErrBadCommissionRate    = sdkerrors.Register(ModuleName, 18, "invalid commission rate")
	ErrBadSigner            = sdkerrors.Register(ModuleName, 19, "invalid signer")
	ErrInvalidCreator       = sdkerrors.Register(ModuleName, 20, "invalid creator")
	ErrInvalidNFTOwner      = sdkerrors.Register(ModuleName, 21, "invalid NFT Owner")
	ErrOrderExists          = sdkerrors.Register(ModuleName, 22, "order already exists")
	ErrOrderNotExists       = sdkerrors.Register(ModuleName, 23, "order not exists")
	ErrInsufficientFunds    = sdkerrors.Register(ModuleName, 24, "insufficient funds")
	ErrInvalidBidPrice      = sdkerrors.Register(ModuleName, 25, "invalid bid price")
	ErrPoolLocking          = sdkerrors.Register(ModuleName, 26, "pool is locking")
	ErrInvalidPollCreator   = sdkerrors.Register(ModuleName, 27, "invalid pool creator")
	ErrInvalidOrder         = sdkerrors.Register(ModuleName, 28, "invalid order")
	ErrInvalidPoolCreator   = sdkerrors.Register(ModuleName, 29, "invalid pool creator")
	ErrNoDelegation         = sdkerrors.Register(ModuleName, 30, "no delegation")
	ErrOutOfDelegation      = sdkerrors.Register(ModuleName, 31, "out of delegate amount")
	ErrPrice                = sdkerrors.Register(ModuleName, 32, "price must be positive")
	ErrInvalidStartID       = sdkerrors.Register(ModuleName, 33, "invalid startID")
	ErrInvalidTokenFormat   = sdkerrors.Register(ModuleName, 34, "invalid token format")
	ErrCannotLockOrFrozen   = sdkerrors.Register(ModuleName, 35, "only free nft can lock or frozen")
	ErrInvalidParam         = sdkerrors.Register(ModuleName, 36, "invalid param")
	ErrNotFree              = sdkerrors.Register(ModuleName, 37, "nft not free")
)
