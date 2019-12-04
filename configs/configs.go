package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	GoAnythingConfigs struct {
		v *viper.Viper
	}
)

var (
	DefaultConfigs    = &GoAnythingConfigs{}
	rootPathForConfig string
	Env               string
)

func init() {
	rootPathForConfig = "./configs"
	DefaultConfigs = newGoAnythingConfig()
}

func newGoAnythingConfig() *GoAnythingConfigs {
	v := viper.New()
	v.AddConfigPath(rootPathForConfig)
	v.SetConfigName("settings")
	v.AddConfigPath("./")
	v.SetConfigType("yaml")
	return &GoAnythingConfigs{v: v}
}

func (G *GoAnythingConfigs) LoadConfigs(key string) interface{} {
	if e := G.v.ReadInConfig(); e != nil {
		log.Println(fmt.Sprintf("configs: LoadConfigs: %s", e.Error()))
		return "-1"
	}

	key = fmt.Sprintf("%s.%s", Env, key)
	log.Println("configs: LoadConfigs: key:", key)
	return G.v.Get(key)
}
