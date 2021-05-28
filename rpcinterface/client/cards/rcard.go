//@author yangzhiwu<578154898@qq.com>
//@date 2020/10/21 14:44
package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Rcard struct {
	client.Baseclient
}

func (r *Rcard) Init() *Rcard {
	r.ServiceName = "rpc_cards"
	r.ServicePath = "Rcard"
	return r
}

//添加充值卡
func (r *Rcard) AddRcard(ctx context.Context, args *cards.ArgsAddRcard, rcardId *int) error {
	return r.Call(ctx, "AddRcard", args, rcardId)
}

//添加充值卡-开放平台
func (r *Rcard) OpenPlatFormV1AddRcard(ctx context.Context, args *cards.ArgsAddRcard, rcardId *int) error {
	return r.Call(ctx, "OpenPlatFormV1AddRcard", args, rcardId)
}

//编辑充值卡
func (r *Rcard) EditRcard(ctx context.Context, args *cards.ArgsEditRcard, reply *bool) error {
	return r.Call(ctx, "EditRcard", args, reply)
}

//获取冲值卡详情
func (r *Rcard) RcardInfo(ctx context.Context, args *cards.ArgsRcardInfo, reply *cards.ReplyRcardInfo) error {
	return r.Call(ctx, "RcardInfo", args, reply)
}

//获取总店的充值卡列表
func (r *Rcard) BusRcardPage(ctx context.Context, args *cards.ArgsBusRcardPage, reply *cards.ReplyRcardPage) error {
	return r.Call(ctx, "BusRcardPage", args, reply)
}

//获取总店的充值卡列表-开放平台
func (r *Rcard) OpenPlatFormV1BusRcardPage(ctx context.Context, args *cards.ArgsBusRcardPage, reply *cards.ReplyRcardPage) error {
	return r.Call(ctx, "OpenPlatFormV1BusRcardPage", args, reply)
}

//设置适用门店
func (r *Rcard) SetRcardShop(ctx context.Context, args *cards.ArgsSetRcardShop, reply *bool) error {
	return r.Call(ctx, "SetRcardShop", args, reply)
}

//总店上下架充值卡
func (r *Rcard) DownUpRcard(ctx context.Context, args *cards.ArgsDownUpRcard, reply *bool) error {
	return r.Call(ctx, "DownUpRcard", args, reply)
}

// 总店删除充值卡
func (r *Rcard) DeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) error {
	return r.Call(ctx, "DeleteRcard", args, reply)
}

// 总店删除充值卡-开放平台
func (r *Rcard) OpenPlatFormV1DeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) error {
	return r.Call(ctx, "OpenPlatFormV1DeleteRcard", args, reply)
}

//子店获取适用本店的充值卡列表
func (r *Rcard) ShopGetBusRcardPage(ctx context.Context, args *cards.ArgsShopGetBusRcardPage, reply *cards.ReplyRcardPage) error {
	return r.Call(ctx, "ShopGetBusRcardPage", args, reply)
}

//子店添加充值卡到自己的店铺
func (r *Rcard) ShopAddRcard(ctx context.Context, args *cards.ArgsShopAddRcard, reply *bool) error {
	return r.Call(ctx, "ShopAddRcard", args, reply)
}

//子店添加充值卡到自己的店铺-开放平台
func (r *Rcard) OpenPlatFormV1ShopAddRcard(ctx context.Context, args *cards.ArgsShopAddRcard, reply *bool) error {
	return r.Call(ctx, "OpenPlatFormV1ShopAddRcard", args, reply)
}

//获取子店的充值卡列表
func (r *Rcard) ShopRcardPage(ctx context.Context, args *cards.ArgsShopRcardPage, reply *cards.ReplyShopRcardPage) error {
	return r.Call(ctx, "ShopRcardPage", args, reply)
}

//获取子店的充值卡列表-开放平台
func (r *Rcard) OpenPlatFormV1ShopRcardPage(ctx context.Context, args *cards.ArgsShopRcardPage, reply *cards.ReplyShopRcardPage) error {
	return r.Call(ctx, "OpenPlatFormV1ShopRcardPage", args, reply)
}

//子店上下架充值卡
func (r *Rcard) ShopDownUpRcard(ctx context.Context, args *cards.ArgsShopDownUpRcard, reply *bool) error {
	return r.Call(ctx, "ShopDownUpRcard", args, reply)
}

//子店上下架充值卡-开放平台
func (r *Rcard) OpenPlatFormV1ShopDownUpRcard(ctx context.Context, args *cards.ArgsShopDownUpRcard, reply *bool) error {
	return r.Call(ctx, "OpenPlatFormV1ShopDownUpRcard", args, reply)
}

// 子店删除充值卡
func (r *Rcard) ShopDeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) error {
	return r.Call(ctx, "ShopDeleteRcard", args, reply)
}

// 子店删除充值卡-开放平台
func (r *Rcard) OpenPlatFormV1ShopDeleteRcard(ctx context.Context, args *cards.ArgsDelRcard, reply *bool) error {
	return r.Call(ctx, "OpenPlatFormV1ShopDeleteRcard", args, reply)
}

//获取充值卡基础数据
func (r *Rcard) GetRcardBaseInfo(ctx context.Context, args *cards.ArgsGetRcardBaseInfo, reply *cards.ReplyGetRcardBaseInfo) error {
	return r.Call(ctx, "GetRcardBaseInfo", args, reply)
}

// 总店新增充值规则
func (r *Rcard) AddRechargeRules(ctx context.Context, args *cards.ArgsAddRechargeRules, id *int) error {
	return r.Call(ctx, "AddRechargeRules", args, id)
}

// 充值卡规则编辑
func (r *Rcard) EditRechargeRules(ctx context.Context, args *cards.ArgsEditRechargeRules, reply *bool) error {
	return r.Call(ctx, "EditRechargeRules", args, reply)
}

// 获取充值卡规则详情
func (r *Rcard) RechargeRulesInfo(ctx context.Context, args *cards.ArgsRechargeRulesInfo, reply *cards.ReplyRechargeRulesInfo) error {
	return r.Call(ctx, "RechargeRulesInfo", args, reply)
}

// 删除充值卡规则
func (r *Rcard) DeleRechargeRules(ctx context.Context, args *cards.ArgsDeleRechargeRules, reply *bool) error {
	return r.Call(ctx, "DeleRechargeRules", args, reply)
}

// 获取充值规则列表
func (r *Rcard) BusRechargeRulesList(ctx context.Context, args *cards.ArgsRechargeRulesList, reply *cards.ReplyRechargerRulesList) error {
	return r.Call(ctx, "BusRechargeRulesList", args, reply)
}
