package keeper

import (
	"context"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAllExecutionLayers(goCtx context.Context, req *types.QueryListAllExecutionLayersRequest) (*types.QueryListAllExecutionLayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	


	// TODO: Process the query
	_ = ctx

	return &types.QueryListAllExecutionLayersResponse{}, nil
}
