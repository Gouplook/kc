package admin

import (
	"context"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	// 启用状态
	ENABLE_STATUS = 1
	// 删除状态
	DELETE_STATUS = 2

	SESSION_EXPIRE int64 = 2592000 // 过期时间暂定30天

	CHANNEL_unknown = 0 // 未知
	CHANNEL_pc      = 1 // pc网站

	DEVICE_unknown = 0 // 未知
	DEVICE_pc      = 1 // pc
)

// 设备数组
func GetDevices() []int {
	return []int{
		DEVICE_unknown,
		DEVICE_pc,
	}
}

// 渠道数组
func GetChannel() []int {
	return []int{
		CHANNEL_unknown,
		CHANNEL_pc,
	}
}

type Channel = int // 渠道  0=未知， 1=pc网站
type Device = int  // 设备  0=未知  1=PC

// 验证渠道
func VerifyChannel(channel Channel) error {
	if functions.InArray(channel, GetChannel()) {
		return nil
	}
	return common.GetInterfaceError(common.CHANNEL_ERR)
}

// 验证设备
func VerifyDevice(device Device) error {
	if functions.InArray(device, GetDevices()) {
		return nil
	}
	return common.GetInterfaceError(common.DEVICE_ERR)
}

// 用户登录Session
type Session struct {
	AdminUserId int // 用户uid
	RoleId      int
	LoginTime   int64 // 第一次登录时间
	Refresh     int64 // 刷新时间
	ExpireTime  int64 // session过期时间
	//Ip         string // 客户端ip
	Device
}

// ArgsAdminUserBase 后台管理员基础信息
type ArgsAdminUserBase struct {
	AdminUserId int    `mapstructure:"admin_user_id"` // 用户uid
	Name        string `mapstructure:"name"`          // 用户名
	NickName    string `mapstructure:"nick_name"`     // 昵称
	RoleId      int    `mapstructure:"role_id"`       // 所属角色id
	Email       string `mapstructure:"email"`         // 邮箱
	Ctime       int    `mapstructure:"c_time"`        // 创建时间
	Utime       int    `mapstructure:"u_time"`        // 更新时间
	CtimeStr    string `mapstructure:"ctime_str"`
	Status      int    `mapstructure:"status"` // 状态:默认为1启用状态
}

// ArgsAddAdminUser 新增管理员入参
type ArgsAddAdminUser struct {
	common.Autoken //用户登录token
	ArgsAdminUserBase
	Password string
}

// ReplyAddAdminUser 新增管理员返回参数
type ReplyAddAdminUser struct {
	AdminUserId int
}

// ArgsUpdateAdminUser 修改管理员入参
type ArgsUpdateAdminUser struct {
	common.Autoken //用户登录token
	ArgsAdminUserBase
	Password    string
	AdminUserId int
}

// ReplyUpdateAdminUser 修改管理员返回参数
type ReplyUpdateAdminUser struct {
	AdminUserId int
}

// AddUpdateAdminUserBase AddUpdateAdminUserBase
type AddUpdateAdminUserBase struct {
	ArgsAdminUserBase
	Password    string
	AdminUserId int
}

// ArgsDeleteAdminUser 删除管理员入参
type ArgsDeleteAdminUser struct {
	common.Autoken //用户登录token
	AdminUserId    int
}

// ReplyDeleteAdminUser 删除管理员返回参数
type ReplyDeleteAdminUser struct {
	AdminUserId int
}

// ArgsAdminUserList 管理员列表入参
type ArgsAdminUserList struct {
	common.Paging
	Name        string
	AdminUserId int
}

// ReplyAdminUserList 管理员列表返回参数
type ReplyAdminUserList struct {
	TotalNum int
	List     []ArgsAdminUserBase
}

// ArgsAdminUserInfo 管理员详情入参
type ArgsAdminUserInfo struct {
	AdminUserId int
}

// ReplyAdminUserInfo 管理员详情返回参数
type ReplyAdminUserInfo struct {
	ArgsAdminUserBase
}

// ArgsAdminUserLogin 管理员登陆
type ArgsAdminUserLogin struct {
	Email    string
	Password string
	Channel
	Device
	//Ip string
}
type ReplyAdminUserLogin struct {
	AdminUserId int // 用户id
	RoleId      int
	Token       string // 登录token
	ArgsAdminUserBase
}

// 验证登录参数
type ArgsAuthLogin struct {
	Channel
	Device
	Token string // 登录token
}

// 验证登录返回数据
type ReplyAuthLogin struct {
	EncodeStr string //加密后的uid
	ArgsAdminUserBase
}

// 退出登录
type ArgsLoginOut struct {
	AdminAuth
	Token string // 用户登录的token
}
type AdminAuth struct {
	common.Autoken // 加密后的uidToken==>UidEncodeStr
	Channel
	Device
}

// ArgsAdminListByIds ArgsAdminListByIds
type ArgsAdminListByIds struct {
	AdminAuth
	AdminIds []int
}
type ReplyAdminListByIds struct {
	List []ArgsAdminUserBase
}

// AdminUser AdminUser
type AdminUser interface {
	// AddAdminUser 新增管理员
	AddAdminUser(ctx context.Context, args *ArgsAddAdminUser, reply *ReplyAddAdminUser) error
	// AdminUserLogin 管理员登陆
	AdminUserLogin(ctx context.Context, args *ArgsAdminUserLogin, reply *ReplyAdminUserLogin) error
	// UpdateAdminUser 修改管理员
	UpdateAdminUser(ctx context.Context, args *ArgsUpdateAdminUser, reply *bool) error
	// DeleteAdminUser 删除管理员
	DeleteAdminUser(ctx context.Context, args *ArgsDeleteAdminUser, reply *bool) error
	// AdminUserList 管理员列表
	AdminUserList(ctx context.Context, args *ArgsAdminUserList, reply *ReplyAdminUserList) error
	// AdminUserInfo 管理员详情
	AdminUserInfo(ctx context.Context, args *ArgsAdminUserInfo, reply *ReplyAdminUserInfo) error
	//	AuthLogin 验证登陆
	AuthLogin(ctx context.Context, args *ArgsAuthLogin, reply *ReplyAuthLogin) error
	//	LoginOut 退出登录
	LoginOut(ctx context.Context, args *ArgsLoginOut, reply *bool) error
	//	AdminListByIds 根据ids批量获取管理员信息
	AdminListByIds(ctx context.Context, args *ArgsAdminListByIds, reply *ReplyAdminListByIds) error
}
