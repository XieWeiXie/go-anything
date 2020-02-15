package Healthz

import (
	"log"
	"net/http"
	"strings"
	"syscall"

	"github.com/kataras/iris/v12"
)

func healthZHandler(c iris.Context) {
	_, e := c.JSON(iris.Map{
		"data":   "pong",
		"status": "ok",
		"code":   http.StatusOK,
	})
	log.Println(e)
}

func newDisk(path string) *Disk {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free
	return &Disk{
		All:  all,
		Used: used,
		Free: free,
	}
}
func (D Disk) GetTotal(unit string) uint64 {
	unit = strings.ToUpper(unit)
	switch unit {
	case "B":
		return D.All
	case "KB":
		return D.All / KB
	case "MB":
		return D.All / MB
	default:
		return D.All / GB
	}
}
func (D Disk) GetUsed(unit string) uint64 {
	unit = strings.ToUpper(unit)
	switch unit {
	case "B":
		return D.Used
	case "KB":
		return D.Used / KB
	case "MB":
		return D.Used / MB
	default:
		return D.Used / GB
	}
}
func (D Disk) GetFree(unit string) uint64 {
	return D.GetTotal(unit) - D.GetUsed(unit)
}

func healthCpuHandler(c iris.Context) {
	disk := newDisk("/")
	_, e := c.JSON(iris.Map{
		"All(GB)":  disk.GetTotal("GB"),
		"Used(GB)": disk.GetUsed("GB"),
		"Free(GB)": disk.GetFree("GB"),
	})
	if e != nil {
		log.Println(e)
	}

}
