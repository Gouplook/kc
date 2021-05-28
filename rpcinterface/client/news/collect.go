package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Collect struct {
	client.Baseclient
}

func (c *Collect) Init() *Collect {
	c.ServiceName = "rpc_news"
	c.ServicePath = "Collect"
	return c
}

// UserCollect 用户收藏
func (c *Collect) UserCollect(ctx context.Context, args *news.ArgsUserCollect, reply *bool) (err error) {
	return c.Call(ctx, "UserCollect", args, reply)
}

// UserCollectList 用户收藏的攻略和探店
func (c *Collect) UserCollectList(ctx context.Context, args *news.ArgsUserCollectList, reply *news.ReplyUserCollectList) (err error) {
	return c.Call(ctx, "UserCollectList", args, reply)
}
