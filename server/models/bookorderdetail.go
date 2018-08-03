package models

type BookOrderDetail struct {
	BookOrder
	UserAddressInfo
	Books []*Book
}
