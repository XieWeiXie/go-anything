package router

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/go-anything/src/Doodles"
	"github.com/wuxiaoxiaoshen/go-anything/src/holiday"
	"github.com/wuxiaoxiaoshen/go-anything/src/zhihu"
	"os"
	"os/signal"
	"syscall"

	"github.com/wuxiaoxiaoshen/go-anything/src/k8s"

	"github.com/wuxiaoxiaoshen/go-anything/src/Healthz"

	"github.com/wuxiaoxiaoshen/go-anything/src/fund"

	"github.com/wuxiaoxiaoshen/go-anything/src/Bing"

	"github.com/wuxiaoxiaoshen/go-anything/src/Jav"

	"github.com/wuxiaoxiaoshen/go-anything/src/weixin"

	"github.com/wuxiaoxiaoshen/go-anything/src/Railway12306"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}
func register(app *iris.Application) {
	app.PartyFunc("/v1/api/status", Healthz.RegisterHealth)
	app.PartyFunc("/v1/api/12306", Railway12306.RegisterRailWay12306)
	app.PartyFunc("/v1/api/weixin", weixin.RegisterWeiXin)
	app.PartyFunc("/v1/api/jav", Jav.RegisterJav)
	app.PartyFunc("/v1/api/bing", Bing.RegisterBing)
	app.PartyFunc("/v1/api/tt", fund.RegisterFund)
	app.PartyFunc("/v1/api/k8s", k8s.RegisterForK8s)
	app.PartyFunc("/v1/api/zhihu", zhihu.RegisterForZhiHu)
	app.PartyFunc("/v1/api/gov", holiday.RegisterForHoliday)
	app.PartyFunc("/v1/api/doodles", Doodles.RegisterWithDoodles)

}
func Run(port string) {
	app := newApp()
	register(app)
	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	go quit(quitSignal)
	app.Run(iris.Addr(fmt.Sprintf(":%s", port)))
}

func quit(c chan os.Signal) {
	for i := range c {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("\nBye!")
			os.Exit(0)
		default:
			fmt.Println(i)

		}
	}
}
