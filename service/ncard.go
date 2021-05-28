package service

import (
	"context"

	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type NCard struct {
	ncl *logics.NCardLogic
}

//初始化服务
func (n *NCard) Init() *NCard {
	n.ncl = new(logics.NCardLogic)
	return n
}

//添加限次卡
func (n *NCard) AddNCard(ctx context.Context, args *cards.ArgsAddNCard, replies *cards.RepliesAddNCard) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
	}
	var nCardID int
	if nCardID, err = n.ncl.AddNCard(ctx, busId, args); err == nil {
		replies.NCardID = nCardID
	}
	return
}

//编辑限次卡信息
func (n *NCard) EditNCard(ctx context.Context, args *cards.ArgsEditNCard, replies *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
	}
	err = n.ncl.EditNCard(ctx, busId, args)
	return
}

//获取限次卡的详情
func (n *NCard) NCardInfo(ctx context.Context, args *cards.ArgsNCardInfo, reply *cards.ReplyNCardInfo) (err error) {
	mNCard := new(logics.NCardLogic)
	*reply, err = mNCard.NCardInfo(ctx, args.NCardID, args.ShopID)
	return

}

//获取商家的限次卡列表
func (n *NCard) BusNCardPage(ctx context.Context, args *cards.ArgsBusNCardPage, reply *cards.ReplyNCardPage) (err error) {
	var busID int
	if busID, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mNCard := new(logics.NCardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	shopId, _ := args.GetShopId()
	*reply, err = mNCard.GetBusPage(ctx, busID, shopId, start, limit, args.IsGround,args.FilterShopHasAdd)
	return
}

//设置限次卡适用门店
func (n *NCard) SetNCardShop(ctx context.Context, args *cards.ArgsSetNCardShop, reply *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
	}
	mNCard := new(logics.NCardLogic)
	err = mNCard.SetNCardShop(ctx, busId, args)
	return
}

//总店上下架限次卡
/*func (n *NCard) DownUpNCard(ctx context.Context, args *cards.ArgsDownUpNCard, reply *cards.EmptyReplies) (err error) {
	//交验是否有总店权限
	if acc, accErr := args.GetBusAcc(); accErr != nil || !acc {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var busId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
	}
	mNCard := new(logics.NCardLogic)
	err = mNCard.DownUpNCard(ctx, busId, args)
	return
}*/

//子店获取适用本店的限次卡列表
func (n *NCard) ShopGetBusNCardPage(ctx context.Context, args *cards.ArgsShopGetBusNCardPage, reply *cards.ReplyNCardPage) (err error) {
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
	mNCard := new(logics.NCardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mNCard.ShopGetBusNCardPage(ctx, busID, shopID, start, limit)
	return
}

//子店添加限次卡到自己的店铺
func (n *NCard) ShopAddNCard(ctx context.Context, args *cards.ArgsShopAddNCard, reply *cards.EmptyReplies) (err error) {
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
	mNCard := new(logics.NCardLogic)
	err = mNCard.ShopAddNCard(ctx, busID, shopID, args)
	return
}

//获取子店的限次卡列表
func (n *NCard) ShopNCardPage(ctx context.Context, args *cards.ArgsShopNCardPage, reply *cards.ReplyNCardPage) (err error) {
	mNCard := new(logics.NCardLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mNCard.ShopNCardPage(ctx, args.ShopID, start, limit, args.Status)
	return
}

//子店上下架限次卡
func (n *NCard) ShopDownUpNCard(ctx context.Context, args *cards.ArgsShopDownUpNCard, reply *cards.EmptyReplies) (err error) {
	mNCard := new(logics.NCardLogic)
	var shopId int
	if shopId, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	err = mNCard.ShopDownUpNCard(ctx, shopId, args)
	return
}

//ShopNcardRpc
func (n *NCard) ShopNcardRpc(ctx context.Context, args *cards.ArgsShopNcardRpc, reply *cards.ReplyShopNcardRpc) (err error) {
	return new(logics.NCardLogic).ShopNcardRpc(ctx, args, reply)
}

//总店-软删除
func (n *NCard) DeleteNCard(ctx context.Context, params *cards.ArgsDeleteNCard, reply *bool) error {
	if busId, _ := params.GetBusId(); busId == 0 || len(params.NcardIds) == 0 {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	return new(logics.NCardLogic).DeleteNCardLogic(ctx, params, reply)
}

//分店-软删除
func (n *NCard) DeleteShopNCard(ctx context.Context, params *cards.ArgsDeleteShopNCard, reply *bool) error {
	if shopId, _ := params.GetShopId(); shopId == 0 || len(params.NcardIds) == 0 {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	return new(logics.NCardLogic).DeleteShopNCardLogic(ctx, params, reply)
}
