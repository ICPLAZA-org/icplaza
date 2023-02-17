package keeper

import (
	"github.com/gauss/gauss/v6/x/fixedprice/types"
)

var _ types.QueryServer = Keeper{}
