/**
 * @Author: YangYun
 * @Date: 2020/8/3 10:22
 */
package pay

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

type Pay struct {
	client.Baseclient
}

func (p *Pay) Init() *Pay {
	p.ServiceName = "rpc_task"
	p.ServicePath = "Pay/Pay"
	return p
}

// 订单支付成功
func (p *Pay) PaySuc(ctx context.Context, orderSn *string, reply *bool) error {
	return p.Call(ctx, "PaySuc", orderSn, reply)
}


//代付成功
func (p *Pay) AgentSuc(ctx context.Context, clearId *string, reply *bool) error {
	return p.Call(ctx, "AgentSuc", clearId, reply)
}