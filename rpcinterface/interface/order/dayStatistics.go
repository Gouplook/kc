package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//获取卡包统计数据入参
type ArgsGetCardPackageStatistics struct {
	common.BsToken
	DayTime       string //今日日期，格式：20060102
	YesterdayTime string //昨日日期，格式：20060102
}
type CardPackageStatisticsBase struct {
	TodayTotalAmount     float64 //今日总金额
	YesterdayTotalAmount float64 //昨日总金额
	UpDown               float64 //较昨日涨跌幅度
}
type ConsumeNumStatisticsBase struct {
	TodayMemberConsumeNum         int     //今日会员消费人数
	TodayNoMemberConsumeNum       int     //今日非会员消费人数
	TodayTotalConsumeNum     int     //今日总消费次数
	YesterdayTotalConsumeNum int //昨日总消费次数
	UpDown                   float64 //较昨日涨跌幅度
}

//获取卡包统计数据出参
type ReplyGetCardPackageStatistics struct {
	Total           CardPackageStatisticsBase //今日总营业额
	ConsumptionCard CardPackageStatisticsBase //今日耗卡总金额
	SellCard        CardPackageStatisticsBase //今日售卡总金额
	Recharge        CardPackageStatisticsBase //今日充值金额
	Single          CardPackageStatisticsBase //今日单项目总金额
	Goods           CardPackageStatisticsBase //今日商品总金额
	Refund          CardPackageStatisticsBase //今日退款总金额
	SuccessOrder    int                       //今日支付成功的订单
	ConsumeNum      ConsumeNumStatisticsBase                       //消费人数
}

type ArgsGetOrderTotalStatic struct {
	common.BsToken
}

type ReplyGetOrderTotalStatic struct {
	TodayOrderTotalNum int // 今日待提单总数目
	PendingTotalNum    int // 待提单待处理总数
}

type DayStatistics interface {
	//确认消费完成,统计消费次数，今日耗卡金额
	StatisConsume(ctx context.Context, consumeLogId *int, reply *bool) error
	//获取卡包统计数据（今日耗卡总金额，今日消费人数，今日完成的订单数）
	GetCardPackageStatistics(ctx context.Context, args *ArgsGetCardPackageStatistics, reply *ReplyGetCardPackageStatistics) error
	//订单支付成功，统计（支付订单数，售卡金额，充值卡金额）
	StatisticsOrderPaySuc(ctx context.Context, orderSn *string, reply *bool) error

	// 获取订单 待提单数据（今日待提单总数目， 待提单待处理总数）
	GetOrderRaisedNum(ctx context.Context, args *ArgsGetOrderTotalStatic, reply *ReplyGetOrderTotalStatic) error
}
