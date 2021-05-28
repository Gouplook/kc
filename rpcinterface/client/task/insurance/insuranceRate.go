/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/18 下午1:00

*******************************************/
package insurance

import (
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/insurance"
	"context"
)

type InsuranceRate struct {
	client.Baseclient
}

func (i *InsuranceRate) Init() *InsuranceRate{
	i.ServicePath = "Insurance/InsuranceRate"
	i.ServiceName = "rpc_task"
	return i
}

// 投保率
func(i *InsuranceRate) SetInsuranceRate(ctx context.Context, args *insurance.ArgsInsuranceRate, reply *bool) error{
	return i.Call(ctx, "SetInsuranceRate",args, reply)
}

