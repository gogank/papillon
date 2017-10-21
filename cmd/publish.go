package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(publishCmd)

}

var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a static blog website to ipfs.",
	Long:  `Generate a private key pem , and save to a specific file path`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}