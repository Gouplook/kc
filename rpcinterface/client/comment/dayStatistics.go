package comment

import (
	"git.900sui.cn/kc/rpcinterface/client"
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type DayStatistics struct {
	client.Baseclient
}

func (d *DayStatistics) Init() *DayStatistics {
	d.ServiceName = "rpc_comment"
	d.ServicePath = "DayStatistics"
	return d
}
// 获取客户评价数据（客户今日完成评价总数目， 待处理总数目）
func (d *DayStatistics) GetCommentTotalStatic(ctx context.Context, args *comment.ArgsGetCommentTotalStatic, reply *comment.ReplyGetCommentTotalStatic) error {

	return d.Call(ctx, "GetCommentTotalStatic", args, reply)
}