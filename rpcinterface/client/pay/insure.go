package pay

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/pay"
)

type Insure struct {
	client.Baseclient
}

func (i *Insure) Init() *Insure {
	i.ServiceName = "rpc_pay"
	i.ServicePath = "Insure"
	return i
}

//获取平台保险和续在不同支付渠道的开户账号
func (i *Insure) GetInsureAcct(ctx context.Context, args *pay.ArgsGetInsureAcct, reply *pay.ReplyGetInsureAcct) error {
	return i.Call(ctx, "GetInsureInfo", args, reply)
}
