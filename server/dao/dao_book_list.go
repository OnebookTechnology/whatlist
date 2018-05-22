package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BookListDao interface {
	// 根据书单ID获得书单内容
	GetList(listID uint64) (*models.List, error)
	// 获得最新的六个书单
	GetLatestSixList()([]*models.List, error)
	// 获得推荐的六个书单
	GetRecommendSixList()([]*models.List, error)
}
