package comment

import "context"

type ServiceComment interface {
	// 设置服务评价ID到交换机中
	SetServiceCommentId(ctx context.Context,serviceCommentId *int,reply *bool)error
}
