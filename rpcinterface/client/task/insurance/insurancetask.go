package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

// @author liyang<654516092@qq.com>
// @date  2020/8/3 10:21


type InsuranceTask struct {
	client.Baseclient
}

//初始化
func (i *InsuranceTask) Init() *InsuranceTask {
	i.ServiceName = "rpc_task"
	i.ServicePath = "Insurance/InsuranceTask"
	return i
}

//将预付卡待出单的保单加入到任务列表中
func (i *InsuranceTask) RunInsuranceTask(ctx context.Context,taskId *int,reply *bool) error{
	return i.Call(ctx, "RunInsuranceTask", taskId, reply)
}

