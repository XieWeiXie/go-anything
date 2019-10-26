package Jav

import "encoding/base64"

func toRawString(v string) string {
	r, _ := base64.StdEncoding.DecodeString(v)
	return string(r)
}
