package types

import (
	"fmt"
	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyAuthBizs = []byte("AuthBizs")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(authBizs DaoBizs) Params {
	return Params {
		AuthBizs: authBizs,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	authBizs := DaoBizs{}
	authBizs = append(authBizs, DaoBiz{Name: IssueErc20BizKey, Fee: sdk.NewCoin(sdk.DefaultBondDenom, sdk.ZeroInt())})

	return Params{
		AuthBizs: authBizs,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAuthBizs, &p.AuthBizs, validateAuthBizs),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateAuthBizs(p.AuthBizs); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateAuthBizs(v interface{}) error {
	authBizs, ok := v.([]DaoBiz)	
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if len(authBizs) > 0 {
		authBizsMap := make(map[string]bool, 0)
		for _, authBiz := range authBizs {
			if authBizsMap[authBiz.Name] {
				return sdkerrors.Wrapf(ErrInvalidParams, "please strip repeated params")
			}
			authBizsMap[authBiz.Name] = true
		}
	}

	return nil
}
