package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgReqAuthorization = "require_authorization"
	TypeMsgWithdrawAuthorization = "withdraw_authorization"
	TypeMsgAddAuthBiz = "add_auth_biz"
	TypeMsgUpdateAuthBiz = "update_auth_biz"
	TypeMsgRemoveAuthBiz = "remove_auth_biz"
)

var (
	_ sdk.Msg = &MsgReqAuthorization{}
	_ sdk.Msg = &MsgWithdrawAuthorization{}
	_ sdk.Msg = &MsgAddAuthBiz{}
	_ sdk.Msg = &MsgUpdateAuthBiz{}
	_ sdk.Msg = &MsgRemoveAuthBiz{}
)

func NewMsgReqAuthorization(creator, authorizer sdk.AccAddress, bizName string, fee sdk.Coin) *MsgReqAuthorization {
	return &MsgReqAuthorization{
		Creator:    creator.String(),
		Authorizer: authorizer.String(),
		BizName:    bizName,
		Fee:        fee,
	}
}

func (msg *MsgReqAuthorization) Route() string {
	return RouterKey
}

func (msg *MsgReqAuthorization) Type() string {
	return TypeMsgReqAuthorization
}

func (msg *MsgReqAuthorization) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReqAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReqAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_,err = sdk.AccAddressFromBech32(msg.Authorizer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authorizer address (%s)", err)
	}
	if msg.BizName == "" {
		return errors.New("biz name can't be empty.")
	}
	if err := msg.Fee.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee (%s)", err)
	}

	return nil
}

func NewMsgWithdrawAuthorization(creator, authorizer sdk.AccAddress, bizName string) *MsgWithdrawAuthorization {
	return &MsgWithdrawAuthorization{
		Creator:    creator.String(),
		Authorizer: authorizer.String(),
		BizName:    bizName,
	}
}

func (msg *MsgWithdrawAuthorization) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawAuthorization) Type() string {
	return TypeMsgWithdrawAuthorization
}

func (msg *MsgWithdrawAuthorization) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_,err = sdk.AccAddressFromBech32(msg.Authorizer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authorizer address (%s)", err)
	}
	if msg.BizName == "" {
		return errors.New("biz name can't be empty.")
	}

	return nil
}

func NewMsgAddAuthBiz(authorizerAddr sdk.AccAddress, bizName string, fee sdk.Coin) *MsgAddAuthBiz {
	return &MsgAddAuthBiz{
		Creator: authorizerAddr.String(),
		BizName: bizName,
		Fee:     fee,
	}
}

func (msg *MsgAddAuthBiz) Route() string {
	return RouterKey
}

func (msg *MsgAddAuthBiz) Type() string {
	return TypeMsgAddAuthBiz
}

func (msg *MsgAddAuthBiz) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAuthBiz) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAuthBiz) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authorizer address (%s)", err)
	}
	if msg.BizName == "" {
		return errors.New("biz name can't be empty.")
	}
	if err := msg.Fee.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee (%s)", err)
	}
	
	return nil
}

func NewMsgUpdateAuthBiz(creator sdk.AccAddress, bizName string, fee sdk.Coin) *MsgUpdateAuthBiz {
	return &MsgUpdateAuthBiz{
		Creator: creator.String(),
		BizName: bizName,
		Fee:     fee,
	}
}

func (msg *MsgUpdateAuthBiz) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAuthBiz) Type() string {
	return TypeMsgUpdateAuthBiz
}

func (msg *MsgUpdateAuthBiz) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAuthBiz) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAuthBiz) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.BizName == "" {
		return errors.New("biz name can't be empty.")
	}
	if err := msg.Fee.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee (%s)", err)
	}
	return nil
}

func NewMsgRemoveAuthBiz(authorizerAddr sdk.AccAddress, bizName string) *MsgRemoveAuthBiz {
	return &MsgRemoveAuthBiz{
		Creator: authorizerAddr.String(),
		BizName: bizName,
	}
}

func (msg *MsgRemoveAuthBiz) Route() string {
	return RouterKey
}

func (msg *MsgRemoveAuthBiz) Type() string {
	return TypeMsgRemoveAuthBiz
}

func (msg *MsgRemoveAuthBiz) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveAuthBiz) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveAuthBiz) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authorizer address (%s)", err)
	}

	if msg.BizName == "" {
		return errors.New("biz name can't be empty.")
	}
	return nil
}
