package keeper

import (
	"context"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerificationKey(goCtx context.Context, req *types.QueryVerificationKeyRequest) (*types.QueryVerificationKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryVerificationKeyResponse{}, nil
}
