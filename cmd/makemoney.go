package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/wuxiaoxiaoshen/go-anything/src/MakeMoney"
	"html/template"
	"log"
	"os"
)

var MakeMoneyCmd = &cobra.Command{
	Use: "makeMoney",
	Run: func(cmd *cobra.Command, args []string) {
		mm := MakeMoney.DefaultRealMakeMoneyAction
		content := mm.Do()
		t := template.New("doc")
		for _, i := range content {
			tem, e := t.ParseFiles("./doc/makemoney/template.md")
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

	},
}

func toSave(title string, content []byte) error {
	f, e := os.Create(title)
	if e != nil {
		log.Println(e)
		return e
	}
	f.Write(content)
	defer f.Close()
	return nil
}
