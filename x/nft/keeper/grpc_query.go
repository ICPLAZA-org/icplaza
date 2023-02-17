package keeper

import (
	"github.com/gauss/gauss/v6/x/nft/types"
)

var _ types.QueryServer = Keeper{}
