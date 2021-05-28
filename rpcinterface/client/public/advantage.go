package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

type Advantage struct {
	client.Baseclient
}

func (a *Advantage) Init() *Advantage {
	a.ServiceName = "rpc_public"
	a.ServicePath = "Advantage"
	return a
}

//获取所有的优势标签
func (a *Advantage) GetAdvantageList(ctx context.Context, args *public.EmptyParams, reply *public.ReplyGetAdvantageList) error {
	return a.Call(ctx, "GetAdvantageList", args, reply)
}

//获取指定的优势标签-RPC
func (a *Advantage) GetAdvantageByIds(ctx context.Context, args *public.ArgsGetAdvantageByIds, reply *public.ReplyGetAdvantageByIds) error {
	return a.Call(ctx, "GetAdvantageByIds", args, reply)
}
