package konfig

import (
	"fmt"
	"github.com/spf13/cobra"
	"konfig/pkg/konfig"
	"os"
)

var endpoint string
var manifestId int
var file string
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "generate local env",
	Long:    "g -e=endpoint -i 123 -f local.env",
	//Args:    cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, _ := cmd.Flags().GetString("endpoint")
		manifestId, _ := cmd.Flags().GetInt("manifestId")
		file, _ := cmd.Flags().GetString("file")
		if manifestId <= 0 {
			fmt.Println("provide a valid manifest id")
			os.Exit(0)
		}

		if endpoint == "" {
			fmt.Println("provide a valid endpoint")
			os.Exit(0)
		}

		if file == "" {
			file = "local.sample.env"
		}

		konfig.Generate(endpoint, manifestId, file)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "endpoint")
	generateCmd.Flags().IntVarP(&manifestId, "manifestId", "i", 0, "manifest id")
	generateCmd.Flags().StringVarP(&file, "file", "f", "", "path")
	rootCmd.AddCommand(generateCmd)
}
