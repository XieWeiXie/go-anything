package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wuxiaoxiaoshen/go-anything/src/weixin"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

var WX = &cobra.Command{
	Use:   "weixin",
	Short: "get wexin command from https://weixin.sogou.com/",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		//ctx, cancel := chromedp.NewContext(
		//	context.Background(),
		//	chromedp.WithLogf(log.Printf),
		//)
		//defer cancel()
		//ChromeNav(ctx)
		//Response("https://weixin.sogou.com/")
		//ChromeResponse(ctx)
		var t weixin.TagsAction
		t.Url = "https://weixin.sogou.com/"
		t.Do()
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func ChromeNav(ctx context.Context) {
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

func ChromeResponse(ctx context.Context) {
	var response string
	var text string
	e := chromedp.Run(ctx,
		chromedp.Navigate("https://weixin.sogou.com/"),
		chromedp.OuterHTML("body", &response),
		chromedp.Text(".fieed-box", &text),
	)
	if e != nil {
		log.Println(e)
		return
	}
	if response == "" {
		log.Println("Response: ", nil)
		return
	}
	log.Println("Response: ", &response)
	log.Println("Text: ", text)
}

func Response(url string) string {
	response, e := http.Get(url)
	if e != nil {
		log.Println(e)
		return ""
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	return ""
}
