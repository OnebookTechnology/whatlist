package server

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/cxt90730/xxtea-go/xxtea"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BookOrderReq struct {
	UserId          string  `json:"user_id" form:"user_id"`
	BookISBNS       []int64 `json:"book_isbns"`
	OrderId         int64   `json:"order_id" form:"order_id"`
	ListId          int     `json:"list_id"`
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
		orderIdStr := nowTimestampString() + RandNumber(4)
		orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
		if err != nil {
			sendFailedResponse(ctx, Err, "Invalid orderId err:", err, "data:", nowTimestampString()+RandNumber(4))
			return
		}
		o := &models.BookOrder{
			UserId:          req.UserId,
			OrderId:         orderId,
			ListId:          req.ListId,
			OriginMoney:     req.OriginMoney,
			Discount:        req.Discount,
			OrderMoney:      req.OrderMoney,
			OrderStatus:     models.OrderWaitPay,
			OrderBeginTime:  nowFormat(),
			OrderUpdateTime: nowFormat(),
			Remark:          req.Remark,
		}

		s, _ := jsoniter.MarshalToString(o)
		uri := "s=" + s
		token := xxtea.EncryptStdToURLString(uri, server.XXTEAKey)

		wxReq := new(WeChatPayRequest)
		wxReq.AppId = server.AppId
		wxReq.Body = "购买图书"
		wxReq.MchId = server.MchId
		wxReq.NonceStr = time.Now().Format("20060102150405")
		wxReq.NotifyUrl = "https://" + server.domain + "/order/paycallback/" + token
		wxReq.OpenId = req.UserId
		wxReq.OutTradeNo = orderIdStr
		wxReq.SpbillCreateIP = strings.Split(ctx.Request.RemoteAddr, ":")[0]
		wxReq.LimitPay = "no_credit"
		wxReq.TotalFee = YuanToFen(req.OrderMoney)
		wxReq.TradeType = "JSAPI"
		wxReq.Sign = genSign(wxReq)

		//encode to xml
		payBody, err := xml.Marshal(wxReq)
		if err != nil {
			logger.Error("marshal xml Pay data err:", err, "req:", wxReq)
			sendJsonResponse(ctx, Err, "marshal xml Pay data err: %s", err.Error())
			return
		}

		logger.Debug(strings.Replace(string(payBody), "WeChatPayRequest", "xml", -1))
		//application/xml; charset=utf-8
		resp, err := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder",
			"application/xml; charset=utf-8",
			bytes.NewBuffer([]byte(strings.Replace(string(payBody), "WeChatPayRequest", "xml", -1))))

		if err != nil {
			sendJsonResponse(ctx, RequestPayInterfaceErr, "Pay request err: %s", err.Error())
			return
		}
		if resp.StatusCode != http.StatusOK {
			sendJsonResponse(ctx, RequestPayInterfaceErr, "Pay request code: %s", resp.StatusCode)
			return
		}
		repBody := resp.Body
		defer repBody.Close()
		payDataBytes, err := ioutil.ReadAll(repBody)
		if err != nil && err != io.EOF {
			sendJsonResponse(ctx, RequestPayInterfaceErr, "read Pay api data err: %s", err.Error())
			return
		}
		payResp := new(WeChatPayResponse)
		err = xml.Unmarshal(payDataBytes, payResp)
		if err != nil {
			logger.Error("unmarshal xml Pay data err:", err, "data:", string(payDataBytes))
			sendJsonResponse(ctx, GetPayIdApiErr, "unmarshal xml Pay data err: %s", err.Error())
			return
		}

		//ReturnCode cannot be FAIL
		if payResp.ReturnCode.Value != PayReturnSuccess {
			logger.Error(string(payDataBytes))
			sendJsonResponse(ctx, GetPayReturnCodeErr, "PayReturn Fail msg: %s", payResp.ReturnMsg)
			return
		}

		//and then ReturnCode cannot be FAIL
		if payResp.ResultCode.Value != PayReturnSuccess {
			logger.Error(payResp)
			sendJsonResponse(ctx, GetPayResultCodeErr, "PayResult Fail. code: %s msg: %s ", payResp.ErrCode, payResp.ErrCodeDes)
			return
		}

		expense := &models.ExpenseCalender{
			UserId:       req.UserId,
			OrderID:      orderIdStr,
			Money:        req.OrderMoney,
			Status:       models.Unpaid,
			StartTime:    nowFormat(),
			BusinessType: models.BiggieBook,
		}

		err = server.DB.AddExpenseCalendar(expense)
		if err != nil {
			sendJsonResponse(ctx, Err, "db error when AddExpenseCalendar expense: %+v err: %s", expense, err.Error())
			return
		}
		// response the request
		payRes := &PayResponse{
			AppId:     payResp.AppId.Value,
			TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
			NonceStr:  payResp.NonceStr.Value,
			Package:   "prepay_id=" + payResp.PrepayId.Value,
			SignType:  "MD5",
			Sign:      payResp.Sign.Value,
		}

		err = server.DB.AddBookOrder(o, req.BookISBNS)
		if err != nil {
			sendFailedResponse(ctx, Err, "AddMallAddressInfo err:", err)
			return
		}

		res := &ResData{
			PayResponse: payRes,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

//mall pay call back
func OrderPayCallback(ctx *gin.Context) {
	token := ctx.Param("token")
	uriStr, err := xxtea.DecryptURLToStdString(token, server.XXTEAKey)
	if err != nil {
		logger.Error("MallPayCallBack DecryptURLToStdString err. token:", token, "err:", err)
		sendJsonResponse(ctx, Err, "MallPayCallBack DecryptURLToStdString err: %s", err.Error())
		return
	}
	paramsMap := getURIParams(uriStr)
	info := paramsMap["s"]

	if info == "" {
		logger.Error("MallPayCallBack lack of params err. paramsMap:", paramsMap)
		sendJsonResponse(ctx, Err, "MallPayCallBack lack of params err.")
		return
	}

	order := new(models.BookOrder)
	jsoniter.UnmarshalFromString(info, order)

	logger.Info("%%% MallPayCallBack OrderID :", order.OrderId, "userId:", order.UserId, "Fee:", order.OrderMoney, "%%%")
	//TODO: LOCK orderId

	orderIdStr := strconv.FormatInt(order.OrderId, 10)
	calender, err := server.DB.FindExpenseCalendarByOrderId(orderIdStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindExpenseCalendarByOrderId orderId: %s err: %s", orderIdStr, err.Error())
		return
	}
	if calender.Status == models.Paid {
		sendJsonResponse(ctx, OK, "paid")
		return
	}

	e := &models.ExpenseCalender{
		UserId:  order.UserId,
		OrderID: orderIdStr,
		Status:  models.Paid,
	}
	//UpdateExpenseCalendar
	err = server.DB.UpdateExpenseCalendar(e)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when UpdateExpenseCalendar orderId: %s err: %s", orderIdStr, err.Error())
		return
	}

	//订单修改为审核中
	o := &models.BookOrder{
		UserId:          order.UserId,
		OrderId:         order.OrderId,
		OrderStatus:     models.OrderReviewing,
		OrderUpdateTime: nowFormat(),
	}
	// 直接兑换产品建立订单
	err = server.DB.UpdateBookOrder(o)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when UpdateBookOrder. error:", err)
		return
	}
	sendJsonResponse(ctx, OK, "ok")
}

