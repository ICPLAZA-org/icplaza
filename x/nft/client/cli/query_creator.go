package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "creator [creator-address] [cate-id]",
		Short: "Query nft by creator",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			cateId := ""
			if len(args) == 2 {
				cateId = args[1]
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCreatorRequest{
				CateId: cateId,
				Creator: args[0],
			}

			res, err := queryClient.Creator(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
