package weixin

import "github.com/kataras/iris"

func RegisterWeiXin(c iris.Party) {
	c.Get("/tags", tagsHandler)
	c.Get("/hotBanner", bannerHandler)
	c.Get("/hotTopic", hotTopicHandler)
	c.Get("/passages", passagesHandler)
}
