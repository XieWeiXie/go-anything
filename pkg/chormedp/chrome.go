package chormedp_helper

import (
	"log"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
)

func GetContextWithBackground() context.Context {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	return ctx
}

func GetContextWithLog() context.Context {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	return ctx
}

func GetPageSource(ctx context.Context, url string) string {
	var response string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("body", &response),
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	return response
}
