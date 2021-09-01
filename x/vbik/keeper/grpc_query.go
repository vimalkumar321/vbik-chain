package keeper

import (
	"github.com/example/vbik/x/vbik/types"
)

var _ types.QueryServer = Keeper{}
