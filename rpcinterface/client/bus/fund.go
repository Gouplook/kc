package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

type Fund struct {
	client.Baseclient
}

func (f *Fund) Init() *Fund {
	f.ServiceName = "rpc_bus"
	f.ServicePath = "Fund"
	return f
}

func (f *Fund) PayCharge(ctx context.Context, fundRecord *bus.FundRecord, reply *bool) error {
	return f.Call(ctx, "PayCharge", fundRecord, reply)
}


//确认消费，商家存管资金变动
func (f *Fund) ConsumeCharge(ctx context.Context, args *bus.ArgsConsumeCharge, reply *bool) error {
	return f.Call(ctx, "ConsumeCharge", args, reply)
}

//处理出账中的存管消费金额
func (f *Fund) DeposNeedOut(ctx context.Context, clearId *int, reply *bool) error {
	return f.Call(ctx, "DeposNeedOut", clearId, reply)
}