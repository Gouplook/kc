package risk

import "context"

/*
	预付卡风险管理系统-企业每月商户统计数据
*/

type ArgsDateBusIdCommon struct {
	BusId    int
	DateTime int64
}

//统计当月发卡数量-入参
type ArgsSalesCardNum struct {
	ArgsDateBusIdCommon
	SalesCardNum int
}

//统计消费人数（消费次数）-入参
type ArgsAddServiceNum struct {
	ArgsDateBusIdCommon
}

//统计当月员工新增率-入参
type ArgsStatisticsStaffAddRate struct {
	StaffId int
}

//统计当月营业额=单项目+售卡+商品金额 入参
type ArgsStatisticsEarnedProfit struct {
	ArgsDateBusIdCommon
	SingleAmount  float64 //单项目金额
	CardAmount    float64 //卡项金额
	ProductAmount float64 //商品金额
}

//统计当月卡项/商品下架率
type ArgsStatisticsCardProductUnderRate struct {
	ShopId int
}

//本月用户活跃度
// 本月复购人数/总人数（去重）
type ArgsStatisticsUserActive struct {
	UserActive float64
	ArgsDateBusIdCommon
}

//当月单售卡张卡最高金额-入参
type ArgsStatisticsMaxOrderAssets struct {
	ArgsDateBusIdCommon
	Assets float64 //单张售卡金额
}

//统计本月售卡/消费金额-入参
type ArgsSalesOrCashCardAssets struct {
	ArgsDateBusIdCommon
	SalesCardAssets float64 //售卡金额
	CashCardAssets  float64 //消费金额
}

//退款-减少营业额(时间取当前订单的创建时间)
type ArgsDescEarnedProfitRpc struct {
	ArgsDateBusIdCommon
	RefundAssets float64 //退款金额
}

type BusBasicMonth interface {
	//统计当月企业新增会员数量
	AddMemberNumRpc(ctx context.Context, memberId *int /*会员id*/, reply *bool) error
	//统计当月发卡数量-（前提是订单支付成功）
	AddSalesCardNumRpc(ctx context.Context, args *ArgsSalesCardNum, reply *bool) error
	//统计当月消费人数（消费次数）
	AddServiceNumRpc(ctx context.Context, args *ArgsAddServiceNum, reply *bool) error
	//统计当月员工新增/离职率
	StatisticsStaffAddRateRpc(ctx context.Context, args *ArgsStatisticsStaffAddRate, reply *bool) error
	//统计当月/年度营业额=单项目+售卡+商品金额
	StatisticsEarnedProfitRpc(ctx context.Context, args *ArgsStatisticsEarnedProfit, reply *bool) error
	//用户退款-减少企业当月/年度营业额
	DescEarnedProfitRpc(ctx context.Context,args *ArgsDescEarnedProfitRpc,reply *bool)error
	//本月用户活跃度
	StatisticsUserActiveRpc(ctx context.Context, args *ArgsStatisticsUserActive, reply *bool) error
	//统计当月单售卡张卡最高金额
	StatisticsMaxOrderAssetsRpc(ctx context.Context, args *ArgsStatisticsMaxOrderAssets, reply *bool) error
	//统计本月售卡/消费金额-订单付款成功/确认消费时调用
	SalesOrCashCardAssetsRpc(ctx context.Context, args *ArgsSalesOrCashCardAssets, reply *bool) error
	//统计当月卡项/商品下架率
	StatisticsCardProductUnderRateRpc(ctx context.Context,args *ArgsStatisticsCardProductUnderRate,reply *bool)error
	//确认消费完成,统计消费次数，今日耗卡金额
	RiskStatisticsConsume(ctx context.Context, consumeLogId *int, reply *bool) error
	//订单支付成功，统计（支付订单数，售卡金额，充值卡金额）
	RiskStatisticsOrderPaySuc(ctx context.Context, orderSn *string, reply *bool) error
}
