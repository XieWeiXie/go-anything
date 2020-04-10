package Doodles

import (
	"github.com/tidwall/gjson"
	"time"
)

func toDate(array []gjson.Result) time.Time {
	y := array[0].Int()
	m := array[1].Int()
	d := array[2].Int()
	return time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.Local)
}
