package weixin

type (
	common struct {
		Topic string `json:"topic"`
		Url   string `json:"url"`
	}
	TagsResponse struct {
		common
	}
	HotSearch struct {
		common
	}
	Banner struct {
		common
	}
	Passage struct {
		common
		SubContent string `json:"sub_content"`
		Author     string `json:"author"`
		Date       string `json:"date"`
	}
	MediumEditor struct {
		common
		MediumName string `json:"medium_name"`
		Date       string `json:"date"`
	}
)
