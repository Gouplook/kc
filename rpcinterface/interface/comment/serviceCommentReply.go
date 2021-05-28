package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type ServiceCommentReplyBase struct {
	Id           int
	CommentId    int    //评论ID
	BusId        int    //企业/商户ID
	ShopId       int    //分店ID
	ReplyContent string //回复内容
	Ctime        int64  //回复时间
}

type ArgsAddServiceCommentReply struct {
	common.BsToken
	ServiceCommentReplyBase
}
type ReplyAddServiceCommentReply struct {
	Id int
}

type ServiceCommentReply interface {
	// 添加回复
	AddServiceCommentReply(ctx context.Context,args *ArgsAddServiceCommentReply,reply *ReplyAddServiceCommentReply)error
}
