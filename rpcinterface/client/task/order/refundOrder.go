/******************************************
@Description:
@Time : 2020/12/2 10:47
@Author :lixiaojun

*******************************************/
package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

type RefundOrder struct {
	client.Baseclient
}

func (r *RefundOrder) Init() *RefundOrder {
	r.ServiceName = "rpc_task"
	r.ServicePath = "Order/RefundOrder"
	return r
}

func (r *RefundOrder) SetRefundOrderId(ctx context.Context, refundOrderId *int, reply *bool) error {
	return r.Call(ctx, "SetRefundOrderId", refundOrderId, reply)
}
