package keeper

import (
	"context"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k Keeper) ShowExecutionLayerById(goCtx context.Context, req *types.QueryShowExecutionLayerByIdRequest) (*types.QueryShowExecutionLayerByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	chain_id := req.Id 
	if chain_id == "" { 
		// no execution layer on this address
		return nil, sdkerrors.ErrKeyNotFound
	}

	exelayer, found := k.GetExelayerById(ctx, chain_id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowExecutionLayerByIdResponse{Exelayer: &exelayer}, nil
}
