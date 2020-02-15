package Jav

import "github.com/kataras/iris/v12"

func RegisterJav(c iris.Party) {
	c.Get("/detail/{code:string}", codeDetailHandler)
}
