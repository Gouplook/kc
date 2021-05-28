package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

/**
type Card interface {
	//添加综合卡
	AddCard(ctx context.Context, params *ArgsAddCard, replies *RepliesAddCard ) error
	//编辑综合卡
	EditCard(ctx context.Context, params *ArgsEditCard, replies *EmptyReplies) error
	//获取综合卡的详情
	CardInfo(ctx context.Context, params *ArgsCardInfo, replies *ReplyCardInfo ) error
	//获取总店的综合卡列表
	BusCardPage(ctx context.Context, params *ArgsBusCardPage, replies *ReplyCardPage ) error
	//设置适用门店
	SetCardShop(ctx context.Context, params *ArgsSetCardShop, replies *EmptyReplies) error
	//总店上下架综合卡
	DownUpCard(ctx context.Context, params *ArgsDownUpCard, replies *EmptyReplies) error
	//子店获取适用本店的综合卡列表
	ShopGetBusCardPage(ctx context.Context, params *ArgsShopGetBusCardPage, replies *ReplyCardPage ) error
	//子店添加综合卡到自己的店铺
	ShopAddCard(ctx context.Context, params *ArgsShopAddCard, replies *EmptyReplies) error
	//获取子店的综合卡列表
	ShopCardPage(ctx context.Context,  params *ArgsShopCardPage, replies *ReplyCardPage ) error
	//子店上下架综合卡
	ShopDownUpCard(ctx context.Context, params *ArgsShopDownUpCard, replies *EmptyReplies) error
}
*/
type Card struct {
	client.Baseclient
}

func (n *Card) Init() *Card {
	n.ServiceName = "rpc_cards"
	n.ServicePath = "Card"
	return n
}

//添加综合卡
func (n *Card) AddCard(ctx context.Context, params *cards.ArgsAddCard, replies *cards.RepliesAddCard) error {
	return n.Call(ctx, "AddCard", params, replies)
}

//编辑综合卡
func (n *Card) EditCard(ctx context.Context, params *cards.ArgsEditCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "EditCard", params, replies)
}

// DeleteCard 总店删除综合卡
func (n *Card) DeleteCard(ctx context.Context, params *cards.ArgsDeleteCard, replies *bool) error {

	return n.Call(ctx, "DeleteCard", params, replies)
}

//获取综合卡的详情
func (n *Card) CardInfo(ctx context.Context, params *cards.ArgsCardInfo, replies *cards.ReplyCardInfo) error {
	return n.Call(ctx, "CardInfo", params, replies)
}

//获取总店的综合卡列表
func (n *Card) BusCardPage(ctx context.Context, params *cards.ArgsBusCardPage, replies *cards.ReplyCardPage) error {
	return n.Call(ctx, "BusCardPage", params, replies)
}

//设置适用门店
func (n *Card) SetCardShop(ctx context.Context, params *cards.ArgsSetCardShop, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "SetCardShop", params, replies)
}

//总店上下架综合卡
func (n *Card) DownUpCard(ctx context.Context, params *cards.ArgsDownUpCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "DownUpCard", params, replies)
}

//子店获取适用本店的综合卡列表
func (n *Card) ShopGetBusCardPage(ctx context.Context, params *cards.ArgsShopGetBusCardPage, replies *cards.ReplyCardPage) error {
	return n.Call(ctx, "ShopGetBusCardPage", params, replies)
}

//子店添加综合卡到自己的店铺
func (n *Card) ShopAddCard(ctx context.Context, params *cards.ArgsShopAddCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "ShopAddCard", params, replies)
}

//获取子店的综合卡列表
func (n *Card) ShopCardPage(ctx context.Context, params *cards.ArgsShopCardPage, replies *cards.ReplyCardPage) error {
	return n.Call(ctx, "ShopCardPage", params, replies)
}

//子店上下架综合卡
func (n *Card) ShopDownUpCard(ctx context.Context, params *cards.ArgsShopDownUpCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "ShopDownUpCard", params, replies)
}

//获取门限可售综合卡详情
func (n *Card) CardsInfo(ctx context.Context, params *cards.ArgsCardsInfo, replies *map[int]*cards.ReplyCardsInfo) error {
	return n.Call(ctx, "CardsInfo", params, replies)
}

// 门店综合卡数据rpc内部调用
func (n *Card) ShopCardListRpc(ctx context.Context, params *cards.ArgsShopCardListRpc, replies *cards.ReplyShopCardListRpc) error {
	return n.Call(ctx, "ShopCardListRpc", params, replies)
}

// 获取所有卡的发布数量 rpc内部调用
func (n *Card) GetAllCardsNum(ctx context.Context, args *cards.ArgsAllCardsNum, reply *cards.ReplyAllCardsNum) error {
	return n.Call(ctx, "GetAllCardsNum", args, reply)
}

// 门店-删除本店的综合卡
func (n *Card) DeleteShopCard(ctx context.Context, args *cards.ArgsDeleteShopCard, replies *bool) error {
	return n.Call(ctx, "DeleteShopCard", args, replies)
}
