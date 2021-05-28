package bus

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// 杉德结算账户信息
type ReplyDeposSandInfo struct {
	BusId        int
	AcctName     string // 子账户名
	AcctNo       string // 子账号
	BankName     string // 留存比例
	BankNo       string // 存管商户编号
	DepositRatio string // 留存比例
}

// 建行结算账户信息
type ReplyDeposCcbInfo struct {
	BusId        int
	AcctName     string // 子账户名
	AcctNo       string // 子账号
	DepositRatio string // 留存比例
	MerchantId   string // 存管商户编号
}

const (
	//存管账户类型 1=上海银行 2=交通银行 3=平安银行 4=工商银行 5=宁波银行
	DEPOS_TYPE_BOSC   = 1
	DEPOS_TYPE_BOCOM  = 2
	DEPOS_TYPE_PINGAN = 3
	DEPOS_TYPE_ICBC   = 4
	DEPOS_TYPE_NBCB   = 5

	//交通银行开户状态 0=已提交 1=待绑卡 2=开户成功 3=失败
	BOCOM_STATUS_COMMIT = 0
	BOCOM_STATUS_BIND   = 1
	BOCOM_STATUS_SUC    = 2
	BOCOM_STATUS_FAIL   = 3
)

//存管账号信息
type AccountInfo struct {
	AccountType   int    //存管账号类型
	MerchantId    string //建设银行 - 商户编号
	OperatorPhone string //上海银行 - 操作员手机号
	AcctNo        string //存管账户户号
	AcctName      string //存管账户名称
	DepositRatio  string //资金留存比例
	BankName      string //结算户开户行
	BankCardNo    string //结算户卡号
	BankCardName  string //结算户账户名
	BankNo        string //结算户开户行联行号

}

//资金信息
type DeposInfo struct {
	TotalAmount        string //总金额
	DepositoryAmount   string //应存管金额
	UsableAmount       string //可提现金额
	UndischargedAmount string //待清算金额
	CashingAmount      string //提现中金额

}

type ArgsGetBusDeposInfo struct {
	common.BsToken
}

type ReplyGetBusDeposInfo struct {
	OpenDepos int         //是否开通存管账户 0=否 1=是
	AcctInfo  AccountInfo //存管账户信息
	DepInfo   DeposInfo   //资金信息
	Cid       int         //商家所在城市id
}

type ArgsBusPage struct {
	common.BsToken
	common.Paging
	StartTime int //开始时间
	EndTime   int //结算时间
	Status    int //状态
}

//获取资金明细
type GetBusAmountLogsPage struct {
	Amount        string //金额
	Type          int    //记录类型
	OrderType     int    //记录来源
	FundType      int    //资金类型
	CreateTime    int    //交易时间
	CreateTimeStr string //交易时间格式化
}

type ReplyGetBusAmountLogs struct {
	TotalNum int
	Lists    []GetBusAmountLogsPage
}

//获取存管明细
type GetBusDeposLogsPage struct {
	Amount        string //金额
	Type          int    //记录类型
	OrderType     int    //账单类型
	CreateTime    int    //记录时间
	CreateTimeStr string //记录时间格式化
	RecordTime    int    //出入账时间
	RecordTimeStr string //出入账时间格式化
}

type ReplyGetBusDeposLogs struct {
	TotalNum int
	Lists    []GetBusDeposLogsPage
}

type GetBusCashLogsPage struct {
	Id         string
	OrderSn    string //订单号
	CashAmount string //提现金额
	Status     int    //提现状态
	BankCardNo string //收款银行卡
	Ctime      int    //申请时间
	CtimeStr   string //申请时间格式化
	Ntime      int    //审核时间
	NtimeStr   string //审核时间格式化
	FailReason string //失败原因

}

type ReplyGetBusCashLogsPage struct {
	TotalNum int
	Lists    []GetBusCashLogsPage
}

//申请提现
type ArgsCashApply struct {
	common.BsToken
	Amount  float64 //提现金额
	Captcha string  //体验验证码 - 上海银行提现需要
}

//企业存管
type Depos interface {
	//获取留存比列
	GetRatio(ctx context.Context, busId *int, ratio *float64) error
	// 批量获取杉德存管账户信息
	GetDeposSandInfos(ctx context.Context, busIds *[]int, replyDeposSandInfo *[]ReplyDeposSandInfo) error
	// 批量获取建行存管账户信息
	GetDeposCcbInfos(ctx context.Context, busIds *[]int, replyDeposCcbInfos *[]ReplyDeposCcbInfo) error
	//获取商家的存管账号和存管资金信息
	GetBusDeposInfo(ctx context.Context, args *ArgsGetBusDeposInfo, reply *ReplyGetBusDeposInfo) error
	//获取资金明细
	GetBusAmountLogs(ctx context.Context, args *ArgsBusPage, reply *ReplyGetBusAmountLogs) error
	//获取存管明细
	GetBusDeposLogs(ctx context.Context, args *ArgsBusPage, reply *ReplyGetBusDeposLogs) error
	//获取提现记录
	GetBusCashLogs(ctx context.Context, args *ArgsBusPage, reply *ReplyGetBusCashLogsPage) error
	//申请提现
	CashApply(ctx context.Context, args *ArgsCashApply, reply *bool) error
	//发送上海银行提现验证码
	BoscCashCaptcha(ctx context.Context, bsToken *common.BsToken, reply *bool) error
	//上海银行提现结果处理业务 - 通过任务查询提现结果
	BoscCashResultQuery(ctx context.Context, orderSn *string, reply *bool) error
	//获取商家的存管账号信息
	GetBusDeposInfoByBusid(ctx context.Context, busId *int, reply *ReplyGetBusDeposInfo) error
}
