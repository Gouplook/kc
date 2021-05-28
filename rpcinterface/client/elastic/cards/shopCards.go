package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	cards4 "git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/task/cards"
)

type ShopCards struct {
	client.Baseclient
}

func (s *ShopCards) Init()  *ShopCards {
	s.ServiceName = "rpc_elastic"
	s.ServicePath = "Cards/ShopCards"
	return s
}

//设置门店的项目到es文档
func (s *ShopCards)  SetItem(ctx context.Context, args *cards.ShopItems, reply *bool) error {
	return s.Call(ctx, "SetItem", args, reply)
}

//九百岁APP查询门店信息
func (s *ShopCards) SearchItem(ctx context.Context, args *cards4.ArgsAppInfos, reply *map[string]interface{}) error {
	return s.Call(ctx, "SearchItem", args, reply)
}

//根据ItemId查询门店信息
func (s *ShopCards) GetShopInfoByItemId(ctx context.Context, args *cards4.ArgsShopList, reply *map[string]interface{}) error {
	return s.Call(ctx, "GetShopInfoByItemId", args, reply)
}