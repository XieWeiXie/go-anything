package cmd

import (
	"context"
	"log"

	"github.com/knq/chromedp"
	"github.com/spf13/cobra"
)

var WX = &cobra.Command{
	Use:   "wexin",
	Short: "get wexin command from https://weixin.sogou.com/",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		ChromeNav()
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func ChromeNav() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	var title string
	var location string
	e := chromedp.Run(ctx,
		chromedp.Navigate("https://weixin.sogou.com/"),
		chromedp.Title(&title),
		chromedp.Location(&location),
	)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println("Title: ", title)
	log.Println("Location: ", location)
}
