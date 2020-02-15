package fund

import "github.com/kataras/iris/v12"

func RegisterFund(r iris.Party) {
	r.Get("/funds/{code:string}", getFundInfoHandler)
	r.Get("/global", getGlobalInfoHandler)
}
