package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

// 定义rpc客户端调用方法
// @author liyang<654516092@qq.com>
// @date  2020/4/7 17:12

type Rules struct {
	client.Baseclient
}
//初始化
func (r *Rules)Init() *Rules {
	r.ServiceName = "rpc_staff"
	r.ServicePath = "Rules"
	return r
}

//添加行为规则
func (r *Rules) AddConductRules(ctx context.Context, args *staff.ArgsAddConductRules, reply *staff.ReplyConductRules) error {
	return r.Call(ctx, "AddConductRules", args, reply)
}

//更新行为规则
func (r *Rules) SetConductRules(ctx context.Context,args *staff.ArgsSetConductRules,reply *staff.ReplyConductRules) error{
	return r.Call(ctx, "SetConductRules", args, reply)
}

func (r *Rules) DeleteConductRules(ctx context.Context,args *staff.ArgsDelConductRules,reply *staff.ReplyConductRules) error{
	return r.Call(ctx, "DeleteConductRules", args, reply)
}

//获取行为规则列表
func (r *Rules) GetConductRules(ctx context.Context,args *staff.ArgsGetConductRules,reply *[]staff.ReplyConductRulesInfo) error{
	return r.Call(ctx, "GetConductRules", args, reply)
}

//添加行为规则记录
func (r *Rules) AddConductRulesLog(ctx context.Context,args *staff.ArgsAddConductRulesLog,reply *staff.ReplyAddConductRulesLog) error{
	return r.Call(ctx,"AddConductRulesLog",args,reply)
}

//获取行为规范记录统计-企业/商户
func (r *Rules) GetConductRulesLogDatasForBus(ctx context.Context,args *staff.ArgsGetConductRuleslogData,reply *staff.ReplyGetConductRuleslogData) error{
	return r.Call(ctx,"GetConductRulesLogDatasForBus",args,reply)
}

//获取行为规范记录统计-分店
func (r *Rules) GetConductRulesLogDatasForShop(ctx context.Context,args *staff.ArgsGetConductRuleslogData,reply *staff.ReplyGetConductRuleslogData) error{
	return r.Call(ctx,"GetConductRulesLogDatasForShop",args,reply)
}

//获取服务表现规则
func (r *Rules) GetServiceRules(ctx context.Context,args *staff.ArgsGetServiceRules,reply *[]staff.ReplyGetServiceRules) error{
	return r.Call(ctx,"GetServiceRules",args,reply)
}

//设置/更新服务表现规范
func (r *Rules) SetServiceRules(ctx context.Context,args *staff.ArgsSetServiceRules,reply *staff.ReplySetServiceRules) error{
	return r.Call(ctx,"SetServiceRules",args,reply)
}

//执行服务表现统计
func (r *Rules) ServiceComment(ctx context.Context,args *staff.ArgsServiceComment,reply *staff.ReplyServiceComment) error{
	return r.Call(ctx,"ServiceComment",args,reply)
}

//获取服务表现规范记录统计-企业/商户
func (r *Rules) GetServiceRulesLogDatasForBus(ctx context.Context,args *staff.ArgsGetServiceRuleslogData,reply *staff.ReplyGetServiceRuleslogData) error{
	return r.Call(ctx,"GetServiceRulesLogDatasForBus",args,reply)
}

//获取服务表现规范记录统计-分店
func (r *Rules) GetServiceRulesLogDatasForShop(ctx context.Context,args *staff.ArgsGetServiceRuleslogData,reply *staff.ReplyGetServiceRuleslogData) error{
	return r.Call(ctx,"GetServiceRulesLogDatasForShop",args,reply)
}


