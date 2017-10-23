package cmd

import (
	"fmt"
	"github.com/gogank/papillon/handler"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(generateCmd)

}

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Genrate a static blog website.",
	Long:  `Genrate the whole static blog website.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("Errors:unnecessary args in cmd!")
			fmt.Println("Example:")
			fmt.Println("papi gen")
			return
		}
		fmt.Println("  ____                           _   _")
		fmt.Println(" / ___| ___ _ __   ___ _ __ __ _| |_(_)_ __   __ _")
		fmt.Println("| |  _ / _ \\ '_ \\ / _ \\ '__/ _` | __| | '_ \\ / _` |")
		fmt.Println("| |_| |  __/ | | |  __/ | | (_| | |_| | | | | (_| |")
		fmt.Println(" \\____|\\___|_| |_|\\___|_|  \\__,_|\\__|_|_| |_|\\__, |")
		fmt.Println("                                              |___/")
		if err := handler.Generate("./config.toml"); err != nil {
			panic(err.Error())
		}
	},
}
