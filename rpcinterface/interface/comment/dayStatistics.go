package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type ArgsGetCommentTotalStatic struct {
	common.BsToken
}

type ReplyGetCommentTotalStatic struct {
	TodayCommentTotalNum int // 今日客户评价总数目
	PendingTotalNum      int // 待处理总数
}

type DayStatistics interface {
	// 获取客户评价数据（客户今日完成评价总数目， 待处理总数目）
	GetCommentTotalStatic(ctx context.Context,  args *ArgsGetCommentTotalStatic, reply *ReplyGetCommentTotalStatic) error
}
