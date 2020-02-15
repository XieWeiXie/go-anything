package k8s

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"

	"github.com/wuxiaoxiaoshen/go-anything/model"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/mysql"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/redis"

	"github.com/kataras/iris/v12"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/kafka"
)

// Mysql

func allUsersHandler(c iris.Context) {
	var users []model.User
	mysql_operator.DefaultMySQLAction.DB.Find(&users)
	if len(users) == 0 {
		_, _ = c.JSON(iris.Map{
			"data": "No data",
		})
		return
	}
	var result []model.UserSerializer
	for _, i := range users {
		result = append(result, i.BasicSerializer())
	}
	_, _ = c.JSON(iris.Map{
		"data": result,
	})
}
func singUserHandler(c iris.Context) {
	id, _ := c.Params().GetInt("id")
	var user model.User
	if dbError := mysql_operator.DefaultMySQLAction.DB.Where("id = ?", id).First(&user).Error; dbError != nil {
		_, _ = c.JSON(iris.Map{
			"data": "No data",
		})
		return
	}
	c.JSON(iris.Map{
		"data": user.BasicSerializer(),
	})

}
func createUserHandler(c iris.Context) {
	/*
		1. mysql
		2. redis
		3. kafka
	*/
	var param CreateUserParam
	if e := c.ReadJSON(&param); e != nil {
		log.Println("k8s: create user handler: ", e.Error())
		return
	}
	if ok := param.IsValid(); !ok {
		log.Println("k8s: create user handler: ", ok)
		return
	}

	red := redis_operator.DefaultRedisAction.Get()
	userKeys := model.UserKeys{
		Name: param.Name,
		Age:  param.Age,
	}
	keys := userKeys.KeysFormat()
	log.Println("k8s: format keys :", keys)
	ok, e := redis.Bool(red.Do("EXISTS", keys))
	if e != nil {
		log.Println("k8s: create user handler :", e.Error())
		return
	}
	if ok {
		_, _ = c.JSON(iris.Map{
			"data": fmt.Sprintf("redis: name: %s is already existing", param.Name),
		})
		return
	}

	_, e = red.Do("SET", keys, param.Age)
	if e != nil {
		_, _ = c.JSON(iris.Map{
			"data": fmt.Sprintf("name: %s set in redis fail : %v", param.Name, e),
		})
		return
	}

	var user model.User

	if dbError := mysql_operator.DefaultMySQLAction.DB.Where("name = ?", param.Name).First(&user).Error; dbError != nil {
		user = model.User{
			Name: param.Name,
			Age:  param.Age,
		}
		if dbError := mysql_operator.DefaultMySQLAction.DB.Save(&user).Error; dbError != nil {
			log.Println("k8s: create user handler :", dbError)
			return
		}
	} else {
		_, _ = c.JSON(iris.Map{
			"data": fmt.Sprintf("db: name: %s is already exists", param.Name),
		})
		return
	}

	go func() {
		userMessage := &model.UserMessage{
			Message: user,
		}
		kafka_operator.DefaultAsyncProducer.Run("go-anything", userMessage)

	}()
	_, _ = c.JSON(iris.Map{
		"data": user.BasicSerializer(),
	})
}
func deleteUserHandler(c iris.Context) {}

// Kafka

func allTopicHandler(c iris.Context) {
	topics := kafka_operator.DefaultKafkaClusterAdminAction.GetTopic()
	if topics == nil {
		_, _ = c.JSON(iris.Map{
			"data": "no topics",
		})
	} else {
		_, _ = c.JSON(iris.Map{
			"data": topics,
		})
	}
}

func createTopicHandler(c iris.Context) {
	var param CreateTopicParam
	e := c.ReadJSON(&param)
	if e != nil {
		log.Println("k8s: create topic handler :", e.Error())
		return
	}
	if ok := param.IsValid(); ok {
		log.Println("k8s: create topic handler :", false)
		return
	}
	result := kafka_operator.DefaultKafkaClusterAdminAction.CreateTopic(
		param.Name,
		param.Partition,
		param.Factor,
	)
	_, _ = c.JSON(iris.Map{
		"data": result,
	})
}
