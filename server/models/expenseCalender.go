package models

type ExpenseCalender struct {
	// user id
	UserId string `json:"user_id"`
	// order id
	OrderID string `json:"order_id"`
	// cost money
	Money float64 `json:"money"`

	BusinessType BusinessType `json:"business_type"`
	// order status 0-unpaid 1-paid
	Status StatusExpense `json:"status"`
	// order start time
	StartTime string `json:"start_time"`
	// order end time
	EndTime string `json:"end_time"`
}

type ListPurchaseRecord struct {
	// order id
	OrderId string `json:"order_id"`
	ListId  int    `json:"list_id"`
	PayTime string `json:"pay_time"`
}

type StatusExpense int

const (
	Unpaid StatusExpense = iota
	Paid
)

type BusinessType string

const (
	BigMan = "BIGMAN"
)
