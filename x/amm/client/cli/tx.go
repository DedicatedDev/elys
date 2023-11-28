package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/elys-network/elys/x/amm/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	FlagDiscount               = "discount"
	FlagRecipient              = "recipient"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
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
	cmd.AddCommand(CmdJoinPool())
	cmd.AddCommand(CmdExitPool())
	cmd.AddCommand(CmdSwapExactAmountIn())
	cmd.AddCommand(CmdSwapExactAmountOut())
	cmd.AddCommand(CmdSwapByDenom())
	cmd.AddCommand(CmdUpdatePoolParams())
	// this line is used by starport scaffolding # 1

	return cmd
}
