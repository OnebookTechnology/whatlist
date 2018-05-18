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
	_, err = tx.Exec("UPDATE User SET nick_name=?, avatar_url=?, gender=?, city=?, province=?, country=?, language=?,"+
		"hobby=?, field1=?, field2=?, field3=?, field4=?, field5=?, field6=?, field7=? WHERE user_id = ?",
		u.NickName, u.AvatarUrl, u.Gender, u.City, u.Province, u.Country, u.Language,
		u.Hobby, u.Field1, u.Field2, u.Field3, u.Field4, u.Field5, u.Field6, u.Field7, userId)
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
	row := m.Db.QueryRow("SELECT user_id, hobby FROM User WHERE user_id=?",
		userId)
	u := new(models.User)
	var hobbies string
	err := row.Scan(&u.UserId, &hobbies)
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
	if err != nil {
		return nil, err
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

	row := tx.QueryRow("SELECT @@IDENTITY")
	err = row.Scan(&lastId)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return lastId, nil
}
