package models

type BookOrderDetail struct {
	BookOrder
	Books []*Book
}
