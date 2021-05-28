package user

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

type UserAuth struct {
	client.Baseclient
}

func (u *UserAuth) Init() *UserAuth {
	u.ServiceName = "rpc_user"
	u.ServicePath = "UserAuth"
	return u
}

//GetUserAuthInfoByUid 根据uid获取用户认证信息
func (u *UserAuth) GetUserAuthInfoByUid(ctx context.Context, uid *int, reply *user.ReplyUserAuthInfo) error {
	return u.Call(ctx, "GetUserAuthInfoByUid", uid, reply)
}

//用户实名认证
func (u *UserAuth) CheckIdentity(ctx context.Context, args *user.ArgsIdentity, reply *int) error {
	return u.Call(ctx, "CheckIdentity", args, reply)
}

//用户实名认证-rpc
func (u *UserAuth) CheckIdentityRpc(ctx context.Context, args *user.ArgsIdentity, reply *int) error {
	return u.Call(ctx, "CheckIdentityRpc", args, reply)
}
