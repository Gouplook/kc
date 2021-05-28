package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Comments struct {
	client.Baseclient
}

func (c *Comments) Init() *Comments {
	c.ServicePath = "Comments"
	c.ServiceName = "rpc_news"
	return c
}

// AddCommnets 添加评论
func (c *Comments) AddCommnets(ctx context.Context, args *news.ArgsAddCommnets, reply *news.ReplyAddComments) error {
	return c.Call(ctx, "AddCommnets", args, reply)
}

// UserCommentList 前台评论列表
func (c *Comments) UserCommentList(ctx context.Context, args *news.ArgsUserCommentList, reply *news.ReplyUserCommentList) error {
	return c.Call(ctx, "UserCommentList", args, reply)
}

// ReplyCommentList 前台评论回复列表
func (c *Comments) ReplyCommentList(ctx context.Context, args *news.ArgsReplyCommentList, reply *news.ReplyCommentList) error {
	return c.Call(ctx, "ReplyCommentList", args, reply)
}

// AddReplyComments 添加评论回复
func (c *Comments) AddReplyComments(ctx context.Context, args *news.ArgsAddReplyComment, reply *news.ReplyAddReplyComment) error {
	return c.Call(ctx, "AddReplyComments", args, reply)
}

// AdminCommentsList 后台攻略/探店评论列表
func (c *Comments) AdminCommentsList(ctx context.Context, args *news.ArgsAdminCommentsList, reply *news.ReplyAdminCommentsList) error {
	return c.Call(ctx, "AdminCommentsList", args, reply)
}

// AuditComment 后台审核用户评论
func (c *Comments) AuditComment(ctx context.Context, args *news.ArgsAuditComment, reply *bool) error {
	return c.Call(ctx, "AuditComment", args, reply)
}

//　AdminDelReplyComment 管理员删除评论回复
func (c *Comments) AdminDelReplyComment(ctx context.Context, args *news.ArgsAdminDelReplyComment, reply *bool) error {
	return c.Call(ctx, "AdminDelReplyComment", args, reply)
}

//　AdminReplyCommentList　后台评论回复列表
func (c *Comments) AdminReplyCommentList(ctx context.Context, args *news.ArgsAdminReplyCommentList, reply *news.ReplyAdminReplyCommentLsit) error {
	return c.Call(ctx, "AdminReplyCommentList", args, reply)
}
