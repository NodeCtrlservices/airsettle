package cli

import (
	"strconv"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerify() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify [id] [batch-number] [inputs]",
		Short: "Query verify",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]
			reqBatchNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			reqInputs := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVerifyRequest{

				Id:          reqId,
				BatchNumber: reqBatchNumber,
				Inputs:      reqInputs,
			}

			res, err := queryClient.Verify(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
