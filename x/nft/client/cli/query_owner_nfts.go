package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gauss/gauss/v6/x/nft/types"
	"github.com/spf13/cobra"
	"strconv"
)

var _ = strconv.Itoa(0)

func CmdOwnerNFTs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft [name]",
		Short: "get the NFTs by Name",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryNftsByNameOrTokenRequest{
				Owner:      clientCtx.GetFromAddress().String(),
				Name:       name,
				TokenId:    "",
				Pagination: pageReq,
			}
			res, err := queryClient.OwnerNfts(cmd.Context(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
