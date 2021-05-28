/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/10 下午3:08

*******************************************/

package risk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/task"
)

//  添加预付卡风险管理系统-企业数据信息入参数
type ArgsStoreInformation struct {
	BusId int // 商户ID

}

// 添加成功返回信息
type ReplyRiskInformation struct {
	Id int // 预付卡风险管理系统id
}

//
type ArgsCashCardAssets struct {
	BusId          int
	CashCardAssets float64 //已兑付金额

}

type ReplyCashCardAssets struct {
	Id                  int     // 预付卡风险管理系统id
	CashCardAssetsTotal float64 // 已兑付金额
	CashRate            float64 // 累计兑付率

}

// 添加 投保率 入参
type ArgsInsuranceStatic struct {
	RelationId      int // 卡包关联ID
	CardPackageId   int // 卡包ID
	CardPackageType int // 卡包类型
}

//type ReplyInsuranceStatic struct {
//	//Id int
//	//InsuranceRate float64 // 投保率
//}

//监管平台直连接口-预付卡消费经营状况/门店信息
type ArgsGetBusCashRateAndShopInfo struct {
	RiskBusId int
}
type ReplyGetBusCashRateAndShopInfo struct {
	CashRate     float64 //兑付率
	ShopNum      int     //企业下的门店数量
	StaffNum     int     //企业下的技师数量
	CardTotalNum int     //发布产品/卡项/商品总数
}

//添加 统计预付卡风险管理系统信息
type ArgsInformation struct {
	BusId            int    // 商户ID
	ShopId           int    // 门店ID
	MemberId         int    // 会员人员ID
	StaffAddId       int    // 企业员工增加ID
	StaffDelId       int    // 企业员工减少ID
	OrdenrSn         string // 订单编号
	ConsumeLogId     int    // 消费者日志ID
	ServiceCommentId int    // 评分ID
}

type ReplyGetBussCashInfo struct {
	SalesCardAssets string //累计发卡金额
	CashCardAssets string  //已兑付金额
	CashRate string //累计兑付率
}

type BusBasicAccumulatice interface {
	//  添加 预付卡风险管理系统--riskId
	BusBasicRisId(ctx context.Context, busId int ,reply *bool) (err error)
	// 添加 预付卡风险管理系统--风控系统商户ID 门店总数量统计
	ShopNumRpc(ctx context.Context, shopId int, reply *bool)error
	// 添加 预付卡风险管理系统-会员总人数
	UserNumRpc(ctx context.Context, memberId *int, reply *bool) error
	// 添加 预付卡风险管理系统-统计企业员工人数
	StaffNumRpc(ctx context.Context, staffId *int, reply *bool)error
	// 添加 预付卡风险管理系统 ---- 会员复购率
	PurchaseRate(ctx context.Context, orderSn *string, reply *bool) error
	//添加 预付卡风险管理系统---统计累计已兑付（消费）金额  更新累计兑付率 所辖门店平均年限
	CashCardAssets(ctx context.Context, consumeLogId *int, reply *bool) error
	//添加 预付卡风险管理系统-消费者评分
	ConsumerEvaluation(ctx context.Context, serviceCommentId *int, reply *bool) error
	// 添加 已兑付金额
	AddCashCardAssets(ctx context.Context, args *ArgsCashCardAssets, reply *ReplyCashCardAssets) (err error)
	// 添加 投保率
	AddInsuranceRateRpc(ctx context.Context, args *ArgsInsuranceStatic, reply *bool) error
	//监管平台直连接口-预付卡消费经营状况
	GetBusCashRateAndShopInfo(ctx context.Context, args *ArgsGetBusCashRateAndShopInfo, reply *ReplyGetBusCashRateAndShopInfo) error
	// 添加 产品/卡项/商品总数
	AddGoodsNumRpc(ctx context.Context,args *task.ArgsAddDelGoods, reply *bool) error
	//监管平台直连接口-多商家的经营状况
	GetBussCashInfo(ctx context.Context, riskBusIds *[]int, reply *map[int]ReplyGetBussCashInfo) error
}
