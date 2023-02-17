package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferNft{}

func NewMsgTransferNft(sender, recipient, cateId, tokenId string) *MsgTransferNft {
	return &MsgTransferNft{
		Sender:    sender,
		Recipient: recipient,
		CateId:    cateId,
		TokenId:   tokenId,
	}
}

func (msg *MsgTransferNft) Route() string {
	return RouterKey
}

func (msg *MsgTransferNft) Type() string {
	return "TransferNft"
}

func (msg *MsgTransferNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	if msg.TokenId == "" {
		return sdkerrors.Wrapf(ErrInvalidParam, "need TokenId")
	}
	return nil
}
