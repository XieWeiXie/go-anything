package weixin

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/kataras/iris/v12"
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
	tags.Find(`div.fieed-box a`).Each(func(i int, selection *goquery.Selection) {
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
	action.Url = HOST
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
	doc.Find(`div.sd-slider a.sd-slider-item`).Each(func(i int, selection *goquery.Selection) {
		var eachBanner Banner
		eachBanner.Topic = strings.TrimSpace(selection.Text())
		eachBanner.Url, _ = selection.Attr("href")
		b.Results = append(b.Results, eachBanner)
	})
}
func (b *BannerAction) String() string {
	r, _ := json.MarshalIndent(b.Results, " ", " ")
	return fmt.Sprintf("Banner: 主页横幅项: %s", string(r))
}

func bannerHandler(c iris.Context) {
	var b BannerAction
	b.Url = HOST
	b.Do()
	c.JSON(iris.Map{
		"data": b.Results,
	})
}

type HotSearchAction struct {
	Url     string `json:"url"`
	Results []HotSearch
}

func (H *HotSearchAction) Do() {
	source := pageSource(H.Url)
	doc, e := goquery.NewDocumentFromReader(strings.NewReader(source))
	if e != nil {
		log.Println(fmt.Sprintf("HotSearch: %s", e.Error()))
		return
	}
	doc.Find(`#topwords li`).Each(func(i int, selection *goquery.Selection) {
		var each HotSearch
		each.Url = selection.Find("a").AttrOr("href", "None Url")
		each.Topic = selection.Find("a").AttrOr("title", "No Title")
		H.Results = append(H.Results, each)
	})
}
func (H *HotSearchAction) String() string {
	r, _ := json.MarshalIndent(H.Results, " ", " ")
	return fmt.Sprintf("HotSearch: 搜索热词: %s", string(r))
}

func hotSearchHandler(c iris.Context) {
	var h HotSearchAction
	h.Url = HOST
	h.Do()
	c.JSON(iris.Map{
		"data": h.Results,
	})
}

type PassageAction struct {
	Url     string `json:"url"`
	Results []Passage
}

func dateFormat(value string) string {
	if id := strings.Index(value, "."); id != -1 {
		value = value[:id]
	}
	i, _ := strconv.ParseInt(value, 10, 64)
	values := time.Unix(i, 0)
	return values.Format("2006-01-02 15:04:05")
}
func (P *PassageAction) Do() {
	source := pageSource(P.Url)
	doc, e := goquery.NewDocumentFromReader(strings.NewReader(source))
	if e != nil {
		log.Println(fmt.Sprintf("Passage: %s", e.Error()))
		return
	}
	doc.Find(`.news-list li div.txt-box`).Each(func(i int, selection *goquery.Selection) {
		var eachPassage Passage
		eachPassage.Topic = strings.TrimSpace(selection.Find("h3").Text())
		eachPassage.Url = selection.Find("h3 a").AttrOr("href", "No Url")
		eachPassage.SubContent = selection.Find(`p.txt-info`).Text()
		eachPassage.Author = selection.Find("div.s-p a.account").Text()
		date, _ := selection.Find("div").Attr("t")
		eachPassage.Date = dateFormat(date)
		P.Results = append(P.Results, eachPassage)
	})
}
func (P *PassageAction) String() string {
	r, _ := json.MarshalIndent(P.Results, " ", " ")
	return fmt.Sprintf("Passage: 文章列表项: %s", string(r))
}

func passagesHandler(c iris.Context) {
	var p PassageAction
	p.Url = HOST
	p.Do()
	c.JSON(iris.Map{
		"data": p.Results,
	})
}

type EditorAction struct {
	Url     string `json:"url"`
	Results []MediumEditor
}

func (E *EditorAction) Do() {
	source := pageSource(E.Url)
	doc, e := goquery.NewDocumentFromReader(strings.NewReader(source))
	if e != nil {
		log.Println(fmt.Sprintf("Editor: %s", e.Error()))
		return
	}
	doc.Find(`.news-list-right li`).Each(func(i int, selection *goquery.Selection) {
		selection = selection.Find("div.txt-box")
		var eachEditor MediumEditor
		eachEditor.Url = selection.Find(`p.p1 a`).AttrOr("href", "No Url")
		eachEditor.Topic = selection.Find(`p.p1 a`).Text()
		eachEditor.Date = selection.Find(`p.p2 span`).Text()
		eachEditor.MediumName = selection.Find(`p.p2 a`).Text()
		E.Results = append(E.Results, eachEditor)
	})
}
func (E *EditorAction) String() string {
	r, _ := json.MarshalIndent(E.Results, " ", " ")
	return fmt.Sprintf("Editor: 编辑推荐项: %s", string(r))
}
func editorHandler(c iris.Context) {
	var e EditorAction
	e.Url = HOST
	e.Do()
	c.JSON(iris.Map{
		"data": e.Results,
	})
}
