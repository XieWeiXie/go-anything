package Railway12306

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wuxiaoxiaoshen/go-anything/pkg/error_http"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/page_source"

	"github.com/tidwall/gjson"

	"github.com/kataras/iris"

	"github.com/kataras/iris/v12/context"
)

func ExportRailWayStationHelper() []StationInfo {
	r, _ := railWayStationHelper()
	return r
}
func railWayStationHelper() ([]StationInfo, *error_http.HttpError) {
	var result []StationInfo
	var e error_http.HttpError
	r := page_source.Request{}
	r.R, _ = http.NewRequest("GET", STATION_URL, nil)
	content, err := page_source.GetPageSource(nil, r)
	if err != nil {
		e = error_http.HttpError{
			Code:   http.StatusBadRequest,
			Reason: "http fail",
		}
		return nil, &e
	}
	replacer := strings.NewReplacer(";", "", "'", "", "'", "", "|2880", "")
	stationsList := strings.Split(replacer.Replace(content), "=")
	stations := strings.Split(stationsList[1], "|")
	var AllStations []StationInfo
	for index := 0; index < len(stations); index += 5 {
		if index < len(stations) {
			var temp StationInfo
			temp = StationInfo{
				Code:          stations[index],
				CH:            stations[index+1],
				EN:            stations[index+2],
				Phonetic:      stations[index+3],
				ShortPhonetic: stations[index+4],
			}
			AllStations = append(AllStations, temp)
		}
	}
	result = AllStations
	return result, nil
}

func railWayStationHandler(c context.Context) {
	result, err := railWayStationHelper()
	if err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{"error": err})
		c.StopExecution()
		return
	}
	c.JSON(iris.Map{"stations": result})
}

func railWayStationIsExistHandler(c context.Context) {
	var params searchParam
	enOk := c.URLParamExists("name")
	if enOk {
		params.Name = c.URLParam("name")
	} else {
		c.JSON(iris.Map{
			"error": error_http.HttpError{
				Code:   http.StatusBadRequest,
				Reason: fmt.Sprintf("add url param: name"),
			},
		})
		c.StatusCode(http.StatusBadRequest)
		return
	}

	result, err := railWayStationHelper()
	if err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{"error": err})
		return
	}
	name := strings.ToLower(params.Name)
	var hitResult StationInfo
	var hit bool
	for _, i := range result {
		if i.Phonetic == name || i.ShortPhonetic == name || i.CH == name || i.EN == name {
			hitResult = i
			hit = true
			break
		}
	}
	if hit {
		c.JSON(iris.Map{
			"data": hitResult,
		})
		return
	}
	c.JSON(iris.Map{
		"data": fmt.Sprintf("no result"),
	})
}

func ticketsHandler(c context.Context) {
	var params ticketsParams
	form := c.URLParams()
	t, _ := c.URLParamInt("type")
	params.Type = t
	params = ticketsParams{
		Date:      form["date"],
		FromPlace: form["from_place"],
		ToPlace:   form["to_place"],
		Type:      t,
	}
	fmt.Println(params)
	if e := params.Valid(); e != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{
			"error": e.Error(),
		})
		return
	}
	var from, to string
	from = codeForStations(params.FromPlace)
	to = codeForStations(params.ToPlace)

	types := strings.ToUpper(TickType[params.Type])
	url := RailWayURL(params.Date, from, to, types)
	//fmt.Println(url)
	r := page_source.Request{}
	r.R, _ = http.NewRequest("GET", url, nil)
	r.AddParams("Cookie", "JSESSIONID=3FAEC8F056F8C60C189BDBA195A74CD8;")
	content, err := page_source.GetPageSource(nil, r)
	if err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	var result []TickStationInfo
	for _, i := range gjson.Parse(content).Get("data.result").Array() {
		list := strings.Split(i.String(), "|")
		//fmt.Println(list)
		var temp TickStationInfo
		temp = TickStationInfo{
			TrainCode:        list[3],
			BeginPlace:       CodePlace[list[4]],
			EndPlace:         CodePlace[list[5]],
			FromPlace:        CodePlace[list[6]],
			ToPlace:          CodePlace[list[7]],
			StartTime:        list[8],
			EndTime:          list[9],
			Over:             list[10],
			HighTicket:       list[SeatBusiness],
			FirstTicket:      list[SeatFirst],
			SecondTicket:     list[SeatSecond],
			HighSoftTicket:   "",
			SoftTicket:       "",
			LowSoftTicket:    list[HXHBed],
			SecondSoftTicket: list[HardBed],
			SoftSeat:         list[SoftSeat],
			HardSeat:         list[HardSeat],
		}
		result = append(result, temp)
	}
	c.JSON(iris.Map{
		"data": result,
	})
}

func typeForTicketsHandler(c context.Context) {
	c.JSON(iris.Map{
		"data": TickType,
	})
}
