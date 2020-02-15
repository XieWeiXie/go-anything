package weixin

import "github.com/kataras/iris/v12"

func RegisterWeiXin(c iris.Party) {
	c.Get("/tags", tagsHandler)
	c.Get("/hotBanner", bannerHandler)
	c.Get("/hotSearch", hotSearchHandler)
	c.Get("/passages", passagesHandler)
	c.Get("/mediumPassage", editorHandler)
}
