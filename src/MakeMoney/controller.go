package MakeMoney

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/wuxiaoxiaoshen/go-anything/model"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
	"log"
)

type (
	RealMakeMoneyAction struct {
		Url string
	}
)

var DefaultRealMakeMoneyAction = NewRealMakeMoneyAction(topicURL)

func NewRealMakeMoneyAction(url string) *RealMakeMoneyAction {
	return &RealMakeMoneyAction{
		Url: url,
	}
}

func (r RealMakeMoneyAction) Do() []model.MakeMoney {
	var body = make(map[string]interface{})
	body = map[string]interface{}{
		"limit":          10,
		"skip":           0,
		"order":          "-time",
		"_method":        "GET",
		"_ApplicationId": "app",
	}

	b, _ := json.Marshal(body)
	reader := bytes.NewReader(b)
	header := make(map[string]string)
	header["Content-Type"] = "text/plain"
	content := chromedp_helper.GetPageSourceHTTPPost(r.Url, reader, header)
	//log.Println(content)
	js := gjson.Parse(content)
	c := NewConCurrency(30)
	var results []model.MakeMoney
	for _, i := range js.Get("results").Array() {
		c.Add(1)
		var (
			makeMoney model.MakeMoney
			title     string
			id        int64
			time      string
		)
		title = i.Get("title").String()
		id = i.Get("topic_id").Int()
		time = i.Get("time").String()
		log.Println(title, id, time, formatBody(id))
		go func(id int64, money model.MakeMoney) {
			defer c.Done()
			newContent := chromedp_helper.GetPageSourceHTTPPost(contentURL, formatBody(id), header)
			//log.Println(newContent)
			newJs := gjson.Parse(newContent).Get("results").Array()[0]
			makeMoney.Title = title
			makeMoney.CreateTime = formatTime(newJs.Get("create_time.iso").String())
			owner := newJs.Get("content.talk.owner")
			makeMoney.Owner.Name = owner.Get("name").String()
			makeMoney.Owner.AvatarURL = owner.Get("avatar_url").String()

			text := newJs.Get("content.talk.text").String()
			makeMoney.Text = formatText(text)

			images := newJs.Get("content.talk.images").Array()
			for _, i := range images {
				var image model.Image
				image.Height = i.Get("large.height").String()
				image.Width = i.Get("large.width").String()
				image.URL = i.Get("large.url").String()
				makeMoney.Images = append(makeMoney.Images, image)
			}

			comments := newJs.Get("content.show_comments").Array()
			for _, i := range comments {
				var comment model.Comment
				comment.AuthorName = i.Get("owner.name").String()
				comment.Content = i.Get("text").String()
				makeMoney.Comments = append(makeMoney.Comments, comment)
			}
			//log.Println(makeMoney)
			results = append(results, makeMoney)

		}(id, makeMoney)
		log.Println()
	}
	c.Wait()
	return results
}
