package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type AskQuiz struct {
	client.Baseclient
}

func (a *AskQuiz) Init() *AskQuiz {
	a.ServiceName = "rpc_elastic"
	a.ServicePath = "Comment/AskQuiz"
	return a
}

func (a *AskQuiz) SetQuiz(ctx context.Context, args *comment.ArgsSetQuiz, reply *bool) error {
	return a.Call(ctx, "SetQuiz", args, reply)
}

func (a *AskQuiz) SearchQuiz(ctx context.Context, args *comment.ArgsGetAskQuiz, reply *comment.ReplyGetAskQuiz) error {
	return a.Call(ctx, "SearchQuiz", args, reply)
}
