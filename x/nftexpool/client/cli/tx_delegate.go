package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gauss/gauss/v6/x/nftexpool/types"
)

var _ = strconv.Itoa(0)

func CmdDelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate [pool_address] [amount]",
		Short: "delegate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPoolAddress := args[0]
			argsAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(argsAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgDelegate(clientCtx.GetFromAddress().String(), argsPoolAddress, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
