package xincai

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"html/template"
	"os"
	"testing"
)

func TestCaiXin(t *testing.T){
	caiXin := NewXinCai("硅谷王川", "https://wangchuan.blog.caixin.com/")
	caiXin.GetTotalPage()
	fmt.Println(caiXin)

}

func TestCaiXinPage(t *testing.T) {
	//caiXin := NewXinCai("硅谷王川", "https://wangchuan.blog.caixin.com/")
	//result, _ := caiXin.GetOnePage("https://wangchuan.blog.caixin.com/archives/245907")
	//fmt.Println(result)

}

func TestTempl(t *testing.T){

	pack := packr.NewBox(".")
	bye, _  := pack.FindString("passage.tmpl")
	tmpl, _ := template.New("passage").Parse(bye)
	f, _  := os.Create("test.md")
	passage := Passage{
		Title:       "硅谷王川",
		PublishTime: "2021-05-15",
		RealURL:     "",
		Content:     "这个时间呢很单纯的",
	}
	tmpl.Execute(f, passage)
}

func TestGetAllPassage(t *testing.T){
	caiXin := NewXinCai("硅谷王川", "https://wangchuan.blog.caixin.com")
	caiXin.GetTotalPage()
	caiXin.GetAllPages()
	//fmt.Println(caiXin.GetOnePage("http://wangchuan.blog.caixin.com/archives/133315"))
}