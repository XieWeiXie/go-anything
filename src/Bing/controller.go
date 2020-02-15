package Bing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/page_source"

	"github.com/kataras/iris/v12"
	"github.com/tidwall/gjson"
)

type ImageAction struct {
	Url     string
	Results []ResponseForImages
}

func (I *ImageAction) Do() {
	rr, _ := http.NewRequest(http.MethodGet, I.Url, nil)
	r := page_source.Request{R: rr}
	source, _ := page_source.GetPageSource(nil, r)
	doc := gjson.ParseBytes([]byte(source))
	for _, i := range doc.Get("images").Array() {
		var eachImages ResponseForImages
		eachImages.Url = fmt.Sprintf("%s%s", BASE, strings.TrimSpace(i.Get("url").String()))
		eachImages.Date = i.Get("enddate").String()
		eachImages.Copyright = i.Get("copyright").String()
		I.Results = append(I.Results, eachImages)
	}
}

func (I *ImageAction) String() string {
	r, _ := json.MarshalIndent(I.Results, " ", "")
	return fmt.Sprintf("BingImages: 每日壁纸: %s", string(r))
}
func imageHandler(c iris.Context) {
	date, _ := c.Params().GetInt("date")
	if date < 0 || date > 7 {
		return
	}
	var I ImageAction
	I.Url = fmt.Sprintf(HOST, date)
	fmt.Println(I.Url)
	I.Do()
	c.JSON(iris.Map{
		"data": I.Results,
	})
}
