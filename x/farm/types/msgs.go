package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TypeMsgCreatePool is the type for MsgCreatePool
	TypeMsgCreatePool = "create_pool"

	// TypeMsgDestroyPool is the type for MsgDestroyPool
	TypeMsgDestroyPool = "destroy_pool"

	// TypeMsgStake is the type for MsgStake
	TypeMsgStake = "stake"

	// TypeMsgUnstake is the type for MsgUnstake
	TypeMsgUnstake = "unbond"

	// TypeMsgHarvest is the type for MsgHarvest
	TypeMsgHarvest = "withdraw"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgDestroyPool{}
	_ sdk.Msg = &MsgStake{}
	_ sdk.Msg = &MsgUnbond{}
	_ sdk.Msg = &MsgWithdraw{}
)

// Route implements Msg
func (msg MsgCreatePool) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgCreatePool) Type() string { return TypeMsgCreatePool }

// ValidateBasic implements Msg
func (msg MsgCreatePool) ValidateBasic() error {
	if err := ValidatePoolName(msg.Name); err != nil {
		return err
	}

	if err := ValidateDescription(msg.Description); err != nil {
		return err
	}

	if err := ValidateLpTokenDenom(msg.MinStaking.Denom); err != nil {
		return err
	}

	if err := ValidateCoins("RewardPerBlock", msg.RewardsPerBlock...); err != nil {
		return err
	}

	if err := ValidateAddress(msg.Creator); err != nil {
		return err
	}

	if err := ValidateCoins("TotalReward", msg.TotalRewards...); err != nil {
		return err
	}
	return ValidateReward(msg.RewardsPerBlock, msg.TotalRewards)
}

// GetSignBytes implements Msg
func (msg MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgDestroyPool) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgDestroyPool) Type() string { return TypeMsgDestroyPool }

// ValidateBasic implements Msg
func (msg MsgDestroyPool) ValidateBasic() error {
	if err := ValidateAddress(msg.Creator); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgDestroyPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgDestroyPool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgStake) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgStake) Type() string { return TypeMsgStake }

// ValidateBasic implements Msg
func (msg MsgStake) ValidateBasic() error {
	if err := ValidateAddress(msg.Sender); err != nil {
		return err
	}

	if err := ValidateCoins("Amount", msg.Amount); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgStake) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgUnbond) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgUnbond) Type() string { return TypeMsgUnstake }

// ValidateBasic implements Msg
func (msg MsgUnbond) ValidateBasic() error {
	if err := ValidateAddress(msg.Sender); err != nil {
		return err
	}

	if err := ValidateCoins("Amount", msg.Amount); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgUnbond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUnbond) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgWithdraw) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgWithdraw) Type() string { return TypeMsgHarvest }

// ValidateBasic implements Msg
func (msg MsgWithdraw) ValidateBasic() error {
	if err := ValidateAddress(msg.Sender); err != nil {
		return err
	}

	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
