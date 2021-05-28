//卡项订单客户端
//@author yangzhiwu<578154898@qq.com>
//@date 2020/7/29 16:47

package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type ItemOrder struct {
	client.Baseclient
}

//初始化
func (i *ItemOrder) Init() *ItemOrder {
	i.ServiceName = "rpc_order"
	i.ServicePath = "ItemOrder"
	return i
}

//saas 创建订单
func (i *ItemOrder) SaasCreateItemOrder(ctx context.Context, args *order.ArgsSaasCreateItemOrder, reply *order.ReplySaasCreateItemOrder) error {
	return i.Call(ctx, "SaasCreateItemOrder", args, reply)
}

//前端用户创建订单
func (i *ItemOrder) UserCreateItemOrder(ctx context.Context, args *order.ArgsUserCreateItemOrder, reply *order.ReplyUserCreateItemOrder ) error  {
	return i.Call(ctx, "UserCreateItemOrder", args, reply)
}

//根据订单号，获取订单的详细-rpc使用
func (i *ItemOrder)  GetOrderInfoByOrderSnRpc(ctx context.Context, orderSn *string, reply *order.ReplyGetOrderInfoByOrderSnRpc) error{
	return i.Call(ctx, "GetOrderInfoByOrderSnRpc", orderSn, reply)
}

// 获取卡项目购卡总数（对应的店铺）
func (i *ItemOrder)GetBuyCardNum(ctx context.Context, args *order.ArgsBuyCardNum, reply *order.ReplyBuyCardNum)(err error){
	return i.Call(ctx, "GetBuyCardNum", args ,reply)
}

//获取商家订单存管明细数据
func (i *ItemOrder)GetOrderDeposLists(ctx context.Context, args *order.ArsgOrderDeposLists, reply *order.ReplyOrderDeposLists) error{
	return i.Call(ctx, "GetOrderDeposLists", args ,reply)
}

//获取卡包存管资金释放明细
func (i *ItemOrder)GetOrderDeposReleaseInfo(ctx context.Context,args *order.ArgsGetOrderDeposReleaseInfo,reply *order.ReplyGetOrderDeposReleaseInfo)error{
	return i.Call(ctx, "GetOrderDeposReleaseInfo", args ,reply)
}

//根据商家id，获取商家订单存管明细 - 监管平台直连接口
func (i *ItemOrder) GetBusOrderListsForGov(ctx context.Context, args *order.ArgsGetBusOrderListsForGov, reply *order.ReplyGetBusOrderListsForGov) error{
	return i.Call(ctx, "GetBusOrderListsForGov", args ,reply)
}

//获取卡包的消费记录 - 监管平台直连接口
func (i *ItemOrder)  GetConsumeLogForGov(ctx context.Context, args *order.ArgsGetConsumeLogForGov, reply *order.ReplyGetConsumeLogForGov ) error {
	return i.Call(ctx, "GetConsumeLogForGov", args ,reply)
}