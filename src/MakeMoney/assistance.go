package MakeMoney

import (
	"bytes"
	"encoding/json"
	"io"
	"regexp"
	"sync"
	"time"
)

func formatBody(topicId int64) io.Reader {
	var body = make(map[string]interface{})
	var where = make(map[string]interface{})
	where = map[string]interface{}{
		"topic_id": topicId,
	}
	body = map[string]interface{}{
		"where":          where,
		"limit":          1,
		"_method":        "GET",
		"_ApplicationId": "app",
	}
	b, _ := json.Marshal(body)
	return bytes.NewReader(b)
}

func formatTime(string2 string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", string2)
	ft := t.Format("2006-01-02 15:04:05")
	t, _ = time.Parse("2006-01-02 15:04:05", ft)
	return t
}

var (
	patternOne = `<e type="hashtag".*?/>`
	patternTwo = `<e type="web".*?/>）`
)

func formatText(string2 string) string {
	re1 := regexp.MustCompile(patternOne)
	s1 := re1.ReplaceAllString(string2, "")
	re2 := regexp.MustCompile(patternTwo)
	s2 := re2.ReplaceAllString(s1, "")
	return s2

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
