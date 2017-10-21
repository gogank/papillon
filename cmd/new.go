package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(newCmd)

}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "New file in a static blog website.",
	Long:  `New a markdown file in local static blog website.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1{
			fmt.Println("Errors:Please specific the markdown file path!")
			fmt.Println("Example:")
			fmt.Println("papi new ./test.md")
			return
		}
		//TODO specific logic
		fmt.Println("Call NewMD cmd!")
	},
}