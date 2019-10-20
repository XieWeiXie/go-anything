package Railway12306

import (
	"go-anything/pkg/error_http"
	"go-anything/pkg/page_source"
	"net/http"
	"strings"

	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
)

func railWayStationHelper() ([]StationInfo, *error_http.HttpError) {
	var result []StationInfo
	var e error_http.HttpError
	content, err := page_source.GetPageSource(STATION_URL, "get", nil)
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
	var params SearchParam

	result, err := railWayStationHelper()
	if err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{"error": err})
		c.StopExecution()
		return
	}
	for _, i := range result {

	}
}
