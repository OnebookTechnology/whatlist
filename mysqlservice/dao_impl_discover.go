package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/json-iterator/go"
	"strings"
)

// 获得发现内容
func (m *MysqlService) GetDiscoverList(pageNum, pageCount int) ([]*models.Discover, error) {
	var discovers []*models.Discover
	rows, err := m.Db.Query("SELECT u.`nick_name`,u.avatar_url ,d.id, d.`title` ,d.`subtitle` ,d.`picture`,d.`publish_time` ,d.`read_num` ,d.`like_num` "+
		" FROM `discover` d LEFT JOIN `user` u ON d.`user_id` = u.`user_id` "+
		" LIMIT ?,?", pageNum-1, (pageNum-1)*5)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var picStr string
		d := new(models.Discover)
		err = rows.Scan(&d.NickName, &d.AvatarUrl, &d.DiscoverId, &d.Title, &d.Subtitle, &picStr, &d.PublishTime, &d.ReadNum, &d.LikeNum)
		if err != nil {
			return nil, err
		}
		picStr = picStr[1 : len(picStr)-1]

		d.Picture = strings.Split(picStr, ",")
		discovers = append(discovers, d)
	}
	return discovers, nil
}

// 获得发现内容
func (m *MysqlService) GetDiscoverDetail(id int) (*models.Discover, error) {
	row := m.Db.QueryRow("SELECT u.`nick_name`,u.avatar_url ,d.id, d.`title` ,d.`subtitle`,d.content, d.`picture`,d.`publish_time` ,d.`read_num` ,d.`like_num` "+
		" FROM `discover` d LEFT JOIN `user` u ON d.`user_id` = u.`user_id` "+
		" WHERE d.id = ?", id)

	d := new(models.Discover)
	var picStr string
	err := row.Scan(&d.NickName, &d.AvatarUrl, &d.DiscoverId, &d.Title, &d.Subtitle, &d.Content, &picStr, &d.PublishTime, &d.ReadNum, &d.LikeNum)
	if err != nil {
		return nil, err
	}
	picStr = picStr[1 : len(picStr)-1]
	d.Picture = strings.Split(picStr, ",")
	return d, nil
}

func (m *MysqlService) AddReadNum(id int) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE discover SET read_num=read_num+1 WHERE id = ?", id)
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

func (m *MysqlService) AddLikeNum(id int) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE discover SET like_num=like_num+1 WHERE id = ?", id)
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
