package mysql

import "github.com/OnebookTechnology/whatlist/server/models"

// 添加图书订单
func (m *MysqlService) AddBookOrder(o *models.BookOrder, bookISBNS []int64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO bookorder(order_id, user_id, address_id, origin_money, discount, order_money,order_status, "+
		" order_begin_time, order_update_time) VALUES(?,?,?,?,?,?,?,?,?)",
		o.OrderId, o.UserId, o.AddressId, o.OriginMoney, o.Discount, o.OrderMoney, o.OrderStatus, o.OrderBeginTime, o.OrderBeginTime)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}

	for _, i := range bookISBNS {
		_, err = tx.Exec("INSERT INTO bookorderdetail(order_id, ISBN) VALUES(?,?)", o.OrderId, i)
		if err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				return rollBackErr
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	return nil
}

// 根据userId查询所有商城订单， 按发起时间排序, 分页
func (m *MysqlService) FindOrdersByUserId(userId string, pageNum, pageItems int) ([]*models.BookOrderDetail, error) {
	rows, err := m.Db.Query("SELECT b.`order_id` ,b.`order_money` ,b.`order_status`, b.`order_begin_time` ,"+
		"FROM `bookorder` b LEFT JOIN `useraddressinfo` u ON b.`address_id` = u.`address_id` "+
		"WHERE b.`user_id` = ? ORDER BY b.`order_begin_time` DESC LIMIT ?,?", userId, (pageNum-1)*pageItems, pageItems)
	if err != nil {
		return nil, err
	}
	var orders []*models.BookOrderDetail
	for rows.Next() {
		order := new(models.BookOrderDetail)
		err = rows.Scan(&order.OrderId, &order.OrderMoney, &order.OrderStatus, &order.OrderBeginTime)
		if err != nil {
			return nil, err
		}
		order.Books, err = m.FindBooksByOrderId(order.OrderId)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// 根据orderId查询所有订单图书内容
func (m *MysqlService) FindBooksByOrderId(orderId int) ([]*models.Book, error) {
	rows, err := m.Db.Query("SELECT b.`ISBN` ,b.`book_name`, b.`book_icon`, b.`price`   FROM `book` b "+
		"LEFT JOIN `bookorderdetail` d ON b.`ISBN` = d.`ISBN` WHERE d.`order_id` = ?", orderId)
	if err != nil {
		return nil, err
	}
	var books []*models.Book
	for rows.Next() {
		book := new(models.Book)
		err = rows.Scan(&book.ISBN, &book.BookName, &book.BookIcon, &book.BookPrice)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// 根据orderId查询所有订单内容
func (m *MysqlService) FindOrderDetailByOrderId(orderId int) (*models.BookOrderDetail, error) {
	row := m.Db.QueryRow("SELECT bo.`order_id` , bo.`origin_money` ,bo.`discount` ,bo.`order_money` ,bo.`order_status`,"+
		"bo.`tracking_number` ,bo.`freight` ,bo.`remark` ,bo.`order_begin_time`, bo.`order_update_time`,"+
		"a.`receiver_number` , a.`receiver_name` , a.`receiver_address` "+
		"FROM `bookorder` bo LEFT JOIN `useraddressinfo` a ON bo.`address_id` = a.`address_id` WHERE bo.`order_id`=?", orderId)

	order := new(models.BookOrderDetail)
	err := row.Scan(&order.OrderId, &order.OriginMoney, &order.Discount, &order.OrderMoney, &order.OrderStatus,
		&order.TrackingNumber, &order.Freight, &order.Remark, &order.OrderBeginTime, &order.OrderUpdateTime,
		&order.ReceiverNumber, &order.ReceiverName, &order.ReceiverAddress)
	if err != nil {
		return nil, err
	}
	order.Books, err = m.FindBooksByOrderId(order.OrderId)
	if err != nil {
		return nil, err
	}

	return order, nil
}
