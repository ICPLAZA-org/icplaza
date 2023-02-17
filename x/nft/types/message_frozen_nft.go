package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFrozenNft{}

func NewMsgFrozenNft(sender, cateId, tokenId string) *MsgFrozenNft {
	return &MsgFrozenNft{
		Sender:  sender,
		CateId:  cateId,
		TokenId: tokenId,
	}
}

func (msg *MsgFrozenNft) Route() string {
	return RouterKey
}

func (msg *MsgFrozenNft) Type() string {
	return "FrozenNft"
}

func (msg *MsgFrozenNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFrozenNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFrozenNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
