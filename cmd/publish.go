package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/gogank/papillon/publish"
)

func init() {
	RootCmd.AddCommand(publishCmd)

}

var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a static blog website to ipfs.",
	Long:  `Publish a new static blog website to ipfs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0{
			fmt.Println("Errors:unnecessary args in cmd!")
			fmt.Println("Example:")
			fmt.Println("papi pub")
			return
		}
		//TODO  specific logic
		fmt.Println("Call Publish cmd!")
		pub := publish.NewPublishImpl()
		hash,err := pub.PublishCmd()
		if err != nil {
			fmt.Println("Errors:Publish Failed.")
			return
		}
		fmt.Println("The Url is https://ipfs.io/ipns/",hash)
	},
}