package mysql

import "github.com/OnebookTechnology/whatlist/server/models"

func (m *MysqlService) AddBook(book *models.Book) error{
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO Book(ISBN,book_name,author_name,press,publication_time,print_time,edition," +
		"price,format,paper,pack,suit,book_brief_intro,author_intro,content_intro,editor_recommend," +
		"book_icon,book_pic,book_detail) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		book.ISBN,book.BookName,book.AuthorName,book.Press,book.PublicationTime,book.PrintTime,book.Edition,
		book.BookPrice,book.Format,book.Paper,book.Pack,book.Suit,book.BookBriefIntro,book.AuthorIntro,
		book.ContentIntro,book.EditorRecommend, book.BookIcon,book.BookPic,book.BookDetail)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}