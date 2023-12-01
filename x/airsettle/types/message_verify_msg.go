package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVerifyMsg = "verify_msg"

var _ sdk.Msg = &MsgVerifyMsg{}

func NewMsgVerifyMsg(creator string, id string, batchNumber uint64, inputs string) *MsgVerifyMsg {
	return &MsgVerifyMsg{
		Creator:     creator,
		Id:          id,
		BatchNumber: batchNumber,
		Inputs:      inputs,
	}
}

func (msg *MsgVerifyMsg) Route() string {
	return RouterKey
}

func (msg *MsgVerifyMsg) Type() string {
	return TypeMsgVerifyMsg
}

func (msg *MsgVerifyMsg) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVerifyMsg) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVerifyMsg) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
