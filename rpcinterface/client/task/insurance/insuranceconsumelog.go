package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

/**
 * 实现保险公司消费记录上传
 * @className insuranceconsumelog
 * @author liyang<654516092@qq.com>
 * @date 2020/9/10 17:21
 */
type InsuranceConsumeLog struct {
	client.Baseclient
}

//初始化
func (i *InsuranceConsumeLog) Init() *InsuranceConsumeLog {
	i.ServiceName = "rpc_task"
	i.ServicePath = "Insurance/InsuranceConsumeLog"
	return i
}

//实现保险公司消费记录上传
func (i *InsuranceConsumeLog) RunInsuranceConsumeLogTask(ctx context.Context,taskId *int,reply *bool) error{
	return i.Call(ctx, "RunInsuranceConsumeLogTask", taskId, reply)
}
