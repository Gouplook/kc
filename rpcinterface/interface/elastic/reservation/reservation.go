package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type ArgsSetReservationInfo struct {
	Id int
}

type Reservation interface {
	// 添加/更新预约数据到ES
	SetReservationInfo(ctx context.Context, args *ArgsSetReservationInfo, reply *bool) error
	//	预约检索
	SearchReservation(ctx context.Context, args *reservation.GetReservationRecordListParams, reply *reservation.GetReservationRecordListReplies) error
	//查询预约健康信息
	SearchHealth(ctx context.Context, args *reservation.ArgsHealthGet, reply *reservation.ReplyReservationHealth) error
}
