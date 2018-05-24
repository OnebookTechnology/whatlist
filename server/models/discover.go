package models

type Discover struct {
	DiscoverId  int      `json:"discover_id"`
	NickName    string   `json:"nick_name"`
	AvatarUrl   string   `json:"avatar_url"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Content     string   `json:"content"`
	PublishTime string   `json:"publish_time"`
	ReadNum     int      `json:"read_num"`
	LikeNum     int      `json:"like_num"`
	Picture     []string `json:"picture"`
	IsThumb     bool     `json:"is_thumb,omitempty"`
}
