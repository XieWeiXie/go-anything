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
		"limit":          1300,
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
		)
		id = i.Get("topic_id").Int()
		title = i.Get("title").String()
		makeMoney.TopicId = id
		go func(id int64, money model.MakeMoney) {
			defer c.Done()
			newContent := chromedp_helper.GetPageSourceHTTPPost(contentURL, formatBody(id), header)
			//log.Println(newContent)
			array := gjson.Parse(newContent).Get("results").Array()
			if len(array) != 1 {
				return
			}
			newJs := array[0]

			makeMoney.Title = title
			makeMoney.CreateTime = formatTime(newJs.Get("create_time.iso").String())
			makeMoney.Type = newJs.Get("content.type").String()
			if makeMoney.Type == "talk" {
				owner := newJs.Get("content.talk.owner")
				makeMoney.Owner.Name = owner.Get("name").String()
				makeMoney.Owner.AvatarURL = owner.Get("avatar_url").String()

				text := newJs.Get("content.talk.text").String()
				makeMoney.Text = formatText(text)
				log.Println(text)

				images := newJs.Get("content.talk.images").Array()
				makeMoney.Images = toImages(images)
			} else if makeMoney.Type == "q&a" {
				owner := newJs.Get("content.question.owner")
				makeMoney.Question.Owner.Name = owner.Get("name").String()
				makeMoney.Question.Owner.AvatarURL = owner.Get("avatar_url").String()
				makeMoney.Question.QuestionText = formatText(newJs.Get("content.question.text").String())
				makeMoney.Question.Images = toImages(newJs.Get("content.question.images").Array())

				answerOwner := newJs.Get("content.answer.owner")
				makeMoney.Answer.Owner.Name = answerOwner.Get("name").String()
				makeMoney.Answer.Owner.AvatarURL = answerOwner.Get("avatar_url").String()
				makeMoney.Answer.Images = toImages(newJs.Get("content.answer.images").Array())
				makeMoney.Answer.AnswerText = formatText(newJs.Get("content.answer.text").String())
			}

			comments := newJs.Get("content.show_comments").Array()
			makeMoney.Comments = toComments(comments)
			log.Println(makeMoney)
			results = append(results, makeMoney)

		}(id, makeMoney)
	}
	c.Wait()
	return results
}

func toImages(array []gjson.Result) []model.Image {
	var results []model.Image
	if len(array) == 0 {
		return results
	}
	for _, i := range array {
		var image model.Image
		image.Height = i.Get("large.height").String()
		image.Width = i.Get("large.width").String()
		image.URL = i.Get("large.url").String()
		results = append(results, image)
	}
	return results
}

func toComments(array []gjson.Result) []model.Comment {
	var results []model.Comment
	if len(array) == 0 {
		return results
	}
	for _, i := range array {
		var comment model.Comment
		comment.AuthorName = i.Get("owner.name").String()
		comment.Content = i.Get("text").String()
		results = append(results, comment)
	}
	return results
}
