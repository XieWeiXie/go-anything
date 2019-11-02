package weixin

import (
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
)

func pageSource(url string) string {
	return chromedp_helper.GetPageSource(chromedp_helper.GetContextWithLog(), url)
}
