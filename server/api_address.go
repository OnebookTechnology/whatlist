package server

import (
	"database/sql"
	"fmt"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	//"strconv"
	"github.com/json-iterator/go"
)

//func GetLatestBiggie(ctx *gin.Context) {
//	crossDomain(ctx)
//pageNumStr := ctx.Query("page_num")
//pageNum, err := strconv.Atoi(pageNumStr)
//if err != nil {
//sendFailedResponse(ctx, Err, "parse page_num error:", err, "page_num:", pageNumStr)
//return
//}
//pageCountStr := ctx.Query("page_count")
//pageCount, err := strconv.Atoi(pageCountStr)
//if err != nil {
//sendFailedResponse(ctx, Err, "parse page_count error:", err, "page_count:", pageCountStr)
//return
//}
//
//}

type AddressReq struct {
	UserId          string `json:"user_id" form:"user_id"`
	AddressId       int    `json:"address_id" form:"address_id"`
	ReceiverNumber  int    `json:"receiver_number"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverAddress string `json:"receiver_address"`
	IsDefault       int    `json:"is_default"`
	CreateTime      string `json:"create_time"`
	Page
}

func AddAddress(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.ShouldBindJSON(&req); err == nil {
		c := &models.UserAddressInfo{
			UserId:          req.UserId,
			ReceiverNumber:  req.ReceiverNumber,
			ReceiverName:    req.ReceiverName,
			ReceiverAddress: req.ReceiverAddress,
			CreateTime:      nowTimestampString(),
		}

		address, err := server.DB.FindDefaultAddressByUserId(c.UserId)
		if err != nil && err != sql.ErrNoRows {
			sendFailedResponse(ctx, Err, fmt.Sprintf("db error when FindDefaultAddressByPhoneNumber. error: %s. data: %s %d %s %s",
				err.Error(), req.UserId, req.ReceiverNumber, req.ReceiverName, req.ReceiverAddress))
			return
		}
		c.IsDefault = models.NotDefaultAddress
		if address == nil {
			c.IsDefault = models.DefaultAddress
		}
		err = server.DB.AddAddressInfo(c)
		if err != nil {
			sendFailedResponse(ctx, Err, "AddMallAddressInfo err:", err)
			return
		}
		sendSuccessResponse(ctx, nil)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func FindDefaultAddress(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		a, err := server.DB.FindDefaultAddressByUserId(req.UserId)
		if err != nil {
			if err == sql.ErrNoRows {
				sendFailedResponse(ctx, NoResultErr, "FindDefaultAddressByUserId err:", err)
				return
			}
			sendFailedResponse(ctx, Err, "FindDefaultAddressByUserId err:", err)
			return
		}
		res := &ResData{
			Address: a,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func FindAddress(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		a, err := server.DB.FindAddressById(req.AddressId)
		if err != nil {
			if err == sql.ErrNoRows {
				sendFailedResponse(ctx, NoResultErr, "FindAddressById err:", err)
				return
			}
			sendFailedResponse(ctx, Err, "FindDefaultAddressByUserId err:", err)
			return
		}
		res := &ResData{
			Address: a,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func UpdateAddressInfo(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.BindJSON(&req); err == nil {
		c := &models.UserAddressInfo{
			UserId:          req.UserId,
			AddressId:       req.AddressId,
			ReceiverNumber:  req.ReceiverNumber,
			ReceiverName:    req.ReceiverName,
			ReceiverAddress: req.ReceiverAddress,
			CreateTime:      nowTimestampString(),
		}

		err := server.DB.UpdateAddressInfo(c)
		if err != nil && err != sql.ErrNoRows {
			sendFailedResponse(ctx, Err, fmt.Sprintf("db error when UpdateAddressInfo. error: %s. data: %s %d %d %s %s",
				err.Error(), req.UserId, req.AddressId, req.ReceiverNumber, req.ReceiverName, req.ReceiverAddress))
			return
		}
		sendSuccessResponse(ctx, nil)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func DeleteAddressInfo(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.ShouldBindJSON(&req); err == nil {
		err := server.DB.DeleteAddressInfo(uint64(req.AddressId), req.UserId)
		if err != nil {
			sendFailedResponse(ctx, Err, fmt.Sprintf("db error when DeleteAddressInfo. error: %s. data: %s %d %d %s %s",
				err.Error(), req.UserId, req.AddressId, req.ReceiverNumber, req.ReceiverName, req.ReceiverAddress))
			return
		}
		sendSuccessResponse(ctx, nil)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func ListAllAddressInfo(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		a, err := server.DB.ListAllAddressInfoByUserId(req.UserId, req.PageNum, req.PageCount)
		if err != nil {
			if err == sql.ErrNoRows {
				sendFailedResponse(ctx, NoResultErr, "FindDefaultAddressByUserId err:", err)
				return
			}
			sendFailedResponse(ctx, Err, "FindDefaultAddressByUserId err:", err)
			return
		}
		res := &ResData{
			Addresses: a,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func UpdateMallAddressInfoToDefault(ctx *gin.Context) {
	crossDomain(ctx)
	var req AddressReq
	if err := ctx.BindJSON(&req); err == nil {
		err := server.DB.UpdateAddressInfoToDefaultByAddressId(req.UserId, uint64(req.AddressId))
		if err != nil && err != sql.ErrNoRows {
			sendFailedResponse(ctx, Err, fmt.Sprintf("db error when UpdateAddressInfo. error: %s. data: %s %d %d %s %s",
				err.Error(), req.UserId, req.AddressId, req.ReceiverNumber, req.ReceiverName, req.ReceiverAddress))
			return
		}
		sendSuccessResponse(ctx, nil)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}
