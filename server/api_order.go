package server

import (
	"database/sql"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookOrderReq struct {
	UserId          string  `json:"user_id" form:"user_id"`
	BookISBNS       []int64 `json:"book_isbns"`
	OrderId         int64   `json:"order_id" form:"order_id"`
	AddressId       int     `json:"address_id"`
	OriginMoney     float64 `json:"origin_money"`
	Discount        float64 `json:"discount"`
	OrderMoney      float64 `json:"order_money"`
	OrderStatus     int     `json:"order_status"`
	TrackingNumber  string  `json:"tracking_number"`
	Freight         float64 `json:"freight"`
	Remark          string  `json:"remark"`
	OrderBeginTime  string  `json:"order_begin_time"`
	OrderUpdateTime string  `json:"order_update_time"`
	Page
}

func AddBookOrder(ctx *gin.Context) {
	crossDomain(ctx)
	var req BookOrderReq
	if err := ctx.BindJSON(&req); err == nil {
		orderId, err := strconv.ParseInt(nowTimestampString()+RandNumber(4), 10, 64)
		if err != nil {
			sendFailedResponse(ctx, Err, "Invalid orderId err:", err, "data:", nowTimestampString()+RandNumber(4))
			return
		}
		o := &models.BookOrder{
			UserId:          req.UserId,
			OrderId:         orderId,
			AddressId:       req.AddressId,
			OriginMoney:     req.OriginMoney,
			Discount:        req.Discount,
			OrderMoney:      req.OrderMoney,
			OrderStatus:     models.OrderReviewing,
			OrderBeginTime:  nowFormat(),
			OrderUpdateTime: nowFormat(),
			Remark:          req.Remark,
		}

		err = server.DB.AddBookOrder(o, req.BookISBNS)
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

func FindOrderDetail(ctx *gin.Context) {
	crossDomain(ctx)
	var req BookOrderReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		a, err := server.DB.FindOrderDetailByOrderId(req.OrderId)
		if err != nil {
			if err == sql.ErrNoRows {
				sendFailedResponse(ctx, NoResultErr, "FindOrderDetailByOrderId err:", err)
				return
			}
			sendFailedResponse(ctx, Err, "FindOrderDetailByOrderId err:", err)
			return
		}
		res := &ResData{
			OrderDetail: a,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func FindOrders(ctx *gin.Context) {
	crossDomain(ctx)
	var req BookOrderReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		a, err := server.DB.FindOrdersByUserId(req.UserId, req.PageNum, req.PageCount)
		if err != nil {
			if err == sql.ErrNoRows {
				sendFailedResponse(ctx, NoResultErr, "FindOrdersByUserId err:", err)
				return
			}
			sendFailedResponse(ctx, Err, "FindOrdersByUserId err:", err)
			return
		}
		res := &ResData{
			OrderDetails: a,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}
