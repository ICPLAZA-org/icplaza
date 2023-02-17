package types

import (
	"time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBox      = "create_box"
	TypeMsgCreatePool     = "create_pool"
	TypeMsgOpenBox        = "open_box"
	TypeMsgRemovePool     = "remove_pool"
	TypeMsgRevokeBoxGroup = "revoke_box_group"
	TypeMsgUpdatePool     = "update_pool"
)

var (
	_ sdk.Msg = &MsgCreateBox{}
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgOpenBox{}
	_ sdk.Msg = &MsgRemovePool{}
	_ sdk.Msg = &MsgRevokeBoxGroup{}
	_ sdk.Msg = &MsgUpdatePool{}
)

func NewMsgCreateBox(creator string, groupId string, boxSize uint64, boxPrice sdk.Coin, startTime time.Time, endTime time.Time, randomMin uint64, randomMax uint64, randomNfts []string, fixedNfts []string, poolId string) *MsgCreateBox {
	return &MsgCreateBox{
		Creator:    creator,
		GroupId:    groupId,
		BoxSize:    boxSize,
		BoxPrice:   &boxPrice,
		StartTime:  startTime,
		EndTime:    endTime,
		RandomMin:  randomMin,
		RandomMax:  randomMax,
		RandomNfts: randomNfts,
		FixedNfts:  fixedNfts,
		PoolId:     poolId,
	}
}

func (msg *MsgCreateBox) Route() string {
	return RouterKey
}

func (msg *MsgCreateBox) Type() string {
	return TypeMsgCreateBox
}

func (msg *MsgCreateBox) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBox) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBox) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.StartTime.Before(time.Now()) {
		return sdkerrors.Wrapf(InvalidParam, "startTime(%s) must after now(%s)", msg.StartTime, time.Now().UTC())
	}
	if msg.EndTime.Before(time.Now()) {
		return sdkerrors.Wrapf(InvalidParam, "endTime(%s) must after now(%s)", msg.EndTime, time.Now().UTC())
	}
	if msg.StartTime.After(msg.EndTime) {
		return sdkerrors.Wrapf(InvalidParam, "startTime should before endTime")
	}
	if msg.BoxSize == 0 {
		return sdkerrors.Wrapf(InvalidParam, "boxSize should greater than 0")
	}
	if len(msg.FixedNfts) == 0 && len(msg.RandomNfts) == 0 {
		return sdkerrors.Wrapf(InvalidParam, "need nft list")
	}
	// trim repeated nft
	if len(msg.FixedNfts) > 0 {
		fixedNftMap := make(map[string]bool, 0)
		for _, nft := range msg.FixedNfts {
			fixedNftMap[nft] = true
		}
		fixedNFTSize := len(fixedNftMap)
		if len(msg.FixedNfts) != fixedNFTSize {
			return sdkerrors.Wrapf(InvalidParam, "please trim repeated nft in fixed list")
		}
		if fixedNFTSize > 0 {
			m := uint64(fixedNFTSize) / msg.BoxSize
			m = uint64(fixedNFTSize) - msg.BoxSize*m
			if m != 0 {
				return sdkerrors.Wrapf(InvalidParam, "boxSize(%d) not match fixedNfts(%d)", msg.BoxSize, fixedNFTSize)
			}
		}
	}
	// trim repeated nft
	if len(msg.RandomNfts) > 0 {
		if msg.RandomMin > msg.RandomMax {
			return sdkerrors.Wrapf(InvalidParam, "randomMin(%d) should less than randomMax(%d)", msg.RandomMin, msg.RandomMax)
		}
		randomNFTMap := make(map[string]bool, 0)
		for _, nft := range msg.RandomNfts {
			randomNFTMap[nft] = true
		}
		randomNFTSize := len(randomNFTMap)
		if len(msg.RandomNfts) != randomNFTSize {
			return sdkerrors.Wrapf(InvalidParam, "please trim repeated nft in random list")
		}
		if msg.RandomMin > 0 {
			minSize := msg.RandomMin * msg.BoxSize
			if minSize > uint64(randomNFTSize) {
				return sdkerrors.Wrapf(InvalidParam, "count of random nft should >= %d", minSize)
			}
		}
		if int(msg.RandomMax)*int(msg.BoxSize) < len(msg.RandomNfts) {
			return sdkerrors.Wrapf(InvalidParam, "count of random nft should less than", int(msg.RandomMax)*int(msg.BoxSize))
		}
	}
	// check repeated nft in fixed and random
	allNFT := make(map[string]bool, 0)
	for _, nft := range msg.FixedNfts {
		allNFT[nft] = true
	}
	for _, nft := range msg.RandomNfts {
		allNFT[nft] = true
	}
	if len(allNFT) != len(msg.FixedNfts)+len(msg.RandomNfts) {
		return sdkerrors.Wrapf(InvalidParam, "some nft are repeated between fixed and random list")
	}

	if msg.PoolId == "" {
		return sdkerrors.Wrapf(InvalidParam, "need poolId")
	}

	if msg.GroupId == "" {
		return sdkerrors.Wrapf(InvalidParam, "need groupId")
	}

	if msg.BoxPrice == nil {
		return sdkerrors.Wrapf(InvalidParam, "invalid price")
	}
	return nil
}

