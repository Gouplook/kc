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
	a.ServiceName = "rpc_comment"
	a.ServicePath = "AskQuiz"
	return a
}

//用户发表提问
func (a *AskQuiz) AddAskQuiz(ctx context.Context, args *comment.ArgsAddAskQuiz, reply *comment.ReplyAddAskId) error {
	return a.Call(ctx, "AddAskQuiz", args, reply)
}

//查看提问列表
func (a *AskQuiz) GetAskQuiz(ctx context.Context, args *comment.ArgsGetAskQuiz, reply *comment.ReplyGetAskQuiz) error {
	return a.Call(ctx, "GetAskQuiz", args, reply)
}

//内部处理提问列表数据
func (a *AskQuiz) DisposeAskQuizRpc(ctx context.Context, args *comment.ArgsDisposeAskQuizRpc, reply *comment.ReplyDisposeAskQuizRpc) error {
	return a.Call(ctx, "DisposeAskQuizRpc", args, reply)
}

//查看提问详情(包含三条默认的提问答案)
func (a *AskQuiz) GetAskQuizInfo(ctx context.Context, args *comment.ArgsGetAskQuizInfo, reply *comment.ReplyGetAskQuizInfo) error {
	return a.Call(ctx, "GetAskQuizInfo", args, reply)
}

//点赞提问
func (a *AskQuiz) PraiseQuiz(ctx context.Context, args *comment.ArgsPraiseQuiz, reply *bool) error {
	return a.Call(ctx, "PraiseQuiz", args, reply)
}

//收藏提问
func (a *AskQuiz) CollectQuiz(ctx context.Context, args *comment.ArgsCollectQuiz, reply *bool) error {
	return a.Call(ctx, "CollectQuiz", args, reply)
}

//我收藏的提问
func (a *AskQuiz) GetMyCollectQuiz(ctx context.Context, args *comment.ArgsGetMyCollectQuiz, reply *comment.ReplyGetMyCollectQuiz) error {
	return a.Call(ctx, "GetMyCollectQuiz", args, reply)
}

//获取问答关键字
func (a *AskQuiz) GetQuizKeywords(ctx context.Context, args *comment.ArgsGetQuizKeywords, reply *comment.ReplyGetQuizKeywords) error {
	return a.Call(ctx, "GetQuizKeywords", args, reply)
}

//我的回复(相当于提问的答案)
func (a *AskQuiz) GetMyQuizAnswer(ctx context.Context, args *comment.ArgsGetMyQuizAnswer, reply *comment.ReplyGetMyQuizAnswer) error {
	return a.Call(ctx, "GetMyQuizAnswer", args, reply)
}

//我的提问
func (a *AskQuiz) GetMyQuiz(ctx context.Context, args *comment.ArgsGetMyQuiz, reply *comment.ReplyGetMyQuiz) error {
	return a.Call(ctx, "GetMyQuiz", args, reply)
}

//提问数据-rpc
func (a *AskQuiz) GetQuizListByIdsRpc(ctx context.Context, quizIds *[]int, reply *[]map[string]interface{}) error {
	return a.Call(ctx, "GetQuizListByIdsRpc", quizIds, reply)
}

//根据id获取简单的提问数据
func (a *AskQuiz) GetSimpleQuizById(ctx context.Context, quizId *int, reply *comment.AskQuizBase) error {
	return a.Call(ctx, "GetSimpleQuizById", quizId, reply)
}

//==============================================后台管理接口============================================

//获取提问列表
func (a *AskQuiz) AdminGetQuizList(ctx context.Context, args *comment.ArgsAdminGetQuizList, reply *comment.ReplyAdminGetQuizList) error {
	return a.Call(ctx, "AdminGetQuizList", args, reply)
}

//添加/修改提问
func (a *AskQuiz) AdminAddOrUpdateQuiz(ctx context.Context, args *comment.ArgsAdminAddOrUpdateQuiz, reply *comment.ReplyAddAskId) error {
	return a.Call(ctx, "AdminAddOrUpdateQuiz", args, reply)
}

//审核提问
func (a *AskQuiz) AdminAuditQuiz(ctx context.Context, args *comment.ArgsAdminAuditQuiz, reply *bool) error {
	return a.Call(ctx, "AdminAuditQuiz", args, reply)
}

//提问详情
func (a *AskQuiz) AdminGetQuizInfo(ctx context.Context, args *comment.ArgsAdminGetQuizInfo, reply *comment.ReplyAdminGetQuizInfo) error {
	return a.Call(ctx, "AdminGetQuizInfo", args, reply)
}

//添加提问关键字
func (a *AskQuiz) AdminAddKeyword(ctx context.Context, args *comment.ArgsAdminAddOrUpdateKeyword, reply *bool) error {
	return a.Call(ctx, "AdminAddKeyword", args, reply)
}

//修改提问关键字
func (a *AskQuiz) AdminUpdateKeyword(ctx context.Context, args *comment.ArgsAdminAddOrUpdateKeyword, reply *bool) error {
	return a.Call(ctx, "AdminUpdateKeyword", args, reply)
}

//关键字列表
func (a *AskQuiz) AdminGetKeywordLists(ctx context.Context, args *comment.ArgsGetQuizKeywords, reply *comment.ReplyAdminGetKeywordLists) error {
	return a.Call(ctx, "AdminGetKeywordLists", args, reply)
}

//删除提问关键字
func (a *AskQuiz) AdminDelKeyword(ctx context.Context, args *comment.ArgsAdminDelKeyword, reply *bool) error {
	return a.Call(ctx, "AdminDelKeyword", args, reply)
}
