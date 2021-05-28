package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type DayStatistics struct {
	client.Baseclient
}

func (d *DayStatistics)Init()*DayStatistics  {
	d.ServiceName = "rpc_order"
	d.ServicePath = "DayStatistics"
	return d
}

//确认消费完成,统计消费次数，今日耗卡金额
func (d *DayStatistics)	StatisConsume(ctx context.Context, consumeLogId *int, reply *bool) error  {
	return d.Call(ctx, "StatisConsume", consumeLogId, reply)
}

//获取卡包统计数据（今日耗卡总金额，今日消费人数，今日完成的订单数）
func (d *DayStatistics)GetCardPackageStatistics(ctx context.Context, args *order.ArgsGetCardPackageStatistics, reply *order.ReplyGetCardPackageStatistics) error{
	return d.Call(ctx, "GetCardPackageStatistics", args, reply)
}

//订单支付成功，统计（支付订单数，售卡金额，充值卡金额）
func (d *DayStatistics)StatisticsOrderPaySuc(ctx context.Context, orderSn *string, reply *bool) error {
	return d.Call(ctx, "StatisticsOrderPaySuc", orderSn, reply)
}

// 获取订单 待提单数据（今日待提单总数目， 待提单待处理总数）
func (d *DayStatistics) GetOrderRaisedNum(ctx context.Context,  args *order.ArgsGetOrderTotalStatic, reply *order.ReplyGetOrderTotalStatic) (err error){
	return d.Call(ctx, "GetOrderRaisedNum", args, reply)
}