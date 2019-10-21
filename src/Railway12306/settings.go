package Railway12306

import (
	"fmt"
	"strings"
)

var STATION_URL = "https://kyfw.12306.cn/otn/resources/js/framework/station_name.js?station_version=1.9114"
var trainUrl = "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=%s&leftTicketDTO.from_station=%s&leftTicketDTO.to_station=%s&purpose_codes=%s"
var priceUrl = "https://kyfw.12306.cn/otn/leftTicket/queryTicketPrice?train_no=%s&from_station_no=%s&to_station_no=%s&seat_types=%s&train_date=%s"
var railwayUrl = "https://kyfw.12306.cn/otn/queryTrainInfo/query?leftTicketDTO.train_no=67000G600401&leftTicketDTO.train_date=2019-10-25&rand_code="
var search = "https://search.12306.cn/search/v1/train/search?keyword=G60&date=20191020"

func RailWayURL(date string, fromStation string, toStation string, codes string) string {
	return fmt.Sprintf(trainUrl, date, strings.ToUpper(fromStation), strings.ToUpper(toStation), strings.ToUpper(codes))
}

func RailWayPrice(number string, date string, fromStationNo string, toStationNo string, seatTypes string) string {
	return fmt.Sprintf(priceUrl, number, fromStationNo, toStationNo, seatTypes, date)
}

var TickType map[int]string

func init() {
	TickType = make(map[int]string)
	TickType[1] = "adult"      // 成人票
	TickType[2] = "child"      // 孩票
	TickType[3] = "student"    // 学生票
	TickType[4] = "disability" // 伤残军人票
}

// 参考：https://github.com/sunhailin-Leo/12306-Go
const (
	HXHBed       int = 33 // 动卧
	SeatBusiness int = 32 // 商务座
	SeatFirst    int = 31 // 一等座
	SeatSecond   int = 30 // 二等座
	HardSeat     int = 29 // 硬座
	HardBed      int = 28 // 硬卧
	NoSeat       int = 26 // 无座
	SeatSpecial  int = 25 // 特等座
	SoftSeat     int = 23 // 软座
)
