package Doodles

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/tidwall/gjson"
	"github.com/wuxiaoxiaoshen/go-anything/model"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
	"log"
	"strings"
)

type GoogleDoodlesAction struct {
	Url string
}

func NewGoogleDoodlesAction(url string) *GoogleDoodlesAction {
	return &GoogleDoodlesAction{
		Url: url,
	}
}
func (G GoogleDoodlesAction) Do() []model.GoogleDoodles {
	content := chromedp_helper.GetPageSourceHTTP(G.Url)
	log.Println(content)
	var results []model.GoogleDoodles

	js := gjson.Parse(content)
	for _, i := range js.Array() {
		var one model.GoogleDoodles
		one.ShareText = strings.TrimSpace(i.Get("share_text").String())
		one.Name = i.Get("name").String()
		one.Title = i.Get("title").String()
		one.Width = i.Get("high_res_width").String()
		one.Height = i.Get("high_res_height").String()
		one.Date = toDate(i.Get("run_date_array").Array())
		one.UrlHigh = fmt.Sprintf("https:" + i.Get("high_res_url").String())
		one.Url = fmt.Sprintf("https:" + i.Get("url").String())
		one.AlternateUrl = i.Get("alternate_url").String()
		log.Println(one)
		results = append(results, one)
	}
	return results
}

func getGoogleDoodlesSearch(c iris.Context) {
	gg := NewGoogleDoodlesAction(c.URLParam("search"))
	gg.Do()

}
func getGoogleDoodlesDate(c iris.Context) {
	gg := NewGoogleDoodlesAction(c.URLParam("date"))
	gg.Do()

}
func getGoogleDoodlesYear(c iris.Context) {
	gg := NewGoogleDoodlesAction(c.URLParam("year"))
	gg.Do()

}
