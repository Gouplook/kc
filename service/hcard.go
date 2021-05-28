package service

import (
	"context"

	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

// 添加 hcard service-------------------

// Hcard Hcard service
type Hcard struct {
}

// 添加rpc方法

// AddHcard 总店添加限时卡
func (h *Hcard) AddHcard(ctx context.Context, args *cards.ArgsAddHcard, reply *cards.ReplyAddHcard) (err error) {
	hcardLogic := new(logics.HcardLogic)
	hcardID, err := hcardLogic.AddHcard(ctx, args)
	reply.HcardID = hcardID
	return
}

// EditHcard 总店修改限时卡
func (h *Hcard) EditHcard(ctx context.Context, args *cards.ArgsEditHcard, reply *bool) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply = true
	if err = hcardLogic.EditHcard(ctx, args); err != nil {
		*reply = false
	}
	return
}

// DeleteHcard 总店删除限时卡
func (h *Hcard) DeleteHcard(ctx context.Context, args *cards.ArgsDeleteHcard, reply *bool) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply = true
	if err = hcardLogic.DeleteHcard(ctx, args); err != nil {
		*reply = false
	}
	return
}

// HcardInfo 获取限时卡详情(总店和分店共用)
func (h *Hcard) HcardInfo(ctx context.Context, args *cards.ArgsHcardInfo, reply *cards.ReplyHcardInfo) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply, err = hcardLogic.HcardInfo(ctx, args.HcardID, args.ShopID)
	return
}

// BusHcardPage 获取总店限时卡列表
func (h *Hcard) BusHcardPage(ctx context.Context, args *cards.ArgsBusHcardPage, reply *cards.ReplyHcardPage) (err error) {
	hcardLogic := new(logics.HcardLogic)
	start, limit := args.GetStart(), args.GetPageSize()
	*reply, err = hcardLogic.BusHcardPage(ctx, args.BusID, args.ShopId, start, limit, args.IsGround,args.FilterShopHasAdd)
	return
}

// SetHcardShop 设置总店限时卡的适用门店
func (h *Hcard) SetHcardShop(ctx context.Context, args *cards.ArgsSetHcardShop, reply *bool) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply = true
	if err = hcardLogic.SetHcardShop(ctx, args); err != nil {
		*reply = false
	}
	return
}

// DownUpHcard 总店上下架限时卡
func (h *Hcard) DownUpHcard(ctx context.Context, args *cards.ArgsDownUpHcard, reply *bool) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply = true
	if err = hcardLogic.DownUpHcard(ctx, args); err != nil {
		*reply = false
	}
	return
}

// ShopGetBusHcardPage 子店获取适用本店的限时卡列表(总店分配给自己店铺的限时卡)
func (h *Hcard) ShopGetBusHcardPage(ctx context.Context, args *cards.ArgsShopGetBusHcardPage, reply *cards.ReplyShopGetBusHcardPage) (err error) {
	hcardLogic := new(logics.HcardLogic)
	start, limit := args.GetStart(), args.GetPageSize()
	*reply, err = hcardLogic.ShopGetBusHcardPage(ctx, args.BusID, args.ShopID, start, limit)
	return
}

// ShopHcardPage 子店限时卡列表
func (h *Hcard) ShopHcardPage(ctx context.Context, args *cards.ArgsShopHcardPage, reply *cards.ReplyHcardPage) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply, err = hcardLogic.ShopHcardPage(ctx, args)
	return
}

// ShopAddHcard 子店添加总部限时卡到自己的店铺
func (h *Hcard) ShopAddHcard(ctx context.Context, args *cards.ArgsShopAddHcard, reply *bool) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply = true
	if err = hcardLogic.ShopAddHcard(ctx, args); err != nil {
		*reply = false
	}
	return
}

// ShopDownUpHcard 子店上下架自己店铺中的限时卡
func (h *Hcard) ShopDownUpHcard(ctx context.Context, args *cards.ArgsShopDownUpHcard, reply *bool) (err error) {
	hcardLogic := new(logics.HcardLogic)
	*reply = true
	if err = hcardLogic.ShopDownUpHcard(ctx, args); err != nil {
		*reply = false
	}
	return
}

// ShopHcardListRpc 子店现时卡列表rpc内部调用
func (h *Hcard) ShopHcardListRpc(ctx context.Context, args *cards.ArgsShopHcardListRpc, reply *cards.ReplyShopHcardListRpc) (err error) {
	return new(logics.HcardLogic).ShopHcardListRpc(ctx, args, reply)
}

// 子店删除限时限次卡
func (h *Hcard) ShopDeleteHcard(ctx context.Context, args *cards.ArgsDeleteHcard, reply *bool) (err error) {
	*reply = true
	err = new(logics.HcardLogic).ShopDeleteHcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}
