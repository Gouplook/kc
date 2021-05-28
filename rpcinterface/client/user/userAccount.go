package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

type UserAccount struct {
	client.Baseclient
}

func (u *UserAccount) Init() *UserAccount {
	u.ServiceName = "rpc_user"
	u.ServicePath = "UserAccount"
	return u
}

// UserAccountInfo 根据uid获取用户帐号信息
func (u *UserAccount) GetUserAccountInfo(ctx context.Context, uid *int, reply *user.ReplyUserAccountInfo) error {
	return u.Call(ctx, "GetUserAccountInfo", uid, reply)
}
