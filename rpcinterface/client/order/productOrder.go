package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

// @author yangzhiwu<578154898@qq.com>
// @date  2020/7/27 16:46

type ProductOrder struct {
	client.Baseclient
}

//初始化
func (p *ProductOrder) Init() *ProductOrder {
	p.ServiceName = "rpc_order"
	p.ServicePath = "ProductOrder"
	return p
}

//saas创建产品订单
func (p *ProductOrder) SaasCreateProductOrder(ctx context.Context, args *order.ArgsSaasCreateProductOrder, reply *order.ReplySaasCreateProductOrder ) error {
	return p.Call(ctx, "SaasCreateProductOrder", args, reply)
}

//用户购买产品订单
func (p *ProductOrder) UserCreateProductOrder(ctx context.Context, args *order.ArgsUserCreateProductOrder, reply *order.ReplyUserCreateProductOrder ) error {
	return p.Call(ctx, "UserCreateProductOrder", args, reply)
}

// 根据用户商品提货单号修改用户商品提货状态
func (p *ProductOrder) SetUserPickUpGoodsStatus(ctx context.Context, args *order.ArgsSetUserPickUpGoodsStatus, reply *order.ReplyBool) error {
	return p.Call(ctx, "SetUserPickUpGoodsStatus", args, reply)
}