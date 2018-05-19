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
	FindAllBooks() (map[int][]*models.Book, error)
	FindBook(isbn uint64) (*models.Book, error)
	//FindAllBooksByCatalogAndPriceOrderByOnlineTime(catalog, lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
	//FindAllBooksByCatalogAndPriceOrderBySales(catalog, lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
	//FindAllBooksByPriceOrderByOnlineTime(lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
	//FindAllBooksByPriceOrderBySales(lo, hi, pageNum, pageItems int) ([]*models.BookDisplay, error)
}
