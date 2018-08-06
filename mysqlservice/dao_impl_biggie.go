package mysql

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"time"
)

func (m *MysqlService) FindBiggieById(id int) (*models.Biggie, error) {
	row := m.Db.QueryRow("SELECT id,name,identity,intro,sendword,weight,signtime,image,collect_count FROM biggie WHERE id=?", id)

	b := new(models.Biggie)
	err := row.Scan(&b.Id, &b.Name, &b.Identity, &b.Intro, &b.Sendword, &b.Weight, &b.Signtime, &b.Image, &b.CollectCount)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *MysqlService) FindBiggieListById(listId int) (*models.BiggieList, error) {
	row := m.Db.QueryRow("SELECT l.list_name, b.name, l.list_intro, l.list_create_time, l.list_click_count, l.list_img, l.list_price "+
		"FROM biggielist l LEFT JOIN biggie b ON l.biggie_id=b.id WHERE l.list_id=?", listId)

	b := new(models.BiggieList)
	err := row.Scan(&b.ListName, &b.BiggieName, &b.ListIntro, &b.ListCreateTime, &b.ListClickCount, &b.ListImg, &b.ListPrice)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *MysqlService) FindBiggieIsCollected(userId string, biggieId int) (string, error) {
	row := m.Db.QueryRow("SELECT collect_time FROM biggiecollect WHERE user_id=? AND biggie_id=?", userId, biggieId)
	var b string
	err := row.Scan(&b)
	if err != nil {
		return "", err
	}
	return b, nil
}

func (m *MysqlService) FindListsByBiggie(userId string, biggieId, pageNum, pageCount int) ([]*models.BiggieList, error) {
	var bs []*models.BiggieList
	var lmap = make(map[int]string)
	//查询支付过的list
	rows2, err := m.Db.Query("SELECT l.list_id FROM `list_purchase_record` r LEFT JOIN `biggielist` l ON r.`list_id` = l.`list_id` "+
		"LEFT JOIN `expensecalender` e ON e.`order_id` = r.`order_id` "+
		"WHERE e.`status` = 1 AND e.`user_id` = ? AND l.`biggie_id` = ? ", userId, biggieId)
	if err != nil {
		return nil, err
	}
	for rows2.Next() {
		var lid int
		err = rows2.Scan(&lid)
		if err != nil {
			return nil, err
		}
		//所有已支付记录
		lmap[lid] = userId
	}

	rows, err := m.Db.Query("SELECT l.list_id,l.biggie_id,l.list_name,l.list_intro,l.list_create_time,l.list_click_count,l.list_img, list_price FROM `biggielist` l WHERE `biggie_id` = ? "+
		"ORDER BY list_create_time DESC LIMIT ?,?",
		biggieId, (pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.BiggieList)
		err = rows.Scan(&b.ListId, &b.BiggieId, &b.ListName, &b.ListIntro, &b.ListCreateTime, &b.ListClickCount, &b.ListImg, &b.ListPrice)
		if err != nil {
			return nil, err
		}
		if lmap[b.ListId] != "" {
			b.IsPayed = true
		}
		bs = append(bs, b)

	}

	return bs, nil
}

