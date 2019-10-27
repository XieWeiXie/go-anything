package router

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/go-anything/src/Bing"

	"github.com/wuxiaoxiaoshen/go-anything/src/Jav"

	"github.com/wuxiaoxiaoshen/go-anything/src/weixin"

	"github.com/wuxiaoxiaoshen/go-anything/src/Railway12306"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}
func register(app *iris.Application) {
	app.PartyFunc("/v1/api/12306", Railway12306.RegisterRailWay12306)
	app.PartyFunc("/v1/api/weixin", weixin.RegisterWeiXin)
	app.PartyFunc("/v1/api/jav", Jav.RegisterJav)
	app.PartyFunc("/v1/api/bing", Bing.RegisterBing)

}
func Run(port string) {
	app := newApp()
	register(app)
	app.Run(iris.Addr(fmt.Sprintf(":%s", port)))
}
