package keeper

import (
	"dichain/x/swap/types"
)

var _ types.QueryServer = Keeper{}
