package models

type UserAddressInfo struct {
	AddressId       int    `json:"address_id"`
	UserId          string `json:"user_id"`
	ReceiverNumber  int    `json:"receiver_number"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverAddress string `json:"receiver_address"`
	IsDefault       int    `json:"is_default"`
	CreateTime      string `json:"create_time"`
}

const (
	NotDefaultAddress = iota
	DefaultAddress
)
