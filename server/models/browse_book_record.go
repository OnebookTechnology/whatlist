package models

type BrowseBookRecord struct {
	BrowseBookId int `json:"browse_book_id"`
	Book
	UserId     string `json:"user_id"`
	BrowseTime string `json:"browse_time"`
}
