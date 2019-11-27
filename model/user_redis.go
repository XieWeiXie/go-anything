package model

import "fmt"

var (
	formatString string
	formatHash   string
)

func init() {
	formatString = "user::keys::%s"
	formatHash = "user::hash"
}

type UserKeys struct {
	Name string `json:"name"` // user::keys::{name}::{age}
	Age  int    `json:"age"`
}

func (U UserKeys) KeysFormat() string {
	return fmt.Sprintf(formatString, U.Name)
}
