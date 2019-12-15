package holiday

import "github.com/kataras/iris"

func RegisterForHoliday(r iris.Party) {
	r.Get("/holiday/search", getHolidayHandler)
}
