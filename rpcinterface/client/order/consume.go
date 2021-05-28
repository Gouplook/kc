/*
 * @Author: your name
 * @Date: 2021-05-19 17:33:27
 * @LastEditTime: 2021-05-19 17:44:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \rpcOpen\rpcinterface\client\order\consume.go
 */
package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
	v1 "git.900sui.cn/kc/rpcinterface/interface/order/openPlatform/v1"
)

type Consume struct {
	client.Baseclient
}

func (c *Consume) Init() *Consume {
	c.ServiceName = "rpc_order"
	c.ServicePath = "Consume"
	return c
}

//确认消费
func (c *Consume) ConsumeService(ctx context.Context, args *order.ArgsConsumeService, reply *bool) error {
	return c.Call(ctx, "ConsumeService", args, reply)
}

//确认消费修改卡包释放金额【rpc】
func (c *Consume) ChangeRelaseAmount(ctx context.Context, args *order.ArgsChangeRelaseAmount, reply *bool) error {
	return c.Call(ctx, "ChangeRelaseAmount", args, reply)
}

//系统自动确认消费限时限次卡包
func (c *Consume) AutoConfirmHNCard(ctx context.Context, dateYMD *int, reply *bool) (err error) {
	return c.Call(ctx, "AutoConfirmHNCard", dateYMD, reply)
}

//系统自动确认消费限时卡包
func (c *Consume) AutoConfirmHCard(ctx context.Context, dateYMD *int, reply *bool) (err error) {
	return c.Call(ctx, "AutoConfirmHCard", dateYMD, reply)
}

//开放平台-v1-充值卡-确认消费
func (c *Consume) OPV1RcardConsumeSrv(ctx context.Context, args *v1.OPV1RcardConsumeSrvRequest, reply *bool) (err error) {
	return c.Call(ctx, "OPV1RcardConsumeSrv", args, reply)
}
