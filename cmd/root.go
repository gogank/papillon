package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//RootCmd is the root cobra cmd
var RootCmd = &cobra.Command{
	Use:   "papi",
	Short: "Papillon is a distribution static blog publish system.",
	Long: `A distribution static blog publish system based on ipfs in Go.
Complete documentation is unavailable yes ;`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" ____             _ _ _")
		fmt.Println("|  _ \\ __ _ _ __ (_) | | ___  _ __")
		fmt.Println("| |_) / _` | '_ \\| | | |/ _ \\| '_ \\")
		fmt.Println("|  __/ (_| | |_) | | | | (_) | | | |")
		fmt.Println("|_|   \\__,_| .__/|_|_|_|\\___/|_| |_|")
		fmt.Println("|_|")
		fmt.Println("Welcome to Papillon;")
		fmt.Println("------------------------")
		cmd.Help()
		fmt.Println("------------------------")
		fmt.Println("2017 (c) GoGank Team")
	},
}
