package market

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

const (
	//权益类型
	EQUITY_TYPE_BUY      = 0 //购卡所得
	EQUITY_TYPE_RECHARGE = 1 //充值所得
	//权益状态
	EQUITY_STATUS_CAN_USE     = 1 //1=可使用
	EQUITY_STATUS_NOT_CAN_USE = 2 //2=已使用完
	//过期状态
	EXPIRE_STATUS_FILTER = 1 //过滤过期的
	EXPIRE_STATUS_UNIQUE = 2 //查询过期的

)

var EquityMap = map[int]string{
	EQUITY_TYPE_BUY:      "购卡所得",
	EQUITY_TYPE_RECHARGE: "充值所得",
}

//添加权益入参
type ArgsAddEquity struct {
	BusId      int
	ShopId     int
	Uid        int    //权益用户id
	EquityType int    //权益类型 默认 0=购卡获得
	EquityName string //权益名称 默认 请传 购卡权益包
	ExpireDate int    //过期时间  传 天数
	RelationId int    //卡包关联id
	Details    []ServerDetailArgs
	Desc       string //使用说明
}

//权益服务详情
type ServerDetailArgs struct {
	SingleId         int
	Num              int
	SingleName       string
	ExpireDate       int64 //过期时间,权益过期时间
	PeriodOfValidity int   //有效期，单位天
}

//查询权益列表
type ArgsGetEquityList struct {
	common.BsToken
	common.Paging
	Uid    int
	Status int
	ExpireStatus int  //0查询全部  1过滤掉已经过期 2只查询过期的
}

type ReplyGetEquityList struct {
	TotalNum int
	Lists    []GetEquityList
}

type GetEquityList struct {
	//根据权益类型去映射 图片
	ImgId            int    //权益卡图片id
	ImgPath          string //权益卡图片路径
	Id               int    //权益卡id
	EquityType       int    //权益类型
	EquityTypeName   string //权益类型
	EquityName       string //权益名称
	Status           int    //权益状态 0可使用 1已使用完
	RelationId       int    //关系id
	ExpireDate       int    //权益过期时间 0表示不过期
	ExpireDateString string //权益国企时间 0表示不过期
	Count            int    //总次数
	SurCount         int    //剩余总次数
	ShopId           int
	ShopName         string
	BranchName       string
	PeriodOfValidity int  	//有效期
	SurTime 		 int   //剩余天数
	IsLate			 int   //是否会过期 0表示不会  1表示有期限
}

//用户查询权益列表
type ArgsGetEquityListByUid struct {
	common.Utoken
	common.Paging
	Status int
	ExpireStatus int
}

//查询一条详情
type ArgsGetOneEquity struct {
	common.Utoken
	Id              int
	EquitySingleIds []int //权益卡下的单项目
}

type ReplyGetOneEquity struct {
	ImgId            int    //权益卡图片id
	ImgPath          string //权益卡图片路径
	Id               int    //权益卡id
	ShopId           int    //门店id
	ShopName         string //门店名称
	BranchName       string //分店名称
	BusId            int
	RelationId       int    //关联id
	EquityType       int    //权益类型
	EquityTypeName   string //权益类型
	EquityName       string //权益名称
	Status           int    //权益状态 0可使用 1已使用完
	ExpireDate       int    //权益国企时间 0表示不过期
	ExpireDateString string //权益国企时间 0表示不过期
	Count            int    //总次数
	SurCount         int    //剩余总次数
	CreateTime       int    //权益获得时间
	CreateTimeStr    string //权益获得时间
	Details          []EquityDetail
	Desc             []string //使用说明
	PeriodOfValidity int  	//有效期
	SurTime 		 int   //剩余天数
	IsLate           int    //是否会过期 0表示不会过期 1表示有期限
}

type EquityDetail struct {
	SingleId      int
	SingleName    string
	Count         int
	SurCount      int //剩余总次数
	ServiceTime   int
	ImgUrl        string
	ImgHash       string
	ExpireDate    int64  //过期时间,权益过期时间
	ExpireDateStr string //过期时间,权益过期时间
	PeriodOfValidity int   //有效期，单位天
	SurTime 		 int   //剩余天数
	IsLate			  int  //是否会过期 0表示不会过期  1表示有期限
	SsId			 int   //单项目在门店的id
	RealPrice   float64
	Price       float64
	MinPrice  string
	MaxPrice  string
	ShopStatus int 	//项目在门店上下架的状态：1=下架 2=上架
	ShopDelStatus int	//项目在门店是否被删除：0-否，1-是
}

//确认消费权益卡
type ArgsConsumeEquity struct {
	common.BsToken
	Uid int
	ConsumeEquityBase
}

//确认消费权益卡-rpc
type ArgsConsumeEquityRpc struct {
	ShopId int
	BusId  int
	Uid    int
	ConsumeEquityBase
}
type ConsumeEquityBase struct {
	Uid         int
	EquityId    int             // 权益表id
	Details     []ConsumeEquity //服务明细
	ConfirmType int             //消费确认类型 1=短信验证 2=动态码验证
	Captcha     string          //短语验证码或者动态码
}
type ConsumeEquity struct {
	SingleId int    //单项目id
	StaffIds string //手艺人组合id
	Num      int    //次数
	EquityId int    // 权益表id,一期优化
}

