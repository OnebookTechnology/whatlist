package models

type BigManRecommendList struct {
	ID              uint64 `json:"id"`
	BigManName      string `json:"big_man_name"`
	RecommendReason string `json:"recommend_reason"`
	RecommendTime   string `json:"recommend_time"`
	ImgUrl          string `json:"img_url"`
	IsRecommending  bool   `json:"is_recommending"`
}
