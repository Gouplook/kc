package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type RcardOrder struct {
	client.Baseclient
}

func (r *RcardOrder)Init()*RcardOrder  {
	r.ServiceName = "rpc_order"
	r.ServicePath = "RcardOrder"
	return r
}

//sass创建充值订单
func (r *RcardOrder)SaasCreateRechargeOrder(ctx context.Context,args *order.ArgsSaasCreateRechargeOrder,reply *order.ReplyCreateRechargeOrder)error  {
	return r.Call(ctx, "SaasCreateRechargeOrder", args, reply)
}

//用户创建充值订单
func (r *RcardOrder)UserCreateRechargeOrder(ctx context.Context,args *order.ArgsUserCreateRechargeOrder,reply *order.ReplyCreateRechargeOrder)error  {
	return r.Call(ctx, "UserCreateRechargeOrder", args, reply)
}