package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BiggieDao interface {
	FindBiggieById(id int) (*models.Biggie, error)
	FindBiggieListById(listId int) (*models.BiggieList, error)
	FindLatestBiggie(pageNum, pageCount int) ([]*models.Biggie, error)
	FindRecommendBiggies(pageNum, pageCount int) ([]*models.Biggie, error)
	FindListsByBiggie(userId string, biggieId, pageNum, pageCount int) ([]*models.BiggieList, error)
	FindBiggieListBooks(listId int) ([]*models.BiggieBooks, error)
	FindLatestBiggieList(pageNum, pageCount int) ([]*models.BiggieList, error)

	AddCollectBiggie(collect *models.BiggieCollect) error
	DeleteCollectBiggie(userId string, biggieId int) error
	FindCollectBiggies(userId string, pageNum, pageCount int) ([]*models.Biggie, error)
	FindBiggieIsCollected(userId string, biggieId int) (string, error)
	AddClickCount(listId int) error
}
