package model

import "time"

type (
	MakeMoney struct {
		Title      string    `json:"title"`
		CreateTime time.Time `json:"create_time"`
		TopicId    int64     `json:"topic_id"`
		Owner      Owner     `json:"owner"`
		Text       string    `json:"text"`
		Images     []Image   `json:"images"`
		Comments   []Comment `json:"comments"`
		Type       string    `json:"type"`
		Question   Question
		Answer     Answer
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
	Question struct {
		Owner        Owner
		QuestionText string `json:"question_text"`
		Images       []Image
	}
	Answer struct {
		Owner      Owner
		AnswerText string `json:"answer_text"`
		Images     []Image
	}
)
