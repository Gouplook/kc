package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/bus"
)

// @author liyang<654516092@qq.com>
// @date  2020/6/9 16:25
// Bus Bus
type Shop struct {
	client.Baseclient
}

// Init Init
func (s *Shop) Init() *Shop {
	s.ServiceName = "rpc_task"
	s.ServicePath = "Bus/Shop"
	return s
}

//实现将分店信息上传到交换机中
func (s *Shop) SetShop(ctx context.Context, shopId *int, reply *bool) error {
	return s.Call(ctx, "SetShop", shopId, reply)
}

//实现将分店信息上传到交换机中,更新门店的销量、平均分、平均消费价格到es中
func (s *Shop)SetShopProperty(ctx context.Context, shopId *int, reply *bool)error  {
	return s.Call(ctx, "SetShopProperty", shopId, reply)
}

//实现将分店状态更新到es中
func (s *Shop)SetShopStatus(ctx context.Context, shopId *int, reply *bool)error  {
	return s.Call(ctx, "SetShopStatus", shopId, reply)
}

//门店审核成功风控统计
func (s *Shop)ShopAudit(ctx context.Context, shopId *int, reply *bool)error  {
	return s.Call(ctx, "ShopAudit", shopId, reply)
}
//分店面积更新上传到交换机做风控统计
func (s *Shop)ShopAreaUpdate(ctx context.Context, args *bus.ArgsShopAreaUpdate , reply *bool) error{
	return s.Call(ctx, "ShopAreaUpdate", args, reply)
}
