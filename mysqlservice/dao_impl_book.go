package mysql

import "github.com/OnebookTechnology/whatlist/server/models"

//ISBN       uint64 `json:"isbn,omitempty"`
//BookName   string `json:"book_name,omitempty"`
//AuthorName string `json:"author_name,omitempty"`
//// 出版社
//Press           string `json:"press,omitempty"`
//PublicationTime string `json:"publication_time,omitempty"`
//// 印刷时间
//PrintTime string `json:"print_time,omitempty"`
//// 版次
//Edition   uint8   `json:"edition,omitempty"`
//BookPrice float64 `json:"price,omitempty"`
//// 开本
//Format string `json:"format,omitempty"`
//// 纸张
//Paper string `json:"paper,omitempty"`
//// 包装
//Pack string `json:"pack,omitempty"`
//// 是否套装
//Suit uint8 `json:"suit,omitempty"`
//// 目录
//TableOfContent string `json:"table_of_content,omitempty"`
//// 图书一句话简介
//BookBriefIntro string `json:"book_brief_intro,omitempty"`
//// 作者信息
//AuthorIntro string `json:"author_intro,omitempty"`
//// 图书内容简介
//ContentIntro string `json:"content_intro,omitempty"`
//// 编辑推荐
//EditorRecommend string `json:"editor_recommend,omitempty"`
//// 一级分类
//FirstClassification uint16 `json:"first_classification,omitempty"`
//// 二级分类
//SecondClassification uint16 `json:"second_classification,omitempty"`
//// 书的总分
//TotalScore uint32 `json:"total_score"`
//// 书的评论次数
//CommentTimes uint32 `json:"comment_times"`
//
//BookIcon   string  `json:"book_icon"`
//BookPic    string  `json:"book_pic"`
//BookDetail string  `json:"book_detail"`

//Category   int     `json:"category"`
//Field1     []int   `json:"field_1,omitempty"` //年龄id范围
//Field2     int     `json:"field_2,omitempty"` //性别
//Field3     int     `json:"field_3,omitempty"` //婚姻状况id
//Field4     int     `json:"field_4,omitempty"` //教育程度
//Field5     []int   `json:"field_5,omitempty"` //收入id范围
//Field6     int     `json:"field_6,omitempty"` //工作行业id
//Field7     float64 `json:"field_7,omitempty"` //身高体重比例

func (m *MysqlService) FindAllBooks() (map[int][]*models.Book, error) {
	rows, err := m.Db.Query("SELECT ISBN,book_name,author_name,press,publication_time,print_time,format,paper,pack," +
		"suit,edition,table_of_content,book_brief_intro,author_intro,content_intro,editor_recommend,first_classification," +
		"second_classification,total_score,comment_times,book_icon,book_pic,book_detail,category," +
		"field1,field2,field3,field4,field5,field6,field7 from book")
	if err != nil {
		return nil, err
	}
	//bookMap := make(map[int][]*models.Book)
	for rows.Next() {
		book := new(models.Book)
		err = rows.Scan(&book.ISBN, &book.BookName, &book.AuthorName, &book.Press, &book.PublicationTime, &book.PrintTime, &book.Format, &book.Paper, &book.Pack,
			&book.Suit, &book.Edition, &book.TableOfContent, &book.BookBriefIntro, &book.AuthorIntro, &book.ContentIntro, &book.EditorRecommend, &book.FirstClassification,
			&book.SecondClassification, &book.TotalScore, &book.CommentTimes, &book.BookIcon, &book.BookPic, &book.BookDetail, &book.Category,
			&book.Field1, &book.Field2, &book.Field3, &book.Field4, &book.Field5, &book.Field6, &book.Field7)
		if err != nil {
			return nil, err
		}
	}
	return nil, err
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
