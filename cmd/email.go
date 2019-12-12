package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/email"

	"github.com/spf13/cobra"
)

var fundURL = "http://push2his.eastmoney.com/api/qt/stock/kline/get?secid=1.000001&fields1=f1,f2,f3,f4,f5&fields2=f51,f52,f53,f54,f55,f56,f57,f58&klt=101&fqt=0&"
var bingURL = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=1"
var EmailCmd = &cobra.Command{
	Use:   "email",
	Short: "send info by email",
	Long:  "send info by email",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Step 1: Email init ...")
		email_operator.EmailInit()
		now := time.Now()
		fundURL = fmt.Sprintf(fundURL+"beg=%s&end=%s", "20190101", now.Format("20060102"))
		log.Println("Step 2: Fund url init ...", fundURL)
		bingURL = fmt.Sprintf(bingURL, 0)
		log.Println("Step 3: Bing url init ...", bingURL)
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start ...")
		send()
		//fmt.Println(bing())
		//fmt.Println(fund(), bing())
		log.Println("End ...")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("End 1: Email end ...")
		email_operator.DefaultEmailAction.Close()
	},
}

type (
	fundData struct {
		Code  string  `json:"code"`
		Name  string  `json:"name"`
		Low   float64 `json:"low"`
		High  float64 `json:"high"`
		Open  float64 `json:"open"`
		Close float64 `json:"close"`
		Rate  string  `json:"rate"`
		Date  string  `json:"date"`
	}
)

func fund() fundData {
	// 14:45
	data := chromedp_helper.GetPageSourceHTTP(fundURL)
	parse := gjson.Parse(data)
	d := parse.Get("data")
	var result fundData
	length := len(d.Get("klines").Array())
	if length <= 0 {
		return fundData{}
	}
	before := strings.Split(d.Get("klines").Array()[length-2].String(), ",")
	lines := d.Get("klines").Array()[length-1]
	list := strings.Split(lines.String(), ",")
	toFloat64 := func(v string) float64 {
		r, _ := strconv.ParseFloat(v, 32)
		return r
	}
	yesterday := toFloat64(before[2])
	rate := func() string {
		return strconv.FormatFloat((toFloat64(list[2])-yesterday)*100/yesterday, 'f', 2, 32) + "%"
	}
	result = fundData{
		Code:  d.Get("code").String(),
		Name:  d.Get("name").String(),
		Low:   toFloat64(list[4]),
		High:  toFloat64(list[3]),
		Open:  toFloat64(list[1]),
		Close: toFloat64(list[2]),
		Rate:  rate(),
		Date:  list[0],
	}
	return result

}
func bing() string {
	// 9:00
	log.Println("bingURL", bingURL)
	data := chromedp_helper.GetPageSourceHTTP(bingURL)
	list := gjson.Parse(data).Get("images").Array()
	if len(list) <= 0 {
		return "-1"
	}
	images := list[0]
	log.Println("images", images)
	url := fmt.Sprintf("%s%s", "https://cn.bing.com", strings.TrimSpace(images.Get("url").String()))
	return url
}

func send() {
	var content struct {
		fundData `json:"fund"`
		Url      string `json:"url"`
	}
	content.fundData = fund()
	content.Url = bing()
	t := template.New("index.html")
	tem, _ := t.Parse(templateForEmail())
	var byt bytes.Buffer
	e := tem.Execute(&byt, content)
	if e != nil {
		log.Println("send email", e.Error())
		return
	}
	fmt.Println(byt.String())
	email_operator.DefaultEmailAction.AddContent(byt.String())
	email_operator.DefaultEmailAction.Run("今日上证指数行情 || 今日壁纸")

}

func templateForEmail() string {
	return `
<html>
 <h1>今日上证指数行情 ||  今日壁纸</h1>
 <br>
 <br>
 <body>
	<table border="1">
	  <tr>
		<th>日期</th>
		<th>代码</th>
		<th>名称</th>
		<th>开市</th>
		<th>闭市</th>
		<th>最高</th>
		<th>最低</th>
		<th>幅度</th>
	  </tr>
	  <tr>
		<td>{{.Date}}</td>
		<td>{{.Code}}</td>
		<td>{{.Name}}</td>
		<td>{{.Open}}</td>
		<td>{{.Close}}</td>
		<td>{{.High}}</td>
		<td>{{.Low}}</td>
		<td>{{.Rate}}</td>
	  </tr>
	</table>
 <br>
 <br>
 <br>
 <img src="{{.Url}}" alt="image">
 </body>
</html>
`
}
