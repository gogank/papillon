package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"certgen/certUtils"
	"github.com/gogank/papillon/utils"
)

func init() {
	RootCmd.AddCommand(publishCmd)

}

var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a static blog website to ipfs.",
	Long:  `Generate a private key pem , and save to a specific file path`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2{
			fmt.Println("Please specific the private key path and the public key path")
			fmt.Println("Example:")
			fmt.Println("certgen gc ./test.priv ./test.pub")
			return
		}
		fmt.Println("Generate a private key")
		privPath := utils.Abs(args[0])
		pubPath := utils.Abs(args[1])
		fmt.Println("save the private key into ",privPath)
		err := certUtils.GeneratePrivKeyFile(privPath,pubPath)
		if err != nil {
			fmt.Println("generate the priv file failed!")
		}
		fmt.Println("generate the priv file success")
	},
}