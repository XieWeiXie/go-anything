package main

import (
	"log"
	"os"

	"github.com/wuxiaoxiaoshen/go-anything/configs"

	"github.com/wuxiaoxiaoshen/go-anything/cmd"
)

var Env string

func Execute() error {
	e := cmd.ROOT.Execute()
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}
func init() {
	log.Println("Env: ", Env)
}
func main() {

	if Env == "" {
		configs.Env = "dev"
	} else {
		configs.Env = Env
	}
	e := Execute()
	if e != nil {
		os.Exit(1)
		return
	}

}
