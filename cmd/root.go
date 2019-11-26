package cmd

import (
	"log"

	kafka_operator "github.com/wuxiaoxiaoshen/go-anything/pkg/kafka"

	redis_operator "github.com/wuxiaoxiaoshen/go-anything/pkg/redis"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/mysql"

	"github.com/wuxiaoxiaoshen/go-anything/router"

	"github.com/spf13/cobra"
)

var ROOT = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("Web Start...")
		log.Println("Step 0: Configs...")
		log.Println("Step 1: Mysql...")
		mysql_operator.MySQLInit()

		log.Println("Step 2: Redis...")
		redis_operator.RedisInit()
		//
		log.Println("Step 3: Kafka...")
		kafka_operator.KafkaInit()

	},
	Run: func(cmd *cobra.Command, args []string) {
		router.Run("8888")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("End Step 0: Mysql close...")
		defer mysql_operator.DefaultMySQLAction.Close()
		log.Println("End Step 1: Redis close...")
		//defer redis_operator.DefaultRedisAction.Close()
		log.Println("End Step 2: Kafka close...")
		//defer kafka_operator.DefaultAsyncProducer.Close()
	},
}

func init() {
	ROOT.AddCommand(STATION)
	ROOT.AddCommand(WX)
	ROOT.AddCommand(JAV)
	ROOT.AddCommand(FundCmd)
	ROOT.AddCommand(ConfigsCmd)
}