func DeleteOrder(ctx *gin.Context) {
	crossDomain(ctx)
	var req BookOrderReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		o := &models.BookOrder{
			OrderId: req.OrderId,
			UserId:  req.UserId,
		}
		err := server.DB.DeleteBookOrder(o)
		if err != nil {
			sendFailedResponse(ctx, Err, "DeleteBookOrder err:", err)
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

func UpdateOrder(ctx *gin.Context) {
	crossDomain(ctx)
	var req BookOrderReq
	if err := ctx.ShouldBindJSON(&req); err == nil {
		o := &models.BookOrder{
			OrderStatus:     req.OrderStatus,
			OrderUpdateTime: nowFormat(),
			TrackingNumber:  req.TrackingNumber,
			Freight:         req.Freight,
			OrderId:         req.OrderId,
			UserId:          req.UserId,
		}
		err := server.DB.UpdateBookOrder(o)
		if err != nil {
			sendFailedResponse(ctx, Err, "UpdateBookOrder err:", err)
			return
		}
		sendSuccessResponse(ctx, nil)
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

func CalculateBookOrderPrice(ctx *gin.Context) {
	crossDomain(ctx)
	var req BookOrderReq
	if err := ctx.ShouldBindJSON(&req); err == nil {
		originMoney, err := server.DB.CalculatePrice(req.BookISBNS)
		if err != nil {

			sendFailedResponse(ctx, Err, "CalculatePrice err:", err)
			return
		}
		c := len(req.BookISBNS)
		if c == 0 {
			sendFailedResponse(ctx, Err, "invalid BookISBNS err:", err)
			return
		}
		discount := 1000 - 50*c
		if discount < 800 {
			discount = 800
		}

		a := new(models.BookOrderDetail)
		a.OriginMoney = originMoney
		a.Discount = FenToYuan(discount)
		a.OrderMoney = FenToYuan(YuanToFen(originMoney * a.Discount))

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
