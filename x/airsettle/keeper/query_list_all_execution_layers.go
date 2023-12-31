package keeper

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) ListAllExecutionLayers(goCtx context.Context, req *types.QueryListAllExecutionLayersRequest) (*types.QueryListAllExecutionLayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var executionLayers []*types.Exelayer

	executionLayerStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))

	pageRes, err := query.Paginate(executionLayerStore, req.Pagination, func(key []byte, value []byte) error {
        var executionLayer types.Exelayer
        if err := k.cdc.Unmarshal(value, &executionLayer); err != nil {
            return err
        }

        executionLayers = append(executionLayers, &executionLayer)
        return nil
    })
    
	// Throw an error if pagination failed
    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

	return &types.QueryListAllExecutionLayersResponse{Exelayer: executionLayers, Pagination: pageRes}, nil
}
