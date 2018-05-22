package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BookListDao interface {
	// 根据书单ID获得书单内容
	GetBookList(listID uint64) (*models.List, error)
}
