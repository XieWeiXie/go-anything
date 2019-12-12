package zhihu

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	patternJson = `</script><script id="js-initialData" type="text/json">(.*?)</script>`
)

// hotZhiHuSource...
// 获取源代码，再使用正则表达式截取
func hotZhiHuSource() string {
	content := chromedp_helper.GetPageSource(chromedp_helper.GetContextWithBackground(), HOT)
	reg, e := regexp.Compile(patternJson)
	if e != nil {
		log.Println("regexp fail")
		return "nil"
	}
	results := reg.FindAllStringSubmatch(content, -1)
	//fmt.Println(results[0][1])
	return results[0][1]
}

// hotZhiHuList...
// 解析 json 数据获取到列表
func hotZhiHuList(content string) []ResultForZhiHu {
	var response []ResultForZhiHu
	result := gjson.Parse(content).Get("initialState.topstory.hotList").Array()
	for _, i := range result {
		collection := i.Get("target")
		var i ResultForZhiHu
		i.Title = collection.Get("titleArea.text").String()
		i.Text = collection.Get("excerptArea.text").String()
		i.Url = urlReplace(collection.Get("link.url").String())
		i.Answer.CreatedAt = hotQuestCreatedAt(urlAPI(i.Url))
		i.Answer.ApiForAnswer = answer(urlAPI(i.Url))
		i.Answer.ApiForQuestions = urlAPI(i.Url)
		i.Answer.QueryKeys = []string{"limit", "offset", "sort_by"}
		response = append(response, i)
	}
	return response
}

func hotQuestCreatedAt(url string) time.Time {
	content := chromedp_helper.GetPageSourceHTTP(url)
	createdAt := gjson.Parse(content).Get("created").String()
	unix, _ := strconv.ParseInt(createdAt, 10, 32)
	return time.Unix(unix, 0)
}

// urlReplace...
// url 编码替换字符
func urlReplace(url string) string {
	replacer := strings.NewReplacer("\u002F", "/")
	return replacer.Replace(url)
}

// urlAPI ...
// restful 风格的 api
func urlAPI(url string) string {
	replacer := strings.NewReplacer("question", "api/v4/questions")
	return replacer.Replace(url)
}

// answer ...
// 答案 api
func answer(url string) string {
	return fmt.Sprintf("%s/answers", url)
}

// HotResultController ...
// 对外暴露, cmd 中需要使用...
func HotResultController() []ResultForZhiHu {
	return hotZhiHuList(hotZhiHuSource())
}
