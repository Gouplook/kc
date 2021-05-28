package risk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/bus"
	"git.900sui.cn/kc/rpcinterface/interface/task"
)

type BusBasicIndusAvg struct {
	client.Baseclient
}

func (b *BusBasicIndusAvg)Init()*BusBasicIndusAvg  {
	b.ServiceName = "rpc_risk"
	b.ServicePath = "BusBasicIndusAvg"
	return b
}
//审核商户
func (b *BusBasicIndusAvg)RiskForAuditBus(ctx context.Context,busId *int,reply *bool)error  {
	return b.Call(ctx, "RiskForAuditBus", busId, reply)
}
//审核店铺
func (b *BusBasicIndusAvg)RiskForAuditShop(ctx context.Context,shopId *int,reply *bool)error{
	return b.Call(ctx, "RiskForAuditShop", shopId, reply)
}
//员工新增、离职、删除
func (b *BusBasicIndusAvg)RiskForSetStaff(ctx context.Context,staffId *int,reply *bool) error {
	return b.Call(ctx, "RiskForSetStaff", staffId,reply)
}
//会员新增
func (b *BusBasicIndusAvg)RiskForAddMember(ctx context.Context,memberId *int,reply *bool)error{
	return b.Call(ctx, "RiskForAddMember", memberId,reply)
}
//订单支付成功
func (b *BusBasicIndusAvg)RiskForPaySuccess(ctx context.Context,orderSn *string, reply *bool)error{
	return b.Call(ctx, "RiskForPaySuccess",orderSn,reply)
}
//店铺面积更新
func (b *BusBasicIndusAvg)RiskForShopAreaUpdate(ctx context.Context,args *bus.ArgsShopAreaUpdate, reply *bool)error{
	return b.Call(ctx, "RiskForShopAreaUpdate",args,reply)
}
//新增删除卡项/商品
func (b *BusBasicIndusAvg)RiskForAddDelGoods(ctx context.Context,args *task.ArgsAddDelGoods ,reply *bool)error{
	return b.Call(ctx, "RiskForAddDelGoods",args,reply)
}

