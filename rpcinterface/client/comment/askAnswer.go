package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type AskAnswer struct {
	client.Baseclient
}

func (a *AskAnswer) Init() *AskAnswer {
	a.ServiceName = "rpc_comment"
	a.ServicePath = "AskAnswer"
	return a
}

//用户回答提问
func (a *AskAnswer) AddAskAnswer(ctx context.Context, args *comment.ArgsAddAskAnswer, reply *comment.ReplyAddAskId) error {
	return a.Call(ctx, "AddAskAnswer", args, reply)
}

//查看指定提问下的全部回答(包含三条最新的答案评论)
func (a *AskAnswer) GetAllAskAnswer(ctx context.Context, args *comment.ArgsGetAllAskAnswer, reply *comment.ReplyGetAllAskAnswer) error {
	return a.Call(ctx, "GetAllAskAnswer", args, reply)
}

//查看答案详情
func (a *AskAnswer) GetAnswerComments(ctx context.Context, args *comment.ArgsGetAnswerComments, reply *comment.ReplyGetAnswerComments) error {
	return a.Call(ctx, "GetAnswerComments", args, reply)
}

//点赞答案
func (a *AskAnswer) PraiseAnswer(ctx context.Context, args *comment.ArgsPraiseAnswer, reply *bool) error {
	return a.Call(ctx, "PraiseAnswer", args, reply)
}

//发表答案评论
func (a *AskAnswer) AddAnswerComment(ctx context.Context, args *comment.ArgsAddAnswerComment, reply *bool) error {
	return a.Call(ctx, "AddAnswerComment", args, reply)
}

//==============================================后台管理接口============================================

//后台答案列表接口出参
func (a *AskAnswer) AdminGetAnswerList(ctx context.Context, args *comment.AdminAuditListBase, reply *comment.ReplyAdminGetAnswerList) error {
	return a.Call(ctx, "AdminGetAnswerList", args, reply)
}

//审核答案/评论入参
func (a *AskAnswer) AdminAuditAnswer(ctx context.Context, args *comment.ArgsAdminAuditAnswerOrComment, reply *bool) error {
	return a.Call(ctx, "AdminAuditAnswer", args, reply)
}

//后台评论列表出参
func (a *AskAnswer) AdminGetCommentList(ctx context.Context, args *comment.AdminAuditListBase, reply *comment.ReplyAdminGetCommentList) error {
	return a.Call(ctx, "AdminGetCommentList", args, reply)
}
