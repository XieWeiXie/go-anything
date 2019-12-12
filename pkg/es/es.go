package es_operator

import (
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/wuxiaoxiaoshen/go-anything/configs"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type (
	EsClient struct {
		Client *elasticsearch.Client
	}
	esSetting struct {
		address  string
		user     string
		password string
	}
)

func newEsClient(address []string, user string, password string) *EsClient {
	cfg := elasticsearch.Config{
		Addresses: address,
		Username:  user,
		Password:  password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}
	client, e := elasticsearch.NewClient(cfg)
	if e != nil {
		log.Println(fmt.Sprintf("elasticsearch client: %s", e.Error()))
		return nil
	}
	return &EsClient{Client: client}
}

var DefaultEsClient = &EsClient{}

func EsInit() {
	e := configs.DefaultConfigs.LoadConfigs("es")
	a := e.(map[string]interface{})
	s := esSetting{
		address:  fmt.Sprintf("http://%s", a["address"].(string)),
		user:     a["user"].(string),
		password: a["password"].(string),
	}
	if os.Getenv(configs.ES_ADDRESS) != "" {
		s.address = os.Getenv(configs.ES_ADDRESS)
	}
	if os.Getenv(configs.ES_USER) != "" {
		s.user = os.Getenv(configs.ES_USER)
	}
	if os.Getenv(configs.ES_PASSWORD) != "" {
		s.password = os.Getenv(configs.ES_PASSWORD)
	}
	DefaultEsClient = newEsClient([]string{s.address}, s.user, s.password)
}

func (E EsClient) Close() {

}
