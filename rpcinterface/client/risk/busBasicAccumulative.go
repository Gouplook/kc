/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/10 下午3:25

*******************************************/
package risk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/risk"
	"git.900sui.cn/kc/rpcinterface/interface/task"
)

type BusBasicAccumulative struct {
	client.Baseclient
}

// 初始化
func (b *BusBasicAccumulative) Init() *BusBasicAccumulative {
	b.ServiceName = "rpc_risk"
	b.ServicePath = "BusBasicAccumulative"
	return b
}
//  添加 预付卡风险管理系统--riskId
func (b *BusBasicAccumulative)BusBasicRisId(ctx context.Context, busId int ,reply *bool) error {
	return b.Call(ctx, "BusBasicRisId", busId, reply)
}
// 添加 预付卡风险管理系统--风控系统商户ID 门店总数量统计
func (b *BusBasicAccumulative)ShopNumRpc(ctx context.Context, shopId int, reply *bool)error{
	return b.Call(ctx, "ShopNumRpc", shopId, reply)
}
// 添加 预付卡风险管理系统-会员总人数
func (b *BusBasicAccumulative)UserNumRpc(ctx context.Context, memberId *int, reply *bool) error{
	return b.Call(ctx, "UserNumRpc", memberId, reply)
}
// 添加 预付卡风险管理系统-统计企业员工人数
func (b *BusBasicAccumulative)StaffNumRpc(ctx context.Context, staffId *int, reply *bool)error{
	return b.Call(ctx, "StaffNumRpc", staffId, reply)
}
// 添加 预付卡风险管理系统 ---- 会员复购率
func (b *BusBasicAccumulative)PurchaseRate(ctx context.Context, orderSn *string, reply *bool) error{
	return b.Call(ctx, "PurchaseRate", orderSn, reply)
}
//添加 预付卡风险管理系统---统计累计已兑付（消费）金额  更新累计兑付率 所辖门店平均年限
func (b *BusBasicAccumulative)CashCardAssets(ctx context.Context, consumeLogId *int, reply *bool) error{
	return b.Call(ctx, "CashCardAssets", consumeLogId, reply)
}
//添加 预付卡风险管理系统-消费者评分
func (b *BusBasicAccumulative)ConsumerEvaluation(ctx context.Context, serviceCommentId *int, reply *bool) error{
	return b.Call(ctx, "ConsumerEvaluation", serviceCommentId, reply)
}
// 添加 已兑付金额
func (b *BusBasicAccumulative)AddCashCardAssets(ctx context.Context, args *risk.ArgsCashCardAssets, reply *risk.ReplyCashCardAssets) error {
	return b.Call(ctx, "AddCashCardAssets", args, reply)
}
// 添加 投保率
func (b *BusBasicAccumulative) AddInsuranceRateRpc(ctx context.Context, args *risk.ArgsInsuranceStatic, reply *bool) error {
	return b.Call(ctx, "AddInsuranceRateRpc", args, reply)
}


//监管平台直连接口-预付卡消费经营状况
func (b *BusBasicAccumulative) GetBusCashRateAndShopInfo(ctx context.Context, args *risk.ArgsGetBusCashRateAndShopInfo, reply *risk.ReplyGetBusCashRateAndShopInfo) error {
	return b.Call(ctx, "GetBusCashRateAndShopInfo", args, reply)
}

// 添加 产品/卡项/商品总数
func (b *BusBasicAccumulative) AddGoodsNumRpc(ctx context.Context,args *task.ArgsAddDelGoods, reply *bool) error {
	return b.Call(ctx, "AddGoodsNumRpc", args, reply)
}

//监管平台直连接口-多商家的经营状况
func (b *BusBasicAccumulative) GetBussCashInfo(ctx context.Context, riskBusIds *[]int, reply *map[int]risk.ReplyGetBussCashInfo) error {
	return b.Call(ctx, "GetBussCashInfo", riskBusIds, reply)
}


