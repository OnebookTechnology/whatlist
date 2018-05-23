package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BookListDao interface {
	// 根据书单ID获得书单内容
	GetListDetail(listID uint64) (*models.List, error)
	// 获得最新的六个书单
	GetLatestSixLists() ([]*models.List, error)
	// 获得推荐的六个书单
	GetRecommendSixLists() ([]*models.List, error)
	// 获得最热的六个书单
	GetHeatSixLists() ([]*models.List, error)
	// 获得大咖推荐书单
	GetBigManRecommendLists() ([]*models.BigManRecommendList, error)
}
