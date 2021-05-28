package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)
//定义rpc调用
// @author liyang<654516092@qq.com>
// @date  2020/3/25 16:37
type BusAuth struct {
	client.Baseclient
}
//初始化
func (b *BusAuth)Init() *BusAuth {
	b.ServiceName = "rpc_bus"
	b.ServicePath = "BusAuth"
	return b
}

//统一鉴权-九百岁SAAS
func (b *BusAuth) BusAuth(ctx context.Context,args *bus.ArgsBusAuth,reply *bus.ReplyBusAuth) error{
	return b.Call(ctx, "BusAuth", args, reply)
}

//验证用户是否绑定商户主体
func (b *BusAuth) UserBindBusAuth(ctx context.Context,uid *int,reply *bool) error{
	return b.Call(ctx, "UserBindBusAuth", uid, reply)
}

//验证用户是否绑定分店
func (b *BusAuth) UserBindShopAuth(ctx context.Context,uid *int,reply *bool) error{
	return b.Call(ctx, "UserBindShopAuth", uid, reply)
}


