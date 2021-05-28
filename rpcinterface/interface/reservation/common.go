package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//通用字段
	OTHER = 1 //其他  通用字段

	//健康状况
	HEALTH_FEVER  = 2 //发热
	HEALTH_COUGH  = 3 //咳嗽
	HEALTH_RESP   = 4 //呼吸道感染
	HEALTH_HEALTH = 0 //健康

	//疫情状况
	EPI_LIFE    = 2 //疫情发源地生活史
	EPI_TRAVEL  = 3 //疫情发源地旅游史
	EPI_CONTACT = 4 //疫情发源地接触史
	EPI_NO      = 0 //无疫情状况

	//交通方式
	TRAFFIC_BUS   = 2 //公交车
	TRAFFIC_METRO = 3 //地铁
	TRAFFIC_WALK  = 4 //步行
	TRAFFIC_TAXI  = 5 //出租车
	TRAFFIC_DRIVE = 6 //驾车

	//出行地点
	ADDRESS_COMPANY = 2 //公司
	ADDRESS_HOME    = 3 //家

	//预约设置状态
	Resvervation_Setting_Disable = 0 // 关闭
	Resvervation_Setting_Enable  = 1 // 开启
)

var HealthMap = map[int]string{
	//健康状况
	HEALTH_FEVER:  "发热",
	HEALTH_COUGH:  "咳嗽",
	HEALTH_RESP:   "呼吸道感染",
	HEALTH_HEALTH: "健康",
	OTHER:         "其他",
}

var EpiMap = map[int]string{
	//疫情状况
	EPI_LIFE:    "疫情发源地生活史",
	EPI_TRAVEL:  "疫情发源地旅游史",
	EPI_CONTACT: "疫情发源地接触史",
	EPI_NO:      "无疫情状况",
	OTHER:       "其他",
}

var TraMap = map[int]string{
	//交通方式
	TRAFFIC_BUS:   "公交车",
	TRAFFIC_METRO: "地铁",
	TRAFFIC_WALK:  "步行",
	TRAFFIC_TAXI:  "出租车",
	TRAFFIC_DRIVE: "驾车",
	OTHER:         "其他",
}

var AddMap = map[int]string{
	//出行地点
	ADDRESS_COMPANY: "公司",
	ADDRESS_HOME:    "家",
	OTHER:           "其他",
}

//客户健康信息新增给入参
type ArgsHealthAdd struct {
	common.BsToken
	common.Utoken
	ShopId   int
	Health   string //健康状况 1代表发热 2代表咳嗽 3代表呼吸道 多个选项用,号隔开
	Epidemic string //疫情状况 1 有无疫情发源地生活史 2旅游史 3接触史 多个选项用,号隔开
	Traffic  int    //交通方式  1公交 2地铁 3步行 4出租车 5驾车 6其他
	GTime    string //出行时间
	GAddress int    //出行地点 1公司 2家 3其他
	ReservationId int //预约id
}

//客户健康信息更新入参
type ArgsHealthUpdate struct {
	common.BsToken
	Id       int
	CId      int    //客户id
	CName    string //客户名
	CPhone   int    //客户手机号
	RTime    int    //预约时间
	OrderNum string //单号
}

//查询客户健康信息入参
type ArgsHealthGet struct {
	common.Paging
	common.BsToken
	common.Utoken
	ShopId int
	RTime  string
	Data   string //全数字就匹配 查手机号或者单号 Y开头就查单号 两者都不是就查名字
	Type int //true用户  false Saas
}

//客户健康信息返回
type ReplyHealth struct {
	TotalNum int            //总数量
	List     []HealthInfo //客户健康信息
}
type HealthInfo struct {
	EsReservationHealth
	//CId          int         //客户id
	//CName        string   //客户名
	//CPhone       string    //客户手机号
	//ShopId       int      //门店id
	//ShopName     string
	//BranchName   string
	//RTime        string    //预约时间`mapstructure:"order_num"` //单号
	//OrderNum	 string
	//RId          int
	IsFever      bool   //是否发热
	IsCough      bool   //是否咳嗽
	IsResp       bool   //是否有呼吸道症状
	IsEpiLife    bool   //是否有疫情地区生活史
	IsEpiTravel  bool   //是否有疫情地区旅游史
	IsEpiContact bool   //是否有疫情地区接触史
	Traffic  string       //交通方式  1公交 2地铁 3步行 4出租车 5驾车 6其他
	GTime    string     //出行时间
	GAddress string     //出行地点 1公司 2家 3其他
	//AId 		int
	AIsFever      bool   //是否发热
	AIsCough      bool   //是否咳嗽
	AIsResp       bool   //是否有呼吸道症状
	//LId			int
	LTraffic  string      //交通方式  1公交 2地铁 3步行 4出租车 5驾车 6其他
	LGTime    string   //出行时间
	LGAddress string     //出行地点 1公司 2家 3其他
}

//根据服务id和手艺人id查询时间
type ArgsTimeGet struct {
	ShopId       int    //门店id
	Cids         []int  //手艺人id
	TimeInterval int    //时间间隔 分钟    默认是 15
	TimeStr      string //时间 格式为： 2020-01-01
}

//查询手艺人占用时间返回
type ReplyTime struct {
	ReserveEndTime      int
	ReserveEndTimeStr   string
	ReserveStartTime    int
	ReserveStartTimeStr string
	UnreservableTime    []map[string]interface{}
}
type Args struct{}

//健康状态
type Healths struct {
	Id   int
	Name string
}

//疫情状态
type Epidemics struct {
	Id   int
	Name string
}

//出行方式
type Traffics struct {
	Id   int
	Name string
}

//出行地点
type Address struct {
	Id   int
	Name string
}

//返回顾客健康信息
type ReservationHealth struct {
	Healths   []Healths
	Epidemics []Epidemics
	Traffics  []Traffics
	Address   []Address
}

type ReplyReservationHealth struct {
	ToTalNum int
	Lists []EsReservationHealth
}

type EsReservationHealth struct {
	ReservationId int
	Uid int
	Name string
	Phone string
	ShopId int
	ShopName string
	BranchName string
	OrderNumber string
	ReachHealthId int
	LeaveHealthId int
	ArriveHealthId int
	ReservationStartTime string
}

type Common interface {
	//添加顾客健康信息
	AddHealth(ctx context.Context, args *ArgsHealthAdd, reply *int) error
	//查询顾客健康信息
	GetHealth(ctx context.Context, args *ArgsHealthGet, reply *ReplyHealth) error
	//根据服务id和手艺人id查询时间
	GetReservationTime(ctx context.Context, args *ArgsTimeGet, reply *ReplyTime) error
	//添加顾客到店健康信息
	AddHealthArrive(ctx context.Context, args *ArgsHealthAdd, reply *int) error
	//添加顾客离店健康信息
	AddHealthLeave(ctx context.Context, args *ArgsHealthAdd, reply *int) error
}
