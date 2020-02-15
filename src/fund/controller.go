package fund

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/tidwall/gjson"

	"github.com/PuerkitoBio/goquery"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"

	"github.com/kataras/iris/v12"
)

type OneFundAction struct {
	Code    string         `json:"code"`
	Result  OneFoundResult `json:"result"`
	urlHome string
	urlJS   string
}

func NewOneFundAction(code string) *OneFundAction {
	o := &OneFundAction{
		Code: code,
	}
	o.formatUrl()
	return o
}

func (O *OneFundAction) formatUrl() {
	O.urlHome = fmt.Sprintf(HOME, O.Code)
	O.urlJS = fmt.Sprintf(JS, O.Code)
}

func (O *OneFundAction) resultPartHTML() {}
func (O *OneFundAction) resultPartJS() {
	source := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithLog(), O.urlJS)
	patternRate := `fund_Rate="(.*?)"`
	r1 := regexp.MustCompile(patternRate)
	rate := r1.FindAllStringSubmatch(source, -1)
	O.Result.BuyRate = rate[0][1]
	patternYear := `syl_1n="(.*?)"`
	patternSixMonth := `syl_6y="(.*?)"`
	patternThreeMonth := `syl_3y="(.*?)"`
	patterMonth := `syl_1y="(.*?)"`
	r2 := regexp.MustCompile(patternYear)
	year := r2.FindAllStringSubmatch(source, -1)
	O.Result.YearForEarn = year[0][1] + "%"
	r3 := regexp.MustCompile(patternSixMonth)
	sixMonth := r3.FindAllStringSubmatch(source, -1)
	O.Result.SixMonthForEarn = sixMonth[0][1] + "%"
	r4 := regexp.MustCompile(patternThreeMonth)
	threeMonth := r4.FindAllStringSubmatch(source, -1)
	O.Result.ThreeMonthForEarn = threeMonth[0][1] + "%"
	r5 := regexp.MustCompile(patterMonth)
	month := r5.FindAllStringSubmatch(source, -1)
	O.Result.MonthForEarn = month[0][1] + "%"
}
func (S *Stock) resultSockets(url string) {
	url = fmt.Sprintf(STOCK, url)
	source := chromedp_helper.GetPageSourceHTTP(url)
	pattern := `\((.*?)\)`
	r1 := regexp.MustCompile(pattern)
	jsonSource := r1.FindAllStringSubmatch(source, -1)[0][1]
	S.Price = gjson.Parse(jsonSource).Get("info.c").String()
}

func (O *OneFundAction) sourcePage() {
	source := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithLog(), O.urlHome)
	doc, e := goquery.NewDocumentFromReader(strings.NewReader(source))
	if e != nil {
		return
	}
	var result OneFoundResult
	result.Code = O.Code
	result.Name = doc.Find(".fundDetail-tit").Text()
	fundInfoItem := doc.Find(".fundInfoItem")
	price := doc.Find("#gz_gsz").Text()

	guessPrice, _ := strconv.ParseFloat(price, 2)
	result.GuessPrice = guessPrice
	guessRate := doc.Find("#gz_gszzl").Text()
	result.GuessRate = guessRate
	realPrice := fundInfoItem.Find("dl.dataItem02 dd.dataNums span")
	result.Price, _ = strconv.ParseFloat(realPrice.Eq(0).Text(), 2)
	result.Rate = realPrice.Eq(1).Text()

	trs := doc.Find(".infoOfFund table tbody tr")
	result.Type = trs.Eq(0).Find("td").Eq(0).Find("a").Text()
	result.Scale = trs.Eq(0).Find("td").Eq(1).Text()
	result.Date = trs.Eq(1).Find("td").Eq(0).Text()
	result.Company = trs.Eq(1).Find("td").Eq(1).Find("td a").Text()

	var stocks []Stock
	var wg sync.WaitGroup
	doc.Find("li#position_shares tbody tr").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			var tmpStock Stock
			tmpStock.Name = strings.TrimSpace(selection.Find("td").Eq(0).Text())
			tmpStock.Rate = selection.Find("td").Eq(1).Text()
			tmpStock.Ratio = selection.Find("td").Eq(2).Find("span").Text()
			code, _ := selection.Find("td").Eq(2).Attr("stockcode")
			tmpStock.Code = strings.Split(code, "_")[1]
			var url string
			if strings.HasPrefix(tmpStock.Code, "0") {
				url = tmpStock.Code + "2"
			} else {
				url = tmpStock.Code + "1"
			}
			wg.Add(1)
			go func(url string, s *Stock) {
				defer wg.Done()
				s.resultSockets(url)
				stocks = append(stocks, tmpStock)
			}(url, &tmpStock)

		}
	})
	wg.Wait()
	result.Stocks = stocks
	O.Result = result
}
func (O *OneFundAction) Do() {
	O.sourcePage()
	O.resultPartJS()

}

func (O *OneFundAction) String() string {
	b, _ := json.MarshalIndent(O.Result, " ", " ")
	return string(b)
}
func getFundInfoHandler(c iris.Context) {
	code := c.Params().GetString("code")
	fund := NewOneFundAction(code)
	fund.Do()
	c.JSON(iris.Map{
		"data": fund.Result,
	})
}

type GlobalAction struct {
	Result GlobalFundMarkets `json:"result"`
}

func NewGlobalAction() *GlobalAction {
	return &GlobalAction{}
}

func (G *GlobalAction) sourcePage() {
	source := chromedp_helper.GetPageSourceHTTP(GLOBAL)
	for i := 0; i < 21; i++ {
		data := gjson.Parse(source).Get(fmt.Sprintf("data.diff.%d", i))
		var tempData GlobalFundMarket
		tempData.Code = data.Get("f12").String()
		tempData.Name = data.Get("f14").String()
		tempData.Rate = strconv.FormatFloat(data.Get("f3").Float()/100, 'f', 2, 32) + "%"
		tempData.Current = strconv.FormatFloat(data.Get("f2").Float()/100, 'f', 2, 32)
		G.Result = append(G.Result, tempData)
	}
}

func (G *GlobalAction) Do() {
	G.sourcePage()
}

func (G *GlobalAction) String() string {
	b, _ := json.MarshalIndent(G.Result, " ", " ")
	return string(b)
}

func getGlobalInfoHandler(c iris.Context) {
	g := NewGlobalAction()
	g.Do()
	c.JSON(iris.Map{
		"data": g.Result,
	})
}
