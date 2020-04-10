package Doodles

import "testing"

func TestDoodles(t *testing.T) {
	d := NewGoogleDoodlesAction("https://www.google.com/doodles/json/2020/4?hl=zh_CN")
	d.Do()
}
