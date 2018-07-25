package models

type BiggieList struct {
	ListId         int     `json:"list_id"`
	BiggieId       int     `json:"biggie_id"`
	ListName       string  `json:"list_name"`
	ListIntro      string  `json:"list_intro"`
	ListCreateTime string  `json:"list_create_time"`
	ListClickCount int     `json:"list_click_count"`
	ListImg        string  `json:"list_img"`
	Books          []*Book `json:"books,omitempty"`
}
