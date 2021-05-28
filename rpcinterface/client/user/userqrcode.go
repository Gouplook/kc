package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

// @author liyang<654516092@qq.com>
// @date  2020/8/7 11:34

type UserQrcode struct {
	client.Baseclient
}

func (u *UserQrcode) Init() *UserQrcode {
	u.ServiceName = "rpc_user"
	u.ServicePath = "UserQrcode"
	return u
}

//获取用户二维码信息
func (u *UserQrcode) GetUserQrcode(ctx context.Context, args *user.ArgsUserQrcode, reply *user.ReplyUserQrcode) error {
	return u.Call(ctx, "GetUserQrcode", args, reply)
}



