package admin

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/admin"
)

type Sugger struct {
	client.Baseclient
}



func (s *Sugger)  Init() *Sugger {
	s.ServiceName = "rpc_admin"
	s.ServicePath = "Sugger"
	return s
}

//添加反馈意见
func (s *Sugger) AddSugger( ctx context.Context, args *admin.ArgsAddSugger, reply *bool ) (err error)  {
	return s.Call(ctx, "AddSugger", args, reply)
}