func (m *MysqlService) FindLatestBiggie(pageNum, pageCount int) ([]*models.Biggie, error) {
	var bs []*models.Biggie
	rows, err := m.Db.Query("SELECT id,name,identity,intro,sendword,weight,signtime,image,collect_count FROM biggie ORDER BY signtime DESC LIMIT ?,?",
		(pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Biggie)
		err = rows.Scan(&b.Id, &b.Name, &b.Identity, &b.Intro, &b.Sendword, &b.Weight, &b.Signtime, &b.Image, &b.CollectCount)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func (m *MysqlService) FindRecommendBiggies(pageNum, pageCount int) ([]*models.Biggie, error) {
	var bs []*models.Biggie
	rows, err := m.Db.Query("SELECT b.id,b.name,b.identity,b.intro,b.sendword,b.weight,b.signtime,b.image,b.latest_list_id,l.`list_name`, b.collect_count "+
		"FROM biggie b LEFT JOIN biggielist l ON b.latest_list_id=l.list_id ORDER BY b.`weight` DESC LIMIT ?,?",
		(pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Biggie)
		l := new(models.BiggieList)
		err = rows.Scan(&b.Id, &b.Name, &b.Identity, &b.Intro, &b.Sendword, &b.Weight, &b.Signtime, &b.Image, &b.LatestListId, &l.ListName, &b.CollectCount)
		if err != nil {
			return nil, err
		}
		b.Lists = append(b.Lists, l)
		bs = append(bs, b)
	}
	return bs, nil
}

func (m *MysqlService) FindBiggieListBooks(listId int) ([]*models.BiggieBooks, error) {
	var bs []*models.BiggieBooks
	rows, err := m.Db.Query("SELECT bb.`list_id` , bb.`ISBN` , bb.`recommend` , b.`book_name`, b.`author_name` , b.`book_icon` "+
		"FROM `biggiebooks` bb LEFT JOIN `book` b ON bb.`ISBN` = b.`ISBN` WHERE bb.`list_id` = ?", listId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.BiggieBooks)
		err = rows.Scan(&b.ListId, &b.ISBN, &b.Recommend, &b.BookName, &b.AuthorName, &b.BookIcon)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func (m *MysqlService) FindLatestBiggieList(pageNum, pageCount int) ([]*models.BiggieList, error) {
	var bs []*models.BiggieList
	rows, err := m.Db.Query("SELECT l.list_id, l.biggie_id, l.list_name, l.list_click_count, l.list_img, l.list_intro, l.list_price, b.name "+
		"FROM `biggielist` l LEFT JOIN biggie b ON b.id= l.biggie_id ORDER BY l.`list_create_time` DESC  LIMIT ?,?",
		(pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.BiggieList)
		err = rows.Scan(&b.ListId, &b.BiggieId, &b.ListName, &b.ListClickCount, &b.ListImg, &b.ListIntro, &b.ListPrice, &b.BiggieName)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func (m *MysqlService) AddCollectBiggie(c *models.BiggieCollect) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO biggiecollect(user_id, biggie_id, collect_time) VALUES(?,?,?)", c.UserId, c.BiggieId, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	_, err = tx.Exec("UPDATE biggie SET collect_count=collect_count+1 WHERE id=?", c.BiggieId)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (m *MysqlService) FindCollectBiggies(userId string, pageNum, pageCount int) ([]*models.Biggie, error) {
	var bs []*models.Biggie
	rows, err := m.Db.Query("SELECT b.id,b.name,b.identity,b.image,b.intro,b.collect_count,l.list_name FROM biggie b "+
		"LEFT JOIN biggiecollect bc ON b.id=bc.biggie_id "+
		"LEFT JOIN biggielist l ON b.latest_list_id=l.list_id "+
		"WHERE bc.user_id=? ORDER BY b.latest_list_id DESC LIMIT ?,?", userId, (pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Biggie)
		l := new(models.BiggieList)
		err = rows.Scan(&b.Id, &b.Name, &b.Identity, &b.Image, &b.Intro, &b.CollectCount, &l.ListName)
		if err != nil {
			return nil, err
		}
		b.Lists = append(b.Lists, l)
		bs = append(bs, b)
	}
	return bs, nil
}

func (m *MysqlService) DeleteCollectBiggie(userId string, biggieId int) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM biggiecollect WHERE user_id=? and biggie_id=?", userId, biggieId)
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

func (m *MysqlService) AddClickCount(listId int) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE biggielist SET list_click_count=list_click_count+1 WHERE list_id=?", listId)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (m *MysqlService) AddBiggieListRecord(orderId string, listId int) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO list_purchase_record VALUES(?,?,?)", orderId, listId, nowStr)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
