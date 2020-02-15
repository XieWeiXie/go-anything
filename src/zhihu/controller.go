package zhihu

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

func getZhiHuHandler(c iris.Context) {
	response := HotResultController()
	if response != nil {
		_, _ = c.JSON(iris.Map{"data": response})
		return
	} else {
		_, _ = c.JSON(iris.Map{
			"error": http.StatusBadRequest,
		})
		return
	}
}

func deleteZhiHuHandler(c iris.Context) {

}