func NewMsgCreatePool(creator string, feeRate sdk.Dec, feeAddress string, poolId string) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:    creator,
		FeeRate:    &feeRate,
		FeeAddress: feeAddress,
		PoolId:     poolId,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.FeeAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fee address (%s)", err)
	}

	if msg.FeeRate.GTE(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(InvalidParam, "feeRate should less than 100%")
	}

	return nil
}

func NewMsgOpenBox(creator string, groupId string, boxIds []string) *MsgOpenBox {
	return &MsgOpenBox{
		Creator: creator,
		GroupId: groupId,
		BoxIds:  boxIds,
	}
}

func (msg *MsgOpenBox) Route() string {
	return RouterKey
}

func (msg *MsgOpenBox) Type() string {
	return TypeMsgOpenBox
}

func (msg *MsgOpenBox) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOpenBox) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOpenBox) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// check repeated boxId
	list := make(map[string]bool, 0)
	for _, boxId := range msg.BoxIds {
		list[boxId] = true
	}
	if len(list) != len(msg.BoxIds) {
		return sdkerrors.Wrapf(InvalidParam, "please strip repeated boxId", err)
	}
	return nil
}

func NewMsgRemovePool(creator string, poolId string) *MsgRemovePool {
	return &MsgRemovePool{
		Creator: creator,
		PoolId:  poolId,
	}
}

func (msg *MsgRemovePool) Route() string {
	return RouterKey
}

func (msg *MsgRemovePool) Type() string {
	return TypeMsgRemovePool
}

func (msg *MsgRemovePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemovePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemovePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgRevokeBoxGroup(creator string, groupId string) *MsgRevokeBoxGroup {
	return &MsgRevokeBoxGroup{
		Creator: creator,
		GroupId: groupId,
	}
}

func (msg *MsgRevokeBoxGroup) Route() string {
	return RouterKey
}

func (msg *MsgRevokeBoxGroup) Type() string {
	return TypeMsgRevokeBoxGroup
}

func (msg *MsgRevokeBoxGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeBoxGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeBoxGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgUpdatePool(creator string, feeRate sdk.Dec, feeAddress string, poolId string) *MsgUpdatePool {
	return &MsgUpdatePool{
		Creator:    creator,
		FeeRate:    &feeRate,
		FeeAddress: feeAddress,
		PoolId:     poolId,
	}
}

func (msg *MsgUpdatePool) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePool) Type() string {
	return TypeMsgUpdatePool
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.FeeAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid FeeAddress (%s)", err)
	}
	if msg.FeeRate.GTE(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(InvalidParam, "fee rate should less than 100%")
	}
	return nil
}
