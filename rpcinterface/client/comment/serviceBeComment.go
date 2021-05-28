package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type ServiceBeComment struct {
	client.Baseclient
}

func (s *ServiceBeComment) Init() *ServiceBeComment {
	s.ServiceName = "rpc_comment"
	s.ServicePath = "ServiceBeComment"
	return s
}

//添加服务待评论-rpc
func (s *ServiceBeComment) AddServiceBeCommentRpc(ctx context.Context, args *comment.ArgsAddServiceBeComment, reply *comment.ReplyAddServiceBeComment) error {
	return s.Call(ctx, "AddServiceBeCommentRpc", args, reply)
}

//根据状态获取待评价/已评价的服务列表
func (s *ServiceBeComment) GetServiceBeComments(ctx context.Context, args *comment.ArgsGetServiceBeComments, reply *comment.ReplyGetServiceBeComments) error {
	return s.Call(ctx, "GetServiceBeComments", args, reply)
}
