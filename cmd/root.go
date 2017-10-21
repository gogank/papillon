package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "papi",
	Short: "Papillon is a distribution static blog publish system",
	Long: `long desc
                `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}