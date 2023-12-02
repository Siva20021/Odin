package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "odin/testutil/keeper"
	"odin/x/odin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OdinKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
