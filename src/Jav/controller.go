package Jav

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"

	"github.com/PuerkitoBio/goquery"

	"github.com/kataras/iris/v12"
)

type CodeDetailAction struct {
	Code   string      `json:"code"`
	Result Designation `json:"result"`
}

func (C *CodeDetailAction) format() string {
	return toRawString(HOST) + fmt.Sprintf("%s/", C.Code)
}

func (C *CodeDetailAction) Do() {
	source := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithLog(), C.format())
	doc, e := goquery.NewDocumentFromReader(strings.NewReader(source))
	if e != nil {
		log.Println(fmt.Sprintf("CodeDetail: %s", e.Error()))
		return
	}
	var each Designation
	img := doc.Find("div.col-md-9.img img")
	src := img.AttrOr("src", "No Image Url")
	each.Image = src
	each.Title = img.AttrOr("alt", "No Title")
	each.Url = C.format()
	each.Code = C.Code
	doc.Find("ul.list-group").Each(func(i int, selection *goquery.Selection) {
		each.Date = selection.Find("li").Eq(1).Text()
		each.Duration = selection.Find("li").Eq(2).Text()
		each.Actor = selection.Find("li").Last().Text()
	})
	C.Result = each

}

func (C *CodeDetailAction) String() string {
	r, _ := json.MarshalIndent(C.Result, " ", " ")
	return fmt.Sprintf("Designation: 详情: %s", string(r))
}

func codeDetailHandler(c iris.Context) {
	var r CodeDetailAction
	r.Code = strings.ToUpper(c.Params().Get("code"))
	r.Do()
	c.JSON(iris.Map{
		"data": r.Result,
	})

}
