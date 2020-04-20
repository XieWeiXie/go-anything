package xc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestXcCity(t *testing.T) {
	var h HotelForXc
	var result interface{}
	content := h.DoForCity()
	json.Unmarshal([]byte(content), &result)
	fmt.Println(result)
}

func TestXcHotel(t *testing.T) {
	var h HotelForXc
	//var result interface{}
	content := h.DoForHotel(1)
	fmt.Println(content)
}

func TestXcDetail(t *testing.T) {
	var h HotelForXc
	content := h.DoForHotelDetail(608516)
	fmt.Println(content)
}

func TestXcTest(t *testing.T) {
	var h HotelForXc
	//https://m.ctrip.com/webapp/hotel/hoteldetail/608516.html&contrl=0&pay=0&latlon=#fromList
	url := "http://m.ctrip.com/restapi/get/list"
	//url := "https://m.ctrip.com/restapi/h5api/searchapp/search?source=mobileweb&action=autocomplete&contentType=json&keyword=北京"
	fmt.Println(h.DoForTest(url))

}
