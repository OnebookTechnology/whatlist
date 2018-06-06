package models

type BigManRecommendList struct {
	ID              uint64  `json:"id,omitempty"`
	ListID          uint64  `json:"list_id,omitempty"`
	BigManName      string  `json:"big_man_name,omitempty"`
	RecommendReason string  `json:"recommend_reason,omitempty"`
	RecommendTime   string  `json:"recommend_time,omitempty"`
	ImgUrl          string  `json:"img_url,omitempty"`
	IsRecommending  bool    `json:"is_recommending,omitempty"`
	Price           float64 `json:"price"`
}
