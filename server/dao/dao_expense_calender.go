package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type ExpenseCalendar interface {
	// add expense record
	AddExpenseCalendar(expense *models.ExpenseCalender) error

	UpdateExpenseCalendar(userId, orderId string, listId, biggieId int, status models.StatusExpense, payType string) error
	// find one user all expense records
	FindAllExpenseCalendar(userId string) ([]*models.ExpenseCalender, error)
	// find by order id
	FindExpenseCalendarByOrderId(orderId string) (*models.ExpenseCalender, error)

	FindListPurchaseRecord(userId string, listId int) (*models.ListPurchaseRecord, error)
}
