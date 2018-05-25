package models

type BrowseListRecord struct {
	BrowseListId int    `json:"browse_list_id"`
	ListId       int    `json:"list_id"`
	UserId       string `json:"user_id"`
	BrowseTime   string `json:"browse_time"`
}
