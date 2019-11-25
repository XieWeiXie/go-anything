package cmd

import (
	"log"

	"github.com/wuxiaoxiaoshen/go-anything/configs"

	"github.com/spf13/cobra"
)

var ConfigsCmd = &cobra.Command{
	Use:     "configs",
	Short:   "Get configs",
	Long:    "Get All Configs by Command",
	Aliases: []string{"c", "-c", "C", "-C"},
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Run ./go-anything configs arg")
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(configs.DefaultConfigs.LoadConfigs(args[0]))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("Say Bye!")
	},
}
