package bus

import "context"

// Bus Bus
type Bus interface {
	// 企业/商户信息添加
	BusAdd(ctx context.Context, busId *int, reply *bool) error
	BusAudit(ctx context.Context, busId *int, reply *bool) error
	BusArea(ctx context.Context, areaStr *string, reply *bool ) error
}
