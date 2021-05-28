package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	v1 "git.900sui.cn/kc/rpcinterface/interface/open/v1"
)

type Depos struct {
	client.Baseclient
}

func (d *Depos) Init() *Depos {
	d.ServiceName = "rpc_bus"
	d.ServicePath = "Depos"
	return d
}

func (d *Depos) GetRatio(ctx context.Context, busId *int, ratio *float64) error {
	return d.Call(ctx, "GetRatio", busId, ratio)
}

// 批量获取建行存管账户信息
func (d *Depos) GetDeposCcbInfos(ctx context.Context, busIds *[]int, replyDeposCcbInfos *[]bus.ReplyDeposCcbInfo) error {
	return d.Call(ctx, "GetDeposCcbInfos", busIds, replyDeposCcbInfos)
}

// 批量获取杉德存管账户信息
func (d *Depos) GetDeposSandInfos(ctx context.Context, busIds *[]int, replyDeposSandInfo *[]bus.ReplyDeposSandInfo) error {
	return d.Call(ctx, "GetDeposSandInfos", busIds, replyDeposSandInfo)
}

//获取商家的存管账号和存管资金信息
func (d *Depos) GetBusDeposInfo(ctx context.Context, args *bus.ArgsGetBusDeposInfo, reply *bus.ReplyGetBusDeposInfo) error {
	return d.Call(ctx, "GetBusDeposInfo", args, reply)
}

//获取资金明细
func (d *Depos) GetBusAmountLogs(ctx context.Context, args *bus.ArgsBusPage, reply *bus.ReplyGetBusAmountLogs) error {
	return d.Call(ctx, "GetBusAmountLogs", args, reply)
}

//获取存管明细
func (d *Depos) GetBusDeposLogs(ctx context.Context, args *bus.ArgsBusPage, reply *bus.ReplyGetBusDeposLogs) error {
	return d.Call(ctx, "GetBusDeposLogs", args, reply)
}

//获取提现记录
func (d *Depos) GetBusCashLogs(ctx context.Context, args *bus.ArgsBusPage, reply *bus.ReplyGetBusCashLogsPage) error {
	return d.Call(ctx, "GetBusCashLogs", args, reply)
}

//申请提现
func (d *Depos) CashApply(ctx context.Context, args *bus.ArgsCashApply, reply *bool) error {
	return d.Call(ctx, "CashApply", args, reply)
}

//发送上海银行提现验证码
func (d *Depos) BoscCashCaptcha(ctx context.Context, bsToken *common.BsToken, reply *bool) error {
	return d.Call(ctx, "BoscCashCaptcha", bsToken, reply)
}

//上海银行提现结果处理业务 - 通过任务查询提现结果
func (d *Depos) BoscCashResultQuery(ctx context.Context, orderSn *string, reply *bool) error {
	return d.Call(ctx, "BoscCashResultQuery", orderSn, reply)
}

//获取商家的存管账号信息
func (d *Depos) GetBusDeposInfoByBusid(ctx context.Context, busId *int, reply *bus.ReplyGetBusDeposInfo) error {
	return d.Call(ctx, "GetBusDeposInfoByBusid", busId, reply)
}

//开放平台-申请提现
func (d *Depos) OpenPlatFormV1CashApply(ctx context.Context, args *v1.ArgsMerchantApplyWithdraw, reply *bool) (err error) {
	return d.Call(ctx, "OpenPlatFormV1CashApply", args, reply)
}

//开放平台-获取提现记录
func (d *Depos) OpenPlatFormV1GetBusCashLogs(ctx context.Context, args *v1.ArgsGetMerchantWithdrawLogs, reply *bus.ReplyGetBusCashLogsPage) (err error) {
	return d.Call(ctx, "OpenPlatFormV1GetBusCashLogs", args, reply)
}

//开放平台-获取商家的存管账号和存管资金信息
func (d *Depos) OpenPlatFormV1GetBusDeposInfo(ctx context.Context, args *v1.ArgsMerchantDeposInfo, reply *bus.ReplyGetBusDeposInfo) (err error) {
	return d.Call(ctx, "OpenPlatFormV1GetBusDeposInfo", args, reply)
}

//开放平台-获取存管明细
func (d *Depos)  OpenPlatFormV1GetBusDeposLogs(ctx context.Context, args *v1.ArgsMerchantDeposLogs, reply *bus.ReplyGetBusDeposLogs) (err error) {
	return d.Call(ctx, "OpenPlatFormV1GetBusDeposLogs", args, reply)
}
