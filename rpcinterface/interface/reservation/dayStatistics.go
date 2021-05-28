package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//根据条件获取预约统计入参
type ArgsGetTotalStatistics struct {
	common.BsToken
}

//根据条件获取预约统计出参
type ReplyGetTotalStatistics struct {
	TodayCompleteTotalNum int //今日预约完成总数
	PendingTotalNum int //待处理总数
}

type DayStatistics interface {
	//根据条件获取预约统计数据(今日预约完成数，总待处理预约数)
	GetTotalStatistics(ctx context.Context, args *ArgsGetTotalStatistics, reply *ReplyGetTotalStatistics) error
}
