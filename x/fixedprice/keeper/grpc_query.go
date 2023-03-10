package keeper

import (
	"github.com/evmos/evmos/v11/x/fixedprice/types"
)

var _ types.QueryServer = Keeper{}
