package konfig

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "konfig",
	Version: version,
	Short:   "konfig - a simple CLI to build local config",
	Long:    `konfig is a super fancy CLI :P`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "!! There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
