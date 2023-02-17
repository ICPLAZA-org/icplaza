package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

var _ sdk.Msg = &MsgMintNft{}

func NewMsgMintNft(sender string, recipient string, tokenId string, catId string, tokenUri string, companyUri string, valueAddedTax sdk.Dec, copyrightTax sdk.Dec, name string, components []*Component) *MsgMintNft {
	return &MsgMintNft{
		Sender:        sender,
		Recipient:     recipient,
		TokenId:       tokenId,
		CateId:        catId,
		TokenUri:      tokenUri,
		CompanyUri:    companyUri,
		ValueAddedTax: &valueAddedTax,
		CopyrightTax:  &copyrightTax,
		Components:    components,
		Name:          name,
	}
}

func (msg *MsgMintNft) Route() string {
	return RouterKey
}

func (msg *MsgMintNft) Type() string {
	return "CreateNft"
}

func (msg *MsgMintNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// check value added tax
	if msg.ValueAddedTax.IsNegative() {
		return sdkerrors.Wrapf(ErrInvalidValueAddedTax, "(%s) should be positive", msg.ValueAddedTax)
	}
	if msg.ValueAddedTax.GT(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(ErrInvalidValueAddedTax, "should not great than (%s)", msg.ValueAddedTax)
	}
	// check copyright tax
	if msg.CopyrightTax.IsNegative() {
		return sdkerrors.Wrapf(ErrInvalidCopyrightTax, "(%s) should be positive", msg.CopyrightTax)
	}
	if msg.CopyrightTax.GT(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(ErrInvalidCopyrightTax, "should not great than (%s)", msg.CopyrightTax)
	}

	if msg.ValueAddedTax.Add(*msg.CopyrightTax).GT(sdk.NewDec(1)) {
		return ErrInvalidSum
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(ErrInvalidParam, "name required")
	}
	if strings.Contains(msg.TokenId, "-") {
		return ErrInvalidTokenFormat
	}
	return nil
}
