package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
)

/**
Task: 查询所有出版社
Author: CX
Return: 返回出版社信息
*/
func (m *MysqlService) GetPresses() ([]*models.Press, error) {
	var presses []*models.Press
	rows, err := m.Db.Query("SELECT p.`id` ,p.`pressName` ,p.`pressImg` FROM `whatlist`.`press` p WHERE p.`isShowing` = 1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		press := new(models.Press)
		err = rows.Scan(&press.PressID, &press.PressName, &press.PressImg)
		if err != nil {
			return nil, err
		}
		presses = append(presses, press)
	}
	return presses, nil
}

/*
Task: 查询出版社推荐书单
Author: CX
Return: 返回出版社推荐书单
*/
func (m *MysqlService) GetPressRecommendLists(pressID uint64) ([]*models.List, error) {
	var lists []*models.List
	rows, err := m.Db.Query("SELECT l.`listID` ,l.`listName` ,l.`listBriefIntro` ,l.`listImg` ,l.`listClickCount` "+
		"FROM `whatlist`.`press` p "+
		"LEFT JOIN `whatlist`.`presslist` pl ON p.`id` = pl.`pressID` "+
		"LEFT JOIN `whatlist`.`list` l ON l.`listID` = pl.`listID` "+
		"WHERE pl.`isShowing` = 1 AND p.`id` = ?", pressID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list := new(models.List)
		err = rows.Scan(&list.ListID, &list.ListName, &list.ListBriefIntro, &list.ListImg, &list.ListClickCount)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}
