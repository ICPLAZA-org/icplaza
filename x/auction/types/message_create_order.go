package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(sender string, tokenId string, poolAddress string, startPrice sdk.Coin, minStep sdk.Coin, startTime time.Time, endTime time.Time, autoAgreePeriod time.Duration) *MsgCreateOrder {
	return &MsgCreateOrder{
		Sender:          sender,
		TokenId:         tokenId,
		PoolAddress:     poolAddress,
		StartPrice:      startPrice,
		MinStep:         minStep,
		StartTime:       startTime,
		EndTime:         endTime,
		AutoAgreePeriod: autoAgreePeriod,
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
	// check min step
	if !msg.MinStep.IsPositive() {
		return sdkerrors.Wrapf(ErrInvalidParam, "need a positive step price")
	}
	// check start time
	if msg.StartTime.Before(time.Now()) {
		return sdkerrors.Wrapf(ErrInvalidParam, "start time need after now ")
	}
	// check end time
	if msg.StartTime.After(msg.EndTime) {
		return sdkerrors.Wrapf(ErrInvalidParam, "start time need before end time ")
	}
	return nil
}


