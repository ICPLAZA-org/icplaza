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

func CmdUndelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "undelegate [delegator_address] [pool_address] [amount]",
		Short: "undelegate ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDelegatorAddress := args[0]
			argsPoolAddress := args[1]
			argsAmount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(argsAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgUndelegate(clientCtx.GetFromAddress().String(), argsDelegatorAddress, argsPoolAddress, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
