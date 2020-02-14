package es_operator

import (
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/wuxiaoxiaoshen/go-anything/configs"
	"io"
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

// EsIndexInterface ...
type EsIndexInterface interface {
	Index() string
	IndexBody() io.Reader
	RecordBody() io.Reader
	SortFiled() string
	Size() int
}

// CreateIndex...
// 结构尽可能的通用，采用 interface 入参形式
func (E EsClient) CreateIndex(index EsIndexInterface) bool {
	response, e := E.Client.Indices.Create(
		index.Index(),
		E.Client.Indices.Create.WithBody(index.IndexBody()),
		E.Client.Indices.Create.WithPretty(),
		E.Client.Indices.Create.WithHuman(),
	)
	if e != nil {
		log.Println(fmt.Sprintf("es create index fail %s", e.Error()))
		return false
	}
	log.Println(response.String())

	log.Println(fmt.Sprintf("es create index response %s", response.String()))
	return true
}

// GetIndex...
// 获取 index 的详细信息
// 类似的 es 接口还有: GetMapping, GetSetting...
func (E EsClient) GetIndex(index EsIndexInterface) (bool, *esapi.Response) {
	response, e := E.Client.Indices.Get(
		[]string{index.Index()},
		E.Client.Indices.Get.WithPretty(),
		E.Client.Indices.Get.WithIncludeDefaults(true),
		E.Client.Indices.Get.WithIncludeTypeName(true),
	)
	if e != nil {
		log.Println(fmt.Sprintf("es get index mapping and setting fail: %s", e.Error()))
		return false, nil
	}
	return true, response

}

func (E EsClient) ExistsIndex(index EsIndexInterface) bool {
	response, e := E.Client.Indices.Exists(
		[]string{index.Index()},
		E.Client.Indices.Exists.WithHuman(),
		E.Client.Indices.Exists.WithPretty(),
	)
	if e != nil {
		log.Println(fmt.Sprintf("es index exists fail"))
		return false
	}
	s := response.String()
	log.Println(s)
	return true
}

func (E EsClient) IndexSettings(index EsIndexInterface) bool {
	return false
}

func (E EsClient) InsertRecord(id string, index EsIndexInterface) (bool, *esapi.Response) {
	response, e := E.Client.Create(
		index.Index(),
		id,
		index.RecordBody(),
		E.Client.Create.WithDocumentType("_doc"),
		E.Client.Create.WithPretty(),
		E.Client.Create.WithHuman())
	if e != nil {
		log.Println(fmt.Sprintf("es index : %s insert record :%s fail ", index.Index(), index.RecordBody()))
		return false, nil
	}
	return true, response
}

func (E EsClient) BulkInsert(index EsIndexInterface) (bool, *esapi.Response) {
	response, e := E.Client.Bulk(
		index.RecordBody(),
		E.Client.Bulk.WithHuman(),
		E.Client.Bulk.WithPretty(),
		E.Client.Bulk.WithIndex(index.Index()),
	)
	if e != nil {
		return false, nil
	}
	return true, response
}

func (E EsClient) Search(index EsIndexInterface) (bool, *esapi.Response) {
	response, e := E.Client.Search(
		E.Client.Search.WithIndex(index.Index()),
		E.Client.Search.WithPretty(),
		E.Client.Search.WithHuman(),
		E.Client.Search.WithBody(index.RecordBody()),
		E.Client.Search.WithSize(index.Size()),
		E.Client.Search.WithSort(index.SortFiled()),
	)
	if e != nil {
		log.Println(fmt.Sprintf("es search fail"))
		return false, nil
	}
	return true, response
}
