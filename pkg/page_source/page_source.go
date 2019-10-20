package page_source

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetPageSource(url string, method string, body io.Reader) (string, error) {
	request, err := http.NewRequest(strings.ToUpper(method), url, body)
	if err != nil {
		log.Fatal(err)
		return "-1", err
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return "-1", err
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return "-1", err
	}
	return string(b), nil
}
