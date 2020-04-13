package Doodles

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/wuxiaoxiaoshen/go-anything/model"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/concurrency"
	"html/template"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDoodles(t *testing.T) {

	c := concurrency.NewConCurrency(30)
	var all []string
	for i := 1991; i < 2021; i++ {
		all = append(all, urlFormat(i)...)
	}
	for _, i := range all {
		year, _ := urlSplit(i)
		y, _ := strconv.Atoi(year)
		c.Add(1)

		go func(year int, i string) {
			defer c.Done()
			toYear(year, i)
		}(y, i)
	}
	c.Wait()

}

func toSave(title string, content []byte) error {
	replacer := strings.NewReplacer("，", "", "\"", "", " ", "",
		"/", "")
	title = replacer.Replace(title)
	f, e := os.Create(fmt.Sprintf("../../doc/doodles/%s.md", title))
	if e != nil {
		log.Println(e)
		return e
	}
	f.Write(content)
	defer f.Close()
	return nil
}

func toYear(year int, url string) {
	d := NewGoogleDoodlesAction(url)
	results := d.Do()
	var returnResult model.MonthGoogleDoodles
	_, returnResult.Month = urlSplit(url)
	returnResult.Doodles = results
	ts := template.New("doodles")
	box := packr.NewBox("../../doc/doodles/template")
	templateText, _ := box.FindString("doodles.md")
	ts = ts.Funcs(template.FuncMap{"toDateFunc": toDateFormat})
	tem, _ := ts.Parse(templateText)
	var byt bytes.Buffer
	e := tem.Execute(&byt, returnResult)
	if e != nil {
		log.Println(e)
		return
	}
	go func(title string, content []byte) {
		toSave(title, content)
	}(fmt.Sprintf("%d-%s", year, returnResult.Month), byt.Bytes())
}

func toDateFormat(date time.Time) string {
	y, m, d := date.Date()
	return fmt.Sprintf("%d年%d月%d日", y, m, d)
}

func urlFormat(year int) []string {
	var result []string
	for i := 12; i > 0; i-- {
		result = append(result, fmt.Sprintf("https://www.google.com/doodles/json/%d/%d?hl=zh_CN", year, i))
	}
	return result
}

func urlSplit(urls string) (string, string) {
	URL, _ := url.Parse(urls)
	fmt.Println(URL.Path)
	list := strings.Split(URL.Path, "/")
	return list[len(list)-2], list[len(list)-1]
}

func TestUrlSplit(t *testing.T) {
	url := "https://www.google.com/doodles/json/2020/4?hl=zh_CN"
	fmt.Println(urlSplit(url))
}
