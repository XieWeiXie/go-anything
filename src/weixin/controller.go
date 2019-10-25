package weixin

import "github.com/kataras/iris"

type TagsAction struct {
}

func (t TagsAction) Do()            {}
func (t TagsAction) String() string {}

func tagsHandler(c iris.Context)     {}
func bannerHandler(c iris.Context)   {}
func hotTopicHandler(c iris.Context) {}
func passagesHandler(c iris.Context) {}
