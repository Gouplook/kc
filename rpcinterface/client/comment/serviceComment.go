package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type ServiceComment struct {
	client.Baseclient
}

func (s *ServiceComment) Init() *ServiceComment {
	s.ServiceName = "rpc_comment"
	s.ServicePath = "ServiceComment"
	return s
}


//根据Saas门店id查询用户评论
func (s *ServiceComment) GetCommentBySaas(ctx context.Context, args *comment.ArgsComment, reply *comment.ReplyComment) error {
	return s.Call(ctx,"GetCommentBySaas",args,reply)
}

//根据门店id查询用户评论
func (s *ServiceComment) GetCommentByShopId(ctx context.Context, args *comment.ArgsComment, reply *comment.ReplyComment) error {
	return s.Call(ctx,"GetCommentByShopId",args,reply)
}

//根据单项目id查询用户评论
func (s *ServiceComment) GetCommentBySingleId(ctx context.Context, args *comment.ArgsComment, reply *comment.ReplyComment) error {
	return s.Call(ctx,"GetCommentBySingleId",args,reply)
}



//用户评论
func (s *ServiceComment) AddServiceComment(ctx context.Context, args *comment.ArgsAddServiceComment, reply *comment.ReplyAddServiceComment) error {
	return s.Call(ctx, "AddServiceComment", args, reply)
}

//根据评价ID获取评价内容
func (s *ServiceComment) GetServiceCommentByID(ctx context.Context, serviceCommentId *int, reply *comment.UserComment) error {
	return s.Call(ctx, "GetServiceCommentByID", serviceCommentId, reply)
}

//根据评价ID获取评价得分
func (s *ServiceComment) GetServiceCommentScoreById(ctx context.Context,serviceCommentId *int ,reply *comment.ReplyGetServiceCommentScore)error{
	return s.Call(ctx, "GetServiceCommentByID", serviceCommentId, reply)
}

//用户点赞评价
func (s *ServiceComment) CommentPraise(ctx context.Context,args *comment.ArgsCommentPraise,reply *bool)error{
	return s.Call(ctx, "CommentPraise", args, reply)
}
// 获取消费者评分
func(s *ServiceComment) GetCousumerEvalutaion(ctx context.Context, args *comment.ArgsCousumerEvalutaion, reply *comment.ReplyCousumerEvalutaion) error{
	return s.Call(ctx,"GetCousumerEvalutaion",args, reply)
}