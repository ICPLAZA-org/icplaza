package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gauss/gauss/v6/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdTransferNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [from_address] [to_address] [nft_token_id]",
		Short: "transfer nft",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			sender := args[0]

			recipient := args[1]

			cateId := ""
			tokenId := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferNft(sender, recipient, cateId, tokenId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
