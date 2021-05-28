package task

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task"
)

type Public struct {
	client.Baseclient
}

func (p *Public)Init() *Public {
	p.ServiceName = "rpc_task"
	p.ServicePath = "Public"
	return p
}

//发短信
func (p *Public) SendSms( ctx context.Context, args *task.SendSmsParams, reply *bool ) error {
	return p.Call( ctx, "SendSms", args, reply )
}

//发邮件
func (p *Public) SendMail( ctx context.Context, args *task.SendMailParams, reply *bool ) error {
	return p.Call( ctx, "SendMail", args, reply )
}

//新增商品或卡项，风控统计
func (p *Public )AddDelGoods( ctx context.Context, args *task.ArgsAddDelGoods, reply *bool )error{
	return p.Call( ctx, "AddDelGoods", args, reply )
}
