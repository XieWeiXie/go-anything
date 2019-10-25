package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/wuxiaoxiaoshen/go-anything/src/Railway12306"

	"github.com/spf13/cobra"
)

var STATION = &cobra.Command{
	Use:   "stations",
	Short: "get station in local",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start...")
		dir := "./src/Railway12306/stations.go"
		os.Remove(dir)
	},
	Run: func(cmd *cobra.Command, args []string) {
		getStations()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("End...")
	},
}

func getStations() bool {
	result := Railway12306.ExportRailWayStationHelper()
	if len(result) == 0 {
		log.Fatal("get stations fail")
		return false
	}
	dir := "./src/Railway12306/stations.go"
	stationsCH := `map[string]string{`
	//stationsEn := `map[string]string{`
	stationCode := `map[string]string{`

	for _, i := range result {
		stationsCH += fmt.Sprintf(`"%s":"%s",`, i.CH, i.EN)
		//stationsEn += fmt.Sprintf(`"%s":"%s",`, i.Phonetic, i.EN)
		stationCode += fmt.Sprintf(`"%s":"%s",`, i.EN, i.CH)
	}
	stationsCH += "}"
	//stationsEn += "}"
	stationCode += "}"

	code := `package Railway12306
var ChPlace = map[string]string{}

var CodePlace = map[string]string{}
func init() {
	ChPlace = %s
	CodePlace = %s
}
`
	goCode := fmt.Sprintf(code, stationsCH, stationCode)
	err := ioutil.WriteFile(dir, []byte(goCode), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return false
	}
	cmd := exec.Command("go", "fmt", dir)
	cmd.Start()
	return true

}
