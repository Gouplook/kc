package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type AskQuiz interface {
	SetQuiz(ctx context.Context,args *comment.ArgsSetQuiz,reply *bool)error
	SearchQuiz(ctx context.Context,args *comment.ArgsGetAskQuiz, reply *comment.ReplyGetAskQuiz)error
}
