package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//监管平台信息对接成功的商户

const (
	//是否已添加到saas 1=否 2=是
	IS_IN_SAAS_no  = 1
	IS_IN_SAAS_yes = 2

	//是否暂停发卡 1=否 2=是
	IS_STOP_SELL_no  = 1
	IS_STOP_SELL_yes = 2
)

type ArgsAnsyBus struct {
	CompanyName     string //公司名称
	RiskBusId       int    //风控系统的商家id
	ParentRiskBusId int    //所属总公司在风控系统的商家id
}

type ArgsGetByCompanyname struct {
	common.BsToken
	CompanyName string //公司名称
	BusId       int    //内部使用
}

type ReplyGetByCompanyname struct {
	GovBusId int
	IsInSaas int
}

type ReplyGetCommentsByRiskbusids struct {
	PriceScore   float64 //价格评分
	ServiceScore float64 //服务评分
	EnvirScore   float64 //环境评分
	CompScore    float64 //综合评分
}

//是否暂停发卡参数
type ArgsStopOrStartSell struct {
	IsStop    int //是否停止发卡
	RiskBusId int //风控商家id
}

type ArgsChangeSellUplimiter struct {
	RiskBusId int    //风控商家id
	Uplimiter string //发卡额度上限 单位 万
}

type ReplyBusGovRule struct {
	BusId         int
	RiskBusId     int
	IsStopSell    int
	SellUplimiter float64
}

type GovBus interface {
	//同步监管平台对接成功的商户数据
	AnsyBus(ctx context.Context, args *ArgsAnsyBus, reply *bool) error
	//根据公司名称获取商家是否已对接
	GetByCompanyname(ctx context.Context, args *ArgsGetByCompanyname, reply *ReplyGetByCompanyname) error
	//根据riskbusIds 获取商家的评论
	GetCommentsByRiskbusids(ctx context.Context, riskBusIds *[]int, reply *map[int]ReplyGetCommentsByRiskbusids) error
	//监管平台添加商家发卡规则 - 暂停发卡
	StopOrStartSell(ctx context.Context, args *ArgsStopOrStartSell, reply *bool) error
	//监管平台添加商家发卡规则 - 设置发卡上限额度
	ChangeSellUplimiter(ctx context.Context, args *ArgsChangeSellUplimiter, reply *bool) error
	//获取商家的发卡状态 发卡额度情况
	BusGovRule(ctx context.Context, busId *int, reply *ReplyBusGovRule) error
}
