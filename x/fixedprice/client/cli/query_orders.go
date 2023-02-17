package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gauss/gauss/v6/x/fixedprice/types"
)

var _ = strconv.Itoa(0)

func CmdOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders [pool_address]",
		Short: "Query orders",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			poolAddress := ""
			if len(args) > 0 {
				poolAddress = args[0]
				_, err := sdk.AccAddressFromBech32(poolAddress)
				if err != nil {
					return err
				}
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryOrdersRequest{
				PoolAddress: poolAddress,
				Pagination:  pageReq,
			}

			res, err := queryClient.Orders(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
