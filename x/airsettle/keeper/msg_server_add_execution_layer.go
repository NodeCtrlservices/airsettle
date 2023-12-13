package keeper

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) AddExecutionLayer(goCtx context.Context, msg *types.MsgAddExecutionLayer) (*types.MsgAddExecutionLayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newUUID := uuid.New().String()

	var exelayer = types.Exelayer{
		Validator:            []string{msg.Creator},
		VotingPower:          []uint64{100},
		LatestBatch:          0,
		LatestMerkleRootHash: "0",
		VerificationKey:      msg.VerificationKey,
		ChainInfo:            msg.ChainInfo,
		Id:                   newUUID,
		Creator:              msg.Creator,
	}

	// save execution layer data.
	Error := k.AddExecutionLayerHelper(
		ctx,
		exelayer,
		msg.Creator,
	)

	if Error != nil {
		return nil, Error
	}

	return &types.MsgAddExecutionLayerResponse{}, nil
}
