package service

import (
	"context"
	"git.900sui.cn/kc/base/common/toolLib"
	_const "git.900sui.cn/kc/rpcCards/lang/const"

	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//ICard 折扣卡管理（身份卡）
type ICard struct {
	iCardLogic *logics.ICardLogic
}

//Init 实例化
func (s *ICard) Init() *ICard {
	s.iCardLogic = new(logics.ICardLogic)
	return s
}

//List 列表
func (s *ICard) List(ctx context.Context, args cards.InputParamsICardList, reply *cards.OutputICardList) (err error) {
	busID, _ := args.BsToken.GetBusId()

	shopId, _ := args.GetShopId()
	err = s.iCardLogic.List(ctx, busID, shopId, args, reply)
	return
}

//Save 保存
func (s *ICard) Save(ctx context.Context, args cards.InputParamsICardSave, reply *cards.OutputParamsICardSave) (err error) {
	busID, err := args.GetBusId()
	if err != nil {
		return err
	}
	err = s.iCardLogic.Save(ctx, busID, args, reply)
	return
}

//Info 详情
func (s *ICard) Info(ctx context.Context, args cards.InputParamsICardInfo, reply *cards.OutputParamsICardInfo) (err error) {
	err = s.iCardLogic.Info(ctx, args, reply)
	return
}

//Delete 删除
func (s *ICard) Delete(ctx context.Context, args cards.InputParamsDelete, reply *bool) (err error) {
	busId, err := args.GetBusId()
	if err != nil {
		return err
	}
	shopId, _ := args.GetShopId()
	*reply, err = s.iCardLogic.Delete(ctx, busId, shopId, args.IcardIds)
	return
}

//Push 推送门店
func (s *ICard) Push(ctx context.Context, args cards.InputParamsICardPush, reply *cards.OutputParamsICardPush) (err error) {
	err = s.iCardLogic.Push(ctx, args, reply)
	return
}

//SetOn 上架
func (s *ICard) SetOn(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	err = s.iCardLogic.SetOn(ctx, args, reply)
	return
}

//SetOff 下架
func (s *ICard) SetOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	err = s.iCardLogic.SetOff(ctx, args, reply)
	return
}

//SetOnOff 上下架
func (s *ICard) SetOnOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	err = s.iCardLogic.SetOnOff(ctx, args, reply)
	return
}

//OurShopList 已添加到门店列表
func (s *ICard) OurShopList(ctx context.Context, args cards.InputParamsICardList, reply *cards.OutputICardList) (err error) {
	err = s.iCardLogic.OurShopList(ctx, args, reply)
	return
}

//ShopList 门店列表
func (s *ICard) ShopList(ctx context.Context, args cards.InputParamsICardList, reply *cards.OutputICardList) (err error) {
	err = s.iCardLogic.ShopList(ctx, args, reply)
	return
}

//ShopSetOnOff 门店上下架
func (s *ICard) ShopSetOnOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	err = s.iCardLogic.ShopSetOnOff(ctx, args, reply)
	return
}

//ShopSetOn 门店上架
func (s *ICard) ShopSetOn(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	err = s.iCardLogic.ShopSetOn(ctx, args, reply)
	return
}

//ShopSetOff 门店下架
func (s *ICard) ShopSetOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	err = s.iCardLogic.ShopSetOff(ctx, args, reply)
	return
}

//AddToShop 添加至本店
func (s *ICard) AddToShop(ctx context.Context, args cards.InputParamsICardAddToShop, reply *cards.OutputParamsICardAddToShop) (err error) {
	shopID, _ := args.BsToken.GetShopId()
	busID, _ := args.BsToken.GetBusId()
	if shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	err = s.iCardLogic.AddToShop(ctx,busID,shopID, args, reply)
	return
}

//UserICardList 查看用户身份卡列表
func (s *ICard) UserICardList(ctx context.Context, args cards.InputParams, reply *cards.PageOutputReply) (err error) {
	*reply = s.iCardLogic.UserICardList(ctx, args)
	return
}

//CanUseICardList 查看用户权益身份卡列表
func (s *ICard) CanUseICardList(ctx context.Context, args cards.InputParamsICardCanUse, reply *cards.OutputParamsICardCanUse) (err error) {
	err = s.iCardLogic.CanUseICardList(ctx, args, reply)
	return
}

//CanUseICardList 查看用户权益身份卡列表
func (s *ICard) CanUseICardListForUser(ctx context.Context, args cards.InputParamsICardCanUseForUser, reply *cards.OutputParamsICardCanUse) (err error) {
	err = s.iCardLogic.CanUseICardListForUser(ctx, args, reply)
	return
}

//获取iCard企业基本信息-风控统计用
func (s *ICard) GetBusBaseInfoRpc(ctx context.Context, iCardId int, reply *cards.ReplyGetBusBaseInfoRpc) (err error) {
	*reply, err = s.iCardLogic.GetBusBaseInfoRpc(iCardId)
	return
}

//获取身份卡的折扣信息
func (s *ICard) GetICardDiscountById(ctx context.Context, iCardId *int, reply *cards.ReplyGetIcardDiscountById) (err error) {
	return s.iCardLogic.GetICardDiscountById(ctx, *iCardId, reply)
}

//获取身份卡备份表中的项目折扣
func (s *ICard) GetICardSingleDiscount(ctx context.Context, args *cards.ArgsGetICardSingleDiscount, reply *cards.ReplyGetICardSingleDiscount) (err error) {
	return s.iCardLogic.GetICardSingleDiscount(ctx, args, reply)
}
