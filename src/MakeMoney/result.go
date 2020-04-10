package MakeMoney

type (
	ResultForMakeMoney struct {
		Title    string
		Content  string
		Comments []Comment
	}

	Comment struct {
		Name string
		Doc  string
	}
)
