package order

import (
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"golang.org/x/net/context"
)

// 定义卡包-单项目数据入参出参数据
// @author liyang<654516092@qq.com>
// @date  2020/7/23 9:48


//获取卡包详情-用户
type ArgsGetUserSingle struct {
	common.Utoken   //用户信息
	Id int          //卡包Id
}

//获取卡包详情-企业/商户分店
type ArgsGetBusSingle struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户/用户信息
	Id  int   //卡包ID
}

//获取卡包详情-RPC
type ArgsGetRpcSingle struct {
	Id int  //卡包ID
}


//卡包详情
type CardSingleInfo struct {
	Id int //卡包ID
	CardSn string //编号
	BusId  int //企业/商户ID
	ShopId int //分店ID
	Uid    int //下单用户ID
	RealPrice float64 //实际金额
	Price     float64 //面值金额
	ConsumePrice float64 //消费面值金额
	ActualConsumePrice float64 //实际消费金额
	Disaccount  float64 //折扣率
	CardId      int     //卡ID
	CardName    string  //卡名称
	SspId int
	SspName string
	Status      int     //状态
	PayTime     int64   //付款时间
	ConsumeingTime int64 //最近一次消费时间
	ConsumeCompTime int64 //消费完成时间
	PayChannel      int   //支付渠道
	FundMode        int   //资金管理方式
	InsuranceChannel int //保险渠道
	DepositRatio     float64 //留存比例
	DepositAmount    float64 //留存金额
	ReleaseAmount    float64 //已释放的留存金额
	Deleted          int     //是否正常显示
	Ctime            int     //生成时间
}

//用户卡包详情
type CardSingleUserInfo struct {
	SsId int //单项目在门店的id
	RealPrice float64 //实际金额
	Price     float64 //面值金额
	ConsumePrice float64 //消费面值金额
	ActualConsumePrice float64 //实际消费金额
	Disaccount  float64 //折扣率
	InsuranceChannel int //保险渠道
	CardName    string  //卡名称
	SingleId    int `mapstructure:"card_id"`
	ImgId       int    //卡封面
	PayTime     int64   //付款日期时间戳
	PayTimeStr  string   //付款日期
	SspId int    //规格ID
	CanUse int   //是否可用 默认=2可用
	ReservationNum int //已被预约的总次数
	SspName string  //规格名称
	ServiceTime int //服务时长
}

//用户卡包中包含的单项目
type CardSingleUserSingle struct {
	SingleId  int      //单项目ID
	SingleName string  //单项目名称
}


//定义接口
type CardPackageSingle interface {
	//获取单项目-用户
	GetInfoByUser(ctx context.Context,args *ArgsGetUserSingle,reply *CardSingleInfo) error
	//获取单项目-企业/商户/分店
	GetInfoByBus(ctx context.Context,args *ArgsGetBusSingle,reply *CardSingleInfo) error
	//获取单项目-rpc内部
	GetInfoByRpc(ctx context.Context,args *ArgsGetRpcSingle,reply *CardSingleInfo) error
}