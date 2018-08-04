package server

import (
	"database/sql"
	"fmt"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	//"strconv"
)

//func GetLatestBiggie(ctx *gin.Context) {
//	crossDomain(ctx)
//	pageNumStr := ctx.Query("page_num")
//	pageNum, err := strconv.Atoi(pageNumStr)
//	if err != nil {
//		sendFailedResponse(ctx, Err, "parse page_num error:", err, "page_num:", pageNumStr)
//		return
//	}
//	pageCountStr := ctx.Query("page_count")
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
			sendFailedResponse(ctx, Err, fmt.Sprintf("db error when FindDefaultAddressByPhoneNumber. error: %s. data: %d %d %s %s",
				err.Error(), req.UserId, req.ReceiverNumber, req.ReceiverName, req.ReceiverAddress))
			return
		}
		c.IsDefault = models.NotDefaultAddress
		if address == nil {
			c.IsDefault = models.DefaultAddress
		}
		err = server.DB.AddMallAddressInfo(c)
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
