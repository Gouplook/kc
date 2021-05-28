package service

import (
	"context"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type Item struct {
}

//为es的shop-item文档获取项目信息
func (i *Item) ItemInfo4Es(ctx context.Context, args *cards.ArgsItemInfo4Es, reply *cards.ReplyItemInfo4Es) (err error) {
	mItem := new(logics.ItemLogic)
	*reply, err = mItem.GetItemInfo4Es(args.ShopId, args.ItemId, args.ItemType)
	return
}

//根据卡项在门店的ids，获取卡项数据
func (i *Item) GetItemsBySsids(ctx context.Context, args *cards.ArgsGetItemsBySsids, reply *map[cards.SsId]cards.ItemBase) (err error) {
	mItem := new(logics.ItemLogic)
	*reply, err = mItem.GetItemsBySsids(ctx, args)
	return
}

//根据条件查询九百岁
func (i *Item) GetInfos(ctx context.Context, args *cards.ArgsAppInfos, reply *cards.ReplyAppInfo) error {
	mItem := new(logics.ItemLogic)
	res, err := mItem.GetInfos(ctx, args)
	if err != nil {
		return err
	}
	*reply = *res
	return nil
}

//根据项目查询详情-适用门店详情
func (i *Item) GetDetailById(ctx context.Context, args *cards.ArgsShopList, reply *[]order.ReplyCableShopInfo) error {
	return new(logics.ItemLogic).GetDetailById(ctx, args, reply)
}

// 门店服务项目/各类预付卡
func (i *Item) GetItemsByShopId(ctx context.Context, args *cards.ArgsGetItemsByShopId, reply *cards.ReplyGetItemsByShopId) (err error) {
	return new(logics.ItemLogic).GetItemsByShopId(ctx, args, reply)
}

//门店详情-推荐的单项目
func (i *Item) GetRecommendSingles(ctx context.Context, args *cards.ArgsGetRecommendSingles, reply *cards.ReplyGetRecommendSingles) (err error) {

	req := &cards.ArgsGetItemsByShopId{Paging: common.Paging{Page: 1, PageSize: args.TotalNum}, ItemType: cards.ITEM_TYPE_single, ShopId: args.ShopId, OrderBy: "sales"}
	var res cards.ReplyGetItemsByShopId
	if err = new(logics.ItemLogic).GetItemsByShopId(ctx, req, &res); err != nil {
		return
	}
	if len(res.Lists) > 0 {
		reply.Lists = res.Lists
		reply.IndexImg = res.IndexImg
	}
	return
}

//购买成功，设置卡项的销量
func (i *Item) IncrItemSales(ctx context.Context, orderSn *string, reply *bool) error {
	return new(logics.ItemLogic).IncrItemSales(ctx, *orderSn, reply)
}

//卡项收藏入参
func (i *Item) CollectItems(ctx context.Context, args *cards.ArgsCollectItems, reply *bool) (err error) {
	uid, err := args.GetUid()
	if err != nil {
		return
	}
	if args.SsId == 0 {
		return toolLib.CreateKcErr(_const.PARAM_ERR)
	}
	args.Uid = uid
	return new(logics.ItemLogic).CollectItems(ctx, args, reply)
}

//卡项收藏状态
func (i *Item) GetCollectStatus(ctx context.Context, args *cards.ArgsCollectStatus, reply *cards.ReplyCollectStatus) (err error) {
	args.Uid, err = args.GetUid()
	if err != nil {
		return
	}
	if args.SsId == 0 || args.ItemId == 0 {
		return toolLib.CreateKcErr(_const.PARAM_ERR)
	}
	return new(logics.ItemLogic).GetCollectStatus(ctx, args, reply)
}

//获取用户收藏的卡项入参
func (i *Item) GetCollectItems(ctx context.Context, args *cards.ArgsGetCollectItems, reply *cards.ReplyGetCollectItems) (err error) {
	uid, err := args.GetUid()
	if err != nil {
		return
	}
	args.Uid = uid
	return new(logics.ItemLogic).GetCollectItems(ctx, args, reply)
}

func (i *Item) GetItemAllXCradsNumByShopId(ctx context.Context, shopId *int, reply *cards.ReplyGetItemAllXCradsNumByShopId) (err error) {
	return new(logics.ItemLogic).GetItemAllXCradsNumByShopId(ctx, *shopId, reply)
}

//获取卡项下的商品信息
func (i *Item) GetCardProductsInfo(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetCardProductsInfo) error {
	return new(logics.ItemLogic).GetCardProductsInfo(ctx, args, reply)
}

//预付卡包含的单项目
func (i *Item) GetItemIncludeSingles(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetItemIncludeSingles) (err error) {
	return new(logics.ItemLogic).GetItemIncludeSingles(ctx, args, reply)
}

//预付卡赠送的单项目
func (i *Item) GetItemGiveSingles(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetItemGiveSingles) error {
	return new(logics.ItemLogic).GetItemGiveSingles(ctx, args, reply)
}

//预付卡默认图片
func (i *Item) GetItemDefaultImgs(ctx context.Context, args *cards.ArgsGetItemDefaultImgs, reply *cards.ReplyGetItemDefaultImgs) (err error) {
	return new(logics.ItemLogic).GetItemDefaultImgs(ctx, args, reply)
}
