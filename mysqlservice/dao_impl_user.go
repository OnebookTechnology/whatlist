package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"strconv"
	"strings"
)

// 添加用户
// NickName  string `json:"nick_name"`
// AvatarUrl string `json:"avatar_url"`
// Gender    string `json:"gender"`
// City      string `json:"city"`
// Province  string `json:"province"`
// Country   string `json:"country"`
// Language  string `json:"language"`

// 更新用户
func (m *MysqlService) UpdateUser(userId string, u *models.User) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	var hobbyStr string
	for i := range u.Hobby {
		hobbyStr = hobbyStr + strconv.Itoa(u.Hobby[i]) + ","
	}

	_, err = tx.Exec("UPDATE User SET hobby=?, field1=?, field2=?, field3=?, field4=?, field5=?, field6=?, field7=? WHERE user_id = ?",
		hobbyStr[:len(hobbyStr)-1], u.Field1, u.Field2, u.Field3, u.Field4, u.Field5, u.Field6, u.Field7, userId)
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

// 查找用户
func (m *MysqlService) FindUser(userId string) (*models.User, error) {
	row := m.Db.QueryRow("SELECT user_id, hobby,field1, field2, field3, field4, field5, field6, field7 FROM User WHERE user_id=?",
		userId)
	u := new(models.User)
	var hobbies string
	err := row.Scan(&u.UserId, &hobbies, &u.Field1, &u.Field2, &u.Field3, &u.Field4, &u.Field5, &u.Field6, &u.Field7)
	if err != nil {
		return nil, err
	}
	hobbyArray := strings.Split(hobbies, ",")
	if hobbyArray[0] == "" {
		u.Hobby = nil
	} else {
		for _, h := range hobbyArray {
			hi, _ := strconv.Atoi(h)
			u.Hobby = append(u.Hobby, hi)
		}
	}
	return u, nil
}

func (m *MysqlService) RegisterUser(user *models.User) (int64, error) {
	tx, err := m.Db.Begin()
	if err != nil {
		return 0, err
	}

	var lastId int64
	// s1. add user
	_, err = tx.Exec("INSERT INTO User(user_id) VALUES (?)", user.UserId)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	row := tx.QueryRow("select COUNT(*)  from `user`")
	err = row.Scan(&lastId)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

/*
task: 添加一本喜爱的图书
author: cx
params: userID 用户唯一标识符，isbn 图书唯一标识符
return: error 如果没有错误返回nil
*/
func (m *MysqlService) AddInterestedBook(userID string, isbn uint64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO "+
		"`whatlist`.`userlike`(`user_id` ,`ISBN` ) "+
		"VALUES(?,?)", userID, isbn)
	if err != nil {
		return err
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

/*
task: 删除一本喜爱的图书
author: cx
params: userID 用户唯一标识符，isbn 图书唯一标识符
*/
func (m *MysqlService) DeleteInterestedBook(userID string, isbn uint64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM `whatlist`.`userlike` WHERE `user_id` = ? AND `ISBN` = ?", userID, isbn)
	if err != nil {
		return err
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
	return nil
}

/*
task: 查询指定用户的所有喜爱图书
author: cx
params: userID 用户唯一表示符。
*/
func (m *MysqlService) GetInterestedBooksByUserID(userID string) ([]*models.Book, error) {
	var books []*models.Book
	rows, err := m.Db.Query("SELECT b.`book_name` , b.`book_icon`, b.`isbn` "+
		"FROM `whatlist`.`userlike` ul "+
		"LEFT JOIN `whatlist`.`book` b ON b.`ISBN` = ul.`ISBN` "+
		"WHERE ul.`user_id`  = ?", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Book)
		err = rows.Scan(&b.BookName, &b.BookIcon, &b.ISBN)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
