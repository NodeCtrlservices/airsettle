package cli

import (
	"strconv"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-batch [id] [batch-number] [merkle-root-hash] [prev-merkle-root-hash] [zk-proof]",
		Short: "Broadcast message add_batch",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argBatchNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argMerkleRootHash := args[2]
			argPrevMerkleRootHash := args[3]
			argZkProof := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddBatch(
				clientCtx.GetFromAddress().String(),
				argId,
				argBatchNumber,
				argMerkleRootHash,
				argPrevMerkleRootHash,
				argZkProof,
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
