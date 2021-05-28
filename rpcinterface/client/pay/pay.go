package pay

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/pay"
)

type Pay struct {
	client.Baseclient
}

func (p *Pay) Init() *Pay {
	p.ServiceName = "rpc_pay"
	p.ServicePath = "Pay"
	return p
}

// 获取支付二维码
func (p *Pay) PayQr(ctx context.Context, args *pay.PayInfo, reply *string) error {
	return p.Call(ctx, "PayQr", args, reply)
}

// 获取支付url连接
func (p *Pay) PayH5(ctx context.Context, args *pay.PayInfo, reply *string) error {
	return p.Call(ctx, "PayH5", args, reply)
}

// 获取支付url连接
func (p *Pay) PayH5New(ctx context.Context, args *pay.PayInfo, reply *pay.PayH5) error {
	return p.Call(ctx, "PayH5New", args, reply)
}

// 获取小程序支付数据
func (p *Pay) PayWxapp(ctx context.Context, args *pay.PayInfo, reply *string) error {
	return p.Call(ctx, "PayWxapp", args, reply)
}

// 微信公众号支付数据
func (p *Pay)PayWxOfficial(ctx context.Context, args *pay.PayInfo, reply *string) error{
	return p.Call(ctx, "PayWxOfficial", args, reply)
}

// 获取app支付串
func (p *Pay) PayApp(ctx context.Context, args *pay.PayInfo, reply *string) error {
	return p.Call(ctx, "PayApp", args, reply)
}

// 获取app支付串
func (p *Pay) PayAppSign(ctx context.Context, args *pay.PayInfo, reply *string) error {
	return p.Call(ctx, "PayAppSign", args, reply)
}

// 异步回调处理
func (p *Pay) Notify(ctx context.Context, args *pay.ArgsNotify, reply *bool) error {
	return p.Call(ctx, "Notify", args, reply)
}

// 获取订单支付信息
func (p *Pay) PayInfo(ctx context.Context, orderSn *string, reply *pay.PayNotify) error {
	return p.Call(ctx, "PayInfo", orderSn, reply)
}

// 支付成功清分资金记账处理 支付成功后消息任务调度
func (p *Pay) PayAgent(ctx context.Context, orderSn *string, reply *bool) error {
	return p.Call(ctx, "PayAgent", orderSn, reply)
}

/**
 * 资金清分【商家】
 * @return void
 */
func (p *Pay) AngelChannel(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "AngelChannel", timeUnix, reply)
}

/**
 * 资金清分【平台】
 * @return void
 */
func (p *Pay) AngelChannelPlat(ctx context.Context,timeUnix *int,reply *bool) error{
	return p.Call(ctx, "AngelChannelPlat", timeUnix, reply)
}

/**
 * 资金清分【保险-正常】
 * @return void
 */
func (p *Pay) AngelChannelInsure(ctx context.Context,timeUnix *int,reply *bool) error{
	return p.Call(ctx, "AngelChannelInsure", timeUnix, reply)
}

/**
 * 资金清分【保险-续保】
 * @return void
 */
func (p *Pay) AngelChannelRenewInsure(ctx context.Context,timeUnix *int,reply *bool) error{
	return p.Call(ctx, "AngelChannelRenewInsure", timeUnix, reply)
}

// 商家建行分账清分
func (p *Pay) CcbAgent(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "CcbAgent", timeUnix, reply)
}

// 商家建行分账结果处理
func (p *Pay) CcbAgentFund(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "CcbAgentFund", timeUnix, reply)
}

// 商家杉德分账清分
func (p *Pay) SandAgent(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgent", timeUnix, reply)
}

// 商家杉德分账异步结果处理
func (p *Pay) SandAgentFund(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentFund", timeUnix, reply)
}

// 平台杉德分账清分
func (p *Pay) SandAgentPlat(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentPlat", timeUnix, reply)
}

// 平台杉德分账异步结果处理
func (p *Pay) SandAgentFundPlat(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentFundPlat", timeUnix, reply)
}

// 保险公司杉德分账清分
func (p *Pay) SandAgentInsure(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentInsure", timeUnix, reply)
}

// 保险公司杉德分账异步结果处理
func (p *Pay) SandAgentFundInsure(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentFundInsure", timeUnix, reply)
}

// 续保公司杉德分账清分
func (p *Pay) SandAgentRenewInsure(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentRenewInsure", timeUnix, reply)
}

// 续保公司杉德分账异步结果处理
func (p *Pay) SandAgentFundRenewInsure(ctx context.Context, timeUnix *int, reply *bool) error {
	return p.Call(ctx, "SandAgentFundRenewInsure", timeUnix, reply)
}

//根据订单号，获取代付状态信息
func (p *Pay) GetAgentInfoByOrderSn(ctx context.Context, orderSn *string, reply *pay.ReplyGetAgentInfoByOrderSn) error {
	return p.Call(ctx, "GetAgentInfoByOrderSn", orderSn, reply)
}

//根据clearId，获取代付状态信息
func (p *Pay) GetAgentInfoByClearid(ctx context.Context, clearId *int, reply *pay.ReplyGetAgentInfoByOrderSn) error {
	return p.Call(ctx, "GetAgentInfoByClearid", clearId, reply)
}

//异步回调响应
func (p *Pay) NotifyResponse(ctx context.Context, args *pay.ArgsNotify, reply *pay.ReplyNotifyResponse) error {
	return p.Call(ctx, "NotifyResponse", args, reply)
}
