/**
 * @Author: Gosin
 * @Date: 2020/4/7 18:24
 */
package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
)

type Member struct {
	client.Baseclient
}

func (m *Member) Init() *Member {
	m.ServiceName = "rpc_elastic"
	m.ServicePath = "Bus/Member"
	return m
}

// 添加会员信息
func (m *Member) AddMember(ctx context.Context, memberId *int, reply *bool) error {
	return m.Call(ctx, "AddMember", memberId, reply)
}

// 店铺会员检索
func (m *Member) SearchMember(ctx context.Context, args *bus2.ArgsMemberParam, reply *bus2.MemberList) error {
	return m.Call(ctx, "SearchMember", args, reply)
}

func (m *Member) EditMember(ctx context.Context, memberBase *bus2.ArgsMemberBase, reply *bool) error {
	return m.Call(ctx, "EditMember", memberBase, reply)
}

//更新会员的累计消费次数，累计消费金额，最后消费时间,持卡量 到es中
func (m *Member)SetMemberConsumeCount(ctx context.Context,memberId *int,reply *bool)error{
	return m.Call(ctx, "SetMemberConsumeCount", memberId, reply)
}

//根据手机号模糊匹配门店会员
func (m *Member)GetUserInfoByPhoneMatch(ctx context.Context,args *bus2.ArgsGetUserInfoByPhoneMatch,reply *bus2.ReplyGetUserInfoByPhoneMatch)error{
	return m.Call(ctx, "GetUserInfoByPhoneMatch",args, reply)
}