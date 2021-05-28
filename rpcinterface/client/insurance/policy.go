package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/insurance"
)
// 保单
// @author liyang<654516092@qq.com>
// @date  2020/7/28 10:57

type Policy struct {
	client.Baseclient
}

func (p *Policy) Init() *Policy {
	p.ServiceName = "rpc_insurance"
	p.ServicePath = "Policy"
	return p
}

//获取预付卡保单列表
func (p *Policy) GetPolicyByRelationId(ctx context.Context, args *insurance.ArgsCardpackagePolicy, reply *[]insurance.ReplyCardPackageList) error {
	return p.Call(ctx, "GetPolicyByRelationId", args, reply)
}

//获取预付卡保单详情
func (p *Policy) GetPolicyById(ctx context.Context, args *insurance.ArgsSinglePolicy, reply *insurance.ReplySinglePolicy) error {
	return p.Call(ctx, "GetPolicyById", args, reply)
}

//获取用户累计保额
func (p *Policy) GetUserByTotoalInsurancAmount(ctx context.Context, args *insurance.ArgsInsuranceUser, reply *insurance.ReplyInsuranceUser) error {
	return p.Call(ctx, "GetUserByTotoalInsurancAmount", args, reply)
}

// 获取保单任务信息
func (p *Policy)GetPolicyTaskInfo(ctx context.Context, args *insurance.ArgsPolicyTask, reply *insurance.ReplyPolicyTask) error {
	return  p.Call(ctx, "GetPolicyTaskInfo", args, reply)
}

// 获取续保任务信息
func (p *Policy)GetRenewPolicyTaskInfo(ctx context.Context, args *insurance.ArgsPolicyTask, reply *insurance.ReplyRenewPolicyTask) error{
	return p.Call(ctx, "GetRenewPolicyTaskInfo", args, reply)
}
