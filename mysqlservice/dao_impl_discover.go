package mysql

import (
	"database/sql"
	"github.com/OnebookTechnology/whatlist/server/models"
	"strings"
)

// 获得发现内容
func (m *MysqlService) GetDiscoverList(userId string, pageNum, pageCount int) ([]*models.Discover, error) {
	var discovers []*models.Discover
	rows, err := m.Db.Query("SELECT u.`nick_name`,u.avatar_url ,d.discover_id, d.`title` ,d.`subtitle` ,d.`picture`,d.`publish_time` ,d.`read_num` ,d.`like_num` "+
		" FROM `discover` d LEFT JOIN `user` u ON d.`user_id` = u.`user_id` "+
		" LIMIT ?,?", (pageNum-1)*pageCount, pageCount)
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
		if picStr != "" {
			d.Picture = strings.Split(picStr, ",")
			for i := range d.Picture {
				d.Picture[i] = d.Picture[i][1 : len(d.Picture[i])-1]
			}
		}

		err = m.GetDiscoverDetailIsThumb(d.DiscoverId, userId)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if err == sql.ErrNoRows {
			d.IsThumb = false
		} else {
			d.IsThumb = true
		}
		discovers = append(discovers, d)
	}
	return discovers, nil
}

// 获得发现内容
func (m *MysqlService) GetDiscoverDetail(id int) (*models.Discover, error) {
	row := m.Db.QueryRow("SELECT u.`nick_name`,u.avatar_url ,d.discover_id, d.`title` ,d.`subtitle`,d.content, d.`picture`,d.`publish_time` ,d.`read_num` ,d.`like_num` "+
		" FROM `discover` d LEFT JOIN `user` u ON d.`user_id` = u.`user_id` "+
		" WHERE d.discover_id = ?", id)

	d := new(models.Discover)
	var picStr string
	err := row.Scan(&d.NickName, &d.AvatarUrl, &d.DiscoverId, &d.Title, &d.Subtitle, &d.Content, &picStr, &d.PublishTime, &d.ReadNum, &d.LikeNum)
	if err != nil {
		return nil, err
	}
	picStr = picStr[1 : len(picStr)-1]
	if picStr != "" {
		d.Picture = strings.Split(picStr, ",")
		for i := range d.Picture {
			d.Picture[i] = d.Picture[i][1 : len(d.Picture[i])-1]
		}
	}
	return d, nil
}

func (m *MysqlService) AddReadNum(id int) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE discover SET read_num=read_num+1 WHERE discover_id = ?", id)
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

	_, err = tx.Exec("UPDATE discover SET like_num=like_num+1 WHERE discover_id = ?", id)
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

func (m *MysqlService) SubLikeNum(id int) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE discover SET like_num=like_num-1 WHERE discover_id = ?", id)
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

// 获得是否已经点赞
func (m *MysqlService) GetDiscoverDetailIsThumb(discoverId int, userId string) error {
	row := m.Db.QueryRow("SELECT user_id, discover_id from thumbs_up_record "+
		" WHERE user_id = ? AND discover_id = ?", userId, discoverId)

	var uid string
	var did int
	err := row.Scan(&uid, &did)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlService) ThumbsUpDiscover(discoverId int, userId string) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO thumbs_up_record(user_id,discover_id) values (?,?)", userId, discoverId)
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

func (m *MysqlService) CancelThumbsUpDiscover(discoverId int, userId string) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM thumbs_up_record "+
		" WHERE user_id = ? AND discover_id = ?", userId, discoverId)
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
