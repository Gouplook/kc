package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type ServiceCommentReply struct {
	client.Baseclient
}

func (s *ServiceCommentReply)Init() *ServiceCommentReply {
	s.ServiceName="rpc_comment"
	s.ServicePath="ServiceCommentReply"
	return s
}

// 添加回复
func (s *ServiceCommentReply)AddServiceCommentReply(ctx context.Context,args *comment.ArgsAddServiceCommentReply,reply *comment.ReplyAddServiceCommentReply)error{
	return s.Call(ctx,"AddServiceCommentReply",args,reply)
}