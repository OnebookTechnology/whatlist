package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BookDao interface {
	// 添加图书
	AddBook(book *models.Book) error
}