package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type OrderDao interface {
	AddBookOrder(o *models.BookOrder, bookISBNS []int64) error
	FindOrdersByUserId(userId string, pageNum, pageItems int) ([]*models.BookOrderDetail, error)
	FindBooksByOrderId(orderId int) ([]*models.Book, error)
	FindOrderDetailByOrderId(orderId int) (*models.BookOrderDetail, error)
}
