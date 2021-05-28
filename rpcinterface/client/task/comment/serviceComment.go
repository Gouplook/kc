package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

type ServiceComment struct {
	client.Baseclient
}

func (s *ServiceComment) Init() *ServiceComment {
	s.ServiceName = "rpc_task"
	s.ServicePath = "Comment/ServiceComment"
	return s
}

//设置服务评价ID到交换机中
func (s *ServiceComment) SetServiceCommentId(ctx context.Context, serviceCommentId *int, reply *bool) error {
	return s.Call(ctx, "SetServiceCommentId", serviceCommentId, reply)
}
