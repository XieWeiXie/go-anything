package Jav

import "github.com/kataras/iris"

func RegisterJav(c iris.Party) {
	c.Get("/detail/{code:string}", codeDetailHandler)
}
