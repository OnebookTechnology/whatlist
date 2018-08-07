package models

type BookOrder struct {
	OrderId         int64   `json:"order_id"`
	UserId          string  `json:"user_id"`
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
	WxAddress
}

type WxAddress struct {
	ErrMsg       string `json:"err_msg"`
	UserName     string `json:"user_name"`
	PostalCode   string `json:"postal_code"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	CountryName  string `json:"country_name"`
	DetailInfo   string `json:"detail_info"`
	NationalCode string `json:"national_code"`
	TelNumber    string `json:"tel_number"`
}

const (
	OrderWaitPay = iota
	OrderReviewing
	OrderHandled
	OrderSending
	OrderFinished
	OrderCanceled
)
