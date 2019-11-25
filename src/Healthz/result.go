package Healthz

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

type (
	Disk struct {
		All  uint64 `json:"all"`
		Used uint64 `json:"used"`
		Free uint64 `json:"free"`
	}
)
