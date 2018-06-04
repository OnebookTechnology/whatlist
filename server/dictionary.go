package server

import "math"

const (
	OK = -iota
	Err
	EmptyErr
	SessionTimeoutErr
	SendSMSFailed
	SendSMSErr
	PutObjErr
	GetObjErr
	SendMQErr
	GetKeyErr
	GenCaptchaErr
	GetCaptchaCodeErr
	ReadRequestErr
	RequestPayInterfaceErr
	GetPayIdApiErr
	GetPayReturnCodeErr
	GetPayResultCodeErr
	FindCouponEmptyErr
	RepaymentFrequentErr
	OutOfDistanceErr
	OutOfMaxBalanceErr
	ExtendDayMaxBorrowTimesErr   //-21
	ExtendMonthMaxBorrowTimesErr //-22
	ApplyRepeatErr
	BalanceNotEnoughErr
	OutOfStockErr
	NoResultErr
)

type ClientType int

const (
	None ClientType = iota
	WeChat
)

const (
	VerifyCodePrefix = "/vcode/"
	CaptchaPrefix    = "/captcha/"
	SessionPrefix    = "/session/"
)

const (
	//年龄范围
	AgeAny = iota
	Age00
	Age95
	Age90
	Age85
	Age80
	Age70
)

var AgeMap = []string{"", "00后", "95后", "90后", "85后", "80后", "70后", "其他"}

const (
	//性别
	SexAny = iota
	SexMan
	SexWoman
)

var SexMap = []string{"", "男", "女"}

const (
	//婚姻状况
	MarriageAny         = iota
	MarriedWithChild    //已婚有子女
	MarriedWithoutChild //已婚无子女
	UnmarriedWithInLove //未婚恋爱中
	UnmarriedWithSingle //未婚单身中
)

var MarriageMap = []string{"", "已婚有子女", "已婚无子女", "未婚恋爱中", "未婚单身中"}

const (
	//教育程度
	EduAny = iota
	EduPrimary
	EduMiddle
	EduHigh
	EduCollege
	EduMaster
	EduDoctor
)

var EduMap = []string{"", "小学", "中学", "高中", "大学", "硕士", "博士"}

const (
	//最小收入
	IncomeAny = iota
	IncomeUnder5k
	Income5k_10k
	Income10k_15k
	Income15k_30k
	Income30k_50k
	IncomeBeyond50k
)

var IncomeMap = []string{"", "5000以下", "5000-10000", "10000-15000", "15000-30000", "30000-50000", "50000以上"}

const (
	//工作行业
	WorkAny = iota
	Work1
	Work2
	Work3
	Work4
	Work5
	Work6
	Work7
	Work8
	Work9
	Work10
	Work11
	Work12
	WorkOther
)

var WorkMap = []string{"", "农、林、牧、渔业", "制造业、建筑业", "电力、热力、燃气及水生产和供应业", "批发和零售业", "交通运输、仓储和邮政业",
	"信息传输、软件和信息技术服务业", "金融业", "房地产业", "商务服务业", "科学研究和技术服务业", "文化、体育和娱乐业", "公共管理、社会保障和社会组织", "其他"}

const (
	//身高体重比例
	WeightAny = iota
	UnderWeight
	NomalWeight
	GeneralObesity
	MildObesity
	ModerateObesity
	SevereObesity
	UnderWeightValue     = 18.5
	NormalWeightValue    = 24
	GeneralObesityValue  = 27
	MildObesityValue     = 30
	ModerateObesityValue = 35
	SevereObesityValue   = math.MaxInt8
)

var WeightMap = []string{"", "体重过轻", "正常体重", "轻微肥胖", "中度肥胖", "重度肥胖", "严重肥胖"}

const (
	//图书品类
	CategoryAny       = iota
	CategoryChild     //童书类
	CategoryEducation //教育类
	CategoryNovel     //小说文学类
	CategoryEconomic  //经济管理类
	CategorySuccess   //成功与励志类
	CategorySocial    //社科人文类
	CategoryLife      //生活类
	CategoryPhoto     //艺术摄影类
	CategoryScience   //科技类
	CategoryComputer  //计算机与互联网类

)

var CategoryMap = []string{"随机", "童书类", "教育类", "小说文学类", "经济管理类", "成功与励志类", "社科人文类", "生活类", "艺术摄影类", "科技类", "计算机与互联网类"}

var WeightRange = []float64{UnderWeightValue, NormalWeightValue, GeneralObesityValue, MildObesityValue,
	ModerateObesityValue, SevereObesityValue}
var CategoryArray = []int{CategoryAny, CategoryChild, CategoryEducation, CategoryNovel, CategoryEconomic, CategorySuccess,
	CategorySocial, CategoryLife, CategoryPhoto, CategoryScience, CategoryComputer}
