package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

type Bus struct {
	client.Baseclient
}

func (bus *Bus)Init()*Bus{
	bus.ServiceName="rpc_elastic"
	bus.ServicePath="Bus/Bus"
	return bus
}
// SetBus 设置商户信息到ES
func (bus *Bus)SetAdminBus(ctx context.Context,busId *int, reply *bool)error  {
	return bus.Call(ctx,"SetAdminBus",busId,reply)
}

// SearchAdminBus ES商户后台搜索
func (bus *Bus)SearchAdminBus(ctx context.Context, args *bus.ArgsAdminBusAuditPage, reply *bus.ReplyAdminBusAuditPage) error  {
	return bus.Call(ctx,"SearchAdminBus",args,reply)
}
