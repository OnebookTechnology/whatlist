package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
)

func (m *MysqlService) FindListRecordByUserId(userId string) ([]*models.BrowseListRecord, error) {
	var lists []*models.BrowseListRecord
	rows, err := m.Db.Query("SELECT l.`listID` ,l.`listName` ,l.`listAuthor` ,l.`listCategoryID` ,"+
		"l.`listIntro` ,l.`listImg`, l.`listCreateTime` ,l.`listLastEditTime` ,l.`listClickCount` "+
		"FROM `browse_list_record` r LEFT JOIN `list` l on l.`listID` =r.`list_id` "+
		"WHERE r.`user_id` = ? ORDER BY r.`browse_time` desc ", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.BrowseListRecord)
		err = rows.Scan(&list.ListID, &list.ListName, &list.ListAuthor, &list.ListCategoryID,
			&list.ListIntro, &list.ListImg, &list.ListCreateTime, &list.ListLastEditTime, &list.ListClickCount)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (m *MysqlService) AddListRecord(userID string, listId uint64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `browse_list_record`(`user_id` ,`list_id` ) "+
		"VALUES(?,?)", userID, isbn)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	return nil
}
