package keeper

import (
	"github.com/evmos/evmos/v11/x/nft/types"
)

var _ types.QueryServer = Keeper{}
