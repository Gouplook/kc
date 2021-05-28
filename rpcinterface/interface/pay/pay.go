/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:38
 */
package pay

import (
	"context"
)

const (
	PAY_Suc  = 1
	PAY_Fail = 2

	CLEAR_STATUS_none = 0
	CLEAR_STATUS_ing  = 1
	CLEAR_STATUS_suc  = 2
	CLEAR_STATUS_fail = 3

	//平安分账信息 - 支付模式 0-冻结支付 1-普通支付
	PINGAN_PAYMODEL_dj     = 0
	PINGAN_PAYMODEL_normal = 1
)

type Wx struct {
	AppId  string // 微信appid
	OpenId string // 微信openid
}

type PayInfo struct {
	OrderSn           string // 订单编号
	BusId             int    // 购买商家总店id
	RealAmount        string // 订单总金额
	InsureAmount      string // 保险费用
	RenewInsureAmount string // 续保费用
	PlatformAmount    string // 平台手续费
	BusAmount         string // 商户收取金额
	PayChannel        int    // 支付渠道
	InsuranceChannel  int    // 保险渠道
	ChosePayType      int    // 支付方式
	Wx
	Version      string
	CreateIP     string
	StoreID      string
	PayExtra     string
	AccsplitFlag string
	SignType     string
	FormUrl      string      // 成功后跳转连接
	Cid          int         //当前被下单商户所属城市，如该商户在上海，则传321
	OrderRemark  []OrderList //分账标签-平安银行渠道适用，其他渠道忽略
}

//订单备注明细结构体
type OrderRemark struct {
	Oderlist     []OrderList `json:"oderlist"`
	SFJOrdertype string      //订单类型 1=子订单
	Remarktype   string      `json:"remarktype"` //备注类型 默认取值 JHS0100000
	PlantCode    string      `json:"plantCode"`  //平台代码
}

//订单备注中的订单订单列表结构体
type OrderList struct {
	PayModel   string //支付模式 0-冻结支付 1-普通支付
	TranFee    string //手续费  单位 ：元； 金额为两位小数格式，例如0.01。注意是字符串、而非数字类型。
	SubAccNo   string //入账会员子账户
	Subamount  string `json:"subamount"`  //子订单金额 单位 ：元； 金额为两位小数格式，例如0.01。注意是字符串、而非数字类型。
	SuborderId string `json:"suborderId"` //子订单号 不可超过22位，且全局唯一
	Object     string `json:"object"`     //子订单简单描述
}

type ArgsNotify struct {
	PayChannel int // 支付渠道
	Data       string
}

//PayH5 PayH5
type PayH5 struct {
	PayURL  string
	PayBoby string
	PaySign string
}

type PayNotify struct {
	OrderSn    string // 订单编号
	PayTime    int64  // 交易时间 时间戳
	PayAmount  string // 支付金额
	PayFee     string // 手续费
	PaySn      string // 支付流水号
	PayStatus  int    // 订单状态
	PayChannel int    //支付渠道
	PayType    int    //付款方式
}

type ReplyGetAgentInfoByOrderSn struct {
	BusClearId int //商户单日资金结算表id
	Status     int //代付状态
	BusId      int //商家id
	AgentType  int //清算类型 2=上海杉德清算 3=杭州建行清算
}

type ReplyNotifyResponse struct {
	ResponseJsonStr string
}

type Pay interface {
	// 获取支付二维码
	PayQr(ctx context.Context, args *PayInfo, reply *string) error
	// 获取H5支付连接
	PayH5(ctx context.Context, args *PayInfo, reply *string) error
	// 获取H5支付连接
	PayH5New(ctx context.Context, args *PayInfo, reply *PayH5) error
	// 获取小程序支付数据
	PayWxapp(ctx context.Context, args *PayInfo, reply *string) error
	// 微信公众号支付数据
	PayWxOfficial(ctx context.Context, args *PayInfo, reply *string) error
	// 获取app支付串
	PayApp(ctx context.Context, args *PayInfo, reply *string) error
	// 获取app支付串
	PayAppSign(ctx context.Context, args *PayInfo, reply *string) error
	// 支付回调
	Notify(ctx context.Context, args *ArgsNotify, reply *bool) error
	//支付回调响应（工行有这个要求）
	NotifyResponse(ctx context.Context, args *ArgsNotify, reply *ReplyNotifyResponse)
	// 获取支付成功的订单信息
	PayInfo(ctx context.Context, orderSn *string, reply *PayNotify) error
	// 支付成功清分资金记账处理 支付成功后消息任务调度
	PayAgent(ctx context.Context, orderSn *string, reply *bool) error
	// 分账数据初始化 每日处理前一日的清分数据 在具体清分操作前调用
	AngelChannel(ctx context.Context, timeUnix *int, reply *bool) error
	// 商家建行分账清分
	CcbAgent(ctx context.Context, timeUnix *int, reply *bool) error
	// 商家建行分账结果处理
	CcbAgentFund(ctx context.Context, timeUnix *int, reply *bool) error
	// 商家杉德分账清分
	SandAgent(ctx context.Context, timeUnix *int, reply *bool) error
	// 商家杉德分账异步结果处理
	SandAgentFund(ctx context.Context, timeUnix *int, reply *bool) error
	// 保险公司杉德分账清分
	SandAgentInsure(ctx context.Context, timeUnix *int, reply *bool) error
	// 保险公司杉德分账异步结果处理
	SandAgentFundInsure(ctx context.Context, timeUnix *int, reply *bool) error
	// 续保公司杉德分账清分
	SandAgentRenewInsure(ctx context.Context, timeUnix *int, reply *bool) error
	// 续保公司杉德分账异步结果处理
	SandAgentFundRenewInsure(ctx context.Context, timeUnix *int, reply *bool) error
	//根据订单号，获取代付状态信息
	GetAgentInfoByOrderSn(ctx context.Context, orderSn *string, reply *ReplyGetAgentInfoByOrderSn) error
	//根据clearId，获取代付状态信息
	GetAgentInfoByClearid(ctx context.Context, clearId *int, reply *ReplyGetAgentInfoByOrderSn) error
}
