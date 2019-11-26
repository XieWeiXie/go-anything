package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);unique_index" json:"name"`
	Age  int    `gorm:"type:integer" json:"age"`
}

func (U User) TableName() string {
	return "go_anything_user"
}
