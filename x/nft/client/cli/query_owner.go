package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/nft/types"
	"github.com/spf13/cobra"
	"strconv"
)

var _ = strconv.Itoa(0)

func CmdOwners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner [accountAddress] [cateId]",
		Short: "get the NFTs owned by an account address",
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

			params := &types.QueryOwnerRequest{
				CateId: cateId,
				Owner:  args[0],
			}
			res, err := queryClient.Owner(cmd.Context(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
