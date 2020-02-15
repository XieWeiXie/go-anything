package Healthz

import "github.com/kataras/iris/v12"

func RegisterHealth(c iris.Party) {
	c.Get("/health", healthZHandler)
	c.Get("/health_cpu", healthCpuHandler)
}
