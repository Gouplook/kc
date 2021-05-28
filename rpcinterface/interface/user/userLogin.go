package user

// 用户登录注册相关的接口定义
// @author yangzhiwu<wyz@900sui.com>
// @date 2020-03-04 10:33:20

import (
	"context"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	LoginType_captcha = 1 // 验证码登录
	LoginType_pwd     = 2 // 密码登录

	SESSION_EXPIRE int64 = 2592000 // 过期时间暂定30天

	CHANNEL_unknown          = 0 // 未知
	CHANNEL_pc               = 1 // pc的Saas系统
	CHANNEL_900sui_app       = 2 // 900岁app
	CHANNEL_kangxiangbao_app = 3 // 康享宝app
	CHANNEL_900sui_wap       = 4 // 900岁wap
	CHANNEL_kadidou_wxapp    = 5 // 卡D兜小程序
	CHANNEL_kadidou_wxofficial=6 // 微信公众号

	DEVICE_unknown  = 0 // 未知
	DEVICE_pc       = 1 // pc
	DEVICE_android  = 2 // android
	DEVICE_ios      = 3 // ios
	DEVICE_winphone = 4 // winphone
	DEVICE_mac      = 5 // mac

	//#登录场景
	//普通登录
	LOGIN_scene_normal = 0
	//主体登录
	LOGIN_scene_bus    = 1
	//分店登录
	LOGIN_scene_shop   = 2
)

// 登录场景数组
func GetLoginScene()[]int{
	return []int{
		LOGIN_scene_normal,
		LOGIN_scene_bus,
		LOGIN_scene_shop,
	}
}

// 设备数组
func GetDevices() []int {
	return []int{
		DEVICE_unknown,
		DEVICE_pc,
		DEVICE_android,
		DEVICE_ios,
		DEVICE_winphone,
		DEVICE_mac,
	}
}

// 渠道数组
func GetChannel() []int {
	return []int{
		CHANNEL_unknown,
		CHANNEL_pc,
		CHANNEL_900sui_app,
		CHANNEL_kangxiangbao_app,
		CHANNEL_900sui_wap,
		CHANNEL_kadidou_wxapp,
		CHANNEL_kadidou_wxofficial,
	}
}

type Channel = int // 注册渠道  0=未知， 1=pc网站 2=900岁app 3=康享宝app 4=900岁wap版，5=卡D兜小程序
type Device = int  // 注册设备  0=未知  1=PC 2=Android 3=Ios 4=WinPhone  5=Mac

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
	Uid        int    // 用户uid
	LoginTime  int64  // 第一次登录时间
	Refresh    int64  // 刷新时间
	ExpireTime int64  // session过期时间
	Ip         string // 客户端ip
	Device
}

// 注册所需参数
type RegParams struct {
	Phone    string // 手机
	Captcha  string // 短信验证码
	Password string // 密码
	Channel
	Device
}

// 快速注册
type RegFast struct {
	Phone string // 手机
	Channel
	Device
}

// 注册返回数据
type RegReply struct {
	Uid int // 用户id
}

// 登录参数
type LoginParams struct {
	Phone     string
	LoginScene int   // 登录场景 0 =普通登录 1=saas总部(主体) 2=分店登录
	LoginType int    // 登录方式 1=短信验证码登录(常量 LoginType_captcha) 2=密码登录（常量 LoginType_pwd）
	Captcha   string // 短信验证码
	Password  string // 密码
	Channel
	Device
	Ip string // 客户端ip
}

// 登录返回数据
type LoginReply struct {
	Uid   int    // 用户id
	Token string // 登录token
}

// 验证登录参数
type CheckLoginParams struct {
	Channel
	Token string // 登录token
}

// 验证登录返回数据
type CheckLoginReply struct {
	UidEncodeStr string //加密后的uid
	Nick         string // 用户昵称
	RealNameAuth int    // 是否实名认证
}

// 忘记密码参数
type FindPwdParams struct {
	Phone    string // 手机号
	Captcha  string // 短信验证码
	Password string // 新密码
}

// 忘记密码返回数据
type FindPwdReply struct {
	Ok bool // 修改状态
}

type ArgsWxLogin struct {
	Appid string
	Code  string
	Channel
	Ip string
}

type ArgsWxBind struct {
	OpenHash string // 未绑定会员时的临时验证
	LoginParams
}

type ReplyWxLogin struct {
	OpenHash   string // 未绑定会员时的临时验证
	OpenId     string // 用户openId
	LoginReply        // 登录成功返回数据
}

type ArgsLoginOut struct {
	common.Utoken
	Channel
	Token string //登录字符串
}

//获取用户微信openid接口入参
type ArgsGetWxOpenid struct {
	Appid string //微信应用的appid
	Code  string //微信返回的用于获取openid的临时code
}

//获取用户微信openid接口返回数据
type  ReplyGetWxOpenid struct {
	Openid string
}

//微信小程序登录
type ArgsWxMappLogin struct {
	Appid string //小程序id
	Code  string //小程序临时code
	Channel //渠道id
	Device //设备
	Ip string //客户端ip
	Iv string  //手机号解密iv 非必传
	EncryptedData string //手机号加密数据 非必传
}

//微信小程序登录返回
type ReplyWxMappLogin struct {
	OpenId     	string // 用户openId
	Token 		string // 登录token
	ReplyUserinfo //用户信息
}

type UserLogin interface {
	// 注册
	Reg(ctx context.Context, args *RegParams, reply *RegReply) error
	// 登录
	Login(ctx context.Context, args *LoginParams, reply *LoginReply) error
	// 验证登录
	CheckLogin(ctx context.Context, args *CheckLoginParams, reply *CheckLoginReply) error
	// 找回密码
	FindPwd(ctx context.Context, args *FindPwdParams, reply *FindPwdReply) error
	// 微信登录
	WxLogin(ctx context.Context, args *ArgsWxLogin, reply *ReplyWxLogin) error
	// 微信登录绑定账号
	WxBind(ctx context.Context, args *ArgsWxBind, reply *LoginReply) error
	//快速注册(已注册则返回用户UID)
	FastReg(ctx context.Context, args *RegFast, reply *RegReply) error
	//退出登录
	LoginOut(ctx context.Context, args *ArgsLoginOut, reply *bool) error
	//获取用户的微信openid
	GetWxOpenid(ctx context.Context, args *ArgsGetWxOpenid, reply *ReplyGetWxOpenid ) error
	//微信小程序登录
	WxMappLogin(ctx context.Context, args *ArgsWxMappLogin, reply *ReplyWxMappLogin  ) error
}
