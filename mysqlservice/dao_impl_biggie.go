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

func (m *MysqlService) FindBiggieIsCollected(userId string, biggieId int) (*models.BiggieCollect, error) {
	row := m.Db.QueryRow("SELECT collect_time FROM biggiecollect WHERE user_id=? AND biggie_id=?", userId, biggieId)
	var b = new(models.BiggieCollect)
	err := row.Scan(&b.CollectTime)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *MysqlService) FindListsByBiggie(biggieId, pageNum, pageCount int) ([]*models.BiggieList, error) {
	var bs []*models.BiggieList
	rows, err := m.Db.Query("SELECT list_id,biggie_id,list_name,list_intro,list_create_time,list_click_count,list_img, list_price FROM `biggielist` WHERE `biggie_id` = ? "+
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
	rows, err := m.Db.Query("SELECT list_id, biggie_id, list_name, list_click_count, list_img "+
		"FROM `biggielist` ORDER BY `list_create_time` DESC  LIMIT ?,?",
		(pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.BiggieList)
		err = rows.Scan(&b.ListId, &b.BiggieId, &b.ListName, &b.ListClickCount, &b.ListImg)
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
	_, err = tx.Exec("UPDATE biggie SET collect_count=collect_count+1")
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

func (m *MysqlService) FindCollectBiggies(userId string) ([]*models.Biggie, error) {
	var bs []*models.Biggie
	rows, err := m.Db.Query("SELECT b.id,b.name,b.identity,b.image,b.intro,b.collect_count,l.list_name FROM biggie b "+
		"LEFT JOIN biggiecollect bc ON b.id=bc.biggie_id "+
		"LEFT JOIN biggielist l ON b.latest_list_id=l.list_id "+
		"WHERE bc.user_id=? ORDER BY b.latest_list_id DESC", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Biggie)
		err = rows.Scan(&b.Id, &b.Name, &b.Identity, &b.Image, &b.Intro, &b.CollectCount)
		if err != nil {
			return nil, err
		}
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
