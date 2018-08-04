package dao

import "github.com/OnebookTechnology/whatlist/server/models"

type AddressDao interface {
	AddAddressInfo(info *models.UserAddressInfo) error
	FindDefaultAddressByUserId(userId string) (*models.UserAddressInfo, error)
	ListAllAddressInfoByUserId(userId string, pageNum, pageCount int) ([]*models.UserAddressInfo, error)
	UpdateAddressInfo(info *models.UserAddressInfo) error
	DeleteAddressInfo(addressId uint64, userId string) error
	UpdateAddressInfoToDefaultByAddressId(userId string, addressId uint64) error
	FindAddressById(addressId int) (*models.UserAddressInfo, error)
}
