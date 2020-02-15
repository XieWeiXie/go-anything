package zhihu

import "github.com/kataras/iris/v12"

func RegisterForZhiHu(c iris.Party) {
	c.Get("/hot/_search", getZhiHuHandler)
	c.Delete("/hot/_delete", deleteZhiHuHandler)
}
