package _interface

import "github.com/OnebookTechnology/WhatList/server/dao"

type ServerDB interface {
	InitialDB(confPath string, tagName string) error
	dao.UserDao
}
