package holiday

type (
	GovResultForHoliday struct {
		Query      string `json:"query"`
		RawUrl     string `json:"raw_url"`
		RawContent string `json:"raw_content"`
	}
	GovResultForHolidays []GovResultForHoliday
	SpecialHoliday       struct {
		Year    string `json:"year"`
		Start   string `json:"start"`
		End     string `json:"end"`
		Content string `json:"content"`
	}
)
