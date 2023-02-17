package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Parameter store keys
var (
	KeyPoolCreationFee   = []byte("PoolCreationFee")
	KeyBurnPercent       = []byte("BurnPercent")
	KeyCommunityPercent  = []byte("CommunityPercent")
	KeyValidatorsPercent = []byte("ValidatorsPercent")
)

var (
	DefaultPoolCreationFee      = sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000)
	DefaultBurnPercent          = sdk.NewDecWithPrec(0, 1)
	DefaultCommunityPercent     = sdk.NewDecWithPrec(0, 1)
	DefaultValidatorsPercent    = sdk.NewDecWithPrec(0, 1)
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default liquidity module parameters.
func DefaultParams() Params {
	return Params{
		PoolCreationFee:   DefaultPoolCreationFee,
		BurnPercent:       DefaultBurnPercent,
		CommunityPercent:  DefaultCommunityPercent,
		ValidatorsPercent: DefaultValidatorsPercent,
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyPoolCreationFee, &p.PoolCreationFee, validatePoolCreationFee),
		paramstypes.NewParamSetPair(KeyBurnPercent, &p.BurnPercent, validateBurnPercent),
		paramstypes.NewParamSetPair(KeyCommunityPercent, &p.CommunityPercent, validateCommunityPercent),
		paramstypes.NewParamSetPair(KeyValidatorsPercent, &p.ValidatorsPercent, validateValidatorsPercent),
	}
}
func validatePoolCreationFee(i interface{}) error {
	return nil
}
func validateBurnPercent(i interface{}) error {
	return nil
}
func validateCommunityPercent(i interface{}) error {
	return nil
}
func validateValidatorsPercent(i interface{}) error {
	return nil
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate validates parameters.
func (p Params) Validate() error {
	return nil
}
