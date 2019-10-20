package Railway12306

import (
	"fmt"
	"time"
)

type searchParam struct {
	Name string `json:"en"`
}

type ticketsParams struct {
	Date      string `json:"date"`
	FromPlace string `json:"from_place"`
	ToPlace   string `json:"to_place"`
	Type      int    `json:"type"`
}

var defaultTicketsParams = ticketsParams{
	Date:      "2019-10-21",
	FromPlace: "IUQ",
	ToPlace:   "HEQ",
	Type:      1,
}

func (t ticketsParams) Valid() error {
	d, e := time.Parse("2006-01-02", t.Date)
	if e != nil {
		return fmt.Errorf("date format error")
	}
	if d.Before(time.Now()) {
		return fmt.Errorf("date should be today or future")
	}
	if _, ok := TickType[t.Type]; !ok {
		return fmt.Errorf("type should be in (1~4)")
	}
	return nil
}
