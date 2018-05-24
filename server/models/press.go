package models

type Press struct {
	PressID   uint64 `json:"press_id"`
	PressName string `json:"press_name"`
	PressImg  string `json:"press_img"`
	IsShowing bool   `json:"is_showing,omitempty"`
}
