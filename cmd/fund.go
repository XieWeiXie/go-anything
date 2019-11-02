package cmd

import (
	"github.com/spf13/cobra"
)

var FundCmd = &cobra.Command{
	Use: "fund",
	Run: func(cmd *cobra.Command, args []string) {
		//source := chromedp_helper.GetPageSourceHTTP("http://pdfm.eastmoney.com/EM_UBG_PDTI_Fast/api/js?id=0006512&rtntype=5")
		//pattern := `\((.*?)\)`
		//r1 := regexp.MustCompile(pattern)
		//jsonSource := r1.FindAllStringSubmatch(source, -1)
		//fmt.Println(jsonSource[0][1])
		//d := gjson.Parse(jsonSource[0][1]).Get("info.c")
		//fmt.Println(d)
	},
}
