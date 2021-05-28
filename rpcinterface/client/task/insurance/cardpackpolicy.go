/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/1 上午10:19

*******************************************/
package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/insurance"
)

type CardPackPolicy struct {
	client.Baseclient
}

func (c *CardPackPolicy) Init() *CardPackPolicy {
	c.ServiceName = "rpc_task"
	c.ServicePath = "Insurance/CardPackPolicy"
	return c
}

// 保单出单成功
func (c *CardPackPolicy) CardPackPolicySuc(ctx context.Context, args *insurance.ArgsCardPackPolicySuc, reply *bool) (err error) {
	return c.Call(ctx, "CardPackPolicySuc", args, reply)
}
