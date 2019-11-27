package cmd

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/wuxiaoxiaoshen/go-anything/model"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/kafka"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/mysql"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/redis"

	"github.com/spf13/cobra"
)

var KubernetesCmd = &cobra.Command{
	Use:   "k8s",
	Short: "migrate db",
	Long:  "use cli to create db",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("k8s: cobra command: step 1: Mysql")
		mysql_operator.MySQLInit()
		log.Println("k8s: cobra command: step 2: Redis")
		redis_operator.RedisInit()
		log.Println("k8s: cobra command: step 3: Kafka")
		kafka_operator.KafkaInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("k8s: cobra command: start migrate db")
		/*
			1: migrate db
			2: redis keys
			3: kafka topic
		*/
		migrateTable()
		log.Println("k8s: cobra command: end migrate db")
		log.Println("k8s: cobra command: start redis ping")
		redisPing()
		log.Println("k8s: cobra command: end redis ping")
		log.Println("k8s: cobra command: start kafka topics")
		kafkaTopic()
		log.Println("k8s: cobra command: end kafka topics")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("k8s: cobra command: End 1: Mysql")
		defer mysql_operator.DefaultMySQLAction.Close()
		log.Println("k8s: cobra command: End 2: Redis")
		defer redis_operator.DefaultRedisAction.Close()
		log.Println("k8s: cobra command: End 3: Kafka")
		defer kafka_operator.DefaultAsyncProducer.Close()
	},
}

func migrateTable() {
	mysql_operator.DefaultMySQLAction.DB.AutoMigrate(
		&model.User{},
	)
}
func redisPing() {
	c := redis_operator.DefaultRedisAction.Get()
	pong, e := redis.String(c.Do("PING"))
	if e != nil {
		log.Println("k8s: cobra command: pong")
		return
	}
	log.Println(pong)
}

func kafkaTopic() {
	r := kafka_operator.DefaultKafkaClusterAdminAction.GetTopic()
	log.Println(fmt.Sprintf("k8s: cobra command: topics : %+v", r))
}
