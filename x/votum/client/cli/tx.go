package cli

import (
	"fmt"

	"github.com/EG-easy/votumchain/x/votum/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	votumTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "votumchain transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	votumTxCmd.AddCommand(client.PostCommands(
		GetCmdIssueToken(cdc),
	)...)

	return votumTxCmd
}

func GetCmdIssueToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "issue [coin]",
		Short: "issue coin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			coin, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}
			fmt.Println("OK")

			fmt.Printf("address:%s", cliCtx.GetFromAddress().String())
			fmt.Println("OK2")

			msg := types.NewMsgIssueToken(cliCtx.GetFromAddress(), coin)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			fmt.Println("OK3")
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
