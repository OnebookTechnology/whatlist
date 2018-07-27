package models

type BiggieSubscribeRecord struct {
	SubscribeId   int    `json:"subscribe_id"`
	UserId        string `json:"user_id"`
	BiggieId      int    `json:"biggie_id"`
	SubscribeTime string `json:"subscribe_time"`
	OrderId       string `json:"order_id"`
}
