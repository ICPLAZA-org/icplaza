package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	ErrInvalidNftOwner   = sdkerrors.Register(ModuleName, 1, "invalid nft owner")
	ErrInvalidParam      = sdkerrors.Register(ModuleName, 2, "invalid param")
	ErrInvalidOrder      = sdkerrors.Register(ModuleName, 3, "order not exists")
	ErrOrderBidding      = sdkerrors.Register(ModuleName, 4, "order is bidding")
	ErrOrderExpired      = sdkerrors.Register(ModuleName, 5, "order is expired")
	ErrOrderNotStart     = sdkerrors.Register(ModuleName, 6, "order is not start")
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 7, "insufficient funds")
	ErrDenomMatch        = sdkerrors.Register(ModuleName, 8, "bid price denom not match order price")
	ErrInvalidBidPrice   = sdkerrors.Register(ModuleName, 9, "invalid bid price")
	ErrInvalidPool       = sdkerrors.Register(ModuleName, 10, "pool not exists")
	// this line is used by starport scaffolding # ibc/errors
)
