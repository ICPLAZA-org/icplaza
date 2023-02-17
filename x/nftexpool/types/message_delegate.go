package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDelegate{}

func NewMsgDelegate(delegator string, poolAddress string, amount sdk.Coin) *MsgDelegate {
	return &MsgDelegate{
		Delegator:   delegator,
		PoolAddress: poolAddress,
		Amount:      amount,
	}
}

func (msg *MsgDelegate) Route() string {
	return RouterKey
}

func (msg *MsgDelegate) Type() string {
	return "Delegate"
}

func (msg *MsgDelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDelegate) ValidateBasic() error {
	// check delegator
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	// check delegator
	_, err = sdk.AccAddressFromBech32(msg.PoolAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid pool address (%s)", err)
	}

	return nil
}
