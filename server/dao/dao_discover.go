package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type DiscoverDao interface {
	GetDiscoverList(userId string, pageNum, pageCount int) ([]*models.Discover, error)
	GetDiscoverDetail(id int) (*models.Discover, error)
	AddReadNum(id int) error
	AddLikeNum(id int) error
	SubLikeNum(id int) error
	ThumbsUpDiscover(discoverId int, userId string) error
	CancelThumbsUpDiscover(discoverId int, userId string) error
	GetDiscoverDetailIsThumb(discoverId int, userId string) error
}
