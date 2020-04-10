package model

type ActionInterface interface {
	First(id int) (interface{}, error)
	Find() ([]interface{}, error)
	Patch(id int) (interface{}, error)
	Post(interface{}) (interface{}, error)
	Delete(id int) (bool, error)
}
