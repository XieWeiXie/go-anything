package internal

type (
	Action interface {
		Do()
		String() string
	}
)
