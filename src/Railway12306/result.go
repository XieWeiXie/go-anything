package Railway12306

type StationInfo struct {
	CH            string `json:"ch"`
	EN            string `json:"en"`
	Phonetic      string `json:"phonetic"`
	ShortPhonetic string `json:"short_phonetic"`
	Code          string `json:"code"`
}

type TickStationInfo struct {
	TrainCode        string `json:"train_code(车次)"`
	BeginPlace       string `json:"begin_place(始站)"`
	EndPlace         string `json:"end_place(终点)"`
	FromPlace        string `json:"from_place(出发站)"`
	ToPlace          string `json:"to_place(到达站)"`
	StartTime        string `json:"start_time(出发时间)"`
	EndTime          string `json:"end_time(到达时间)"`
	Over             string `json:"over(历时)"`
	HighTicket       string `json:"high_ticket(商务座/特等座)"`
	FirstTicket      string `json:"first_ticket(一等座)"`
	SecondTicket     string `json:"second_ticket(二等座)"`
	HighSoftTicket   string `json:"high_soft_ticket(高级软卧)"`
	SoftTicket       string `json:"soft_ticket(软卧)"`
	LowSoftTicket    string `json:"low_soft_ticket(动卧)"`
	SecondSoftTicket string `json:"hard_soft_ticket(硬卧)"`
	SoftSeat         string `json:"soft_seat(软座)"`
	HardSeat         string `json:"hard_seat(硬座)"`
}
