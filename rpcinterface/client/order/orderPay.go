//支付信息客户端
//@author yangzhiwu<578154898@qq.com>
//@date 2020/7/29 16:52

package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type OrderPay struct {
	client.Baseclient
}

//初始化
func (o *OrderPay) Init() *OrderPay {
	o.ServiceName = "rpc_order"
	o.ServicePath = "OrderPay"
	return o
}

//获取支付签名
func (o *OrderPay) GetPaySign(ctx context.Context, args *order.ArgsGetPaySign, reply *order.ReplyGetPayInfo) error {
	return o.Call(ctx, "GetPaySign", args, reply)
}

//获取支付信息
func (o *OrderPay) GetPayInfo(ctx context.Context, args *order.ArgsGetPayInfo, reply *order.ReplyGetPayInfo) error {
	return o.Call(ctx, "GetPayInfo", args, reply)
}

//支付成功业务处理
func (o *OrderPay) PaySuc(ctx context.Context, orderSn *string, reply *bool) error {
	return o.Call(ctx, "PaySuc", orderSn, reply)
}

//获取订单分账信息
func (o *OrderPay) GetOrderSplitBill(ctx context.Context, orderSn *string, reply *order.ReplyGetOrderSplitBill) error {
	return o.Call(ctx, "GetOrderSplitBill", orderSn, reply)
}

//单项目现金支付
func (o *OrderPay) CashPaySuc(ctx context.Context, args *order.ArgsCashPaySuc, reply *bool) error {
	return o.Call(ctx, "CashPaySuc", args, reply)
}

//获取支付状态信息 0=待支付 1=支付成功 2=支付失败
func (o *OrderPay) QueryPayStatus(ctx context.Context, orderSn *string, reply *int) error {
	return o.Call(ctx, "QueryPayStatus", orderSn, reply)
}

//订单-获取服务订单列表
func (o *OrderPay) GetServiceOrderList(ctx context.Context, args *order.ArgsServiceOrderList, reply *order.ReplyServiceOrderList) error {
	return o.Call(ctx, "GetServiceOrderList", args, reply)
}

//获取商品订单信息
func (o *OrderPay) GetProductOrderList(ctx context.Context, args *order.ArgsProductOrder, reply *order.ReplyProductOrder) error {
	return o.Call(ctx, "GetProductOrderList", args, reply)
}

//获取用户提货单
func (o *OrderPay) GetUserPickUpGoods(ctx context.Context, args *order.ArgsPickUpGoods, reply *order.ReplyPickUpGoods) error {
	return o.Call(ctx, "GetUserPickUpGoods", args, reply)
}

//订单超时未支付，关闭订单
func (o *OrderPay) CloseOrder(ctx context.Context, orderId *int, reply *bool) error {
	return o.Call(ctx, "CloseOrder", orderId, reply)
}

//获取一条商品订单详情信息
func (o *OrderPay) GetOneProductDetail(ctx context.Context, args int, reply *order.ReplyOneProductDetail) error {
	return o.Call(ctx, "GetOneProductDetail", args, reply)
}

//获取一条服务订单详情信息
func (o *OrderPay) GetOneServerDetail(ctx context.Context, args *order.ArgsGetOneServerDetail, reply *order.ReplyOneServerDetail) error {
	return o.Call(ctx, "GetOneServerDetail", args, reply)
}

// 统计用户复购率
func (o *OrderPay) GetUsePurchaseRate(ctx context.Context, args *order.ArgsUserPurchaseRate, reply *order.ReplyUserPurchaseRate) error {
	return o.Call(ctx, "GetUsePurchaseRate", args, reply)
}
