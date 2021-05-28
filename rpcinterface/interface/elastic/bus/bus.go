package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)
// ES 搜索
type Bus interface {
	// SetAdminBus 设置商户信息到ES
	SetAdminBus(ctx context.Context,busId *int, reply *bool)error
	// SearchAdminBus ES商户后台搜索
	SearchAdminBus(ctx context.Context, args *bus.ArgsAdminBusAuditPage, reply *bus.ReplyAdminBusAuditPage) error
}
