package financial

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/financial"
)

type Gov struct {
	client.Baseclient
}

func (g *Gov) Init() *Gov {
	g.ServiceName = "rpc_financial"
	g.ServicePath = "Gov"
	return g
}

//同步监管平台的银行，保险公司，第三方支付
func (g *Gov) AnsyGovInfo(ctx context.Context, args *financial.ArgsAnsyGovInfo, reply *bool) error{
	return g.Call(ctx, "AnsyGovInfo", args, reply)
}
