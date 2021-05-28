//订单相关任务入队列
//@author yangzhiwu<578154898@qq.com>
//@date 2020/8/5 14:38
package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/order"
)

type Order struct {
	client.Baseclient
}

func (o *Order) Init() *Order  {
	o.ServiceName = "rpc_task"
	o.ServicePath = "Order/Order"
	return o
}

//将卡包主表的主键id加入交换机
func (o *Order) SetCardPackageRelationId(ctx context.Context, cardPackageRelationId *int, reply *bool ) error {
	return o.Call(ctx, "SetCardPackageRelationId", cardPackageRelationId, reply)
}

//订单超时未付款订单id加入交换机
func (o *Order) TimeOutOrderid(ctx context.Context, args *order.ArgsTimeOutOrderid, reply *bool) (err error)  {
	return o.Call(ctx, "TimeOutOrderid", args, reply)
}

//挂单超时取消挂单
func (o *Order) TimeOutTemp(ctx context.Context, args *order.ArgsTimeOutTemp, reply *bool) error {
	return o.Call(ctx, "TimeOutTemp", args, reply)
}