package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

type Member interface {
	// SetMember 设置用户会员到es
	SetMember(ctx context.Context, uid *int, reply *bool) error
	// SearchMember 搜索用户会员-用于后台搜索
	SearchMember(ctx context.Context, args *user.ArgsMemberList, reply *user.ReplyMemberList) error
	//根据手机号模糊匹配门店会员
	GetUserInfoByPhoneMatch(ctx context.Context,args *user.ArgsGetUserInfoByPhoneMatch,reply *user.ReplyGetUserInfoByPhoneMatch)error
}
