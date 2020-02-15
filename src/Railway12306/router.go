package Railway12306

import "github.com/kataras/iris/v12"

func RegisterRailWay12306(c iris.Party) {
	c.Get("/stations", railWayStationHandler)               // 站点
	c.Get("/stations/search", railWayStationIsExistHandler) // 搜索站点
	c.Get("/tickets", ticketsHandler)                       // 查询车票
	c.Get("/type_for_tickets", typeForTicketsHandler)       // 票的类型
	//c.Get("/ticketPrice")                                   // 票价
	//c.Get("/")                                              // 车次经过的站点
}
