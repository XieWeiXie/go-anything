package error_http

import "fmt"

type HttpError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func (h HttpError) Error() string {
	return fmt.Sprintf("Status Code: %d\nReason: %s", h.Code, h.Reason)
}
