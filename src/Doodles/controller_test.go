package Doodles

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/wuxiaoxiaoshen/go-anything/model"
	"html/template"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDoodles(t *testing.T) {
	d := NewGoogleDoodlesAction("https://www.google.com/doodles/json/2020/4?hl=zh_CN")
	results := d.Do()
	var returnResult model.MonthGoogleDoodles
	returnResult.Month = "4"
	returnResult.Doodles = results
	ts := template.New("doodles")
	box := packr.NewBox("../../doc/doodles/template")
	templateText, _ := box.FindString("doodles.md")
	tem, _ := ts.Parse(templateText)
	var byt bytes.Buffer
	e := tem.Execute(&byt, returnResult)
	if e != nil {
		log.Println(e)
		return
	}
	toSave(returnResult.Month, byt.Bytes())
}

func toSave(title string, content []byte) error {
	replacer := strings.NewReplacer("，", "", "\"", "", " ", "",
		"/", "")
	title = replacer.Replace(title)
	f, e := os.Create(fmt.Sprintf("../../doc/doodles/%s.md", title))
	if e != nil {
		log.Println(e)
		return e
	}
	f.Write(content)
	defer f.Close()
	return nil
}
