package models

type BiggieBooks struct {
	ListId     int     `json:"list_id"`
	ISBN       int64   `json:"isbn"`
	BookPrice  float64 `json:"price"`
	Recommend  string  `json:"recommend"`
	BookName   string  `json:"book_name"`
	AuthorName string  `json:"author_name"`
	BookIcon   string  `json:"book_icon"`
}
