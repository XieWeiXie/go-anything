package internal

type (
	Action interface {
		Do() interface{}
		String() string
	}
)
