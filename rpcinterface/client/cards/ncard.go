package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

/**
type NCard interface {
	//添加限次卡
	AddNCard(ctx context.Context, params *ArgsAddNCard, replies *RepliesAddNCard ) error
	//编辑限次卡
	EditNCard(ctx context.Context, params *ArgsEditNCard, replies *EmptyReplies) error
	//获取限次卡的详情
	NCardInfo(ctx context.Context, params *ArgsNCardInfo, replies *ReplyNCardInfo ) error
	//获取总店的限次卡列表
	BusNCardPage(ctx context.Context, params *ArgsBusNCardPage, replies *ReplyNCardPage ) error
	//设置适用门店
	SetNCardShop(ctx context.Context, params *ArgsSetNCardShop, replies *EmptyReplies) error
	//总店上下架限次卡
	DownUpNCard(ctx context.Context, params *ArgsDownUpNCard, replies *EmptyReplies) error
	//子店获取适用本店的限次卡列表
	ShopGetBusNCardPage(ctx context.Context, params *ArgsShopGetBusNCardPage, replies *ReplyNCardPage ) error
	//子店添加限次卡到自己的店铺
	ShopAddNCard(ctx context.Context, params *ArgsShopAddNCard, replies *EmptyReplies) error
	//获取子店的限次卡列表
	ShopNCardPage(ctx context.Context,  params *ArgsShopNCardPage, replies *ReplyNCardPage ) error
	//子店上下架限次卡
	ShopDownUpNCard(ctx context.Context, params *ArgsShopDownUpNCard, replies *EmptyReplies) error
}
*/
type NCard struct {
	client.Baseclient
}

func (n *NCard) Init() *NCard {
	n.ServiceName = "rpc_cards"
	n.ServicePath = "NCard"
	return n
}

//添加限次卡
func (n *NCard) AddNCard(ctx context.Context, params *cards.ArgsAddNCard, replies *cards.RepliesAddNCard) error {
	return n.Call(ctx, "AddNCard", params, replies)
}

//编辑限次卡
func (n *NCard) EditNCard(ctx context.Context, params *cards.ArgsEditNCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "EditNCard", params, replies)
}

//获取限次卡的详情
func (n *NCard) NCardInfo(ctx context.Context, params *cards.ArgsNCardInfo, replies *cards.ReplyNCardInfo) error {
	return n.Call(ctx, "NCardInfo", params, replies)
}

//获取总店的限次卡列表
func (n *NCard) BusNCardPage(ctx context.Context, params *cards.ArgsBusNCardPage, replies *cards.ReplyNCardPage) error {
	return n.Call(ctx, "BusNCardPage", params, replies)
}

//设置适用门店
func (n *NCard) SetNCardShop(ctx context.Context, params *cards.ArgsSetNCardShop, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "SetNCardShop", params, replies)
}

//总店上下架限次卡
func (n *NCard) DownUpNCard(ctx context.Context, params *cards.ArgsDownUpNCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "DownUpNCard", params, replies)
}

//子店获取适用本店的限次卡列表
func (n *NCard) ShopGetBusNCardPage(ctx context.Context, params *cards.ArgsShopGetBusNCardPage, replies *cards.ReplyNCardPage) error {
	return n.Call(ctx, "ShopGetBusNCardPage", params, replies)
}

//子店添加限次卡到自己的店铺
func (n *NCard) ShopAddNCard(ctx context.Context, params *cards.ArgsShopAddNCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "ShopAddNCard", params, replies)
}

//获取子店的限次卡列表
func (n *NCard) ShopNCardPage(ctx context.Context, params *cards.ArgsShopNCardPage, replies *cards.ReplyNCardPage) error {
	return n.Call(ctx, "ShopNCardPage", params, replies)
}

//子店上下架限次卡
func (n *NCard) ShopDownUpNCard(ctx context.Context, params *cards.ArgsShopDownUpNCard, replies *cards.EmptyReplies) error {
	return n.Call(ctx, "ShopDownUpNCard", params, replies)
}

//ShopNcardRpc
func (n *NCard) ShopNcardRpc(ctx context.Context, params *cards.ArgsShopNcardRpc, replies *cards.ReplyShopNcardRpc) error {
	return n.Call(ctx, "ShopNcardRpc", params, replies)
}

//总店-删除
func (n *NCard) DeleteNCard(ctx context.Context, params *cards.ArgsDeleteNCard, reply *bool) error {
	return n.Call(ctx, "DeleteNCard", params, reply)
}

//分店-删除
func (n *NCard) DeleteShopNCard(ctx context.Context, params *cards.ArgsDeleteShopNCard, reply *bool) error {
	return n.Call(ctx, "DeleteShopNCard", params, reply)
}
