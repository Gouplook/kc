package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/elastic/reservation"
	reservation2 "git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type Reservation struct {
	client.Baseclient
}

func (r *Reservation) Init() *Reservation {
	r.ServiceName = "rpc_elastic"
	r.ServicePath = "Reservation/Reservation"
	return r
}

//添加/更新预约数据到ES
func (r *Reservation) SetReservationInfo(ctx context.Context, args *reservation.ArgsSetReservationInfo, reply *bool) error {
	return r.Call(ctx, "SetReservationInfo", args, reply)
}
//预约检索
func (r *Reservation) SearchReservation(ctx context.Context, args *reservation2.GetReservationRecordListParams, reply *reservation2.GetReservationRecordListReplies) error {
	return r.Call(ctx, "SearchReservation", args, reply)
}

//查询预约健康信息
func (r *Reservation) SearchHealth(ctx context.Context, args *reservation2.ArgsHealthGet, reply *reservation2.ReplyReservationHealth) error {
	return r.Call(ctx, "SearchHealth", args, reply)
}