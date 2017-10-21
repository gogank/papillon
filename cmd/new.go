package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(newCmd)

}

var newCmd = &cobra.Command{
	Use:   "gen",
	Short: "New file in a static blog website.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1{
			fmt.Println("Errors:")
			return
		}

	},
}