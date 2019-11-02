package chromedp_helper

import (
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
)

func GetContextWithBackground() context.Context {
	ctx, _ := chromedp.NewContext(context.Background())
	return ctx
}

func GetContextWithLog() context.Context {
	ctx, _ := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	return ctx
}

// 只通过 Chromedp 获取到网页源代码
func GetPageSource(ctx context.Context, url string) string {
	var response string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("body", &response),
	})
	if err != nil {
		log.Println(fmt.Sprintf("GetPageSource: %s", err.Error()))
		return ""
	}
	return response
}
