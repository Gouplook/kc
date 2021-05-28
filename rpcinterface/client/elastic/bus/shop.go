package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

// @author liyang<654516092@qq.com>
// @date  2020/6/9 15:50

type Shop struct {
	client.Baseclient
}

func (s *Shop)Init()*Shop{
	s.ServiceName="rpc_elastic"
	s.ServicePath="Bus/Shop"
	return s
}

//新增、更新门店基本信息
func (s *Shop) SetShop(ctx context.Context, shopId *int, reply *bool) error {
	return s.Call(ctx, "SetShop", shopId, reply)
}

//设置门店的销量、平均分、平均消费价格
func (s *Shop) SetShopProperty(ctx context.Context, shopId *int, reply *bool) error {
	return s.Call(ctx, "SetShopProperty", shopId, reply)
}

//更新门店status
func (s *Shop)SetShopStatus(ctx context.Context, shopId *int, reply *bool)error  {
	return s.Call(ctx, "SetShopStatus", shopId, reply)
}

//更新门店安全码颜色
func (s *Shop)SetShopSafeCode(ctx context.Context, busId *int, reply *bool)error  {
	return s.Call(ctx, "SetShopSafeCode", busId, reply)
}

//根据门店id查询es 门店信息
func (s *Shop) GetShopInfo(ctx context.Context,shopId int, reply *news.ShopInfo) error {
	return s.Call(ctx, "GetShopInfo", shopId, reply)
}