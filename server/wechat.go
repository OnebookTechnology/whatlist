package server

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const WeChatKey = "onebooktechonebooktechonebooktec"

type WeChatLoginInfo struct {
	OpenId             string `json:"openid"`
	SessionKey         string `json:"session_key"`
	UnionId            string `json:"unionid"`
	ErrCode            int    `json:"errcode"`
	ErrMsg             string `json:"errmsg"`
	PhoneNumber        uint64 `json:"phone_number"`
	UpdateUTCTimestamp int64  `json:"timestamp"` //13 bit
}

type LoginInfo struct {
	WeChatLoginInfo
}

type CdataString struct {
	Value string `xml:",cdata"`
}

type WeChatRegisterInfo struct {
	ErrMsg        string      `json:"errMsg"`
	RawData       string      `json:"rawData"`
	UserInfo      *WeChatUser `json:"userInfo"`
	Signature     string      `json:"signature"`
	EncryptedData string      `json:"encryptedData"`
	IV            string      `json:"iv"`
}

type WeChatUser struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
}

type WeChatPayRequest struct {
	AppId          string `xml:"appid"`
	MchId          string `xml:"mch_id"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	SignType       string `xml:"sign_type,omitempty"`
	Body           string `xml:"body"`
	Detail         string `xml:"detail,omitempty"`
	Attach         string `xml:"attach,omitempty"`
	OutTradeNo     string `xml:"out_trade_no"`
	FeeType        string `xml:"fee_type,omitempty"`
	TotalFee       int    `xml:"total_fee" json:"total_fee"`
	SpbillCreateIP string `xml:"spbill_create_ip"`
	TimeStart      string `xml:"time_start,omitempty"`
	TimeExpire     string `xml:"time_expire,omitempty"`
	GoodsTag       string `xml:"goods_tag,omitempty"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`          // JSAPI
	LimitPay       string `xml:"limit_pay,omitempty"` // no_credit
	OpenId         string `xml:"openid,omitempty"`
}

type WeChatPayResponse struct {
	ReturnCode CdataString `xml:"return_code"`
	ReturnMsg  CdataString `xml:"return_msg"`
	AppId      CdataString `xml:"appid"`
	MchId      CdataString `xml:"mch_id"`
	DeviceInfo CdataString `xml:"device_info"`
	NonceStr   CdataString `xml:"nonce_str"`
	Sign       CdataString `xml:"sign"`
	ResultCode CdataString `xml:"result_code"`
	ErrCode    CdataString `xml:"err_code"`
	ErrCodeDes CdataString `xml:"err_code_des"`
	TradeType  CdataString `xml:"trade_type"`
	PrepayId   CdataString `xml:"prepay_id"`
}

//return third_key, err
func GetWeChatInfo(jsCode string) (*WeChatLoginInfo, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session"+
		"?appid=%s"+
		"&secret=%s"+
		"&js_code=%s"+
		"&grant_type=authorization_code",
		server.AppId, server.AppSecret, jsCode)
	logger.Debug("get Wechat session. url:", url, "app:", server.AppId, "secret:", server.AppSecret, "code:", jsCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}
	logger.Debug("request success! url:", url, "body:", string(body))
	wx := new(WeChatLoginInfo)
	err = jsoniter.Unmarshal(body, wx)
	if err != nil {
		return nil, err
	}
	if wx.ErrCode != 0 {
		return nil, errors.New(strconv.Itoa(wx.ErrCode) + ":" + wx.ErrMsg)
	}

	wx.UpdateUTCTimestamp = nowTimestampMs()
	return wx, nil

}

func genWeChatInfo(info *LoginInfo) (wx *WeChatLoginInfo) {
	wx = new(WeChatLoginInfo)
	wx.PhoneNumber = info.PhoneNumber
	wx.ErrCode = info.ErrCode
	wx.OpenId = info.OpenId
	wx.UpdateUTCTimestamp = info.UpdateUTCTimestamp
	wx.ErrMsg = info.ErrMsg
	wx.SessionKey = info.SessionKey
	wx.UnionId = info.UnionId
	return
}
