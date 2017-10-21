package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(generateCmd)

}

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Genrate a static blog website.",
	Long:  `Generate a private key pem , and save to a specific file path`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}