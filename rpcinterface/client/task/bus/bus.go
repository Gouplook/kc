package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

// Bus Bus
type Bus struct {
	client.Baseclient
}

// Init Init
func (bus *Bus) Init() *Bus {
	bus.ServiceName = "rpc_task"
	bus.ServicePath = "Bus/Bus"
	return bus
}

// BusAdd 企业/商户添加
func (bus *Bus) BusAdd(ctx context.Context, busId *int, reply *bool) error {
	return bus.Call(ctx, "BusAdd", busId, reply)
}

// BusAudit 审核商户
func (bus *Bus) BusAudit(ctx context.Context, busId *int, reply *bool) error {
	return bus.Call(ctx, "BusAudit", busId, reply)
}


func (bus *Bus) BusArea(ctx context.Context, areaStr *string, reply *bool ) error  {
	return bus.Call(ctx, "BusArea", areaStr, reply)
}