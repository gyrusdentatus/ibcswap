package cli

import (
	"fmt"
	// "strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ibcswap/ibcswap/v6/modules/apps/101-interchain-swap/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group interchainswap queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListInterchainLiquidityPool())
	cmd.AddCommand(CmdShowInterchainLiquidityPool())
	cmd.AddCommand(CmdListInterchainMarketMaker())
	cmd.AddCommand(CmdShowInterchainMarketMaker())
	// this line is used by starport scaffolding # 1

	return cmd
}
