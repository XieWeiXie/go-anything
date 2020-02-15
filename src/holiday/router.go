package holiday

import "github.com/kataras/iris/v12"

func RegisterForHoliday(r iris.Party) {
	r.Get("/holiday/search", getHolidayHandler)
}
