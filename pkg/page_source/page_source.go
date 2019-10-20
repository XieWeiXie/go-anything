package page_source

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	R *http.Request
}

func (r *Request) AddParams(key string, v string) {
	r.R.Header.Add(key, v)
}

func GetPageSource(body io.Reader, r Request) (string, error) {
	client := http.DefaultClient
	response, err := client.Do(r.R)
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
