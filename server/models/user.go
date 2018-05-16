package models

type User struct {
	UserId uint64 `json:"user_id"`
	// phone number
	PhoneNumber uint64 `json:"phone_number"`
	// register time
	RegisterTime string `json:"register_time"`

	NickName            string  `json:"nick_name"`
	AvatarUrl           string  `json:"avatar_url"`
	Gender              int     `json:"gender"`
	City                string  `json:"city"`
	Province            string  `json:"province"`
	Country             string  `json:"country"`
	Language            string  `json:"language"`
	RegisterRank        int64   `json:"register_rank"`
	ReadCoin            uint64  `json:"read_coin"`
	BorrowLeft          uint8   `json:"borrow_left"`
	Vip                 uint8   `json:"vip"`
	Hobby               []int   `json:"hobby"`             //喜好id
	Field1              int     `json:"field_1,omitempty"` //年龄id
	Field2              int     `json:"field_2,omitempty"` //性别
	Field3              int     `json:"field_3,omitempty"` //婚姻状况id
	Field4              int     `json:"field_4,omitempty"` //教育程度
	Field5              int     `json:"field_5,omitempty"` //收入id
	Field6              int     `json:"field_6,omitempty"` //工作行业id
	Field7              float64 `json:"field_7,omitempty"` //身高体重比例
	NeedUpdateRecommend bool    //是否更新过数据
}
