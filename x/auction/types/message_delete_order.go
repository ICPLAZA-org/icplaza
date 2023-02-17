package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteOrder{}

func NewMsgDeleteOrder(creator, poolAddress, tokenId string) *MsgDeleteOrder {
	return &MsgDeleteOrder{
		Creator:     creator,
		TokenId:     tokenId,
		PoolAddress: poolAddress,
	}
}

func (msg *MsgDeleteOrder) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOrder) Type() string {
	return "DeleteOrder"
}

func (msg *MsgDeleteOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// pool address
	_, err = sdk.AccAddressFromBech32(msg.PoolAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid pool address (%s)", err)
	}
	return nil
}
