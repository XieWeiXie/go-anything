package cmd

import (
	"go-anything/router"
	"log"

	"github.com/spf13/cobra"
)

var ROOT = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Web Start...")

	},
	Run: func(cmd *cobra.Command, args []string) {
		router.Run("8888")
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	ROOT.AddCommand(STATION)
	ROOT.AddCommand(WX)
}
