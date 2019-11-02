package fund

import "github.com/kataras/iris"

func RegisterFund(r iris.Party) {
	r.Get("/funds/{code:string}", getFundInfoHandler)
	r.Get("/global", getGlobalInfoHandler)
}
