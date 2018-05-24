package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type DiscoverDao interface {
	GetDiscoverList(pageNum, pageCount int) ([]*models.Discover, error)
	GetDiscoverDetail(id int) (*models.Discover, error)
	AddReadNum(id int) error
	AddLikeNum(id int) error
}
