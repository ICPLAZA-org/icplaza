package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/auction/types"
	"github.com/spf13/cobra"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCreateOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-order [token-id] [pool-address] [start-price] [start-time] [end-time] [min-step] [auto-agree-period]",
		Short: "create order",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenId := args[0]
			poolAddress := args[1]
			startPrice, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			startTimeInt64, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}
			startTime := time.Unix(startTimeInt64, 0)

			endTimeInt64, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}
			endTime := time.Unix(endTimeInt64, 0)

			minStep, err := sdk.ParseCoinNormalized(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			autoAgreePeriod, err := time.ParseDuration(args[6])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOrder(clientCtx.GetFromAddress().String(), tokenId, poolAddress, startPrice, minStep, startTime, endTime, autoAgreePeriod)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
