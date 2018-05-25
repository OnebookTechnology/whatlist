package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type RecordDao interface {
	FindListRecordByUserId(userId string, pageNum, pageCount int) ([]*models.BrowseListRecord, error)
	FindBookRecordByUserId(userId string, pageNum, pageCount int) ([]*models.BrowseBookRecord, error)
	AddBookRecord(userID string, isbn uint64) error
	AddListRecord(userID string, listId uint64) error
}
