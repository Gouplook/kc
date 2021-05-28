/**
 * @Author: Gosin
 * @Date: 2020/4/3 14:48
 */
package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

type Member interface {
	// 添加会员信息
	AddMember(ctx context.Context, memberId *int, reply *bool) error
	// 修改会员信息
	EditMember(ctx context.Context, memberBase *bus.ArgsMemberBase, reply *bool) error
	// 店铺会员检索
	SearchMember(ctx context.Context, args *bus.ArgsMemberParam, reply *bus.MemberList) error
	//更新会员的累计消费次数，累计消费金额，最后消费时间，持卡量到es中
	SetMemberConsumeCount(ctx context.Context,memberId *int,reply *bool)error
}
