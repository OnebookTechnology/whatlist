package dao

import "github.com/OnebookTechnology/WhatList/server/models"

type UserDao interface {
	// 添加用户
	RegisterUser(user *models.User) (int64, error)
	// 更新用户
	UpdateUser(userId string, user *models.User) error
	// 查找用户
	FindUser(userId string) (*models.User, error)
}
