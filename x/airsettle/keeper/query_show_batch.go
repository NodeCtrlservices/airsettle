package keeper

import (
	"context"

	"airsettle/x/airsettle/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowBatch(goCtx context.Context, req *types.QueryShowBatchRequest) (*types.QueryShowBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := k.GetbatchById(ctx, req.Id, req.BatchNumber)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	batch_data := types.Batch{
		BatchNumber:        batch.BatchNumber,
		MerkleRootHash:     batch.MerkleRootHash,
		PrevMerkleRootHash: batch.PrevMerkleRootHash,
		ZkProof:            batch.ZkProof,
	}

	return &types.QueryShowBatchResponse{Batch: batch_data}, nil
}

