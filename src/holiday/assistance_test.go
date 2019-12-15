package holiday

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestPageSource(t *testing.T) {

	v := url.Values{}
	v.Add("q", fmt.Sprintf(HOLIDAY_STRING, 2020))
	url_ := fmt.Sprintf(HOST, v.Encode())
	log.Println(url_)
	content := pageSource(url_)
	log.Println(content)
}

func TestGetContent(t *testing.T) {
	short := urlEncode(map[string]string{
		"q": fmt.Sprintf(HOLIDAY_STRING, 2019),
	})
	content := pageSource(fmt.Sprintf(HOST, short))
	newContent := getContent(content, fmt.Sprintf(HOST, short))
	fmt.Println(newContent)
}
