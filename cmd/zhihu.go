package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/es"
	"github.com/wuxiaoxiaoshen/go-anything/src/zhihu"
	"log"
	"strconv"
	"time"
)

var ZhiHuCmd = &cobra.Command{
	Use: "zhihuhot",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Step 1: ElasticSearch...")
		es_operator.EsInit()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("End 1: ElasticSearch...")
		es_operator.DefaultEsClient.Close()
	},
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	index := zhihu.ResultForZhiHu{}
	if ok := es_operator.DefaultEsClient.ExistsIndex(index); ok {
		log.Println(fmt.Sprintf("index exists: %s", index.Index()))
	} else {
		if ok := es_operator.DefaultEsClient.CreateIndex(index); !ok {
			log.Println(fmt.Sprintf("index %s create fail", index.Index()))
			return
		}
	}
	for j, i := range zhihu.HotResultController() {
		id := fmt.Sprintf("%d_"+strconv.FormatFloat(float64(time.Now().UnixNano()), 'f', 0, 32), j)
		fmt.Println(es_operator.DefaultEsClient.InsertRecord(id, i))
	}
}
