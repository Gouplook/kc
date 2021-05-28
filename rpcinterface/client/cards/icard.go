package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//ICardClient 折扣卡管理（身份卡）
type ICardClient struct {
	client.Baseclient
}

//Init 初始化
func (c *ICardClient) Init() *ICardClient {
	c.ServiceName = "rpc_cards"
	c.ServicePath = "ICard"
	return c
}

//List 列表
func (c *ICardClient) List(ctx context.Context, args *cards.InputParamsICardList, reply *cards.OutputICardList) error {
	return c.Call(ctx, "List", args, reply)
}

//Save 保存
func (c *ICardClient) Save(ctx context.Context, args *cards.InputParamsICardSave, reply *cards.OutputParamsICardSave) error {
	return c.Call(ctx, "Save", args, reply)
}

//Info 详情
func (c *ICardClient) Info(ctx context.Context, args *cards.InputParamsICardInfo, reply *cards.OutputParamsICardInfo) error {
	return c.Call(ctx, "Info", args, reply)
}

//Delete 删除
func (c *ICardClient) Delete(ctx context.Context, args *cards.InputParamsDelete, reply *bool) error {
	return c.Call(ctx, "Delete", args, reply)
}

//Push 推送门店
func (c *ICardClient) Push(ctx context.Context, args *cards.InputParamsICardPush, reply *cards.OutputParamsICardPush) error {
	return c.Call(ctx, "Push", args, reply)
}

//OurShopList 已添加到门店列表
func (c *ICardClient) OurShopList(ctx context.Context, args *cards.InputParamsICardList, reply *cards.OutputICardList) error {
	return c.Call(ctx, "OurShopList", args, reply)
}

//ShopList 门店列表
func (c *ICardClient) ShopList(ctx context.Context, args *cards.InputParamsICardList, reply *cards.OutputICardList) error {
	return c.Call(ctx, "ShopList", args, reply)
}

//SetOnOff 上下架
func (c *ICardClient) SetOnOff(ctx context.Context, args *cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) error {
	return c.Call(ctx, "SetOnOff", args, reply)
}

//SetOn 上架
func (c *ICardClient) SetOn(ctx context.Context, args *cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) error {
	return c.Call(ctx, "SetOn", args, reply)
}

//SetOff 上下架
func (c *ICardClient) SetOff(ctx context.Context, args *cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) error {
	return c.Call(ctx, "SetOff", args, reply)
}

//ShopSetOnOff 门店上下架
func (c *ICardClient) ShopSetOnOff(ctx context.Context, args *cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) error {
	return c.Call(ctx, "ShopSetOnOff", args, reply)
}

//ShopSetOn 门店上架
func (c *ICardClient) ShopSetOn(ctx context.Context, args *cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) error {
	return c.Call(ctx, "ShopSetOn", args, reply)
}

//ShopSetOff 门店下架
func (c *ICardClient) ShopSetOff(ctx context.Context, args *cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) error {
	return c.Call(ctx, "ShopSetOff", args, reply)
}

//AddToShop 添加至本店
func (c *ICardClient) AddToShop(ctx context.Context, args *cards.InputParamsICardAddToShop, reply *cards.OutputParamsICardAddToShop) error {
	return c.Call(ctx, "AddToShop", args, reply)
}

//UserICardList 查看用户身份卡列表
func (c *ICardClient) UserICardList(ctx context.Context, args *cards.InputParams, reply *cards.OutputReply) error {
	return c.Call(ctx, "UserICardList", args, reply)
}

//CanUseICardList 查看用户身份卡列表
func (c *ICardClient) CanUseICardList(ctx context.Context, args *cards.InputParamsICardCanUse, reply *cards.OutputParamsICardCanUse) error {
	return c.Call(ctx, "CanUseICardList", args, reply)
}

//CanUseICardListForUser 查看用户身份卡列表
func (c *ICardClient) CanUseICardListForUser(ctx context.Context, args *cards.InputParamsICardCanUseForUser, reply *cards.OutputParamsICardCanUse) error {
	return c.Call(ctx, "CanUseICardListForUser", args, reply)
}

//获取iCard企业基本信息-风控统计用
func (c *ICardClient) GetBusBaseInfoRpc(ctx context.Context, iCardId *int, reply *cards.ReplyGetBusBaseInfoRpc) error {
	return c.Call(ctx, "GetBusBaseInfoRpc", iCardId, reply)
}

//获取身份卡的折扣信息
func (c *ICardClient) GetIcardDiscountById(ctx context.Context, iCardId *int, reply *cards.ReplyGetIcardDiscountById) error {
	return c.Call(ctx, "GetIcardDiscountById", iCardId, reply)
}

//获取身份卡备份表中的项目折扣
func (c *ICardClient) GetICardSingleDiscount(ctx context.Context, args *cards.ArgsGetICardSingleDiscount, reply *cards.ReplyGetICardSingleDiscount) error {
	return c.Call(ctx, "GetICardSingleDiscount", args, reply)
}
