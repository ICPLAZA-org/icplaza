package keeper

import (
	"github.com/evmos/evmos/v11/x/auction/types"
)

var _ types.QueryServer = Keeper{}
