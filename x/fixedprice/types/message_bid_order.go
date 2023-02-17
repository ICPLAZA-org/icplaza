package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/nft/types"
)

var _ sdk.Msg = &MsgBidOrder{}

func NewMsgBidOrder(sender string, tokenId string, price sdk.Coin, poolAddress string) *MsgBidOrder {
	return &MsgBidOrder{
		Sender:      sender,
		TokenId:     tokenId,
		Price:       &price,
		PoolAddress: poolAddress,
	}
}

func (msg *MsgBidOrder) Route() string {
	return RouterKey
}

func (msg *MsgBidOrder) Type() string {
	return "BidOrder"
}

func (msg *MsgBidOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBidOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBidOrder) ValidateBasic() error {
	// check sender
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// check pool
	_, err = sdk.AccAddressFromBech32(msg.PoolAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid pool address (%s)", err)
	}
	// check price
	if msg.Price.IsNegative() {
		return sdkerrors.Wrapf(types.ErrInvalidBidPrice, msg.Price.String())
	}
	return nil
}
