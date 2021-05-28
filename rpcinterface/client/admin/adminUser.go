package admin

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/admin"
)

// AdminUser AdminUser
type AdminUser struct {
	client.Baseclient
}

func (adminUser *AdminUser) Init() *AdminUser {
	adminUser.ServiceName = "rpc_admin"
	adminUser.ServicePath = "AdminUser"
	return adminUser
}

// AddAdminUser 新增管理员
func (adminUser *AdminUser) AddAdminUser(ctx context.Context, args *admin.ArgsAddAdminUser, reply *admin.ReplyAddAdminUser) error {
	return adminUser.Call(ctx, "AddAdminUser", args, reply)
}

// AdminUserLogin 管理员登陆
func (adminUser *AdminUser) AdminUserLogin(ctx context.Context, args *admin.ArgsAdminUserLogin, reply *admin.ReplyAdminUserLogin) error {
	return adminUser.Call(ctx, "AdminUserLogin", args, reply)
}

// UpdateAdminUser 修改管理员
func (adminUser *AdminUser) UpdateAdminUser(ctx context.Context, args *admin.ArgsUpdateAdminUser, reply *bool) error {
	return adminUser.Call(ctx, "UpdateAdminUser", args, reply)
}

// DeleteAdminUser 删除管理员
func (adminUser *AdminUser) DeleteAdminUser(ctx context.Context, args *admin.ArgsDeleteAdminUser, reply *bool) error {
	return adminUser.Call(ctx, "DeleteAdminUser", args, reply)
}

// AdminUserList 管理员列表
func (adminUser *AdminUser) AdminUserList(ctx context.Context, args *admin.ArgsAdminUserList, reply *admin.ReplyAdminUserList) error {
	return adminUser.Call(ctx, "AdminUserList", args, reply)
}

// AdminUserInfo 管理员详情
func (adminUser *AdminUser) AdminUserInfo(ctx context.Context, args *admin.ArgsAdminUserInfo, reply *admin.ReplyAdminUserInfo) error {
	return adminUser.Call(ctx, "AdminUserInfo", args, reply)
}

//	CheckLogin 验证登陆
func (adminUser *AdminUser) AuthLogin(ctx context.Context, args *admin.ArgsAuthLogin, reply *admin.ReplyAuthLogin) error {
	return adminUser.Call(ctx, "AuthLogin", args, reply)
}

//	LoginOut 退出登录
func (adminUser *AdminUser) LoginOut(ctx context.Context, args *admin.ArgsLoginOut, reply *bool) error {
	return adminUser.Call(ctx, "LoginOut", args, reply)
}

// AdminListByIds 根据ids批量获取管理员信息
func (adminUser *AdminUser)AdminListByIds(ctx context.Context,args *admin.ArgsAdminListByIds,reply *admin.ReplyAdminListByIds)error{
	return adminUser.Call(ctx, "AdminListByIds", args, reply)
}