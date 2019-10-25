package weixin

import "github.com/wuxiaoxiaoshen/go-anything/pkg/chormedp"

func pageSource(url string) string {
	return chormedp_helper.GetPageSource(chormedp_helper.GetContextWithLog(), url)
}
