package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type ExpenseCalendar interface {
	// add expense record
	AddExpenseCalendar(expense *models.ExpenseCalender) error

	UpdateExpenseCalendar(orderId string, status models.StatusExpense) error
	// find one user all expense records
	FindAllExpenseCalendar(userId string) ([]*models.ExpenseCalender, error)
	// find by order id
	FindExpenseCalendarByOrderId(orderId string) (*models.ExpenseCalender, error)
}
