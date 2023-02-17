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

func CmdUpdatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [pool-address] [commission-rate] [commission-address] [valueAddedTax-address]",
		Short: "update pool",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			poolAddress := args[0]
			commissionRateX := args[1]
			commissionRate, err := sdk.NewDecFromStr(commissionRateX)
			if err != nil {
				return err
			}
			commissionAddress := args[2]
			valueAddedTaxAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePool(clientCtx.GetFromAddress().String(), poolAddress, commissionRate, commissionAddress, valueAddedTaxAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
