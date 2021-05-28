//充值卡服务
//@author yangzhiwu<578154898@qq.com>
//@date 2020/10/21 14:47
package service

import (
	"context"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Rcard struct {
}

//添加充值卡
func (r *Rcard) AddRcard(ctx context.Context, args *cards.ArgsAddRcard, rcardId *int) (err error) {
	_, err = args.GetBusAcc()
	if err != nil {
		return
	}
	busId, err := args.GetBusId()
	if err != nil {
		return err
	}
	*rcardId, err = new(logics.RcardLogic).AddRcard(ctx, busId, args)
	return
}

//添加充值卡-开放平台
func (r *Rcard) OpenPlatFormV1AddRcard(ctx context.Context, args *cards.ArgsAddRcard, rcardId *int) (err error) {
	*rcardId, err = new(logics.RcardLogic).AddRcard(ctx, args.BusID, args)
	return
}

//编辑充值卡
func (r *Rcard) EditRcard(ctx context.Context, args *cards.ArgsEditRcard, reply *bool) (err error) {
	_, err = args.GetBusAcc()
	if err != nil {
		return
	}
	busId, err := args.GetBusId()
	if err != nil {
		return err
	}
	*reply = true
	err = new(logics.RcardLogic).EditRcard(ctx, busId, args)
	if err != nil {
		*reply = false
	}
	return
}

//获取冲值卡详情
func (r *Rcard) RcardInfo(ctx context.Context, args *cards.ArgsRcardInfo, reply *cards.ReplyRcardInfo) (err error) {
	*reply, err = new(logics.RcardLogic).RcardInfo(ctx, args)
	return
}

//获取总店的充值卡列表
func (r *Rcard) BusRcardPage(ctx context.Context, args *cards.ArgsBusRcardPage, reply *cards.ReplyRcardPage) (err error) {
	*reply, err = new(logics.RcardLogic).BusRcardPage(ctx, args)
	return
}

//获取总店的充值卡列表-开放平台
func (r *Rcard) OpenPlatFormV1BusRcardPage(ctx context.Context, args *cards.ArgsBusRcardPage, reply *cards.ReplyRcardPage) (err error) {
	*reply, err = new(logics.RcardLogic).BusRcardPage(ctx, args)
	return
}

//设置适用门店
func (r *Rcard) SetRcardShop(ctx context.Context, args *cards.ArgsSetRcardShop, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).SetRcardShop(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

//总店上下架充值卡
func (r *Rcard) DownUpRcard(ctx context.Context, args *cards.ArgsDownUpRcard, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).DownUpRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

// 总店删除充值卡
func (r *Rcard) DeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).DeleteRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

// 总店删除充值卡-开放平台
func (r *Rcard) OpenPlatFormV1DeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).DeleteRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

//子店获取适用本店的充值卡列表
func (r *Rcard) ShopGetBusRcardPage(ctx context.Context, args *cards.ArgsShopGetBusRcardPage, reply *cards.ReplyRcardPage) (err error) {
	*reply, err = new(logics.RcardLogic).ShopGetBusRcardPage(ctx, args)
	return
}

//子店添加充值卡到自己的店铺
func (r *Rcard) ShopAddRcard(ctx context.Context, args *cards.ArgsShopAddRcard, reply *bool) (err error) {
	var busId, shopId int
	if busId, err = args.GetBusId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	*reply = true
	err = new(logics.RcardLogic).ShopAddRcard(ctx, shopId, busId, args.RcardIds)
	if err != nil {
		*reply = false
	}
	return
}

//子店添加充值卡到自己的店铺-开放平台
func (r *Rcard) OpenPlatFormV1ShopAddRcard(ctx context.Context, args *cards.ArgsShopAddRcard, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).ShopAddRcard(ctx, args.ShopId, args.BusId, args.RcardIds)
	if err != nil {
		*reply = false
	}
	return
}

//获取子店的充值卡列表
func (r *Rcard) ShopRcardPage(ctx context.Context, args *cards.ArgsShopRcardPage, reply *cards.ReplyShopRcardPage) (err error) {
	*reply, err = new(logics.RcardLogic).ShopRcardPage(ctx, args)
	return
}

//获取子店的充值卡列表-开放平台
func (r *Rcard) OpenPlatFormV1ShopRcardPage(ctx context.Context, args *cards.ArgsShopRcardPage, reply *cards.ReplyShopRcardPage) (err error) {
	*reply, err = new(logics.RcardLogic).ShopRcardPage(ctx, args)
	return
}

//子店上下架充值卡
func (r *Rcard) ShopDownUpRcard(ctx context.Context, args *cards.ArgsShopDownUpRcard, reply *bool) (err error) {
	var shopId int
	if shopId, err = args.GetShopId(); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId <= 0 {
		err = toolLib.CreateKcErr(_const.SHOPID_NTL)
		return
	}
	args.ShopId = shopId
	*reply = true
	err = new(logics.RcardLogic).ShopDownUpRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

//子店上下架充值卡-开放平台
func (r *Rcard) OpenPlatFormV1ShopDownUpRcard(ctx context.Context, args *cards.ArgsShopDownUpRcard, reply *bool) (err error) {
	if args.ShopId <= 0 {
		err = toolLib.CreateKcErr(_const.SHOPID_NTL)
		return
	}
	*reply = true
	err = new(logics.RcardLogic).ShopDownUpRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

//子店删除充值卡
func (r *Rcard) ShopDeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) (err error) {
	shopId, err := args.GetShopId()
	if err != nil {
		return
	}
	args.ShopId = shopId
	*reply = true
	err = new(logics.RcardLogic).ShopDeleteRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

// 子店删除充值卡-开放平台
func (r *Rcard) OpenPlatFormV1ShopDeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).ShopDeleteRcard(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

// 总店新增充值规则
func (r *Rcard) AddRechargeRules(ctx context.Context, args *cards.ArgsAddRechargeRules, id *int) (err error) {
	*id, err = new(logics.RcardLogic).AddRechargeRules(ctx, args)
	if err != nil {
		return
	}
	return
}

// 充值卡规则编辑
func (r *Rcard) EditRechargeRules(ctx context.Context, args *cards.ArgsEditRechargeRules, reply *bool) (err error) {
	*reply = true
	if err = new(logics.RcardLogic).EditRechargeRules(ctx, args); err != nil {
		*reply = false
	}
	return
}

// 获取充值卡规则详情
func (r *Rcard) RechargeRulesInfo(ctx context.Context, args *cards.ArgsRechargeRulesInfo, reply *cards.ReplyRechargeRulesInfo) (err error) {
	err = new(logics.RcardLogic).RechargeRulesInfo(ctx, args, reply)
	return
}

// 删除充值卡规则
func (r *Rcard) DeleRechargeRules(ctx context.Context, args *cards.ArgsDeleRechargeRules, reply *bool) (err error) {
	*reply = true
	err = new(logics.RcardLogic).DeleRechargeRules(ctx, args)
	if err != nil {
		*reply = false
	}
	return
}

// 获取充值规则列表
func (r *Rcard) BusRechargeRulesList(ctx context.Context, args *cards.ArgsRechargeRulesList, reply *cards.ReplyRechargerRulesList) (err error) {
	err = new(logics.RcardLogic).BusRechargeRulesList(ctx, args, reply)
	return
}

//获取充值卡基础数据
func (r *Rcard) GetRcardBaseInfo(ctx context.Context, args *cards.ArgsGetRcardBaseInfo, reply *cards.ReplyGetRcardBaseInfo) (err error) {
	return new(logics.RcardLogic).GetRcardBaseInfo(ctx, args, reply)
}
