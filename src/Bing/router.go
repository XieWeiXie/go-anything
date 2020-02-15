package Bing

import "github.com/kataras/iris/v12"

func RegisterBing(c iris.Party) {
	c.Get("/images/{date:int}", imageHandler)
}
