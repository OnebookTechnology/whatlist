package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"time"
)

// add expense record
func (m *MysqlService) AddExpenseCalendar(expense *models.ExpenseCalender) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO ExpenseCalender(order_id, user_id,money,business_type, status,start_time,end_time) VALUES(?,?,?,?,?,?,?)",
		expense.OrderID, expense.UserId, expense.Money, expense.BusinessType, expense.Status, expense.StartTime, expense.EndTime)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (m *MysqlService) UpdateExpenseCalendar(userId, orderId string, listId int, status models.StatusExpense) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	_, err = tx.Exec("UPDATE ExpenseCalender SET status=?, end_time=? WHERE order_id=?", status, nowStr, orderId)
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO list_purchase_record VALUES(?,?,?)", orderId, listId, nowStr)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// find all expense records
func (m *MysqlService) FindAllExpenseCalendar(userId string) ([]*models.ExpenseCalender, error) {
	rows, err := m.Db.Query("SELECT user_id,order_id,money,status,start_time,end_time, business_type FROM ExpenseCalender WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	var expenseRecords []*models.ExpenseCalender
	for rows.Next() {
		expenseRecord := new(models.ExpenseCalender)
		err = rows.Scan(&expenseRecord.UserId, &expenseRecord.OrderID, &expenseRecord.Money, &expenseRecord.Status, &expenseRecord.StartTime, &expenseRecord.EndTime, &expenseRecord.BusinessType)
		if err != nil {
			return nil, err
		}
		expenseRecords = append(expenseRecords, expenseRecord)
	}
	return expenseRecords, nil
}

// find expense records by OrderId
func (m *MysqlService) FindExpenseCalendarByOrderId(orderId string) (*models.ExpenseCalender, error) {
	row := m.Db.QueryRow("SELECT user_id,order_id,money,status,start_time,end_time,business_type FROM ExpenseCalender WHERE order_id=?", orderId)
	ec := new(models.ExpenseCalender)
	err := row.Scan(&ec.UserId, &ec.OrderID, &ec.Money, &ec.Status, &ec.StartTime, &ec.EndTime, &ec.BusinessType)
	if err != nil {
		return nil, err
	}
	return ec, nil
}

// find expense records by OrderId
func (m *MysqlService) FindListPurchaseRecord(userId string, listId int) (*models.ListPurchaseRecord, error) {
	row := m.Db.QueryRow("SELECT e.`order_id`"+
		"FROM `list_purchase_record` r LEFT JOIN `expensecalender` e ON r.`order_id` = e.`order_id` where e.`user_id`=? AND r.list_id=?",
		userId, listId)
	ec := new(models.ListPurchaseRecord)
	err := row.Scan(&ec.OrderId)
	if err != nil {
		return nil, err
	}
	return ec, nil
}
