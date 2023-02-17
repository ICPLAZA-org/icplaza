package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gauss/gauss/v6/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token [token_id]",
		Short: "Query token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqTokenId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryNftokenRequest{
				TokenId: reqTokenId,
			}

			res, err := queryClient.Nftoken(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
