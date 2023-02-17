package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(sender string, tokenId string, poolAddress string, startPrice, endPrice sdk.Coin, startTime time.Time, endTime time.Time) *MsgCreateOrder {
	return &MsgCreateOrder{
		Sender:      sender,
		TokenId:     tokenId,
		PoolAddress: poolAddress,
		StartPrice:  startPrice,
		EndPrice:    endPrice,
		StartTime:   startTime,
		EndTime:     endTime,
	}
}

func (msg *MsgCreateOrder) Route() string {
	return RouterKey
}

func (msg *MsgCreateOrder) Type() string {
	return "CreateOrder"
}

func (msg *MsgCreateOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	// check sender
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	// check pool
	_, err = sdk.AccAddressFromBech32(msg.PoolAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid pool address (%s)", err)
	}
	// check start price
	if !msg.StartPrice.IsPositive() {
		return sdkerrors.Wrapf(ErrInvalidParam, "need a positive start price")
	}
	// check end price
	if !msg.EndPrice.IsPositive() {
		return sdkerrors.Wrapf(ErrInvalidParam, "need a positive end price")
	}
	// check start time
	if msg.StartTime.Before(time.Now()) {
		return sdkerrors.Wrapf(ErrInvalidParam, "start time need after now ")
	}
	// check end time
	if msg.EndTime.Before(msg.StartTime) {
		return sdkerrors.Wrapf(ErrInvalidParam, "start time need after end time ")
	}
	// check denom
	if msg.StartPrice.Denom != msg.EndPrice.Denom {
		return sdkerrors.Wrapf(ErrInvalidParam, "startPrice denom need same as endPrice denom")
	}
	return nil
}
