package keeper

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowExecutionLayerByAddress(goCtx context.Context, req *types.QueryShowExecutionLayerByAddressRequest) (*types.QueryShowExecutionLayerByAddressResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	address := req.Address
	chain_id, found := k.GetExelayerIdByAddress(ctx, address)
	if chain_id == "" || found == false { // any one condition is enough too
		// no execution layer on this address
		return nil, sdkerrors.ErrKeyNotFound
	}

	exelayer, found := k.GetExelayerById(ctx, chain_id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowExecutionLayerByAddressResponse{Exelayer: &exelayer}, nil

}