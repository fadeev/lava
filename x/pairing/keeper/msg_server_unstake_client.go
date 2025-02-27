package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/pairing/types"
)

func (k msgServer) UnstakeClient(goCtx context.Context, msg *types.MsgUnstakeClient) (*types.MsgUnstakeClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.UnstakeEntry(ctx, false, msg.ChainID, msg.Creator, types.UnstakeDescriptionClientUnstake)
	return &types.MsgUnstakeClientResponse{}, err
}
