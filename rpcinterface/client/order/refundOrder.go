package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type RefundOrder struct {
	client.Baseclient
}

func (r *RefundOrder) Init() *RefundOrder {
	r.ServiceName = "rpc_order"
	r.ServicePath = "RefundOrder"
	return r
}

//根据支付订单id和订单类型 计算当前订单可退款总金额
func (r *RefundOrder) CalculateSingleOrCardRefundAmount(ctx context.Context, args *order.ArgsCalculateSingleOrCardRefundAmount, reply *order.ReplyCalculateSingleOrCardRefundAmount) error {
	return r.Call(ctx, "CalculateSingleOrCardRefundAmount", args, reply)
}

//退款申请
func (r *RefundOrder) RefundApply(ctx context.Context, args *order.ArgsRefundApply, reply *bool) error {
	return r.Call(ctx, "RefundApply", args, reply)
}

//退款订单列表
func (r *RefundOrder) GetRefundOrderList(ctx context.Context, args *order.ArgsGetRefundOrderList, reply *order.ReplyGetRefundOrderList) error{
	return r.Call(ctx, "GetRefundOrderList", args, reply)
}

//退款详情
func (r *RefundOrder)GetRefundOrderInfoById(ctx context.Context, refundOrderId *int, reply *order.ReplyGetRefundOrderInfoById) error{
	return r.Call(ctx, "GetRefundOrderInfoById", refundOrderId, reply)
}