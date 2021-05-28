package user

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/user"
)

type Member struct {
	client.Baseclient
}

func (m *Member) Init() *Member {
	m.ServiceName = "rpc_elastic"
	m.ServicePath = "User/Member"
	return m
}

// SetMember 设置用户会员到es
func (m *Member) SetMember(ctx context.Context, uid *int, reply *bool) error {
	return m.Call(ctx, "SetMember", uid, reply)
}

// SearchMember 搜索用户会员-用于后台搜索
func (m *Member) SearchMember(ctx context.Context, args *user.ArgsMemberList, reply *user.ReplyMemberList) error {
	return m.Call(ctx, "SearchMember", args, reply)
}
