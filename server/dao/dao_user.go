package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type UserDao interface {
	// 添加用户
	RegisterUser(user *models.User) (int64, error)
	// 更新用户
	UpdateUser(userId string, user *models.User) error
	// 查找用户
	FindUser(userId string) (*models.User, error)
	/*
		task: 添加一本喜爱的图书
		author: cx
		params: userID 用户唯一标识符，isbn 图书唯一标识符
		return: error 如果没有错误返回nil
	*/
	AddInterestedBook(userID string, isbn uint64) error

	/*
		task: 删除一本喜爱的图书
		author: cx
		params: userID 用户唯一标识符，isbn 图书唯一标识符
		return: error 如果没有错误返回nil
	*/
	DeleteInterestedBook(userID string, isbn uint64) error

	/*
		task: 查询指定用户的所有喜爱图书
		author: cx
		params: userID 用户唯一表示符
		return:
	*/
	GetInterestedBooksByUserID(userID string) ([]*models.Book, error)
}
