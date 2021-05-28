package cards

import (
	"context"
	cards2 "git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/task/cards"
)



//门店卡项
type ShopCards interface {
	//设置门店的卡项
	SetItem(ctx context.Context, args *cards.ShopItems, reply *bool) error

	//九百岁APP查询门店信息
	SearchItem(ctx context.Context, args *cards2.ArgsAppInfos, reply *map[string]interface{}) error

	//根据ItemId查询门店信息
	GetShopInfoByItemId(ctx context.Context, args *cards2.ArgsShopList, reply *map[string]interface{}) error
}