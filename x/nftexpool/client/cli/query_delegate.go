package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

var _ = strconv.Itoa(0)

func CmdQuerydelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegation [delegator_address] [pool_address]",
		Short: "Query my delegate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqdelegatorAddress := args[0]
			reqPoolAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDelegateRequest{
				Delegator:   reqdelegatorAddress,
				PoolAddress: reqPoolAddress,
			}

			res, err := queryClient.QDelegate(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
