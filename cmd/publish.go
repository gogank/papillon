package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(publishCmd)

}

var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a static blog website to ipfs.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2{
			fmt.Println("Errors:")
			return
		}
	},
}