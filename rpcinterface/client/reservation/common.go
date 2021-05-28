package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type Common struct {
	client.Baseclient
}

func (c *Common) Init() *Common {
	c.ServiceName = "rpc_reservation"
	c.ServicePath = "Common"
	return c
}

//添加顾客健康信息
func (c *Common) AddHealth(ctx context.Context, args *reservation.ArgsHealthAdd, reply *int) error {
	return c.Call(ctx, "AddHealth", args, reply)
}

//查询顾客健康信息
func (c *Common) GetHealth(ctx context.Context, args *reservation.ArgsHealthGet, reply *reservation.ReplyHealth) error {
	return c.Call(ctx, "GetHealth", args, reply)
}

//根据服务id和手艺人id查询时间
func (c *Common) GetReservationTime(ctx context.Context, args *reservation.ArgsTimeGet, reply *reservation.ReplyTime) error {
	return c.Call(ctx, "GetReservationTime", args, reply)
}

//添加顾客到店健康信息
func (c *Common) AddHealthArrive(ctx context.Context, args *reservation.ArgsHealthAdd, reply *int) error {
	return c.Call(ctx, "AddHealthArrive", args, reply)
}

//添加顾客离店健康信息
func (c *Common) AddHealthLeave(ctx context.Context, args *reservation.ArgsHealthAdd, reply *int) error {
	return c.Call(ctx, "AddHealthLeave", args, reply)
}