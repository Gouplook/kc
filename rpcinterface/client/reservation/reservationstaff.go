package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type ReservationStaff struct {
	client.Baseclient
}

func (r *ReservationStaff) Init() *ReservationStaff {
	r.ServiceName = "rpc_reservation"
	r.ServicePath = "ReservationStaff"
	return r
}

//获取技师被预约的时间
func (r *ReservationStaff) GetReservationStaff(ctx context.Context, args *reservation.ArgsReservationStaff, reply *map[int][]reservation.ReplyReservationStaff) (err error) {
	err = r.Call(ctx, "GetReservationStaff", args, reply)
	return
}