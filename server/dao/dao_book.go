package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type BookDao interface {
	// 添加图书
	AddBook(book *models.Book) error
	//// 删除书
	//DeleteBook(ISBN uint64) error
	//// 更新书
	//UpdateBook(ISBN uint64, b *models.Book) error
	//// 修改书的评论次数
	//UpdateBookCommentTimes(ISBN, times uint64) error
	//// 修改书的分数
	//UpdateBookScore(ISBN uint64, score float64) error
	//// 按照ISBN查找一本书
	//FindBook(ISBN uint64) (*models.Book, error)
	//// 按照rfid查找一本书
	//FindRealPriceByRFID(rfid uint64) (float64, error)
	//// 按照价格范围检索图书
	//FindBookByPrice(lo, hi uint64) ([]*models.Book, error)
	// find all books
	// 分类书单
	FindBookByCateGory(categoryId, pageNum, pageCount int) ([]*models.Book, error)
	FindAllBooks() (map[int][]*models.Book, error)
	FindBook(isbn uint64) (*models.Book, error)
	/*
		task: 查询一本书是否被某用户喜欢
		auth: cx
		params: isbn图书唯一标识， userID用户唯一标识
		return: 返回记录唯一标识，0则为未找到结果
	*/
	IsBookInterested(isbn uint64, userID string) (uint64, error)
	//FindAllBooksByCatalogAndPriceOrderByOnlineTime(catalog, lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
	//FindAllBooksByCatalogAndPriceOrderBySales(catalog, lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
	//FindAllBooksByPriceOrderByOnlineTime(lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
	//FindAllBooksByPriceOrderBySales(lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)

	CalculatePrice(ISBNs []int64) (float64, error)
}
