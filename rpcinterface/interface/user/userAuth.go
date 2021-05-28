package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type ReplyUserAuthInfo struct {
	Uid        int
	Name       string // 真实姓名
	CardNo     string // 身份证号码
	CardNoExpire string //身份证过期日期、格式：2037-10-01
	FaceData   string // 人脸图片base64编码数据
	Similarity int    // 人脸相识度 1-100 值越高越准确
	Ctime      int64  // 认证时间
	CtimeStr   string
}

//实名认证入参
type ArgsIdentity struct {
	common.Utoken   //用户登录信息
	RealName string //真实姓名
	CardNo   string //身份证号
	Uid int //rpc内部校验
}

type UserAuth interface {
	// GetUserAuthInfoByUid 根据uid获取用户认证信息
	GetUserAuthInfoByUid(ctx context.Context, uid *int, reply *ReplyUserAuthInfo) error
}
