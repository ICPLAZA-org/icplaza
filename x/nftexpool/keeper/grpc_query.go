package keeper

import (
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

var _ types.QueryServer = Keeper{}
