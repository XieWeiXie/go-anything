package holiday

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"strings"
)

func getHolidayHandler(c iris.Context) {
	if !c.URLParamExists("query") {
		c.JSON(iris.Map{
			"error": "url params should contain query",
		})
		return
	}
	var results []*GovResultForHoliday
	query := c.URLParams()["query"]
	years := strings.Split(strings.TrimSpace(query), ",")
	for _, year := range years {
		v := urlEncode(map[string]string{
			"q": fmt.Sprintf(HOLIDAT_INTEGER, year),
		})
		url := fmt.Sprintf(HOST, v)
		results = append(results, &GovResultForHoliday{RawUrl: url, Query: fmt.Sprintf(HOLIDAT_INTEGER, year)})
	}
	channel := NewConCurrency(len(years))
	for index, i := range results {
		channel.Add(1)
		go func(index int, i *GovResultForHoliday) {
			log.Println(index, i.RawUrl)
			content := pageSource(i.RawUrl)
			newContent := getContent(content, i.RawUrl)
			log.Println(newContent)
			results[index].RawContent = newContent
			defer channel.Done()
		}(index, i)
	}
	channel.Wait()
	c.JSON(iris.Map{
		"data": results,
	})

}

func getHolidayByYearHandler(c iris.Context) {}
