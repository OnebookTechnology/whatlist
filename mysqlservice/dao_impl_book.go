package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"strconv"
	"strings"
)

func (m *MysqlService) FindAllBooks() (map[int][]*models.Book, error) {
	rows, err := m.Db.Query("SELECT ISBN,book_name,author_name,press,publication_time,print_time,format,paper,pack," +
		"suit,edition,table_of_content,book_brief_intro,author_intro,content_intro,editor_recommend,first_classification," +
		"second_classification,total_score,comment_times,book_icon,book_pic,book_detail,category," +
		"field1,field2,field3,field4,field5,field6,field7 from book")
	if err != nil {
		return nil, err
	}
	bookMap := make(map[int][]*models.Book)
	for rows.Next() {
		book := new(models.Book)
		var age, income, marry, job, edu string
		err = rows.Scan(&book.ISBN, &book.BookName, &book.AuthorName, &book.Press, &book.PublicationTime, &book.PrintTime, &book.Format, &book.Paper, &book.Pack,
			&book.Suit, &book.Edition, &book.TableOfContent, &book.BookBriefIntro, &book.AuthorIntro, &book.ContentIntro, &book.EditorRecommend, &book.FirstClassification,
			&book.SecondClassification, &book.TotalScore, &book.CommentTimes, &book.BookIcon, &book.BookPic, &book.BookDetail, &book.Category,
			&age, &book.Field2, &marry, &edu, &income, &job, &book.Field7)
		if err != nil {
			return nil, err
		}
		ageArray := strings.Split(age, ",")
		for _, data := range ageArray {
			d, _ := strconv.Atoi(data)
			book.Field1 = append(book.Field1, d)
		}

		incomeArray := strings.Split(income, ",")
		for _, data := range incomeArray {
			d, _ := strconv.Atoi(data)
			book.Field5 = append(book.Field5, d)
		}

		marryArray := strings.Split(marry, ",")
		for _, data := range marryArray {
			d, _ := strconv.Atoi(data)
			book.Field3 = append(book.Field3, d)
		}

		jobArray := strings.Split(job, ",")
		for _, data := range jobArray {
			d, _ := strconv.Atoi(data)
			book.Field6 = append(book.Field6, d)
		}

		eduArray := strings.Split(edu, ",")
		for _, data := range eduArray {
			d, _ := strconv.Atoi(data)
			book.Field4 = append(book.Field4, d)
		}

		if books, ok := bookMap[book.Category]; ok {
			books = append(books, book)
			bookMap[book.Category] = books
		} else {
			books = append(books, book)
			bookMap[book.Category] = books
		}
	}
	return bookMap, err
}

func (m *MysqlService) FindBook(isbn uint64) (*models.Book, error) {
	row := m.Db.QueryRow("SELECT ISBN,book_name,author_name,press,publication_time,print_time,format,paper,pack,"+
		"suit,edition,table_of_content,book_brief_intro,author_intro,content_intro,editor_recommend,first_classification,"+
		"second_classification,total_score,comment_times,book_icon,book_pic,book_detail,category from book where ISBN=?", isbn)
	book := new(models.Book)
	err := row.Scan(&book.ISBN, &book.BookName, &book.AuthorName, &book.Press, &book.PublicationTime, &book.PrintTime, &book.Format, &book.Paper, &book.Pack,
		&book.Suit, &book.Edition, &book.TableOfContent, &book.BookBriefIntro, &book.AuthorIntro, &book.ContentIntro, &book.EditorRecommend, &book.FirstClassification,
		&book.SecondClassification, &book.TotalScore, &book.CommentTimes, &book.BookIcon, &book.BookPic, &book.BookDetail, &book.Category)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (m *MysqlService) CalculatePrice(ISBNs []int64) (float64, error) {
	var isbns string
	var sum float64
	for _, isbn := range ISBNs {
		isbns += strconv.FormatInt(isbn, 10) + ","
	}
	row := m.Db.QueryRow("SELECT SUM(`price`)  FROM `book` WHERE ISBN IN (?)", isbns[:len(isbns)-1])
	err := row.Scan(&sum)
	if err != nil {
		return 0.00, err
	}
	return sum, nil
}

func (m *MysqlService) FindBookByCateGory(categoryId, pageNum, pageCount int) ([]*models.Book, error) {
	var list []*models.Book
	rows, err := m.Db.Query("SELECT ISBN,book_name,author_name,press,publication_time,print_time,format,paper,pack,"+
		"suit,edition,table_of_content,book_brief_intro,author_intro,content_intro,editor_recommend,first_classification,"+
		"second_classification,total_score,comment_times,book_icon,book_pic,book_detail,category from book where category=? "+
		"LIMIT ?,?", categoryId, (pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := new(models.Book)
		err := rows.Scan(&book.ISBN, &book.BookName, &book.AuthorName, &book.Press, &book.PublicationTime, &book.PrintTime, &book.Format, &book.Paper, &book.Pack,
			&book.Suit, &book.Edition, &book.TableOfContent, &book.BookBriefIntro, &book.AuthorIntro, &book.ContentIntro, &book.EditorRecommend, &book.FirstClassification,
			&book.SecondClassification, &book.TotalScore, &book.CommentTimes, &book.BookIcon, &book.BookPic, &book.BookDetail, &book.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, book)
	}
	return list, nil
}

func (m *MysqlService) AddBook(book *models.Book) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO Book(ISBN,book_name,author_name,press,publication_time,print_time,edition,"+
		"price,format,paper,pack,suit,book_brief_intro,author_intro,content_intro,editor_recommend,"+
		"book_icon,book_pic,book_detail) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		book.ISBN, book.BookName, book.AuthorName, book.Press, book.PublicationTime, book.PrintTime, book.Edition,
		book.BookPrice, book.Format, book.Paper, book.Pack, book.Suit, book.BookBriefIntro, book.AuthorIntro,
		book.ContentIntro, book.EditorRecommend, book.BookIcon, book.BookPic, book.BookDetail)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

/*
task: 查询一本书是否被某用户喜欢
auth: cx
params: isbn图书唯一标识， userID用户唯一标识
return: 返回记录唯一标识，0则为未找到结果
*/
func (m *MysqlService) IsBookInterested(isbn uint64, userID string) (uint64, error) {
	var flag uint64 = 0
	row := m.Db.QueryRow("SELECT ul.`id` FROM `whatlist`.`userlike` ul "+
		"WHERE ul.`ISBN` = ? AND ul.`user_id` = ? LIMIT 1", isbn, userID)
	err := row.Scan(&flag)
	if err != nil {
		return 0, err
	}
	return flag, nil
}