//批量确认消费权益卡-rpc
type ArgsBatchConsumeEquity struct {
	common.BsToken
	IsAuthBsToken bool //web调用传true,此时shopId和busId可忽略,rpc调用传false，此时shopId和busId可必填
	ShopId        int
	BusId         int
	Uid           int
	ConfirmType   int    //消费确认类型 1=短信验证 2=动态码验证
	Captcha       string //短语验证码或者动态码
	Details       []ConsumeEquity
}

//根据权益ids批量获取权益列表基础数据
type ArgsGetEquityListsByIds struct {
	EquityIds []int
}
type ReplyGetEquityListsByIds struct {
	Lists []GetEquityListsByIdsBase
}
type GetEquityListsByIdsBase struct {
	Id            int //权益卡id
	ShopId        int //门店id
	BusId         int
	RelationId    int //关联id
	EquityType    int //权益类型
	Status        int //权益状态 0可使用 1已使用完
	Count         int //总次数
	SurCount      int //剩余总次数
	CreateTime    int //权益获得时间
	Details       []EquityDetail
	ExpireDate    int    //权益国企时间 0表示不过期
	ExpireDateStr string //过期时间,权益过期时间
}

//获取用户权益卡二维码信息入参
type ArgsGetUserEquityQrcode struct {
	common.Utoken
	EquityId int
}

//权益卡二维码扫一扫监测
type ArgsScanQrcodeCheck struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户信息
	EquityId       int
}

type ReplyScanQrcodeCheck struct {
	EquityId int
}

type ReplyGetQrcodeByConsumeCode struct {
	Uid      int
	EquityId int
}

//用户可使用权益卡统计
type ArgsGetUserEquityCountRpc struct {
	BusId  int
	ShopId int
	Uid    int
}
type ReplyGetUserEquityCountRpc struct {
	CanUserEquity int //可使用的权益卡
}

type ArgsGetEquityItemList struct {
	common.BsToken
	EquityId []int
	SingleId []int
}
type RepliesEquityDetailList struct {
	Id            int
	EquityId      int
	SingleId      int
	SingleName    string
	Count         int
	TransferCount int    //剩余总次数
	ExpireDate    int64  //过期时间,权益过期时间
	ExpireDateStr string //过期时间,权益过期时间
	PeriodOfValidity int   //有效期，单位天
}

type ArgsGetMemberDetailAsEquity struct {
	BusId int
	UId int
}

type ReplyGetMemberDetailAsEquity struct {
	AllRights int
	UseRights int
}

type Equity interface {
	//添加权益
	AddEquity(ctx context.Context, args *ArgsAddEquity, reply *int) error
	//查询权益列表
	GetEquityList(ctx context.Context, args *ArgsGetEquityList, reply *ReplyGetEquityList) error
	//用户查询权益列表
	GetEquityListByUid(ctx context.Context, args *ArgsGetEquityListByUid, reply *ReplyGetEquityList) error
	//查询一条详情
	GetOneEquity(ctx context.Context, args *ArgsGetOneEquity, reply *ReplyGetOneEquity) error
	//内部使用-查询一条详情
	GetOneEquityOfInternal(ctx context.Context, args *ArgsGetOneEquity, reply *ReplyGetOneEquity) error
	//确认消费权益卡
	ConsumeEquity(ctx context.Context, args *ArgsConsumeEquity, reply *bool) error
	//确认消费权益卡rpc
	ConsumeEquityRpc(ctx context.Context, args *ArgsConsumeEquityRpc, reply *bool) error
	//批量确认消费权益卡rpc
	BatchConsumeEquity(ctx context.Context, args *ArgsBatchConsumeEquity, reply *bool) error
	//根据权益ids批量获取权益列表基础数据
	GetEquityListsByIds(ctx context.Context, args *ArgsGetEquityListsByIds, reply *ReplyGetEquityListsByIds) error
	//获取用户权益卡二维码信息
	GetUserEquityQrcode(ctx context.Context, args *ArgsGetUserEquityQrcode, reply *string) error
	//根据消费码获取权益卡信息
	GetQrcodeByConsumeCode(ctx context.Context, args *order.ArgsCardPackageQrcode, reply *ReplyGetQrcodeByConsumeCode) error
	//用户可使用权益卡统计
	GetUserEquityCountRpc(ctx context.Context, args *ArgsGetUserEquityCountRpc, reply *ReplyGetUserEquityCountRpc) error
	//获取权益包下项目列表
	GetEquityItemList(ctx context.Context, args *ArgsGetEquityItemList, replies *[]RepliesEquityDetailList) error
	//bus member 查询 详情 使用rpc
	GetMemberDetailAsEquity(ctx context.Context, args *ArgsGetMemberDetailAsEquity, replies *ReplyGetMemberDetailAsEquity) error
}
