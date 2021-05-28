//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/15 14:44
package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Sm struct {
	client.Baseclient
}

func (s *Sm) Init() *Sm {
	s.ServiceName = "rpc_cards"
	s.ServicePath = "Sm"
	return s
}

//添加套餐
func (s *Sm) AddSm(ctx context.Context, sm *cards.ArgsAddSm, smId *int) error {
	return s.Call(ctx, "AddSm", sm, smId)
}

//编辑套餐信息
func (s *Sm) EditSm(ctx context.Context, sm *cards.ArgsEditSm, reply *bool ) error{
	return s.Call( ctx, "EditSm", sm, reply )
}

//获取套餐的详情
func (s *Sm) SmInfo(ctx context.Context, args *cards.ArgsSmInfo, reply *cards.ReplySmInfo) error {
	return s.Call(ctx, "SmInfo", args, reply)
}

//获取商家的套餐列表
func (s *Sm) BusSmPage(ctx context.Context, args *cards.ArgsBusSmPage, reply *cards.ReplySmPage) error {
	return s.Call(ctx, "BusSmPage", args, reply)
}

//设置套餐适用门店
func (s *Sm) SetSmShop(ctx context.Context, args *cards.ArgsSetSmShop, reply *bool) error {
	return s.Call(ctx, "SetSmShop", args, reply)
}

//总店上下架套餐
func (s *Sm) DownUpSm(ctx context.Context, args *cards.ArgsDownUpSm, reply *bool) error {
	return s.Call(ctx, "DownUpSm", args, reply)
}

//子店获取适用本店的套餐列表
func (s *Sm) ShopGetBusSmPage(ctx context.Context, args *cards.ArgsShopGetBusSmPage, reply *cards.ReplyShopGetBusSmPage) error {
	return s.Call(ctx, "ShopGetBusSmPage", args, reply)
}

//子店添加套餐到自己的店铺
func (s *Sm) ShopAddSm(ctx context.Context, args *cards.ArgsShopAddSm, reply *bool) error {
	return s.Call(ctx, "ShopAddSm", args, reply)
}

//获取子店的套餐列表
func (s *Sm) ShopSmPage(ctx context.Context, args *cards.ArgsShopSmPage, reply *cards.ReplyShopSmPage) error {
	return s.Call(ctx, "ShopSmPage", args, reply)
}

//子店上下架套餐
func (s *Sm) ShopDownUpSm(ctx context.Context, args *cards.ArgsShopDownUpSm, reply *bool) error {
	return s.Call(ctx, "ShopDownUpSm", args, reply)
}

//总店-删除套餐
func (s *Sm) DeleteSm(ctx context.Context, args *cards.ArgsDeleteSm, reply *bool) error {
	return s.Call(ctx, "DeleteSm", args, reply)
}

//分店-删除套餐
func (s *Sm) DeleteShopSm(ctx context.Context, args *cards.ArgsDeleteShopSm, reply *bool) error {
	return s.Call(ctx, "DeleteShopSm", args, reply)
}
// 子店添加套餐（一期优化，去掉总店推送，改为门店自动拉取)
func (s *Sm)ShopAddToSm(ctx context.Context,args *cards.ArgsShopAddSm,reply *bool) error {
	return s.Call(ctx, "ShopAddToSm", args, reply)
}
