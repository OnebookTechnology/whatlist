package server

import (
	"bytes"
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

const PayReturnSuccess = "SUCCESS"

type PayResponse struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	Sign      string `json:"sign"`
}

type SessionInfo struct {
	OpenId string
	Type   ClientType
}

func Pay(ctx *gin.Context) {
	crossDomain(ctx)
	var err error
	userId := ctx.Query("user_id")
	businessType := ctx.Query("business_type")
	listId := ctx.Query("list_id")
	biggieId := ctx.Query("biggie_id")
	//WeChatLoginInfo

	body := ctx.Request.Body
	defer body.Close()
	dataBytes, err := ioutil.ReadAll(body)
	if err != nil && err != io.EOF {
		sendJsonResponse(ctx, ReadRequestErr, "read VerifyVCode api data err: %s", err.Error())
		return
	}

	req := new(WeChatPayRequest)
	err = jsoniter.Unmarshal(dataBytes, req)
	if err != nil {
		sendJsonResponse(ctx, ReadRequestErr, "unmarshal verify data err: %s data: %s", err.Error(), string(dataBytes))
		return
	}

	tradeNo := strconv.FormatInt(time.Now().Unix(), 10) + RandNumber(4)
	uri := "trade_no=" + tradeNo + "&business_type=" + businessType + "&fee=" + strconv.Itoa(req.TotalFee) +
		"&uid=" + userId + "&list_id=" + listId + "&bid=" + biggieId
	token := xxtea.EncryptStdToURLString(uri, server.XXTEAKey)

	req.AppId = server.AppId
	req.Body = "onebooktech"
	req.MchId = server.MchId
	req.NonceStr = time.Now().Format("20060102150405")
	req.NotifyUrl = "https://" + server.domain + "/whatlist/pay/callback/" + token
	req.OpenId = userId
	req.OutTradeNo = tradeNo
	req.SpbillCreateIP = strings.Split(ctx.Request.RemoteAddr, ":")[0]
	req.LimitPay = "no_credit"
	//req.TotalFee
	req.TradeType = "JSAPI"
	req.Sign = genSign(req)

	//encode to xml
	//payBody, err := xml.MarshalIndent(req, "", " ")
	payBody, err := xml.Marshal(req)
	if err != nil {
		logger.Error("marshal xml Pay data err:", err, "req:", req)
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

	//orderId, err := strconv.ParseUint(req.OutTradeNo, 10, 64)
	//if err != nil {
	//	sendJsonResponse(ctx, Err, "req.OutTradeNo ParseUint %s err: %s phone: %s", req.OutTradeNo, err.Error(), phoneStr)
	//	return
	//}

	expense := &models.ExpenseCalender{
		UserId:       userId,
		OrderID:      tradeNo,
		Money:        FenToYuan(req.TotalFee),
		Status:       models.Unpaid,
		StartTime:    nowFormat(),
		BusinessType: businessType,
	}

	err = server.DB.AddExpenseCalendar(expense)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when AddExpenseCalendar expense: %+v err: %s", expense, err.Error())
		return
	}
	// response the request
	res := &PayResponse{
		AppId:     payResp.AppId.Value,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  payResp.NonceStr.Value,
		Package:   "prepay_id=" + payResp.PrepayId.Value,
		SignType:  "MD5",
		Sign:      payResp.Sign.Value,
	}

	data, _ := jsoniter.MarshalToString(res)
	sendJsonResponse(ctx, OK, data)

}

//WeChat call back
func PayCallback(ctx *gin.Context) {
	token := ctx.Param("token")
	uriStr, err := xxtea.DecryptURLToStdString(token, server.XXTEAKey)
	if err != nil {
		logger.Error("PayCallback DecryptURLToStdString err. token:", token, "err:", err)
		sendJsonResponse(ctx, Err, "PayCallback DecryptURLToStdString err: %s", err.Error())
		return
	}
	paramsMap := getURIParams(uriStr)
	orderId := paramsMap["trade_no"]
	businessType := paramsMap["business_type"]
	feeStr := paramsMap["fee"]
	listIdStr := paramsMap["list_id"]
	userId := paramsMap["uid"]
	biggieIdStr := paramsMap["bid"]

	if len(orderId) == 0 || len(businessType) == 0 || len(feeStr) == 0 || len(userId) == 0 {
		logger.Error("PayCallback lack of params err. paramsMap:", paramsMap)
		sendJsonResponse(ctx, Err, "PayCallback lack of params err.")
		return
	}

	if listIdStr == "" {
		listIdStr = "0"
	}
	listId, err := strconv.ParseInt(listIdStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "PayCallback invalid listId: %s", listIdStr)
		return
	}

	if biggieIdStr == "" {
		biggieIdStr = "0"
	}
	biggieId, err := strconv.ParseInt(biggieIdStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "PayCallback invalid biggieIdStr: %s", biggieIdStr)
		return
	}

	feeInt, err := strconv.ParseInt(feeStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "PayCallback invalid fee: %s", feeStr)
		return
	}
	fee := FenToYuan(int(feeInt)) //to CHN ¥

	logger.Info("%%% PayCallBack OrderID :", orderId, "BusinessType:", businessType, "Fee:", fee, "%%%")

	//TODO: LOCK orderId

	calender, err := server.DB.FindExpenseCalendarByOrderId(orderId)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindExpenseCalendarByOrderId orderId: %s err: %s", orderId, err.Error())
		return
	}
	if calender.Status == models.Paid {
		sendJsonResponse(ctx, OK, "paid")
		return
	}
	err = server.DB.UpdateExpenseCalendar(userId, orderId, int(listId), int(biggieId), models.Paid, businessType)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when AfterPay orderId: %s err: %s", orderId, err.Error())
		return
	}

	//TODO: Update Status
	sendJsonResponse(ctx, OK, "ok")
}

func genSign(req *WeChatPayRequest) string {
	totalFee := strconv.Itoa(req.TotalFee)
	stringA := "appid=" + req.AppId +
		"&body=" + req.Body +
		"&limit_pay=" + req.LimitPay +
		"&mch_id=" + req.MchId +
		"&nonce_str=" + req.NonceStr +
		"&notify_url=" + req.NotifyUrl +
		"&openid=" + req.OpenId +
		"&out_trade_no=" + req.OutTradeNo +
		"&spbill_create_ip=" + req.SpbillCreateIP +
		"&total_fee=" + totalFee +
		"&trade_type=" + req.TradeType

	stringSignTemp := stringA + "&key=" + WeChatKey
	sign := strings.ToUpper(doMD5FromString(stringSignTemp))
	logger.Debug("stringSignTemp:", stringSignTemp, "sign:", sign)
	return sign
}
