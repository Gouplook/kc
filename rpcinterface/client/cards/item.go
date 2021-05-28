package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type Item struct {
	client.Baseclient
}

func (i *Item) Init() *Item {
	i.ServiceName = "rpc_cards"
	i.ServicePath = "Item"
	return i
}

//为es的shop-item文档获取项目信息
func (i *Item) ItemInfo4Es(ctx context.Context, args *cards.ArgsItemInfo4Es, reply *cards.ReplyItemInfo4Es) error {
	return i.Call(ctx, "ItemInfo4Es", args, reply)
}

//根据门店ids获取卡项信息
func (i *Item) GetItemsBySsids(ctx context.Context, args *cards.ArgsGetItemsBySsids, reply *map[cards.SsId]cards.ItemBase) error {
	return i.Call(ctx, "GetItemsBySsids", args, reply)
}

//根据条件查询九百岁
func (i *Item) GetInfos(ctx context.Context, args *cards.ArgsAppInfos, reply *cards.ReplyAppInfo) error {
	return i.Call(ctx, "GetInfos", args, reply)
}

//根据项目查询详情-适用门店详情
func (i *Item) GetDetailById(ctx context.Context, args *cards.ArgsShopList, reply *[]order.ReplyCableShopInfo) error {
	return i.Call(ctx, "GetDetailById", args, reply)
}

// 门店拥有的项目-上架的项目
func (i *Item) GetItemsByShopId(ctx context.Context, args *cards.ArgsGetItemsByShopId, reply *cards.ReplyGetItemsByShopId) error {
	return i.Call(ctx, "GetItemsByShopId", args, reply)
}

//门店详情-推荐的单项目
func (i *Item) GetRecommendSingles(ctx context.Context, args *cards.ArgsGetRecommendSingles, reply *cards.ReplyGetRecommendSingles) error {
	return i.Call(ctx, "GetRecommendSingles", args, reply)
}

//卡项收藏入参
func (i *Item) CollectItems(ctx context.Context, args *cards.ArgsCollectItems, reply *bool) error {
	return i.Call(ctx, "CollectItems", args, reply)
}

//获取用户收藏的卡项入参
func (i *Item) GetCollectItems(ctx context.Context, args *cards.ArgsGetCollectItems, reply *cards.ReplyGetCollectItems) error {
	return i.Call(ctx, "GetCollectItems", args, reply)
}

//卡项收藏状态
func (i *Item) GetCollectStatus(ctx context.Context, args *cards.ArgsCollectStatus, reply *cards.ReplyCollectStatus) error {
	return i.Call(ctx, "GetCollectStatus", args, reply)
}

//获取当月下架的卡项
func (i *Item) GetItemAllXCradsNumByShopId(ctx context.Context, shopId *int, reply *cards.ReplyGetItemAllXCradsNumByShopId) error {
	return i.Call(ctx, "GetItemAllXCradsNumByShopId", shopId, reply)
}

//购买成功，设置卡项的销量
func (i *Item) IncrItemSales(ctx context.Context, orderSn *string, reply *bool) error {
	return i.Call(ctx, "IncrItemSales", orderSn, reply)
}

//获取卡项下的商品信息
func (i *Item) GetCardProductsInfo(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetCardProductsInfo) error {
	return i.Call(ctx, "GetCardProductsInfo", args, reply)
}

//预付卡包含的单项目
func (i *Item) GetItemIncludeSingles(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetItemIncludeSingles) error {
	return i.Call(ctx, "GetItemIncludeSingles", args, reply)
}

//预付卡赠送的单项目
func (i *Item) GetItemGiveSingles(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetItemGiveSingles) error {
	return i.Call(ctx, "GetItemGiveSingles", args, reply)
}

//预付卡默认图片
func (i *Item) GetItemDefaultImgs(ctx context.Context, args *cards.ArgsGetItemDefaultImgs, reply *cards.ReplyGetItemDefaultImgs) error {
	return i.Call(ctx, "GetItemDefaultImgs", args, reply)
}
