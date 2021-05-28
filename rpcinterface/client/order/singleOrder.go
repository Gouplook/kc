package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

// @author yangzhiwu<578154898@qq.com>
// @date  2020/7/27 16:46

type SingleOrder struct {
	client.Baseclient
}

//初始化
func (s *SingleOrder) Init() *SingleOrder {
	s.ServiceName = "rpc_order"
	s.ServicePath = "SingleOrder"
	return s
}

//saas 创建订单
func (s *SingleOrder) SaasCreateSingleOrder(ctx context.Context, args *order.ArgsSaasCreateSingleOrder, reply *order.ReplySaasCreateSingleOrder ) error  {
	return s.Call(ctx, "SaasCreateSingleOrder", args, reply)
}

//前端用户创建订单
func (s *SingleOrder )  UserCreateSingleOrder(ctx context.Context, args *order.ArgsUserCreateSingleOrder, reply *order.ReplyUserCreateSingleOrder ) error {
	return s.Call(ctx, "UserCreateSingleOrder", args, reply)
}