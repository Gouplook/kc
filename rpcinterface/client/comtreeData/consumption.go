/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/3/25 14:32
@Description:

*********************************************/
package comtreeData

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
)

type Consumption struct {
	client.Baseclient
}

func (c *Consumption) Init() *Consumption {
	c.ServiceName = "rpc_comtreedata"
	c.ServicePath = "Consumption"
	return c
}


// 添加 预付卡消费 信息
func (c *Consumption)AddConsumptionRpc(ctx context.Context, consumeLogId *int, reply *bool) error {
	return c.Call(ctx, "AddConsumptionRpc",consumeLogId, reply)
}
//添加  预付卡保险出单 信息
func (c *Consumption)AddInsuranceBillRpc(ctx context.Context, transNo *string, reply *bool)error{
	return c.Call(ctx, "AddInsuranceBillRpc",transNo, reply)
}
// 添加 预付卡交易 信息
func (c *Consumption)AddTransactionsRpc(ctx context.Context, orderSn *string, reply *bool) error{
	return c.Call(ctx, "AddTransactionsRpc",orderSn, reply)
}
// 添加 全国市场规模 信息
func (c *Consumption)AddScaleNationalMarketRpc(ctx context.Context, busId *int , reply *bool) error{
	return c.Call(ctx, "AddScaleNationalMarketRpc",busId, reply)
}
//添加 业务规模 信息
func (c *Consumption)AddScaleBusinessRpc(ctx context.Context, busId *int, reply *bool) error {
	return c.Call(ctx, "AddScaleBusinessRpc", busId, reply)
}
