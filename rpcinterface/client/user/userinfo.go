package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

type Userinfo struct {
	client.Baseclient
}

func (userinfo *Userinfo) Init() *Userinfo {
	userinfo.ServiceName = "rpc_user"
	userinfo.ServicePath = "Userinfo"
	return userinfo
}

func (userinfo *Userinfo) GetUserinfoByUid(ctx context.Context, uid *int, reply *user.ReplyUserinfo) error {
	return userinfo.Call(ctx, "GetUserinfoByUid", uid, reply)
}

func (userinfo *Userinfo) GetBaseUserinfoByUids(ctx context.Context, uids *[]int, reply *[]user.ReplyBaseUserinfo) error {
	return userinfo.Call(ctx, "GetBaseUserinfoByUids", uids, reply)
}

func (userinfo *Userinfo) UpdateUserinfoByUid(ctx context.Context, args *user.ArgsUpdateUserinfoByUid, reply *bool) error {
	return userinfo.Call(ctx, "UpdateUserinfoByUid", args, reply)
}

func (userinfo *Userinfo) GetUserinfoByUtoken(ctx context.Context, utoken *common.Utoken, reply *user.ReplyUserinfo) error {
	return userinfo.Call(ctx, "GetUserinfoByUtoken", utoken, reply)
}

func (userinfo *Userinfo) UpdateUserinfoByUtoken(ctx context.Context, args *user.ArgsUpdateUserinfoByUtoken, reply *bool) error {
	return userinfo.Call(ctx, "UpdateUserinfoByUtoken", args, reply)
}

// 根据用户手机号获取用户信息
func (userinfo *Userinfo) GetUserInfoByPhone(ctx context.Context, args *user.ArgsGetUserInfoByPhone, reply *user.ReplyGetUserInfoByPhone) error {
	return userinfo.Call(ctx, "GetUserInfoByPhone", args, reply)
}

// 根据nick获取用户信息
func (userinfo *Userinfo) GetUserInfoByNick(ctx context.Context, args *user.ArgsGetUserInfoByNick, reply *[]user.ReplyBaseUserinfo) error {
	return userinfo.Call(ctx, "GetUserInfoByNick", args, reply)
}

// 会员列表
func (userinfo *Userinfo) GetMemberList(ctx context.Context, args *user.ArgsMemberList, reply *user.ReplyMemberList) error {
	return userinfo.Call(ctx, "GetMemberList", args, reply)
}

// 会员详情
func (userinfo *Userinfo) GetMemberInfo(ctx context.Context, uid *int, reply *user.ReplyMemberInfo) error {
	return userinfo.Call(ctx, "GetMemberInfo", uid, reply)
}

// 更改会员状态
func (userinfo *Userinfo) UpdateMemberStatus(ctx context.Context, args *user.ArgsUpdateMemberStatus, reply *bool) error {
	return userinfo.Call(ctx, "UpdateMemberStatus", args, reply)
}

//根据多个用户id查询用户信息-rpc
func (userinfo *Userinfo) GetUserInfosByUids(ctx context.Context, uids *[]int, reply *[]user.ReplyUserInfo) error {
	return userinfo.Call(ctx, "GetUserInfosByUids", uids, reply)
}