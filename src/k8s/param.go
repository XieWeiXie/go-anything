package k8s

// CreateTopicParam
type (
	CreateTopicParam struct {
		Name      string `json:"name" form:"name"`
		Partition int32  `json:"partition" form:"partition"`
		Factor    int16  `json:"factor" form:"factor"`
	}
)

func (C CreateTopicParam) IsValid() bool {
	if C.Name == "" {
		return false
	}
	if C.Partition < 0 || C.Partition > 100 {
		return false
	}
	return true
}

// CreateUserParam
type (
	CreateUserParam struct {
		Name string `json:"name" form:"name"`
		Age  int    `json:"age" form:"age"`
	}
)

func (C CreateUserParam) IsValid() bool {
	if C.Name == "" || C.Age == 0 {
		return false
	}
	return true
}
