package models

type EveryDayRecommend struct {
	ID             uint64 `json:"id,omitempty"`
	RecommendTime  string `json:"recommend_time,omitempty"`
	ImgUrl         string `json:"img_url,omitempty"`
	ListID         uint64 `json:"list_id,omitempty"`
	ListName       string `json:"list_name,omitempty"`
	ListBriefIntro string `json:"list_brief_intro,omitempty"`
}
