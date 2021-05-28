package service

import (
	"context"
	"strconv"

	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Card struct {
	ncl *logics.CardLogic
}

//初始化服务
func (n *Card) Init() *Card {
	n.ncl = new(logics.CardLogic)
	return n
}

//添加综合卡
func (n *Card) AddCard(ctx context.Context, args *cards.ArgsAddCard, replies *cards.RepliesAddCard) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
	}
	var cardID int
	if cardID, err = n.ncl.AddCard(ctx, busID, args); err == nil {
		replies.CardID = cardID
	}
	return
}

//编辑综合卡信息
func (n *Card) EditCard(ctx context.Context, args *cards.ArgsEditCard, replies *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
	}
	err = n.ncl.EditCard(ctx, busID, args)
	return
}

//获取综合卡的详情
func (n *Card) CardInfo(ctx context.Context, args *cards.ArgsCardInfo, reply *cards.ReplyCardInfo) (err error) {
	mCard := new(logics.CardLogic)
	*reply, err = mCard.CardInfo(ctx, args.CardID, args.ShopID)
	return

}

//获取商家的综合卡列表
func (n *Card) BusCardPage(ctx context.Context, args *cards.ArgsBusCardPage, reply *cards.ReplyCardPage) (err error) {
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mCard := new(logics.CardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	shopId, _ := args.GetShopId()
	*reply, err = mCard.GetBusPage(ctx, busID, shopId, start, limit, args.IsGround,args.FilterShopHasAdd)
	return
}

//设置综合卡适用门店
func (n *Card) SetCardShop(ctx context.Context, args *cards.ArgsSetCardShop, reply *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	mCard := new(logics.CardLogic)
	err = mCard.SetCardShop(ctx, busID, args)
	return
}

//总店上下架综合卡
/*func (n *Card) DownUpCard(ctx context.Context, args *cards.ArgsDownUpCard, reply *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mCard := new(logics.CardLogic)
	err = mCard.DownUpCard(ctx, busID, args)
	return
}*/

//子店获取适用本店的综合卡列表
func (n *Card) ShopGetBusCardPage(ctx context.Context, args *cards.ArgsShopGetBusCardPage, reply *cards.ReplyCardPage) (err error) {
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
	mCard := new(logics.CardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mCard.ShopGetBusCardPage(ctx, busID, shopID, start, limit)
	return
}

//子店添加综合卡到自己的店铺
func (n *Card) ShopAddCard(ctx context.Context, args *cards.ArgsShopAddCard, reply *cards.EmptyReplies) (err error) {
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
	mCard := new(logics.CardLogic)
	err = mCard.ShopAddCard(ctx, busID, shopID, args)
	return
}

//获取子店的综合卡列表
func (n *Card) ShopCardPage(ctx context.Context, args *cards.ArgsShopCardPage, reply *cards.ReplyCardPage) (err error) {
	mCard := new(logics.CardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	status, _ := strconv.Atoi(args.Status)
	*reply, err = mCard.ShopCardPage(ctx, args.ShopID, start, limit, status)
	return
}

//子店上下架综合卡
func (n *Card) ShopDownUpCard(ctx context.Context, args *cards.ArgsShopDownUpCard, reply *cards.EmptyReplies) (err error) {
	mCard := new(logics.CardLogic)
	var shopId int
	if shopId, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	err = mCard.ShopDownUpCard(ctx, shopId, args)
	return
}

//获取门限可售综合卡详情
func (n *Card) CardsInfo(ctx context.Context, args *cards.ArgsCardsInfo, reply *map[int]*cards.ReplyCardsInfo) (err error) {
	var shopId int
	shopId = 1
	//if shopId, err = args.GetShopId(); err != nil {
	//	err = toolLib.CreateKcErr(_const.POWER_ERR)
	//	return
	//}
	*reply, err = new(logics.CardLogic).GetCardsInfo(ctx, shopId, args.CardIds)
	return
}

// 门店综合卡数据rpc内部调用 rpc内部调用
func (n *Card) ShopCardListRpc(ctx context.Context, params *cards.ArgsShopCardListRpc, replies *cards.ReplyShopCardListRpc) (err error) {
	return new(logics.CardLogic).ShopCardListRpc(ctx, params, replies)
}

// 获取所有卡的发布数量 rpc内部调用
func (n *Card) GetAllCardsNum(ctx context.Context, args *cards.ArgsAllCardsNum, reply *cards.ReplyAllCardsNum) (err error) {
	return new(logics.CardLogic).GetAllCardsNum(ctx, args, reply)
}

//总店-删除综合卡
func (n *Card) DeleteCard(ctx context.Context, args *cards.ArgsDeleteCard, reply *bool) (err error) {
	if _, err := args.GetBusId(); err != nil {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	if len(args.CardIds) == 0 {
		return toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
	}
	return new(logics.CardLogic).DeleteCard(ctx, args, reply)
}

//分店-删除综合卡
func (n *Card) DeleteShopCard(ctx context.Context, args *cards.ArgsDeleteShopCard, reply *bool) (err error) {
	if _, err := args.GetShopId(); err != nil {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	if len(args.CardIds) == 0 {
		return toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
	}
	return new(logics.CardLogic).DeleteShopCard(ctx, args, reply)
}
