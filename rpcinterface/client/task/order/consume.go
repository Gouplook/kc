//消费相关任务入队列
//@author yangzhiwu<578154898@qq.com>
//@date 2020/8/5 14:38
package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

type Consume struct {
	client.Baseclient
}

func (c *Consume) Init() *Consume  {
	c.ServiceName = "rpc_task"
	c.ServicePath = "Order/Consume"
	return c
}

//将消费记录索引主表的主键id加入交换机
func (c *Consume) SetLogRelationId(ctx context.Context, logRelationId *int, reply *bool ) error {
	return c.Call(ctx, "SetLogRelationId", logRelationId, reply)
}
