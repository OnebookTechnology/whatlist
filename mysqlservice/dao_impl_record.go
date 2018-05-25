package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
)

func (m *MysqlService) FindListRecordByUserId(userId string, pageNum, pageCount int) ([]*models.BrowseListRecord, error) {
	var lists []*models.BrowseListRecord
	rows, err := m.Db.Query("SELECT l.`listID` ,l.`listName` ,l.`listAuthor` ,l.`listCategoryID` ,"+
		"l.`listIntro` ,l.`listImg`, l.`listCreateTime` ,l.`listLastEditTime` ,l.`listClickCount` "+
		"FROM `browse_list_record` r LEFT JOIN `list` l on l.`listID` =r.`list_id` "+
		"WHERE r.`user_id` = ? ORDER BY r.`browse_time` desc "+
		"LIMIT ?,?", userId, pageNum-1, pageCount)
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

func (m *MysqlService) FindBookRecordByUserId(userId string, pageNum, pageCount int) ([]*models.BrowseBookRecord, error) {
	var lists []*models.BrowseBookRecord
	rows, err := m.Db.Query("SELECT b.`ISBN` , b.`book_name` , b.`author_name` ,b.`press` ,b.`publication_time` ,b.`book_icon`,rr.`user_id` "+
		"FROM `browse_book_record` rr LEFT JOIN `book` b on rr.`ISBN` = b.`ISBN` WHERE rr.`user_id` = ? "+
		"ORDER BY rr.`browse_time` DESC "+
		"LIMIT ?,?", userId, pageNum-1, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.BrowseBookRecord)
		err = rows.Scan(&list.ISBN, &list.BookName, &list.AuthorName, &list.Press, &list.PublicationTime, &list.BookIcon, &list.UserId)
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
		"VALUES(?,?)", userID, listId)
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

func (m *MysqlService) AddBookRecord(userID string, isbn uint64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO `browse_book_record`(`user_id` ,`isbn` ) "+
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
