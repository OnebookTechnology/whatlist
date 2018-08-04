package mysql

import "github.com/OnebookTechnology/whatlist/server/models"

//添加地址
func (m *MysqlService) AddAddressInfo(info *models.UserAddressInfo) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO useraddressinfo(user_id, receiver_number, receiver_name, receiver_address, is_default, create_time) "+
		"VALUES(?,?,?,?,?,?)",
		info.UserId, info.ReceiverNumber, info.ReceiverName, info.ReceiverAddress, info.IsDefault, info.CreateTime)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	return nil
}

//查询默认地址
func (m *MysqlService) FindDefaultAddressByUserId(userId string) (*models.UserAddressInfo, error) {
	row := m.Db.QueryRow("SELECT address_id, user_id, receiver_number, receiver_name, receiver_address, is_default, create_time "+
		"FROM useraddressinfo "+
		"WHERE user_id=? AND is_default=?", userId, models.DefaultAddress)
	info := new(models.UserAddressInfo)
	err := row.Scan(&info.AddressId, &info.UserId, &info.ReceiverNumber, &info.ReceiverName, &info.ReceiverAddress, &info.IsDefault, &info.CreateTime)
	if err != nil {
		return nil, err
	}
	return info, nil
}

//查询默认地址
func (m *MysqlService) FindAddressById(addressId int) (*models.UserAddressInfo, error) {
	row := m.Db.QueryRow("SELECT address_id, user_id, receiver_number, receiver_name, receiver_address, is_default, create_time "+
		"FROM useraddressinfo WHERE address_id=?", addressId)
	info := new(models.UserAddressInfo)
	err := row.Scan(&info.AddressId, &info.UserId, &info.ReceiverNumber, &info.ReceiverName, &info.ReceiverAddress, &info.IsDefault, &info.CreateTime)
	if err != nil {
		return nil, err
	}
	return info, nil
}

//查询所有地址
func (m *MysqlService) ListAllAddressInfoByUserId(userId string, pageNum, pageCount int) ([]*models.UserAddressInfo, error) {
	rows, err := m.Db.Query("SELECT address_id, user_id, receiver_number, receiver_name, receiver_address, is_default, create_time "+
		"FROM useraddressinfo "+
		"WHERE user_id=? "+
		"ORDER BY is_default DESC, create_time asc LIMIT ?,? ", userId, (pageNum-1)*pageCount, pageCount)
	if err != nil {
		return nil, err
	}
	var infos []*models.UserAddressInfo
	for rows.Next() {
		info := new(models.UserAddressInfo)
		err = rows.Scan(&info.AddressId, &info.UserId, &info.ReceiverNumber, &info.ReceiverName, &info.ReceiverAddress, &info.IsDefault, &info.CreateTime)
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}
	return infos, nil
}

func (m *MysqlService) UpdateAddressInfo(info *models.UserAddressInfo) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE useraddressinfo SET receiver_number=?, receiver_name=?, receiver_address=?, create_time=? WHERE address_id=?",
		info.ReceiverNumber, info.ReceiverName, info.ReceiverAddress, info.CreateTime, info.AddressId)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlService) DeleteAddressInfo(addressId uint64, userId string) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM useraddressinfo WHERE address_id=? AND user_id=?",
		addressId, userId)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

//更新默认地址
func (m *MysqlService) UpdateAddressInfoToDefaultByAddressId(userId string, addressId uint64) error {
	tx, err := m.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE useraddressinfo SET is_default=? WHERE user_id=?",
		models.NotDefaultAddress, userId)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}

	_, err = tx.Exec("UPDATE useraddressinfo SET is_default=? WHERE address_id=? AND user_id=?",
		models.DefaultAddress, addressId, userId)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return rollBackErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
