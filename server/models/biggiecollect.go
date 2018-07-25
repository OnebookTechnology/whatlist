package models

type BiggieCollect struct {
	CollectId int `json:"collect_id"`
	UserId    int `json:"user_id"`
	BiggieId  int `json:"biggie_id"`
}
