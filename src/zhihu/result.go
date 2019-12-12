package zhihu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type (
	ResultForZhiHu struct {
		Title   string               `json:"title"`
		Text    string               `json:"text"`
		Url     string               `json:"url"`
		Answer  ResultForZhiHuAnswer `json:"answer"`
		mapping io.Reader
	}
	ResultForZhiHuAnswer struct {
		CreatedAt       time.Time `json:"created_at"`
		ApiForQuestions string    `json:"api_for_questions"`
		ApiForAnswer    string    `json:"api_for_answer"`
		QueryKeys       []string  `json:"query_keys"`
	}
)

// 实现 interface 方便操作数据
func (R ResultForZhiHu) Index() string {
	return "zhihu_billboard"
}
func (R *ResultForZhiHu) Mapping() {
	// todo: why settings not work?
	body := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"title": map[string]interface{}{
					"type": "keyword",
				},
				"text": map[string]interface{}{
					"type": "text",
				},
				"url": map[string]interface{}{
					"type":  "text",
					"index": false,
				},
				"answer": map[string]interface{}{
					"properties": map[string]interface{}{
						"created_at": map[string]interface{}{
							"type":  "date",
							"index": true,
						},
						"api_for_questions": map[string]interface{}{
							"type": "text",
						},
						"api_for_answer": map[string]interface{}{
							"type": "text",
						},
					},
				},
			},
		},
		"settings": map[string]interface{}{
			"number_of_shards":   2,
			"number_of_replicas": 2,
		},
	}
	m, _ := json.Marshal(body)
	R.mapping = bytes.NewReader(m)
}
func (R ResultForZhiHu) IndexBody() io.Reader {
	return R.mapping
}
func (R ResultForZhiHu) RecordBody() io.Reader {
	m, _ := json.Marshal(R)
	return bytes.NewReader(m)
}
func (R ResultForZhiHu) Size() int {
	// Default 10
	return 10
}
func (R ResultForZhiHu) SortFiled() string {
	return fmt.Sprintf(R.Answer.CreatedAt.Format("2006-01-02 15:04:05"))
}
