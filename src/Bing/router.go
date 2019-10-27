package Bing

import "github.com/kataras/iris"

func RegisterBing(c iris.Party) {
	c.Get("/images/{date:int}", imageHandler)
}
