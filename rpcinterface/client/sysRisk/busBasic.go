package sysRisk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/sysRisk"
)

type BusBasic struct {
	client.Baseclient
}

func (b *BusBasic) Init() *BusBasic {
	b.ServiceName = "rpc_sysrisk"
	b.ServicePath = "Bus"
	return b
}

//修改商家的区域数据
func (b *BusBasic) ChangeBusArea(ctx context.Context, args *sysRisk.ArgsChangeBusArea, reply *bool) error {
	return b.Call(ctx, "ChangeBusArea", args, reply)
}
