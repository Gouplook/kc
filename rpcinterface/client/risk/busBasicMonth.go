package risk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/risk"
)

type BusBasicMonth struct {
	client.Baseclient
}

func (b *BusBasicMonth) Init() *BusBasicMonth {
	b.ServiceName = "rpc_risk"
	b.ServicePath = "BusBasicMonth"
	return b
}

//统计企业当月新增会员数量
func (b *BusBasicMonth) AddMemberNumRpc(ctx context.Context, memberId *int /*会员id*/, reply *bool) error {
	return b.Call(ctx, "AddMemberNumRpc", memberId, reply)
}

//统计当月发卡数量-（前提是订单支付成功）
func (b *BusBasicMonth) AddSalesCardNumRpc(ctx context.Context, args *risk.ArgsSalesCardNum, reply *bool) error {
	return b.Call(ctx, "AddSalesCardNumRpc", args, reply)
}

//统计当月消费人数（消费次数）
func (b *BusBasicMonth) AddServiceNumRpc(ctx context.Context, args *risk.ArgsAddServiceNum, reply *bool) error {
	return b.Call(ctx, "AddServiceNumRpc", args, reply)
}

//统计当月员工新增/离职率
func (b *BusBasicMonth) StatisticsStaffAddRateRpc(ctx context.Context, args *risk.ArgsStatisticsStaffAddRate, reply *bool) error {
	return b.Call(ctx, "StatisticsStaffAddRateRpc", args, reply)
}

//统计当月/年度营业额=单项目+售卡+商品金额
func (b *BusBasicMonth) StatisticsEarnedProfitRpc(ctx context.Context, args *risk.ArgsStatisticsEarnedProfit, reply *bool) error {
	return b.Call(ctx, "StatisticsEarnedProfitRpc", args, reply)
}

//用户退款-减少企业当月/年度营业额
func (b *BusBasicMonth) DescEarnedProfitRpc(ctx context.Context, args *risk.ArgsDescEarnedProfitRpc, reply *bool) error {
	return b.Call(ctx, "DescEarnedProfitRpc", args, reply)
}

//本月用户活跃度
func (b *BusBasicMonth) StatisticsUserActiveRpc(ctx context.Context, args *risk.ArgsStatisticsUserActive, reply *bool) error {
	return b.Call(ctx, "StatisticsUserActiveRpc", args, reply)
}

//统计当月单售卡张卡最高金额
func (b *BusBasicMonth) StatisticsMaxOrderAssetsRpc(ctx context.Context, args *risk.ArgsStatisticsMaxOrderAssets, reply *bool) error {
	return b.Call(ctx, "StatisticsMaxOrderAssetsRpc", args, reply)
}

//统计本月售卡/消费金额-订单付款成功/确认消费时调用
func (b *BusBasicMonth) SalesOrCashCardAssetsRpc(ctx context.Context, args *risk.ArgsSalesOrCashCardAssets, reply *bool) error {
	return b.Call(ctx, "SalesOrCashCardAssetsRpc", args, reply)
}

//统计当月卡项/商品下架率
func (b *BusBasicMonth) StatisticsCardProductUnderRateRpc(ctx context.Context, args *risk.ArgsStatisticsCardProductUnderRate, reply *bool) error {
	return b.Call(ctx, "StatisticsCardProductUnderRateRpc", args, reply)
}

//确认消费完成,统计消费次数，今日耗卡金额
func (b *BusBasicMonth) RiskStatisticsConsume(ctx context.Context, consumeLogId *int, reply *bool) error {
	return b.Call(ctx, "RiskStatisticsConsume", consumeLogId, reply)
}

//订单支付成功统计 单项目，卡项，商品金额
func (b *BusBasicMonth) RiskStatisticsOrderPaySuc(ctx context.Context, orderSn *string, reply *bool) error {
	return b.Call(ctx, "RiskStatisticsOrderPaySuc", orderSn, reply)
}
