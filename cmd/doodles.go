package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wuxiaoxiaoshen/go-anything/src/Doodles"
)

var GoogleDoodles = &cobra.Command{
	Use: "doodles",
	Run: func(cmd *cobra.Command, args []string) {
		gg := Doodles.NewGoogleDoodlesAction("https://www.google.com/doodles/json/2020/2?hl=zh_CN")
		gg.Do()
	},
}
