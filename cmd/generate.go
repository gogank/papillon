package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(generateCmd)

}

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Genrate a static blog website.",
	Long:  `Genrate the whole static blog website.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0{
			fmt.Println("Errors:unnecessary args in cmd!")
			fmt.Println("Example:")
			fmt.Println("papi gen")
			return
		}
		//TODO  specific logic
		fmt.Println("Call Generate cmd!")
	},
}