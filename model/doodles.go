package model

import "time"

type GoogleDoodles struct {
	Id           uint      `json:"id"`
	ShareText    string    `json:"share_text"`
	Name         string    `json:"name"`
	Title        string    `json:"title"`
	Width        string    `json:"width"`
	Height       string    `json:"height"`
	Date         time.Time `json:"date"`
	UrlHigh      string    `json:"url_high"`
	Url          string    `json:"url"`
	AlternateUrl string    `json:"alternate_url"`
}

func (g GoogleDoodles) First(id int) (GoogleDoodles, error) {
	return GoogleDoodles{}, nil
}

func (g GoogleDoodles) Find() ([]GoogleDoodles, error) {
	return []GoogleDoodles{}, nil
}

func (g GoogleDoodles) Patch(id int) (GoogleDoodles, error) {
	return GoogleDoodles{}, nil
}

func (g GoogleDoodles) Delete(id int) (GoogleDoodles, error) {
	return GoogleDoodles{}, nil
}
func (g GoogleDoodles) Post(doodles GoogleDoodles) (GoogleDoodles, error) {
	return GoogleDoodles{}, nil
}
