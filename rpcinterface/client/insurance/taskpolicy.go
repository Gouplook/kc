package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/insurance"
)

// @author liyang<654516092@qq.com>
// @date  2020/8/3 11:54

type TaskPolicy struct {
	client.Baseclient
}

func (t *TaskPolicy) Init() *TaskPolicy {
	t.ServiceName = "rpc_insurance"
	t.ServicePath = "TaskPolicy"
	return t
}

//将卡包信息Push到出单任务中
func (t *TaskPolicy) PushRelationIdToTask(ctx context.Context, relationId *int, reply *bool) error {
	return t.Call(ctx, "PushRelationIdToTask", relationId, reply)
}

//执行预付卡出单-正常出单
func (t *TaskPolicy) RunIssuePolicy(ctx context.Context, taskId *int, reply *string) error {
	return t.Call(ctx, "RunIssuePolicy", taskId, reply)
}

//执行预付卡出单-续保出单
func (t *TaskPolicy) RunRenewIssuePolicy(ctx context.Context, taskId *int, reply *string) error {
	return t.Call(ctx, "RunRenewIssuePolicy", taskId, reply)
}

//获取未上传保险公司的消费记录
func (t *TaskPolicy) RunConsumeDataTask(ctx context.Context,args *insurance.ArgsConsumeTask,reply *[]insurance.ReplyConsumeTask) error {
	return t.Call(ctx, "RunConsumeDataTask", args, reply)
}

//更新上传保险公司的消费记录为已跑批
func (t *TaskPolicy) UpdateConsumeDataTask(ctx context.Context,logId *int,reply *bool) error {
	return t.Call(ctx, "UpdateConsumeDataTask", logId, reply)
}

//执行上传消费记录到保险公司
func (t *TaskPolicy) RunConsumeData(ctx context.Context,logId *int,reply *bool) error {
	return t.Call(ctx, "RunConsumeData", logId, reply)
}


