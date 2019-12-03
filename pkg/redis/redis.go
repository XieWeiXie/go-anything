package redis_operator

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wuxiaoxiaoshen/go-anything/configs"

	"github.com/gomodule/redigo/redis"
)

type poolAction struct {
	Pool redis.Pool
}

var (
	DefaultRedisAction *poolAction
	addr               string
	auth               string
)

type (
	redisSettings struct {
		port string
		auth string
		host string
	}
)

func init() {

}

func RedisInit() {
	r := configs.DefaultConfigs.LoadConfigs("redis")
	a := r.(map[string]interface{})
	log.Println(fmt.Sprintf("Keys: Redis: %#v", r))
	setting := redisSettings{
		port: a["port"].(string),
		auth: a["auth"].(string),
		host: a["host"].(string),
	}
	if os.Getenv(configs.REDIS_HOST) != "" {
		setting.host = os.Getenv(configs.REDIS_HOST)

	}
	if os.Getenv(configs.REDIS_PORT) != "" {
		setting.port = os.Getenv(configs.REDIS_PORT)
	}
	addr = fmt.Sprintf("%s:%s", setting.host, setting.port)
	auth = setting.auth
	DefaultRedisAction = newPoolAction(addr, auth)
}

func newPoolAction(addr string, auth string) *poolAction {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			if auth == "" {
				return redis.Dial("tcp", addr)
			} else {
				c, _ := redis.Dial("tcp", addr)
				_, e := c.Do("AUTH", auth)
				if e != nil {
					log.Println("redis_operator: redis auth:", e.Error())
				}
				return c, nil
			}
		},
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return &poolAction{Pool: pool}
}

func (P *poolAction) Close() {
	defer P.Pool.Close()
}

func (P *poolAction) Get() redis.Conn {
	return P.Pool.Get()
}
