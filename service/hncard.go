package service

import (
	"context"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type HNCard struct {
	ncl *logics.HNCardLogic
}

//初始化服务
func (n *HNCard) Init() *HNCard {
	n.ncl = new(logics.HNCardLogic)
	return n
}

//添加限时限次卡
func (n *HNCard) AddHNCard(ctx context.Context, args *cards.ArgsAddHNCard, replies *cards.RepliesAddHNCard) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var nCardID int
	if nCardID, err = n.ncl.AddHNCard(ctx, busId, args); err == nil {
		replies.HNCardID = nCardID
	}
	return
}

//编辑限时限次卡信息
func (n *HNCard) EditHNCard(ctx context.Context, args *cards.ArgsEditHNCard, replies *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	err = n.ncl.EditHNCard(ctx, busId, args)
	return
}

//获取限时限次卡的详情
func (n *HNCard) HNCardInfo(ctx context.Context, args *cards.ArgsHNCardInfo, reply *cards.ReplyHNCardInfo) (err error) {
	mHNCard := new(logics.HNCardLogic)
	*reply, err = mHNCard.HNCardInfo(ctx, args.HNCardID, args.ShopID)
	return
}

//获取商家的限时限次卡列表
func (n *HNCard) BusHNCardPage(ctx context.Context, args *cards.ArgsBusHNCardPage, reply *cards.ReplyHNCardPage) (err error) {
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mHNCard := new(logics.HNCardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	shopId, _ := args.GetShopId()
	*reply, err = mHNCard.GetBusPage(ctx, busID, shopId, start, limit, args.IsGround,args.FilterShopHasAdd)
	return
}

//设置限时限次卡适用门店
func (n *HNCard) SetHNCardShop(ctx context.Context, args *cards.ArgsSetHNCardShop, reply *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mHNCard := new(logics.HNCardLogic)
	err = mHNCard.SetHNCardShop(ctx, busId, args)
	return
}

//总店上下架限时限次卡
func (n *HNCard) DownUpHNCard(ctx context.Context, args *cards.ArgsDownUpHNCard, reply *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mHNCard := new(logics.HNCardLogic)
	err = mHNCard.DownUpHNCard(ctx, busId, args)
	return
}

//子店获取适用本店的限时限次卡列表
func (n *HNCard) ShopGetBusHNCardPage(ctx context.Context, args *cards.ArgsShopGetBusHNCardPage, reply *cards.ReplyHNCardPage) (err error) {
	var busID int
	var shopID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopID, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mHNCard := new(logics.HNCardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mHNCard.ShopGetBusHNCardPage(ctx, busID, shopID, start, limit)
	return
}

//子店添加限时限次卡到自己的店铺
func (n *HNCard) ShopAddHNCard(ctx context.Context, args *cards.ArgsShopAddHNCard, reply *cards.EmptyReplies) (err error) {
	var busID int
	var shopID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopID, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mHNCard := new(logics.HNCardLogic)
	err = mHNCard.ShopAddHNCard(ctx, busID, shopID, args)
	return
}

//获取子店的限时限次卡列表
func (n *HNCard) ShopHNCardPage(ctx context.Context, args *cards.ArgsShopHNCardPage, reply *cards.ReplyHNCardPage) (err error) {
	mHNCard := new(logics.HNCardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mHNCard.ShopHNCardPage(ctx, args.ShopID, start, limit, args.Status)
	return
}

//子店上下架限时限次卡
func (n *HNCard) ShopDownUpHNCard(ctx context.Context, args *cards.ArgsShopDownUpHNCard, reply *cards.EmptyReplies) (err error) {
	var shopID int
	if shopID, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mHNCard := new(logics.HNCardLogic)
	err = mHNCard.ShopDownUpHNCard(ctx, shopID, args)
	return
}

// 子店限时限次卡-rpc
func (n *HNCard) ShopHncardRpc(ctx context.Context, args *cards.ArgsShopHncardRpc, list *cards.ReplyShopHncardRpc) (err error) {
	return new(logics.HNCardLogic).ShopHncardRpc(ctx, args, list)
}

// 总店删除限时限次卡
func (n *HNCard) DeleteHNCard(ctx context.Context, args *cards.ArgsDelHNCard, reply *bool) (err error) {
	*reply = true
	err = new(logics.HNCardLogic).DeleteHNCard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

// 子店删除限时限次卡
func (n *HNCard) ShopDeleteHNCard(ctx context.Context, args *cards.ArgsDelHNCard, reply *bool) (err error) {
	*reply = true
	err = new(logics.HNCardLogic).ShopDeleteHNCard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}
