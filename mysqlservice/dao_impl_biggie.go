package mysql

import "github.com/OnebookTechnology/whatlist/server/models"

func (m *MysqlService) FindBiggieById(id int) (*models.Biggie, error) {
	row := m.Db.QueryRow("SELECT id,name,identity,intro,sendword,weight,signtime,image FROM biggie")

	b := new(models.Biggie)
	err := row.Scan(&b.Id, &b.Name, &b.Identity, &b.Intro, &b.Sendword, &b.Weight, &b.Signtime, &b.Image)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *MysqlService) FindLatestBiggie(limit int) ([]*models.Biggie, error) {
	var bs []*models.Biggie
	rows, err := m.Db.Query("SELECT id,name,identity,intro,sendword,weight,signtime,image FROM biggie ORDER BY signtime DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Biggie)
		err = rows.Scan(&b.Id, &b.Name, &b.Identity, &b.Intro, &b.Sendword, &b.Weight, &b.Signtime, &b.Image)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func (m *MysqlService) FindLatestBiggieList(pageNum, pageCount int) ([]*models.Biggie, error) {
	var bs []*models.Biggie
	rows, err := m.Db.Query("SELECT b.id,b.name,b.identity,b.intro,b.sendword,b.weight,b.signtime,b.image,l.`list_name` "+
		"FROM biggie b LEFT JOIN biggielist l ON b.latest_list_id=l.list_id ORDER BY b.`weight` DESC LIMIT ?,?",
		(pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := new(models.Biggie)
		l := new(models.BiggieList)
		err = rows.Scan(&b.Id, &b.Name, &b.Identity, &b.Intro, &b.Sendword, &b.Weight, &b.Signtime, &b.Image, &b.LatestListId, &l.ListName)
		if err != nil {
			return nil, err
		}
		b.Lists = append(b.Lists, l)
		bs = append(bs, b)
	}
	return bs, nil
}
