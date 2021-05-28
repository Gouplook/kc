package market

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/market"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type Equity struct {
	client.Baseclient
}

func (m *Equity) Init() *Equity {
	m.ServiceName = "rpc_market"
	m.ServicePath = "Equity"
	return m
}

//添加权益
func (m *Equity) AddEquity(ctx context.Context, args *market.ArgsAddEquity, reply *int) error {
	return m.Call(ctx, "AddEquity", args, reply)
}

//查询权益列表
func (m *Equity) GetEquityList(ctx context.Context, args *market.ArgsGetEquityList, reply *market.ReplyGetEquityList) error {
	return m.Call(ctx, "GetEquityList", args, reply)
}

//用户查询权益列表
func (m *Equity) GetEquityListByUid(ctx context.Context, args *market.ArgsGetEquityListByUid, reply *market.ReplyGetEquityList) error {
	return m.Call(ctx, "GetEquityListByUid", args, reply)
}

//查询一条详情
func (m *Equity) GetOneEquity(ctx context.Context, args *market.ArgsGetOneEquity, reply *market.ReplyGetOneEquity) error {
	return m.Call(ctx, "GetOneEquity", args, reply)
}

//内部使用-查询一条数据
func (m *Equity) GetOneEquityOfInternal(ctx context.Context, args *market.ArgsGetOneEquity, reply *market.ReplyGetOneEquity) error {
	return m.Call(ctx, "GetOneEquity", args, reply)
}

//确认消费权益卡
func (m *Equity) ConsumeEquity(ctx context.Context, args *market.ArgsConsumeEquity, reply *bool) error {
	return m.Call(ctx, "ConsumeEquity", args, reply)
}

//确认消费权益卡rpc
func (m *Equity) ConsumeEquityRpc(ctx context.Context, args *market.ArgsConsumeEquityRpc, reply *bool) error {
	return m.Call(ctx, "ConsumeEquityRpc", args, reply)
}

//批量确认消费权益卡rpc
func (m *Equity) BatchConsumeEquity(ctx context.Context, args *market.ArgsBatchConsumeEquity, reply *bool) error {
	return m.Call(ctx, "BatchConsumeEquity", args, reply)
}

//根据权益ids批量获取权益列表基础数据
func (m *Equity) GetEquityListsByIds(ctx context.Context, args *market.ArgsGetEquityListsByIds, reply *market.ReplyGetEquityListsByIds) error {
	return m.Call(ctx, "GetEquityListsByIds", args, reply)
}

//获取用户权益卡二维码信息
func (m *Equity) GetUserEquityQrcode(ctx context.Context, args *market.ArgsGetUserEquityQrcode, reply *string) error {
	return m.Call(ctx, "GetUserEquityQrcode", args, reply)
}

//根据消费码获取权益卡信息
func (m *Equity) GetQrcodeByConsumeCode(ctx context.Context, args *order.ArgsCardPackageQrcode, reply *market.ReplyGetQrcodeByConsumeCode) error {
	return m.Call(ctx, "GetQrcodeByConsumeCode", args, reply)
}

//用户可使用权益卡统计
func (m *Equity) GetUserEquityCountRpc(ctx context.Context, args *market.ArgsGetUserEquityCountRpc, reply *market.ReplyGetUserEquityCountRpc) error {
	return m.Call(ctx, "GetUserEquityCountRpc", args, reply)
}

//获取权益包下项目列表
func (m *Equity) GetEquityItemList(ctx context.Context, args *market.ArgsGetEquityItemList, replies *[]market.RepliesEquityDetailList) error {
	return m.Call(ctx, "GetEquityItemList", args, replies)
}

//bus member 查询 详情 使用rpc--次数
func (m *Equity) GetMemberDetailAsEquity(ctx context.Context, args *market.ArgsGetMemberDetailAsEquity, replies *market.ReplyGetMemberDetailAsEquity) error {
	return m.Call(ctx, "GetMemberDetailAsEquity", args, replies)
}

//bus member 查询 详情 使用rpc--个数
func (m *Equity) GetMemberDetailAsEquity2(ctx context.Context, args *market.ArgsGetMemberDetailAsEquity, replies *market.ReplyGetMemberDetailAsEquity) error {
	return m.Call(ctx, "GetMemberDetailAsEquity2", args, replies)
}
