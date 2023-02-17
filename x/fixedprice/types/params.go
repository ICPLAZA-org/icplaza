package types

import (
	"fmt"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
	"time"
)

// Parameter store keys
var (
	KeyAdjustPricePeriod = []byte("AdjustPricePeriod")
)

var (
	DefaultAdjustPricePeriod, _ = time.ParseDuration("3600s")
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default liquidity module parameters.
func DefaultParams() Params {
	return Params{
		AdjustPricePeriod: DefaultAdjustPricePeriod,
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyAdjustPricePeriod, &p.AdjustPricePeriod, validateAdjustPricePeriod),
	}
}
func validateAdjustPricePeriod(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.Seconds() <= 0 {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate validates parameters.
func (p Params) Validate() error {
	if p.AdjustPricePeriod.Seconds() < 0 {
		return fmt.Errorf("invalid AdjustPricePeriod type: %s", p.AdjustPricePeriod.String())
	}
	return nil
}
