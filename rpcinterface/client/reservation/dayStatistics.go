package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type DayStatistics struct {
	client.Baseclient
}

func (d *DayStatistics) Init() *DayStatistics {
	d.ServiceName = "rpc_reservation"
	d.ServicePath = "DayStatistics"
	return d
}

//根据条件获取预约统计数据(今日预约完成数，总待处理预约数)
func (d *DayStatistics) GetTotalStatistics(ctx context.Context, args *reservation.ArgsGetTotalStatistics, reply *reservation.ReplyGetTotalStatistics) error {
	return d.Call(ctx, "GetTotalStatistics", args, reply)
}
