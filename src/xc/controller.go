package xc

import (
	"bytes"
	"encoding/json"
	"github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
)

var (
	header    map[string]string
	head      map[string]interface{}
	extension map[string]string
)

func init() {
	header = make(map[string]string)
	header["Content-Type"] = "text/plain"

	var extension = make(map[string]string)

	extension = map[string]string{
		"name":  "appId",
		"value": "11048657",
	}

	head = map[string]interface{}{
		"cid":       "09031015211484252533",
		"ctok":      "",
		"cver":      "",
		"lang":      "01",
		"sid":       "",
		"syscode":   "30",
		"auth":      "",
		"sauth":     "",
		"extension": []interface{}{extension},
	}
}

type HotelForXc struct {
}

func (H HotelForXc) DoForCity() string {
	var body = make(map[string]interface{})
	body = map[string]interface{}{
		"fun":   "convertCity",
		"param": "{}",
		"head":  head,
	}
	bt, _ := json.Marshal(body)

	content := chromedp_helper.GetPageSourceHTTPPost(cityInfoApi, bytes.NewReader(bt), header)
	return content

}

func (H HotelForXc) DoForHotel(cityId int) string {
	var filterInfo = make(map[string]interface{})
	filterInfo = map[string]interface{}{
		"lowestPrice":      0,
		"highestPrice":     0,
		"filterItemList":   []interface{}{},
		"locationItemList": []interface{}{},
		"starItemList":     []interface{}{},
	}
	var body = make(map[string]interface{})
	body = map[string]interface{}{
		"cityID":         cityId,
		"districtID":     0,
		"checkinDate":    "",
		"checkoutDate":   "",
		"pageIndex":      0,
		"pageSize":       10,
		"orderItem":      "sort-45|1",
		"filterInfo":     filterInfo,
		"channel":        1,
		"userCoordinate": make(map[string]interface{}),
		"sourceFromTag":  "",
		"sessionId":      "aa894df4-3688-92f3-5992-6c7951455807",
		"preHotelIds":    "",
		"preCount":       0,
		"head":           head,
	}
	bt, _ := json.Marshal(body)
	return chromedp_helper.GetPageSourceHTTPPost(listUrl, bytes.NewReader(bt), header)
}

func (H HotelForXc) DoForHotelDetail(ht int) string {
	var body = make(map[string]interface{})
	body = map[string]interface{}{
		"hotelId":      ht,
		"checkinDate":  "",
		"checkoutDate": "",
		"head":         head,
	}
	bt, _ := json.Marshal(body)
	return chromedp_helper.GetPageSourceHTTPPost(detailApi, bytes.NewReader(bt), header)
}

func (H HotelForXc) DoForTest(url string) string {

	return chromedp_helper.GetPageSource(chromedp_helper.GetContextWithBackground(), "https://m.ctrip.com/webapp/hotel/hoteldetail/608516.html?atime=20200419&daylater=0&days=1&contrl=0&pay=0&discount=&latlon=&listindex=6&userLocationSearch=false")

}
