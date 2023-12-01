package cli

import (
	"strconv"

	"airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerifyMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-msg [id] [batch-number] [inputs]",
		Short: "Broadcast message verify_msg",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argBatchNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argInputs := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVerifyMsg(
				clientCtx.GetFromAddress().String(),
				argId,
				argBatchNumber,
				argInputs,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
