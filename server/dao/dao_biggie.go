package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BiggieDao interface {
	FindBiggieById(id int) (*models.Biggie, error)
	FindLatestBiggie(pageNum, pageCount int) ([]*models.Biggie, error)
	FindRecommendBiggies(pageNum, pageCount int) ([]*models.Biggie, error)
	FindListsByBiggie(biggieId, pageNum, pageCount int) ([]*models.BiggieList, error)
	FindBiggieListBooks(listId int) ([]*models.BiggieBooks, error)
	FindLatestBiggieList(pageNum, pageCount int) ([]*models.BiggieList, error)

	AddCollectBiggie(collect *models.BiggieCollect) error
	DeleteCollectBiggie(userId string, biggieId int) error
	FindCollectBiggies(userId string) ([]*models.Biggie, error)
	FindBiggieIsCollected(userId string, biggieId int) (string, error)
	AddClickCount(listId int) error
}
