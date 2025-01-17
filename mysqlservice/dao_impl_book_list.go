package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
)

// 根据书单ID,获得指定书单的内容
func (m *MysqlService) GetListDetail(listID uint64) (*models.List, error) {
	list := new(models.List)
	// 查询书单元信息
	row := m.Db.QueryRow("SELECT l.`listID` ,l.`listName` ,l.`listAuthor` ,l.`listCategoryID` , c.`categoryName`,"+
		"l.`listIntro` ,l.`listImg`, l.`listCreateTime` ,l.`listLastEditTime` ,l.`listClickCount`"+
		" FROM whatlist.`list` l"+
		" LEFT JOIN `whatlist`.`category` c ON c.`categoryID` = l.`listCategoryID`"+
		" WHERE l.`listID` = ?", listID)
	err := row.Scan(&list.ListID, &list.ListName, &list.ListAuthor, &list.ListCategoryID, &list.ListCategoryName,
		&list.ListIntro, &list.ListImg, &list.ListCreateTime, &list.ListLastEditTime, &list.ListClickCount)
	if err != nil {
		return nil, err
	}
	var books []*models.Book
	// 查询书单包含图书的元信息
	rows, err := m.Db.Query("SELECT b.ISBN,b.book_name,b.author_name,b.press,b.publication_time,b.print_time,b.format,b.paper,b.pack,"+
		" b.suit,b.edition,b.table_of_content,b.book_brief_intro,b.author_intro,b.content_intro,b.editor_recommend,b.first_classification,"+
		" b.second_classification,b.total_score,b.comment_times,b.book_icon,b.book_pic,b.book_detail,b.category "+
		" FROM `whatlist`.`booklist` bl"+
		" LEFT JOIN `whatlist`.`book` b ON b.`ISBN` = bl.`ISBN`"+
		" LEFT JOIN `whatlist`.`list` l ON l.`listID` = bl.`listID`"+
		" WHERE l.`listID`  = ?", listID)
	for rows.Next() {
		book := new(models.Book)
		err = rows.Scan(&book.ISBN, &book.BookName, &book.AuthorName, &book.Press, &book.PublicationTime, &book.PrintTime, &book.Format, &book.Paper, &book.Pack,
			&book.Suit, &book.Edition, &book.TableOfContent, &book.BookBriefIntro, &book.AuthorIntro, &book.ContentIntro, &book.EditorRecommend, &book.FirstClassification,
			&book.SecondClassification, &book.TotalScore, &book.CommentTimes, &book.BookIcon, &book.BookPic, &book.BookDetail, &book.Category)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	list.ListBooks = books
	return list, nil
}

// 获得最新的六个书单
func (m *MysqlService) GetLatestSixLists(index uint64) ([]*models.List, error) {
	var lists []*models.List
	rows, err := m.Db.Query("SELECT l.`listID` ,l.`listName` ,l.`listImg`, l.`listClickCount`, l.`listBriefIntro` "+
		" FROM `whatlist`.`list` l"+
		" ORDER BY l.`listCreateTime` DESC"+
		" LIMIT ?,6", index*6)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.List)
		err = rows.Scan(&list.ListID, &list.ListName, &list.ListImg, &list.ListClickCount, &list.ListBriefIntro)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

// 获得推荐的六个书单
func (m *MysqlService) GetRecommendSixLists(index uint64) ([]*models.List, error) {
	var lists []*models.List
	rows, err := m.Db.Query("SELECT l.`listID` , l.`listName` ,l.`listImg` ,l.`listClickCount`, l.`listBriefIntro`  "+
		"FROM `whatlist`.`recommendlist` r "+
		"LEFT JOIN `whatlist`.`list` l ON r.`listID` = l.`listID` "+
		"WHERE r.`isRecommending` = 1 LIMIT ?,6;", index*6)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.List)
		err = rows.Scan(&list.ListID, &list.ListName, &list.ListImg, &list.ListClickCount, &list.ListBriefIntro)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

// 获得最热的六个书单
func (m *MysqlService) GetHeatSixLists() ([]*models.List, error) {
	var lists []*models.List
	rows, err := m.Db.Query("SELECT l.`listID` ,l.`listName` ,l.`listImg`, l.`listClickCount`, l.`listBriefIntro` " +
		" FROM `whatlist`.`list` l" +
		" ORDER BY l.`listClickCount` DESC" +
		" LIMIT 6")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.List)
		err = rows.Scan(&list.ListID, &list.ListName, &list.ListImg, &list.ListClickCount, &list.ListBriefIntro)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

// 获得大咖推荐书单
func (m *MysqlService) GetBigManRecommendLists() ([]*models.BigManRecommendList, error) {
	var lists []*models.BigManRecommendList
	rows, err := m.Db.Query("SELECT b.`id` ,b.`imgUrl`, b.`listID`, b.`price` " +
		"FROM `whatlist`.`bigmanrecommend` b " +
		"WHERE b.`isRecommending` = 1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.BigManRecommendList)
		err = rows.Scan(&list.ID, &list.ImgUrl, &list.ListID, &list.Price)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

// 获得每日推荐书单
func (m *MysqlService) GetEveryDayRecommendList(index uint64) (*models.EveryDayRecommend, error) {
	row := m.Db.QueryRow("SELECT e.`id` ,e.`recommendTime`,e.`imgUrl` ,e.`listID` , l.`listName` , l.`listBriefIntro` "+
		"FROM `whatlist`.`everydayrecommend` e "+
		"LEFT JOIN `whatlist`.`list` l ON l.`listID` = e.`listID` "+
		"ORDER BY e.`recommendTime` DESC LIMIT ?,1", index)
	everyDayRecommend := new(models.EveryDayRecommend)
	err := row.Scan(&everyDayRecommend.ID, &everyDayRecommend.RecommendTime, &everyDayRecommend.ImgUrl,
		&everyDayRecommend.ListID, &everyDayRecommend.ListName, &everyDayRecommend.ListBriefIntro)
	if err != nil {
		return nil, err
	}
	return everyDayRecommend, nil
}

// 获得轮播图
func (m *MysqlService) GetCarousel() ([]*models.Carousel, error) {
	var carousels []*models.Carousel
	rows, err := m.Db.Query("SELECT c.`id` ,c.`imgUrl` ,c.`ISBN`  " +
		"FROM `whatlist`.`carousel` c " +
		"WHERE c.`isShowing` = 1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := new(models.Carousel)
		err = rows.Scan(&c.ID, &c.ImgUrl, &c.ISBN)
		if err != nil {
			return nil, err
		}
		carousels = append(carousels, c)
	}
	return carousels, nil
}

/*
task：添加书单浏览量
author：陈曦
params：listID书单唯一标识符
*/
func (m *MysqlService) AddListClickCount(listID uint64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE `whatlist`.`list` l SET l.`listClickCount`  = l.`listClickCount`  + 1 "+
		"WHERE l.`listID` = ?", listID)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
