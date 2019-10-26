package weixin

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/kataras/iris"
)

type TagsAction struct {
	Url     string `json:"url"`
	Results []TagsResponse
}

func eachTagsResponse(selection *goquery.Selection) TagsResponse {
	var each TagsResponse
	attrs, ok := selection.Attr("id")
	if ok {
		each.Url = fmt.Sprintf("%s/pcindex/pc/%s", HOST, attrs)
	}
	each.Topic = strings.TrimSpace(selection.Text())
	return each
}

func (t *TagsAction) Do() {
	source := pageSource(t.Url)
	//fmt.Println(source)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(source))
	tags := doc.Find("#type_tab")
	tags.Find(`div[@class="fieed-box"] a`).Each(func(i int, selection *goquery.Selection) {
		t.Results = append(t.Results, eachTagsResponse(selection))
	})
	tags.Find(`#hide_tab a`).Each(func(i int, selection *goquery.Selection) {
		t.Results = append(t.Results, eachTagsResponse(selection))
	})
	//fmt.Println(t)

}

func (t *TagsAction) String() string {
	r, _ := json.MarshalIndent(t.Results, " ", " ")
	return fmt.Sprintf("Tags: 主页标签项: %s", string(r))
}

func tagsHandler(c iris.Context) {
	var action TagsAction
	action.Url = "https://weixin.sogou.com/"
	action.Do()
	c.JSON(iris.Map{
		"data": action.Results,
	})
}

type BannerAction struct {
	Url     string `json:"url"`
	Results []Banner
}

func (b *BannerAction) Do() {
	source := pageSource(b.Url)
	doc, e := goquery.NewDocumentFromReader(strings.NewReader(source))
	if e != nil {
		log.Println(fmt.Sprintf("Banner error : %s", e.Error()))
		return
	}
	doc.Find()
}
func (b *BannerAction) String() string {
	r, _ := json.MarshalIndent(b.Results, " ", " ")
	return fmt.Sprintf("Banner: 主页横幅项: %s", string(r))
}

func bannerHandler(c iris.Context)   {}
func hotTopicHandler(c iris.Context) {}
func passagesHandler(c iris.Context) {}
