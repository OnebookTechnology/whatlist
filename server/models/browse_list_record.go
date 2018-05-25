package models

type BrowseListRecord struct {
	BrowseListId int `json:"browse_list_id"`
	List
	UserId     string `json:"user_id"`
	BrowseTime string `json:"browse_time"`
}
