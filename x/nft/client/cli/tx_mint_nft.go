package cli

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gauss/gauss/v6/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdMintNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [recipient] [token_id] [cate_id] [token_uri] [company_uri] [value_added_tax] [copyright_tax] [name] [components]",
		Short: `Broadcast message createNft,components is json string format,example:[{"class_id":"collection",min_amount:1,max_amount:1,type:1}]`,
		Args:  cobra.MinimumNArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecipient := args[0]
			argsTokenId := args[1]
			argsCatId := args[2]
			argsTokenUri := args[3]
			argsCompanyUri := args[4]
			argsValueAddedTax, err := sdk.NewDecFromStr(args[5])
			if err != nil {
				return err
			}
			argsCopyrightTax, err := sdk.NewDecFromStr(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}
			name := args[7]
			var components []*types.Component
			if len(args) > 8 {
				componentsStr := args[8]
				if componentsStr != "" {
					err := json.Unmarshal([]byte(componentsStr), &components)
					if err != nil {
						return err
					}
				}
			}
			msg := types.NewMsgMintNft(
				clientCtx.GetFromAddress().String(),
				argsRecipient,
				argsTokenId,
				argsCatId,
				argsTokenUri,
				argsCompanyUri,
				argsValueAddedTax,
				argsCopyrightTax,
				name,
				components,
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
