package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUndelegate{}

func NewMsgUndelegate(sender string, delegator string, poolAddress string, amount sdk.Coin) *MsgUndelegate {
	return &MsgUndelegate{
		Sender:      sender,
		Delegator:   delegator,
		PoolAddress: poolAddress,
		Amount:      amount,
	}
}

func (msg *MsgUndelegate) Route() string {
	return RouterKey
}

func (msg *MsgUndelegate) Type() string {
	return "Undelegate"
}

func (msg *MsgUndelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUndelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUndelegate) ValidateBasic() error {
	// check address
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	// check address
	_, err = sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	// check address
	_, err = sdk.AccAddressFromBech32(msg.PoolAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid pool address (%s)", err)
	}
	// check amount
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(ErrInvalidParam, " amount should be greater than 0")
	}
	return nil
}
