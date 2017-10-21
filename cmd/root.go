package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var RootCmd = &cobra.Command{
	Use:   "papi",
	Short: "Papillon is a distribution static blog publish system.",
	Long: `A distribution static blog publish system based on ipfs in Go.
Complete documentation is unavailable yes ;(`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wecome to Papillon;)")
		fmt.Println("------------------------")
		cmd.Help()
		fmt.Println("------------------------")
		fmt.Println("2017 (c) GoGank Team")
	},
}