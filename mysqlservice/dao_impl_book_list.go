package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
)

// 根据书单ID,获得指定书单的内容
func (m *MysqlService) GetList(listID uint64) (*models.List, error) {
	list := new(models.List)
	// 查询书单元信息
	row := m.Db.QueryRow("SELECT l.`listID` ,l.`listName` ,l.`listAuthor` ,l.`listCategoryID` , c.`categoryName`," +
	"l.`listIntro` ,l.`listImg`, l.`listCreateTime` ,l.`listLastEditTime` ,l.`listClickCount`" +
	" FROM whatlist.`list` l" +
	" LEFT JOIN `whatlist`.`category` c ON c.`categoryID` = l.`listCategoryID`" +
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
func (m *MysqlService) GetLatestSixList()([]*models.List, error) {
	var lists []*models.List
	rows, err := m.Db.Query("SELECT l.`listID` ,l.`listName` ,l.`listImg`, l.`listClickCount`" +
	" FROM `whatlist`.`list` l" +
	" ORDER BY l.`listCreateTime` DESC" +
	" LIMIT 6")
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		list := new(models.List)
		err = rows.Scan(&list.ListID, &list.ListName, &list.ListImg, &list.ListClickCount)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}
// 获得推荐的六个书单
func (m *MysqlService) GetRecommendSixList()(*[]models.List, error) {
	var lists *[]models.List
	return lists, nil
}