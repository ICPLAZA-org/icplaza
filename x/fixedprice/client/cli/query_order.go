package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gauss/gauss/v6/x/fixedprice/types"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order [tokenId]",
		Short: "Query order",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenId := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryOrderRequest{
				TokenId: tokenId,
			}
			res, err := queryClient.Order(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
