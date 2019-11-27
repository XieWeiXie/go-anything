package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);unique_index" json:"name"`
	Age  int    `gorm:"type:integer" json:"age"`
}

func (U User) TableName() string {
	return "go_anything_user"
}

type UserSerializer struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func (U User) BasicSerializer() UserSerializer {
	return UserSerializer{
		Id:        U.ID,
		CreatedAt: U.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: U.UpdatedAt.Format("2006-01-02 15:04:05"),
		Name:      U.Name,
		Age:       U.Age,
	}
}

type UserMessage struct {
	Message User
	encode  []byte
	err     error
}

func (U *UserMessage) Length() int {
	b, e := json.Marshal(U.Message)
	U.encode = b
	U.err = e
	return len(b)
}
func (U *UserMessage) Encode() ([]byte, error) {
	return U.encode, U.err
}
