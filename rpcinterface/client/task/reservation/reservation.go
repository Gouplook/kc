package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/reservation"
)

type Reservation struct {
	client.Baseclient
}

func (r *Reservation) Init() *Reservation {
	r.ServiceName = "rpc_task"
	r.ServicePath = "Reservation/Reservation"
	return r
}

// 添加预约记录信息
func (r *Reservation) ReservationAdd(ctx context.Context, args *reservation.ArgsReservationAdd, reply *bool) error {
	return r.Call(ctx, "ReservationAdd", args, reply)
}

//预约的项目消费完成
func (r *Reservation) ReservationComplete(ctx context.Context, reservationId *int, reply *bool) error  {
	return r.Call(ctx, "ReservationComplete", reservationId, reply)
}

//预约到店超时将id加入交换机
func (r *Reservation)TimeOutReservationId(ctx context.Context,args *reservation.ArgsTimeOutReservationId,reply *bool)error{
	return r.Call(ctx, "TimeOutReservationId", args, reply)
}