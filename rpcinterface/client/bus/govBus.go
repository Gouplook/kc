package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

//监管平台信息对接成功的商户

type GovBus struct {
	client.Baseclient
}

func (g *GovBus) Init() *GovBus {
	g.ServiceName = "rpc_bus"
	g.ServicePath = "GovBus"
	return g
}

//同步监管平台对接成功的商户数据
func (g *GovBus) AnsyBus(ctx context.Context, args *bus.ArgsAnsyBus, reply *bool) error {
	return g.Call(ctx, "AnsyBus", args, reply)
}

//根据公司名称获取商家是否已对接
func (g *GovBus) GetByCompanyname(ctx context.Context, args *bus.ArgsGetByCompanyname, reply *bus.ReplyGetByCompanyname) error {
	return g.Call(ctx, "GetByCompanyname", args, reply)
}

//根据公司名称获取商家是否已对接-开放平台
func (g *GovBus) OpenPlatFormV1GetByCompanyname(ctx context.Context, args *bus.ArgsGetByCompanyname, reply *bus.ReplyGetByCompanyname) error {
	return g.Call(ctx, "OpenPlatFormV1GetByCompanyname", args, reply)
}

//根据riskbusIds 获取商家的评论
func (g *GovBus) GetCommentsByRiskbusids(ctx context.Context, riskBusIds *[]int, reply *map[int]bus.ReplyGetCommentsByRiskbusids) error {
	return g.Call(ctx, "GetCommentsByRiskbusids", riskBusIds, reply)
}

//监管平台添加商家发卡规则 - 暂停发卡/允许发卡
func (g *GovBus) StopOrStartSell(ctx context.Context, args *bus.ArgsStopOrStartSell, reply *bool) error {
	return g.Call(ctx, "StopOrStartSell", args, reply)
}

//监管平台添加商家发卡规则 - 设置发卡上限额度
func (g *GovBus) ChangeSellUplimiter(ctx context.Context, args *bus.ArgsChangeSellUplimiter, reply *bool) error {
	return g.Call(ctx, "ChangeSellUplimiter", args, reply)
}

//获取商家的发卡状态 发卡额度情况
func (g *GovBus) BusGovRule(ctx context.Context, busId *int, reply *bus.ReplyBusGovRule) error {
	return g.Call(ctx, "BusGovRule", busId, reply)
}
