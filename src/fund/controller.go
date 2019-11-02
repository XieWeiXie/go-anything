package fund

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"

	"github.com/kataras/iris"
)

type OneFundAction struct {
	Code    string         `json:"code"`
	Result  OneFoundResult `json:"result"`
	urlHome string
	urlJS   string
}

func NewOneFundAction(code string) *OneFundAction {
	return &OneFundAction{
		Code: code,
	}
}

func (O *OneFundAction) formatUrl() {
	O.urlHome = fmt.Sprintf(HOME, O.Code)
	O.urlJS = fmt.Sprintf(JS, O.Code)
}

func (O *OneFundAction) SourcePage() {
	source := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithLog(), O.urlHome)
	fmt.Println(source)
}

func getFundInfoHandler(c iris.Context) {
	c.JSON(iris.Map{
		"data": "data",
	})
}
