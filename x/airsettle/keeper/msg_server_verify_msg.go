package keeper

import (
	"context"

	"airsettle/x/airsettle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) VerifyMsg(goCtx context.Context, msg *types.MsgVerifyMsg) (*types.MsgVerifyMsgResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid msguest")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get verification key
	str_verificationKey := k.LocalGetvkey(ctx, msg.Id)
	if str_verificationKey == "" {
		Log("verification key not found")
		return &types.MsgVerifyMsgResponse{
			Result:  false,
			Message: "verification key not found",
		}, sdkerrors.ErrKeyNotFound
	}

	batch, found := k.GetbatchById(ctx, msg.Id, msg.BatchNumber)
	str_zkproof := batch.ZkProof
	if !found {
		Log("Proof not found")
		return &types.MsgVerifyMsgResponse{
			Result:  false,
			Message: "proof not found",
		}, sdkerrors.ErrKeyNotFound
	}


	// verify
	result, message := k.Verifier(ctx, str_zkproof, str_verificationKey, msg.Inputs)
	Log(message)

	return &types.MsgVerifyMsgResponse{
		Result:  result,
		Message: message,
	}, nil
}
