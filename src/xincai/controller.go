package xincai

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gobuffalo/packr"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/handle"
	"golang.org/x/sync/errgroup"
	"html/template"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type XinCai struct {
	Author string
	RootUrl string
	Des string
	Page int
	urls []string
}

func NewXinCai(author string, root string) XinCai{
	return XinCai{
		Author: author,
		RootUrl: root,
	}
}


func (x XinCai) Do() (err error){
	return nil
}

func (x *XinCai) GetTotalPage() (err error) {
	page := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithBackground(), x.RootUrl)
	doc, docErr:= goquery.NewDocumentFromReader(strings.NewReader(page))
	if docErr != nil {
		err = docErr
		return docErr
	}
	content := doc.Find(".pages").Text()
	pattern := `\d+`
	re := regexp.MustCompile(pattern)
	results := re.FindAllString(content, -1)
	if len(results)!=2 {
		errors.New("no pages tag")
		return
	}
	pages, strconvErr := strconv.Atoi(results[1])
	if strconvErr != nil {
		err = strconvErr
		return
	}
	x.Page = pages
	x.Des = doc.Find("div[class=sidebar]  div[class=widget_body]").Eq(0).Find("p").Text()
	return
}

func (x *XinCai) GetAllUrls() (err error)  {
	g := new(errgroup.Group)
	for i:=1;i<=x.Page;i++{
		one := i
		g.Go(func() error {
			page := fmt.Sprintf(x.RootUrl + "/page/" + strconv.Itoa(one))
			fmt.Println(page)
			err := x.GetOnePageUrl(page)
			if err!= nil {
				log.Println("GetAllUrls", err)
			}
			return err
		})
	}
	if err = g.Wait();err!=nil {
		log.Println("GetAllUrls", err)
	}
	return  err
}

func (x *XinCai) GetOnePageUrl(url string) (err error){
	page := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithBackground(), url)
	doc, docErr := goquery.NewDocumentFromReader(strings.NewReader(page))
	if docErr!=nil {
		log.Println("docErr", docErr)
		err = docErr
		return
	}

	doc.Find("h2 a").Each(func(i int, selection *goquery.Selection) {
		 href , ok := selection.Attr("href")
		 if ok {
			 x.urls = append(x.urls, href)
		 }
	})
	return err
}

func (x *XinCai) GetOnePage(url string) (passage Passage, err error){
	page := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithBackground(), url)
	doc, docErr := goquery.NewDocumentFromReader(strings.NewReader(page))
	if docErr != nil {
		log.Println("GetOnePage", docErr)
		err = docErr
		return
	}
	bigContent := doc.Find("#content")
	passage.Title = bigContent.Find("h2").Text()
	passage.PublishTime = bigContent.Find(".post_date").Text()
	bigContent.Find(".blog_content").Each(func(i int, selection *goquery.Selection) {
		passage.Content += selection.Text()
	})
	var (
		hand handle.Handle
	)
	hand = &handle.StringHandle{Content: passage.Content}
	content, _ := hand.Handler()
	passage.Content = content.(string)
	return
}

func (x XinCai) GetAllPages() (err error){
	x.GetAllUrls()
	pack := packr.NewBox(".")
	bye, _  := pack.FindString("passage.tmpl")
	tmpl, _ := template.New("passage").Parse(bye)
	os.Mkdir("./passages", os.ModePerm)
	for _, i := range x.urls {
		if !strings.Contains(i, "archives") {
			continue
		}
		page := i
		one, _ := x.GetOnePage(page)
		f, err := os.Create("./passages/" + one.Title + "-" + one.PublishTime + ".md")
		if err != nil {
			log.Println("create", err)
		}
		err = tmpl.Execute(f, one)
		if err != nil {
			log.Println("markdown", err)
		}
	}
	return
}