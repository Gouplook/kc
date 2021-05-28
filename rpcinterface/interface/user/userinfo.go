/**
 * @Author: Gosin
 * @Date: 2019/12/12 15:38
 */
package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//性别
	SEX_unknow = 0
	SEX_male   = 1
	SEX_female = 2

	//实名认证
	REAL_NAME_AUTH_no  = 0
	REAL_NAME_AUTH_yes = 1
)

type ReplyUserinfo struct {
	Uid            int    // uid
	Nick           string // 用户昵称
	AvatarImgId    int    // 图片id
	AvatarUrl      string //头像地址
	Sex            int    // 性别 0=未定义  1=女 2=男
	IsRealNameAuth int    // 实名认证 0=未认证  1=已认证
	Name           string // 真实姓名
	Phone          string //手机号
	PreLoginTime   int    //上次登录时间
	Ctime          int64  // 注册时间
}

//用户基础信息
type ReplyBaseUserinfo struct {
	Uid         int    // uid
	Nick        string // 用户昵称
	AvatarImgId int    // 图片id
	Sex         int    // 性别 0=未定义  1=女 2=男
}
type ArgsUpdateUserinfoByUid struct {
	Uid           int    // uid 必须
	Nick          string // 用户昵称
	AvatarImgHash string // 图片hash
	Sex           int    // 性别 0=未定义  1=女 2=男
}

type ArgsUpdateUserinfoByUtoken struct {
	Utoken        common.Utoken //用户登录token
	Nick          string        // 用户昵称
	AvatarImgHash string        // 图片hash
	Sex           int           // 性别 0=未定义  1=女 2=男

}

//================后台接口================
// ArgsGetUserInfoByPhone ArgsGetUserInfoByPhone
type ArgsGetUserInfoByPhone struct {
	Phone string
}

// ReplyGetUserInfoByPhone ReplyGetUserInfoByPhone
type ReplyGetUserInfoByPhone struct {
	Phone string
	Uid   int
}

// ArgsGetUserInfoByNick ArgsGetUserInfoByNick
type ArgsGetUserInfoByNick struct {
	Nick string
}

// 会员列表入参
type ArgsMemberList struct {
	common.Paging
	Nick       string // 用户名
	Phone      string // 手机号
	Type       string // 用户类型
	Status     string // 用户状态
	CtimeStart int64  // 注册开始时间
	CtimeEnd   int64  // 注册结束时间
}
type MemberListBase struct {
	Uid            int     // 用户UID
	Nick           string  // 用户名
	Phone          string  // 手机号
	Name           string  // 真实姓名
	Email          string  // 邮箱
	Type           int     // 用户类型
	AccountBalance float64 // 帐户余额
	CardNo         string  // 身份证号码
	Ctime          int64   // 注册时间
	CtimeStr       string  // 注册时间字符串
	Ip             string  // 注册IP
}

// 会员列表出参
type ReplyMemberList struct {
	TotalNum int
	List     []MemberListBase
}

// ArgsUpdateMemberStatus 更新会员状态入参
type ArgsUpdateMemberStatus struct {
	Uid    int
	Status int // 会员状态:开启/禁用
}

// 会员详情出参
type ReplyMemberInfo struct {
	Uid            int
	Nick           string  // 用户名
	Phone          string  // 手机号
	Email          string  // 邮箱
	Name           string  // 真实姓名
	Type           int     // 会员/用户类型
	Address        string  // 详细地址
	QqNo           string  // qq帐号
	CardNo         string  // 身份证号
	Sex            int     // 性别 0=未定义  1=女 2=男
	IsRealNameAuth int     // 实名认证 0=未认证  1=已认证
	AccountBalance float64 // 帐户余额
}

type ReplyUserInfo struct {
	Uid int
	Nick string //昵称
	Name string //真实姓名
	Phone string
	AvatarImgId int
	Sex int
}

type Userinfo interface {
	// 根据uid获取用户单条详细数据
	GetUserinfoByUid(ctx context.Context, uid *int, reply *ReplyUserinfo) error
	// 根据uids获取用户基础数据
	GetBaseUserinfoByUids(ctx context.Context, uids *[]int, reply *[]ReplyBaseUserinfo) error
	// 修改用户信息
	UpdateUserinfoByUid(ctx context.Context, args *ArgsUpdateUserinfoByUid, reply *bool) error
	//更具登录信息获取用户单条详细数据
	GetUserinfoByUtoken(ctx context.Context, utoken *common.Utoken, reply *ReplyUserinfo) error
	// 修改当前登录用户的信息
	UpdateUserinfoByUtoken(ctx context.Context, args *ArgsUpdateUserinfoByUtoken, reply *bool) error
	// 根据用户手机号获取用户信息
	GetUserInfoByPhone(ctx context.Context, args *ArgsGetUserInfoByPhone, reply *ReplyGetUserInfoByPhone) error
	// 根据nick获取用户信息
	GetUserInfoByNick(ctx context.Context, args *ArgsGetUserInfoByNick, reply *[]ReplyBaseUserinfo) error
	// 会员列表
	GetMemberList(ctx context.Context, args *ArgsMemberList, reply *ReplyMemberList) error
	// 会员详情
	GetMemberInfo(ctx context.Context, uid *int, reply *ReplyMemberInfo) error
	// 更改会员状态
	UpdateMemberStatus(ctx context.Context, args *ArgsUpdateMemberStatus, reply *bool) error
	//根据多个用户id查询用户信息-rpc
	GetUserInfosByUids(ctx context.Context,uids *[]int, reply *[]ReplyUserInfo) error
}
