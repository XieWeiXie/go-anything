package Healthz

import "github.com/kataras/iris"

func RegisterHealth(c iris.Party) {
	c.Get("/health", healthZHandler)
	c.Get("/health_cpu", healthCpuHandler)
}
