package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/cards"
)

//门店项目

type ShopItems struct {
	client.Baseclient
}

// Init Init
func (s *ShopItems) Init() *ShopItems {
	s.ServiceName = "rpc_task"
	s.ServicePath = "Cards/ShopItems"
	return s
}

//设置门店的项目 到交换机中
func (s *ShopItems) SetItems(ctx context.Context, args *cards.ShopItems, reply *bool) error {
	return s.Call(ctx, "SetItems", args, reply)
}

//身份卡折扣同步 到交换机中
func (s *ShopItems) SetIcardDiscount(ctx context.Context, icardId *int, reply *bool) error {
	return s.Call(ctx, "SetIcardDiscount", icardId, reply)
}
