package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)


const (
	//挂单支付状态
	TEMP_STATUS_PAYMENT_AWAIT = 1  		//待支付
	TEMP_STATUS_PAYMENT_PAYED = 2		//已支付
	TEMP_STATUS_PAYMENT_DELETED = 3		//已删除

	//挂单类型
	TEMP_TYPE_SINGLE = 1		//单项目
	TEMP_TYPE_PRODUCT = 2		//商品

)

//新增挂单入参
type ArgsAddTemp struct {
	common.BsToken
	common.Utoken
	Type int   		//1服务 2商品
	ReservationId int
	Uid int
	TotalMoney float64
	ServerContent []ServerContent		//服务明细
	ProductContent []ProductContent		//商品明细
}

type ServerContent struct {
	Ssid int		//单项目在门店的id
	SspId int		//单项目规格价格id
	SingleId int   // 单项目id
	Name string		//单项目名称
	SpecIds string	//单项目规格id组合
	SpecNames string	//单项目规格名称组合
	StaffId string		//手艺人id组合 多个用,号隔开
	StaffName string
	SalesStaffId string	//销售人id组合 多个用,号隔开
	SalesStaffName string
	Price float64		//单价
	Num int			//数量
}

type ProductContent struct {
	ProductId int		//商品id
	DetailId int		//明细id
	Name string			//商品名称
	SpecIds string	//单项目规格id组合
	SpecNames string	//单项目规格名称组合
	SalesStaffId string	//销售人id
	SalesStaffName string
	Price float64		//单价
	Num int			//数量
}

type ArgsGetOneTemp struct {
	common.BsToken
	TempId int
}

type ReplyGetOneTemp struct {
	TempId int
	BusId int
	ShopId int
	AdminUid int
	Uid int
	Uname string
	Unick string
	Phone string
	Sex int
	MemberInt int
	MemberName string
	MemberLevel int
	UimgId int
	UimgPath string
	Type int
	ReservationId int
	Status int
	TotalMoney float64
	Ctime int
	CtimeStr string
	ServerContent []ServerContent		//服务明细
	ProductContent []ProductContent		//商品明细
}

type ArgsGetTempList struct {
	common.BsToken
	common.Paging
	Status int  //0查询全部  1服务 2商品
	Phone string //手机号模糊查询
}

type ReplyGetTempList struct {
	TotalNum int
	Lists []GetTempList
}

type GetTempList struct {
	TempId int
	Ctime int
	CtimeStr string
	Uid int
	Uname string
	//NickName string
	Phone string
	ProductNames string
	ProductContent []ProductContent
	ServerNames string
	ServerContent []ServerContent
	StaffNames string
	SalesStaffNames string
	TotalMoney float64
	Status int
	Type int
}

type ArgsCancelTemp struct {
	common.BsToken
	common.Utoken
	TempId int
}

type Temp interface {
	//添加一条 挂单
	AddTemp(ctx context.Context,args *ArgsAddTemp,reply *int) error
	//延迟队列处理挂单
	MqCancelTemp(ctx context.Context,id int,reply *bool) error
	//取单
	GetOneTemp (ctx context.Context,args *ArgsGetOneTemp,reply *ReplyGetOneTemp) error
	//获取挂单列表
	GetTempList(ctx context.Context,args *ArgsGetTempList, reply *ReplyGetTempList) error
	//取消一条挂单
	CancelTemp(ctx context.Context, args *ArgsCancelTemp, reply *bool) error
}
