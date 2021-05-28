package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type Member struct {
	client.Baseclient
}

//初始化
func (m *Member) Init() *Member {
	m.ServiceName = "rpc_bus"
	m.ServicePath = "Member"
	return m
}

// 添加会员标签
func (m *Member) AddTag(ctx context.Context, args *bus.ArgsAddTag, reply *bool) error {
	return m.Call(ctx, "AddTag", args, reply)
}

// 移除会员标签
func (m *Member) RemoveTag(ctx context.Context, args *bus.ArgsRemoveTag, reply *bool) error {
	return m.Call(ctx, "RemoveTag", args, reply)
}

// 获取全部会员标签
func (m *Member) GetTagList(ctx context.Context, args *common.BsToken, reply *[]bus.ReplyTag) error {
	return m.Call(ctx, "GetTagList", args, reply)
}

// 添加会员
func (m *Member) AddMember(ctx context.Context, args *bus.ArgsMemberInfo, reply *bool) error {
	return m.Call(ctx, "AddMember", args, reply)
}

// 获取会员列表
func (m *Member) GetMemberList(ctx context.Context, args *bus.ArgsMemberParam, reply *bus.MemberList) error {
	return m.Call(ctx, "GetMemberList", args, reply)
}

// 获取会员信息接口
func (m *Member) GetMemberInfo(ctx context.Context, memberId *int, reply *bus.ReplyMemberInfo) error {
	return m.Call(ctx, "GetMemberInfo", memberId, reply)
}

// 获取会员详情接口
func (m *Member) GetMemberDetail(ctx context.Context, memberId *int, reply *bus.ReplyMemberDetail) error {
	return m.Call(ctx, "GetMemberDetail", memberId, reply)
}

// 平台会员快速成为门店会员
func (m *Member) UserToMember(ctx context.Context, args *bus.ArgsUserToMember, memberId *int) error {
	return m.Call(ctx, "UserToMember", args, memberId)
}

// 根据uid获取店铺会员折扣信息
func (m *Member) GetMemberRebateByUid(ctx context.Context, args *bus.ArgsMemberRebate, reply *bus.BusLevelDetail) error {
	return m.Call(ctx, "GetMemberRebateByUid", args, reply)
}

//根据手机号获取店铺会员信息-rpc
func (m *Member) GetUserInfoByPhoneRpc(ctx context.Context, args *bus.ArgsGetUserInfoByPhone, reply *bus.ReplyGetUserInfoByPhone) error {
	return m.Call(ctx, "GetUserInfoByPhoneRpc", args, reply)
}

// 获取会员基础档案信息
func (m *Member) GetMemberBase(ctx context.Context, memberId *int, reply *bus.ReplyMemberBase) error {
	return m.Call(ctx, "GetMemberBase", memberId, reply)
}

//编辑会员基础档案信息
func (m *Member) EditMemberBase(ctx context.Context, memberInfo *bus.ArgsMemberBase, reply *bool) error {
	return m.Call(ctx, "EditMemberBase", memberInfo, reply)
}

//根据用户ID查会员详情
func (m *Member) GetMemberByUid(ctx context.Context, args *bus.ArgsGetByUid, reply *bus.ReplyMemberDetail) (err error) {
	return m.Call(ctx, "GetMemberByUid", args, reply)
}

//发送会员短信入参
func (m *Member) SendMemberSms(ctx context.Context, args *bus.ArgsSendMemberSms, reply *bool) error {
	return m.Call(ctx, "SendMemberSms", args, reply)
}

//确认消费完成,添加店铺会员消费数据统计
func (m *Member) AddMemberConsumeCount(ctx context.Context, consumeLogId *int, reply *bus.ReplyAddMemberConsumeCount) error {
	return m.Call(ctx, "AddMemberConsumeCount", consumeLogId, reply)
}

//获取店铺会员消费数据
func (m *Member) GetMemberConsumeCount(ctx context.Context, memberId *int, reply *bus.ReplyGetMemberConsumeCount) error {
	return m.Call(ctx, "GetMemberConsumeCount", memberId, reply)
}

//购买卡包后，更新会员持卡的量,如果会员不存在则增加一条数据
func (m *Member)AddMemberCardNum(ctx context.Context,relationId *int,reply *bus.ReplyAddMemberConsumeCount)error{
	return m.Call(ctx, "AddMemberCardNum", relationId, reply)
}

//支付成功后，用户自动加入商家会员
func (m *Member) PayUserJoin(ctx context.Context, orderSn *string, reply *bool ) (err error)  {
	return m.Call(ctx, "PayUserJoin", orderSn, reply)
}
// 获取会员总人数
func (m *Member) GetMemberNum(ctx context.Context, args *bus.ArgsRiskMemberNum, reply *bus.ReplyRiskMemberNum) error {
	return m.Call(ctx, "GetMemberNum",args, reply)
}

//根据指定店铺(shopId 有值就查询shopId)的多个uid查询会员信息  -rpc
func (m *Member) GetMemberInfos(ctx context.Context, args *bus.ArgsGetMemberInfo, reply *[]bus.ReplyGetMemberInfo) error {
	return m.Call(ctx, "GetMemberInfos",args, reply)
}

//根据手机号模糊匹配门店会员
func (m *Member)GetUserInfoByPhoneMatch(ctx context.Context,args *bus.ArgsGetUserInfoByPhoneMatch,reply *bus.ReplyGetUserInfoByPhoneMatch)error{
	return m.Call(ctx, "GetUserInfoByPhoneMatch",args, reply)
}