package model

import (
	"fmt"
	"log"

	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
)

type UserKafka struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	err    error
	encode []byte
}

func (U *UserKafka) Length() int {
	b, e := json.Marshal(U)
	if e != nil {
		log.Println(fmt.Sprintf("model: UserKafka: %s", e.Error()))
		return -1
	}
	U.encode = b
	U.err = e
	return len(b)
}

func (U *UserKafka) Encode() ([]byte, error) {
	return U.encode, U.err
}
