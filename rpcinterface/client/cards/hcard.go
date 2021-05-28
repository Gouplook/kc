package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

// 现时卡 RPC Client

// Hcard Hcard
type Hcard struct {
	client.Baseclient
}

// Init Init
func (h *Hcard) Init() *Hcard {
	h.ServiceName = "rpc_cards"
	h.ServicePath = "Hcard"
	return h
}

// AddHcard 总店添加限时卡
func (h *Hcard) AddHcard(ctx context.Context, args *cards.ArgsAddHcard, reply *cards.ReplyAddHcard) error {
	return h.Call(ctx, "AddHcard", args, reply)
}

// EditHcard 总店修改限时卡
func (h *Hcard) EditHcard(ctx context.Context, args *cards.ArgsEditHcard, reply *bool) error {
	return h.Call(ctx, "EditHcard", args, reply)
}

// DeleteHcard 总店删除限时卡
func (h *Hcard) DeleteHcard(ctx context.Context, args *cards.ArgsDeleteHcard, reply *bool) error {
	return h.Call(ctx, "DeleteHcard", args, reply)
}

// HcardInfo 获取限时卡详情(总店和分店共用)
func (h *Hcard) HcardInfo(ctx context.Context, args *cards.ArgsHcardInfo, reply *cards.ReplyHcardInfo) error {
	return h.Call(ctx, "HcardInfo", args, reply)
}

// BusHcardPage 获取总店限时卡列表
func (h *Hcard) BusHcardPage(ctx context.Context, args *cards.ArgsBusHcardPage, reply *cards.ReplyHcardPage) error {
	return h.Call(ctx, "BusHcardPage", args, reply)
}

// SetHcardShop 设置总店限时卡的适用门店
func (h *Hcard) SetHcardShop(ctx context.Context, args *cards.ArgsSetHcardShop, reply *bool) error {
	return h.Call(ctx, "SetHcardShop", args, reply)
}

// DownUpHcard 总店上下架限时卡
func (h *Hcard) DownUpHcard(ctx context.Context, args *cards.ArgsDownUpHcard, reply *bool) error {
	return h.Call(ctx, "DownUpHcard", args, reply)
}

// ShopGetBusHcardPage 子店获取适用本店的限时卡列表(总店分配给自己店铺的限时卡)
func (h *Hcard) ShopGetBusHcardPage(ctx context.Context, args *cards.ArgsShopGetBusHcardPage, reply *cards.ReplyShopGetBusHcardPage) error {
	return h.Call(ctx, "ShopGetBusHcardPage", args, reply)
}

// ShopHcardPage 子店限时卡列表
func (h *Hcard) ShopHcardPage(ctx context.Context, args *cards.ArgsShopHcardPage, reply *cards.ReplyHcardPage) error {
	return h.Call(ctx, "ShopHcardPage", args, reply)
}

// ShopAddHcard 子店添加总部限时卡到自己的店铺
func (h *Hcard) ShopAddHcard(ctx context.Context, args *cards.ArgsShopAddHcard, reply *bool) error {
	return h.Call(ctx, "ShopAddHcard", args, reply)
}

// ShopDownUpHcard 子店上下架自己店铺中的限时卡
func (h *Hcard) ShopDownUpHcard(ctx context.Context, args *cards.ArgsShopDownUpHcard, reply *bool) error {
	return h.Call(ctx, "ShopDownUpHcard", args, reply)
}

// ShopHcardListRpc 子店现时卡列表rpc内部调用
func (h *Hcard) ShopHcardListRpc(ctx context.Context, args *cards.ArgsShopCardListRpc, reply *cards.ReplyShopHcardListRpc) error {
	return h.Call(ctx, "ShopHcardListRpc", args, reply)
}

// 子店删除限时卡
func (h *Hcard)ShopDeleteHcard(ctx context.Context, args *cards.ArgsDeleteHcard, reply *bool) error {
	return h.Call(ctx, "ShopDeleteHcard", args, reply)
}
