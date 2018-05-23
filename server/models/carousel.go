package models

type Carousel struct {
	ID        uint64 `json:"id,omitempty"`
	ImgUrl    string `json:"img_url,omitempty"`
	ISBN      uint64 `json:"isbn,omitempty"`
	IsShowing bool   `json:"is_showing,omitempty"`
	AddTime   string `json:"add_time,omitempty"`
}
