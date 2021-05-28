package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type CommentTag struct {
	client.Baseclient
}

func (c *CommentTag) Init() *CommentTag {
	c.ServiceName = "rpc_comment"
	c.ServicePath = "CommentTag"
	return c
}

//评价标签
func (c *CommentTag) GetCommentTag(ctx context.Context, args *comment.ArgsGetCommentTag, reply *comment.ReplyGetCommentTag) error {
	return c.Call(ctx, "GetCommentTag", args, reply)
}
