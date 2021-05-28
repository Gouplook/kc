package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

/**
 * @className busshoprole
 * @author liyang<654516092@qq.com>
 * @date 2020/9/2 10:49
 */

type BusShopRole struct {
	client.Baseclient
}
//初始化
func (b *BusShopRole)Init() *BusShopRole {
	b.ServiceName = "rpc_bus"
	b.ServicePath = "BusShopRole"
	return b
}

//获取分店角色信息
func (b *BusShopRole) GetShopRole(ctx context.Context, args *bus.ArgsShopRole, reply *[]bus.ReplyShopRole) error {
	return b.Call(ctx, "GetShopRole",args,reply)
}

//获取分店角色用户信息
func (b *BusShopRole) GetShopRoleUser(ctx context.Context, args *bus.ArgsShopRoleUser, reply *bus.ReplyShopRoleUser) error {
	return b.Call(ctx, "GetShopRoleUser",args,reply)
}

//分配分店角色用户
func (b *BusShopRole) AddShopRoleUser(ctx context.Context,args *bus.ArgsAddShopUser,reply *bus.ReplyShopUser) error{
	return b.Call(ctx, "AddShopRoleUser",args,reply)
}
//编辑分店角色用户
func (b *BusShopRole) EditShopRoleUser(ctx context.Context,args *bus.ArgsEditShopUser,reply *bus.ReplyShopUser) error{
	return b.Call(ctx, "EditShopRoleUser",args,reply)
}

//删除分店角色用户
func (b *BusShopRole) DelShopRoleUser(ctx context.Context,args *bus.ArgsDelShopUser,reply *bus.ReplyShopUser) error{
	return b.Call(ctx, "DelShopRoleUser",args,reply)
}

//获取门店员工角色用户-rpc
func (b *BusShopRole) GetUserByShopIdAndStaffId(ctx context.Context,args *bus.ArgsCheckShopUser,reply *bus.ReplyCheckShopUser) error {
	return b.Call(ctx, "GetUserByShopIdAndStaffId",args,reply)
}