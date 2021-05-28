package user

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
)

type User struct {
	client.Baseclient
}

func (u *User) Init() *User {
	u.ServiceName = "rpc_task"
	u.ServicePath = "User/User"
	return u
}

func (u *User) UserReg(ctx context.Context, uid *int, reply *bool) error {
	return u.Call(ctx, "UserReg", uid, reply)
}

// 用户信息
func (u *User) SetUserInfo(ctx context.Context, uid *int, reply *bool) error {
	return u.Call(ctx, "SetUserInfo", uid, reply)
}

//comtree data consumer table
func (u *User) ConsumerComtreeData(ctx context.Context, uid *int, reply *bool) error {
	return u.Call(ctx, "ConsumerComtreeData", uid, reply)
}
