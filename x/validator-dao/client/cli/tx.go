package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gauss/gauss/v6/x/validator-dao/types"
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
	cmd.AddCommand(CmdReqAuthorization())
	cmd.AddCommand(CmdAddAuthBiz())
	cmd.AddCommand(CmdUpdateAuthBiz())
	cmd.AddCommand(CmdRemoveAuthBiz())
	// this line is used by starport scaffolding # 1

	return cmd 
}

func CmdReqAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "req-authorization [authorizer-address] [biz-name] [fee]",
		Short: "Broadcast message req-authorization",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			authorizerAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			fee, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgReqAuthorization(
				clientCtx.GetFromAddress(),
				authorizerAddr,
				args[1],
				fee,
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

func CmdAddAuthBiz() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-auth-biz [biz-name] [fee]",
		Short: "Broadcast message addAuthBiz",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fee, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddAuthBiz(
				clientCtx.GetFromAddress(),
				args[0],
				fee,
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

func CmdUpdateAuthBiz() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-auth-biz [biz-name] [fee]",
		Short: "Broadcast message updateAuthBiz",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fee, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAuthBiz(
				clientCtx.GetFromAddress(),
				args[0],
				fee,
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

func CmdRemoveAuthBiz() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-auth-biz [biz-name]",
		Short: "Broadcast message removeAuthBiz",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveAuthBiz(
				clientCtx.GetFromAddress(),
				args[0],
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
