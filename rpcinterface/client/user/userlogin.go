package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

type UserLogin struct {
	client.Baseclient
}

func (u *UserLogin) Init() *UserLogin {
	u.ServiceName = "rpc_user"
	u.ServicePath = "UserLogin"
	return u
}

//注册
func (u *UserLogin) Reg(ctx context.Context, args *user.RegParams, reply *user.RegReply) error {
	return u.Call(ctx, "Reg", args, reply)
}

//登录
func (u *UserLogin) Login(ctx context.Context, args *user.LoginParams, reply *user.LoginReply) error {
	return u.Call(ctx, "Login", args, reply)
}

//验证登录
func (u *UserLogin) CheckLogin(ctx context.Context, args *user.CheckLoginParams, reply *user.CheckLoginReply) error {
	return u.Call(ctx, "CheckLogin", args, reply)
}

//找回密码
func (u *UserLogin) FindPwd(ctx context.Context, args *user.FindPwdParams, reply *user.FindPwdReply) error {
	return u.Call(ctx, "FindPwd", args, reply)
}

// 微信登录
func (u *UserLogin) WxLogin(ctx context.Context, args *user.ArgsWxLogin, reply *user.ReplyWxLogin) error {
	return u.Call(ctx, "WxLogin", args, reply)
}

// 微信登录绑定账号
func (u *UserLogin) WxBind(ctx context.Context, args *user.ArgsWxBind, reply *user.LoginReply) error {
	return u.Call(ctx, "WxBind", args, reply)
}

// 快速注册
func (u *UserLogin) FastReg(ctx context.Context, args *user.RegFast, reply *user.RegReply) error {
	return u.Call(ctx, "FastReg", args, reply)
}

// 退出登录
func (u *UserLogin) LoginOut(ctx context.Context, args *user.ArgsLoginOut, reply *bool) error {
	return u.Call(ctx, "LoginOut", args, reply)
}

//获取用户的微信openid
func (u *UserLogin) GetWxOpenid(ctx context.Context, args *user.ArgsGetWxOpenid, reply *user.ReplyGetWxOpenid ) error {
	return u.Call(ctx, "GetWxOpenid", args, reply)
}

//微信小程序登录
func (u *UserLogin)  WxMappLogin(ctx context.Context, args *user.ArgsWxMappLogin, reply *user.ReplyWxMappLogin  ) error {
	return u.Call(ctx, "WxMappLogin", args, reply)
}