package models

type Carousel struct {
	ID        uint64 `json:"id"`
	ImgUrl    string `json:"img_url"`
	ISBN      uint64 `json:"isbn"`
	IsShowing bool   `json:"is_showing"`
	AddTime   string `json:"add_time"`
}
