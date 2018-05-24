package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type PressDao interface {
	/**
	Task: 查询所有出版社
	Author: CX
	Return: 返回出版社信息
	*/
	GetPresses() ([]*models.Press, error)
	/*
		Task: 查询出版社推荐书单
		Author: CX
		Return: 返回出版社推荐书单
	*/
	GetPressRecommendLists(pressID uint64) ([]*models.List, error)
}
