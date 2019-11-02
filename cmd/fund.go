package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wuxiaoxiaoshen/go-anything/src/fund"
)

var FundCmd = &cobra.Command{
	Use: "fund",
	Run: func(cmd *cobra.Command, args []string) {
		f := fund.NewOneFundAction("005063")
		f.SourcePage()
	},
}
