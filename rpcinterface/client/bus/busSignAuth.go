package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

/**
 * @className busSignAuth
 * @author liyang<654516092@qq.com>
 * @date 2020/11/9 15:29
 */

type BusSignAuth struct {
	client.Baseclient
}

//初始化
func (b *BusSignAuth) Init() *BusSignAuth {
	b.ServiceName = "rpc_bus"
	b.ServicePath = "BusSignAuth"
	return b
}

//绑定账户
func (b *BusSignAuth) SupplyAccount(ctx context.Context, args *bus.ArgsSupplyAccount, reply *bus.ReplyAccount) error {
	return b.Call(ctx, "SupplyAccount", args, reply)
}

//验证账户
func (b *BusSignAuth) AuthAccount(ctx context.Context, args *bus.ArgsAuthAccount, reply *bus.ReplyAccount) error {
	return b.Call(ctx, "AuthAccount", args, reply)
}

//sass绑定账户
func (b *BusSignAuth) SassSupplyAccount(ctx context.Context, args *bus.ArgsSassSupplyAccount, reply *bool) error {
	return b.Call(ctx, "SassSupplyAccount", args, reply)
}

//sass查询绑定账户
func (b *BusSignAuth) QuerySassAccount(ctx context.Context, args *bus.ArgsQuerySassAccount, reply *bus.ReplyQuerySassAccount) error {
	return b.Call(ctx, "QuerySassAccount", args, reply)
}
