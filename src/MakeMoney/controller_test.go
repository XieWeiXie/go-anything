package MakeMoney

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr"
	"html/template"
	"log"
	"os"
	"strings"
	"testing"
)

func TestRealMakeMoneyAction_Do(t *testing.T) {

	mm := NewRealMakeMoneyAction(topicURL)
	content := mm.Do()

	for _, i := range content {
		ts := template.New("doc")
		box := packr.NewBox("../../doc/makemoney")
		templateText, e := box.FindString("template.md")
		tem, e := ts.Parse(templateText)
		if e != nil {
			log.Println(e)
			return
		}
		var byt bytes.Buffer
		e = tem.Execute(&byt, i)
		if e != nil {
			log.Println(e)
			return
		}
		toSave(i.Title, byt.Bytes())

	}

}

func toSave(title string, content []byte) error {
	replacer := strings.NewReplacer("，", "", "\"", "", " ", "",
		"/", "")
	title = replacer.Replace(title)
	f, e := os.Create(fmt.Sprintf("../../doc/makemoney/%s.md", title))
	if e != nil {
		log.Println(e)
		return e
	}
	f.Write(content)
	defer f.Close()
	return nil
}
