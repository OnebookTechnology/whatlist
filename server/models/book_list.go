package models

type BookList struct {
	BookListID uint64 `json:"book_list_id"`
	BookISBN uint64 `json:"book_isbn"`
	ListID uint64 `json:"list_id"`
}
