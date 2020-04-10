package model

import "time"

type (
	MakeMoney struct {
		Title      string    `json:"title"`
		CreateTime time.Time `json:"create_time"`
		Owner      Owner     `json:"owner"`
		Text       string    `json:"text"`
		Images     []Image   `json:"images"`
		Comments   []Comment `json:"comments"`
	}
	Owner struct {
		Name      string `json:"name"`
		AvatarURL string `json:"avatar_url"`
	}
	Image struct {
		Width  string `json:"width"`
		Height string `json:"height"`
		URL    string `json:"url"`
	}
	Comment struct {
		AuthorName string `json:"name"`
		Content    string `json:"content"`
	}
)
