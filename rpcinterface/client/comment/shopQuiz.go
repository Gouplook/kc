package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type ShopQuiz struct {
	client.Baseclient
}
//初始化
func (s *ShopQuiz)Init() *ShopQuiz {
	s.ServiceName = "rpc_comment"
	s.ServicePath = "ShopQuiz"
	return s
}

//用户添加提问
func(s *ShopQuiz) AddQuiz(ctx context.Context, args *comment.ArgsAddQuiz,reply *int) error {
	return s.Call(ctx,"AddQuiz",args,reply)
}
//用户回答提问
func(s *ShopQuiz) AddReply(ctx context.Context, args *comment.ArgsAddReply, reply *int) error {
	return s.Call(ctx,"AddReply",args,reply)
}

//获取门店问答列表
func(s *ShopQuiz) GetQuizList(ctx context.Context, args *comment.ArgsGetQuizList, reply *comment.ReplyGetQuizList) error {
	return s.Call(ctx,"GetQuizList",args,reply)
}
//获取详细问答
func(s *ShopQuiz) GetQuizDetails(ctx context.Context, args *comment.ArgsGetQuizDetails, reply *comment.ReplyGetQuizDetails) error {
	return s.Call(ctx,"GetQuizDetails",args,reply)
}

//门店回答点赞
func(s *ShopQuiz) GiveALike(ctx context.Context, args *comment.ArgsGiveALick, reply *bool) error {
	return s.Call(ctx,"GiveALike",args,reply)
}