package Doodles

import "github.com/kataras/iris/v12"

func RegisterWithDoodles(party iris.Party) {
	party.Get("/search/{search:string}", getGoogleDoodlesSearch)  // 通过标题搜索
	party.Get("/date/{date:string}", getGoogleDoodlesDate)        // 通过时间搜索
	party.Get("/year/{year:int}/{top:int}", getGoogleDoodlesYear) // 通过年份搜索
}
