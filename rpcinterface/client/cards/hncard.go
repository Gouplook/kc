package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

/**
type HNCard interface {
	//添加限时限次卡
	AddHNCard(ctx context.Context, params *ArgsAddHNCard, replies *RepliesAddHNCard ) error
	//编辑限时限次卡
	EditHNCard(ctx context.Context, params *ArgsEditHNCard, replies *EmptyReplies) error
	//获取限时限次卡的详情
	HNCardInfo(ctx context.Context, params *ArgsHNCardInfo, replies *ReplyHNCardInfo ) error
	//获取总店的限时限次卡列表
	BusHNCardPage(ctx context.Context, params *ArgsBusHNCardPage, replies *ReplyHNCardPage ) error
	//设置适用门店
	SetHNCardShop(ctx context.Context, params *ArgsSetHNCardShop, replies *EmptyReplies) error
	//总店上下架限时限次卡
	DownUpHNCard(ctx context.Context, params *ArgsDownUpHNCard, replies *EmptyReplies) error
	//子店获取适用本店的限时限次卡列表
	ShopGetBusHNCardPage(ctx context.Context, params *ArgsShopGetBusHNCardPage, replies *ReplyHNCardPage ) error
	//子店添加限时限次卡到自己的店铺
	ShopAddHNCard(ctx context.Context, params *ArgsShopAddHNCard, replies *EmptyReplies) error
	//获取子店的限时限次卡列表
	ShopHNCardPage(ctx context.Context,  params *ArgsShopHNCardPage, replies *ReplyHNCardPage ) error
	//子店上下架限时限次卡
	ShopDownUpHNCard(ctx context.Context, params *ArgsShopDownUpHNCard, replies *EmptyReplies) error
}
 */
type HNCard struct {
	client.Baseclient
}
func (n *HNCard) Init() *HNCard {
	n.ServiceName = "rpc_cards"
	n.ServicePath = "HNCard"
	return n
}

//添加限时限次卡
func (n *HNCard)AddHNCard(ctx context.Context, params *cards.ArgsAddHNCard, replies *cards.RepliesAddHNCard ) error{
	return n.Call(ctx, "AddHNCard", params, replies)
}
//编辑限时限次卡
func (n *HNCard)EditHNCard(ctx context.Context, params *cards.ArgsEditHNCard, replies *cards.EmptyReplies) error{
	return n.Call(ctx, "EditHNCard", params, replies)
}
//获取限时限次卡的详情
func (n *HNCard)HNCardInfo(ctx context.Context, params *cards.ArgsHNCardInfo, replies *cards.ReplyHNCardInfo ) error{
	return n.Call(ctx, "HNCardInfo", params, replies)
}
//获取总店的限时限次卡列表
func (n *HNCard)BusHNCardPage(ctx context.Context, params *cards.ArgsBusHNCardPage, replies *cards.ReplyHNCardPage ) error{
	return n.Call(ctx, "BusHNCardPage", params, replies)
}
//设置适用门店
func (n *HNCard)SetHNCardShop(ctx context.Context, params *cards.ArgsSetHNCardShop, replies *cards.EmptyReplies) error{
	return n.Call(ctx, "SetHNCardShop", params, replies)
}
//总店上下架限时限次卡
func (n *HNCard)DownUpHNCard(ctx context.Context, params *cards.ArgsDownUpHNCard, replies *cards.EmptyReplies) error{
	return n.Call(ctx, "DownUpHNCard", params, replies)
}
//子店获取适用本店的限时限次卡列表
func (n *HNCard)ShopGetBusHNCardPage(ctx context.Context, params *cards.ArgsShopGetBusHNCardPage, replies *cards.ReplyHNCardPage ) error{
	return n.Call(ctx, "ShopGetBusHNCardPage", params, replies)
}
//子店添加限时限次卡到自己的店铺
func (n *HNCard)ShopAddHNCard(ctx context.Context, params *cards.ArgsShopAddHNCard, replies *cards.EmptyReplies) error{
	return n.Call(ctx, "ShopAddHNCard", params, replies)
}
//获取子店的限时限次卡列表
func (n *HNCard)ShopHNCardPage(ctx context.Context,  params *cards.ArgsShopHNCardPage, replies *cards.ReplyHNCardPage ) error{
	return n.Call(ctx, "ShopHNCardPage", params, replies)
}
//子店上下架限时限次卡
func (n *HNCard)ShopDownUpHNCard(ctx context.Context, params *cards.ArgsShopDownUpHNCard, replies *cards.EmptyReplies) error{
	return n.Call(ctx, "ShopDownUpHNCard", params, replies)
}
// 子店限时限次卡-rpc
func (n *HNCard)ShopHncardRpc(ctx context.Context,params *cards.ArgsShopHncardRpc,replies *cards.ReplyShopHncardRpc)error{
	return n.Call(ctx, "ShopHncardRpc", params, replies)
}

// 总店删除限时限次卡
func (n *HNCard)DeleteHNCard(ctx context.Context, args *cards.ArgsDelHNCard, reply *bool) error {
	return n.Call(ctx, "DeleteHNCard", args, reply)
}
// 子店删除限时限次卡
func (n *HNCard)ShopDeleteHNCard(ctx context.Context, args *cards.ArgsDelHNCard,reply *bool)error {
	return n.Call(ctx, "ShopDeleteHNCard", args, reply)
}