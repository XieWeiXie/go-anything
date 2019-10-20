package Railway12306

import "github.com/kataras/iris"

func RegisterRailWay12306(c iris.Party) {
	c.Get("/stations", railWayStationHandler)
	c.Get("/stations/search", railWayStationIsExistHandler)

}
