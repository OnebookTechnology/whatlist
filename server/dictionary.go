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

const (
	//性别
	SexAny = iota
	SexMan
	SexWoman
)
const (
	//婚姻状况
	MarriageAny         = iota
	MarriedWithChild    //已婚有子女
	MarriedWithoutChild //已婚无子女
	UnmarriedWithInLove //未婚恋爱中
	UnmarriedWithSingle //未婚单身中
)
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
const (
	//工作行业
	WorkAny = iota
	Work1   //1、农、林、牧、渔业
	Work2   //2、制造业、建筑业
	Work3   //3、电力、热力、燃气及水生产和供应业
	Work4   //4、批发和零售业
	Work5   //5、交通运输、仓储和邮政业
	Work6   //6、信息传输、软件和信息技术服务业
	Work7   //7、金融业
	Work8   //8、房地产业
	Work9   //9、商务服务业
	Work10  //10、科学研究和技术服务业
	Work11  //11、文化、体育和娱乐业
	Work12  //12、公共管理、社会保障和社会组织
	WorkOther
)
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
const (
	//图书品类
	CategoryChild     = iota + 1 //童书类
	CategoryEducation            //教育类
	CategoryNovel                //小说文学类
	CategoryEconomic             //经济管理类
	CategorySuccess              //成功与励志类
	CategorySocial               //社科人文类
	CategoryLife                 //生活类
	CategoryPhoto                //艺术摄影类
	CategoryScience              //科技类
	CategoryComputer             //计算机与互联网类

)

var WeightRange = []float64{UnderWeightValue, NormalWeightValue, GeneralObesityValue, MildObesityValue,
	ModerateObesityValue, SevereObesityValue}
var CategoryArray = []int{CategoryChild, CategoryEducation, CategoryNovel, CategoryEconomic, CategorySuccess,
	CategorySocial, CategoryLife, CategoryPhoto, CategoryScience, CategoryComputer}
