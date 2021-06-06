package handle

import (
	"strings"
)

type Handle interface {
	Handler() (interface{}, error)
}

const (
	StringHand = iota
)
func NewHandle(_type int) Handle {
	switch _type {
	case StringHand:
		return &StringHandle{}
	default:
		return nil
	}
}

type StringHandle struct {
	Content string
}
func (s *StringHandle) Handler() (interface{},  error){
	replacer := strings.NewReplacer("&nbsp", "")
	s.Content = replacer.Replace(s.Content)
	return s.Content, nil
}
