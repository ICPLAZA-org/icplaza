package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gauss/gauss/v6/x/blindbox/types"
)


// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePool())
	cmd.AddCommand(CmdUpdatePool())
	cmd.AddCommand(CmdRemovePool())
	cmd.AddCommand(CmdCreateBox())
	cmd.AddCommand(CmdRevokeBoxGroup())
	cmd.AddCommand(CmdOpenBox())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdCreateBox() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-box [pool-id] [group-id] [box-size] [box-price] [start-time] [end-time] (--random-min [min] --random-max [max] --random-nfts [nfts]) (--fixed-nfts [nfts]) [flags]",
		Short: "create boxes",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			poolId := args[0]

			argGroupId := args[1]
			argBoxSize, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return
			}
			argBoxPrice, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return
			}

			argStartTime, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return
			}
			startTime := time.Unix(argStartTime, 0)

			argEndTime, err := strconv.ParseInt(args[5], 10, 64)
			if err != nil {
				return
			}
			endTime := time.Unix(argEndTime, 0)

			argRandomMin, err := cmd.Flags().GetUint64(FlagRandomMin)
			if err != nil {
				return
			}
			argRandomMax, err := cmd.Flags().GetUint64(FlagRandomMax)
			if err != nil {
				return
			}

			argRandomNftsStr, err := cmd.Flags().GetString(FlagRandomNfts)
			if err != nil {
				return err
			}
			var argRandomNfts []string
			if argRandomNftsStr != "" {
				argRandomNfts = strings.Split(argRandomNftsStr, ",")
			}
			var argFixedNfts []string
			argFixedNftsStr, err := cmd.Flags().GetString(FlagFixedNfts)
			if err != nil {
				return err
			}
			if argFixedNftsStr != "" {
				argFixedNfts = strings.Split(argFixedNftsStr, ",")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBox(
				clientCtx.GetFromAddress().String(),
				argGroupId,
				argBoxSize,
				argBoxPrice,
				startTime,
				endTime,
				argRandomMin,
				argRandomMax,
				argRandomNfts,
				argFixedNfts,
				poolId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().Uint64(FlagRandomMin, 0, "min count of nft per box")
	cmd.Flags().Uint64(FlagRandomMax, 0, "max count of nft per box")
	cmd.Flags().String(FlagRandomNfts, "", "random nft list")
	cmd.Flags().String(FlagFixedNfts, "", "fixed nft list")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdOpenBox() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-box [group-id] [box-id]",
		Short: "Broadcast message openBox",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGroupId := args[0]
			argBoxId := args[1]
			argBoxIds := strings.Split(argBoxId, ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOpenBox(
				clientCtx.GetFromAddress().String(),
				argGroupId,
				argBoxIds,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [fee-rate] [fee-address] [pool-id]",
		Short: "Broadcast message createPool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFee, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return
			}
			_,err = sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return
			}
			argFeeAddress := args[1]

			argPoolId := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(
				clientCtx.GetFromAddress().String(),
				argFee,
				argFeeAddress,
				argPoolId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pool [fee] [fee-address] [pool-id]",
		Short: "Broadcast message updatePool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFee ,err:= sdk.NewDecFromStr(args[0])
			if err != nil {
				return
			}
			argFeeAddress := args[1]
			argPoolId := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePool(
				clientCtx.GetFromAddress().String(),
				argFee,
				argFeeAddress,
				argPoolId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemovePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-pool [pool-id]",
		Short: "Broadcast message removePool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPoolId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemovePool(
				clientCtx.GetFromAddress().String(),
				argPoolId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRevokeBoxGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-box-group [group-id]",
		Short: "Broadcast message revokeBoxGroup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGroupId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeBoxGroup(
				clientCtx.GetFromAddress().String(),
				argGroupId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
