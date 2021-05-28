package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Exploration struct {
	client.Baseclient
}

func (e *Exploration) Init() *Exploration {
	e.ServiceName = "rpc_news"
	e.ServicePath = "Exploration"
	return e
}

// 后台添加探店
func (e *Exploration) AddExploration(ctx context.Context, req *news.AddExplorationReq, resp *news.ExplorationID) (err error) {
	return e.Call(ctx, "AddExploration", req, resp)
}

// 后台更新探店
func (e *Exploration) UpdateExploration(ctx context.Context, req *news.ExplorationReq, resp *news.Empty) (err error) {
	return e.Call(ctx, "UpdateExploration", req, resp)
}

// 后台探店详情
func (e *Exploration) ExplorationInfo(ctx context.Context, req *news.ExplorationID, resp *news.ExplorationResp) (err error) {
	return e.Call(ctx, "ExplorationInfo", req, resp)
}

func (e *Exploration) ExplorationInfo2(ctx context.Context, req *news.ExplorationID, resp *news.ExplorationResp) (err error) {
	return e.Call(ctx, "ExplorationInfo2", req, resp)
}

// 后台探店列表
func (e *Exploration) ExplorationList(ctx context.Context, req *news.ListReq, resp *news.ExplorationList) (err error) {
	return e.Call(ctx, "ExplorationList", req, resp)
}

// ExplorationListByAuthId 前台获取探员下的探店
func (e *Exploration) ExplorationListByAuthId(ctx context.Context, args *news.ArgsExplorationListByAuthId, reply *news.ReplyUserExplorationList) (err error) {
	return e.Call(ctx, "ExplorationListByAuthId", args, reply)
}

// UserExplorationInfo 前台探店详情
func (e *Exploration) UserExplorationInfo(ctx context.Context, args *news.ArgsUserExplorationInfo, reply *news.ReplyExplorationInfo) (err error) {
	return e.Call(ctx, "UserExplorationInfo", args, reply)
}

// UserExplorationList 前台探店列表
func (e *Exploration) UserExplorationList(ctx context.Context, args *news.ArgsUserExplorationList, reply *news.ReplyUserExplorationList) (err error) {
	return e.Call(ctx, "UserExplorationList", args, reply)
}

// ExplorationRecommend 前台探店推荐
func (e *Exploration) ExplorationRecommend(ctx context.Context, args *news.ArgsExplorationRecommend, reply *news.ReplyUserExplorationList) (err error) {
	return e.Call(ctx, "ExplorationRecommend", args, reply)
}
