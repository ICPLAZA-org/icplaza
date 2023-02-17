package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintBatch{}

func NewMsgMintBatch(sender string, recipient string, items []*MintBatchItem) *MsgMintBatch {
	return &MsgMintBatch{
		Sender:    sender,
		Recipient: recipient,
		Items:     items,
	}
}

func (msg *MsgMintBatch) Route() string {
	return RouterKey
}

func (msg *MsgMintBatch) Type() string {
	return "MintBatch"
}

func (msg *MsgMintBatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receipt address (%s)", err)
	}
	// 检测每个nft具体信息
	for _, item := range msg.Items {
		// check value added tax
		if item.ValueAddedTax.IsNegative() {
			return sdkerrors.Wrapf(ErrInvalidValueAddedTax, "(%s) should be positive", item.ValueAddedTax)
		}
		if item.ValueAddedTax.GT(sdk.NewDec(1)) {
			return sdkerrors.Wrapf(ErrInvalidValueAddedTax, "should not great than (%s)", item.ValueAddedTax)
		}
		// check copyright tax
		if item.CopyrightTax.IsNegative() {
			return sdkerrors.Wrapf(ErrInvalidCopyrightTax, "(%s) should be positive", item.CopyrightTax)
		}
		if item.CopyrightTax.GT(sdk.NewDec(1)) {
			return sdkerrors.Wrapf(ErrInvalidCopyrightTax, "should not great than (%s)", item.CopyrightTax)
		}
		if item.Name == "" {
			return sdkerrors.Wrap(ErrInvalidParam, "name required")
		}
		if item.ValueAddedTax.Add(*item.CopyrightTax).GT(sdk.NewDec(1)) {
			return ErrInvalidSum
		}
		if item.TokenId == "" {
			return sdkerrors.Wrap(ErrInvalidParam, " TokenId is null")
		}
		if item.CateId == "" {
			return sdkerrors.Wrap(ErrInvalidParam, " CateId is null")
		}
		if item.TokenUri == "" {
			return sdkerrors.Wrap(ErrInvalidParam, " TokenUri is null")
		}
		if item.CompanyUri == "" {
			return sdkerrors.Wrap(ErrInvalidParam, " CompanyUri is null")
		}
	}

	return nil
}
