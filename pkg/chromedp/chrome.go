package chromedp_helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

func GetPageSourceHTTP(url string, header ...map[string]string) string {
	client := http.DefaultClient
	request, e := http.NewRequest(http.MethodGet, url, nil)
	if len(header) != 0 {
		for _, v := range header {
			for key, value := range v {
				request.Header.Add(key, value)
			}
		}
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	if e != nil {
		log.Println(e)
		return ""
	}
	response, e := client.Do(request)
	if e != nil {
		log.Println(e)
		return ""
	}
	defer response.Body.Close()
	content, e := ioutil.ReadAll(response.Body)
	if e != nil {
		log.Println(e)
		return ""
	}
	return string(content)

}
