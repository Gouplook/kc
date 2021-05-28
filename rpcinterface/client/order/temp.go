package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

type Temp struct {
	client.Baseclient
}

//初始化
func (t *Temp) Init() *Temp {
	t.ServiceName = "rpc_order"
	t.ServicePath = "Temp"
	return t
}

//添加一条 挂单
func(t *Temp) AddTemp(ctx context.Context,args *order.ArgsAddTemp,reply *int) error {
	return t.Call(ctx,"AddTemp",args,reply)
}

//延迟队列处理挂单
func(t *Temp) MqCancelTemp(ctx context.Context,id int,reply *bool) error {
	return t.Call(ctx,"MqCancelTemp",id,reply)
}

//取单
func(t *Temp) GetOneTemp (ctx context.Context,args *order.ArgsGetOneTemp,reply *order.ReplyGetOneTemp) error {
	return t.Call(ctx,"GetOneTemp",args,reply)
}

//获取挂单列表
func(t *Temp) GetTempList(ctx context.Context,args *order.ArgsGetTempList, reply *order.ReplyGetTempList) error {
	return t.Call(ctx,"GetTempList",args,reply)
}

//取消一条挂单
func(t *Temp) CancelTemp(ctx context.Context, args *order.ArgsCancelTemp, reply *bool) error {
	return t.Call(ctx,"CancelTemp",args,reply)
}