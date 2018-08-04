package models

type BookOrder struct {
	OrderId         int64   `json:"order_id"`
	UserId          string  `json:"user_id"`
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
}

const (
	OrderReviewing = iota
	OrderHandled
	OrderSending
	OrderFinished
	OrderCanceled
)
