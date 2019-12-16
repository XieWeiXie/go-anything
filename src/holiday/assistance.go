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

func urlDecode(v string) url.Values {
	r, e := url.ParseQuery(v)
	if e != nil {
		log.Println(e)
		return nil
	}
	return r
}

func getContent(content string, urls string) string {
	reader := bytes.NewReader([]byte(content))
	doc, e := goquery.NewDocumentFromReader(reader)
	if e != nil {
		log.Println(e)
		return "nil"
	}
	q, _ := url.ParseQuery(urls)
	log.Println("q", q.Encode())
	query := urlDecode(q.Encode()).Get("q")
	var result string
	doc.Find(".result ul li.res-list").Each(func(i int, selection *goquery.Selection) {
		log.Println(selection.Find("h3 a").Text(), query)
		if selection.Find("h3 a").Text() == query {
			log.Println(selection.Find("h3 a").Text(), query)
			targetUrl, ok := selection.Find("h3 a").Attr("href")
			if !ok {
				log.Println("no found target url")
				return
			}
			log.Println(targetUrl)
			newContent := pageSource(targetUrl)
			newDoc, e := goquery.NewDocumentFromReader(strings.NewReader(newContent))
			if e != nil {
				log.Println(e)
				return
			}
			result = newDoc.Find("td.b12c").Text()
		}
	})
	return result

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
