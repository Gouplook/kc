//套餐服务
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/15 14:47
package service

import (
	"context"

	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Sm struct {
}

//添加套餐
func (s *Sm) AddSm(ctx context.Context, sm *cards.ArgsAddSm, smId *int) (err error) {
	mSm := new(logics.SmLogic)
	*smId, err = mSm.AddSm(ctx, sm)
	if err != nil {
		return
	}
	return nil
}

//编辑套餐信息
func (s *Sm) EditSm(ctx context.Context, sm *cards.ArgsEditSm, reply *bool) (err error) {
	mSm := new(logics.SmLogic)
	*reply = true
	err = mSm.EditSm(ctx, sm)
	if err != nil {
		*reply = false
		return
	}
	return nil
}

//获取套餐的详情
func (s *Sm) SmInfo(ctx context.Context, args *cards.ArgsSmInfo, reply *cards.ReplySmInfo) (err error) {
	mSm := new(logics.SmLogic)
	*reply, err = mSm.SmInfo(ctx, args.SmId, args.ShopId)
	return

}

//获取商家的套餐列表
func (s *Sm) BusSmPage(ctx context.Context, args *cards.ArgsBusSmPage, reply *cards.ReplySmPage) (err error) {
	mSm := new(logics.SmLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mSm.GetBusPage(ctx, args.BusId, args.ShopId, start, limit, args.IsGround,args.FilterShopHasAdd)
	return
}

//设置套餐适用门店
func (s *Sm) SetSmShop(ctx context.Context, args *cards.ArgsSetSmShop, reply *bool) (err error) {
	mSm := new(logics.SmLogic)
	*reply = true
	err = mSm.SetSmShop(ctx, args)
	return
}

//总店上下架套餐
/*func (s *Sm) DownUpSm(ctx context.Context, args *cards.ArgsDownUpSm, reply *bool) (err error) {
	mSm := new(logics.SmLogic)
	*reply = true
	err = mSm.DownUpSm(ctx, args)
	return
}*/

//子店获取适用本店的套餐列表
func (s *Sm) ShopGetBusSmPage(ctx context.Context, args *cards.ArgsShopGetBusSmPage, reply *cards.ReplyShopGetBusSmPage) (err error) {
	mSm := new(logics.SmLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mSm.ShopGetBusSmPage(ctx, args.BusId, args.ShopId, start, limit)
	return
}

//子店添加套餐到自己的店铺
func (s *Sm) ShopAddSm(ctx context.Context, args *cards.ArgsShopAddSm, reply *bool) (err error) {
	mSm := new(logics.SmLogic)
	*reply = true
	err = mSm.ShopAddSm(args)
	return
}

//获取子店的套餐列表
func (s *Sm) ShopSmPage(ctx context.Context, args *cards.ArgsShopSmPage, reply *cards.ReplyShopSmPage) (err error) {
	mSm := new(logics.SmLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mSm.ShopSmPage(ctx, args.ShopId, start, limit, args.Status)
	return
}

//子店上下架套餐
func (s *Sm) ShopDownUpSm(ctx context.Context, args *cards.ArgsShopDownUpSm, reply *bool) (err error) {
	mSm := new(logics.SmLogic)
	*reply = true
	err = mSm.ShopDownUpSm(ctx, args)
	return
}

//总店-删除套餐

func (s *Sm) DeleteSm(ctx context.Context, args *cards.ArgsDeleteSm, reply *bool) (err error) {
	if _, err := args.GetBusId(); err != nil {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	if len(args.SmIds) == 0 {
		return toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
	}

	if bools := new(logics.SmLogic).DeleteSmLogic(ctx, args, reply); !bools {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	return
}

//分店-删除套餐
func (s *Sm) DeleteShopSm(ctx context.Context, args *cards.ArgsDeleteShopSm, reply *bool) (err error) {
	if _, err := args.GetShopId(); err != nil {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	if len(args.SmIds) == 0 {
		return toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
	}
	if bools := new(logics.SmLogic).DeleteShopSmLogic(ctx, args, reply); !bools {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	return
}

// 子店添加套餐（一期优化，去掉总店推送，改为门店自动拉取)
func (s *Sm) ShopAddToSm(ctx context.Context, args *cards.ArgsShopAddSm, reply *bool) (err error) {
	mSm := new(logics.SmLogic)
	*reply = true
	err = mSm.ShopAddToSm(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}
