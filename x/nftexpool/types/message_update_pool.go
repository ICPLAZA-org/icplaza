package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdatePool{}

func NewMsgUpdatePool(creator string, poolAddress string, commissionRate sdk.Dec, commissionAddress string, valueAddedTaxAddress string) *MsgUpdatePool {
	return &MsgUpdatePool{
		Creator:              creator,
		PoolAddress:          poolAddress,
		CommissionRate:       commissionRate,
		CommissionAddress:    commissionAddress,
		ValueAddedTaxAddress: valueAddedTaxAddress,
	}
}

func (msg *MsgUpdatePool) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePool) Type() string {
	return "UpdatePool"
}

func (msg *MsgUpdatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePool) ValidateBasic() error {
	// check creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// check pool address
	_, err = sdk.AccAddressFromBech32(msg.PoolAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid pool address (%s)", err)
	}
	// check commission address
	_, err = sdk.AccAddressFromBech32(msg.CommissionAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid commission address (%s)", err)
	}
	// check value added address
	_, err = sdk.AccAddressFromBech32(msg.ValueAddedTaxAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid value added tax address (%s)", err)
	}
	// check CommissionRate
	if msg.CommissionRate.IsNegative() {
		return sdkerrors.Wrap(ErrInvalidParam, "commission rate should be positive")
	}
	// CommissionRate should be less than 100%
	if msg.CommissionRate.GTE(sdk.NewDec(1)) {
		return sdkerrors.Wrap(ErrInvalidParam, "commission rate should be less than 100%")
	}
	return nil
}
