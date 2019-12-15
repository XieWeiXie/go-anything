package holiday

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
	"log"
	"net/url"
	"strings"
	"sync"
)

func pageSource(url string) string {
	content := chromedp_helper.GetPageSourceHTTP(url, map[string]string{
		"Host": "sousuo.gov.cn",
	})
	return content
}

func urlEncode(v map[string]string) string {
	r := url.Values{}
	for key, value := range v {
		r.Add(key, value)
	}
	return r.Encode()
}

func getContent(content string) string {
	reader := bytes.NewReader([]byte(content))
	doc, e := goquery.NewDocumentFromReader(reader)
	if e != nil {
		log.Println(e)
		return "nil"
	}
	list := doc.Find(".result ul li.res-list")
	targetUrl, ok := list.Eq(0).Find("h3 a").Attr("href")
	if !ok {
		log.Println("no found target url")
		return "nil"
	}
	newContent := pageSource(targetUrl)
	newDoc, e := goquery.NewDocumentFromReader(strings.NewReader(newContent))
	if e != nil {
		log.Println(e)
		return "nil"
	}
	return newDoc.Find("td.b12c").Text()

}

type C struct {
	number int
	ch     chan struct{}
	wg     *sync.WaitGroup
}

// New is used to initial a concurrent control object
func NewConCurrency(limit int) *C {
	return &C{
		wg:     &sync.WaitGroup{},
		ch:     make(chan struct{}, limit),
		number: limit,
	}
}

// Add is used to add a task
func (c *C) Add(n int) {
	c.wg.Add(n)
	for n > 0 {
		n--
		c.ch <- struct{}{}
	}
}

// Done is used to accomplish a task
func (c *C) Done() {
	c.wg.Done()
	<-c.ch
}

// Wait is used to wg for all tasks to be completed
func (c *C) Wait() {
	c.wg.Wait()
}
