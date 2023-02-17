package keeper

import (
	"github.com/gauss/gauss/v6/x/auction/types"
)

var _ types.QueryServer = Keeper{}
