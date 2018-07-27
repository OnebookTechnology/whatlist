package models

type BiggieCollect struct {
	CollectId   int    `json:"collect_id"`
	UserId      string `json:"user_id"`
	BiggieId    int    `json:"biggie_id"`
	CollectTime string `json:"collect_time"`
}
