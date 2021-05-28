package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

/**
 * @className PinganThirdOrder
 * @author liyang<654516092@qq.com>
 * @date 2021/4/27 9:38
 */

type PinganThirdOrder struct {
	client.Baseclient
}

func (p *PinganThirdOrder) Init() *PinganThirdOrder  {
	p.ServiceName = "rpc_task"
	p.ServicePath = "Order/PinganThirdOrder"
	return p
}

//将卡包主表的主键id加入交换机
func (p *PinganThirdOrder) PayThirdOrderSuc(ctx context.Context, id *int, reply *bool ) error {
	return p.Call(ctx, "PayThirdOrderSuc", id, reply)
}