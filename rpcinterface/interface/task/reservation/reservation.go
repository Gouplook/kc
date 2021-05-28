package reservation

import "context"

type ArgsReservationAdd struct {
	ReservationID int
}

type ArgsTimeOutReservationId struct {
	ReservationID int
	ExpireTimeOut int // 超时时间,单位毫秒
}

type Reservation interface {
	// 添加预约信息
	ReservationAdd(ctx context.Context,args *ArgsReservationAdd,reply *bool)error
	//预约的项目消费完成
	ReservationComplete(ctx context.Context, reservationId *int, reply *bool) error
	//预约到店超时将id加入交换机
	TimeOutReservationId(ctx context.Context,args *ArgsTimeOutReservationId,reply *bool)error
}